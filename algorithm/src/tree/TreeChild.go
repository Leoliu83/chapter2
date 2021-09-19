package tree

import (
	"bytes"
	"fmt"
	"log"
	"runtime"
)

/*
	树的孩子表示法
	每个结点有多个指针域，其中每个指针指向一棵树的根节点，我们把这种方法称为多重链表表示法。
	TODO 待改进
*/
/*
	结点的数据类型
*/
type dataTypeC int

/*
	孩子结点
*/
type ChildNode struct {
	child int        // 是孩子结点在数组中的下标
	next  *ChildNode // 下一个孩子
	_     struct{}   // 强制使使用者用属性名进行初始化
}

/*
	结点
*/
type Cbox struct {
	IsInit     bool       // 是否初始化，默认为false
	idx        int        // *数组下标位置，用于打印与child比对，无实际意义
	data       dataTypeC  // 结点的数据
	firstChild *ChildNode // 孩子链表的头指针
	_          struct{}   // 强制使使用者用属性名进行初始化
}

/*
	孩子结点表示法的树
*/
type CTree struct {
	debug   bool
	flag    int    // 位图标志，用于快速定位可用的数组下标(每一个bit含义: 0空 1非空)
	maxsize int    // 最大结点数
	nodes   []Cbox // 结点数组
	r, n    int    // 根位置和节点数
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

/*
	初始化空树
*/
func CInitTree(maxsize int) *CTree {
	ctree := CTree{maxsize: maxsize, nodes: make([]Cbox, maxsize), r: 0, n: 0, flag: 0, debug: false}
	/*
		下面这段代码用于注册垃圾回收时所需要执行的函数
		类似于java中的Object中的finalize()方法
		主要含义是:
			对象可以关联一个SetFinalizer函数， 当gc检测到unreachable对象有关联的SetFinalizer函数时，会执行关联的SetFinalizer函数，
		同时取消关联。 这样当下一次gc的时候，对象重新处于unreachable状态并且没有SetFinalizer关联， 就会被回收。

		还有几个需要注意的点：
		1.即使程序正常结束或者发生错误， 但是在对象被 gc 选中并被回收之前，SetFinalizer 都不会执行， 所以不要在SetFinalizer中执行将内存中的内容flush到磁盘这种操作
		2.SetFinalizer 最大的问题是延长了对象生命周期。在第一次回收时执行 Finalizer 函数，且目标对象重新变成可达状态，直到第二次才真正 “销毁”。这对于有大量对象分配的高并发算法，可能会造成很大麻烦
		3.指针构成的 "循环引⽤" 加上 runtime.SetFinalizer 会导致内存泄露
	*/
	runtime.SetFinalizer(&ctree, func(t *CTree) {
		log.Println("Ctree struct Finalized!")
	})
	return &ctree
}

/*
	销毁数，在这里无法像c一样释放内存，因此只能将所有参数初始化
*/
func (t *CTree) Destory() {
	// 清空最大结点数量，为-1代表该树无法再次使用
	t.maxsize = -1
	t.Clear()
	log.Printf("%v,%p", &t, t)
}

/*
	保留树最大结点数量
*/
func (t *CTree) Clear() {
	t.n = 0
	t.r = 0
	t.nodes = make([]Cbox, 0)
}

func CCreateCBox(data dataTypeC) Cbox {
	return Cbox{IsInit: true, data: data, firstChild: nil}
}

/*
	树是否为空
*/
func CTreeEmpty(t *CTree) bool {
	// maxsize如果为0 ，说明树可能已经被销毁了
	if t.n == 0 && t.maxsize > 0 {
		return true
	} else {
		return false
	}
}

/*
	寻找深度
*/
func (t *CTree) TreeDepth() {
	for _, node := range t.nodes {
		log.Println(node)
	}
}

/*
	插入结点
	插入c为树T中p指结点的第i棵子树
	@param p 在哪个结点插入子节点
	@param c 需要插入的子节点
	@param i 在子节点的什么位置插入子节点，从0开始
*/
func (tree *CTree) InsertChild(p *Cbox, c Cbox, i int) {
	if tree == nil || tree.maxsize == 0 {
		if tree.debug {
			log.Println("Tree is nil or has been destoryed!")
		}
		return
	}
	if tree.debug {
		log.Println("InsertChild starting......")
	}
	idx, ok := tree.FindFreeIdx()
	if !ok {
		return
	}
	if tree.debug {
		log.Println("Finding position to insert the child....")
	}
	// 如果firstchild为空，则直接赋值给firstchild
	if p.firstChild == nil {
		p.firstChild = &ChildNode{child: idx}
	} else {
		// 如果不为空，则将firstnode赋值为当前node，将当前node的next赋值为原firstnode的next
		// 在不讲究顺序的情况下可以这么做，比较方便，相当于新node作为firstnode，这里不这么做，这里通过传入的i来实现对指定位置的插入
		node := p.firstChild
		/*
			分情况讨论：
			  1. 当j==-1 ，传入为0，则需要将原node变为新node的next，而firstnode变为新node
			则
		*/
		var j int
		/*
			这里j选择-1，因为如果当node==nil的时候，j就指向了当前位置的前一个结点，处理方便
			例如 i=2
			则当判断到i==2时，j==1，在处理的时候，只需要将j所指的当前结点的next置为新结点，然后新结点的next指向当前节点的next
		*/
		for j = -1; node != nil; j++ {
			if j+1 == i {
				if tree.debug {
					log.Println("Index is set to zero.")
				}
				break
			}
			if node.next != nil {
				node = node.next
			}
		}
		if tree.debug {
			log.Printf("i[%d] : j[%d] : node[%+v]", i, j, node)
		}
		/*
			该判断用于防止输入的i大于子节点的数量，例如，如果子节点只有2个，而i输入的是3，则是无效输入
			i输入2是有效输入，因此i从0开始
		*/
		if j < i {
			log.Printf("Illegal input 'i': [%d], child count: [%d], i start with not 1 but 0 .", i, j)
			return
		}
		if i == 0 {
			if tree.debug {
				log.Println("Change insert new child to the first position.")
			}
			p.firstChild = &ChildNode{child: idx, next: node}
		} else {
			node.next = &ChildNode{child: idx, next: node.next}
		}

		if tree.debug {
			log.Printf("%p,%+v", node, node)
			log.Printf("%p,%+v", p.firstChild, p.firstChild)
		}
	}
	// 增加结点
	c.idx = idx
	tree.nodes[idx] = c
	// 修改位图标志位，将对应位置更新为1
	tree.flag |= (1 << idx)
	tree.n++
	// log.Printf("%08b, %+v", tree.flag, tree)
	// log.Printf("%08b, %+v, %+v", tree.flag, tree.nodes[0].firstChild, tree.nodes[0].firstChild.next)
	log.Println("InsertChild finished......")
}

/*
	删除结点p的第i棵子树
*/
func (tree *CTree) DeleteChild(p *Cbox, i int) {

}

/*
	创建根结点
	@param c 根节点
*/
func (tree *CTree) InsertRoot(c Cbox) {
	tree.nodes[0] = c
	tree.flag = 1
	tree.r = 0
	tree.n = 1
}

/*
	寻找数组空闲元素的下标
	@return int  数组下标位置，-1表示已满
	@return bool 是否成功，如果已满，没有可用的下标，则返回false
*/
func (tree *CTree) FindFreeIdx() (int, bool) {
	if tree.debug {
		log.Println("Finding index....")
	}
	flag := tree.flag
	idx := 0
	// 寻找位图中0的位置
	for (flag & 1) == 1 {
		flag >>= 1
		idx++
	}
	// 如果索引值大于树的最大大小，返回false
	if idx >= tree.maxsize {
		log.Printf("Illegal index value: %d(idx) %d(now) %d(max)", idx, tree.n, tree.maxsize)
		return -1, false
	}
	if tree.debug {
		log.Printf("Found index: [%d] ;flag: [%08b]", idx, tree.flag)
	}
	// 检查tree的nodes数组是否需要扩容（废弃），现在由IsInit判断struct的初始化
	// if idx > tree.n {
	// 	tree.nodes = append(tree.nodes, Cbox{data: -1})
	// }
	return idx, true
}

/*
	清除数组下标位置，用于在删除结点时，将数组下标的bit位置为0
	@param 需要清除的下标位置
*/
func (tree *CTree) CleanFlag(idx int) {
	if tree.debug {
		log.Printf("%b,%+v", tree.flag, tree)
	}
	// 左移 idx 位，位图所表示的数组下标从右侧开始,&^是golang中的清位符
	// 位图：11111
	// 下标：43210
	tree.flag &^= 1 << (idx)
	if tree.debug {
		log.Printf("%b,%+v", tree.flag, tree)
	}
}

/*
	打印树的结构
*/
func (tree *CTree) Print() {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("\n[ tree bit: %08b\n", tree.flag))
	for i := range tree.nodes {
		node := tree.nodes[i]
		if node.IsInit {
			buffer.WriteString(fmt.Sprintf("\t[%p] %+v\n", &node, node))
			child := node.firstChild
			for child != nil {
				buffer.WriteString(fmt.Sprintf("\t\t|--[%p] %+v\n", child, *child))
				child = child.next
			}
		}
	}
	buffer.WriteString(" ] \n")
	log.Println(buffer.String())
}
