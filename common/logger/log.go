package logger

type ConsoleLog struct {
	Level         LogLevel `json:"level"`
	CorrelationID string   `json:"correlationID"`
	Message       string   `json:"message"`
}

type LogLevel string

const (
	LogLevelError LogLevel = "Error"
	LogLevelWarn  LogLevel = "Warn"
	LogLevelInfo  LogLevel = "Info"
	LogLevelDebug LogLevel = "Debug"
)
