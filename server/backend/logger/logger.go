package logger

import (
	"log/slog"
	"os"
)

// InitC initializes a standard JSON logger
func InitC() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)
	return logger
}
