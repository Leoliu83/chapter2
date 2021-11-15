package queue

import (
	"fmt"
	"log"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedQueue struct {
	front *Node // 首结点
	rear  *Node // 尾结点
}

/*
	初始化空的队列时候，尾结点都指向首结点
*/
func (lq *LinkedQueue) InitQueue() {
	lq.front = &Node{}
	lq.rear = lq.front
}

/*
	插入结点流程：
	1. 将原rear的next指向新结点
	2. 将当前rear指向新结点
*/
func (lq *LinkedQueue) EnterQueue(i int) {
	n := Node{Data: i, Next: nil}
	lq.rear.Next = &n // 将当前尾结点指针指向新结点
	lq.rear = &n      // rear 指向新结点
}

/*
	删除结点流程：
	0. 判断是否为空，为空则打印队列为空，结束
	1. 将当前front指向下一个结点
	2. 将头结点指向当前front (若无头结点则无此步骤)
*/
func (lq *LinkedQueue) DeleteQueue() {
	if lq.front == lq.rear {
		log.Println("Queue is empty!")
		return
	}
	lq.front = lq.front.Next // 将下一个结点复制给当前结点
}

// 打印整个链表
func (lq *LinkedQueue) Print() {
	j := 1
	log.Println("----------------------------------------------")
	log.Printf("[HEAD]: %+v -> ", lq.front)
	for p := lq.front.Next; p != nil; p = p.Next {
		if j == 50 {
			break
		}
		fmt.Printf("[%d]: %+v -> ", j, p)
		j++
	}
	// for 循环中，p==nil时，j多加了1
	fmt.Println("END", "size: ", j-1)
}
