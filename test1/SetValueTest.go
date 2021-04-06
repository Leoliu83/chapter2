package test1

import (
	"fmt"
	"os"
)

var x int     // Define int variable ,default value 0
var y = false // 自动推断为boolean类型变量

// VariableExample is for Test variable
// go语言中首字母大写函数表示public，外部可以调用，首字母小写函数为private函数，外部不可见
func VariableExample() {
	fmt.Println("Hello world!")
	fmt.Println("x=", &x, x)
	// = 表示赋值，如果变量未定义，会产生错误
	x = 100
	fmt.Println("x=", &x, x)
	// := 表示声明变量并赋值，由系统自动推断其类型
	z := 100
	// x会申明新的变量
	x, x1 := 200, "abc"
	fmt.Println("x=", &x, x)
	fmt.Println("z=", &z, z)
	fmt.Println("d=", &x1, x1)
	// z 不会申明新的变量，:= 退化成了赋值操作
	/*
		:= 操作符退化操作条件：
		  1. 必须有一个新的变量被申明，在这里是a，
		  2. z的申明必须在同一个作用域，上面的x, d := 200, "abc"中的x又申明了新的变量，是因为，x之前的赋值是在main函数外
	*/
	z, z1 := 1, 2
	fmt.Println("z=", &z, z)
	fmt.Println("a=", &z1, z1)

	f, err := os.Open("D:\\Users\\liubin18\\Desktop\\Test.java")
	buf := make([]byte, 1024)
	// err 退化为赋值，n为新变量
	n, err := f.Read(buf)
	fmt.Println(err)
	fmt.Println(n)
}
