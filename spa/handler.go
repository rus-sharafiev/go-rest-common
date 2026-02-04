package spa

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"core"
)

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(core.Config.StaticDir, r.URL.Path)
	index := "index.html"
	fileType := filepath.Ext(path)

	acceptGzip := strings.Contains(r.Header.Get("Accept-Encoding"), "gzip")

	// Check whether a file exists or is a directory
	if fi, err := os.Stat(path); os.IsNotExist(err) || fi.IsDir() {
		htmlFile := filepath.Join(core.Config.StaticDir, index)
		if acceptGzip {
			if _, err := os.Stat(htmlFile + ".gz"); err == nil {
				w.Header().Add("Content-Encoding", "gzip")
				w.Header().Add("Content-Type", "text/html")
				htmlFile += ".gz"
			}
		}

		// Serve SPA
		http.ServeFile(w, r, htmlFile)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serve gziped
	if acceptGzip && (fileType == ".js" || fileType == ".css") {
		if _, err := os.Stat(path + ".gz"); err == nil {

			w.Header().Add("Content-Encoding", "gzip")
			if fileType == ".js" {
				w.Header().Add("Content-Type", "text/javascript")
			}
			if fileType == ".css" {
				w.Header().Add("Content-Type", "text/css")
			}

			http.ServeFile(w, r, filepath.Join(path+".gz"))
			return
		}
	}

	w.Header().Add("Cache-Control", "no-cache")
	http.FileServer(http.Dir(core.Config.StaticDir)).ServeHTTP(w, r)
}

var Handler = &handler{}
