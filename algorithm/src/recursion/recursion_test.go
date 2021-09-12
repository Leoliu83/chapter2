package recursion

import "testing"

func BenchmarkTestFibonacciR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciR(0, 1, 0)
	}
}

func BenchmarkTestFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci()
	}
}

func TestFibonacciR(t *testing.T) {
	FibonacciR(0, 1, 0)
}

func TestFibonacci(t *testing.T) {
	Fibonacci()
}

// 393248 ns/op
//  69631 ns/op
