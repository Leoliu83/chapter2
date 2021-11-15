package queue

import "log"

/*
	循环队列的go实现，但不支持并发处理
*/
type LoopQueue struct {
	Data    []int
	Front   int  // 队列头
	Rear    int  // 队列尾
	MaxLen  int  // 最大长度
	IsFull  bool // 队列是否已满
	IsEmpty bool // 是否为空队列
}

func (lq *LoopQueue) InitQueue(n int) {
	lq.Front = 0
	lq.Rear = 0
	lq.Data = make([]int, n)
	lq.MaxLen = n
	lq.IsFull = false
	lq.IsEmpty = true
}

/*
	入队，这里引入 IsFull 和 IsEmpty 变量，
	或者也可以使用变量count来保存元素个数，如果count=MaxLen 则说明满了，如果count=0说明空了
	而不像文中那样直接用 (lq.Rear+1)%lq.MaxLen == lq.Front 判断
	因为如果直接判断会导致最后一个数组位置无法填充数字，造成浪费
	即，当 lq.Rear = 3 时，执行完EnterQueue，此时lq.Rear变为4，
	在一开始判断的时候，lq.Rear+1变成了5，(lq.Rear+1)%lq.MaxLen == lq.Front 为true，因此下标为4的位置无法填充
	原书代码如下：
	if (lq.Rear + 1) % lq.MaxLen {
		log.Println("Queue is full!")
		return
	}
	lq.Data[lq.Rear] = e
	lq.Rear = (lq.Rear + 1) % lq.MaxLen

*/
func (lq *LoopQueue) EnterQueue(e int) {
	if lq.IsFull {
		log.Println("Queue is full!")
		return
	}
	lq.Data[lq.Rear] = e
	lq.Rear = (lq.Rear + 1) % lq.MaxLen
	lq.IsEmpty = false
	if (lq.Rear)%lq.MaxLen == lq.Front {
		lq.IsFull = true
	}
}

/*
	出队
*/
func (lq *LoopQueue) DeleteQueue() {
	if lq.IsEmpty {
		log.Println("Queue is Empty!")
		return
	}
	lq.Data[lq.Front] = 0
	lq.Front = (lq.Front + 1) % lq.MaxLen
	if lq.Rear == lq.Front {
		lq.IsEmpty = true
	}
}
