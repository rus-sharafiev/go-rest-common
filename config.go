package common

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

type converterConfig struct {
	Whitelist         *[]string          `json:"whitelist,omitempty"`
	Blacklist         *[]string          `json:"blacklist,omitempty"`
	UploadPath        *string            `json:"uploadPath,omitempty"`
	UploadPathPrefix  *string            `json:"uploadPathPrefix,omitempty"`
	UseUserSubfolder  *bool              `json:"useUserSubfolder,omitempty"`
	UploadPathByRoute *map[string]string `json:"uploadPathByRoute,omitempty"`
	OptimizeImages    *map[string][]int  `json:"optimizeImages,omitempty"`
}

type authConfig struct {
	Prefix *[]string `json:"prefix,omitempty"`
}

type telegram struct {
	BotToken  string `json:"bot_token"`
	ApiSecret string `json:"api_secret"`
}

type config struct {
	Port              *string          `json:"port,omitempty"`
	DbConnString      *string          `json:"dbConnString,omitempty"`
	StaticDir         *string          `json:"staticDir,omitempty"`
	JwtKey            *string          `json:"jwtKey,omitempty"`
	RecaptchaSecret   *string          `json:"recaptchaSecret,omitempty"`
	MailLogin         *string          `json:"mailLogin,omitempty"`
	MailPassword      *string          `json:"mailPassword,omitempty"`
	MailHost          *string          `json:"mailHost,omitempty"`
	RefreshCookiePath *string          `json:"refreshCookiePath,omitempty"`
	ConverterConfig   *converterConfig `json:"formdataConverter,omitempty"`
	Auth              *authConfig      `json:"auth,omitempty"`
	Telegram          *telegram        `json:"telegram,omitempty"`
}

var Config config

func LoadConf() {
	const confFile = "./config.json"

	if _, err := os.Stat(confFile); err != nil && os.IsNotExist(err) {
		return
	} else if err != nil {
		fmt.Printf("\nConfig file: \x1b[31m%v: %v\x1b[0m", "Error opening the file", err)
		return
	}

	data, err := os.ReadFile(confFile)
	if err != nil {
		fmt.Printf("\nConfig file: \x1b[31m%v: %v\x1b[0m", "Error reading the file", err)
		return
	}

	var config config
	if err = json.Unmarshal(data, &config); err != nil {
		fmt.Printf("\nConfig file: \x1b[31m%v: %v\x1b[0m", "Unmarshalling error", err)
		return
	}

	Config = config
}

func (c config) IsNotValid() bool {
	mandatotyFields := []*string{
		c.DbConnString,
		c.Port,
		c.JwtKey,
	}
	return slices.Contains(mandatotyFields, nil)
}
