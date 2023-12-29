package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/frasnym/go-telegram-bot-vercel-boilerplate/common"
	"github.com/frasnym/go-telegram-bot-vercel-boilerplate/common/ctxdata"
	"github.com/frasnym/go-telegram-bot-vercel-boilerplate/common/logger"
	"github.com/frasnym/go-telegram-bot-vercel-boilerplate/common/session"
	"github.com/frasnym/go-telegram-bot-vercel-boilerplate/config"
	"github.com/frasnym/go-telegram-bot-vercel-boilerplate/pkg/telebot"
	"github.com/frasnym/go-telegram-bot-vercel-boilerplate/repository"
	"github.com/frasnym/go-telegram-bot-vercel-boilerplate/service"
)

// WebhookHandler handles incoming HTTP requests for a Telegram bot's webhook.
// It processes updates, handles commands, and takes appropriate actions based on the received messages.
// If any errors occur during the process, they are logged.
// After processing the request, it writes a "Webhook OK" message to the response writer (w).
func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	now := time.Now()
	ctx := ctxdata.EnsureCorrelationIDExist(r)

	// Log any errors and write "Webhook OK" as the API response
	defer func() {
		logger.LogHandler(ctx, "WebhookHandler", err, &now)
		fmt.Fprint(w, "WebhookHandler OK")
	}()

	// Create a new bot repository with the application's configuration and Telegram bot
	cfg := config.GetConfig()
	botRepo := repository.NewBotRepository(cfg, telebot.GetBot())
	exampleSvc := service.NewExampleService(botRepo)

	// Get the update from the request body
	update, err := botRepo.GetUpdate(ctx, r.Body)
	if err != nil {
		err = fmt.Errorf("err botRepo.GetUpdate: %w", err)
		return
	}

	// Handle messages and commands
	if update.Message != nil {
		userID := update.Message.From.ID
		chatID := update.Message.Chat.ID

		// Handle commands
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case common.CommandExample:
				if err = exampleSvc.Request(ctx, userID, chatID); err != nil {
					err = fmt.Errorf("err exampleSvc.Request: %w", err)
				}
				return
			default:
				err = fmt.Errorf("invalid command: %s", update.Message.Command())
				return
			}
		}

		// Get the user's current action
		action, errSession := session.GetAction(userID)
		if errSession != nil {
			logger.Warn(ctx, fmt.Sprintf("session.GetAction: %s", errSession.Error()))
			return
		}

		// Handle requests based on the user's current action
		switch action {
		case common.CommandExample:
			if err = exampleSvc.Processor(ctx, userID, update.Message.Text); err != nil {
				err = fmt.Errorf("err exampleSvc.Processor: %w", err)
			}
			return
		}

		err = fmt.Errorf("unprocessable text: %s", update.Message.Text)
	}
}
