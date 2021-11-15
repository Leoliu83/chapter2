package tree

import (
	"fmt"
	"testing"
)

func TestCreateBinaryTree(t *testing.T) {
	bt := CreateRoot(1)
	CreateBinaryTree(bt, 1)
	PrintBinaryTreeOrigin(bt, true)
	PrintBinaryTreeChar(bt, true)
}

func TestPreOrderBinaryTree(t *testing.T) {
	bt := CreateRoot(1)
	CreateBinaryTree(bt, 1)
	PrintBinaryTreeChar(bt, true)
	PreOrder(bt)
	fmt.Println()
}

func TestMiddleOrderBinaryTree(t *testing.T) {
	bt := CreateRoot(1)
	CreateBinaryTree(bt, 1)
	PrintBinaryTreeChar(bt, true)
	MiddleOrder(bt)
	fmt.Println()
}

func TestPostOrderBinaryTree(t *testing.T) {
	bt := CreateRoot(1)
	CreateBinaryTree(bt, 1)
	PrintBinaryTreeChar(bt, true)
	PostOrder(bt)
	fmt.Println()
}

func TestLevelOrderBinaryTree(t *testing.T) {
	bt := CreateRoot(1)
	CreateBinaryTree(bt, 1)
	PrintBinaryTreeChar(bt, true)
	LevelOrder(bt)
	fmt.Println()
}

func TestLevelOrderQueueBinaryTree(t *testing.T) {
	bt := CreateRoot(1)
	CreateBinaryTree(bt, 1)
	PrintBinaryTreeChar(bt, true)
	LevelOrderQueue(bt)
	fmt.Println()
}

func TestPrintBinaryTreeChar(t *testing.T) {
	type args struct {
		bt     *BinaryTreeNode
		isroot bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintBinaryTreeChar(tt.args.bt, tt.args.isroot)
		})
	}
}
