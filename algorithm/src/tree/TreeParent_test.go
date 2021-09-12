package tree

import "testing"

func BenchmarkTree(b *testing.B) {
	a := 1
	for i := 0; i < b.N; i++ {
		a = a + 1
	}
}

func addOne(i int) int {
	return i + 1
}

func BenchmarkTest2(b *testing.B) {
	a := 1
	for i := 0; i < b.N; i++ {
		a = addOne(a)
	}
}
