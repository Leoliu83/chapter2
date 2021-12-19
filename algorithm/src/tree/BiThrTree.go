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

/*
	init中尽量不要放复杂逻辑，因为不好调试，出了类似死循环或者hang死的逻辑，比较难定位问题
*/
func init() {
	// InitQueueCnt()
}

/*
	根据树的高度计算队列长度
*/
func InitQueueCnt() {
	queuecnt := 0
	for i := (2 << MAXLEVEL); i != 0; i >>= 1 {
		// log.Println(queuecnt)
		queuecnt += i
	}
	log.Println("queuecnt: ", queuecnt)
	halfBranchLenList = CaculateHalfBranchLenth(MAXLEVEL)
	queue.Init(queuecnt)
}

var queue NodeQueue

/*
	线索二叉树定义
*/
type BiThrTreeNode struct {
	Data           int            // 数据
	lchild, rchild *BiThrTreeNode // 左右指针域
	ltype          int            // 左指针域类型，枚举值（Link，Thread）
	rtype          int            // 右指针域类型，枚举值（Link，Thread）
}

func CreateBiThrTreeNode(root *BiThrTreeNode, lv int) {
	if lv >= MAXLEVEL {
		return
	}
	// if root == nil {
	// 	return
	// }
	// if root.Data == 4 {
	// 	root.lchild = nil
	// 	root.rchild = nil
	// } else {
	// 	root.lchild = &BiThrTreeNode{Data: root.Data * 2}
	// 	root.rchild = &BiThrTreeNode{Data: root.Data*2 + 1}

	// }
	root.lchild = &BiThrTreeNode{Data: root.Data * 2}
	root.rchild = &BiThrTreeNode{Data: root.Data*2 + 1}
	// log.Printf("currentLevel: %d", lv)
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
				subRes = append(subRes, array[0].lchild.Data)
			}
			if array[0].rchild != nil {
				array = append(array, array[0].rchild)
				subRes = append(subRes, array[0].rchild.Data)
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
	打印二叉树，未完成
	两个枝之间只有一个空格，未处理
*/
func (root *BiThrTreeNode) Print(maxLevel int) {
	// 层数是从0开始，因此如果总层数是4，那么最高层数就是3，而不是4
	maxLevel = maxLevel - 1
	fmt.Println(halfBranchLenList)
	// 二叉树是标准的，的每层的结点树是可以计算出来的
	var nodeCount int = 1
	var currLevel int = 0
	var spaceCount, halfBranchLen int = 0, 0
	var str bytes.Buffer
	var formtStr string
	queue.Push(root)

	// idx 从1开始, 打印一个结点
	printOne := func(maxLv int, currentLv int, idx int, s string) {
		// fmt.Printf("currentLv: %d", currentLv)
		halfBranchLen = halfBranchLenList[maxLevel-currLevel]
		if currLevel == maxLevel {
			spaceCount = 0
		} else {
			spaceCount = halfBranchLen
		}
		if idx&1 == 1 { // 非2的倍数，也就是左枝
			// 左空格补齐(由于fmt.Sprintf不能处理0空格的情况，至少会出现一个空格，因此这里用循环)
			for i := 0; i < spaceCount; i++ {
				str.WriteString(" ")
			}
			// 例如底层的枝 ┌───┴───┐ halfBranchLen 就是 ┌─── 中的 ─ 的数量也就是3
			//
			// 向后补空格，halfBranchLen+1 因为除了空格，还必须包含1位数字
			formtStr = "%-" + fmt.Sprintf("%d", halfBranchLen+1) + "s"
			str.WriteString(fmt.Sprintf(formtStr, s))
		} else { // 如果是右枝
			// 向前补空格，halfBranchLen+2 因为除了空格以及┴所占用的1个字节，还必须包含1位数字
			formtStr = "%" + fmt.Sprintf("%d", halfBranchLen+2) + "s"
			str.WriteString(fmt.Sprintf(formtStr, s))
			// 向后补空
			formtStr = "%-" + fmt.Sprintf("%d", spaceCount+1) + "s"
			str.WriteString(fmt.Sprintf(formtStr, " "))
		}
	}

	for i := 1; !queue.isEmpty(); i++ {
		// log.Printf("max: %d,current: %d", maxLevel, currLevel)
		if currLevel > maxLevel {
			break
		}
		currentNode, ok := queue.Pop()
		// fmt.Printf("%+v \n", currentNode)
		if !ok {
			return
		}

		if currentNode == nil {
			printOne(maxLevel, currLevel, i, "-")
		} else {
			printOne(maxLevel, currLevel, i, fmt.Sprintf("%d", currentNode.Data))
		}

		if currentNode == nil {
			continue
		}
		queue.Push(currentNode.lchild)
		queue.Push(currentNode.rchild)
		if i == nodeCount {
			i = 0
			nodeCount = nodeCount * 2
			currLevel++
			fmt.Println(str.String())
			str.Reset()
			PrintBranch(maxLevel, currLevel)
		}
		// fmt.Printf("nodeCount: %d \n", nodeCount)
	}
	fmt.Println()

}

var halfBranchLenList []int

/*
	打印枝干
	也是就是树结构中的  ┌───┴───┐ 这些非数据部分
*/
func PrintBranch(maxLevel int, currLevel int) {
	if currLevel > maxLevel {
		return
	}
	nodeCount := 1 << currLevel
	idx := maxLevel - currLevel
	// log.Print(maxLevel, currLevel, idx)
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
