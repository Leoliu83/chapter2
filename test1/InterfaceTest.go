package test1

import (
	"fmt"
	"log"
)

/*
	接口表示一种调用契约，是多个方法声明的集合
	在某些动态语言里，接口（interface）也被称之为协议（protocol）。
	准备交互的双方，共同遵守事先约定的规则，使得在无需知道对方身份的情况下进行协作。
	接口要实现的是做什么，而不关心怎么做，谁来做。
	接口除了类型依赖，有助于减少用户可视方法，屏蔽内部结构和实现细节，但接口的实现机制会有运行时开销。
	对于相同的包，不会频繁变化的内部模块之间，并不需要抽象出接口来强行分离。
	接口最常见的使用场景是对包外提供访问，或者预留扩展空间。
*/
/*
	go接口的内部实现：
		type iface struct{
			tab *itab // 类型信息
			data unsafe.Pointer // 实际对象指针
		}
	限制：
		·不能有字段
		·不能定义自己的方法
		·只能声明方法，不能实现
		·可嵌入其他接口类型
*/
/*
	go接口运行机制：
	接口使用一个名为itab的结构存储运行期间所需的相关类型信息。
	itab结构体如下：
	type itab struct{
		inter *interfacetype // 接口类型
		_type *_type     	 // 实际对象类型
		fun   [1]uintptr 	 // 实际对象方法地址
	}
	详细《Go语言学习笔记》p146
*/
// InterfaceTest 测试接口类型
func InterfaceTest() {
	var o SubscriberOne
	o.notice()

	var t testerImpl
	// var t1 tester = t  //cannot use t (variable of type testerImpl) as tester value in variable
	var t1 tester = &t
	t1.test()
	println(t1.string())

	// interface{} 表示空接口，类似于java中的 Object
	var it1, it2 interface{}
	println(it1 == nil, it1 == it2)
	it1, it2 = 100, 100
	println(it1 == it2)
	it1, it2 = map[string]int{}, map[string]int{}
	println(it1 == it2) // error  map不能比较，只能和nil比较

	var nti newtesterImpl
	// receiver 为 value 包含了所有接口的方法，因此这里可以不适用指针赋值
	var nti1 newtester = nti
	p(nti1) // 隐式转换为子集接口（stringer.string(SB)）

}

/*
	支持匿名接口，有点绕
*/
func AnonymousInterfaceTest() {
	// 定义一个变量d 是匿名接口类型 'interface {string() string}'
	var d interface {
		string() string
	} = stringerImpl{} // 给这个变量赋值为 stringerImpl结构体实例
	// 创建一个 anonymous 实例，其属性 data 是一个匿名接口，将变量d赋值给该属性
	n := anonymous{data: d}
	// 调用n的data属性的 string() 方法
	println(n.data.string())
}

// 将对象赋值给接口变量时，会复制该对象
func InterfaceInternalTest() {
	type outdata struct {
		x int
	}

	d := outdata{100}
	log.Printf("variabel[d] address: %p, value: %+v", &d, d)
	var t interface{} = d
	// t.(outdata) 其实就是获取了接口对象中的outdata的副本
	println(t.(outdata).x)
	// println(t.x)
	// 不能直接读取t.(outdata) 这个 outdata{100} 的副本
	// p := &t.(outdata) // invalid operation: cannot take address of t.(outdata)
	// t.(outdata).x = 200 // cannot assign to t.(outdata).x (value of type int)
	d.x++
	log.Printf("variabel[d] type: %T, address: %p, value: %+v", d, &d, d)
	log.Printf("variabel[t1] type: %T, address: %p, value: %+v", t, &t, t)
	// 下一个打印会产生一个告警，由于 t.(outdata) 是不可寻址的，因此 %p 会产生告警
	log.Printf("variabel[t.(outdata)] address: %p, value: %+v", t.(outdata), t.(outdata))
	// 解决办法就是把对象的指针赋值给接口变量，这样接口变量中就保存的是指针的副本
	log.Println("------------------------ use pointer ----------------------------")
	var t1 interface{} = &d
	p := t1.(*outdata)
	p.x = 10
	t1.(*outdata).x = 11
	log.Printf("variabel[d] type: %T, address: %p, value: %+v", d, &d, d)
	log.Printf("variabel[t1] type: %T, address: %p, value: %+v", t1, t1, t1)
	log.Printf("variabel[t1.(*outdata)] type: %T, address: %p, value: %+v", t1.(*outdata), t1.(*outdata), t1.(*outdata))

}

/*
	只有当接口变量内部的两个指针（itab,data）都为nil时，接口变量才为nil
*/
func NilInterfaceTest() {
	var a interface{} = nil
	var b interface{} = (*int)(nil)
	println(nil == (*int)(nil))
	/*
		虽然两个都是nil, 并且 (*int)(nil) == nil 是 true
		但是 b 的 itab._type 中保留了类型信息(*int)，因此 b != nil
		这个特性使得在 函数返回error时候 特别容易出错
	*/
	println("a is nil? ", a == nil, " | ", "b si nil? ", b == nil)

	/* 自定义错误处理的-错误做法 */
	f := func(x int) (int, error) {
		// 这里的err有了类型，因此永远不是nil
		var err *TestError
		if x < 0 {
			err = new(TestError)
			x = 0
		} else {
			x += 100
		}
		return x, err
	}
	// 按照逻辑应该执行 x+=100
	x, e := f(100)
	// 但是下面的判断永远执行 log.Fatal(e)
	if e == nil {
		log.Println("f -> ", x)
	} else {
		log.Println("f err -> ", e)
	}
	/* 自定义错误处理的-正确做法 */
	f1 := func(x int) (int, error) {
		// 这里的err有了类型，因此永远不是nil
		var err *TestError
		if x < 0 {
			err = new(TestError)
			x = 0
			return x, err
		} else {
			x += 100
			// 这里显式的返回 nil
			return x, nil
		}
	}
	x, e = f1(100)
	if e == nil {
		log.Println("f1 -> ", x)
	} else {
		log.Fatal("f1 err -> ", e)
	}
}

// TODO 未完成
func InterfaceTypeTransform() {

}

/*-------------- 自定义error并且实现error接口 ------------------*/
type TestError struct{}

func (*TestError) Error() string {
	return "error"
}

/*--------------------------------------------------------------*/
// Subscriber 是一个接口类型，定义了订阅者需要实现的方法,其中包含一个需要实现的方法“通知”
type Subscriber interface {
	notice() bool
}

// SubscriberOne 表示一个订阅者
type SubscriberOne struct{}

func (s SubscriberOne) notice() bool {
	fmt.Println("I am SubscriberOne. I receive the notice.")
	return true
}

// 接口通常以er为名称后缀
type tester interface {
	test()
	string() string
}

/*
	编译器会根据方法集来判断是否实现了接口
	在这里，只有*testerImpl才符合tester的要求，testerImpl不符合要求
	原因：
		之前说过 *T 包含所有 receiver 为 T 和 *T 的方法（详见MethodTest.go）
*/
type testerImpl struct{}

func (testerImpl) test() {
	log.Println("test()")
}
func (*testerImpl) string() string {
	return "string()"
}

/*
	嵌入其他接口类型，相当于将其声明的方法集导入。
	·这就要求不能有同名方法，因为不支持重载。
	·不能嵌入自身或循环嵌入，会导致递归错误
*/
type stringer interface {
	string() string
	// stringer // illegal cycle in declaration of stringer
}

type stringerImpl struct{}

func (stringerImpl) string() string { return "" }

/*
	查看objdump
		TEXT %22%22.newtester.string(SB) gofile..<autogenerated>
		TEXT %22%22.newtester.test(SB) gofile..<autogenerated>
		TEXT %22%22.(*newtesterImpl).string(SB) gofile..<autogenerated>
		TEXT %22%22.(*newtesterImpl).test(SB) gofile..<autogenerated>
		TEXT %22%22.stringer.string(SB) gofile..<autogenerated>
	可以发现，编译器实际为newtester也生成了string方法
		·%22%22.stringer.string(SB)
		·%22%22.newtester.string(SB)
	这两个方法同时存在
*/
type newtester interface {
	stringer // 嵌入其他接口，stringer是接口名，不是属性名，这是个匿名属性，和匿名接口区分开
	test()
}

type newtesterImpl struct{}

// 这里 receiver 为 value 定义两个方法
func (nti newtesterImpl) test() {
	log.Println("newtest()")
}
func (nti newtesterImpl) string() string {
	log.Println("newstring()")
	return "newstring()"
}

func p(s stringer) {
	s.string()
}

/*
	在结构体内部可以定义匿名接口类型
*/
type anonymous struct {
	// 匿名接口类型，表示 anonymousInterface 属性 是一个 'interface {string() string}' 接口
	// anonymousInterface 是属性名，不是接口名，和匿名属性区分开
	data interface {
		string() string
	}
}
