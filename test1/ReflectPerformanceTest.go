package test1

import (
	// "log"
	"reflect"
)

/*
	反射在带来“方便”的同时，也造成了很大的困扰。很多人对反射避之不及，
	因为认为它会造成很大的性能损失。
	由于反射的性能影响非常大，因此对于性能要求高的应用系统，尽可能不要使用反射
*/
type data105 struct {
	X int
}

func (x data105) Inc() {
	x.X++
}

var d data105

/*
	Running tool: D:\DeveloperEnv\google\Go\bin\go.exe test -benchmem -run=^$ -bench ^(BenchmarkReflectPerformanceTest1)$ golangStudy/test1

	goos: windows
	goarch: amd64
	pkg: golangStudy/test1
	BenchmarkReflectPerformanceTest1-8   	1000000000	         0.274 ns/op	       0 B/op	       0 allocs/op
	PASS
	ok  	golangStudy/test1	0.548s
*/
func Set105(x int) {
	d.X = x
}

/*
	Running tool: D:\DeveloperEnv\google\Go\bin\go.exe test -benchmem -run=^$ -bench ^(BenchmarkReflectPerformanceTest2)$ golangStudy/test1

	goos: windows
	goarch: amd64
	pkg: golangStudy/test1
	BenchmarkReflectPerformanceTest2-8   	15042022	        87.6 ns/op	      15 B/op	       1 allocs/op
	PASS
	ok  	golangStudy/test1	1.459s
*/
func Rset105_1(x int) {
	// 必须是指针
	// newd := new(data105)
	// %T可以显示变量的实际类型
	// log.Printf("%T", newd) // 等同于 &d
	// rd := reflect.ValueOf(&d).Elem()
	rd := reflect.ValueOf(&d).Elem()
	fx := rd.FieldByName("X")
	fx.Set(reflect.ValueOf(x))
}

// 将反射数据放到内存中，而不是放入方法栈
var rd = reflect.ValueOf(&d).Elem()
var fx = rd.FieldByName("X")

/*
	Running tool: D:\DeveloperEnv\google\Go\bin\go.exe test -benchmem -run=^$ -bench ^(BenchmarkReflectPerformanceTest3)$ golangStudy/test1

	goos: windows
	goarch: amd64
	pkg: golangStudy/test1
	BenchmarkReflectPerformanceTest3-8   	44959984	        26.9 ns/op	       8 B/op	       0 allocs/op
	PASS
	ok  	golangStudy/test1	1.323s
*/
func Rset105_2(x int) {
	fx.Set(reflect.ValueOf(x))
}

func ReflectPerformanceTest1() {
	Set105(1)
	Rset105_1(1)
	Rset105_2(1)
}

/*
	Running tool: D:\DeveloperEnv\google\Go\bin\go.exe test -benchmem -run=^$ -bench ^(BenchmarkReflectPerformanceTest4)$ golangStudy/test1

	goos: windows
	goarch: amd64
	pkg: golangStudy/test1
	BenchmarkReflectPerformanceTest4-8   	1000000000	         0.266 ns/op	       0 B/op	       0 allocs/op
	PASS
	ok  	golangStudy/test1	0.589s
*/
func Call105() {
	d.Inc()
}

/*
	Running tool: D:\DeveloperEnv\google\Go\bin\go.exe test -benchmem -run=^$ -bench ^(BenchmarkReflectPerformanceTest5)$ golangStudy/test1

	goos: windows
	goarch: amd64
	pkg: golangStudy/test1
	BenchmarkReflectPerformanceTest5-8   	 9250329	       139 ns/op	       8 B/op	       1 allocs/op
	PASS
	ok  	golangStudy/test1	1.749s
*/
var rd1 = reflect.ValueOf(&d).Elem()
var fx1 = rd1.MethodByName("Inc")

func Rcall105_1() {
	fx1.Call(nil)
}
