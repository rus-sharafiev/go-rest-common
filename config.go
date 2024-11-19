package common

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

type converterConfig struct {
	Whitelist         *[]string          `json:"whitelist"`
	Blacklist         *[]string          `json:"blacklist"`
	UploadPath        *string            `json:"uploadPath"`
	UploadPathPrefix  *string            `json:"uploadPathPrefix"`
	UseUserSubfolder  *bool              `json:"useUserSubfolder"`
	UploadPathByRoute *map[string]string `json:"uploadPathByRoute"`
	OptimizeImages    *map[string][]int  `json:"optimizeImages"`
}

type config struct {
	Db                *string          `json:"db"`
	Port              *string          `json:"port"`
	StaticDir         *string          `json:"staticDir"`
	JwtKey            *string          `json:"jwtKey"`
	TelegramKey       *string          `json:"telegramKey"`
	TelegramAuthToken *string          `json:"telegramAuthToken"`
	RecaptchaSecret   *string          `json:"recaptchaSecret"`
	MailLogin         *string          `json:"mailLogin"`
	MailPassword      *string          `json:"mailPassword"`
	MailHost          *string          `json:"mailHost"`
	RefreshCookiePath *string          `json:"refreshCookiePath"`
	ConverterConfig   *converterConfig `json:"formdataConverter"`
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
		c.Db,
		c.Port,
		c.JwtKey,
	}
	return slices.Contains(mandatotyFields, nil)
}
