package test1

import (
	"testing"
)

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
