package string

import (
	"fmt"
	"log"
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

func GetNext(sub string) {
	l := len(sub)
	subArray := make([]int, l+1)
	subArray[0] = l
	runeArray := []rune(sub)
	log.Printf("%d", len(runeArray))
	// copy(subArray[1:], *pt)
	// log.Printf("%+v", subArray)
	// next := make([]int, l)
	// var i, j int = 1, 0
	// next[1] = 0
	// for i < l {
	// 	// sub[j]表示前缀的单个字符
	// 	// sub[i]表示后缀单个字符
	// 	if j == 0 || sub[i] == sub[j] {
	// 		i++
	// 		j++
	// 		next[i] = j
	// 	}
	// }
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
