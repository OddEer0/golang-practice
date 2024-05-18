package stacktrace

import (
	"context"
	"encoding/json"
)

const (
	traceKey = "stack_trace_id"
)

type Entity struct {
	Pkg    string `json:"pkg"`
	Type   string `json:"type"`
	Method string `json:"method"`
}

type Trace struct {
	stack  []*Entity
	isLock bool
}

func InitWithCap(ctx context.Context, cap int) context.Context {
	trace := &Trace{
		stack:  make([]*Entity, 0, cap),
		isLock: false,
	}

	return context.WithValue(ctx, traceKey, trace)
}

func Add(ctx context.Context, pkg, methodType, method string) {
	trace, ok := ctx.Value(traceKey).(*Trace)
	if ok {
		trace.stack = append(trace.stack, &Entity{Pkg: pkg, Type: methodType, Method: method})
	}
}

func Done(ctx context.Context) {
	trace, ok := ctx.Value(traceKey).(*Trace)
	if ok {
		trace.stack = trace.stack[:len(trace.stack)-1]
	}
}

func Get(ctx context.Context) string {
	trace, ok := ctx.Value(traceKey).(*Trace)
	if ok {
		jsonRes, err := json.Marshal(trace.stack)
		if err != nil {
			return "Not trace"
		}

		return string(jsonRes)
	}
	return "Not trace"
}

func Lock(ctx context.Context) {
	trace, ok := ctx.Value(traceKey).(*Trace)
	if ok {
		trace.isLock = true
	}
}

func Unlock(ctx context.Context) {
	trace, ok := ctx.Value(traceKey).(*Trace)
	if ok {
		trace.isLock = false
	}
}

func IsLock(ctx context.Context) bool {
	trace, ok := ctx.Value(traceKey).(*Trace)
	if ok {
		return trace.isLock
	}
	return false
}
