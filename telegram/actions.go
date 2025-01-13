package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	common "github.com/rus-sharafiev/go-rest-common"
)

func (c *controller) SendMessage(message MessageToSend) error {
	msgJson, err := json.Marshal(message)
	if err != nil {
		return err
	}

	telegramApiUrl := "https://api.telegram.org/bot" + common.Config.Telegram.BotToken
	resp, err := http.Post(telegramApiUrl+"/sendMessage", "application/json", bytes.NewBuffer(msgJson))
	if err != nil {
		return err
	}

	var result BasicResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	if !result.Ok && result.Description != nil {
		return errors.New(*result.Description)
	}

	return nil
}
