package test1

import (
	"log"
	"unicode/utf8"
)

/*
	Unicode 码点（code point）,它是int32的别名，相当于UCS-4/UTF-32编码格式，
	使用单引号的字面量，其默认类型就是rune
	字符串中字符的数量可以由 utf8.RuneCountInString 获取
*/
func UnicodeTest() {
	r := '我'
	log.Printf("%T \n", r)

	s := "我是谁（who am i）"
	log.Printf("字符串[s]的字节长度为: %d,字符数量为: %d", len(s), utf8.RuneCountInString(s))
}

/*
	可以直接在 rune、byte、string 之间进行转换
	string不能直接转rune
*/
func UnicodeTransformTest() {
	r := '我'
	// rune 转 string
	s := string(r)
	/*
		rune 转 byte
		由于 rune 是 int32，而byte是int8 在强转的时候，将会值保留第8位
		例如：
			'我'(rune) -> 25105(int) -> 110 0010 0001 0001(binary)
			转byte后：17(int) -> 0001 0001(binary)
	*/
	b := byte(r)
	// byte 转 string
	s1 := string(b)
	// byte 转 rune
	r1 := rune(b)

	log.Printf("变量[r]的类型是: %T, 值是: %d \n", r, r)
	log.Printf("变量[s]的类型是: %T, 值是: '%s', 字节是: %d \n", s, s, []byte(s))
	log.Printf("变量[b]的类型是: %T, 值是: %d \n", b, b)
	log.Printf("变量[s1]的类型是: %T, 值是: %s \n", s1, s1)
	log.Printf("变量[r1]的类型是: %T, 值是: %c \n", r1, r1)

	i := 25105
	b1 := (byte)(i)
	log.Printf("变量[b1]的类型是: %T, 值是: %d \n", b1, b1)
}
