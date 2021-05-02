package test1

import (
	"log"
	"reflect"
)

/*
	和Type获取的信息不同，Value专注于对象实例数据的读写
*/
type obj102 struct {
	x    int    // unexported field
	Y    int    // exported field
	Name string // exported field
	code int    // unexported field
	_    struct{}
}

/*
	复习：
		当把一个对象赋值给一个接口类型变量时，会复制该对象，并且该对象不可寻址，不可被修改
		如果需要修改目标对象，必须使用指针
	ifaceV.(obj102).y = 200 // 编译错误
	同样需要 .Elem() 获取目标对象，因为被接口本身存储的指针本身是不能寻址和进行设置操作的
*/
func ReflectValueTest1() {
	obj1 := obj102{x: 1, Y: 2}
	var ifaceV interface{} = obj1
	var ifaceP interface{} = &obj1
	ifaceType1, ifaceType2 := reflect.ValueOf(ifaceV.(obj102)), reflect.ValueOf(ifaceP.(*obj102)).Elem()
	log.Println("'ifaceType1' addressable? ", ifaceType1.CanAddr())
	log.Println("'ifaceType1' can be set? ", ifaceType1.CanSet())
	log.Println("'ifaceType2' addressable? ", ifaceType2.CanAddr())
	log.Println("'ifaceType2' can be set? ", ifaceType2.CanSet())
	log.Printf("%+v", ifaceV.(obj102))
	ifaceP.(*obj102).x = 100
	ifaceP.(*obj102).Y = 200
	log.Println(obj1)
	log.Println("==================================================")
	// 指针
	obj2 := new(obj102)
	v := reflect.ValueOf(obj2).Elem()
	name := v.FieldByName("Name")
	code := v.FieldByName("code")
	// 只能设置导出(exported) 字段，不能设置非导出 (unexported) 字段
	log.Println("Field 'Name' can be set? ", name.CanSet()) // true
	log.Println("Field 'code' can be set? ", code.CanSet()) // false
	name.SetString("Leoliu")
	code.SetInt(1001) // panic: reflect: reflect.Value.SetInt using value obtained using unexported field
	log.Printf("%+v", obj2)
}
