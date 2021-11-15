package leetcode

/*
	原题：https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
	二叉树的层序遍历
	给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
*/
const (
	MAXLEVEL = 3
)

/*
	二叉树的定义
*/
type BiThrTreeNode struct {
	data           int            // 数据
	lchild, rchild *BiThrTreeNode // 左右指针域
}

/*
	初始化树
*/
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

/*************** 后面是解答 *******************/
/*
	递归实现，性能较差，后面有循环的实现，性能比对：
	递归：
	BenchmarkPrintNonQueueRecursion-8
    2288410	       548 ns/op	     576 B/op	      17 allocs/op
	循环：
	BenchmarkPrintNonQueueLoop-8
    2394727	       492 ns/op	     576 B/op	      17 allocs/op
	每次操作差了 50 ns
	注意递归中得 'return'：
	    没有返回值的 'return' 会直接跳出整个递归，结束执行
		而又返回值的 'return' 只会跳过一次递归调用
*/
func (bt *BiThrTreeNode) PrintNonQueueRecursion() [][]int {
	result := make([][]int, 0)
	if bt == nil {
		return result
	}
	var array []*BiThrTreeNode = make([]*BiThrTreeNode, 0, 20)
	var print func(bt *BiThrTreeNode)
	array = append(array, bt)
	print = func(bt *BiThrTreeNode) {
		subRes := make([]int, 0, 2)
		if bt != nil {
			if bt.lchild != nil {
				array = append(array, bt.lchild)
				subRes = append(subRes, bt.lchild.data)
			}
			if bt.rchild != nil {
				array = append(array, bt.rchild)
				subRes = append(subRes, bt.rchild.data)
			}
			if len(subRes) > 0 {
				result = append(result, subRes)
			}
		}
		array = array[1:]
		size := len(array)
		if size == 0 {
			return
		}
		print(array[0])
	}
	print(bt)
	// log.Println(result)
	return result
}

/*
    https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
	循环实现，后面有循环的实现，性能比对：
*/
func (root *BiThrTreeNode) PrintNonQueueLoop() [][]int {
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
