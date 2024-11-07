package formdata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/kolesa-team/go-webp/encoder"
	cwebp "github.com/kolesa-team/go-webp/webp"
	common "github.com/rus-sharafiev/go-rest-common"
	"github.com/rus-sharafiev/go-rest-common/exception"
	"golang.org/x/image/draw"
	"golang.org/x/image/webp"
)

type uploadErrors struct {
	UploadErrors []uploadError `json:"uploadErrors"`
}

type uploadError struct {
	Filename string `json:"filename"`
	Error    string `json:"error"`
}

// Form Data interceptor for custom JSON to Form Data converter.
// The source code of the converter can be found at
// https://github.com/rus-sharafiev/fetch-api/blob/master/src/fetch-api.ts#L180
func Interceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentPath := strings.Split(strings.Trim(r.URL.Path, "/"), "/")[0]
		config := common.Config.ConverterConfig

		// Check whether to use the converter
		if config == nil {
			next.ServeHTTP(w, r)
			return
		}

		// Check for whitea and black lists
		if wl := config.Whitelist; wl != nil && !inList(*wl, r) {
			next.ServeHTTP(w, r)
			return
		} else if bl := config.Blacklist; wl == nil && bl != nil && inList(*bl, r) {
			next.ServeHTTP(w, r)
			return
		}

		// Select the upload dir
		uploadPath := "uploads"

		if config.UploadPath != nil {
			uploadPath = *config.UploadPath
		}

		if config.UploadPathByRoute != nil {
			pathByRouteMap := *config.UploadPathByRoute
			if pathByRoute := pathByRouteMap[currentPath]; len(pathByRoute) > 0 {
				uploadPath = pathByRoute
			}
		}

		// Use subfolders
		fullPath := uploadPath
		if config.UploadPathPrefix != nil {
			fullPath = path.Join(*config.UploadPathPrefix, fullPath)
		}

		if config.UseUserSubfolder != nil && *config.UseUserSubfolder {

			userDir := ""
			if userID := r.Header.Get("userID"); len(userID) != 0 {
				userDir = userID
			} else {
				exception.UnauthorizedError(w, fmt.Errorf("only authorized users can save files"))
				return
			}

			fullPath = path.Join(fullPath, userDir)
		}

		// Check if the dir exists
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			if err := os.Mkdir(fullPath, 0755); err != nil {
				exception.InternalServerError(w, err)
				return
			}
		}

		// Check whether the request contains multipart/form-data
		if mr, err := r.MultipartReader(); err == nil {

			// Read the form
			form, err := mr.ReadForm(32 << 20)
			if err != nil {
				exception.InternalServerError(w, err)
				return
			}

			resultChan := make(chan []string)
			errChan := make(chan uploadError)
			fileList := []string{}

			// Iterate over the range of multipart file fields
			for name, values := range form.File {
				for index, fileHeader := range values {
					fileList = append(fileList, name+"-"+strconv.Itoa(index))
					fileHeader := fileHeader
					name := name

					go func() {

						// Get file
						file, err := fileHeader.Open()
						if err != nil {
							errChan <- uploadError{Filename: fileHeader.Filename, Error: err.Error()}
							return
						}

						// Generate unique file name
						id, err := uuid.NewRandom()
						if err != nil {
							errChan <- uploadError{Filename: fileHeader.Filename, Error: err.Error()}
							return
						}

						fileName := path.Join(fullPath, id.String()+filepath.Ext(fileHeader.Filename))

						// Write
						outFile, err := os.Create(fileName)
						if err != nil {
							errChan <- uploadError{Filename: fileHeader.Filename, Error: err.Error()}
							return
						}
						defer outFile.Close()

						if config.OptimizeImages == nil {

							// Write original file
							_, err = io.Copy(outFile, file)
							if err != nil {
								errChan <- uploadError{Filename: fileHeader.Filename, Error: err.Error()}
								return
							}

							// Create optimized images
						} else {

							// Write original file and fill the buffer for image.Image
							var fileBuf bytes.Buffer
							fileReader := io.TeeReader(file, &fileBuf)
							_, err = io.Copy(outFile, fileReader)
							if err != nil {
								errChan <- uploadError{Filename: fileHeader.Filename, Error: err.Error()}
								return
							}

							routesMap := *config.OptimizeImages
							var sizesArr *[]int
							for route, v := range routesMap {
								if route == currentPath {
									sizesArr = &v
								}
							}

							contentType := fileHeader.Header.Get("Content-Type")
							allowedTypes := []string{"image/png", "image/jpeg", "image/webp"}
							isImage := slices.Contains(allowedTypes, contentType)

							if sizesArr != nil && isImage {
								isWebP := false
								var img image.Image
								switch contentType {
								case "image/png":
									img, _ = png.Decode(&fileBuf)
								case "image/jpeg":
									img, _ = jpeg.Decode(&fileBuf)
								case "image/webp":
									img, _ = webp.Decode(&fileBuf)
									isWebP = true
								}

								imgWidth := img.Bounds().Bounds().Max.X
								imgHeight := img.Bounds().Bounds().Max.Y
								aspectRatio := float32(imgWidth) / float32(imgHeight)

								// Create webp options object
								webpOptions, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
								if err != nil {
									errChan <- uploadError{Filename: fileHeader.Filename, Error: err.Error()}
									return
								}

								// Create optimized images with required width
								sizes := *sizesArr
								if len(sizes) > 0 {

									imageSizesChan := make(chan []string)
									errorSizesChan := make(chan error)

									for _, size := range sizes {
										size := size

										go func() {
											sizeString := strconv.Itoa(size)

											// Create output file
											optimizedName := path.Join(fullPath, id.String()+"_"+sizeString+"_optimized.webp")
											outFile, err := os.Create(optimizedName)
											if err != nil {
												errChan <- uploadError{Filename: fileHeader.Filename, Error: err.Error()}
												return
											}
											defer outFile.Close()

											if size == 0 || size >= imgWidth {

												// if source img is webp, just copy
												if isWebP {

													_, err = io.Copy(outFile, file)
													if err != nil {
														errorSizesChan <- err
														return
													}

													// otherwise encode to webp
												} else {

													if err := cwebp.Encode(outFile, img, webpOptions); err != nil {
														errorSizesChan <- err
														return
													}
												}

											} else {

												// Create new image with long side target size
												dst := image.NewRGBA(image.Rect(0, 0, size, int(float32(size)/aspectRatio)))

												// Resize
												draw.BiLinear.Scale(dst, dst.Rect, img, img.Bounds(), draw.Over, nil)

												// Encode to jpeg
												if err := cwebp.Encode(outFile, dst, webpOptions); err != nil {
													errorSizesChan <- err
													return
												}
											}

											imageSizesChan <- []string{sizeString, optimizedName}
										}()
									}

									imageSizes := make(map[string]string)
									var errorSizes error

									for range sizes {
										select {
										case sizes := <-imageSizesChan:
											imageSizes[sizes[0]] = sizes[1]
										case err := <-errorSizesChan:
											errorSizes = err
										}
									}

									if errorSizes != nil {
										errChan <- uploadError{Filename: fileHeader.Filename, Error: errorSizes.Error()}
										return
									}
								}
							}
						}

						resultChan <- []string{name, "/" + strings.Replace(fileName, fullPath, uploadPath, 1)}
					}()
				}

			}

			// Create uploaded files map and errors slice
			filesUrlMap := make(map[string]interface{})
			errorSlice := []uploadError{}

			for range fileList {
				select {
				case result := <-resultChan:
					name := result[0]
					fileName := result[1]

					// Add uploaded file url to map as string or append to existing slice
					if filesUrlMap[name] == nil {

						filesUrlMap[name] = fileName

					} else {

						if existingStr, ok := filesUrlMap[name].(string); ok {
							filesUrlMap[name] = []string{existingStr, fileName}
						}

						if existingSlice, ok := filesUrlMap[name].([]string); ok {
							filesUrlMap[name] = append(existingSlice, fileName)
						}
					}

				case err := <-errChan:
					errorSlice = append(errorSlice, err)
				}
			}

			var resultJson string

			// Convert map with files urls to JSON object
			if len(filesUrlMap) != 0 {
				fileList, err := json.Marshal(filesUrlMap)
				if err != nil {
					exception.InternalServerError(w, err)
					return
				}
				resultJson = string(fileList)
			}

			// Add errors to result
			if len(errorSlice) != 0 {
				uploadErrors, err := json.Marshal(uploadErrors{UploadErrors: errorSlice})
				if err != nil {
					exception.InternalServerError(w, err)
					return
				}
				if len(resultJson) != 0 {
					resultJson = strings.TrimSuffix(resultJson, "}") + "," + strings.TrimPrefix(string(uploadErrors), "{")
				} else {
					resultJson = string(uploadErrors)
				}
			}

			// Add `serialized-json` field value and JSON object with files urls to result JSON
			if jsonField := form.Value["serialized-json"]; len(jsonField) != 0 {
				if len(resultJson) != 0 {
					resultJson = strings.TrimSuffix(jsonField[0], "}") + "," + strings.TrimPrefix(resultJson, "{")
				} else {
					resultJson = jsonField[0]
				}
			}

			// Add errors if exists

			// Write result JSON string to request body
			w.Header().Set("Content-Type", "application/json")
			r.Body = io.NopCloser(strings.NewReader(resultJson))
		}

		next.ServeHTTP(w, r)
	})
}

// Check if one of the list strings matches the request URL path
func inList(list []string, r *http.Request) bool {
	path := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	return slices.ContainsFunc(list, func(s string) bool {
		return slices.Contains(path, s)
	})
}
