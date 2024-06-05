package eer0StackTracer

import (
	"context"
	"sync/atomic"
)

func Add(ctx context.Context, value StackTraceValuer) {
	sTrace, ok := ctx.Value(DefaultStackKeyName).(*stackTrace)
	sTrace.mu.Lock()
	defer sTrace.mu.Unlock()
	if ok {
		sTrace.stack = append(sTrace.stack, value)
	}
}

func Done(ctx context.Context) {
	sTrace, ok := ctx.Value(DefaultStackKeyName).(*stackTrace)
	sTrace.mu.Lock()
	defer sTrace.mu.Unlock()
	if ok && len(sTrace.stack) > 0 {
		sTrace.stack = sTrace.stack[:len(sTrace.stack)-1]
	}
}

func Get(ctx context.Context) string {
	sTrace, ok := ctx.Value(DefaultStackKeyName).(*stackTrace)
	sTrace.mu.Lock()
	defer sTrace.mu.Unlock()
	if ok {

		return sTrace.handler(sTrace.stack...)
	}

	return DefaulGetterNotFoundStackTraceInContext
}

func Lock(ctx context.Context) {
	sTrace, ok := ctx.Value(DefaultStackKeyName).(*stackTrace)
	if ok {
		atomic.AddInt32(&sTrace.lock, 1)
	}
}

func Unlock(ctx context.Context) {
	sTrace, ok := ctx.Value(DefaultStackKeyName).(*stackTrace)
	if ok {
		atomic.AddInt32(&sTrace.lock, -1)
	}
}

func IsLock(ctx context.Context) bool {
	sTrace, ok := ctx.Value(DefaultStackKeyName).(*stackTrace)
	if ok {
		if sTrace.lock != 0 {
			return true
		}
	}
	return false
}

func IsUnlock(ctx context.Context) bool {
	sTrace, ok := ctx.Value(DefaultStackKeyName).(*stackTrace)
	if ok {
		if sTrace.lock == 0 {
			return true
		}
	}
	return false
}
