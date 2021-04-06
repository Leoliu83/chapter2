package test1

import (
	"fmt"
	"unsafe"
)

/*
MixedFractions 表示带分数结构体，numerator 分子，denominator 分母
也可简写为：
type MixedFractions struct {
	numerator,denominator int32
}
*/
type MixedFractions struct {
	wholeNum    int32 // 整数部分
	numerator   int32 // 真分数部分的分子
	denominator int32 // 真分数部分的分母
}

// StructTest 方法用来测试结构体
func StructTest() {
	// 结构体初始化方法1
	var f MixedFractions
	f.denominator = 2
	f.numerator = 1
	f.wholeNum = 1
	fmt.Printf("%+v,%T,%d \n", f, f, unsafe.Sizeof(f))
	// 结构体初始化方法2
	f = MixedFractions{1, 2, 4}
	// %+v 可以打印结构体
	fmt.Printf("%+v,%T,%d \n", f, f, unsafe.Sizeof(f))

}
