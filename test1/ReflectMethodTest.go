package test1

import (
	"fmt"
	"log"
	"reflect"
)

/*
	动态调用方法，只需要根据所有的In列表准备好参数，所有的参数必须是reflect.Value类型
*/
type user103 struct {
	Name string
	Age  int
	id   int
}

func (user103) Test(x, y int, strs ...string) (int, error) {
	log.Printf("%+v", strs)
	return x + y, fmt.Errorf("err: %d", x+y)
}

func ReflectMethodTest1() {
	var u user103
	v := reflect.ValueOf(&u)
	m := v.MethodByName("Test")
	inParam := []reflect.Value{
		reflect.ValueOf(100),
		reflect.ValueOf(200),
		// Call 需要一个一个写，后续使用CallSlice更方便
		reflect.ValueOf("a"),
		reflect.ValueOf("b"),
		reflect.ValueOf("c"),
	}
	// out返回一个slice，slice是引用类型，因此是一个指针
	out := m.Call(inParam)
	// reflect.TypeOf(out) 为空，说明out是一个指针类型
	println(reflect.TypeOf(out).Elem().Name(), reflect.TypeOf(out).Elem().Kind()) // Value
	// out 是一个返回值列表
	for _, v := range out {
		log.Println(v)
	}

	// 对于可变长参数，CallSlice更方便，只需要最后是一个数组就可以，可以是string，可以是interface，
	// 需要与调用方法的参数类型匹配
	out = m.CallSlice([]reflect.Value{
		reflect.ValueOf(300),
		reflect.ValueOf(400),
		reflect.ValueOf([]string{"a", "b", "c"}),
	})

	for _, v := range out {
		log.Println(v)
	}
}
