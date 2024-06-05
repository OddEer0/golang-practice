package eer0StackTracer

import (
	"context"
	"sync"
)

type (
	Handle func(...StackTraceValuer) string

	Option struct {
		Handler   Handle
		StackSize uint
	}

	StackTraceValuer interface {
		StackTraceValue() string
	}

	stackTrace struct {
		stack   []StackTraceValuer
		mu      sync.Mutex
		handler Handle
		lock    int32
	}
)

func Init(ctx context.Context, opt *Option) context.Context {
	sTrace := stackTrace{
		stack:   make([]StackTraceValuer, 0, opt.StackSize),
		mu:      sync.Mutex{},
		handler: opt.Handler,
	}

	return context.WithValue(ctx, DefaultStackKeyName, &sTrace)
}
