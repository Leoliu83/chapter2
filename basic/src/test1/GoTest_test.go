package test1

import (
	"fmt"
	"log"
	"testing"
)

func Add111(x, y int) int {
	return x + y
}

/*
	TestAdd 中将测试数据域测试逻辑分离，更便于维护。
	另外，使用 Errorf 是为了让整个表全部完成测试，以便知道具体是哪组条件出现问题
	这种测试方式称为：table driven （测试数据表驱动测试）
*/
func TestAdd(t *testing.T) {
	var tests = []struct {
		x      int
		y      int
		expect int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{3, 2, 5},
		{7, 8, 11},
	}

	for _, tt := range tests {
		actual := Add111(tt.x, tt.y)
		if actual != tt.expect {
			t.Errorf("add111(%d,%d): expected %d, actual %d", tt.x, tt.y, tt.expect, actual)
		}
	}
}

/*
	Example作用是导入GoDoc等工具来生成帮助文档，它通过对比输出（stdout）结果和内部 Output注释（// Output:）是否一致来判断是否成功
	Example函数命名方式为  Example+<函数名>：
		例如 函数 Testfunc() 的 Example函数为 ExampleTestfunc() 否则会出现编译告警
	如果没有 Output注释，Add111 函数将不执行
	下面的Example会产生错误，因为 Add111(2, 3) 结果应该是5 而Output注释中写的是4
	错误输出如下：
	--- FAIL: ExampleAdd111 (0.00s)
	got:
	3
	5
	want:
	3
	4
*/
func ExampleAdd111() {
	fmt.Println(Add111(1, 2))
	fmt.Println(Add111(2, 3))

	// Output:
	// 3
	// 4
}

/*
	pat 表示命令行参数 -run 提供的过滤条件
	str testing.InternalTest.Name
	原文的案例 match 现在已经无法作为第一个参数传递
	go 1.15中testing.MainStart的第一个参数改成了 testing.testDeps 内部接口
	golang 并不建议直接调用MainStart:
	原文代码如下：
	match := func(pat, str string) (bool, error) {
		return true, nil
	}

	tests := []testing.InternalTest{
		{"b", TestB},
		{"a", TestA},
	}
	benchmarks := []testing.InternalBenchmark{}
	examples := []testing.InternalExample{}

	m = testing.MainStart(match, tests, benchmarks, examples)
*/

/*
	只能在一个文件中定义 TestMain 方法，
	在当前文件所在目录打开命令行，并且在命令行执行 go test 的时候，会默认执行 TestMain
	e.g.
		D:\DeveloperProjects\go\golangstudy\test1> go test
		2021/05/09 09:29:05 Test_test.go:39: begin
		1620523745268457900
		2 : 3
		1620523746276519600
		Run time: 1008 ms
		--- FAIL: TestAdd (0.00s)
			Test_test.go:33: add111(7,8): expected 11, actual 15
		FAIL
		2021/05/09 09:29:06 Test_test.go:41: end
		exit status 1
		FAIL    golangstudy/test1       1.312s
	google  并不建议书中所说的使用 MainStart
*/
// func TestMain(m *testing.M) {
// 	log.Println("begin")
// 	code := m.Run() // m.Run() 会调用具体的测试用例
// 	log.Println("end")
// 	os.Exit(code)

// }

func TestA(t *testing.T) {
	log.Println("TestA")
}

func TestB(t *testing.T) {
	log.Println("TestB")
}
