package test1

import (
	"testing"
)

func BenchmarkChannelTest(b *testing.B) {
	for i := 1; i < b.N; i++ {
		// ChannelSelectTest()
		// transOneBlock()
		transOneInt()
	}
}

func BenchmarkChannelTransPerformOneBlockTest(b *testing.B) {
	for i := 1; i < b.N; i++ {
		transOneBlock()
	}
}

func BenchmarkChannelTransPerformOneIntTest(b *testing.B) {
	for i := 1; i < b.N; i++ {
		transOneInt()
	}
}

func TestChannelTest(t *testing.T) {
	ChannelSelectTest()
}
