package innerpkg

import (
	"golangStudy/test1/innerpkg/internal"
	"log"
)

type data93 struct {
	x int // 非导出字段(小写)
	Y int // 可导出字段(大写)
}

func Newdata() *data93 {
	return new(data93)
}

// 无法访问
func innerfunc() {
	log.Println("inner func")
}

/*
	internal 包的上级目录下的所有包都可以调用internal包
*/
func InternalTest() {
	ib := internal.GetInternalObj()
	println(ib)
}
