package test1

import (
	"testing"
)

func BenchmarkChannelTest(b *testing.B) {
	for i := 1; i < b.N; i++ {
		ChannelSelectTest()
	}
}

func TestChannelTest(t *testing.T) {
	ChannelSelectTest()
}
