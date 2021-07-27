package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	log.Println("Hello Go")
	test()
}

func test() {
	var f float64 = 1.1
	fb := fmt.Sprintf("%b", f)
	log.Println(fb)
	ui := math.Float64bits(f)
	log.Printf("%b,%d", ui, ui)
	log.Println(math.Pow(5, 2))
}
