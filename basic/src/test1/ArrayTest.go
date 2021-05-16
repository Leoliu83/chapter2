package test1

import (
	"log"
	"unsafe"
)

/*
	定义数组时，数组长度必须是非负整数常量表达式（非负、整数、常量），
	长度是类型的组成部分，也就是说，数组类型相同，但长度不同的数组不属于同一类
*/
func ArrayTest() {
	// 定义三元素int数组，元素自动初始化为0
	var a0 [3]int
	var a1 [3]int
	log.Printf("数组[a1]值为 [%d] \n", a1)
	var a2 [4]int
	// cannot use a2 (variable of type [2]int) as [3]int value in assignment
	// a1 = a2

	// 初始化三元素int数组，数据为 [2,5,1]
	a1 = [3]int{2, 5, 1}
	log.Printf("数组[a1]值为 [%d] \n", a1)
	// 初始化三元素int数组，第一个元素为1 下标为2的元素为10（2: 10），不指定则为0
	a1 = [3]int{1, 2: 10}
	log.Printf("数组[a1]值为 [%d] \n", a1)
	// 初始化int数组（...表示根据元素初始化值得数量确定数组长度），数据为 [1,2,3]
	a1 = [...]int{1, 2, 3}
	log.Printf("数组[a1]值为 [%d] \n", a1)
	// 初始化int数组（...表示根据元素初始化值得数量确定数组长度），第一个元素为10，下标为3的元素为100（3: 100），不指定则为0
	a2 = [...]int{10, 3: 100}
	log.Printf("数组[a2]值为 [%d] \n", a2)

	// 如果数组是结构体等复杂类型，则初始化元素时，可以省略类型标签
	s := [...]Student{
		// 这里可以不写 Student{..., ...}
		{Sid: 1, Sname: "Tom"},
		{Sid: 2, Sname: "Jack"},
	}
	log.Printf("数组[s]值为 [%v] \n", s)

	// 定义多维数组时，仅仅允许第一维使用 '...'
	ma := [...][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	log.Printf("数组[ma]值为: %d \n", ma)
	// len 和 cap 只返回第一维度的长度
	log.Printf("数组[ma]的长度(len): %d, 容量(cap): %d \n", len(ma), cap(ma))
	// 三维数组
	ma1 := [...][2][2]int{
		{
			{1, 2},
			{3, 4},
		},
		{
			{5, 6},
			{7, 8},
		},
	}
	log.Printf("数组[ma1]值为: %d \n", ma1)
	log.Printf("数组[ma1]的长度(len): %d, 容量(cap): %d \n", len(ma1), cap(ma1))

	// 同类型数组（长度相等，元素类型相同）之间，可以使用 == 或者 != 进行比较
	isEqual := (a0 == a1)
	log.Printf("数组[a0] 和 数组[a1] 是否相等: %t \n", isEqual)
}

func ArrayPointTest() {
	a0 := [...]int{1, 2}
	log.Printf("数组[a0]的地址是: %p，值为: %d \n", &a0, a0)
	log.Printf("数组[a0]的元素的地址分别是: [%p %p] \n", &a0[0], &a0[1])
	// 指针可以直接用于元素操作
	a0p := &a0
	a0p[1] += 10
	log.Printf("数组[a0]的地址是: %p, 值为: %d \n", &a0, a0)
	log.Printf("数组[a0]的元素的地址分别是: [%p %p] \n", &a0[0], &a0[1])
	/*
		通过 unsafe.Pointer 可以将原有数组类型强转成更长数组类型，以实现数组越界访问
		例如：这里的 a0 是一个二维数组，通过 unsafe.Printer 强转成了三维数组，并且可以访问第三个元素，
			数值可能是原内存中得值
		需求！？
	*/
	a0up := unsafe.Pointer(&a0)
	a1 := (*[3]int)(a0up)
	log.Printf("数组[a0]的地址是: %p, 值为: %d \n", &a1, a1)
	log.Printf("数组[a0]的元素的地址分别是: [%p %p %p] \n", &a1[0], &a1[1], &a1[2])
}

/*
	go的数组是值类型，与c不同，因此数组的赋值和传参都会将整个数组进行拷贝
	ArrayCopyTest 中将 a0 做为参数传递给 arrayAsParam，可以看到地址发生了变化
	可以改用指针或者切片，以避免数据复制
*/
func ArrayCopyTest() {
	a0 := [...]int{1, 2}
	log.Printf("数组变量[a0]的地址是: %p, 值为: %d \n", &a0, a0)
	log.Printf("数组变量[a0]的元素的地址分别是: [%p %p] \n", &a0[0], &a0[1])
	arrayAsParam(a0)
	arrayAsPointParam(&a0)
}

func arrayAsParam(a [2]int) {
	log.Printf("数组参数[a]的地址是: %p, 值为: %d \n", &a, a)
	log.Printf("数组参数[a]的元素的地址分别是: [%p %p] \n", &a[0], &a[1])
}

// a1 作为指针，地址不再是数组首元素的地址，而是指针变量地址
func arrayAsPointParam(a1 *[2]int) {
	log.Printf("数组参数[a1]的地址是: %p, 值为: %d \n", &a1, a1)
	log.Printf("数组参数[a1]的元素的地址分别是: [%p %p] \n", &a1[0], &a1[1])
}
