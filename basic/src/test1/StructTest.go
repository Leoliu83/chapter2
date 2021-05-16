package test1

import (
	"log"
	"os"
	"reflect"
	"unsafe"
)

/*
	结构体内存布局：
	不管结构体含有多少个字段，内存都是一次性分配的，各个字段在相邻的地址空间*按定义*顺序排列
	对于引用类型（map，slice，channel），字符串，和指针,结构内存中只包含其基本（头部）数据，所有匿名成员也被包含在内。
*/
/*
	MixedFractions 表示带分数结构体，numerator 分子，denominator 分母
	也可简写为：
	type MixedFractions struct {
		numerator,denominator int32
	}
*/
type MixedFractions struct {
	wholeNum    int32 // 整数部分
	numerator   int32 // 真分数部分的分子
	denominator int32 // 真分数部分的分母
}

/*
	·结构体字段名必须唯一
	·可以使用 _ 进行补位
	·可以使用自身类型指针成员变量
	·使用 _ struct{} 强制使用命名初始化结构体
*/
type node struct {
	_    int // 补位
	id   int
	next *node    // 自身类型指针成员变量
	_    struct{} // 强制使用命名初始化结构体
}

// 在结构体中定义结构体属性
type file struct {
	fid   int
	fname string
	attr  struct {
		owner  int
		permit int
	}
}

type cannotcomp struct {
	id int
	m  map[int]int
}

// StructTest 方法用来测试结构体
func StructTest() {
	// 结构体初始化方法1, 只要定义了结构体变量，结构体就会被初始化，并且所有的属性都有了默认值
	var f MixedFractions
	log.Printf("f struct addr: %p, value: %+v", &f, f)
	f.denominator = 2
	f.numerator = 1
	f.wholeNum = 1
	log.Printf("%+v,%T,%d \n", f, f, unsafe.Sizeof(f))
	// 结构体初始化方法2
	f = MixedFractions{1, 2, 4}
	// %+v 可以打印结构体
	log.Printf("%+v,%T,%d \n", f, f, unsafe.Sizeof(f))

	// 命名初始化
	n1 := node{id: 1, next: nil}
	// 顺序初始化,这里使用了 _ struct{} 无法顺序初始化
	// n2 := node{1, 1, nil}
	log.Println("n1: ", n1)

	// 定义匿名结构体
	n3 := struct {
		name string
		age  int
	}{
		"leo",
		30,
	}
	log.Println("n3: ", n3)

	// 初始化结构体中的结构体属性
	/*
		// 错误的初始化例子：
		n4 := file{
			fid:   1,
			fname: "d:\\123.txt",
			attr: {
				1,
				1,
			},
		}
	*/
	// 正确的初始化方式：
	n4 := file{
		fid:   1,
		fname: "d:\\123.txt",
	}
	n4.attr.owner = 1
	n4.attr.permit = 7

}

// 只有所有属性都支持比较才可以进行比较
// slice，map和function都不支持比较
func StructCompareTest() {
	// 比较
	n5 := MixedFractions{
		wholeNum:  1,
		numerator: 2,
	}

	n6 := MixedFractions{
		wholeNum:  2,
		numerator: 4,
	}

	n7 := MixedFractions{
		wholeNum:  1,
		numerator: 2,
	}

	println(n5 == n6)
	println(n5 == n7)

	// 增加打印，不显示unused告警
	log.Printf("%+v", cannotcomp{id: 1, m: map[int]int{1: 1, 2: 2}})
	/*
		// cannot compare n8 == n9 (operator == not defined for cannotcomp)
		n8 := cannotcomp{
			id: 1,
		}

		n9 := cannotcomp{
			id: 2,
		}

		println(n8 == n9)
	*/
}

/*
	可以直接使用指针操作结构体字段，但不能使用多级指针
*/
func StructPointTest() {
	type user struct {
		id   int
		name string
	}

	p := &user{id: 1, name: "leo"}
	p.name = "marry"
	p.id = 2
	log.Printf("point address: %p, struct address: %p, value: %+v", &p, p, p)

	/*
		// 二级指针报错：p1.id undefined (type **user has no field or method id)
		p1 := &p
		*p1.id = 1
	*/
}

/*
	空结构体
*/
func StructEmptyTest() {
	// 空结构体内存使用为0
	var s struct{}
	var s1 [100]struct{}
	log.Printf("size of struct [s] is: %d", unsafe.Sizeof(s))
	log.Printf("size of [100]struct [s1] is: %d", unsafe.Sizeof(s1))

	// 所有长度为0的对象，通常都指向runtime.zerobase变量
	s2 := s1[:]
	s1[1] = struct{}{}
	s2[0] = struct{}{}
	s3 := [0]int{}
	log.Printf("size of struct [s1] is: %d, len: %d, cap: %d, address: %p, %T", unsafe.Sizeof(s1), len(s1), cap(s1), &s1, s1)
	log.Printf("size of struct [s2] is: %d, len: %d, cap: %d, address: %p, %T", unsafe.Sizeof(s2), len(s2), cap(s2), &s2, s2)
	log.Printf("s1[1] address: %p", &s1[0])
	log.Printf("s2[0] address: %p", &s2[0])
	log.Printf("s3 address: %p", &s3)

	// 空结构体可用作通道元素类型，用于事件通知
	exit := make(chan struct{})
	go func() {
		println("Hello!")
		exit <- struct{}{}
	}()
	<-exit
	println("end.")
}

/*
	匿名字段，指没有名字，仅有类型的字段，也被称为嵌入字段或者嵌入类型
	A field declared with a type but no explicit field name is an anonymous field,
	also called an embedded field or an embedding of the type in the struct.
	An embedded type must be specified as a type name T or as a pointer to a non-interface type name *T,
	and T itself may not be a pointer type. The unqualified type name acts as the field name.
	未命名类型没有名字标识，因此不能作为匿名字段
	go里有两种类型的类型
	1. 命名类型
		类型可以通过标识符来表示，这种类型称为命名类型（ Named Type ）。
		Go 语言的基本类型中有 20 个预声明简单类型都是命名类型，
		Go 语言还有一种命名类型一一用户自定义类型。
	2. 未命名类型
		一个类型由预声明类型、关键字和操作符综合决定，例如数组由数组元素的类型，长度确定是否为同一类型，[2]int 和 [3]int 的变量就不属于同一类型
		这个类型称为未命名类型（ Unamed Type ）。未命名类型又称为类型字面量（ Type Literal ）。
		Go 语言的基本类型中的复合类型：数组（ array ）、切片（ slice ）、字典（ map ）、通道（ channel ）、指针（ pointer ） 、函数字面量（ function ）、结构（ struct ）和接口（ interface ）都属于类型字面量，也都是未命名类型。
		所以 *int , []int , [2]int , map[k]v 都是未命名类型。
*/
func AnonymousFiledTest() {
	type attr struct {
		perm int
	}

	type file1 struct {
		name string
		attr // 仅有类型名
	}

	// 引入外部包类型的匿名字段
	type data struct {
		os.File
	}

	// 除了接口指针和多级指针以外的任何命名类型都可以作为匿名字段
	type newdata struct {
		*int
		string
		// int //不能将基础类型和其指针类型同时嵌入，因为两者隐式名字都是 int
	}

	f := file1{
		name: "test.txt",
		attr: attr{ // 显式的初始化匿名字段
			perm: 0755,
		},
	}
	f.perm = 0644                               // 直接设置匿名字段成员
	log.Printf("f.perm: %d, f: %+v", f.perm, f) // 直接读取匿名字段成员

	// 隐式字段名不包含包名
	dd := data{
		File: os.File{}, // 不包含包名
	}
	log.Printf("f.File: %v, f: %#v", dd.File, dd)

	x := 100
	nd := newdata{
		int:    &x,
		string: "abc",
	}
	log.Printf("nd.int: %d, nd: %#v", *nd.int, nd)

	// type a *int
	// type b **int
	type c interface{}
	type d []int64

	type s struct {
		// a // 不可以是指针
		// *a // 不可以是指针的指针
		// b  // 同上
		// *c // 不可以是接口的指针
		c // 可以是接口
		// []int64 // 不可以是slice，也不可以使别名，因为slice是未命名类型
		d       // 通过type 将slice变为了命名类型，就可以使用了
		Command string
		// 这样使用匿名字段，s结构体就拥有了 log.Logger的所有方法，例如Println等等
		*log.Logger
	}
	// 初始化后便可以直接使用 Pringln 进行打印
	s1 := &s{d: []int64{1, 2, 3},
		c:       struct{}{},
		Command: "test",
		Logger:  log.New(os.Stderr, "Job: ", log.Ldate),
	}
	s1.Println("123...")
	// s1.Printf("%+v \n", s1.d)

	/*
		虽然可以像普通成员那样访问匿名字段成员，但会存在重名
		编译器会从当前显示命名字段开始，逐步向内查找匿名字段成员
		例如 newuser中的name 和 address中的name冲突，因此在使用的时候，必须显示字段名 newuser.address.name
		当多个字段包含相同的成员名，必须使用显示字段名
	*/
	type address struct {
		name   string
		addrid int
		no     string
	}
	type idcard struct {
		cname string
		no    string
	}
	type newuser struct {
		name string
		address
		idcard
	}

	nu := newuser{
		name:    "leo",
		address: address{"test", 1, "0001"},
	}

	nu.name = "Leo"
	nu.addrid = 10              // 不需要显示字段名
	nu.address.name = "Address" // name 重名，因此需要显示字段名
	nu.address.no = "002"       // 与icard中的no重名，因此需要显示字段名
	nu.idcard.no = "0x123123123"
	nu.idcard.cname = "Leo"
	log.Printf("%+v", nu)

}

/*
	tag是对字段进行描述的元数据，常常被用作格式校验，数据库关系映射等。可以在反射时拿到
	struct tag 的格式应该是合法的键值对的格式，否则 GET 将无法获取到值。
	格式形如：`json1:"foo" json2:"bar"`(
		1.首尾使用符号'`'
		2.':'左右没有空格
		3.key需要用""号包起来
		4.两个k-v之间用空格隔开，不是逗号或者其他符号
	)
	虽然tag不是合法键值对格式在编译期间不会有异常，通过go vet可以发现
	下面的结构体有两条告警信息：
		·struct field tag `x` not compatible with reflect.StructTag.Get: bad syntax for struct tag
		·struct field tag `y` not compatible with reflect.StructTag.Get: bad syntax for struct tag
	但编译和初始化及使用均无异常。
	sum 后是合法的tag
	struct tag 是类型的组成部分，例如下面的 stag1 和 stag2 就不是同一类型
*/
func StructTagTest() {
	type stag struct {
		x   int `x`
		y   int `y:1`
		z   int `z1:"1" z2:"2"`
		sum int `sum:"y+x"`
	}

	st := stag{
		x:   1,
		y:   2,
		z:   1,
		sum: 1 + 2,
	}
	log.Printf("st: %#v", st)

	// stag1 和 stag2 不是同一类型，因为tag不同
	var stag1 struct {
		z   int `z1:"1" z2:"2"`
		sum int `sum:"y+x"`
	}
	var stag2 struct {
		z   int
		sum int
	}
	/*
		// cannot compare stag1 == stag2 (mismatched types struct{z int "z1:\"1\" z2:\"2\""; sum int "sum:\"y+x\""} and struct{z int; sum int})
		log.Println(stag1 == stag2)
	*/
	log.Println(stag1)
	log.Println(stag2)

	r := reflect.ValueOf(stag1)
	t := r.Type()
	for i, n := 0, t.NumField(); i < n; i++ {
		log.Printf("Tag is : [%s],Value is: %+v", t.Field(i).Tag, t.Field(i))
	}

}

func StructMemoryTest() {
	type point struct {
		x, y int
	}

	type value struct {
		id   int
		name string
		data []byte
		next *value
		point
	}

	v := value{
		id:    1,
		name:  "test",
		data:  []byte{1, 2, 3, 4},
		point: point{x: 100, y: 100},
	}

	/*
		      |-------- name --------|------------- data ---------------|       |------ point --------|
		+-----+-----------+----------+----------+-----------+-----------+-------+----------+----------+
		|  id |  name.ptr | name.len | data.ptr |  data.len |  data.cap |  next |  point.x |  point.y |
		+-----+-----------+----------+----------+-----------+-----------+-------+----------+----------+
		0     8           16         24         32         40          48       56        64         72
	*/

	format := `
		v: %p ~ %x, size: %d, align: %d
		field   address                offset  size
		------+---------------------+---------+------
		  id      %p           %d        %d
		name      %p           %d        %d
		data      %p           %d       %d
		next      %p           %d       %d
		   x      %p           %d       %d
		   y      %p           %d       %d
	`
	/*
		说明：
		uintptr(unsafe.Pointer(&v)): 将v结构体的地址转换为 uintptr 便于数学运算
		unsafe.Sizeof(v): v结构体的内存占用大小
		unsafe.Alignof(v): 对齐指数

	*/
	log.Printf(format,
		&v, uintptr(unsafe.Pointer(&v))+unsafe.Sizeof(v), unsafe.Sizeof(v), unsafe.Alignof(v),
		&v.id, unsafe.Offsetof(v.id), unsafe.Sizeof(v.id),
		&v.name, unsafe.Offsetof(v.name), unsafe.Sizeof(v.name),
		&v.data, unsafe.Offsetof(v.data), unsafe.Sizeof(v.data),
		&v.next, unsafe.Offsetof(v.next), unsafe.Sizeof(v.next),
		&v.x, unsafe.Offsetof(v.x), unsafe.Sizeof(v.x),
		&v.y, unsafe.Offsetof(v.y), unsafe.Sizeof(v.y),
	)
}

/*
	内存对齐
*/
func StructMemoryAlgnment() {
	/*
		1. golang中内存以最长的*基础类型*宽度为标准。
		   基础类型：例如 string由指向底层数组的指针和标识长度的int组成，
		   基础类型就是指针（8 byte），长度（8 byte），所以string的对齐值是8 byte而不是16 byte
		2. golang的struct属性将小长度属性进行合并来对齐最长长度
		例如（用#代表对齐补位）：
		// b与c靠的最近，因此在b后补00来增加2 byte，最终与c(4 byte)对齐=> a(1 byte)+b(1 byte)+补位(2 byte)
		a = 1 byte
		b = 1 byte
		# = 2 byte
		c = 4 byte
		all: 8 byte
		+-----+--------+----------+
		|  a  |    b   |     c    |
		+-----+--------+----------+
		0     1        4          8
	*/
	v1 := struct {
		a byte
		b byte
		c int32 // 对齐宽度4
	}{}
	// v1 和 v11 只是属性换了位置，但从打印可以看出，v11占用内存是12字节：a(4 byet)+c(4 byte)+b(4 byte)
	// 因为无法将相邻的两个小长度属性进行合并对齐，导致了内存的增长，因此在设计struct时，可以从小字节属性写到大字节属性，以便优化struct内存
	v11 := struct {
		a byte
		c int32 // 对齐宽度4
		b byte
	}{}

	v2 := struct {
		a byte
		b byte // 对齐宽度1
	}{}

	v3 := struct {
		a byte
		b []int // 基础类型 int，对齐宽度8｛point，len，cap｝
		c byte
	}{}

	// i(8 byte) => a(1 byte) + b(1 byte) + c(1 byte) + d(1 byte) + e(1 byte) + f(1 byte) + g(1 byte) + h(1 byte)
	v4 := struct {
		a byte
		b byte
		c byte
		d byte
		e byte
		f byte
		g byte
		h byte
		i []int // 基础类型 int，对齐宽度8｛point，len，cap｝
	}{}

	log.Printf("v1 algn: %d, v2 algn: %d,v3 algn: %d",
		unsafe.Alignof(v1), unsafe.Alignof(v2), unsafe.Alignof(v3))
	log.Printf(`[v1] algn: %d,size of [v1]: %d, address of [v1]: %p, 
		size of [v1.a]: %d, address of [v1.a]: %p,  
		size of [v1.b]: %d, address of [v1.b]: %p,  
		size of [v1.c]: %d, address of [v1.c]: %p`,
		unsafe.Alignof(v1), unsafe.Sizeof(v1), &v1,
		unsafe.Sizeof(v1.a), &v1.a,
		unsafe.Sizeof(v1.b), &v1.b,
		unsafe.Sizeof(v1.c), &v1.c)
	log.Printf(`[v11] algn: %d,size of [v11]: %d, address of [v11]: %p, 
		size of [v11.a]: %d, address of [v11.a]: %p,  
		size of [v11.b]: %d, address of [v11.b]: %p,  
		size of [v11.c]: %d, address of [v11.c]: %p`,
		unsafe.Alignof(v11), unsafe.Sizeof(v11), &v11,
		unsafe.Sizeof(v11.a), &v11.a,
		unsafe.Sizeof(v11.b), &v11.b,
		unsafe.Sizeof(v11.c), &v11.c)
	log.Printf(`[v2] algn: %d,size of [v2]: %d, address of [v2]: %p, 
		size of [v2.a]: %d, address of [v2.a]: %p,  
		size of [v2.b]: %d, address of [v2.b]: %p`,
		unsafe.Alignof(v2), unsafe.Sizeof(v2), &v2,
		unsafe.Sizeof(v2.a), &v2.a,
		unsafe.Sizeof(v2.b), &v2.b)

	log.Printf(`[v3] algn: %d,size of [v3]: %d, address of [v3]: %p, 
		size of [v3.a]: %d, address of [v3.a]: %p,  
		size of [v3.b]: %d, address of [v3.b]: %p,  
		size of [v3.c]: %d, address of [v3.c]: %p`,
		unsafe.Alignof(v3), unsafe.Sizeof(v3), &v3,
		unsafe.Sizeof(v3.a), &v3.a,
		unsafe.Sizeof(v3.b), &v3.b,
		unsafe.Sizeof(v3.c), &v3.c)

	log.Printf(`[v4] algn: %d,size of [v4]: %d, address of [v4]: %p, 
		size of [v4.a]: %d, address of [v4.a]: %p,  
		size of [v4.b]: %d, address of [v4.b]: %p,  
		size of [v4.c]: %d, address of [v4.c]: %p,
		size of [v4.d]: %d, address of [v4.d]: %p,
		size of [v4.e]: %d, address of [v4.e]: %p,
		size of [v4.f]: %d, address of [v4.f]: %p,
		size of [v4.g]: %d, address of [v4.g]: %p,
		size of [v4.h]: %d, address of [v4.h]: %p,
		size of [v4.i]: %d, address of [v4.i]: %p,
		`,
		unsafe.Alignof(v4), unsafe.Sizeof(v4), &v4,
		unsafe.Sizeof(v4.a), &v4.a,
		unsafe.Sizeof(v4.b), &v4.b,
		unsafe.Sizeof(v4.c), &v4.c,
		unsafe.Sizeof(v4.d), &v4.d,
		unsafe.Sizeof(v4.e), &v4.e,
		unsafe.Sizeof(v4.f), &v4.f,
		unsafe.Sizeof(v4.g), &v4.g,
		unsafe.Sizeof(v4.h), &v4.h,
		unsafe.Sizeof(v4.i), &v4.i,
	)

	/*
		空结构体属性, 如果空结构体在最后一个字段，
		那么编译器会将其作为长度为1的类型做对齐处理, 以便其地址不会越界, 避免引发垃圾回收错误
		v5 长度为  a(0 byte)+b(8 byte)+c(8 byte)
		c 视为长度为1的属性（例如 byte），对齐至8
	*/
	var v5 = struct {
		a struct{} // 长度为0 不对齐
		b int      // 长度为8
		c struct{} // 视为长度为1的属性（例如 byte），对齐至8
	}{}
	log.Printf(`[v5] algn: %d,size of [v5]: %d, address of [v1]: %p, 
		size of [v5.a]: %d, address of [v5.a]: %p, offset [v5.a]: %d,
		size of [v5.b]: %d, address of [v5.b]: %p, offset [v5.b]: %d,
		size of [v5.c]: %d, address of [v5.c]: %p, offset [v5.c]: %d
		`,
		unsafe.Alignof(v5), unsafe.Sizeof(v5), &v5,
		unsafe.Sizeof(v5.a), &v5.a, unsafe.Offsetof(v5.a),
		unsafe.Sizeof(v5.b), &v5.b, unsafe.Offsetof(v5.b),
		unsafe.Sizeof(v5.c), &v5.c, unsafe.Offsetof(v5.c))
	/*
		如果只有一个空结构体，则同样按1对齐，但长度为0，且指向runtime.zerobase
	*/
	var v6 = struct {
		a struct{} // 视为长度为1的属性（例如 byte）
	}{}
	log.Printf(`[v6] algn: %d,size of [v6]: %d, address of [v6]: %p, 
		size of [v6.a]: %d, address of [v6.a]: %p, offset [v6.a]: %d`,
		unsafe.Alignof(v6), unsafe.Sizeof(v6), &v6,
		unsafe.Sizeof(v6.a), &v6.a, unsafe.Offsetof(v6.a),
	)
}
