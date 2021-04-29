package test1

import (
	"golangStudy/test1/innerpkg"
	"log"
	"unsafe"
)

/*
	通过unsafe 指针转换访问自有字段，
*/
func PermissionTest() {
	d := innerpkg.Newdata()
	d.Y = 100
	// d.x = 200 // 不可访问

	// 访问私有字段
	dp := unsafe.Pointer(d)
	dps := (*struct {
		x int
		y int
	})(dp)
	dps.x = 300
	log.Printf("%+v", d)
}
