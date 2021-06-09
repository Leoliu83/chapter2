package stack

import (
	"log"
)

func Fbi(i int) {

}

/*
	For循环版的斐波那数列
*/
func FbiFor() {
	var y int
	x1 := 0
	x2 := 1
	for i := 0; i < 40; i++ {
		x1 = x2
		x2 = y
		y = x1 + x2
		log.Println(y)
	}
}

/*
	递归版本的斐波那数列
*/
func FbiRecursion(x1 int, x2 int, n int) {
	if n == 40 {
		return
	}
	n++
	log.Println(x2)
	FbiRecursion(x2, x1+x2, n)
}
