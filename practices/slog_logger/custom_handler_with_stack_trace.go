package slogLogger

import (
	"context"
	"log/slog"

	stacktrace "github.com/OddEer0/golang-practice/practices/stack_trace"
)

type CtxHandler struct {
	slog.Handler
}

func (c CtxHandler) Handle(ctx context.Context, r slog.Record) error {
	if !stacktrace.IsLock(ctx) {
		r.AddAttrs(slog.String("stack_trace", stacktrace.Get(ctx)))
	}
	if r.Level == slog.LevelError {
		stacktrace.Lock(ctx)
	}
	return c.Handler.Handle(ctx, r)
}
