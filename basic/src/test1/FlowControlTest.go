package test1

import (
	"errors"
	"log"
	"strconv"
)

/*
	FlowControlTestBad 中，if块承担了两种逻辑，错误处理逻辑和正常操作逻辑
	基于重构原则，应该保持代码块功能单一性，因此该方法中得实现是比较糟糕的
*/
func FlowControlTestBad() {
	x := 100
	if err := check(x); err == nil {
		x++
		log.Printf("%d \n", x)
	} else {
		log.Fatalln(err)
	}
}

/*
	FlowControlTestGood 中，if块承担了单一逻辑，也就是错误逻辑
	剩余的正常逻辑都在一个层中，而且单一功能可以提升代码可维护性，利于拆分重构
	* 尽可能使正常逻辑处在同一层
*/
func FlowControlTestGood() {
	x := 100
	if err := check(x); err != nil {
		log.Fatalln(err)
	}

	x++
	log.Printf("x++")

}

/*
	将 FlowControlTestComplicateBad 中得复杂逻辑进行了封装
	封装成了 flowControlTestComplicateCheck
*/
func FlowControlTestComplicateGood() {
	s := "9"
	if err := flowControlTestComplicateCheck(s); err != nil {
		log.Fatalln(err)
	}
	log.Println("ok")
}

/*
	if中的逻辑过于复杂，建议重构成函数
*/
func FlowControlTestComplicateBad() {
	s := "9"

	if n, err := strconv.ParseInt(s, 10, 64); err != nil || n < 0 || n > 10 || n%2 != 0 {
		log.Fatalln(err)
	}

	log.Println("ok")

}

/*
	将复杂逻辑封装成函数
*/
func flowControlTestComplicateCheck(s string) error {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil || n < 0 || n > 10 || n%2 != 0 {
		// 不可以写 Invlide number!
		// 1. 不可以包含大写字母
		// 2. 不可以包含标点符号
		return errors.New("invlid number")
	}

	return nil

}

/*

 */
func check(x int) error {
	if x > 10 {
		return errors.New("x>10")
	}

	return nil
}

/*
	switch同样支持表达式，而且支持常量和非常量值，因此变得非常灵活
	switch初始化语句从上到下，从左到右执行，最后执行default
	即使default放在最开始，也会在最后执行，建议放在最后
	不能出现重复的case值
	无需显式的写break;自动中断
*/
func SwitchTest(x int) {
	a, b, c, d, e := 1, 2, 3, 4, 5
	var dont bool = true
	switch y := 10; y - x {
	default:
		log.Println("unknow")
	case a, b: // a 或者 b，两者满足其一即可
		log.Println("ok")
	case c:
		log.Println("warning")
	case d, e:
		log.Println("error")
		if dont { //如果进入这一句，则break直接跳出case，不会执行fallthrough
			break
		}
		// 表示继续执行下一个case 6, 不需要匹配case条件表达式
		// fallthrough 必须放在case块尾，可以被break中断
		fallthrough
	case 6: // 单条件，内容为空时，隐式“case 6: break;”
		// fallthrough // cannot fallthrough final case in switch

	}

}

func ForTest() {
	data := [3]int{1, 2, 3}
	log.Printf("%p , %d", &data[0], data[0])
	log.Printf("%p , %d", &data[1], data[1])
	log.Printf("%p , %d", &data[2], data[2])
	/*
		for range 会从data的复制品种取值，也就是说k是data[i]的副本，
		对k的加减并不影响data[i]，对data[i]的变化也不影响k，
		而且k永远都是一个地址，该地址数据会不停的变化，所以 for循环结束后，k的值永远都是最后一次循环的值
	*/
	for i, k := range data {
		data[i] += 100
		log.Printf("%p , %d -> %p , %d", &data[i], data[i], &k, k)
	}
	log.Println("=============================================")
	/*
		使用data[:]，则由for range一个数组，变成了for range一个slice，有本质区别
		查看for range 源码，for range 语法糖是先将 range 的变量赋值给 for_temp （值传递）再进行循环
		[for range for slice]
		// The loop we generate:
		//   for_temp := range
		//   len_temp := len(for_temp)
		//   for index_temp = 0; index_temp < len_temp; index_temp++ {
		//           value_temp = for_temp[index_temp]
		//           index = index_temp
		//           value = value_temp
		//           original body
		//   }
		则k只复制slice，而不复制底层数组
		[slice结构体]
		type slice struct {
			array unsafe.Pointer
			len   int
			cap   int
		}
		也就是说把slice中的数组指针array，len值，cap值都复制了，数组只是指针，
		由于数组是指针，因此使用的相当于是指针传递，访问的还是原数组
		数组就不同了，数组赋值是把数组中所有的元素赋值给了新的变量,见后面的【数组赋值】
		data[:] means `creates a slice` including elements 0 through len(data) of data
		data[:] 等价于 data[0:len(data)]
	*/
	for i, k := range data[:] {
		data[i] += 100
		log.Printf("%p , %d -> %p , %d", &data[i], data[i], &k, k)
	}

	log.Println("=============================================")
	/*
		range 中得函数表达式只执行一次
		dataArray 只执行一次
	*/
	for i, k := range dataArray() {
		data[i] += 100
		log.Printf("%p , %d -> %p , %d", &data[i], data[i], &k, k)
	}
	// 【数组赋值】
	data1 := data
	log.Printf("data1[0]: %p , data[0]: %p", &data1[0], &data[0])
	log.Printf("data1[1]: %p , data[1]: %p", &data1[1], &data[1])
	log.Printf("data1[2]: %p , data[2]: %p", &data1[2], &data[2])
}

func dataArray() []int {
	return []int{1, 2, 3}
}

/*
	goto的label有以下特点：
	1. 未使用的label会引发编译错误
	2. 不能跳转到其他函数或者内层代码块
	break 可以指定label，中断指定label的for循环
	continue 可以指定label，跳过指定label的for循环

*/
func GotoTest() {
	for i := 0; i < 100; i++ {
		if i == 50 {
			goto label1
		}
		if i > 50 {
			goto label2
		}
	}
label1:
	log.Println("50")
label2:
	log.Println("exit")
	// 未使用的label（label3）会引发编译错误
	// label3:
	// 	log.Println("label3")
}
