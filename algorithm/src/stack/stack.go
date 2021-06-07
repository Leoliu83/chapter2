package stack

import (
	"gostudy/algorithm/src/util"
	"log"
	"reflect"
)

// 这里也可以将Element 作为 int的别名，这样int也可以传入Element参数也可以使用，但这里使用自定义类型
type Element int

type Stack struct {
	data []Element
}

func (s *Stack) Init() {
	s.data = make([]Element, 0, 10)
	log.Printf("Init: %p,%v", s, s)
}

func (s Stack) InitNotRight() {
	util.PrintSliceHeader(s.data, reflect.Int)
	s.data = make([]Element, 4, 10)
	util.PrintSliceHeader(s.data, reflect.Int)
}

func (s Stack) PushNotRight(e Element) {
	s.data = append(s.data, e)
}

func (s *Stack) Push(e Element) {
	s.data = append(s.data, e)
}
