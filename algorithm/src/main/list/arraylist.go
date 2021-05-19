package list

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
	if !arrayList.checkPos(i) {
		return nil, false
	}
	e := arrayList.data[i]
	return e, true
}

func (arrayList *ArrayList) ListInsert(e Element, i int) bool {
	if i >= len(arrayList.data) {
		return false
	}

	arrayList.data[i] = e
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

func (arrayList *ArrayList) checkPos(i int) bool {
	if i >= arrayList.len {
		return false
	}
	return true
}
