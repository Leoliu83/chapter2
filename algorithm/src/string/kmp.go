package string

import (
	"fmt"
	"log"
	"unsafe"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

/*
	字符串KMP模式匹配算法
*/
func KmpWithIndex(src string, sub string) int {
	// srcLen := len(src)
	// subLen := len(sub)

	// // 主串下标
	// i := 0
	// // 子串下标
	// j := 0
	// // 子串回溯位
	// k := 0
	// if sub[j] == sub[k-1] {
	// 	k++
	// }

	// if src[i] == sub[j] {
	// 	i++
	// 	j++
	// }

	return -1
}

/*
	生成next数组
*/
func GetNextArray(sub string) {
	l := len(sub)
	printStr(sub)
	// 主要处理逻辑
	next := make([]int, len(sub))
	k := 0
	j := 2

	next[0] = 0
	next[1] = 1

	for j < l {
		/*
			逐项比对，当j=2时候，从 sub[1]与sub[0] 比，
			如果比对成功，则继续 sub[2] 与 sub[1]比对
			因为每一位的next值是他前一位的比对结果
		*/
		if sub[k] == sub[j-1] {
			next[j] = next[j-1] + 1
		} else {
			// 如果不相等，则子串的下一位重新与第一位比较
			k = 0
			if sub[k] == sub[j-1] {
				next[j] = 2
			} else {
				next[j] = 1
			}
		}
		k++
		j++
	}
	log.Printf("%+v", next)
}

/*
	原书中代码
*/
func GetNext(sub string) {
	runeArray := []rune(sub)
	log.Printf("%+v", runeArray)
	l := len(runeArray)
	// 打印地址
	// printRuneArray(runeArray)
	// 新数组 下标0位放长度，后面开始放数据
	subArray := make([]int32, l+1)
	// 将原始string的字符数组长度保存在数组的0号位置
	subArray[0] = int32(l)
	log.Printf("%d", len(runeArray))
	intArrayP := (*[]int32)(unsafe.Pointer(&runeArray))
	copy(subArray[1:], *intArrayP)
	log.Printf("%+v", subArray)

	// 定义i和j
	var i, j int32 = 1, 0
	// 创建next数组
	next := make([]int32, l+1)
	// 书中的next从1开始，书中用subArray[0]字符串表示长度
	next[1] = 0
	for i < subArray[0] {
		/*
			subArray[j]表示`前缀`的单个字符
			subArray[i]表示`后缀`的单个字符
		*/
		// 当j=0，也就是前缀为0
		// 当 前缀单个字符 = 后缀单个字符
		if j == 0 || subArray[i] == subArray[j] {
			i++
			j++
			next[i] = j
		}
	}
}

/*
	将字符串打印成数组，并且标出下标，从1开始
	例如：abcdex
	[1 2 3 4 5 6]
    [a b c d e x]
*/
func printStr(sub string) {
	l := len(sub)
	sbegin := "["
	send := "]"
	ibegin := "["
	iend := "]"
	for i := 1; i <= l; i++ {
		sbegin += sub[i-1 : i]
		ibegin += fmt.Sprintf("%d", i)
		if i < l {
			sbegin += " "
			ibegin += " "
		}
	}
	sbegin += send
	ibegin += iend
	log.Printf("%s", ibegin)
	log.Printf("%s", sbegin)
}

func printRuneArray(arr []rune) {
	l := len(arr)
	for i := 0; i < l; i++ {
		log.Printf("%p:%d(%T)", &arr[i], arr[i], arr[i])
	}
}
