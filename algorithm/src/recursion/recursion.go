package recursion

/*
	递归的斐波那契数列
*/
func FibonacciR(a int, b int, i int) {
	// log.Println(a)
	if i == 100000 {
		return
	}
	FibonacciR(b, a+b, i+1)
}

/*
	for循环的斐波那契数列
*/
func Fibonacci() {
	a := 0
	b := 1
	b1 := 0
	for i := 0; i <= 100000; i++ {
		b1 = b
		b = a + b
		a = b1
		// log.Println(a)
	}
}
