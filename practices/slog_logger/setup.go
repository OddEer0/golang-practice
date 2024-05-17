package sloglogger

import (
	"log/slog"
	"os"
)

func SetupLogger() *slog.Logger {
	opt := &slog.HandlerOptions{Level: slog.LevelDebug}
	logger := slog.New(CtxHandler{(slog.NewJSONHandler(os.Stdout, opt))})
	return logger
}
