package tree

import (
	"bytes"
	"fmt"
	"log"
	// "time"
)

/*
	枚举值，定义了树的结点左右指针域的类型
	Link = 0 表示类型为 "左孩子,右孩子"
	Thread = 1 表示类型为 "前驱,后继"
*/
const (
	Link = iota
	Thread
)

func init() {
	halfBranchLenList = CaculateHalfBranchLenth(10)
}

/*
	线索二叉树定义
*/
type BiThrTreeNode struct {
	data           int            // 数据
	lchild, rchild *BiThrTreeNode // 左右指针域
	ltype          int            // 左指针域类型，枚举值（Link，Thread）
	rtype          int            // 右指针域类型，枚举值（Link，Thread）
}

func CreateBiThrTreeNode(root *BiThrTreeNode, lv int) {
	if lv > MAXLEVEL {
		return
	}
	if root.data == 4 {
		root.lchild = nil
		root.rchild = nil
	} else {
		root.lchild = &BiThrTreeNode{data: root.data * 2}
		root.rchild = &BiThrTreeNode{data: root.data*2 + 1}

	}

	CreateBiThrTreeNode(root.lchild, lv+1)
	CreateBiThrTreeNode(root.rchild, lv+1)

}

/*
	中序遍历将二叉树线索化
	处理方式：
	如果线索化在处理叶子结点之前，则将处理写在 MiddleOrderThread 之前：
	doThread()
	MiddleOrderThread(root.lchild)
	如果线索化在处理叶子结点之后，则将处理写在 MiddleOrderThread 之后：
	MiddleOrderThread(root.lchild)
	doThread()

	doThread处理逻辑：
	if G的左叶子==nil {
		G的左叶子 = 刚处理的结点
	}
	if G的右叶子==nil {
		G的右叶子 = 下一个要处理的结点
	}
	问题点在于，如何获取 "刚处理的结点" 和 "下一个要处理的结点" 的地址
	在递归中，"下一个要处理的结点" 是在递归函数之后的处理操作


*/
func MiddleOrderThread(root *BiThrTreeNode) {
	if root == nil {
		return
	}
	MiddleOrderThread(root.lchild)
	fmt.Printf("%+v", root)
	MiddleOrderThread(root.rchild)
}

func (root *BiThrTreeNode) PrintNonQueue() [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	var array []*BiThrTreeNode = make([]*BiThrTreeNode, 0, 20)
	// 初始化根
	array = append(array, root)
	for len(array) > 0 {
		subRes := make([]int, 0, 2)
		if array[0] != nil {
			if array[0].lchild != nil {
				array = append(array, array[0].lchild)
				subRes = append(subRes, array[0].lchild.data)
			}
			if array[0].rchild != nil {
				array = append(array, array[0].rchild)
				subRes = append(subRes, array[0].rchild.data)
			}
			if len(subRes) > 0 {
				result = append(result, subRes)
			}
			array = array[1:]
		}
	}

	// fmt.Printf("%+v \n", result)
	return result

}

/*
	TODO
	打印二叉树，未实现
	两个枝之间只有一个空格，未处理
*/
func (root *BiThrTreeNode) Print(maxLevel int) {
	fmt.Println(halfBranchLenList)
	// 二叉树是标准的，的每层的结点树是可以计算出来的
	var nodeCount int = 1
	var currLevel int = 0
	var spaceCount, halfBranchLen int = 0, 0
	var str bytes.Buffer
	var queue NodeQueue
	queue.Init(20)
	queue.Push(root)
	for i := 1; !queue.isEmpty(); i++ {
		if currLevel > maxLevel {
			break
		}
		currentNode, ok := queue.Pop()
		// fmt.Printf("%+v \n", currentNode)
		if !ok {
			return
		}
		// fmt.Printf("currLevel: %d -> ", currLevel)
		halfBranchLen = halfBranchLenList[maxLevel-currLevel]
		if currLevel == maxLevel {
			spaceCount = 0
		} else {
			spaceCount = halfBranchLen
		}

		for j := 0; j < spaceCount; j++ {
			str.WriteString(" ")
		}
		if currentNode == nil {
			str.WriteString("-")
			for j := 0; j < halfBranchLen+1; j++ {
				str.WriteString(" ")
			}
			continue
		}

		str.WriteString(fmt.Sprintf("%d", currentNode.data))
		for j := 0; j < halfBranchLen+1; j++ {
			str.WriteString(" ")
		}
		queue.Push(currentNode.lchild)
		queue.Push(currentNode.rchild)
		if i == nodeCount {
			i = 0
			nodeCount = nodeCount * 2
			currLevel++
			fmt.Println(str.String())
			str.Reset()
		}
		// fmt.Printf("nodeCount: %d \n", nodeCount)
	}
	fmt.Println()

}

var halfBranchLenList []int

func PrintBranch(maxLevel int, currLevel int) {
	nodeCount := 1 << currLevel
	idx := maxLevel - currLevel
	halfBranchLen := halfBranchLenList[idx]
	var spaceCount int
	// 如果是最底层，则左边的空格数为1，否则左边的空格数等于
	if currLevel == maxLevel {
		spaceCount = 0
	} else {
		spaceCount = halfBranchLen
	}

	var str bytes.Buffer
	writeOneBranch := func() {
		for i := 0; i < spaceCount; i++ {
			str.WriteString(" ")
		}
		str.WriteString("┌")
		for i := 0; i < halfBranchLen; i++ {
			str.WriteString("─")
		}
		str.WriteString("┴")
		for i := 0; i < halfBranchLen; i++ {
			str.WriteString("─")
		}
		str.WriteString("┐")
		for i := 0; i < spaceCount+1; i++ {
			str.WriteString(" ")
		}
		fmt.Print(str.String())
		str.Reset()
	}

	for i := 0; i < nodeCount/2; i++ {
		writeOneBranch()
	}
	fmt.Println()
}

/*
	计算打印树枝 ┌────┴────┐ 中 '─' 数量的一半,并返回数组
	推导：
	0 表示最低层，打印叶子结点所需要的 ┌────┴────┐
	f(0) = 3
	f(1) = f(0)+1 // 倒数第2层
	f(2) = f(0)+1+f(1)+1 // 倒数第3层
	f(3) = f(0)+1+f(1)+1+f(2)+1 = f(2)+f(2)+1 // 倒数第4层
	f(4) = f(0)+1+f(1)+1+f(2)+1+f(3)+1 = f(3)+f(3)+1 // 倒数第5层
	......
	一直到 root 也就是根

	@param cnt 表示打印多少层
*/
func CaculateHalfBranchLenth(cnt int) []int {
	hbl := make([]int, cnt)
	hbl[0] = 3
	hbl[1] = 4
	for i := 2; i < cnt; i++ {
		// 等于 [前一个数]×2+1
		hbl[i] = hbl[i-1]<<1 + 1
	}
	// log.Println(hbl)
	return hbl
}

func PrintBinaryBad(bt *BiThrTreeNode) {
	fmt.Printf("%19s%d\n", "", bt.data)
	// level 1
	ln, rn := bt.lchild, bt.rchild
	// level 2
	// ┌───────── 长度为10，叶子结点长度的1/4
	fmt.Printf("%9s┌─────────┴─────────┐\n", "")
	// %19s是┌─────────┴─────────┐长度-2也就是去掉两头的┌和┐
	fmt.Printf("%9s%d%19s%d \r\n", "", ln.data, "", rn.data)
	// level 3
	// ┌──── 长度为5
	fmt.Printf("%4s┌────┴────┐%9s┌────┴────┐\r\n", "", "")
	// %9s是┌────┴────┐长度-2也就是去掉两头的┌和┐
	fmt.Printf("%4s%d%9s%d%9s%d%9s%d\r\n", "", ln.lchild.data, "", ln.rchild.data, "", rn.lchild.data, "", rn.rchild.data)
	// level 4
	// ┌─── 长度为4
	fmt.Printf("%0s┌───┴───┐%1s┌───┴───┐%1s┌───┴───┐%1s┌───┴───┐\r\n", "", "", "", "")
	fmt.Printf("%0s%-2d%5s%2d%1s%-2d%5s%2d%1s%-2d%5s%2d%1s%-2d%5s%2d \r\n",
		"", ln.lchild.lchild,
		"", ln.lchild.rchild,
		"", ln.rchild.lchild.data,
		"", ln.rchild.rchild.data,
		"", rn.lchild.lchild.data,
		"", rn.lchild.rchild.data,
		"", rn.rchild.lchild.data,
		"", rn.rchild.rchild.data,
	)

}

/************** 下面是结点队列的实现 *******************/

type NodeQueue struct {
	list  []*BiThrTreeNode
	front int
	rear  int
	max   int
	cnt   int
}

func (queue *NodeQueue) Init(capacity int) {
	queue.list = make([]*BiThrTreeNode, capacity)
	queue.max = cap(queue.list)
	queue.cnt = 0
	queue.front = 0
	queue.rear = 0
}
func (queue *NodeQueue) Push(node *BiThrTreeNode) {
	// 如果元素个数等于所能容纳的最大值
	if queue.isFull() {
		log.Println("Queue is full.")
		return
	}
	queue.list[queue.rear] = node
	// 获取rear的下一个索引位置
	queue.rear = (queue.rear + 1) % queue.max
	// 元素个数+1
	queue.cnt++
}

/*
	弹出第一个元素并返回元素地址
	这里不返回地址，小对象复制性能也不差，而且可以减少堆垃圾
*/
func (queue *NodeQueue) Pop() (node *BiThrTreeNode, ok bool) {
	// 如果元素个数等于0
	if queue.isEmpty() {
		log.Println("Queue is empty.")
		return nil, false
	}
	node = queue.list[queue.front]
	queue.front = (queue.front + 1) % queue.max
	queue.cnt--
	return node, true
}
func (queue *NodeQueue) isEmpty() bool {
	return queue.cnt <= 0
}

func (queue *NodeQueue) isFull() bool {
	return queue.cnt >= queue.max
}

func (queue *NodeQueue) Print() {
	start := queue.front % queue.max
	loop := queue.cnt + start
	fmt.Printf("Queue[")
	for i := start; i < loop; i++ {
		fmt.Printf("{idx: %d, %p},", i%queue.max, queue.list[i%queue.max])
	}
	fmt.Printf("]")
	fmt.Printf("{front: %d, rear: %d, cnt: %d}", queue.front, queue.rear, queue.cnt)
	fmt.Println()
}
