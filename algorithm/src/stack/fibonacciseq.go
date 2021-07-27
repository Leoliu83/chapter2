package stack

import (
	"log"
)

/*
	递归是栈运用的一种方式
	简单的说就是:
	    在*前行*阶段，对每一层的递归，函数局部变量，参数值，以及返回地址都被压入栈中。
	    在*退回*阶段，位于栈顶的局部变量，参数值，以及返回地址被弹出，用于返回调用层次中执行代码的其余部分
*/

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
