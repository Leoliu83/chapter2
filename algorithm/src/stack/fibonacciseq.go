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
	递归版本的斐波那数列 Fibonacci Sequence
	@param int 打印多少个数
*/
func FbiRecursion(cnt int) {
	i := 0
	// 定义一个递归函数
	var once func(a int, b int)
	once = func(a int, b int) {
		if i >= cnt {
			return
		}
		i++
		once(b, a+b)
	}
	// 递归调用，求Fibonacci Sequence
	once(0, 1)
}
