package main

import (
	"fmt"
	"github.com/pkg/errors"
)

var StackDepth = 10

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func main() {
	err := errors.WithStack(errors.New("kekew"))
	fmt.Printf("%+v", err)
}
