package string

import "log"

/*
	字符串的朴素模式匹配算法
*/
func BruteForce(src string, sub string) int {
	srcLen := len(src)
	subLen := len(sub)
	matched := false
	loop := srcLen - subLen + 1
	if subLen > srcLen {
		log.Fatalf("Error len src[%d],sub [%d]", srcLen, subLen)
	}

	for i := 0; i < loop; i++ {
		if matched {
			break
		}
		if src[i] == sub[0] {
			// println(i, src[i], sub[0])
			j := 1 // j 必须放在for循环外层，可以用于后续判断
			for ; j < subLen; j++ {
				// log.Println(src[int(i)+j], sub[j])
				if src[i+j] != sub[j] {
					break
				}
			}
			// log.Print(j, subLen)
			// 如果 j < subLen 说明循环没有完成，如果上一步循环完成，j 应该等于 subLen，否则说明中途break了
			if j < subLen {
				continue
			}
			matched = true
			// log.Println(matched, i)
			return i
		}
	}
	return -1
}

/*
	使用游标方式实现
*/
func BruteForceWithIndex(src string, sub string) int {
	// 主字符串的游标
	srcIdx := 0
	// 子字符串游标
	subIdx := 0
	srcLen := len(src)
	subLen := len(sub)
	/*
	  如果主字符的下标大于了主字符串长度，则退出循环
	  由于 srcIdx 是从0开始的，因此 srcIdx 不会大于 srcLen
	*/
	for srcIdx < srcLen {
		// log.Println("Src: ", srcIdx, ", Sub:", subIdx, " => ", src[srcIdx], ":", sub[subIdx])
		// 正常来说 游标是从0开始，因此子游标最大应该满足 subIdx = subLen-1，因此subIdx = subLen，就说明匹配完全了。
		if subIdx == subLen {
			log.Printf("Found substring on index : %d", srcIdx-subIdx)
			return srcIdx - subIdx
		}
		// 如果匹配，则游标双方加一
		if src[srcIdx] == sub[subIdx] {
			srcIdx++
			subIdx++
		} else {
			/*
				如果不匹配，则
				1. subIdx置0
				2. srcIdx退回与子串匹配的第一位的后一位
				例如: 如果src的第一位是5,sub第一位是0，当比对到sub的第4位（即subIdx=3,此时srcIdx=8）发现不匹配
				      则退回 5 即 8 - 4 + 1
				5 6 7 8
				0 1 2 3
			*/
			srcIdx = srcIdx - subIdx + 1
			subIdx = 0
		}
	}
	return -1
}
