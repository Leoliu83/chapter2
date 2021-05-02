package test1

import "log"

type tree struct {
	v int
}

// 如果t为nil则会抛出 panic: runtime error: invalid memory address or nil pointer dereference
func (t *tree) Sum() int {
	return t.v
}

// 如果方法内部，不调用t.v，则即使t为nil，方法也是可以被调用的
// tree指针类型的nil和java的null不一样，前者可以调用方法，而后者会抛出空指针异常
func (t *tree) PrintSth() {
	log.Println("PrintSth: ", t)
}

// 如果t的receiver是值，go会将该类型的属性自动初始化，例如这里tree中得v就会初始化为0
func (t tree) PrintSthV() {
	log.Println("PrintSthV: ", t)
}

func NilTest1() {
	var t1 *tree
	var t2 tree
	// 调用正常，打印<nil>
	t1.PrintSth()
	// 调用正常，打印{0}
	t2.PrintSthV()
	// 抛出panic
	t1.Sum()
}
