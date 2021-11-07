package charset

import (
	"io"
	"log"
	"reflect"
	"sync"
)

/*
	asciiDetector is implemention of CharsetDetector
	manual create instance is not allowed, must use GetASCIIDetector() function
	@see CharsetDetector
*/
type asciiDetector struct{ name string }

// 保证某方法只执行一次
var onceASCIID sync.Once

// 私有 asciiDetector 结构体变量
var asciiD *asciiDetector

/*
	由于该结构体没有共享的变量，因此只需要在内存中存储一次
*/
func GetASCIIDetector() *asciiDetector {
	onceASCIID.Do(func() {
		log.Println("Initialize asciiDetector instance.")
		asciiD = &asciiDetector{name: "ASCII-Detector"}
	})
	return asciiD
}

/*
	ascii字符集探测
*/
func (detector asciiDetector) DetectReadSeekerWithSize(r io.ReadSeeker, size int) (Charset, error) {
	buf := make([]byte, size)
	n, err := r.Read(buf)
	// 如果错误不为空且不是io.EOF说明出现了异常
	if err != nil && err != io.EOF {
		val := reflect.ValueOf(r)
		typ := reflect.Indirect(val).Type()
		log.Printf("Error when read []byte from Reader[%s]. \n ERROR:[%s]", typ, err)
		return UNKNOWN, err
	}

	// offset 找出当前读取的字节数，作为回退的偏移量，因为需要基于当前量回退，因此要乘以-1
	offset := -1 * int64(n)
	log.Printf("Offset is: %d", offset)
	// Seek 设置下一次 Read 或 Write 的偏移量为 offset
	// 它的解释取决于 whence：
	//     io.SeekStart 表示相对于文件的起始处
	//     io.SeekCurrent: 表示相对于当前的偏移
	//     io.SeekEnd 表示相对于其结尾处。
	// Seek 返回新的偏移量和一个错误。
	r.Seek(offset, io.SeekCurrent)
	for i := 0; i < n; i++ {
		if buf[i] > 0x7f {
			return UNKNOWN, nil
		}
	}
	return US_ASCII, nil
}

/*
	默认校验，读取4096个字节
*/
func (detector asciiDetector) DetectReadSeekerDefault(r io.ReadSeeker) (Charset, error) {
	return detector.DetectReadSeekerWithSize(r, 4096)
}
