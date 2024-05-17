package sloglogger

import (
	"context"
	"strings"
)

const (
	stackTraceKey = "request_id"
)

type Trace struct {
	stack []string
	block bool
}

func IsLockStackTrace(ctx context.Context) bool {
	val, ok := ctx.Value(stackTraceKey).(*Trace)
	if ok {
		return val.block
	}
	return false
}

func LockStackTrace(ctx context.Context) {
	val, ok := ctx.Value(stackTraceKey).(*Trace)
	if ok {
		val.block = true
	}
}

func UnlockStackTrace(ctx context.Context) {
	val, ok := ctx.Value(stackTraceKey).(*Trace)
	if ok {
		val.block = false
	}
}

func AddStackTrace(ctx context.Context, trace string) {
	val, ok := ctx.Value(stackTraceKey).(*Trace)
	if ok {
		val.stack = append(val.stack, trace)
	}
}

func DoneStackTrace(ctx context.Context) {
	val, ok := ctx.Value(stackTraceKey).(*Trace)
	if ok {
		val.stack = val.stack[:len(val.stack)-1]
	}
}

func InitStackTrace(ctx context.Context) context.Context {
	return context.WithValue(ctx, stackTraceKey, &Trace{
		stack: make([]string, 0, 10),
	})
}

func InitStackTraceWithCap(ctx context.Context, cap int) context.Context {
	return context.WithValue(ctx, stackTraceKey, &Trace{
		stack: make([]string, 0, cap),
	})
}

func GetStackTrace(ctx context.Context) string {
	val, ok := ctx.Value(stackTraceKey).(*Trace)
	if ok {
		var s strings.Builder
		s.WriteString("[")
		for i, tr := range val.stack {
			s.WriteString(tr)
			if i != len(val.stack)-1 {
				s.WriteString(", ")
			}
		}
		s.WriteString("]")
		return s.String()
	}
	return "not trace"
}
