package test1

import (
	"fmt"
	"reflect"
)

// 常量组中如果不指定类型及值，则与上一行非空常量右值保持一致，也就是说这里的b的类型和值和a保持一致
const (
	a uint16 = 120
	b
	c string = "abc"
)

/*
  iota 可以表示成一组枚举值，从0开始
  a1 = 0 , a2 = 1, a3 = 2
*/
const (
	a1 = iota
	a2
	a3
)

// KB = 1<<10*1, MB = 1<<10*2, GB=1<<10*3
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
)

/*
  一行可以定义多个 iota，这多个iota单独计算
  b1=1 c1=10
  b2=2 c2=20
  ...
*/
const (
	_, _ = iota, iota * 10
	b1, c1
	b2, c2
	b3, c3
)

/*
iota中断必须显式恢复
并且后续自增按行序自增
也就是说这里的
d1 = 1 ,
d2 = 2 ,
d3 = 100,
d4 = 100,
d5 = 5 ,
d6 = 6
*/
const (
	_ = iota
	d1
	d2
	d3 = 100
	d4
	d5 = iota
	d6
)

/*
  iota 默认类型为 int
  可以显式的指定类型,数字依然安航自动累加
  e1 int
  e2 float32
  e3 float32
*/
const (
	_ = iota
	e1
	e2 float32 = iota
	e3
)

// 可以使用明确的自定义类型来定义枚举类型
type color byte

const (
	black color = iota
	red
	green
)

// ChkConst is used for checking const variables
func ChkConst() {
	fmt.Println("a1=", a1, reflect.TypeOf(a1))
	fmt.Println("a2=", a2, reflect.TypeOf(a2))
	fmt.Println("a3=", a3, reflect.TypeOf(a3))

	fmt.Println("b1=", b1, reflect.TypeOf(b1))
	fmt.Println("b2=", b2, reflect.TypeOf(b2))
	fmt.Println("b3=", b3, reflect.TypeOf(b3))

	fmt.Println("c1=", c1, reflect.TypeOf(c1))
	fmt.Println("c2=", c2, reflect.TypeOf(c2))
	fmt.Println("c3=", c3, reflect.TypeOf(c3))

	fmt.Println("d1=", d1, reflect.TypeOf(d1))
	fmt.Println("d2=", d2, reflect.TypeOf(d2))
	fmt.Println("d3=", d3, reflect.TypeOf(d3))
	fmt.Println("d4=", d4, reflect.TypeOf(d4))
	fmt.Println("d5=", d5, reflect.TypeOf(d5))
	fmt.Println("d6=", d6, reflect.TypeOf(d6))

	fmt.Println("a=", a, reflect.TypeOf(a))
	fmt.Println("b=", b, reflect.TypeOf(b))
	fmt.Println("c=", c, reflect.TypeOf(c))

	fmt.Println("e1=", e1, reflect.TypeOf(e1))
	fmt.Println("e2=", e2, reflect.TypeOf(e2))
	fmt.Println("e3=", e3, reflect.TypeOf(e3))

	/*
		下面的代码会有编译错误：invalid operation: cannot take address of a (constant 120 of type uint16)
		因为a为常量，常量会在编译器在预处理阶段直接展开，作为指令数据使用
		fmt.Println("a", &a)
	*/

	test(red)
	test(100)
	/*
		下面的代码会产生编译错误，x必须是color类型
		var x byte=1
		test(x)
	*/
}

func test(c color) {
	fmt.Println(c)
}
