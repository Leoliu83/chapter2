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
	s.Init()
	for i := 0; i < 3; i++ {
		log.Println("--- Push: ", i)
		s.Push(Element(i))
		util.PrintSliceHeader(s.data, reflect.Int)
		log.Println("+++ Push: ", i)
	}
	util.PrintSliceHeader(s.data, reflect.Int)

	// log.Printf("%+v", s)
}

func TestInit(t *testing.T) {
	var s Stack
	log.Printf("%p,%p", &s, &s.data)
	/*
		下面这种写法与 (&s).Init() 该写法是等价的
		由于Init定义的receiver是一个指针，receiver会自动获取指针
	*/
	s.Init()
	log.Printf("%+v, cap: %d, len: %d", s.data, cap(s.data), len(s.data))
}

func TestPop(t *testing.T) {
	var s Stack
	s.Init()
	for i := 0; i < 3; i++ {
		s.Push(Element(i))
	}
	for i := 0; i < 5; i++ {
		e, ok := s.Pop()
		if ok {
			log.Println(e)
			util.PrintSliceHeader(s.data, reflect.Int)
		}
	}
}

func TestDistory(t *testing.T) {
	var s Stack
	s.Init()
	util.PrintSliceHeader(s.data, reflect.Int)
	s.Distory()
	util.PrintSliceHeader(s.data, reflect.Int)
}

func TestInitDoubleStack(t *testing.T) {
	var ds DoubleStack
	ds.Init()
}
