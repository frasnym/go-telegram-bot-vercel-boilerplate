package common

import "errors"

var (
	ErrNoChanges = errors.New("no changes")
	ErrTimeout   = errors.New("timeout")
	ErrNoSession = errors.New("no active session")
)
