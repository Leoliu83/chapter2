package test1

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
	"unsafe"
)

/*
	ChannelTest 函数用于测试 channel 类型
	go 鼓励使用CSP通道
	通过消息来避免竞态的除了CSP 还有 Actor（akka异步就是基于Actor实现的）
	CSP:
		作为CSP的核心，通道（channel）是显式的，要求操作双方必须知道*具体通道*和*数据类型*，并不关心操作者身份和数量。
	如果另一端未准备妥当，或者消息未能及时处理，会阻塞当前端。
		通道只是一个队列，同步模式下，发送和接收双方配对，然后直接复制数据给对方。如果配对失败，则置入等待队列，直到另一方出现后才被唤醒。
		异步模式抢夺的只是缓冲槽，发送双方要求有空槽可写入，而接收方则要求有缓冲数据可读。需求不符合时，同样加入等待队列，直到另一方写入数据或腾出空槽后被唤醒。
	Actor:
		相比起来，Actor是透明的，它不在乎数据类型及通道，只要知道接受者信箱即可，默认就是异步方式，发送方对消息是否被接收或者被处理并不关心

	<-channel 通道接收数据操作只执行一次，不会反复接收，需要使用for循环多次调用 <-channel
*/
func ChannelSyncTest() {
	done := make(chan struct{})
	c := make(chan string)

	go func() {
		s := <-c // 接收消息
		log.Println("receive: <- ", s)
		close(done) // 关闭通道
	}()

	c <- "hi" // 发送消息
	<-done    // 等待消息，如无消息则阻塞
}

/*

 */
func ChannelAsyncTest() {
	c := make(chan int, 3) // 创建带有3个缓冲池的channel
	c <- 1
	c <- 2

	log.Println("i1: ", <-c) // 接收消息
	log.Println("i2: ", <-c) // 接收消息

	log.Println("Finished.")

}

/*
	缓冲区大小属于内部属性，不属于类型的组成部分。
	另外，通道变量是指针，可以用相等判断符判断是否是相同对象或nil
	len 和 cap 函数可以用来获取 channel 的当前缓冲的数量，以及最大缓冲数量（异步通道）
	对于同步通道来说，都返回0，因此可以使用cap函数来判断是否是异步channel
*/
func ChannelCompareTest() {

	var a, b chan int = make(chan int, 3), make(chan int)
	var c chan bool
	println(a == b)
	println(c == nil)

	log.Printf("%p,%d\n", a, unsafe.Sizeof(a))

	a <- 1
	a <- 2
	log.Printf("chann[a]当前缓冲区长度: %d, 最大缓冲区长度: %d\n", len(a), cap(a))
	log.Printf("chann[b]当前缓冲区长度: %d, 最大缓冲区长度: %d\n", len(b), cap(b))

}

/*
	除了使用简单的接收和发送符号，还可以使用ok-idom或range模式处理数据
	如果有多个 goroutine 使用同一个channel，则发送的数据由多个channel中的随机一个接收，
	谁来接收完全看goroutine调度
	对于循环接收数据，range 模式更简洁一些
*/
func ChannelReceiveTest() {
	done := make(chan struct{})
	c := make(chan int)
	go func() {
		defer close(done)
		for { // for 无限循环，反复接收数据
			x, ok := <-c
			if !ok { // 如果 通道被 close() 则ok值为false, 接收的x为0，因此不能仅仅靠接收值来判断是否关闭
				println("for -> !ok:", x)
				return
			}
			println("for -> ok:", x)
		}
	}()

	go func() { // range 的语法更简洁
		for x := range c {
			println("range -> x:", x)
		}
	}()

	c <- 0
	c <- 1
	c <- 2
	c <- 3
	close(c) // 及时使用close引发通道关闭，否则可能会导致死锁：fatal error: all goroutines are asleep - deadlock!

	<-done
}

/*
	多goroutine通知
*/
func ChannelMultiReceiveTest() {
	var wg sync.WaitGroup
	var cinfo = make(chan string)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			println(id, "wait info")
			info := <-cinfo
			println(id, "get info", info, "running")
		}(i)
	}
	time.Sleep(time.Second)
	println("main send info")
	// cinfo <- "go" // 发送信号只能有其中一个goroutine接收
	close(cinfo) // 只能使用关闭来作为信号，关闭可以让所有gorouting接收到关闭消息
	wg.Wait()
}

/*
	通过 sync.Cond 实现单播和组播
*/
func ChannelMultiNoticeTest() {
	var wg sync.WaitGroup
	var locker sync.Mutex
	// 这里必须使用指针传递
	var cond *sync.Cond = sync.NewCond(&locker)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			cond.L.Lock()
			log.Println(id, "Now i am waiting.")
			cond.Wait()
			log.Println(id, "I receive the msg.")
			cond.L.Unlock()
		}(i)
	}
	time.Sleep(2 * time.Second) // 等待所有goroutine启动
	// 唤醒一个goroutine
	println("Call one")
	cond.Signal()
	time.Sleep(2 * time.Second) // 等待一个被唤醒
	// 唤醒一个goroutine
	println("Call one")
	cond.Signal()
	time.Sleep(2 * time.Second) // 等待一个被唤醒
	// 唤醒所有的goroutime
	println("Call all")
	cond.Broadcast()

	// 等待所有的wg done
	wg.Wait()
}

/*
	对于closed 或者 nil 通道，发送和接收操作都有相应的规则
		·向已关闭的通道发送数据，会引发panic
		·从已关闭的通道接收数据，返回已缓冲的数据或者0
		·无论收发，nil通道都会阻塞
	主线程如果发生永久阻塞，则会产生 deadlock错误
*/
func ChannelClosedAndNilTest() {
	var nilc chan int
	log.Println(nilc)
	c := make(chan int, 3)
	c <- 1
	c <- 2
	close(c)
	/*
		每次获取数据按缓冲区数据的先后顺序
		缓冲区无数据，则ok返回false，value为0
		超过缓冲区大小的，ok返回false，value为0
	*/
	for i := 0; i < cap(c)+1; i++ {
		v, ok := <-c
		println("ok?", ok, ", value: ", v)
	}
	go func() {
		println("goroutine ready to run.")
		nilc <- 1 // 这里会阻塞，因此永远不会执行下一个println
		println("goroutine run.")
	}()

	go func() {
		println("goroutine ready to run.")
		v, ok := <-nilc // 这里同样会阻塞，因此永远不会执行下一个println
		println("goroutine run,", ok, v)
	}()

	// 休眠
	time.Sleep(5 * time.Second)
}

/*
	channel 默认是双向的，并不区分发送和接收，有时候，我们可以限制收发操作的方向来获得更严谨的操作逻辑
	尽管可以使用make创建单向通道，但是没有意义，通常使用类型转换来获取单向通道，并分别赋予操作双方
	使用make为何没有意义（下面的例子）：
		通过make定义了两个单向通道，writer（只写）和reader（只读），但是两个通道并不是一个通道，也就是说
	reader无法获取数据（reader不能获取），writer写入通道的数据没有接收者（writer不能接收），因此没有意义。
	而通过类型转换的单向通道，由于属于同一个channel，因此sender发送的数据可以被recv接收到
*/
func ChannelSimplexTest() {
	var wg sync.WaitGroup
	wg.Add(4)
	writer := make(chan<- int)
	reader := make(<-chan int)
	c := make(chan int)
	var sender chan<- int = c
	var recv <-chan int = c

	go func() {
		defer wg.Done()
		// close(recv) // close 不能用于接收端
		for x := range recv {
			log.Println("recv: ", x)
		}
		log.Println("Recv finished.")
	}()

	go func() {
		defer wg.Done()
		defer close(sender)
		// defer close(writer)
		for i := 0; i < 10; i++ {
			sender <- i
		}
		log.Println("Sender finished.")
	}()

	// 休眠5秒后，再启动后面的goroutine，让输出更明显
	time.Sleep(5 * time.Second)
	/*
		下面两个都会引发异常：fatal error: all goroutines are asleep - deadlock!
		因为接收者（reader）永远获取不到数据，导致永久阻塞
		发送者（writer）永远没有接收者接收数据，导致永久阻塞
		因为同步channel是一对一的
	*/
	go func(r <-chan int) {
		defer wg.Done()
		v, ok := <-reader
		log.Println("reader: ", ok, v)
	}(reader)

	go func(w chan<- int) {
		defer wg.Done()
		writer <- 100
	}(writer)

	wg.Wait()

}

/*
	如果要同时处理多个channel，可以使用select语句
	select语句会随机选择一个*可用的通道*做收发操作
*/
func ChannelSelectTest() {
	// var (
	// 	name string
	// 	x    int
	// 	ok   bool
	// )

	var wg sync.WaitGroup
	wg.Add(2)

	a, b := make(chan int), make(chan int)
	go func() { // 接收端
		defer wg.Done()
		for {
			// 不知为何申明变量要写在for循环内部，根据测试结果，写在内部比写在外部每次申请内存少3次。TODO
			var (
				name string
				x    int
				ok   bool
			)
			select {
			case x, ok = <-a:
				name = "a"
				// log.Printf("a[name]: %p", &name)
			case x, ok = <-b:
				name = "b"
				// log.Printf("b[name]: %p", &name)
			case x, ok = <-b: // 即使是同一通道也会随机选择
				name = "b1"
				// log.Printf("b1[name]: %p", &name)
			}

			if !ok { // 任一通道关闭则停止接收
				return
			}
			// println(name, x)
			_, _ = name, x
		}
	}()

	go func() { // 发送端
		defer wg.Done()
		defer close(a)
		defer close(b)
		for i := 0; i < 10; i++ {
			select { // 随机选择发送 channel
			case a <- i:
			case b <- i * 10:
			}
		}
	}()

	wg.Wait()
}

/*
	如果需要等待多个channel都处理完成，可以将完成的channel设置为nil，
	这样它就会被阻塞，不再被select选中
*/
func ChannelSelectTest1() {
	var wg sync.WaitGroup
	wg.Add(3)
	a, b := make(chan int), make(chan int)
	i := 0
	go func() {
		defer wg.Done()
		for {
			i++
			select {
			case x, ok := <-a:
				if !ok { // 如果a被close，则将a置为nil，则该case将会阻塞，select不会再选择该case
					log.Println(i, "a: not ok!")
					a = nil
					break // 添加break，保证关闭后不执行后面的命令
					// return // 使用 return 可以直接跳出 select 跳出 for
				}
				log.Println(i, "a: ", ok, x)
			case x, ok := <-b: // 如果b被close，则将b置为nil，则该case将会阻塞，select不会再选择该case
				if !ok {
					log.Println(i, "b: not ok!")
					b = nil
					break // 添加break，保证关闭后不执行后面的命令
				}
				log.Println(i, "b: ", ok, x)
			}
			// 如果不做判断退出，则会陷入select阻塞
			if a == nil && b == nil { // 如果两个都关闭了，则退出
				log.Println("for end.")
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		defer close(a)
		for i := 0; i < 3; i++ {
			log.Println("Send a-> ", i)
			a <- i
			time.Sleep(time.Second)
		}
	}()

	go func() {
		defer wg.Done()
		defer close(b)
		for i := 0; i < 5; i++ {
			log.Println("Send b-> ", i)
			b <- i
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
}

/*
	由于实验版本是1.15，原书版本是1.5，故实验时存在差异，在此记录
	原文说明(1.5)：
		select中如果没有可用的通道，则会进入default，添加default分支可以防止select阻塞
		也可以在default中添加一些处理逻辑，比如增加通道
	实验结果(1.15)：
		1. default分支并不能防止select阻塞
		2. 而且即使select中得所有通道都可用，也会进入分支default
		3. return会跳出select 也会跳出for，因此select多通道时，有一个通道return，后续所有发送通道将无接收者，会产生 deadlock错误
	2021/4/23:
		select中default的作用：如果发现某个通道不可用（不可用的定义包括等待数据时产生阻塞）时，会执行default，并直接进入下一次循环
		如果没有default，会阻塞在随机选择的通道。
	增加default逻辑能够避免阻塞的情况是：
	e.g.：
		阻塞情况：
			for i:=0;i<10;i++{
				select{
				case x,ok:=<-c1
				case x,ok:=<-c2
				}
			}
			上述代码中，每一次for循环会随机选择一个可用的channel，可能是c1 也可能是c2
			但如果c1和c2均不可用时，select就会阻塞。
		非阻塞情况：
			for i:=0;i<10;i++{
				select{
				case x,ok:=<-c1
				case x,ok:=<-c2
				default:
				}
			}
			上述代码中，每一次for循环会随机选择一个可用的channel，可能是c1 也可能是c2（与阻塞情况相同）
			但如果c1和c2均不可用时，select会选择default，*然后进入下一次循环*（浪费循环次数）
		就像轮询机制
*/
func ChannelSelectTest2() {
	done := make(chan struct{})
	a := make(chan int)
	i := 0
	go func() {
		defer close(done)
		for {
			i++
			select {
			case x, ok := <-a:
				if !ok { // 如果a被close，则将a置为nil，则该case将会阻塞，select不会再选择该case
					log.Println(i, "a: not ok!")
					a = nil
					return // 使用 return 可以直接跳出 select 跳出 for
				}
				log.Println(i, "a: ", ok, x)
				// default: // 增加default分支也可以避免select阻塞
				// 	log.Println(i, "default")
				// 	time.Sleep(time.Second)
				// 	// return
			}
		}
	}()

	go func() {
		defer close(a)
		for i := 0; i < 3; i++ {
			log.Println("Send a-> ", i)
			a <- i
			time.Sleep(time.Second)
		}
	}()

	<-done
}

/*
	一般使用工厂方法模式将通道和goroutine绑定
*/
type receiver82 struct {
	sync.WaitGroup
	data chan int
}

/*
	由于receiver中有 sync.WaitGroup ，而 sync.WaitGroup 中又有 nocopy
	因此receiver只能是指针形式
*/
func newReceiver() *receiver82 {
	r := &receiver82{
		data: make(chan int),
	}
	r.Add(1)
	go func() {
		defer r.Done()
		for x := range r.data { // 接收消息，直到通道被关闭
			println("recv: ", x)
		}
	}()
	return r
}

func (r receiver82) close() {
	close(r.data)
}

func ChannelFactoryTest() {
	r := newReceiver()
	defer r.Done()
	go func() {
		for d := range r.data {
			log.Printf("Receive data value: %d", d)
		}
	}()

	r.data <- 1
	r.data <- 2

	r.close()
}

/*
	由于通道本身就是一个并发安全的队列，可以作为id generator，或者pool
*/
// 定义字节数组通道
type pool chan []byte

// 通道不会复制底层数据结构，可以使用值传递，相同的还有 map slice
func newPool(cap int) pool {
	return make(pool, cap)
}

func (p pool) get() []byte {
	var v []byte
	select {
	case v = <-p:
	case v = <-p:
		// default:
		// 	v = make([]byte, 10)
	}
	return v
}

func (p pool) put(b []byte) {
	select {
	case p <- b:
	default:
	}
}

func ChannelPoolTest() {
	var wg sync.WaitGroup
	wg.Add(3)

	p := newPool(10)
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			v := p.get()
			log.Printf("get(): %+v", v)
		}
	}()

	go func() {
		defer wg.Done()
		b := make([]byte, 0, 10)
		for i := 0; i < 10; i++ {
			// bi := byte(i)
			// b = append(b, bi)
			b = append(b, byte(i))
			p.put(b)
			log.Printf("put1(): %p -> %d -> %+v", b, i, b)
			// 这里必须要加sleep，如果不加sleep，发送数据的goroutine结束了之后，无论channel
			time.Sleep(time.Second)
		}
	}()

	go func() {
		defer wg.Done()
		b := make([]byte, 0, 10)
		for i := 0; i < 10; i++ {
			// bi := byte(i)
			// b = append(b, bi)
			b = append(b, byte(i+100))
			p.put(b)
			log.Printf("put2(): %p -> %d -> %+v", b, i, b)
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
}

/*
	channel 模拟信号量
	相当于利用了 异步channel的阻塞原理，异步channel的缓冲区填满，但数据没有被读取，后续对该channel的读写将会被阻塞
	缓冲区的数量，就是信号量的数量
*/
func ChannelSemaphoreTest() {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	sem := make(chan struct{}, 2) // 最大支持两个并发同时执行
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sem <- struct{}{}        // acquire 获取信号量（向异步channel缓冲区放入数据）
			defer func() { <-sem }() // 释放信号量 （从异步channel的缓冲区取出数据）
			time.Sleep(2 * time.Second)
			log.Println(id, time.Now())
		}(i)
	}
	wg.Wait()

}

/*
	time包实现了 tick channel 和 timeout
*/
func TimeAndTickChannelTest() {
	go func() {
		select {
		case <-time.After(5 * time.Second): // 2秒后超时
			log.Println("timeout......")
			os.Exit(0)
		}
	}()

	go func() {
		tick := time.Tick(time.Second)
		for {
			select {
			case <-tick:
				log.Println(time.Now())
			}
		}
	}()

	<-(chan struct{})(nil) // s会用nil channel阻塞
}

/*
	实现接收  INT 和 TERM 信号
	实现 atexit 函数
*/
var exits = &struct {
	sync.RWMutex
	funcs   []func()
	signals chan os.Signal
}{}

/*
	atexit函数可以登记这些函数。 exit调⽤终⽌处理函数的顺序和atexit登记的顺序相反（网上很多说造成顺序相反的原因是参数压栈造成的，
	参数的压栈是先进后出，和函数的栈帧相同），如果⼀个函数被多次登记，也会被多次调⽤。
*/
func atexit(f func()) {
	exits.Lock()                         // 获取读写锁
	defer exits.Unlock()                 // 延迟释放读写锁（defer可以保证一定执行）
	exits.funcs = append(exits.funcs, f) // 添加func
	log.Printf("%+v", exits.funcs)
}

/*
	singnal包中的部分源码部分解读
	// 信号量总共是62个，没有32和33，因此用两个32位数字来标志信号，所以在这里用了一个两个元素的数组array[0]保存1-31 array[1]保存32-62
	// 1-31号信号都是不可靠信号，不支持排队，32号开始，信号都是可靠信号，支持排队
	type handler struct {
    	mask [(numSig + 31) / 32]uint32
    }

    func (h *handler) want(sig int) bool {
        return (h.mask[sig/32]>>uint(sig&31))&1 != 0
    }

	//	假设 sig 为 3
	//	1. 3&31=3
	//	2. 1<<3 = 1000
	//	3. h.mask[0] | 1000  = 0000 0000 0000 0000 0000 0000 0000 1000
	//                                                               ↑ 0号位
    func (h *handler) set(sig int) {
        h.mask[sig/32] |= 1 << uint(sig&31)
    }

	//  &^ 是 golang 特有的清位符
    func (h *handler) clear(sig int) {
        h.mask[sig/32] &^= 1 << uint(sig&31)
    }
*/
func waitExit() {
	if exits.signals == nil {
		exits.signals = make(chan os.Signal)
		// 第一个参数为 chan<- os.Signal类型，表示是  os.Signal类型的channel
		// 表示将 syscall.SIGINT, syscall.SIGTERM 发送给 exits.signals
		signal.Notify(exits.signals, syscall.SIGINT, syscall.SIGTERM)
	}
	log.Printf("%+v", exits.signals)
	// 获取读锁, 因为for循环在读取exits.funcs中的内容
	exits.RLock()
	for _, f := range exits.funcs {
		defer f() // 即使某些函数panic，也能确保后续函数被执行
	} // 延迟调用按照FILO顺序
	// 释放读锁
	exits.RUnlock()
	switch <-exits.signals {
	case syscall.SIGINT:
		log.Printf("get SIGINT")
	case syscall.SIGTERM: // windows下无法测试该信号
		log.Printf("get SIGTERM")
	default:
		log.Printf("unsupport signal")
	}
}

/*
	获取信号测试并实现atexit方法
*/
func ChannelINTandTERMandAtexitTest() {
	atexit(func() { println("exit-1") })
	atexit(func() { println("exit-2") })
	waitExit()
}

/*
	将发往通道的数据打包，可以减少发送次数，提高性能
	通道队列用的仍然是锁同步机制
*/
/*
	单个int发送
	goos: windows
	goarch: amd64
	pkg: chapter2/test1
	cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
	BenchmarkChannelTransPerformOneIntTest-8   	*** Test killed: ran too long (11m0s).
	exit status 1
	FAIL	chapter2/test1	661.550s
	FAIL
*/
func transOneInt() {
	// log.Print("start")
	const (
		max     = 500e6
		bufsize = 100
		block   = 500
	)
	c := make(chan int, bufsize)
	done := make(chan struct{})
	go func() {
		count := 0
		for x := range c {
			count += x
		}
		// log.Printf("count = %d", count)
		close(done)
	}()

	for i := 0; i < max; i++ {
		c <- i
	}
	close(c)
	<-done
}

/*
	通过block批量发送
	Running tool: D:\MyTool\Go\bin\go.exe test -benchmem -run=^$ -bench ^(BenchmarkChannelTransPerformOneBlockTest)$ chapter2/test1

	goos: windows
	goarch: amd64
	pkg: chapter2/test1
	cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
	BenchmarkChannelTransPerformOneBlockTest-8   	     100	 243187330 ns/op	  397600 B/op	       3 allocs/op
	PASS
	ok  	chapter2/test1	25.330s
*/
func transOneBlock() {
	// log.Print("start")
	const (
		max       = 50000000
		bufsize   = 100
		blocksize = 500
	)
	count := 0
	c := make(chan [blocksize]int, bufsize)
	done := make(chan struct{})
	go func() {
		for x := range c {
			for i := 0; i < blocksize; i++ {
				count += x[i]
			}
			// log.Println("count => ", count)
			// time.Sleep(1 * time.Second)
		}
		// log.Printf("count = %d", count)
		close(done)
	}()

	var block [blocksize]int
	pos := -1
	for i := 0; i < max; i++ {
		pos = (i + 1) % blocksize
		// log.Println(pos)
		block[pos] = i
		if pos == 0 {
			// time.Sleep(time.Second)
			c <- block
			// log.Println("send block")
		}
	}
	// log.Println("send block over!")
	/*
		close(c)非常重要！
		如果不关闭会引发Deadlock，因为向c发送数据的goroutine都已经完成，但c并没有close
		那么上面的匿名goroutine将会一直等待数据（for x := range c），从而永远无法执行close(done)
		而主线程由于<-done，也将永远被阻塞，因此引发了deadlock
	*/
	close(c)
	<-done
	// log.Println("done!")
}
func ChannelPerformanceTest() {
	transOneBlock()
	transOneInt()
}

/*
	通道可能引发 goroutine leak，就是说，goroutine在处于发送或者接收阻塞状态，但一直未被唤醒。
	垃圾回收器并不收集此类资源，导致他们会长时间在等待队列里休眠，造成资源泄露
	使用 gotrace 查看详细gc信息(未测试)
	GODEBUG="gotrace=1,schedtrace=1000,scheddetail=1"
*/
func ChannelGarbageTest() {
	c := make(chan int)
	for i := 1; i < 100; i++ {
		go func() {
			<-c // 造成资源泄露
		}()
	}

	for i := 0; i < 30; i++ {
		time.Sleep(time.Second)
		runtime.GC() // 强制垃圾回收也无法回收
	}
}
