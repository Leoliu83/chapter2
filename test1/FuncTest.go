package test1

import (
	"fmt"
	"unsafe"
)

// Calculate 表示一种运算函数类型
type Calculate func(x int32, y int32) int32

/*
	定义一个结构体，里面的属性都是函数
*/
type MultiCalculate struct {
	Add   func(a int, b int) int
	Minus func(a int, b int) int
}

var add Calculate = func(x, y int32) int32 {
	fmt.Printf("%d + %d = %d \n", x, y, x+y)
	return x + y
}

func run(c Calculate, x int32, y int32) {
	c(x, y)
}

// FuncTest 方法用于测试函数类型
func FuncTest() {
	fmt.Printf("%T,%d \n", add, unsafe.Sizeof(add))
	run(add, 5, 6)

	mc := MultiCalculate{
		Add: func(a int, b int) int {
			return a + b
		},
		Minus: func(a int, b int) int {
			return a - b
		},
	}
	fmt.Printf("a + b = %d \n", mc.Add(1, 2))
	fmt.Printf("a - b = %d \n", mc.Minus(1, 2))
	ClosureTest()
	DeferTest()
}

/*
	闭包测试，闭包是指，外部函数可以访问当前函数的内部函数的变量
	在这里就是 ClosureTest 调用 ClosureFunc 所返回的匿名函数，还能正确的访问到x
	闭包得以实现，是因为，ClosureFunc返回的不仅仅是函数，还有x的变量的指针
*/
func ClosureFunc(x int) (func(), []func()) {
	fmt.Printf("outter -> x -> %p \n", &x)
	var s []func()
	// for循环的i是复用的，因此i的地址永远是一个不变
	for i := 0; i < 3; i++ {
		fmt.Printf("for -> i -> %p \n", &i)
		// x每次都会分配一个新的地址来放值，如果不使用x 返回的函数列表中的函数所调用的i都将为最终值
		j := i
		s = append(s, func() {
			x += i
			fmt.Printf("i -> %p,%d \n", &i, i)
			fmt.Printf("j -> %p,%d \n", &j, j)
			fmt.Printf("x -> %p,%d \n", &x, x)
		})
	}
	return func() {
		fmt.Println(x)
	}, s
}

func ClosureTest() {
	f1, f2 := ClosureFunc(5)
	f1()
	for _, f := range f2 {
		f()
	}
	// 上面for循环中将上下文环境中的x进行了+i操作，因此影响了下面f1()函数的值
	f1()

}

func DeferFunc() (x int) {
	defer func() {
		fmt.Printf("defer: %p,%d \n", &x, x)
		x += 100
	}()
	return 100
}

func DeferTest() {
	fmt.Printf("test: %d \n", DeferFunc())
}
