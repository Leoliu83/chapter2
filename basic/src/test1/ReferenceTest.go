package test1

import (
	"log"
)

/*
	引用类型主要指slice map channel这三种预定义类型
	初始化方法：
		new ：new内置函数按照指定类型长度分配0值内存，返回指针
		make：make函数可以转换为目标类型专用的创建函数,以确保所有的内存分配和初始化参数
	因此，引用类型必须使用make
*/
func NewMakeTest() {
	// new产生的slice不能够使用append，因为它认为数组不是一个slice
	// slice1 := new([]int)
	// slice1 = append(slice1, 1)
	// 使用make创建slice
	slice2 := make([]int, 1, 10)
	slice2 = append(slice2, 1)
	log.Printf("%+v,%d \n", slice2, len(slice2))
	// 用new初始化的map仅仅分配了字典类型本身所需的内存，并没有分配键值存储内存，也没有初始化散列桶等内部属性
	map1 := new(map[string]int) // 返回指针
	map2 := *map1
	map2["a"] = 1 // 报错 nil map
	log.Printf("%+v \n", map2)

}
