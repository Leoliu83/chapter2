package test1

import (
	"log"
	"runtime"
	"testing"
)

func BenchmarkGoRoutineTest(b *testing.B) {
	maxProcs := runtime.GOMAXPROCS(2)
	log.Printf("Max processes count is: %d", maxProcs)
	log.Printf("Current processes count is: %d", runtime.GOMAXPROCS(0))
	// GoRoutineParallelTest(ps, false)
}
