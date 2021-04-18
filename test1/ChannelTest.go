package test1

import (
	"log"
	"sync"
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
	缓冲区大小属于内部属性，不属于类型的组成部分。另外，通道变量是指针，可以用相等判断符判断是否是相同对象或nil
	len 和 cap 函数可以用来获取 channel 的当前缓冲的数量，以及最大缓冲数量（异步通道）
	对于同步通道来说，都返回0
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
	如果有多个 goroutine 使用同一个channel，则发送的数据由多个channel中的随机一个接收，谁来接收完全看goroutine调度
*/
func ChannelReceiveTest() {
	done := make(chan struct{})
	c := make(chan int)
	go func() {
		defer close(done)
		for { // for 无限循环，反复接收数据
			x, ok := <-c
			if !ok { // 如果 close() 则ok值为false, 接收的x为0
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
	TODO
*/
func ChannelMultiNotice() {
	var locker sync.Mutex
	// 这里必须使用指针传递
	var cond *sync.Cond = sync.NewCond(&locker)
	for i := 0; i < 10; i++ {
		go func(id int) {
			cond.L.Lock()
			cond.Wait()
			log.Println(id, "I receive the msg.")
			cond.L.Unlock()
		}(i)
	}
}
