package runners

import (
	"context"
	"log/slog"

	sloglogger "github.com/OddEer0/golang-practice/practices/slog_logger"
)

func innerFn(ctx context.Context, logger *slog.Logger) {
	sloglogger.AddStackTrace(ctx, "fn: RunLogSlogPractice")
	defer sloglogger.DoneStackTrace(ctx)
	logger.ErrorContext(ctx, "inner func")
}

func RunLogSlogPractice() {
	ctx := context.Background()
	ctx = sloglogger.InitStackTrace(ctx)
	sloglogger.AddStackTrace(ctx, "fn: RunLogSlogPractice")
	defer sloglogger.DoneStackTrace(ctx)

	logger := sloglogger.SetupLogger()

	logger.Info(
		"message",
		slog.String("first_attr", "first attr value"),
		slog.Group("image",
			slog.String("width", "4000px"),
			slog.String("height", "2500px"),
		),
	)

	logger.InfoContext(ctx, "with ctx")

	innerFn(ctx, logger)
}
