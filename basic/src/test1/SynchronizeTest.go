package test1

import (
	"log"
	"net/http"
	"sync"
	"time"
)

/*
	通道用于解决逻辑层的并发处理架构，而锁在全局范围内保证数据安全
	将Mutex作为匿名字段时，相关方法必须实现为 pointer-receiver，否则会因为复制导致锁机制失效
*/
type data83 struct {
	sync.Mutex
}

type data83_1 struct {
	*sync.Mutex
}

/*
	当receiver为值类型时候（d data83），会导致锁失效
	因此必须是 pointer-receiver（d *data83）
	编译器会产生告警，可以通过告警发现问题
	也可以使用 *Mutex 来避免复制问题，但那需要专门初始化(在SynchronizeTest1中演示)
*/
func (d *data83) test(s string) {
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		log.Println(s, i)
		time.Sleep(time.Second)
	}
}

// data83_1中使用的是：*sync.Mutex，因此避免了使用 value-receiver 导致的锁失效
func (d data83_1) test(s string) {
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		log.Println(s, i)
		time.Sleep(time.Second)
	}
}

func SynchronizeTest() {
	var wg sync.WaitGroup
	wg.Add(2)
	var d data83
	go func() {
		defer wg.Done()
		d.test("read")
	}()

	go func() {
		defer wg.Done()
		d.test("write")
	}()

	wg.Wait()

}

func SynchronizeTest1() {
	var wg sync.WaitGroup
	wg.Add(2)
	var d data83_1
	// 初始化mutex
	d.Mutex = new(sync.Mutex)
	go func() {
		defer wg.Done()
		d.test("read")
	}()

	go func() {
		defer wg.Done()
		d.test("write")
	}()

	wg.Wait()

}

/*
	Mutex 不支持递归锁，即便在同一goroutine下也会导致死锁
	递归锁(Recursive Lock)也称为可重入互斥锁(reentrant mutex),是互斥锁的一种,同一线程对其多次加锁不会产生死锁。
	下面的代码就会引发死锁
	fatal error: all goroutines are asleep - deadlock!
*/
func MutexTest() {
	var mutex sync.Mutex
	mutex.Lock()
	{
		mutex.Lock()
		mutex.Unlock()
	}
	mutex.Unlock()
}

/*
	Mutex 锁粒度应该控制在最小范围内
*/
var c83 cache83

// 正确的用法
func MutexTestRight() {
	var mutex sync.Mutex
	mutex.Lock()
	url := c83.data[0]
	mutex.Unlock()
	http.Get(url)
}

// 错误的用法
func MutexTestFalse() {
	var mutex sync.Mutex
	mutex.Lock()
	url := c83.data[0]
	http.Get(url) // 该操作不属于需要加锁的最小粒度，因此这里是不合适的
	mutex.Unlock()
}

/*
	容易忽视的递归锁
	下面的代码中，get() 调用了 count() 就形成了递归锁
*/
type cache83 struct {
	sync.Mutex
	data []string
}

func (c *cache83) count() int {
	c.Lock()
	n := len(c.data)
	c.Unlock()
	return n
}

func (c *cache83) get() string {
	c.Lock()
	defer c.Unlock()
	var d string
	if n := c.count(); n > 0 {
		d = c.data[0]
		c.data = c.data[1:]
	}
	return d
}

func MutexRecursiveTest1() {
	c := cache83{
		data: []string{"1", "2", "3"},
	}
	print(c.get())
}
