package test1

import (
	"log"
	"reflect"
)

/*
	反射可以让我们在运行期探知对象的类型信息和内存结构，一定程度弥补了静态语言在动态行为上的不足，也是实现元编程的重要手段
	和C一样，GO对象头部并没有类型指针，通过其自身无法获取任何类型相关信息的。反射操作所需的全部信息都源自接口变量。接口变量除了存储自身类型外，还会保存实际的类型数据。
	func TypeOf(i interface{}) Type
	func ValueOf(i interface{}) Value
	在面对类型时候，要区分Type和Kind，Type表示真实类型（静态类型），Kind表示基础结构（底层类型）类别
*/
func ReflectTest() {
	type T int
	// a 的type是 T ，a的kind是 int
	var a T = 100
	t := reflect.TypeOf(a)
	log.Printf("Type: %s, Kind: %s", t.Name(), t.Kind())
}

/*
	在类型的判断上需要使用正确的方式
*/
func ReflectTest1() {
	type T int
	type Y int
	var a, b T = 100, 200
	var c Y = 300
	ta, tb, tc := reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c)
	log.Println(ta == tb, tb == tc)                             // true false
	log.Println(ta.Kind() == tb.Kind(), tb.Kind() == tc.Kind()) // true true
	// 基类型和指针类型不是同一类型
	log.Println(ta, reflect.TypeOf(&a), ta == reflect.TypeOf(&a))
}

/*
	reflect可以构造基础复合类型
*/
func ReflectTest2() {
	a := reflect.ArrayOf(10, reflect.TypeOf(byte(0)))
	m := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0))
	log.Println(a, m)
}

/*
	匿名字段结构显示
*/
func ReflectTest3() {
	type user struct {
		id   int
		name string
	}
	type manager struct {
		user
		title string
	}

	var m manager
	t := reflect.TypeOf(m)
	numf := t.NumField()
	for i := 0; i < numf; i++ {
		f := t.Field(i)
		if f.Anonymous { // 判断是否匿名结构
			for j := 0; j < f.Type.NumField(); j++ {
				af := f.Type.Field(j)
				log.Println(" ", af, af.Type.Name(), af.Type.Kind())
			}
		}
	}

	fn, ok := t.FieldByName("name")
	log.Println("通过名称查找匿名字段：", fn, ok)
	id := t.FieldByIndex([]int{0, 0}) // 第一个 0 表示第一个字段，也就是匿名字段user，第二个0是匿名字段user中的第1个字段 id
	log.Println("通过多级索引查找匿名字段：", id)
}

/*
	输出方法集的时候，也区分基础类型和指针类型
*/
type A101 int
type B101 struct {
	A101
}

func (A101) a101v()  {}
func (*A101) a101p() {}
func (B101) b101v()  {}
func (*B101) b101p() {}

func ReflectTest4() {
	var b B101
	t := reflect.TypeOf(&b)
	s := []reflect.Type{t, t.Elem()}
	log.Println(s)
}
