package runners

import (
	"context"
	"log/slog"

	slogLogger "github.com/OddEer0/golang-practice/practices/slog_logger"
	stacktrace "github.com/OddEer0/golang-practice/practices/stack_trace"
)

func innerFn(ctx context.Context, logger *slog.Logger) {
	stacktrace.Add(ctx, "runners", "-", "innerFn")
	defer stacktrace.Done(ctx)
	logger.ErrorContext(ctx, "inner func")
}

func RunLogSlogPractice() {
	ctx := context.Background()
	ctx = stacktrace.InitWithCap(ctx, 10)
	stacktrace.Add(ctx, "runners", "-", "RunLogSlogPractice")
	defer stacktrace.Done(ctx)

	logger := slogLogger.SetupLogger()
	defer slogLogger.CloseLogger()

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
