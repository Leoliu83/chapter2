package test1

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"reflect"
	"unsafe"
)

type Student struct {
	Sid   int
	Sname string
}

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
	//  将unsafe.Pointer 转换为 uintptr(其实就是一个uint，用于存放指针地址)
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
	// 下面语句会有编译错误：cannot take address of (s[0])
	// log.Printf("字符串[s]的首字节是: %p", &(s[0]))

	log.Printf("%s", (*reflect.StringHeader)(unsafe.Pointer(&s1)))
}

/*
	转码 UTF8 -> GBK
*/
func TransCharacter() {
	s1 := "你好??"
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
