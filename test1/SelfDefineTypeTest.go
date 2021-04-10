package test1

import (
	"log"
)

func SelfDefineTypeTest() {
	// 自定义类型 mytype 是一个int类型
	// 但是 mytype和int不能做隐式类型转换，因为这两个类型被认为是不同的类型
	// 也不能做比较
	type mytype int
	var a mytype = 1
	var b int = 2

	// cannot compare a == b (mismatched types mytype and int)
	// if a == b {
	// 	log.Println(true)
	// }

	// cannot use a (variable of type mytype) as int value in variable declaration
	// var b int = a
	log.Printf("%T,%T \n", a, b)
}
