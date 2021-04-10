package test1

/*
	go中的默认引用类型只有：
		·函数（func）
		·切片（slice）
		·字典（map）
		·通道（channel）
*/

/*
	banchmark 结果说明
	e.g.：
	[
		goos: windows
		goarch: amd64
		pkg: chapter2/test1
		BenchmarkSliceTest-8   	  878136	      1288 ns/op	    8192 B/op	       1 allocs/op
		PASS
		ok  	chapter2/test1	1.335s
	]
	* BenchmarkSliceTest-8
		BenchmarkSliceTest 是测试函数名，8表示线程数
	* 878136
		表示一共执行次数，也就是b.N的值
	* 1288 ns/op
		表示平均每次操作花费了1288纳秒
	* 8192 B/op
		表示平均每次操作申请了8192 Byte内存
	* 1 allocs/op
		表示每次操作申请了1次内存
	测试运行参数：
	1. 参数-bench，它指明要测试的函数；点字符意思是测试当前所有以Benchmark为前缀函数
	2. 参数-benchmem，性能测试的时候显示测试函数的内存分配大小，内存分配次数的统计信息
	3. 参数-count n,运行测试和性能多少此，默认一次
	如果不希望把某些操作纳入计时，那么在这些操作后执行 b.ResetTimer()
*/
