package tree

import (
	"testing"
)

func TestMiddleOrderThread(t *testing.T) {
	InitQueueCnt()
	root := &BiThrTreeNode{Data: 1}
	t.Logf("%+v", root)
	CreateBiThrTreeNode(root, 1)
	root.Print(MAXLEVEL)
}

func TestPrintBranch(t *testing.T) {
	treeLevel := 3
	for i := 0; i <= treeLevel; i++ {
		PrintBranch(treeLevel, i)
	}

}

func TestNodeQueue(t *testing.T) {
	queue := &NodeQueue{}
	queue.Init(10)
	for i := 0; i < 5; i++ {
		queue.Push(&BiThrTreeNode{Data: i})
		queue.Print()
	}

	for i := 0; i < 10; i++ {
		queue.Pop()
		queue.Print()
	}
}
