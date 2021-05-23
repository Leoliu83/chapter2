package list

import (
	"testing"
)

func TestInitList(t *testing.T) {
	var list ArrayList
	ok := list.InitList()
	t.Logf("%t,%+v", ok, list.data)
}

func TestListAppend(t *testing.T) {
	var list ArrayList
	ok := list.InitList()
	if !ok {
		return
	}
	for i := 0; i < 11; i++ {
		list.ListAppend(i)
		t.Logf("list: %+v", list)
	}
}

func TestListDelete(t *testing.T) {
	var list ArrayList
	ok := list.InitList()
	if !ok {
		return
	}
	for i := 0; i < 11; i++ {
		list.ListAppend(i)
	}
	t.Logf("list: %+v", list)
	var e Element
	e, ok = list.ListDelete(5)
	t.Logf("Delete element: %#v,Delete success? %t,list: %+v", e, ok, list)
	e, ok = list.ListDelete(7)
	t.Logf("Delete element: %#v,Delete success? %t,list: %+v", e, ok, list)
	e, ok = list.ListDelete(8)
	t.Logf("Delete element: %#v,Delete success? %t,list: %+v", e, ok, list)
}

func TestListEmpty(t *testing.T) {
	var list ArrayList
	isEmpty := list.ListEmpty()
	t.Logf("isEmpty: %t", isEmpty)
}

func TestGetElem(t *testing.T) {
	var list ArrayList
	ok := list.InitList()
	t.Log(len(list.data))
	e, ok := list.GetElem(9)
	if ok {
		t.Logf("ok? %t,Element: %+v", ok, e)
	}
}

func TestListInsert(t *testing.T) {
	var list ArrayList
	list.InitList()
	for i := 0; i < 5; i++ {
		list.ListAppend(i)
	}
	t.Logf("list: %+v", list)
	var ok bool
	ok = list.ListInsert(-2, 2)
	t.Logf("ok? %t,Element: %+v", ok, list)
	ok = list.ListInsert(-4, 4)
	t.Logf("ok? %t,Element: %+v", ok, list)
	ok = list.ListInsert(-7, 7)
	t.Logf("ok? %t,Element: %+v", ok, list)
	ok = list.ListInsert(-8, 9)
	if ok {
		t.Logf("ok? %t,Element: %+v", ok, list)
	}
	ok = list.ListInsert(-9, 9)
	if ok {
		t.Logf("ok? %t,Element: %+v", ok, list)
	}
	ok = list.ListInsert(-10, 10)
	if ok {
		t.Logf("ok? %t,Element: %+v", ok, list)
	}
}

func TestLocalElem(t *testing.T) {
	var list ArrayList
	list.InitList()
	_ = list.ListAppend("1")
	_ = list.ListAppend("2")
	_ = list.ListAppend("3")
	t.Logf("list: %+v", list)
	i := list.LocateElem("2")
	t.Logf("i = %d", i)
}

func TestUnion(t *testing.T) {
	var list ArrayList
	var list1 ArrayList
	list.InitList()
	list1.InitList()
	_ = list.ListAppend("1")
	_ = list.ListAppend("2")
	_ = list.ListAppend("3")

	_ = list1.ListAppend("1")
	_ = list1.ListAppend("a")
	_ = list1.ListAppend("b")
	newList := list.Union(list1)
	t.Logf("newList: %+v", *newList)
}
