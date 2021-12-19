package algsort

import (
	"fmt"
)

// BubbleSort 函数是'冒泡排序'的实现
/*
	冒泡排序的步骤数量计算如下：
	设数组长度为：8
	对数组中的元素的访问次数为：7+6+...+1
	设数组为'反序'，反序的意思是：如果排序是要求从小到大，那么当前数组的元素是从大到小
	则，需要交换数组元素的次数为：7+6+...+1
	所以总步数为：2*(7+6+...+1)
	因此，冒泡排序的事件复杂度为 O(N^2)，也就是说用O(N^2)的程序处理N个元素需要N^2步
*/
func BubbleSort() ([]int16, int) {
	// needSorted := []int16{9, 5, 1, 8, 12, 3, 23, 8}
	// 极限
	needSorted := []int16{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	// 计步器
	steps := 0
	// 初始循环次数
	loopCnt := len(needSorted) - 1
	/*
		最外层循环表示: 每次内层循环遍历完成之后，数组中需要遍历的元素中，最后一个元素一定是最大的，
		下一次对数组遍历可以排除最后一个元素，例如，数组初始为[9, 5, 1, 8]
		第一次内层循环遍历数组排序完成之后，'9'一定是最后一个元素,下一次内层循环只需要遍历前3个元素
		第二次内层循环遍历数组排序完成之后，'8'一定是前3个元素的最后一个元素，下一次内层循环遍历只需要遍历前2个元素
		...
		以此类推
	*/
	for j := loopCnt; j > 0; j-- {
		// fmt.Println(steps)
		for i := 0; i < j; i++ {
			steps++
			// fmt.Println("<-- ", needSorted[i], needSorted[i+1])
			if needSorted[i] > needSorted[i+1] {
				steps++
				needSorted[i], needSorted[i+1] = needSorted[i+1], needSorted[i]
				// fmt.Println("--> ", needSorted[i], needSorted[i+1])
			}
		}
	}

	fmt.Printf("[BubbleSort] Steps: %d, array: %+v \n", steps, needSorted)
	return needSorted, steps
}
