package telegram

import (
	"core"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"core/exception"
)

type ResponseMessage struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func (b *Bot) handleUpdates(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if core.Config.Telegram.ApiSecret != r.Header.Get("X-Telegram-Bot-Api-Secret-Token") {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		fmt.Println(err)
		return
	}

	if b.log {
		b, _ := json.MarshalIndent(update, "", "  ")
		fmt.Println(string(b))
	}

	if message := update.Message; message != nil {
		if text := message.Text; text != nil {
			for commandPattern, action := range *b.actions {
				if b.log {
					fmt.Print(commandPattern)
				}
				if matched, _ := regexp.MatchString(commandPattern, *text); matched {
					action(b, update.Message)
					if b.log {
						fmt.Printf(": match with %v", *text)
					}
				}
				if b.log {
					fmt.Println()
				}
			}
		}
	}

}

func (b *Bot) sendMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if core.Config.Telegram.ApiSecret != r.Header.Get("X-Telegram-Bot-Api-Secret-Token") {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	telegramApiUrl := "https://api.telegram.org/bot" + core.Config.Telegram.BotToken
	resp, err := http.Post(telegramApiUrl+"/sendMessage", "application/json", r.Body)
	if err != nil {
		exception.InternalServerError(w, err)
		return
	}

	var result BasicResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		exception.InternalServerError(w, err)
		return
	}

	if !result.Ok && result.Description != nil {
		exception.InternalServerError(w, fmt.Errorf("telegram error"))
		return
	}

	json.NewEncoder(w).Encode(ResponseMessage{
		StatusCode: http.StatusOK,
		Message:    "sent successfully",
	})
}
