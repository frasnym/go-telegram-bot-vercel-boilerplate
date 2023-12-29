package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/frasnym/go-telegram-bot-vercel-boilerplate/common/ctxdata"
	"github.com/frasnym/go-telegram-bot-vercel-boilerplate/common/logger"
	"github.com/frasnym/go-telegram-bot-vercel-boilerplate/config"
	"github.com/frasnym/go-telegram-bot-vercel-boilerplate/pkg/telebot"
	"github.com/frasnym/go-telegram-bot-vercel-boilerplate/repository"
)

// IndexHandler handles incoming HTTP requests and sets up a webhook for a Telegram bot.
// It takes an HTTP response writer (w) and a request (r), and ensures that the bot's webhook is properly configured.
// If any errors occur during the process, they are logged.
// After the webhook is set up successfully, it writes an "Index OK" message to the response writer (w).
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	now := time.Now()
	ctx := ctxdata.EnsureCorrelationIDExist(r)

	// Log any errors and write "Index OK" as the API response
	defer func() {
		logger.LogHandler(ctx, r, err, &now)
		fmt.Fprintf(w, "%s OK", r.URL.Path)
	}()

	if r.URL.Path == "/favicon.ico" {
		return
	}

	// Create a new bot repository with the application's configuration and Telegram bot
	botRepo := repository.NewBotRepository(config.GetConfig(), telebot.GetBot())

	// Set up the webhook for the bot
	err = botRepo.SetWebhook(ctx)
	if err != nil {
		err = fmt.Errorf("err botRepo.SetWebhook: %w", err)
	}
}
