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
