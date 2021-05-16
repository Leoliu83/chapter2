package test1

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
	"unsafe"
)

type Student struct {
	Sid   int
	Sname string
	_     struct{} // 表示在初始化结构体时，必须显示的写出属性名，即不允许{1,"leo"}这种写法，必须是{Sid:1,Sname:"leo"}
}

/*
	go中string的结构如下:
	type stringStruct struct {
		// 头部指向字节数组
		str unsafe.Pointer
		// 长度
		len int
	}
	字符串操作通常会在堆上分配内存，届时会有大量的字符串对象需要进行垃圾回收
	可以使用[]byte缓存池，或在栈上自行拼装等方式来实现 zero-gargage
*/
func StringTest() {
	/*
		string 字符串允许:
		16进制: \x61
		8进制: \142
		unicode: \u0041
	*/
	s := "测试\x61\142\u0041"
	log.Printf("% x \n %s \n %T \n", s, s, s)

	/*
		默认string的值为 "" 而不是nil
	*/
	var s1 string
	log.Printf("% x \n %s \n %T \n", s1, s1, s1)

	/*
		使用 `` 处理不需要转义的字符串，而且支持换行
	*/
	s2 := `abcde \n \s \d \r\n
	换行
不带空格换行`
	log.Printf("%s \n", s2)
}

func StringForLoopTest() {
	s := "测试\x61\142\u0041"
	/*
		基于byte的遍历
	*/
	for i := 0; i < len(s); i++ {
		log.Printf("%d,%c \n", i, s[i])
	}
	log.Println("==================================")
	/*
		基于字符（rune）的遍历
		返回 i(数组序号),c(unicode字符)
	*/
	for i, c := range s {
		log.Printf("%d,%c,%c \n", i, s[i], c)
	}

}

/*
	go的字符串头结构和数组的头结构*部分相同*
*/
func StringTransformTest() {
	s := "hello world!"

	// 字符串转数组，数组转字符串，转换一定会导致重新分配内存
	/*
		%x 为 s 变量所表示的地址
		%p 为 s 指向的最终地址，由于不是指针类型，因此没有指向地址，因此产生告警
	*/
	log.Printf("字符串变量[s]的地址: 0x%x \n", &s)
	log.Printf("字符串变量[s]的值: %s \n", s)
	//b1为指针类型(*string)
	b1 := &s
	/*
		%x 为b1变量所表示的地址为
		%p 为 s 指向的最终地址,也就是真正string值所在的地址
	*/
	log.Printf("字符串指针[b1]的地址: 0x%x \n", &b1)
	log.Printf("字符串指针[b1]所指向的真正的数据地址: %p \n", b1)
	log.Printf("字符串指针[b1]的值: %s \n", *b1)

	// 数组的地址就是首地址，也就是 &b2[0]
	b2 := []byte(s)
	// 变量地址值取不到，因为数组变量的地址就是首元素的地址
	log.Printf("数组变量[b2]的地址: 0x%x \n", &b2)
	log.Printf("数组变量[b2]的值: [% x] \n", b2)
	// 首元素的值和字符串地址的值不同，说明做了数组拷贝，因为字符串的地址也是底层数组首元素的值
	log.Printf("数组变量[b2<首元素>]的地址: 0x%x \n", &b2[0])
	log.Printf("数组变量[b2]的值: %s \n", b2)

	s = string(b2)
	log.Printf("字符串变量[s]的地址: 0x%x \n", &s)
	log.Printf("字符串变量[s]的值: %s \n", s)
	/*
		使转换不产生数组拷贝
	*/
	// 将s指针地址转换为 unsafe.Pointer
	p := unsafe.Pointer(&s)
	// 将unsafe.Pointer 转换为 uintptr(其实就是一个uint，用于存放指针地址)，并且可以进行数学运算
	// 这里只做演示，转换过程中可以不需要该步骤
	pp := uintptr(p)
	// 可以看到 pp的值其实就是 s的地址
	log.Printf("uintptr变量[pp]的地址: 0x%x \n", &pp)
	log.Printf("uintptr变量[pp]的值: %d \n", pp)
	// 将 unsafe.Pointer 类型强转成 字节数组指针
	bp := (*[]byte)(p)
	// 可以看到打印出来 bp 所指向的地址就是s的地址
	//  &(*bp)[0] 表示 bp取真实的数组，之后找数组第一个元素的地址
	log.Printf("字节数组指针[bp]的地址: 0x%x \n", &bp)
	log.Printf("字节数组指针[bp]所指向的真正的数据地址: %p \n", bp)
	log.Printf("字节数组指针[bp]的数组值: [% x] \n", *bp)
	log.Printf("字节数组指针[bp]的字符串值: %s \n", *bp)
	// go里允许直接利用下标index访问字符串的字节数组，但不能取地址
	log.Printf("字符串[s]的首字节是: %d", s[0])
	s1 := s[:3]
	s2 := s[5:]
	s3 := s[2:3]
	log.Printf("字符串[s]的字节数组是: [% x] , 字符串是: %s", s1, s1)
	log.Printf("字符串[s]的字节数组是: [% x] , 字符串是: %s", s2, s2)
	log.Printf("字符串[s]的字节数组是: [% x] , 字符串是: %s", s3, s3)
	// 下面语句会有编译错误: cannot take address of (s[0])
	// log.Printf("字符串[s]的首字节是: %p", &(s[0]))
	// 直接将s1的字符串地址转成Pointer，然后强转字符串（b2是字节数组）
	log.Printf("将b2直接转换成字符串: %s", *(*string)(unsafe.Pointer(&b2)))
	// 可以使用append直接将字符串添加到字节数组中去，但是由于string不可变，因此该操作会产生分配新内存并复制数据
	var b3 []byte
	b3 = append(b3, s...)
	log.Printf("数组[b3]地址为: %p,值为: [% x]", &b3, b3)
	/*
		编译器会对以下两种情况进行优化，以避免额外的分配和复制操作:
		  ·在将[]byte转换成string key，去map[string]中查询的时候
		  ·将string转换成[]byte,进行for range迭代的时候，直接取字节赋值给局部变量的时候
		@TODO 未验证
	*/
	m1 := map[string]string{
		"key1": "value",
	}
	key := []byte("key1")
	log.Printf("字节数组变量[key]的地址是: %p", &key[0])
	x, ok := m1[string(key)]
	log.Printf("map[\"key1\"]的值是: %s, 获取map[\"key1\"]的数据是否成功(bool): %t", x, ok)

	log.Printf("s string header: %#v", (*reflect.StringHeader)(unsafe.Pointer(&s)))
}

// 以+号的方式拼接字符串，每一次都要重新分配内存，构建大型字符串时候容易造成性能问题
func StringPlusTest() string {
	var s string
	for i := 0; i < 1000; i++ {
		s += "*"
	}
	return s
}

/*
	strings.Join 方法会统计所有参数长度，并一次性完成内存分配([]byte)
	之后再用copy方式拷贝数据到 []byte
*/
func StringJoinTest() string {
	s := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		s[i] = "*"
	}
	return strings.Join(s, "")
}

/*
	使用bytes.Buffer 拼接字符串，性能较好
*/
func StringByteBufTest() string {
	var b bytes.Buffer
	b.Grow(1000)
	for i := 0; i < 1000; i++ {
		b.WriteString("*")
	}
	return b.String()
}

/*
	转码 UTF8 -> GBK
*/
func TransCharacter() {
	s1 := "你好🧭"
	s2 := "你好"
	// vscode默认utf8
	// 可以看到 中文3个字节，表情4个字节，所以两个中文+表情一共是10个字节
	bp := (*[]byte)(unsafe.Pointer(&s1))
	log.Printf("[UTF8]->s1-> %x,%+v \n", &bp, bp)
	// 转码,转GBK
	bp = (*[]byte)(unsafe.Pointer(&s2))
	log.Printf("[UTF8]->s1-> %x,%+v \n", &bp, bp)
	reader := transform.NewReader(bytes.NewReader(([]byte)(s1)), simplifiedchinese.GBK.NewEncoder())
	newBytes, _ := ioutil.ReadAll(reader)
	result := string(newBytes)
	log.Printf("Result: %s,[% x]", result, newBytes)
}
