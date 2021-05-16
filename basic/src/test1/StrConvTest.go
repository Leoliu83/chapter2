package test1

import (
	"fmt"
	"strconv"
)

// StrConvTest 方法用于测试 strconv 包
func StrConvTest() {
	a, _ := strconv.ParseInt("100", 2, 32)
	b, _ := strconv.ParseInt("0144", 8, 32)
	c, _ := strconv.ParseInt("64", 16, 32)
	fmt.Println(a, b, c)

	d := strconv.FormatInt(100, 2)
	e := strconv.FormatInt(100, 8)
	f := strconv.FormatInt(100, 16)
	fmt.Println(d, e, "0x"+f)
}
