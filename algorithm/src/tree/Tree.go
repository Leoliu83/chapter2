package tree

import (
	"bytes"
	"fmt"
	"log"
	"sync"
)

/*
	通用树结构，Data为interface
*/
type TreeNode struct {
	_              struct{}    // 强制必须使用属性名进行初始化(放在首位，作为0字节)
	MaxLevel       int         // 最大层数 (8)
	lchild, rchild *TreeNode   // 左右指针域 (8,8)
	Data           interface{} // 数据 (interface是复核类型，由两个指针合成，因此是 8,8)
}

/*
	树结构
*/
type Tree struct {
	maxLevel  int
	firstNode *TreeNode
	queuePool sync.Pool
}

/*
	初始化树，包括最大层，头结点，队列
*/
func (tree *Tree) Init(root *TreeNode, maxLv int) {
	tree.maxLevel = maxLv
	tree.firstNode = root
	// 容纳树结点所需要的队列长度
	// queuecnt := 0
	// for i := (2 << maxLv); i != 0; i >>= 1 {
	// 	queuecnt += i
	// }
	// 最大队列数等于叶子结点数+1
	queuecnt := (2 << maxLv) + 1
	// 对象池，由于队列都是临时对象且可以反复重用，因此使用对象池，避免垃圾回收带来的影响
	tree.queuePool = sync.Pool{
		New: func() interface{} {
			return InitQueue(queuecnt)
		},
	}
}

/*
	初始化样例树
*/
func (tree *Tree) CreateSampleTree(root *TreeNode, lv int) {
	if lv >= tree.maxLevel {
		return
	}
	root.lchild = &TreeNode{Data: root.Data.(int) * 2}
	root.rchild = &TreeNode{Data: root.Data.(int)*2 + 1}
	// log.Printf("currentLevel: %d", lv)
	tree.CreateSampleTree(root.lchild, lv+1)
	tree.CreateSampleTree(root.rchild, lv+1)

}

/*
	打印二叉树结构
*/
func (tree *Tree) Print() {
	halfBranchLenList = CaculateHalfBranchLenth(tree.maxLevel)
	// 层数是从0开始，因此如果总层数是4，那么最高层数就是3，而不是4
	maxLevel := tree.maxLevel - 1
	// 二叉树是标准的，的每层的结点树是可以计算出来的
	// 变量说明：
	//   1. nodeCount     结点数量
	//   2. currLevel     当前树的层数
	//   3. spaceCount    结点与结点之间空格的数量
	//   4. halfBranchLen 就是 ┌─── 中的 ─ 的数量
	var nodeCount, currLevel, spaceCount, halfBranchLen int = 1, 0, 0, 0
	var str bytes.Buffer
	var formtStr string
	treeNodeQueue := tree.queuePool.Get().(*treeNodeQueueInternal)
	defer tree.queuePool.Put(treeNodeQueue)

	treeNodeQueue.Push(tree.firstNode)

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

	for i := 1; !treeNodeQueue.isEmpty(); i++ {
		// log.Printf("max: %d,current: %d", maxLevel, currLevel)
		if currLevel > maxLevel {
			break
		}
		currentNode, ok := treeNodeQueue.Pop()
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
		treeNodeQueue.Push(currentNode.lchild)
		treeNodeQueue.Push(currentNode.rchild)
		if i == nodeCount { // 如果i等于当前层结点的数量，表示当前层遍历完了
			i = 0
			nodeCount = nodeCount * 2 // 当前结点数量乘以2，因为每一层结点的数量，就是上一层结点数×2
			currLevel++               // 层数+1
			fmt.Println(str.String()) // 打印当前字符串缓冲中的数据
			str.Reset()               // 重置字符串缓存
			PrintBranch(maxLevel, currLevel)
		}
		// fmt.Printf("nodeCount: %d \n", nodeCount)
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

/************** 下面是结点队列的实现 *******************/
/*
	树结点队列变量
	内部使用，不可export
*/
// var treeNodeQueue treeNodeQueueInternal

type treeNodeQueueInternal struct {
	list  []*TreeNode
	front int
	rear  int
	max   int
	cnt   int
}

/*
	初始化结点队列
	@pram capacity:int 队列最大长度
*/
func InitQueue(capacity int) treeNodeQueueInternal {
	treeNodeQueue := treeNodeQueueInternal{}
	treeNodeQueue.list = make([]*TreeNode, capacity)
	treeNodeQueue.max = cap(treeNodeQueue.list)
	treeNodeQueue.cnt = 0
	treeNodeQueue.front = 0
	treeNodeQueue.rear = 0
	return treeNodeQueue
}

/*
	将结点推入队列
	@pram node:*TreeNode 结点指针
*/
func (treeNodeQueue *treeNodeQueueInternal) Push(node *TreeNode) {
	// 如果元素个数等于所能容纳的最大值
	if treeNodeQueue.isFull() {
		log.Println("[Warning]: Queue is full.")
		log.Printf("%d , %+v \n", treeNodeQueue.cnt, treeNodeQueue.list)
		return
	}
	treeNodeQueue.list[treeNodeQueue.rear] = node
	// 获取rear的下一个索引位置
	treeNodeQueue.rear = (treeNodeQueue.rear + 1) % treeNodeQueue.max
	// 元素个数+1
	treeNodeQueue.cnt++
}

/*
	弹出第一个元素并返回元素地址
	TODO: 考虑不返回地址，小对象复制性能也不差，而且可以减少堆垃圾，降低垃圾回收的成本，对象逃逸会产生堆垃圾
	@param node:*TreeNode 结点指针
	@param ok:bool 弹出出是否成功
*/
func (treeNodeQueue *treeNodeQueueInternal) Pop() (node *TreeNode, ok bool) {
	// 如果元素个数等于0
	if treeNodeQueue.isEmpty() {
		log.Println("[Warning]: Queue is empty.")
		return nil, false
	}
	node = treeNodeQueue.list[treeNodeQueue.front]
	treeNodeQueue.front = (treeNodeQueue.front + 1) % treeNodeQueue.max
	treeNodeQueue.cnt--
	return node, true
}

/*
	判断是否为空队列
*/
func (treeNodeQueue *treeNodeQueueInternal) isEmpty() bool {
	return treeNodeQueue.cnt <= 0
}

/*
	判断队列是否已满
*/
func (treeNodeQueue *treeNodeQueueInternal) isFull() bool {
	return treeNodeQueue.cnt >= treeNodeQueue.max
}

/*
	打印队列
*/
func (treeNodeQueue *treeNodeQueueInternal) Print() {
	var buffer bytes.Buffer
	start := treeNodeQueue.front % treeNodeQueue.max
	loop := treeNodeQueue.cnt + start
	buffer.WriteString("Queue[")
	for i := start; i < loop; i++ {
		buffer.WriteString(fmt.Sprintf("{idx: %d, %p},", i%treeNodeQueue.max, treeNodeQueue.list[i%treeNodeQueue.max]))

	}
	buffer.WriteString("]")
	buffer.WriteString(fmt.Sprintf("{front: %d, rear: %d, cnt: %d}", treeNodeQueue.front, treeNodeQueue.rear, treeNodeQueue.cnt))
	log.Println(buffer.String())
}
