##### Slice 注意事项
1. 慎用append
在很多代码中可以看到无论slice是否需要扩展，都会调用append
例如：
```go
var slicea = make([]int,100)
slicea= append(slicea, j)
```
这样做确实方便了很多，如果不需要扩展，则直接追加内容，如果需要扩展，则扩展新的slice
但是如果了解一下append
(函数原型在go/src/cmd/compile/internal/gc/cgen.go中，函数名为 cgen_append(n, res *Node) )
做一下append和下标赋值的性能可以发现，即使不扩展，append也会带来额外的开销。
下面这种用下标赋值的方式在不需要扩展slice时，和append拥有同样的效果，但开销远远的小于append
``` go
for j := 0; j < size; j++ {
	slicea[j] = j
}
```
当size=100000时候
benchmark结果如下：
``` go
// BenchmarkTestSliceInsert-8
// 1107	    1134451 ns/op	    6635539 B/op	    5 allocs/op
func sliceLoopAppend() {
	var sliceb = make([]int, size)
	for j := 0; j < size; j++ {
		sliceb = append(sliceb, j)
	}

}

// BenchmarkTestSliceInsert-8
// 5509	    240670 ns/op	    802818 B/op	    1 allocs/op
func sliceLoopIndex() {
	var sliceb = make([]int, size)
	for j := 0; j < size; j++ {
		sliceb[j] = j
	}
}
```
差距可以看的出来

再看下数组：
``` go
// BenchmarkTestSliceInsert-8
// 30484	    42177 ns/op	    0 B/op	    0 allocs/op
func arrayLoop() {
	var arrayb [size]int
	for j := 0; j < size; j++ {
		arrayb[j] = j
	}
}
```
数组的性能就更好了.
