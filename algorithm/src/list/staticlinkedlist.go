package list

type Component struct {
	data Element
	cur  int
}

type StaticLinkedList [10]Component

/*将一维数组space中各个分量链成一个备用链表*/
/*space[0].cur 为头指针，"0"表示空指针*/
func (list *StaticLinkedList) InitList() bool {
	for i := 0; i < 10; i++ {
		list[i].cur = i + 1
	}
	list[9].cur = 0 /*目前静态列表为空*/
	return true
}

/*
	分配空闲空间
*/
func (list *StaticLinkedList) Malloc_SLL(i int, c Component) (int, bool) {
	p := list[0].cur // 获取空闲空间
	if p == 0 {      // 如果为0 则表示没有空间了
		list[0].cur = list[p].cur
		return p, false
	}
	return p, true
}
