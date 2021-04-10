package test1

import (
	"log"
	"sync"
	"time"
	"unsafe"
)

type User struct {
	age  int
	name string
	_    struct{}
}

// MapTest 方法用来测试map
func MapTest() {
	// defer 关键词在return前执行，执行顺序按照defer的相反顺序进行
	defer log.Println("执行完成.")
	defer func() {
		if r := recover(); r != nil {
			log.Printf("发现异常：%s \n", r)
		}
	}()

	// 这里必须使用make进行初始化，因为默认是nil，不执行初始化，会报：panic: assignment to entry in nil map
	var stuScoreMap = make(map[string]int32)
	stuScoreMap["Leo"] = 90
	stuScoreMap["Liu"] = 95
	log.Printf("%+v,%d \n", stuScoreMap, unsafe.Sizeof(stuScoreMap))
	for k, v := range stuScoreMap {
		log.Printf("%s -> %d \n", k, v)
	}
	// new 返回指针,只是分配了map类型本身所需要的内存，但不分配键值对存放的内存区域
	m := new(map[string]int)
	stuScoreMapNew := *m
	// 由于没有键值对存放的内存区域,因此在给键赋值的时候回抛出异常:panic: assignment to entry in nil map
	stuScoreMapNew["Leo"] = 100
	log.Println(stuScoreMap)

	m0 := make(map[string]int)

	// for i := 0; i < 10; i++ {
	// 	m0[string('a'+i)] = i
	// }

	// map 的读取时乱序的
	for k, v := range m0 {
		log.Println(k, ":", v)
	}

	m1 := map[int]User{
		1: {age: 19, name: "leo"},
		2: {age: 30, name: "liu"},
	}

	l := len(m1)
	log.Printf("Map [m1] 元素个数: %d", l)
	/*
		// cap 对map不适用，以下操作会产生变异错误
		c := cap(m1)
	*/

	/*
		因内存访问安全和hash算法等缘故，map属于 not addressable，因此无法直接修改value的成员
		以下操作会产生编译错误：
		m1[1].age += 10
	*/
	// 正确做法如下
	u := m1[1]
	u.age += 5
	m1[1] = u
	// 或者直接将map的值设置为指针类型
	m2 := map[int]*User{
		1: {age: 19, name: "leo"},
		2: {age: 30, name: "liu"},
	}
	m2[1].age += 10
	// 对于nil值可以读取，不可以写入
	// m2[3]={}
	var m3 map[int]int
	log.Printf("Map [m3] 值是 nil : %+v", m3)
	// 下面操作会产生错误：panic: assignment to entry in nil ma
	// m3[1] = 1

	// 在迭代期间对map进行删除、新增操作是安全的
	for k, v := range m1 {
		if k == 1 {
			m1[3] = User{age: 50, name: "nobody"}
		}
		if k == 2 {
			delete(m1, k)
		}
		log.Println(k, v)
	}

}

/*
	map的并发测试，map不可以并发进行读写操作，
	某个任务正在对map进行*写*操作，那么其他任务就不可以对该map进行*读写*操作
	跑出异常：fatal error: concurrent map read and map write
*/
func ConcurrentTest() {
	m := make(map[string]int)
	// 并行执行匿名函数，不断的向map写入
	go func() {
		for {
			m["a"] += 1 // 写
			log.Println("Write: ", m["a"])
			time.Sleep(time.Millisecond)
		}
	}()

	// 并行执行匿名函数，不断的从map读取
	go func() {
		for {
			_ = m["b"] // 读
			log.Println("Read: ", m["a"])
			time.Sleep(time.Millisecond)
		}
	}()

	// 阻止进程退出
	select {}
}

/*
	利用sync.RWMutex实现同步，
	保证在同一时间内不会右多个任务同时对map进行读写操作
*/
func ConcurrentMutexTest() {
	// 使用读写锁以获得最佳性能
	var lock sync.RWMutex
	// 定义个map
	m := make(map[string]int)

	// 写
	go func() {
		for {
			lock.Lock()
			m["a"] += 1
			log.Println("Write: ", m["a"])
			lock.Unlock()

			time.Sleep(time.Second)
		}
	}()

	// 读
	go func() {
		for {
			lock.Lock()
			log.Println("Read: ", m["a"])
			lock.Unlock()

			time.Sleep(time.Second)
		}
	}()

	// 阻止进程退出
	select {}
}

/*
	性能比对，预留空间和不预留空间的区别
*/
// map 不预先准备足够的空间情况
func Performance1Test() {
	m1 := make(map[int]int)
	for i := 0; i < 1000; i++ {
		m1[i] = i
	}
}

// map 预先准备1000容量情况
func Performance2Test() {
	m2 := make(map[int]int, 1000)
	for i := 0; i < 1000; i++ {
		m2[i] = i
	}
}
