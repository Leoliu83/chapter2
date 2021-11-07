package tree

import (
	"fmt"
	// "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

var MAXLEVEL int = 3
var OFFSET rune = 64

/*
	标注每个属性的字节，以保证内存对齐情况下，结构体内存使用最优
*/
type BinaryTreeNode struct {
	// 这里使用rune，因为可以作为int处理，又可以直接作为ascii转成字符串
	data  rune            // 8 byte
	left  *BinaryTreeNode // 8 byte
	right *BinaryTreeNode // 8 byte
}

/*
	创建根结点
	@param data 根节点的数据值
	@return 返回一个BinaryTreeNode指针
*/
func CreateRoot(data rune) *BinaryTreeNode {
	return &BinaryTreeNode{data: data}
}

/*
	初始化一棵二叉树树，默认是满二叉树
	@param bt 根结点
	@param level 根属于level几，可以是0，也可以是1
*/
func CreateBinaryTree(bt *BinaryTreeNode, level int) {
	if level > MAXLEVEL {
		return
	}
	// 由于是满二叉树，因此按顺序编号相对比较容易，有一定的规律
	bt.left = &BinaryTreeNode{data: bt.data * 2}
	bt.right = &BinaryTreeNode{data: bt.data*2 + 1}
	CreateBinaryTree(bt.left, level+1)
	CreateBinaryTree(bt.right, level+1)
}

/*
	前序遍历
*/
func PreOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%s -> ", string(OFFSET+root.data))
	PreOrder(root.left)
	PreOrder(root.right)
}

/*
	中序遍历
*/
func MiddleOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	MiddleOrder(root.left)
	fmt.Printf("%s -> ", string(OFFSET+root.data))
	MiddleOrder(root.right)
}

/*
	后续遍历
*/
func PostOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	PostOrder(root.left)
	PostOrder(root.right)
	fmt.Printf("%s -> ", string(OFFSET+root.data))
}

/*
	层序遍历
*/
func LevelOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%s -> ", string(OFFSET+root.left.data))
	fmt.Printf("%s -> ", string(OFFSET+root.right.data))
	LevelOrder(root.left)
	LevelOrder(root.right)
	// fmt.Printf("%s -> ", string(OFFSET+root.data))
}

/*
	暂时写死，作测试使用，后续修改成动态打印
	只写了3层打印，主要是一些二叉树的操作等
*/
func PrintBinaryTreeOrigin(bt *BinaryTreeNode, isroot bool) {
	fmt.Printf("%19s%d\n", "", bt.data)
	// level 1
	ln, rn := bt.left, bt.right
	// level 2
	// ┌───────── 长度为10，叶子结点长度的1/4
	fmt.Printf("%9s┌─────────┴─────────┐\n", "")
	// %19s是┌─────────┴─────────┐长度-2也就是去掉两头的┌和┐
	fmt.Printf("%9s%d%19s%d \r\n", "", ln.data, "", rn.data)
	// level 3
	// ┌──── 长度为5
	fmt.Printf("%4s┌────┴────┐%9s┌────┴────┐\r\n", "", "")
	// %9s是┌────┴────┐长度-2也就是去掉两头的┌和┐
	fmt.Printf("%4s%d%9s%d%9s%d%9s%d\r\n", "", ln.left.data, "", ln.right.data, "", rn.left.data, "", rn.right.data)
	// level 4
	// ┌─── 长度为4
	fmt.Printf("%0s┌───┴───┐%1s┌───┴───┐%1s┌───┴───┐%1s┌───┴───┐\r\n", "", "", "", "")
	fmt.Printf("%0s%-2d%5s%2d%1s%-2d%5s%2d%1s%-2d%5s%2d%1s%-2d%5s%2d \r\n",
		"", ln.left.left.data,
		"", ln.left.right.data,
		"", ln.right.left.data,
		"", ln.right.right.data,
		"", rn.left.left.data,
		"", rn.left.right.data,
		"", rn.right.left.data,
		"", rn.right.right.data,
	)

}

/*
	暂时写死，作测试使用，后续修改成动态打印
	只写了3层打印，主要是一些二叉树的操作等
*/
func PrintBinaryTreeChar(bt *BinaryTreeNode, isroot bool) {
	fmt.Printf("%19s%s\n", "", string(OFFSET+bt.data))
	// level 1
	ln, rn := bt.left, bt.right
	// level 2
	// ┌───────── 长度为10，叶子结点长度的1/4
	fmt.Printf("%9s┌─────────┴─────────┐\n", "")
	// %19s是┌─────────┴─────────┐长度-2也就是去掉两头的┌和┐
	fmt.Printf("%9s%s%19s%s \r\n",
		"", string(OFFSET+ln.data),
		"", string(OFFSET+rn.data))
	// level 3
	fmt.Printf("%4s┌────┴────┐%9s┌────┴────┐\r\n", "", "")
	fmt.Printf("%4s%s%9s%s%9s%s%9s%s\r\n",
		"", string(OFFSET+ln.left.data),
		"", string(OFFSET+ln.right.data),
		"", string(OFFSET+rn.left.data),
		"", string(OFFSET+rn.right.data))
	// level 4
	fmt.Printf("%0s┌───┴───┐%1s┌───┴───┐%1s┌───┴───┐%1s┌───┴───┐\r\n", "", "", "", "")
	fmt.Printf("%0s%-2s%5s%2s%1s%-2s%5s%2s%1s%-2s%5s%2s%1s%-2s%5s%2s \r\n",
		"", string(OFFSET+ln.left.left.data),
		"", string(OFFSET+ln.left.right.data),
		"", string(OFFSET+ln.right.left.data),
		"", string(OFFSET+ln.right.right.data),
		"", string(OFFSET+rn.left.left.data),
		"", string(OFFSET+rn.left.right.data),
		"", string(OFFSET+rn.right.left.data),
		"", string(OFFSET+rn.right.right.data),
	)

}
