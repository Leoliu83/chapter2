package charset

import (
	"os"
	"strings"
	// "time"
	// "sync"
	"testing"
)

func TestGetASCIIDetector(t *testing.T) {
	// var wg sync.WaitGroup
	// for i := 0; i <= 1000; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		d := GetASCIIDetector()
	// 		t.Logf("%p", d)
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()
	d1 := GetASCIIDetector()
	// time.Sleep(2 * time.Second)
	d2 := GetUnicodeDetector()
	t.Logf("%p,%p", d1, d2)
	t.Logf("%+v,%+v", d1, d2)
}

func TestASCIIDetectorFile(t *testing.T) {
	buf := make([]byte, 1024)
	f, _ := os.Open("G:\\Users\\leoliu\\Desktop\\使用xterm.txt")
	// 先读取两个字节 验证Seek函数
	f.Read(make([]byte, 2))
	defer f.Close()
	detector := GetASCIIDetector()
	s, err := detector.DetectReadSeekerDefault(f)
	if err != nil {
		t.Log("FAILED: ", s)
	}
	n, _ := f.Read(buf)
	t.Logf("SUCCESS: (%d) %s", n, string(buf[0:n]))
}

func TestASCIIDetectorString(t *testing.T) {
	buf := make([]byte, 1024)
	r := strings.NewReader("ab我cde") // 我
	// 先读取两个字节 验证Seek函数
	r.Read(make([]byte, 2))
	detector := GetASCIIDetector()
	s, err := detector.DetectReadSeekerDefault(r)
	if err != nil {
		t.Log("FAILED: ", s)
	}
	n, _ := r.Read(buf)
	t.Logf("SUCCESS: %d,%s", n, string(buf[0:n]))
}

func TestASCIIDetector(t *testing.T) {
	m := make(map[string]bool)
	m["a"] = true
	t.Logf("%t,%t", m["b"], m["a"])
}
