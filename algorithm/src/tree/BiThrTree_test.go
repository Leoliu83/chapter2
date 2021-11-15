package tree

import (
	"testing"
)

func TestMiddleOrderThread(t *testing.T) {
	root := &BiThrTreeNode{data: 1}
	CreateBiThrTreeNode(root, 1)
	PrintBinaryBad(root)
	root.Print(MAXLEVEL)
}

func TestPrintBranch(t *testing.T) {
	treeLevel := 5
	for i := 0; i <= treeLevel; i++ {
		PrintBranch(treeLevel, i)
	}

}

func TestNodeQueue(t *testing.T) {
	queue := &NodeQueue{}
	queue.Init(10)
	for i := 0; i < 5; i++ {
		queue.Push(&BiThrTreeNode{data: i})
		queue.Print()
	}

	for i := 0; i < 10; i++ {
		queue.Pop()
		queue.Print()
	}
}
