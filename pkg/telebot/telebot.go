package telebot

import (
	"errors"
	"fmt"

	"github.com/frasnym/go-telegram-bot-vercel-boilerplate/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var bot *tgbotapi.BotAPI

func init() {
	cfg := config.GetConfig()

	newBot, err := tgbotapi.NewBotAPI(cfg.TelegramBotToken)
	if err != nil {
		panic(fmt.Errorf("unable to init tgbotapi: %w", err))
	}

	bot = newBot
}

func GetBot() *tgbotapi.BotAPI {
	if bot == nil {
		panic(errors.New("please init bot first"))
	}

	return bot
}
