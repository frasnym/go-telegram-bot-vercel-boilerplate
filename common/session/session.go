package session

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/frasnym/go-telegram-bot-vercel-boilerplate/common"
	"github.com/frasnym/go-telegram-bot-vercel-boilerplate/common/logger"
)

// User sessions to store interaction data
var (
	userSessions     = make(map[int]Session)
	userSessionMutex sync.Mutex
)

// NewSession creates a new user session with the given userID, chatID, and action.
func NewSession(userID int, chatID int64, action string) {
	session := Session{
		Action:    action,
		ChatID:    chatID,
		StartTime: time.Now(),
	}
	setUserSession(userID, &session)
}

// ResetTimer resets the session timer for a user's session.
func ResetTimer(userID int) error {
	session, exist := getUserSession(userID)
	if !exist {
		return common.ErrNoSession
	}
	session.StartTime = time.Now() // Renew the timer
	setUserSession(userID, session)

	return nil
}

// GetAction retrieves the current action for a user's session.
func GetAction(userID int) (string, error) {
	session, exist := getUserSession(userID)
	if !exist {
		return "", common.ErrNoSession
	}

	return session.Action, nil
}

// GetChatID retrieves the chat ID for a user's session.
func GetChatID(userID int) (int64, error) {
	session, exist := getUserSession(userID)
	if !exist {
		return 0, common.ErrNoSession
	}

	return session.ChatID, nil
}

// DeleteUserSession deletes a user's session when it's no longer needed.
func DeleteUserSession(userID int) {
	userSessionMutex.Lock()
	defer userSessionMutex.Unlock()

	delete(userSessions, userID)
}

// setUserSession sets the user's session data in a thread-safe manner.
func setUserSession(userID int, newSession *Session) {
	userSessionMutex.Lock()
	defer func() {
		userSessionMutex.Unlock()
		if r := recover(); r != nil {
			logger.Error(context.TODO(), fmt.Errorf("err setUserSession: %v", r))
		}
	}()
	userSessions[userID] = *newSession
}

// getUserSession retrieves the user's session data in a thread-safe manner.
func getUserSession(userID int) (*Session, bool) {
	userSessionMutex.Lock()
	defer userSessionMutex.Unlock()

	session, exists := userSessions[userID]
	return &session, exists
}

// IsInteractionTimedOut checks if a user's session has timed out due to inactivity.
func IsInteractionTimedOut(userID int) bool {
	session, exist := getUserSession(userID)
	if !exist {
		return true
	}

	elapsed := time.Since(session.StartTime)
	return elapsed > common.SessionTimeout
}
