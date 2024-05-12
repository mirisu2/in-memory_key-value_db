package logger

import (
	"log/slog"
	"os"
)

var Log *slog.Logger

func init() {
	levelStr := os.Getenv("LOG_LEVEL")
	var level slog.Level = slog.LevelInfo

	switch levelStr {
	case "DEBUG":
		level = slog.LevelDebug
	case "INFO":
		level = slog.LevelInfo
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	}

	Log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
	}))
}
