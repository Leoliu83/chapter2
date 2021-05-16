package test1

import "log"

type M int

/*
	方法也可以作为变量或者参数传递
	该方法测试的是当receiver是值的时候的情况
*/
func MethodAsParamReceiverIsValueTest() {
	var m M = 25
	log.Printf("MethodAsParamTest().m: %p, %d", &m, m)
	/*
		通过*类型*引用的方法表达式（method expression）会被还原成普通函数样式，recevier必须是第一参数，
		调用时需显示传递，类型可以是T 可以是 *T
	*/
	log.Println("----------- Trans as variable -------------")
	// T
	f1 := M.test
	f1(m)
	// *T
	f2 := (*M).test
	f2(&m)
	/*
		也可以直接以表达式方式调用
	*/
	M.test(m)
	(*M).test(&m)
	/*
		基于*实例*或者*指针*引用的方法值（method value），参数签名不变，依旧按照正常方式调用
	*/
	m++
	// 在（method value）方式作为变量传递时，先计算并复制 receiver *对象*（值），
	// ↓因此无论何时调用f3，m的值固定在了当前这个赋值的时间点,在这里为 m++ = 26
	f3 := m.test // 复制receiver对象（值）
	m++
	// 在（method value）方式作为参数传递时，先计算并复制 receiver *对象*（值），
	// ↓因此无论何时调用f4，m的值固定在了当前这个赋值的时间点,在这里为 m++ = 27
	f4 := (&m).test // 复制receiver对象（值）
	log.Printf("MethodAsParamTest().m: %p, %d", &m, m)

	f3() // <-- 26
	f4() // <-- 27
	log.Println("----------- Trans as parameter -------------")
	/*
		在作为参数传递时候, 会复制含receiver在内的整个（method value）
	*/
	// receiver为值
	m++
	call(m.test)
	m++
	call((&m).test)

	// receiver为指针
	m++
	call(m.testPtr)
	m++
	call((&m).testPtr)
}

/*
	方法也可以作为变量或者参数传递
	该方法测试的是当receiver是指针的时候的情况
*/
func MethodAsParamReceiverIsPointerTest() {
	var m M = 25
	log.Printf("MethodAsParamTest().m: %p, %d", &m, m)
	/*
		通过*类型*引用的方法表达式（method expression）会被还原成普通函数样式，recevier必须是第一参数，
		调用时需显示传递，类型可以是T 可以是 *T
	*/
	log.Println("----------- Trans as variable -------------")
	/*
		receiver是指针时，不能使用M.testPtr，必须使用*M.testPtr
		// T
		f1 := M.testPtr // error
		f1(m)
	*/
	// *T
	f2 := (*M).testPtr
	f2(&m)
	/*
		也可以直接以表达式方式调用
	*/
	(*M).testPtr(&m)
	/*
		基于*实例*或者*指针*引用的方法值（method value），参数签名不变，依旧按照正常方式调用
	*/
	m++
	// 在（method value）方式作为变量传递时，先计算并复制 receiver *指针*，
	// ↓因此无论何时调用f3，获取的都是调用前的最新值
	f3 := m.testPtr // 复制receiver指针
	m++
	// 在（method value）方式作为参数传递时，先计算并复制 receiver *指针*，
	// ↓因此无论何时调用f4，获取的都是调用前的最新值
	f4 := (&m).testPtr // 复制receiver指针
	log.Printf("MethodAsParamTest().m: %p, %d", &m, m)

	// 获取的是两次 m++ 后的值
	f3() // <-- 27
	f4() // <-- 27
	log.Println("----------- Trans as parameter -------------")
	/*
		在作为参数传递时候, 会复制含receiver在内的整个（method value）
	*/
	// receiver为指针
	m++
	call(m.testPtr)
	m++
	call((&m).testPtr)
	// 只要类型正确，nil同样可以调用
	log.Println("----------- nil call -------------")
	var p *M
	// p.test() // invalid memory address or nil pointer dereference, 因为test()的receiver不是指针
	p.testPtr() // no error
	// (*M)(nil).test() // invalid memory address or nil pointer dereference, 因为test()的receiver不是指针
	(*M)(nil).testPtr() // no error
	(*M).testPtr(nil)   // no error

}

func (m M) test() {
	log.Printf("test().m: %p, %d", &m, m)
}

func (m *M) testPtr() {
	if m != nil {
		log.Printf("*test().m: %p, %d", m, *m)
	} else {
		log.Println("nil run")
	}
}

func call(m func()) {
	m()
}
