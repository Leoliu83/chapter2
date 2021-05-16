package test1

import (
	"fmt"
	"log"
	"reflect"
	"unsafe"
)

/*
	和Type获取的信息不同，Value专注于对象实例数据的读写
*/
type obj102 struct {
	x    int      // unexported field
	Y    int      // exported field
	Name string   // exported field
	code int      // unexported field
	P    *int     // 指针
	_    struct{} // 强制使用属性名进行初始化（复习）
}

/*
	复习：
		当把一个对象赋值给一个接口类型变量时，会复制该对象，并且该对象不可寻址，不可被修改
		如果需要修改目标对象，必须使用指针
	*由于ValueOf(i interface{}) 参数是一个 interface{} 因此将对象传递给ValueOf就相当于赋值给了接口变量！
	ifaceV.(obj102).y = 200 // 编译错误
	同样需要 .Elem() 获取目标对象，因为被接口本身存储的指针本身是不能寻址和进行设置操作的
	reflect.Value.Pointer 和 reflect.Value.Int 等方法类似，将 reflect.Value.data 存储的数据转换为指针，目标必须是指针类型
*/
func ReflectValueTest1() {
	a := 1
	log.Printf("'a' address: %p", &a)
	obj1 := obj102{x: 1, Y: 2, P: &a}
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
	obj2 := new(obj102) // 等同于 var o obj102; obj2=&o
	obj2.P = &a
	v := reflect.ValueOf(obj2).Elem()
	name := v.FieldByName("Name")
	code := v.FieldByName("code")
	// 只能*直接设置*导出(exported) 字段，不能*直接设置*非导出 (unexported) 字段
	log.Println("Field 'Name' can be set? ", name.CanSet())   // true
	log.Println("Field 'code' can be set? ", code.CanSet())   // false
	log.Println("Field 'Name' addressable? ", name.CanAddr()) // true
	log.Println("Field 'code' addressable? ", code.CanAddr()) // false

	if name.CanSet() {
		name.SetString("Leoliu")
	}
	if code.CanSet() {
		// 如果设置前不执行判断，则会出现如下错误
		// panic: reflect: reflect.Value.SetInt using value obtained using unexported field
		code.SetInt(1001)
	}
	log.Printf("%+v", obj2)
	// 通过unsafe包，可以对addressable的字段进行间接设置
	if code.CanAddr() {
		// fmt.Sprintf("%#x",123) 可以将10进制转成16进制，#会带上0x
		log.Println(code.UnsafeAddr(), fmt.Sprintf("%#x", (int)(code.UnsafeAddr())))
		// UnsafeAddr 可以返回任何 CanAddr的Value.data地址（相当于&取址操作），
		// 比如执行Elem()之后的Value，以及字段成员地址
		// 以结构体中得指针字段为例，Pointer返回该字段所保存(指向)的地址，
		// 而UnsafeAddr返回该字段本身的地址(结构体对象地址+偏移量)
		*(*int)(unsafe.Pointer(code.UnsafeAddr())) = 100
	}
	fa := v.FieldByName("P")
	log.Printf("Address of field 'a' is: %#x, field 'a' pointer to: %#x", fa.Pointer(), fa.UnsafeAddr())
	log.Printf("%+v", obj2)

}

/*
	使用Interface方法做类型推断和转换
*/
func ReflectValueTest2() {
	type user struct {
		Name string
		Age  int
	}

	u := user{
		"leo",
		30,
	}

	v := reflect.ValueOf(&u)
	if !v.CanInterface() {
		println("Can interface failed!")
		return
	}
	// v.Interface() 等同于 var i interface{} = v
	// 也就是将对象赋值给接口变量（复习：如果使用的值，那么接口变量中会保存对象的副本，且无法修改！）
	// 也可以使用Value.Int Value.Bool 等方法进行类型转换，但是失败时候会引发panic，而且不能使用ok-idiom模式
	p, ok := v.Interface().(*user)
	if !ok {
		println("Interface failed!")
		return
	}

	p.Age++
	log.Printf("%+v", u)
}

/*
	复合类型对象以及nil对象
*/
func ReflectValueTest3() {
	c := make(chan int, 4)
	// 引用类型都是指针，因此不需要 & 取址
	v := reflect.ValueOf(c)
	// v.TrySend(x) 方法试图在通道v上发送一个x，且不会阻塞
	// TrySend 参数必须是个 reflect.Value
	if v.TrySend(reflect.ValueOf(100)) {
		// v.TryRecv() 方法试图从通道v接收数据，且不会阻塞（ok-idiom模式）
		log.Println(v.TryRecv())
	}

	var a interface{} = nil
	// 接口有两种nil状态
	// 复习，由于b中包含了类型信息，因此b不等于nil
	var b interface{} = (*int)(nil)
	println(a == nil, b == nil)
	// 由于a是真正的nil，因此不能使用 reflect.ValueOf(a).IsNil() 会引发panic: reflect: call of reflect.Value.IsNil on zero Value
	println(reflect.ValueOf(b).IsNil()) // true
	// 也可以使用unsafe转换后直接判断iface.data 是否为0值
	// 复习：接口类型实际上保存了真实类型和数据指针, 详见 InterfaceTest.go
	iface := (*[2]uintptr)(unsafe.Pointer(&b))
	log.Println(b, iface, iface[1] == 0)

	// Value里的方法并未都实现ok-idiom 方法，或者返回error，因此需要自己判断是否是zero value
	v = reflect.ValueOf(struct{ name string }{})
	log.Printf("%+v", v)
	println("name field is valid? ", v.FieldByName("name").IsValid())
	println("xxx field is valid? ", v.FieldByName("xxx").IsValid())
}
