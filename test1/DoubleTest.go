package test1

import (
	"fmt"
)

// DoubleTest 函数用于测试浮点精度
func DoubleTest() {
	// float32的有效位数是 小数位7位+整数位1位
	var a float32 = 1.123456789123456789
	var b float32 = 1.123456789
	var c float32 = 1.12345

	fmt.Printf("%v,%v,%v \n", a, b, c)
	fmt.Println(a == b, b == c, a == c)

	// 类型转换
	var d float64 = float64(c) // go 语言要求强制的显示类型转换
	fmt.Printf("%v", d)

	p := 123456
	e := (*int)(&p)
	fmt.Printf("%d,%p", e, e)

}
