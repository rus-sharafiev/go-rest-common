package telegram

import (
	"core"
	"net/http"

	"core/db"
)

type Bot struct {
	db      *db.Postgres
	actions *map[string]func(b *Bot, message *Message)
	log     bool
}

func (b *Bot) Handler(mux *http.ServeMux, actions *map[string]func(b *Bot, message *Message), logMessages bool) {
	if core.Config.Telegram == nil {
		return
	}

	// Set actions for webhooks
	if actions != nil {
		b.actions = actions
	} else {
		return
	}

	b.log = logMessages

	// Handle updates from telegram webhook
	mux.HandleFunc("POST /telegram/updates", b.handleUpdates)
	mux.HandleFunc("POST /telegram/updates/{$}", b.handleUpdates)

	// Handle updates from telegram webhook
	mux.HandleFunc("POST /telegram/message", b.sendMessage)
	mux.HandleFunc("POST /telegram/message/{$}", b.sendMessage)

}

var Controller = &Bot{db: &db.Instance}
