package session

import (
	"time"
)

type Session struct {
	Action    string
	ChatID    int64
	StartTime time.Time
}
