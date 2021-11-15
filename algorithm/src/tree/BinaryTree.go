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
	层序遍历（递归实现法）
	思想：将所有同一层的结点同时处理，因此递归的函数参数是一个同一层结点的数组.
         对于如下的树来说
                   A
         ┌─────────┴─────────┐
         B                   C
    ┌────┴────┐         ┌────┴────┐
    D         E         F         G
┌───┴───┐ ┌───┴───┐ ┌───┴───┐ ┌───┴───┐
H       I J       K L       M N       O
遍历A获取A下所有子节点，得到B,C作为下一次递归的参数
遍历B,C获取B,C所有子节点，得到D,E,F,G作为下一次递归的参数
遍历D,E,F,G获取D,E,F,G所有子节点，结果作为下一次递归的参数
...
以此类推
*/
func LevelOrder(root ...*BinaryTreeNode) {
	nilCnt := 0
	rootlen := len(root)
	// 对于二叉树来说，一定是每一层的结点树一定是上一层结点的2倍，2倍问题都可以用移位替代
	subNodes := make([]*BinaryTreeNode, rootlen<<1)
	if root == nil {
		return
	}

	for i, n := range root {
		if n == nil {
			nilCnt++
			continue
		}
		fmt.Printf("%s -> ", string(OFFSET+n.data))
		if n.left != nil {
			subNodes[i<<1] = n.left
		}
		if n.right != nil {
			subNodes[(i<<1)+1] = n.right
		}
	}
	if nilCnt == rootlen {
		return
	}
	// 尾递归，虽然golang没有尾递归优化，不过golang的栈空间可以达到2G
	LevelOrder(subNodes...)
}

/* 队列的简单实现 ---- BEGIN */
type QueueNode struct {
	data *BinaryTreeNode
	next *QueueNode
}
type InnerQueue struct {
	front *QueueNode // 首结点
	rear  *QueueNode // 尾结点
}

func (q *InnerQueue) InitQueue() {
	qn := &QueueNode{}
	q.front = qn
	q.rear = q.front
}

func (q *InnerQueue) Push(node *BinaryTreeNode) {
	qn := &QueueNode{data: node}
	q.rear.next = qn
	q.rear = qn
}

func (q *InnerQueue) Pop() {
	q.front = q.front.next
}

func (q *InnerQueue) isEmpty() bool {
	return q.front == q.rear
}

/* 队列的简单实现 ---- END */

/*
	层序遍历(队列实现法)
	思想，将每一层的结点放入队列，然后依次取出
         对于如下的树来说
                   A
         ┌─────────┴─────────┐
         B                   C
    ┌────┴────┐         ┌────┴────┐
    D         E         F         G
┌───┴───┐ ┌───┴───┐ ┌───┴───┐ ┌───┴───┐
H       I J       K L       M N       O
先将A结点Push进队列，则A为队列的front.next；因为这里用的队列的第一个元素为空元素
获取当前front.next(也就是B)的子节点B,C，依次放入队列，此时队列里是 A->B->C
将A结点Pop出队列，此时B变为front.next，此时队列里是B->C
获取当前front.next(也就是B)的子节点D,E，依次放入队列，此时队列里是B->C->D->E
将B结点Pop出队列，此时C变为front.next，此时队列里是C->D->E
获取当前front.next(也就是C)的子节点F,G，依次放入队列，此时队列里是C->D->E->F->G
将C结点Pop出队列，此时D变为front.next，此时队列里是D->E->F->G
...
以此类推
*/
func LevelOrderQueue(root *BinaryTreeNode) {
	var q InnerQueue
	q.InitQueue()
	q.Push(root)
	if root == nil {
		return
	}

	// 当队列不为空时，循环
	for !q.isEmpty() {
		fmt.Printf("%s -> ", string(OFFSET+q.front.next.data.data))
		if q.front.next.data.left != nil {
			q.Push(q.front.next.data.left)
		}
		if q.front.next.data.right != nil {
			q.Push(q.front.next.data.right)
		}
		q.Pop()
	}
	fmt.Printf("%s\n", "END")
}

/*
	TODO
	打印二叉树，未实现
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
