package stack

// 同package下的test不需要写import

import (
	"gostudy/algorithm/src/util"
	"log"
	"reflect"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestPush(t *testing.T) {
	var s Stack
	// s.data = make([]Element, 0, 10)
	util.PrintSliceHeader(s.data, reflect.Int)
	log.Printf("%p", s.data)
	// (&s).Init()
	s.InitNotRight()
	log.Printf("%p", s.data)
	util.PrintSliceHeader(s.data, reflect.Int)
	for i := 0; i < 3; i++ {
		s.Push(Element(i))
	}
	util.PrintSliceHeader(s.data, reflect.Int)

	// log.Printf("%+v", s)
}
