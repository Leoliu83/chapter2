package test1

import (
	"log"
	"reflect"
)

/*
	三种获取变量类型的方式：
	·reflect.TypeOf
	·在Printf中使用%T
	·switch中使用 <变量>.(type) 这种方式前提是 <变量>类型是一个interface{}
*/
func GetVariableType() {
	// 在这里不可以使用 err:=DivError{x: 1, y: 0}
	// 如果使用了上面这种方式，err则变成了一个DivError类型，而不是一个接口类型,error是一个接口类型
	// 不是接口类型无法使用 err.(type)，会产生编译错误：(variable of type DivError) is not an interface
	var err error = DivError{x: 1, y: 0}
	t := reflect.TypeOf(err)
	log.Printf("type of t is: %T,type of e is: %s", t, t)
	switch e := err.(type) {
	case DivError:
		log.Fatal(e)
	default:
		log.Fatal("default")
	}
}
