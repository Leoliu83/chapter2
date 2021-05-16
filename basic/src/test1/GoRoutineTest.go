package test1

import (
	"log"
	"math"
	"runtime"
	"sync"
	"time"
)

/*

 */
func GoroutineTest() {
	/*
		定义一个通道（channel)结构体，借助通道接收数据的等待状态来让主进程挂起
		在子进程中 close(channel) 就可以解除挂起
		或者向channel中写入数据，也可以解除挂起
		也可以使用select{}挂起主进程
	*/
	exit := make(chan struct{})
	go println("go: Hello parallel!")

	go func(s string) {
		println(s)
	}("func: Hello parallel!")

	/*
		goroutine 和 defer一样会因为延迟执行而立即计算并复制执行参数
	*/
	a := 100
	go func(x, y int) {
		// sleep 1 秒  保证后面的任务先执行
		time.Sleep(time.Second)
		println("go: ", x, y)
		// 关闭通道（channel）,发出信号
		close(exit)
	}(a, counter())

	a += 100
	println("main: ", a, counter())

	log.Println("Main...")
	// 通道等待接收数据
	<-exit
	log.Println("Main exit!")
}

/*
	如果要等待多个任务结束，推荐使用sync.WaitGroup。通过设定计数器，让每个goroutine在结束前递减，直至归0时解除阻塞
	有点像java 的  countDownLatch
	建议将wg.Add 累加放在 goroutine外部,防止 goroutine还没启动并设置 wg.Add 而主线程的wait就已经结束了
	for i := 0; i < 10; i++ {

		go func(id int) {
			defer wg.Done() // 递减计数器
			wg.Add(1) // 累加计数器 来不及设置
			log.Printf("goroutine -> id = %d, done!", id)
		}(i)
	}

	wg.wait() <-- 比goroutine更快结束
*/
func GoroutineWaitTest() {
	var wg sync.WaitGroup
	// 循环开启goroutine
	for i := 0; i < 10; i++ {
		wg.Add(1) // 累加计数器
		go func(id int) {
			defer wg.Done() // 递减计数器
			time.Sleep(time.Second)
			log.Printf("goroutine -> id = %d, done!", id)
		}(i)
	}
	log.Println("main")
	wg.Wait() // 等待归0，解除阻塞
	log.Println("main done!")
}

func GoRoutineParallelTest(n int, isParallel bool) {
	// 获取cpu的数量
	maxProcs := runtime.NumCPU()
	if isParallel {
		if n > maxProcs { // 如果n超过cpu数量，线程数设置为cpu数量，否则设置为n
			runtime.GOMAXPROCS(maxProcs)
			// runtime.GOMAXPROCS(1) // 测试代码
		} else {
			runtime.GOMAXPROCS(n)
		}

		// 获取当前并行线程数
		currProcs := runtime.GOMAXPROCS(0)
		log.Printf("Is parallel? %t", isParallel)
		log.Printf("Max processes count is: %d", maxProcs)
		log.Printf("Current processes count is: %d", currProcs)
		log.Printf("Loop count is: %d", n)
		var wg sync.WaitGroup
		wg.Add(n)
		for i := 0; i < n; i++ {
			go func() {
				doSth81()
				wg.Done()
			}()
		}
		wg.Wait()
	} else {
		log.Printf("Is parallel? %t", isParallel)
		log.Printf("Max processes count is: %d", maxProcs)
		log.Printf("Current processes count is: %d", 1)
		log.Printf("Loop count is: %d", n)
		for i := 0; i < n; i++ {
			doSth81()
		}
	}

}

/*
	TLS(局部存储)的go功能实现
*/
func LocalStorageTest() {
	var wg sync.WaitGroup
	var tls [5]struct {
		id     int
		result int
	}

	/*
		下面这种写法是不对的：
		tls[i].id = id
		tls[i].result = (id + 1) * 100
		会产生编译告警：loop variable i captured by func literal
		因为主线程很快就循环完了，但是goroutine还没有开始，等开始的时候，i已经不是当时的值了
		所以并不会得到期望的tls[1] tls[2] tls[3]

		i 作为id传入，因为goroutine和defer会延迟调用，因此id是先复制了i的值后才注册调用的，因此id永远是循环时候的i值
	*/
	for i := 0; i < len(tls); i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			tls[id].id = id
			tls[id].result = (id + 1) * 100
		}(i)
	}
	wg.Wait()
	log.Printf("%+v", tls)
}

/*
	runtime.Gosched 可以让出当前线程的占用，让线程去执行其他任务。当前队列被放回队列，等待下次调度时恢复执行
	该函数比较少使用，因为运行时，会主动向长时间运行（10ms）的任务发出抢占调度。（未证实）
*/
func GoschedTest() {
	exit := make(chan struct{})
	runtime.GOMAXPROCS(1)
	go func() { // goroutine A
		go func() { // goroutine B （放在A内是为了保证A比B优先执行，如果放在A外，不能保证A先与B执行）
			log.Println("B")
		}()
		for i := 0; i < 10; i++ {
			if i == 3 { // 当i==3 时候，A让出线程，这时，gorouting B开始执行
				runtime.Gosched()
			}
			log.Println("A", i)
		}
		close(exit)
	}()
	<-exit
}

/*
	Goexit会立即终止当前任务，运行时*确保*所有已经注册的延迟调用被执行，
	该函数不会影响其他并发任务，不会引发panic，自然也就无法捕获
	Goexit无论在那一层都可以直接终止整个堆栈调用，并且执行所有注册的延迟调用
	os.Exit 直接终止进程，但不执行注册的任何延迟调用

*/
func GoExitTest() {
	exit := make(chan struct{})
	go func() {
		defer close(exit)
		defer println("Defer 1")
		func() {
			defer func() {
				log.Printf("level 1 panic? %t", recover() == nil)
			}()
			println("level 1 start...")
			func() {
				defer func() {
					log.Printf("level 2 panic? %t", recover() == nil)
				}()
				println("level 2 start...")
				func() {
					defer func() {
						log.Printf("level 3 panic? %t", recover() == nil)
					}()
					println("level 3 start...")
					runtime.Goexit()          // 终止整个堆栈调用
					println("level 3 end...") // 不执行
				}()
				println("level 2 end...") // 不执行
			}()
			println("level 1 end...") // 不执行
		}()
	}()
	<-exit
}

func doSth81() {
	x := 0
	for i := 0; i < math.MaxUint32; i++ {
		x += i
	}
}

var v81 int = 0

func counter() int {
	v81++
	return v81
}
