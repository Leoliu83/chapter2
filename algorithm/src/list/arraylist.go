package list

import "log"

/*
	线性表的数组实现 v0.0.1
*/

/*
	数组中的元素类型为interface
	注意：参数传递给interface{}会保留其类型信息，使得interface变量不为真正的nil
*/
type Element interface{}

type ArrayList struct {
	data [10]Element // 数据区
	len  int
	cap  int
}

func (arrayList *ArrayList) InitList() bool {
	arrayList.len = 0
	arrayList.cap = 10
	return true
}

func (arrayList *ArrayList) ListEmpty() (isEmpty bool) {
	isEmpty = true
	if arrayList.len != 0 {
		isEmpty = false
	}
	return
}

func (arrayList *ArrayList) ClearList() bool {
	arrayList.len = 0
	var newData [10]Element
	arrayList.data = newData
	return true
}

/*
	获取数据,返回的Element作为副本，因为取值不应该允许对原元素产生影响
*/
func (arrayList *ArrayList) GetElem(i int) (Element, bool) {
	if !arrayList.checkPosInRange(i) {
		return nil, false
	}
	e := arrayList.data[i]
	return e, true
}

func (arrayList *ArrayList) ListInsert(e Element, i int) bool {
	// 不允许跳跃插入数据，例如len=3 不允许插入数据位置到4
	if i > arrayList.len || i < 0 || i >= arrayList.cap {
		log.Printf("[ERROR]: Illegal parameter i: %d", i)
		return false
	}
	// 如果i是最后一个元素，则直接替换元素，并将长度设置为最大值
	if i == arrayList.cap-1 {
		arrayList.data[i] = e
		arrayList.len = arrayList.cap
		return true
	}
	// 如果i的位置不是是最后一个元素，则所有元素后移，不是则不需要后移操作
	if i < arrayList.cap-1 {
		// 从倒数第二个元素开始遍历,直到i，每个元素向后移位
		for j := arrayList.len - 1; j >= i; j-- {
			arrayList.data[j+1] = arrayList.data[j]
		}
	}
	arrayList.data[i] = e
	// i小于等于元素长度，则长度+1
	if i <= arrayList.len {
		arrayList.len++
	} else { // 如果i超过元素长度，则元素长度设置为i+1 例如：已存在2个元素，插入元素位置到3，则元素长度为4，而不是3
		arrayList.len = i + 1
	}
	return true
}

/*
	往列表末位添加元素
*/
func (arrayList *ArrayList) ListAppend(e Element) bool {
	if arrayList.len == 10 {
		return false
	}
	arrayList.len++
	arrayList.data[arrayList.len-1] = e
	return true
}

/*
	删除一个元素
*/
func (arrayList *ArrayList) ListDelete(i int) (Element, bool) {
	if i >= arrayList.len || i < 0 {
		return nil, false
	}
	e := arrayList.data[i]
	for j := i; j < arrayList.len; j++ {
		if j == arrayList.len-1 {
			arrayList.data[j] = nil
		} else {
			arrayList.data[j] = arrayList.data[j+1]
		}
	}
	arrayList.len--
	return e, true
}

/*
	获取长度
*/
func (arrayList *ArrayList) ListLength() int {
	return arrayList.len
}

func (arrayList *ArrayList) LocateElem(e Element) int {
	for i, d := range arrayList.data {
		if d == e {
			return i
		}
	}
	return -1
}

func (arrayList *ArrayList) Union(arrayList1 ArrayList) *ArrayList {
	var newArrayList ArrayList

	return &newArrayList
}

/*
	判断获取元素位置参数是否有效
*/
func (arrayList *ArrayList) checkPosInRange(i int) bool {
	// 如果线性表长度为0或者参数i超过线性表长度或者i小于0，则返回false
	if arrayList.len == 0 || i >= arrayList.len || i < 0 {
		return false
	}
	return true
}
