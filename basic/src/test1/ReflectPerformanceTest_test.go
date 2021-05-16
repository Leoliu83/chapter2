package test1

import "testing"

func BenchmarkReflectPerformanceTest1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Set105(i)
	}
}

func BenchmarkReflectPerformanceTest2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rset105_1(i)
	}
}

func BenchmarkReflectPerformanceTest3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rset105_2(i)
	}
}

func BenchmarkReflectPerformanceTest4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Call105()
	}
}

func BenchmarkReflectPerformanceTest5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rcall105_1()
	}
}
