package test1

import (
	"log"
	"reflect"
	// "strings"
)

/*
	反射库提供了内置函数make和new的对应操作
	使用MakeFunc可用于实现通用模板，适应不同的数据类型
	实际上，用另一种方式实现了泛型
*/
/*
	add103 是一个通用方法模板：
	args []reflect.Value 参数值列表（slice）
	results []reflect.Value 返回值列表（slice）
	MakeFunc中第一个参数 reflect.Type（Type可以看成是“函数签名”）的参数和返回值必须和 模板一致
*/
func add103(args []reflect.Value) (results []reflect.Value) {
	if len(args) == 0 {
		return nil
	}
	var ret reflect.Value
	switch args[0].Kind() {
	case reflect.Int:
		n := 0
		for _, a := range args {
			n += int(a.Int())
		}
		ret = reflect.ValueOf(n)
	case reflect.String:
		n := make([]string, 0, len(args))
		for _, a := range args {
			n = append(n, a.String())
		}
		// ret = reflect.ValueOf(strings.Join(n, ","))
		// 这里的n是个slice，因此MakeFunc的第一个参数的签名返回值必须是 string 的slice
		ret = reflect.ValueOf(n)
	}
	results = append(results, ret)
	log.Println("results: ", results)
	return
}

func makeAdd103(fptr interface{}) {
	fn := reflect.ValueOf(fptr).Elem()
	// log.Println(fn, fn.Type())
	v := reflect.MakeFunc(fn.Type(), add103)
	fn.Set(v)
}

func ReflectCreateTest1() {
	var intAdd func(x, y int) int
	var strAdd func(x, y string) []string
	makeAdd103(&intAdd)
	makeAdd103(&strAdd)

	// intAdd(1, 2)
	log.Println(strAdd("a", "b"))

}
