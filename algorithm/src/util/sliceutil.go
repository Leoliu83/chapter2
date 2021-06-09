package util

import (
	"log"
	"reflect"
	"unsafe"
)

/*
	interface参数会复制数据，还会保留类型信息，这里复制的是头
	如果是[]int 就可以通过 s.([]int)中 slice的副本，也就是SliceHeader的副本
	reflect.Kind 参数不适用传递自定义类型
	**这里没有考虑传入指针的情况

*/
func PrintSliceHeader(s interface{}, t reflect.Kind) {
	var isSlice bool
	v := reflect.ValueOf(s)
	// 通过Kind判断是否是slice
	// 如果传入的是指针，这里可以使用 v.Elem().Kind()
	switch v.Kind() {
	case reflect.Slice:
		isSlice = true
	default:
		isSlice = false
	}
	if !isSlice {
		log.Println("Not a slice")
		log.Printf("%+v", v)
		return
	}

	var slice unsafe.Pointer
	// switch中其他类型后续可以添加
	switch t {
	case reflect.Int:
		tmp, ok := s.([]int)
		if !ok {
			log.Println("[ERROR]: Transform to []int error.")
			return
		}
		log.Printf("%p,%s", tmp, reflect.TypeOf(tmp))
		slice = unsafe.Pointer(&tmp)
	case reflect.String:
		tmp, ok := s.([]string)
		if !ok {
			log.Println("[ERROR]: Transform to []int error.")
			return
		}
		log.Printf("%p,%s", tmp, reflect.TypeOf(tmp))
		slice = unsafe.Pointer(&tmp)
	default:
		log.Printf("[ERROR]: Unsupport type: %d, %#v", t, v)
		return
	}
	// 强转成*reflect.SliceHeader指针类型
	header := (*reflect.SliceHeader)(slice)
	log.Printf("%+v", header)
	log.Printf("%p", unsafe.Pointer(header.Data))
	// log.Println((*slice)[0])
	log.Println("------------------------------------------------------")
	log.Printf(">> [Address]s: %p", s)
	log.Printf(">> [Address]First_Element: %p", s)
	log.Printf(">> [Address]header.Data: %p", &header.Data)
	log.Printf(">> [Data]header.Data: %#v", header.Data)
	log.Printf(">> [Data]header.Cap: %d", header.Cap)
	log.Printf(">> [Data]header.Len: %d", header.Len)
	log.Printf(">> [Data]Slice: %+v", s)
	log.Println("------------------------------------------------------")
}
