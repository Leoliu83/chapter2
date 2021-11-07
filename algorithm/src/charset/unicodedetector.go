package charset

import (
	"io"
	"log"
	"reflect"
	"sync"
)

type unicodeDetector struct{ name string }

// 保证某方法只执行一次
var onceUnicodeD sync.Once

// 私有 unicodeDetector 结构体变量
var unicodeD *unicodeDetector

/*
	由于该结构体没有共享的变量，因此只需要在内存中存储一次
*/
func GetUnicodeDetector() *unicodeDetector {
	onceUnicodeD.Do(func() {
		log.Println("Initialize UnicodeDetector instance.")
		unicodeD = &unicodeDetector{name: "Unicode-Detector"}
	})
	return unicodeD
}

func (detector unicodeDetector) DetectReadSeekerWithSize(r io.ReadSeeker, size int) (Charset, error) {
	buf := make([]byte, size)
	n, err := r.Read(buf)
	// 如果错误不为空且不是io.EOF说明出现了异常
	if err != nil && err != io.EOF {
		val := reflect.ValueOf(r)
		typ := reflect.Indirect(val).Type()
		log.Printf("Error when read []byte from Reader[%s]. \n ERROR:[%s]", typ, err)
		return UNKNOWN, err
	}
	log.Print(n)
	return UTF_8, nil
}
