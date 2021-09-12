##### go测试相关的说明
**go自带测试工具**
```
golang的工具链和标准库自带了单元测试框架:
    ·测试代码必须放在当前包，以“_test.go”为结尾的文件中。
    ·测试函数以Test为名称前缀
    ·测试命令（go test）忽略以 "_" 或者 "." 开头的测试文件
    ·正常编译操作（go build/install）会忽略测试文件
标准库 testing 提供了专用类型T来控制测试结果和行为：
方法                        说明                    相关
------------+----------------------------------+-----------
Fail           失败: 继续执行当前测试函数
FailNow        失败: 立即终止执行当前测试函数        Failed
SkipNow        跳过: 停止执行当前测试函数            Skip,Skipf,Skipped
Log            输出错误信息。仅失败或 -v 时输出      Logf
Parallel       与有相同设置的测试函数并行执行
Error          Fail + Log                          Errorf
Fatal          FailNow + Log                       Fatalf

golang没有setup 和 teardown机制
说明：
setUp：
    函数是在众多函数或者说是在一个类类里面最先被调用的函数，
    而且每执行完一个函数都要从setUp()调用开始后再执行下一个函数，
    有几个函数就调用他几次，与位置无关，随便放在那里都是他先被调用。
tearDown：
    函数是在众多函数执行完后他才被执行，意思就是不管这个类里面有多少函数，
    他总是最后一个被执行的，与位置无关，放在那里都行，
    最后不管测试函数是否执行成功都执行tearDown()方法；如果setUp()方法失败，
    则认为这个测试项目失败，不会执行测试函数也不执行tearDown()方法。
```


**banchmark 结果说明**
```
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
3. 参数-count n，运行测试和性能多少此，默认一次
4. 参数-benchtime，运行测试时间，默认1s
如果不希望把某些操作纳入计时，那么在这些操作后执行 b.ResetTimer()
```