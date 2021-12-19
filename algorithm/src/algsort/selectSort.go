package algsort

import (
	"fmt"
)

// SelectSort 函数是'选择排序'的实现
/*
	该排序的事件复杂度为 O(N^2)
	虽然它处理N个元素需要N^2/2步，但是对于事件复杂度来说大O记法是忽略常量的，
	因此事件复杂度也为O(N^2)
	大O记法主要是用来区分算法在不同大O下的比较
*/
func SelectSort() {
	// 计步器
	steps := 0

	// needSorted := []int16{9, 5, 1, 8, 12, 3, 23, 8}
	// 极限
	needSorted := []int16{2, 3, 4, 5, 6, 7, 8, 9, 10, 1}

	smallestIdx := 0
	loopCnt := len(needSorted) - 1
	for i := 0; i < loopCnt; i++ {
		// fmt.Println(steps)
		// 将 smallestIdx 设置为下一次循环的第一个元素的下标
		smallestIdx = i
		// i+1 表示 第一个元素不需要和smallestIdx比对，因为每次外循环都会讲smallestIdx设置为第一个元素的下标
		for j := i + 1; j <= loopCnt; j++ {
			steps++
			if needSorted[j] < needSorted[smallestIdx] {
				smallestIdx = j
			}
		}
		if i != smallestIdx {
			steps++
			needSorted[i], needSorted[smallestIdx] = needSorted[smallestIdx], needSorted[i]
		}
	}
	fmt.Printf("[SelectSort] Steps: %d, array: %+v \n", steps, needSorted)

}
