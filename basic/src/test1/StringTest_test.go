package test1

import "testing"

/*
	测试文件名必须以 原文件名_test.go 构成
	函数必须以 Benchmark 开头
*/
func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// StringPlusTest()
		// StringJoinTest()
		StringByteBufTest()
	}
}
