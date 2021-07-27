package string

import (
	"testing"
)

func TestBruteForce(t *testing.T) {
	s1 := "abcdefghijasscdeeabc"
	s2 := "cdee"
	idx := BruteForce(s1, s2)
	t.Log(idx)
}

func TestBruteForceWithIndex(t *testing.T) {
	s1 := "abcdefghijasscdeeabc"
	s2 := "cdee"
	BruteForceWithIndex(s1, s2)
}

func BenchmarkBruteForce(b *testing.B) {
	s1 := "abcdefghijasscdeeabc"
	s2 := "cdee"
	// idx := BruteForce(s1, s2)
	// b.Log(idx)
	for i := 0; i < b.N; i++ {
		BruteForce(s1, s2)
	}
}
