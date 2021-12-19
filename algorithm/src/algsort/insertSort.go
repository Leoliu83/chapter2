package algsort

import (
	"fmt"
)

// InsertSort 函数是'插入排序'的实现
/*
 */
func InsertSort() {
	var pos int
	var tempValue int16
	steps := 0
	// needSorted := []int16{9, 5, 1, 8, 12, 3, 23, 8}
	// 极限
	needSorted := []int16{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	loopCnt := len(needSorted)
	for i := 1; i < loopCnt; i++ {
		pos = 0
		// 依次将第2-n个值分别取出放到临时变量tempValue里
		tempValue = needSorted[i]
		// 将0~(n-1)个值与n依次进行比较
		for j := i - 1; j >= 0; j-- {
			steps++
			// 如果0~(n-1)个值中有值大于tempValue，则将该值向右移位
			if tempValue < needSorted[j] {
				// 向右移位
				needSorted[j+1] = needSorted[j]
				steps++
			} else {
				pos = j + 1
				break
			}
		}
		needSorted[pos] = tempValue
	}
	fmt.Printf("[InsertSort] Steps: %d, array: %+v \n", steps, needSorted)
}
