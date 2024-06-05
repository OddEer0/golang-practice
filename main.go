package main

import "fmt"

type Valuer interface {
	Value() string
}

type strucc struct {
	Val string
}

func (s *strucc) Value() string {
	return s.Val
}

func testFn(val interface{}) string {
	switch val.(type) {
	case string:
		return val.(string)
	case Valuer:
		ival := val.(Valuer)
		return ival.Value()
	}
	return "not"
}

func main() {
	// runners.RunLogSlogPractice()
	re := &strucc{Val: "kekekekek"}
	fmt.Println(testFn(re))
}
