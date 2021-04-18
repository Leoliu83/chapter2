package main

// 导入自定义的包，chapter2/test1 包下的所有函数皆可使用'包名.xxx'的方式调用
// 一个包导入对应一个目录，chapter2在GOPATH路径下，因此可以使用相对路径直接导入
// pprof 为原生性能分析包
import (
	"chapter2/test1"
	// "runtime"
	// "time"
	// "runtime/pprof"
)

/**
  只有在package main下的main函数才可以运行
  如果需要为最终的可执行文件添加命令行参数，可以使用 flag 包
  os.Args 是一个string切片，用于存储所有的命令行参数
  查看所有参数方法：
  for i,v := range os.Args {
	  log.Pringf("args[%v] = %v", i, v)
  }
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
	// test1.DeferTest()
	// test1.DeferParamTest()
	// test1.FuncSignatureTest()
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
	// test1.NilInterfaceTest()
	// test1.InterfaceTypeTransform()
	// test1.goroutineTest()
	// test1.GoroutineWaitTest()
	// test1.GoRoutineParallelTest(16, true)
	// test1.LocalStorageTest()
	// test1.GoschedTest()
	// test1.GoExitTest()
	// test1.ChannelSyncTest()
	// test1.ChannelAsyncTest()
	// test1.ChannelCompareTest()
	// test1.ChannelReceiveTest()
	test1.ChannelMultiReceiveTest()

	// {
	// 	time.Sleep(time.Second)
	// 	runtime.Goexit() // 主进程调用该函数，会等待其他所有goroutine任务执行完成后，让进程崩溃
	// }
}
