package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/frasnym/go-telegram-bot-vercel-boilerplate/common/ctxdata"
)

// TODO: Use Library
func Error(ctx context.Context, err error) {
	printToConsole(ctx, LogLevelError, err.Error())
}

func Warn(ctx context.Context, msg string) {
	printToConsole(ctx, LogLevelWarn, msg)
}

func Info(ctx context.Context, msg string) {
	printToConsole(ctx, LogLevelInfo, msg)
}

func Debug(ctx context.Context, msg string) {
	printToConsole(ctx, LogLevelDebug, msg)
}

func printToConsole(ctx context.Context, level LogLevel, msg string) {
	log := ConsoleLog{
		Level:         level,
		CorrelationID: ctxdata.GetCorrelationID(ctx),
		Message:       msg,
	}
	logByte, _ := json.Marshal(log)
	fmt.Println(string(logByte))
}

func LogService(ctx context.Context, name string, err error, startTime *time.Time) {
	if err != nil {
		Error(ctx, fmt.Errorf("[SERVICE] %s error: %v. %v elapsed", name, err, time.Since(*startTime)))
	} else {
		Info(ctx, fmt.Sprintf("[SERVICE] %s success. %v elapsed", name, time.Since(*startTime)))
	}
}

func LogHandler(ctx context.Context, name string, err error, startTime *time.Time) {
	if err != nil {
		Error(ctx, fmt.Errorf("[HANDLER] %s error: %v. %v elapsed", name, err, time.Since(*startTime)))
	} else {
		Info(ctx, fmt.Sprintf("[HANDLER] %s success. %v elapsed", name, time.Since(*startTime)))
	}
}

func LogRepository(ctx context.Context, name string, err error, startTime *time.Time) {
	if err != nil {
		Error(ctx, fmt.Errorf("[REPOSITORY] %s error: %v. %v elapsed", name, err, time.Since(*startTime)))
	} else {
		Info(ctx, fmt.Sprintf("[REPOSITORY] %s success. %v elapsed", name, time.Since(*startTime)))
	}
}

func LogPkg(ctx context.Context, name string, err error, startTime *time.Time) {
	if err != nil {
		Error(ctx, fmt.Errorf("[PKG] %s error: %v. %v elapsed", name, err, time.Since(*startTime)))
	} else {
		Info(ctx, fmt.Sprintf("[PKG] %s success. %v elapsed", name, time.Since(*startTime)))
	}
}
