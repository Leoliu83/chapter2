package test1

import (
	"os"
	"runtime/pprof"
	"testing"
)

var f1 *os.File
var f2 *os.File

func init() {
	f1, _ = os.Create("D:\\MyProject\\Go\\gostudy\\basic\\src\\test1\\SliceTest_test.cpu.prof")
	f1.Truncate(0)
	f2, _ = os.Create("D:\\MyProject\\Go\\gostudy\\basic\\src\\test1\\SliceTest_test.heap.prof")
	f2.Truncate(0)
}

/*
	slice创建和array创建的性能测试
*/
func BenchmarkSliceTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreateSliceTest()
	}
}

func BenchmarkArrayTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreateArrayTest()
	}
}

const size int = 100000

// 1107	   1134451 ns/op	 6635539 B/op	       5 allocs/op
func sliceLoopAppend() {
	var sliceb = make([]int, size)
	for j := 0; j < size; j++ {
		sliceb = append(sliceb, j)
	}

}

// 5509	    240670 ns/op	  802818 B/op	       1 allocs/op
func sliceLoopIndex() {
	var sliceb = make([]int, size)
	for j := 0; j < size; j++ {
		sliceb[j] = j
	}
}

// 2290	    527569 ns/op	 1810435 B/op	       2 allocs/op
func sliceLoopCheck() {
	var sliceb = make([]int, size)
	var c int
	for j := 0; j < size+1; j++ {
		c = cap(sliceb)
		if j >= c {
			sliceb = append(sliceb, j)
		} else {
			sliceb[j] = j
		}

	}
}

var slicea = make([]int, size)

func TestSliceInsert(t *testing.T) {
	// pprof.StartCPUProfile(f1)
	for j := 0; j < size; j++ {
		slicea = append(slicea, j)
		// slicea[j] = j
	}
	// runtime.GC()
	pprof.WriteHeapProfile(f2)
	// pprof.StopCPUProfile()
}

func BenchmarkTestSliceInsert(b *testing.B) {
	for j := 0; j < b.N; j++ {
		// sliceLoopAppend()
		// sliceLoopIndex()
		sliceLoopCheck()
	}
}

var arraya [size]int

func arrayLoop() {
	var arrayb [size]int
	for j := 0; j < size; j++ {
		arrayb[j] = j
	}
}

func TestArrayInsert(t *testing.T) {
	// 14536	     82063 ns/op	       0 B/op	       0 allocs/op
	for j := 0; j < size; j++ {
		arraya[j] = j
	}
	pprof.WriteHeapProfile(f2)
}

func BenchmarkTestArrayInsert(b *testing.B) {
	for j := 0; j < b.N; j++ {
		arrayLoop()
	}
}
