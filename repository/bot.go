package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/frasnym/go-telegram-bot-vercel-boilerplate/common"
	"github.com/frasnym/go-telegram-bot-vercel-boilerplate/common/logger"
	"github.com/frasnym/go-telegram-bot-vercel-boilerplate/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// BotRepository is an interface for managing interactions with a Telegram bot.
type BotRepository interface {
	SetWebhook(ctx context.Context) error
	GetUpdate(ctx context.Context, r io.Reader) (*tgbotapi.Update, error)
	SendMessage(ctx context.Context, c tgbotapi.Chattable) (*tgbotapi.Message, error)
	SendTextMessage(ctx context.Context, chatID int64, text string) (*tgbotapi.Message, error)
	DeleteMessage(ctx context.Context, chatID int64, messageID int) (*tgbotapi.Message, error)
}

type botRepo struct {
	bot *tgbotapi.BotAPI
	cfg *config.Config
}

// SendMessage sends a message using the Telegram bot.
func (s *botRepo) SendMessage(ctx context.Context, c tgbotapi.Chattable) (*tgbotapi.Message, error) {
	var err error
	now := time.Now()
	defer func() {
		logger.LogRepository(ctx, "BotSendMessage", err, &now)
	}()

	msg, err := s.bot.Send(c)
	if err != nil {
		err = fmt.Errorf("err bot.Send: %w", err)
		return nil, err
	}

	return &msg, nil
}

// GetUpdate decodes an update from the provided reader.
func (*botRepo) GetUpdate(ctx context.Context, r io.Reader) (*tgbotapi.Update, error) {
	var err error
	now := time.Now()
	defer func() {
		logger.LogRepository(ctx, "BotGetUpdate", err, &now)
	}()

	update := tgbotapi.Update{}
	if err := json.NewDecoder(r).Decode(&update); err != nil {
		err = fmt.Errorf("err json.NewDecoder(r).Decode: %w", err)
		return nil, err
	}

	return &update, nil
}

// SendTextMessage sends a text message to a specific chat.
func (s *botRepo) SendTextMessage(ctx context.Context, chatID int64, text string) (*tgbotapi.Message, error) {
	var err error
	now := time.Now()
	defer func() {
		logger.LogRepository(ctx, "BotSendTextMessage", err, &now)
	}()

	stringMsg := tgbotapi.NewMessage(chatID, text)
	msg, err := s.bot.Send(stringMsg)
	if err != nil {
		err = fmt.Errorf("err bot.Send: %w", err)
		return nil, err
	}

	return &msg, nil
}

// SetWebhook sets up the bot's webhook for receiving updates.
func (s *botRepo) SetWebhook(ctx context.Context) error {
	var err error
	now := time.Now()
	defer func() {
		logger.LogRepository(ctx, "BotSetWebhook", err, &now)
	}()

	webhookURL := fmt.Sprintf("https://%s/webhook", s.cfg.VercelUrl)

	info, err := s.bot.GetWebhookInfo()
	if err != nil {
		err = fmt.Errorf("err bot.GetWebhookInfo: %w", err)
		return err
	}
	if info.URL == webhookURL {
		return common.ErrNoChanges
	}

	_, err = s.bot.SetWebhook(tgbotapi.NewWebhook(webhookURL))
	if err != nil {
		err = fmt.Errorf("err bot.SetWebhook: %w", err)
		return err
	}

	return nil
}

func (r *botRepo) DeleteMessage(ctx context.Context, chatID int64, messageID int) (*tgbotapi.Message, error) {
	var err error
	now := time.Now()
	defer func() {
		logger.LogRepository(ctx, "BotDeleteMessage", err, &now)
	}()

	del := tgbotapi.NewDeleteMessage(chatID, messageID)
	msg, err := r.SendMessage(ctx, del)
	if err != nil {
		err = fmt.Errorf("err SendMessage: %w", err)
		return nil, err
	}

	return msg, nil
}

// NewBotRepository creates a new BotRepository using the provided configuration and Telegram bot.
func NewBotRepository(cfg *config.Config, bot *tgbotapi.BotAPI) BotRepository {
	return &botRepo{cfg: cfg, bot: bot}
}
