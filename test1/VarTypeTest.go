package test1

import (
	"fmt"
	"reflect"
	"unsafe"
)

// VarTest 用来查看所有的基本类型的长度
func VarTest() {
	var a bool = false
	var b byte = 1
	// int 和 uint 为默认整数类型，根据目标平台自动为 32 或者 64
	var c int = 1
	var d uint = 1
	// int8 和 uint8 占用1个字节 范围是：int8[-128 ~ 127]，uint8[0 ~ 255]
	var e1 int8 = 1
	var e2 uint8 = 1
	// int32 和 uint32 占用4个字节 范围是：int8[-32768 ~ 32767]，uint8[0 ~ 65535]
	var f1 int32 = 1
	var f2 uint32 = 1
	// int64 和 uint64 占用8个字节 范围是：int64[-32768 ~ 32767]，[0 ~ 65535]
	var g1 int64 = 1
	var g2 uint64 = 1
	var h1 float32 = 1.0
	var h2 float64 = 1.0
	/*
		64位复数,复数由 "实数部分" "+" "虚数部分" "i" 四个部分组成
		其中 "实数部分" "虚数部分" 分别为float32 类型
	*/
	var i1 complex64 = 1.0 + 2.0i
	/*
		128位复数,复数由 "实数部分" "+" "虚数部分" "i" 四个部分组成
		其中 "实数部分" "虚数部分" 分别为float64 类型
	*/
	var i2 complex128 = 1.0 + 2.0i

	var j string = "我"
	/*
		将字符串转化为字符数组 rune类型等同于int32，占用4个字节，用来处理单个unicode字符,只是int32的别名
	*/
	var k1 rune = '我'
	var k2 int32 = '我'
	/*
	  byte 和 rune 类似，只是int8 的别名，等同于int8,用来处理单个ascii字符
	*/
	var l1 byte = 'a'
	var l2 int8 = 'a'

	/*
		uintptr是一个整数类型，可以用于存放指针，即使uintptr变量仍然有效，该变量所表示的地址的数据也会被GC回收
		unsafe.Pointer是一个指针类型，unsafe.Pointer的值不能被取消引用，如果unsafe.Pointer变量有效，该变量所表示的地址处的数据不会被GC回收
	*/
	var m1 = unsafe.Pointer(&l1)
	var m2 uintptr = uintptr(m1)

	fmt.Println("a: ", reflect.TypeOf(a), unsafe.Sizeof(a))
	fmt.Println("b: ", reflect.TypeOf(b), unsafe.Sizeof(b))
	fmt.Println("c: ", reflect.TypeOf(c), unsafe.Sizeof(c))
	fmt.Println("d: ", reflect.TypeOf(d), unsafe.Sizeof(d))
	fmt.Println("e1: ", reflect.TypeOf(e1), unsafe.Sizeof(e1))
	fmt.Println("e2: ", reflect.TypeOf(e2), unsafe.Sizeof(e2))
	fmt.Println("f1: ", reflect.TypeOf(f1), unsafe.Sizeof(f1))
	fmt.Println("f2: ", reflect.TypeOf(f2), unsafe.Sizeof(f2))
	fmt.Println("g1: ", reflect.TypeOf(g1), unsafe.Sizeof(g1))
	fmt.Println("g2: ", reflect.TypeOf(g2), unsafe.Sizeof(g2))
	fmt.Println("h1: ", reflect.TypeOf(h1), unsafe.Sizeof(h1))
	fmt.Println("h2: ", reflect.TypeOf(h2), unsafe.Sizeof(h2))
	fmt.Println("i1: ", reflect.TypeOf(i1), unsafe.Sizeof(i1))
	fmt.Println("i2: ", reflect.TypeOf(i2), unsafe.Sizeof(i2))
	// Sizeof 可以查看所有类型变量所占用的空间，而len只针对string类型或者数组类型
	fmt.Println("j: ", reflect.TypeOf(j), unsafe.Sizeof(j), len(j))
	fmt.Println("k: ", reflect.TypeOf(k1), unsafe.Sizeof(k1))
	fmt.Printf("%d , %T \n", k1, k1)
	fmt.Printf("%d , %T \n", k2, k2)
	fmt.Printf("%d , %T \n", l1, l1)
	fmt.Printf("%d , %T \n", l2, l2)
	fmt.Printf("%d , %T \n", m1, m1)
	fmt.Printf("%d , %T \n", m2, m2)
}
