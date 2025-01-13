package telegram

import (
	"net/http"

	common "github.com/rus-sharafiev/go-rest-common"
	"github.com/rus-sharafiev/go-rest-common/db"
)

type controller struct {
	db      *db.Postgres
	actions *map[string]func(c *controller, message *Message)
}

func (c *controller) Handler(mux *http.ServeMux, actions *map[string]func(c *controller, message *Message)) {
	if common.Config.Telegram == nil {
		return
	}

	// Set actions for webhooks
	if actions != nil {
		c.actions = actions
	} else {
		return
	}

	// Handle updates from telegram webhook
	mux.HandleFunc("POST /telegram/updates", c.handleUpdates)
	mux.HandleFunc("POST /telegram/updates/{$}", c.handleUpdates)

	// Handle updates from telegram webhook
	mux.HandleFunc("POST /telegram/message", c.sendMessage)
	mux.HandleFunc("POST /telegram/message/{$}", c.sendMessage)

}

var Controller = &controller{db: &db.Instance}
