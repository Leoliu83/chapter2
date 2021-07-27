package stack

import (
	"fmt"
	"log"
)

/*
	这里也将Element 作为 int的别名，这样int也可以传入Element参数也可以使用，
	这里只是为了测试方便使用，实际将使用自定义类型，即 type Element int
*/
type Element = int

type noCopy struct{}

// Lock 方法并不需要使用，只是用于 `go vet`的 -copylocks 检测，也就是表示，该对象不应该使用值传递
func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

/*
	Stack 增加noCopy类型的属性，表示该对象不应该使用值传递
	虽然slice类型的data在值传递时，只是复制头，底层数组不会被复制，但是，值传递中的len和cap也会被复制，
	从而导致调用者传递的slice的头部的len和cap不会发生变化，无法获取append扩容所造成的len和cap的变化
*/
type Stack struct {
	data []Element
	_    noCopy
}

/*
	初始化栈，主要用于初始化栈最大值即 cap ，当前元素个数 0
*/
func (s *Stack) Init() {
	s.data = make([]Element, 0, 2)
	log.Printf("Init: %p, %+v, %p", s, s, &s)
}

func (s *Stack) Distory() {
	*s = Stack{}
}

func (s *Stack) Push(e Element) {
	s.data = append(s.data, e)
}

func (s *Stack) Pop() (e Element, ok bool) {
	e = 0
	ok = false
	l := len(s.data)
	if l == 0 {
		log.Println("Stack is empty!")
		return
	}
	e = s.data[l-1]
	s.data = s.data[0 : l-1]
	ok = true
	return
}

/*
	两栈共享存储空间的实现
*/
type DoubleStack struct {
	data []Element
	top1 int
	top2 int
}

func (ds *DoubleStack) Init() {
	bs := fmt.Sprintf("%b", 9)
	log.Println(bs)
	rs := 0
	for i := 9; i > 0; i >>= 1 {
		rs += (i & 1)
	}
	log.Println(rs)
}
