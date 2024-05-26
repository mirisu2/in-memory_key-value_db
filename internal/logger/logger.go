package logger

import (
	"client-server-db/internal/config"
	"log/slog"
	"os"
)

func NewLogger(cfg *config.Config) (*slog.Logger, error) {
	var log *slog.Logger
	var handler slog.Handler

	var logLevel slog.Level
	switch cfg.Logging.Level {
	case "debug":
		logLevel = slog.LevelDebug
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	var logOutput *os.File
	switch cfg.Logging.Output {
	case "file":
		file, err := os.OpenFile(cfg.Logging.LogFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return nil, err
		}
		logOutput = file
	default:
		logOutput = os.Stdout
	}

	var ho *slog.HandlerOptions
	switch cfg.Logging.Source {
	case "false":
		ho = &slog.HandlerOptions{Level: logLevel}
	default:
		ho = &slog.HandlerOptions{Level: logLevel, AddSource: true}
	}

	switch cfg.Logging.Format {
	case "json":
		handler = slog.NewJSONHandler(logOutput, ho)
	default:
		handler = slog.NewTextHandler(logOutput, ho)
	}

	log = slog.New(handler)
	return log, nil
}
