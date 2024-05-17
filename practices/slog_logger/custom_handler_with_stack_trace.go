package sloglogger

import (
	"context"
	"log/slog"
)

type CtxHandler struct {
	slog.Handler
}

func (c CtxHandler) Handle(ctx context.Context, r slog.Record) error {
	if !IsLockStackTrace(ctx) {
		r.AddAttrs(slog.String("stack_trace", GetStackTrace(ctx)))
	}
	if r.Level == slog.LevelError {
		LockStackTrace(ctx)
	}
	return c.Handler.Handle(ctx, r)
}
