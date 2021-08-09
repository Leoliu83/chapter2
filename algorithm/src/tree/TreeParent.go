package tree

/*
	树的双亲表示法
	在每个结点中，附设一个指针指向其双亲结点到链表中得位置，在这里是数组下标
*/
/*
	结点的数据类型
*/
type dataType int

/*
	定义树的结点结构体
*/
type TreeNode struct {
	data   dataType // 结点数据
	parent int      // 结点双亲所在的数组下标
	_      struct{}
}

/*
	定义树结构体
*/
type Tree struct {
	maxsize int        // 最大结点数量
	nodes   []TreeNode // 结点数组
	r, n    int        // r表示根的位置，n表示结点数
}
