package test1

import (
	"testing"
	"time"
)

/*
	命令行参数为 go test -bench .   .表示匹配所有测试的正则表达式
	它通过逐步调整B.N的值，反复执行测试函数，知道能够获得准确的测量结果
	-run=NONE 可以忽略所有的单元测试用例，而仅仅执行性能测试
	-gcflags "-N -l" 表示进制内联和优化，以观察结果
	go test -bench=.    // 对所有的进行基准测试
	go test -bench='fib$'     // 只运行以 fib 结尾的基准测试，-bench 可以进行正则匹配
	go test -bench=. -benchtime=6s  // 基准测试默认时间是 1s，-benchtime 可以指定测试时间
	go test -bench=. -benchtime=50x  // 参数 -benchtime 除了指定时间，还可以指定运行的次数
	go test -bench=. -benchmem // 进行时间、内存的基准测试
*/
func BenchmarkAdd111(b *testing.B) {
	time.Sleep(time.Second)
	b.ResetTimer() // 重置计时器
	println("B.N = ", b.N)

	for i := 0; i < b.N; i++ {
		_ = Heap()
		if i == 1 {
			b.StopTimer() // 暂停计时器
			time.Sleep(time.Second)
			b.StartTimer() // 恢复计时器
		}
	}
}

func BenchmarkHeap(b *testing.B) {
	b.ReportAllocs() // 设置总是输出内存信息
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Heap()
	}
}

func Heap() []byte {
	return make([]byte, 1024*10)
}
