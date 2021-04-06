package test1

import (
	"fmt"
	"testing"
	"time"
)

// Test variable
func Test(t *testing.T) {
	timeRunStart := time.Now().UnixNano()
	fmt.Println(timeRunStart)
	time.Sleep(time.Duration(1) * time.Second)
	x, y := 1, 2
	a, b := x+1, y+1
	fmt.Println(a, ":", b)
	timeRunEnd := time.Now().UnixNano()
	fmt.Println(timeRunEnd)
	fmt.Printf("Run time: %d ms \r\n", (timeRunEnd-timeRunStart)/int64(time.Millisecond))
}
