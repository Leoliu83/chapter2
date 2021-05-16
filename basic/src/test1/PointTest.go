package test1

import (
	"fmt"
	"unsafe"
)

// PointTest 函数用于测试指针相关数据类型的使用 unsafe.Pointer 和 uintptr
// unsafe.Pointer 不可用于地址计算
func PointTest() {
	// 定义一个4个元素的数组
	a := [4]int{1, 2, 3, 4}
	// 定义一个指针变量，指向a数组中的第二个元素地址，&是取址符
	p1 := unsafe.Pointer(&a[1])
	// 定义一个指针变量，指向a数组中的第四个元素 uintptr(p1) 表示地址的整数,可以用于地址计算，后移两个数组元素长度的位置，即可以得到第四个元素的地址
	p2 := unsafe.Pointer(uintptr(p1) + 2*unsafe.Sizeof(a[0]))
	fmt.Printf("%p,%d \n", p2, p2)
	// 将p2转化为int类型指针变量，然后设置该指针变量对应内存中的值
	*(*int)(p2) = 6
	fmt.Println("a= ", a)
}
