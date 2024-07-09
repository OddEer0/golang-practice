package pgxLogger

import (
	"github.com/OddEer0/golang-practice/resources/domain"
	"io"
	log2 "log"
	"log/slog"
	"os"
)

const (
	EnvLocal   = "local"
	EnvDev     = "dev"
	EnvProd    = "prod"
	EnvTest    = "test"
	stdoutPath = "stdout"
)

var log domain.Logger = nil

func MustLoad(env string, filePath string) domain.Logger {
	if log != nil {
		return log
	}

	out := os.Stdout
	if filePath != "" && filePath != stdoutPath {
		file, err := os.OpenFile(filePath, os.O_WRONLY, 0666)
		if err != nil {
			log2.Fatalf("log file does not exist: %s", err)
		}
		out = file
	}

	switch env {
	case EnvLocal:
		log = setupLocalLog(out)
	case EnvDev:
		log = setupDevLog(out)
	case EnvProd:
		log = setupProdLog(out)
	case EnvTest:
		log = setupTestLog(out)
	}

	return log
}

func setupLocalLog(out io.Writer) domain.Logger {
	jsonHandler := slog.NewJSONHandler(out, &slog.HandlerOptions{Level: slog.LevelDebug})
	return slog.New(jsonHandler)
}

func setupDevLog(out io.Writer) domain.Logger {
	handler := slog.NewJSONHandler(out, &slog.HandlerOptions{Level: slog.LevelInfo})
	return slog.New(handler)
}

func setupProdLog(out io.Writer) domain.Logger {
	jsonHandler := slog.NewJSONHandler(out, &slog.HandlerOptions{Level: slog.LevelError})
	return slog.New(jsonHandler)
}

func setupTestLog(out io.Writer) domain.Logger {
	handler := slog.NewJSONHandler(out, &slog.HandlerOptions{Level: slog.LevelError})
	return slog.New(handler)
}
