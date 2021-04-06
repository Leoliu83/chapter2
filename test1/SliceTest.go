package test1

import (
	"fmt"
	"unsafe"
)

// SliceTest 函数做切片测试
func SliceTest() {
	// 初始化一个slice，1表示当前数组元素个数，初始化为0，3为最大容量
	var s = make([]int, 1, 3)
	fmt.Printf("s: %+v,%d \n", s, unsafe.Sizeof(s))
	s[0] = 1
	var s1 = append(s, 1)
	fmt.Printf("s: %+v,%d,cap: %d \n", s, unsafe.Sizeof(s), cap(s))
	fmt.Printf("s1: %+v,%d,cap: %d \n", s1, unsafe.Sizeof(s), cap(s1))
	// s1 append之后，slice长度变为2 没有超过3，因此s1底层的数组元素和s的数组元素是共享的，因此s1[0]改变，s[0]也会改变
	s1[0] = 2
	fmt.Printf("s: %+v,%d,cap: %d \n", s, unsafe.Sizeof(s), cap(s))
	fmt.Printf("s1: %+v,%d,cap: %d \n", s1, unsafe.Sizeof(s), cap(s1))
	// 都以s为append的对象的时候，s3会覆盖s2所扩展的元素，所以s2和s3打印元素都为 3
	var s2 = append(s, 2)
	var s3 = append(s, 3)
	fmt.Printf("s2: %+v,%d,cap: %d \n", s2, unsafe.Sizeof(s2), cap(s2))
	fmt.Printf("s3: %+v,%d,cap: %d \n", s3, unsafe.Sizeof(s3), cap(s3))
	// s4 基于s3进行扩容 所以容量扩到3，s5基于s4扩容，容量扩到4
	var s4 = append(s3, 4)
	var s5 = append(s4, 5)
	fmt.Printf("s4: %+v,%d,cap: %d \n", s4, unsafe.Sizeof(s4), cap(s4))
	fmt.Printf("s5: %+v,%d,cap: %d \n", s5, unsafe.Sizeof(s5), cap(s5))
	/*
		由于初始化的时候，总容量初始化为3，因此当s4扩容完成后，刚好达到总容量3，
		所以此时，s4和s3、s2、s1、s是共享内存的，所以修改了s4[0]也就是第一个元素，
		s3、s2、s1、s的第一个元素也会同时改变
		但是当s5扩展时，容量超过了3，变为了4，因此s5扩展指向了新的切片，也就是新的内存空间，
		因此修改s5[0]不会影响s3、s2、s1、s、s4
	*/
	s4[0] = 444
	s5[0] = 555
	fmt.Printf(" s: %+v,%d,cap: %d \n", s, unsafe.Sizeof(s), cap(s))
	fmt.Printf("s1: %+v,%d,cap: %d \n", s1, unsafe.Sizeof(s1), cap(s1))
	fmt.Printf("s2: %+v,%d,cap: %d \n", s2, unsafe.Sizeof(s2), cap(s2))
	fmt.Printf("s3: %+v,%d,cap: %d \n", s3, unsafe.Sizeof(s3), cap(s3))
	fmt.Printf("s4: %+v,%d,cap: %d \n", s4, unsafe.Sizeof(s4), cap(s4))
	fmt.Printf("s5: %+v,%d,cap: %d \n", s5, unsafe.Sizeof(s5), cap(s5))
	// s4仍然和s3,s2,s1,s公用数组，且与s5互不相关
	s4[0] = 666
	fmt.Printf(" s: %+v,%d,cap: %d \n", s, unsafe.Sizeof(s), cap(s))
	fmt.Printf("s1: %+v,%d,cap: %d \n", s1, unsafe.Sizeof(s1), cap(s1))
	fmt.Printf("s2: %+v,%d,cap: %d \n", s2, unsafe.Sizeof(s2), cap(s2))
	fmt.Printf("s3: %+v,%d,cap: %d \n", s3, unsafe.Sizeof(s3), cap(s3))
	fmt.Printf("s4: %+v,%d,cap: %d \n", s4, unsafe.Sizeof(s4), cap(s4))
	fmt.Printf("s5: %+v,%d,cap: %d \n", s5, unsafe.Sizeof(s5), cap(s5))
}
