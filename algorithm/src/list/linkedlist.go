package list

import (
	"fmt"
	"log"
)

/*
	定义结点
*/
type Node struct {
	data Element
	next *Node
}

/*
	这里可以使用 type LinkedList = Node
	receiver 就可以使用 *LinkedList
*/
type LinkedList = *Node

/*
	获取元素
*/
func (linkedList LinkedList) GetElem(i int) (Element, bool) {
	var j int = 1
	var p *Node = linkedList.next
	// 当p不为空且j<i时候循环。p为nil，则说明是链表最后一个元素
	// j<i 表示执行i-1次循环，第i-1次时，p为第i个元素
	for p != nil && (j < i) {
		p = p.next
		j++
	}
	// p==nil 说明整个链表扫描完成之后没有扫描到指定位置元素
	// j>i 说明循环次数已经超过了i (原书中有该掉件，但在这里去掉，因为j不可能大于i，j只可能等于i)
	if p == nil /*|| j > i*/ {
		log.Printf("[ERROR]: Invalid index: %d, list size: %d", i, j)
		return nil, false
	}
	return p.data, true
}

/*
	原文代码中的参数i表示在第i个元素之后插入元素
	这里按照原文逻辑实现
	我理解的是在i位置插入元素，而将原来的元素后移，逻辑如下图（0表示头元素）：
	+---+---+---+---+
	| 0 | 1 | 2 | 3 |
	+---+---+---+---+
	ListInsert(2,e)
	+---+---+     +---+---+
	| 0 | 1 |     | 2 | 3 |
	+---+---+     +---+---+
	           ↑
		     +---+
		     | e |
		     +---+
*/
func (linkedList LinkedList) ListInsert(i int, e Element) bool {
	var j int = 1
	// 1. 声明一结点p指向链表第一个节点，初始化j从1开始
	var p *Node = linkedList.next
	// 当p不为nil（也就是尾节点），如果按照我的逻辑，这里要改成 j<i-1
	// 2. 当$j<i$时，就遍历链表，让p的指针向后移动，不断指向下一节点，j++
	for p != nil && j < i {
		p = p.next
		j++
	}
	// 3. 若到链表末尾p为空，则说明第i个元素不存在
	if p == nil || j > i {
		log.Printf("[ERROR]: Invalid index: %d, list size: %d", i, j-1)
		return false
	}
	// 4. 否则查找成功，在系统中生成一个空节点s；
	// 5. 将元素e赋值给s->data
	s := Node{data: e}
	// 6. 单链表的插入标准语句 s->next=p->next; p->next=s
	s.next = p.next
	p.next = &s
	// 7. 返回成功
	return true

}

/*
	删除元素
	原文代码中的参数i表示删除第i个元素后的元素
	这里按照原文逻辑实现
*/
func (linkedList LinkedList) ListDelete(i int) (Element, bool) {
	// 1. 声明一结点p指向链表第一个节点，初始化j从1开始
	p := linkedList.next
	j := 1
	// 2. 当$j<i$时，就遍历链表，让p的指针向后移动，不断指向下一个节点，j累加1
	for p.next != nil && j < i {
		p = p.next
		j++
	}
	// 3. 若到链表末尾p为空，则说明第i个元素不存在
	if p.next == nil || j > i {
		log.Printf("[ERROR]: Invalid index: %d, list size: %d, index valide range [%d,%d]", i, j, 0, j-1)
		return nil, false
	}
	// 4. 否则查找成功，将欲删除的结点p->next赋值给q；
	q := p.next
	// 5. 单链表的删除标准语句p->next = q->next
	p.next = q.next
	// 6. 将q结点中得数据赋值给e，作为返回；
	e := q.data
	// 7. 释放q结点(在这里栈内存执行完成后自动释放)
	// 8. 返回成功
	return e, true
}

/*
	初始化一个链表，数据从ascii:65开始
	这里使用的是插队法
	也就是每一次都将元素插入到头结点的后面：
	p.next = L.next
	L.next = &p
	也可以使用尾插法，也就是每一次都将元素插入上一个元素的后面
	定义一个指向尾部元素的变量 var lastp *Node = L
	for循环中
	lastp.next = &p
	lastp = &p
*/
func (linkedList LinkedList) CreateListHead(n int) (LinkedList, bool) {
	// 如果当前链表结构体不为nil，说明已经初始化了，则不需要再初始化
	if linkedList != nil {
		log.Println("[WARNINT]: Current linkedList is not nil.")
		linkedList.Print()
		return nil, false
	}
	// 1. 声明一结点p和计数器变量
	// 2. 初始化一空链表L；
	// 3. 让L的头结点的指针指向NULL,即建立一个带头结点的单链表
	L := LinkedList(&Node{data: nil})
	var i int
	for i = 0; i < n; i++ {
		p := Node{data: string(rune(65 + i))}
		p.next = L.next
		L.next = &p
	}
	return L, true

}

/*
	单链表整表删除
	这里变量q的意义：
	如果没有变量q，直接将p设置为nil，那么将无法找到下一个节点
*/
func (linkedList LinkedList) ClearList() {
	var q *Node
	var p *Node
	p = linkedList.next
	for p != nil {
		q = p.next
		p = nil
		p = q
	}
}

// 打印整个链表
func (linkedList LinkedList) Print() {
	j := 1
	log.Println("----------------------------------------------")
	fmt.Printf("[HEAD]: %+v -> ", linkedList)
	for p := linkedList.next; p != nil; p = p.next {
		if j == 50 {
			break
		}
		fmt.Printf("[%d]: %+v -> ", j, p)
		j++
	}
	// for 循环中，p==nil时，j多加了1
	fmt.Println("END", "size: ", j-1)
}
