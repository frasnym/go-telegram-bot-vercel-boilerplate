package service

import (
	"context"
	"fmt"
	"time"

	"github.com/frasnym/go-telegram-bot-vercel-boilerplate/common"
	"github.com/frasnym/go-telegram-bot-vercel-boilerplate/common/logger"
	"github.com/frasnym/go-telegram-bot-vercel-boilerplate/common/session"
	"github.com/frasnym/go-telegram-bot-vercel-boilerplate/repository"
)

// ExampleService is an interface for managing Example-related actions.
type ExampleService interface {
	Request(ctx context.Context, userID int64, chatID int64) error
	Processor(ctx context.Context, userID int64, input string) error
}

type exampleSvc struct {
	botRepo repository.BotRepository
}

func (s *exampleSvc) Request(ctx context.Context, userID int64, chatID int64) error {
	var err error
	now := time.Now()
	defer func() {
		logger.LogService(ctx, "ExampleRequest", err, &now)
	}()

	// Start a new session for the user
	session.NewSession(userID, chatID, common.CommandExample)

	// Send a request for the phone number
	replyTxt := "Hello, how are you?"
	_, err = s.botRepo.SendTextMessage(ctx, chatID, replyTxt)
	if err != nil {
		err = fmt.Errorf("error sending text message: %w", err)
		return err
	}

	return nil
}

func (s *exampleSvc) Processor(ctx context.Context, userID int64, input string) error {
	var err error
	now := time.Now()
	defer func() {
		session.DeleteUserSession(userID)
		logger.LogService(ctx, "ExampleProcessor", err, &now)
	}()

	if session.IsInteractionTimedOut(userID) {
		err = s.notifyError(ctx, userID, "Request Timeout")
		if err != nil {
			err = fmt.Errorf("err notifyError: %w", err)
		}

		return err
	}

	chatID, err := session.GetChatID(userID)
	if err != nil {
		err = fmt.Errorf("err session.GetChatID: %w", err)
		return err
	}

	// Send a request for the phone number
	replyTxt := fmt.Sprintf("Okay, so you'are %s huh?", input)
	_, err = s.botRepo.SendTextMessage(ctx, chatID, replyTxt)
	if err != nil {
		err = fmt.Errorf("error sending text message: %w", err)
		return err
	}

	return nil
}

// NewExampleService creates a new ExampleService using the provided bot repository.
func NewExampleService(botRepo repository.BotRepository) ExampleService {
	return &exampleSvc{botRepo}
}

func (s *exampleSvc) notifyError(ctx context.Context, userID int64, msg string) error {
	var err error
	now := time.Now()
	defer func() {
		logger.LogService(ctx, "ExampleNotifyError", err, &now)
	}()

	chatID, err := session.GetChatID(userID)
	if err != nil {
		err = fmt.Errorf("err session.GetChatID: %w", err)
		return err
	}

	_, err = s.botRepo.SendTextMessage(ctx, chatID, msg)
	if err != nil {
		err = fmt.Errorf("err botRepo.SendMessage: %w", err)
		return err
	}

	return nil
}
