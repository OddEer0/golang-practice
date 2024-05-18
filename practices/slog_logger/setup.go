package sloglogger

import (
	"log/slog"
	"os"
)

var file *os.File

func SetupLogger() *slog.Logger {
	file, err := os.OpenFile("logger.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil
	}
	opt := &slog.HandlerOptions{Level: slog.LevelDebug}
	logger := slog.New(CtxHandler{(slog.NewJSONHandler(file, opt))})
	return logger
}

func CloseLogger() {
	file.Close()
}
