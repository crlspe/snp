package logging

import (
	"log/slog"
	"os"
)

func SetLogLevel(level int) {
	var opts = &slog.HandlerOptions{
		Level: Level(level),
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(logger)
}

func Debug(message string, args ...any) {
	slog.Debug(message, args...)
}

func Info(message string, args ...any) {
	slog.Info(message, args...)
}

func Level(level int) slog.Level {
	switch level {
	case -4:
		return slog.LevelDebug
	case 0:
		return slog.LevelInfo
	case 4:
		return slog.LevelWarn
	case 8:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
