package list

import (
	"fmt"
	"log"
	"testing"
)

var list LinkedList

/*
	在init()中初始化链表
*/
func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// 头结点
	list := LinkedList(&Node{data: nil})
	// 存放上一个节点，便于将当前节点赋值给上一个节点的next属性
	// string(rune(65 + i)) 表示将rune转换成字符
	var lastp *Node = list
	for i := 0; i < 5; i++ {
		p := Node{data: string(rune(65 + i))}
		lastp.next = &p
		lastp = &p
	}
}

func TestLinkedListCreateListHead(t *testing.T) {
	var list LinkedList
	list, ok := list.CreateListHead(7)
	if ok {
		list.Print()
	}
}

func TestLinkedListGetElem(t *testing.T) {
	list.Print()
	e, ok := list.GetElem(1)
	if ok {
		t.Logf("%+v", e)
	}
}

func TestLinkedListListInsert(t *testing.T) {
	list.Print()
	fmt.Println("------------------------------------")
	list.ListInsert(3, "new")
	list.Print()
}

func TestLinkedListListDelete(t *testing.T) {
	list.Print()
	fmt.Print("------------------------------")
	list.ListDelete(2)
	list.Print()
}
