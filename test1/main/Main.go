package main

// 导入自定义的包，chapter2/test1 包下的所有函数皆可使用'包名.xxx'的方式调用
// 一个包导入对应一个目录，chapter2在GOPATH路径下，因此可以使用相对路径直接导入
// pprof 为原生性能分析包
import (
	"chapter2/test1"
	// "runtime/pprof"
)

/**
  只有在package main下的main函数才可以运行
*/
func main() {
	// VariableExample 函数在 chapter2/test1/run.go 中
	// test1.VariableExample()
	// Test1 函数在 chapter2/test1/run1.go 中
	// test1.Test1()
	// test1.ChkConst()
	// test1.VarTest()
	// test1.PointTest()
	// test1.StructTest()
	// test1.FuncTest()
	// test1.InterfaceTest()
	// test1.MapTest()
	// test1.MapCompareTest()
	// test1.SliceTest()
	// test1.StrConvTest()
	// test1.DoubleTest()
	// test1.PointTestSenior()
	// test1.ConcurrentTest()
	// test1.ConcurrentMutexTest()
	// test1.StructCompareTest()
	// test1.StructPointTest()
	// test1.StructEmptyTest()
	// test1.AnonymousFiledTest()
	// test1.StructTagTest()
	// test1.StructMemoryTest()
	// test1.StructMemoryAlgnment()
	// test1.MethodTest()
	// test1.MethodAsParamReceiverIsValueTest()
	// test1.MethodAsParamReceiverIsPointerTest()
	// test1.InterfaceTest()
	// test1.InterfaceInternalTest()
	test1.NilInterfaceTest()
}
