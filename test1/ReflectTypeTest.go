package test1

import (
	"fmt"
	"log"
	"reflect"
)

/*
	反射可以让我们在运行期探知对象的类型信息和内存结构，一定程度弥补了静态语言在动态行为上的不足，也是实现元编程的重要手段
	和C一样，GO对象头部并没有类型指针，通过其自身无法获取任何类型相关信息的。反射操作所需的全部信息都源自接口变量。接口变量除了存储自身类型外，还会保存实际的类型数据。
	func TypeOf(i interface{}) Type
	func ValueOf(i interface{}) Value
	在面对类型时候，要区分Type和Kind，Type表示真实类型（静态类型），Kind表示基础结构（底层类型）类别
	TypeOf函数可以从一个任何*非接口类型*的值创建一个reflect.Type值
*/
func ReflectTypeTest() {
	type T int
	// a 的type是 T ，a的kind是 int
	var a T = 100
	t := reflect.TypeOf(a)
	log.Printf("Type: %s, Kind: %s", t.Name(), t.Kind())
}

/*
	在类型的判断上需要使用正确的方式
*/
func ReflectTypeTest1() {
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
func ReflectTypeTest2() {
	a := reflect.ArrayOf(10, reflect.TypeOf(byte(0)))
	m := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0))
	log.Println(a, m)
}

/*
	匿名字段结构显示
*/
func ReflectTypeTest3() {
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

// 这里必须使用可导出方法（首字母大写），否则NumMethod()为0
func (A101) A101v()  {}
func (*A101) A101p() {}
func (B101) B101v()  {}
func (*B101) B101p() {}

func ReflectTypeTest4() {
	var b B101
	// &b 产生了一个 B101的指针，也就是 *B101
	t := reflect.TypeOf(&b)
	// t.Elem() 可以获取t底层指针值引用的值的类型
	log.Println(t.Elem())
	log.Println("Type of 't' is: ", t.Name()) // 指针类型没有名字
	log.Println("'t' is a kind of: ", t.Kind())
	log.Println("Type 't.Elem()' is: ", t.Elem().Name()) // 实际的类型是个结构体
	log.Println("'t.Elem()' is kind of: ", t.Elem().Kind())
	// 判断类型是否是个指针类型
	log.Printf("t is a pointer? [%t], t.Elem is a pointer? [%t]",
		t.Kind() == reflect.Ptr, t.Elem().Kind() == reflect.Ptr)
	/*
		reflect.Elem() 方法获取*指针*指向的元素类型,这个获取过程被称为取元素，
		如果对于*值*类型执行该方法，则会抛出：panic: reflect: call of reflect.Value.Elem on struct Value
		例如: t := reflect.TypeOf(b)
		      t.Elem() <-- panic
		等效于对指针类型变量做了一个*操作（取值）
		s为一个数组，数组元素类型为reflect.Type，里面的两个元素分别是 t【指针类型】 和 t.Elem()【指针指向的实际类型】
	*/
	s := []reflect.Type{t, t.Elem()}
	log.Println(s)
	for _, t := range s {
		log.Println(t, " :")
		// 指针类型没有 NumField(),
		// println(t.NumField()) // panic: reflect: NumField of non-struct type *test1.B101
		for i := 0; i < t.NumMethod(); i++ {
			log.Println(" ", t.Method(i))
		}
	}
	log.Println(t.Elem().NumField())
}

/*
	利用反射提取TAG，还可以自动分解，常用作 ORM 映射，或做数据格式验证
*/
type user101 struct {
	name string `field:"name" type:"varchar2(50)"`
	age  int    `field:"age" type:"int"`
}

func ReflectTypeTest5() {
	var user user101
	t := reflect.TypeOf(user)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		log.Println(f.Tag.Get("field"), f.Tag.Get("type"))
	}
}

/*
	辅助判断方法 Implements、ConvertibalTo、AssignableTo 都是运行期间进行动态调用和赋值做必须的
*/
type X101 int

func (X101) String() string {
	return ""
}
func ReflectTypeTest6() {
	var a X101
	t := reflect.TypeOf(a)
	/*
		Implements 参数必须是一个接口类型，接口类型，可以使用nil构造接口类型
		(fmt.Stringer)(nil)为nil type类型
		(fmt.Stringer)(nil)：panic: reflect: nil type passed to Type.Implements
		st := reflect.TypeOf((fmt.Stringer)(nil))
		// 指针类型是一个独立的类型，在这里并不是interface类型,Implements必须传接口类型
		(*fmt.Stringer)(nil): panic: reflect: non-interface type passed to Type.Implements
		st := reflect.TypeOf((*fmt.Stringer)(nil))
		// 通过nil接口指针获取真实的类型
		st := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	*/
	st := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	log.Println("Type 't' is type of 'st'? ", t.Implements(st))

	// ConvertibleTo 判断一个类型的值是否可以转换为另一个类型
	it := reflect.TypeOf(0)
	log.Println(t.ConvertibleTo(it))
	// AssignableTo 判断一个类型的值是否可以赋值给参数指定的类型
	log.Println(t.AssignableTo(st), t.AssignableTo(it))
}
