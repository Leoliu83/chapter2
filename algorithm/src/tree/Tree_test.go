package tree

import (
	"testing"
	"unsafe"
)

func TestStruct(t *testing.T) {
	node := TreeNode{}
	size := unsafe.Sizeof(node)
	alig := unsafe.Alignof(node)
	t.Logf("Size: %d, Align: %d", size, alig)
	var a interface{}
	t.Logf("Size: %d", unsafe.Sizeof(a))
}

func TestTreeInit(t *testing.T) {
	var tree Tree
	tree.Init(&TreeNode{Data: 1}, 4)
	t.Logf("%+v", tree)
	tree.CreateSampleTree(tree.firstNode, 1)
	tree.Print()
}
