package eer0StackTracerTest

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	eer0StackTracer "github.com/OddEer0/golang-practice/practices/eer0_stack_tracer"
)

type mockStackTraceFunc struct {
	Package  string `json:"package"`
	Function string `json:"function"`
}

func (m *mockStackTraceFunc) StackTraceValue() string {
	result, err := json.Marshal(m)
	if err != nil {
		return "error marshaling to json"
	}
	return string(result)
}

func TestBaseFunctional(t *testing.T) {
	ctx := eer0StackTracer.Init(context.Background(), &eer0StackTracer.Option{
		StackSize: 10,
		Handler: func(stackEntities ...eer0StackTracer.StackTraceValuer) string {
			if len(stackEntities) == 0 {
				return "Empty"
			}
			res := strings.Builder{}
			res.WriteString("[\n")
			for i, s := range stackEntities {
				if i != 0 {
					res.WriteString(",\n")
				}
				res.WriteString(s.StackTraceValue())
			}
			res.WriteString("\n]")
			return res.String()
		},
	})
	res := eer0StackTracer.Get(ctx)
	if res != "Empty" {
		t.Fatal("incorrect get stack trace getter with empty stack trace")
	}
	mock1 := &mockStackTraceFunc{
		Package:  "mock_1",
		Function: "mockFn_1",
	}
	mock2 := &mockStackTraceFunc{
		Package:  "mock_2",
		Function: "mockFn_2",
	}
	eer0StackTracer.Add(ctx, mock1)
	res = eer0StackTracer.Get(ctx)
	if res != "[\n"+mock1.StackTraceValue()+"\n]" {
		t.Fatal("incorrect get stack trace getter with empty stack trace")
	}
	eer0StackTracer.Add(ctx, mock2)
	res = eer0StackTracer.Get(ctx)
	if res != "[\n"+mock1.StackTraceValue()+",\n"+mock2.StackTraceValue()+"\n]" {
		t.Fatal("incorrect get stack trace getter with empty stack trace")
	}
	eer0StackTracer.Done(ctx)
	res = eer0StackTracer.Get(ctx)
	if res != "[\n"+mock1.StackTraceValue()+"\n]" {
		t.Fatal("incorrect get stack trace getter with empty stack trace")
	}
	eer0StackTracer.Done(ctx)
	res = eer0StackTracer.Get(ctx)
	if res != "Empty" {
		t.Fatal("incorrect get stack trace getter with empty stack trace")
	}
}
