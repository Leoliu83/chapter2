package test1

import (
	"testing"
)

func BenchmarkMapTest(b *testing.B) {
	// for i := 0; i < b.N; i++ {
	// 	Performance1Test()
	// }
	for i := 0; i < b.N; i++ {
		Performance2Test()
	}
}
