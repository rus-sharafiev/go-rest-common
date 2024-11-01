package common

import (
	"slices"
)

type converterConfig struct {
	Whitelist         *[]string          `json:"whitelist"`
	Blacklist         *[]string          `json:"blacklist"`
	UploadPath        *string            `json:"uploadPath"`
	UseUserSubfolder  *bool              `json:"useUserSubfolder"`
	UploadPathByRoute *map[string]string `json:"uploadPathByRoute"`
}

type config struct {
	Db                *string          `json:"db"`
	Port              *string          `json:"port"`
	StaticDir         *string          `json:"staticDir"`
	JwtKey            *string          `json:"jwtKey"`
	RecaptchaSecret   *string          `json:"recaptchaSecret"`
	MailLogin         *string          `json:"mailLogin"`
	MailPassword      *string          `json:"mailPassword"`
	MailHost          *string          `json:"mailHost"`
	RefreshCookiePath *string          `json:"refreshCookiePath"`
	ConverterConfig   *converterConfig `json:"formdataConverter"`
}

var Config config

func (c config) IsNotValid() bool {
	mandatotyFields := []*string{
		c.Port,
		c.StaticDir,
		c.JwtKey,
	}
	return slices.Contains(mandatotyFields, nil)
}
