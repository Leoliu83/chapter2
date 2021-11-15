##### 并行与并发
并行不同于并发
1. 并行：物理上具备处理多个任务的能力
2. 并发：逻辑上具备处理多个任务的能力

- 协程
协程与线程不同，线程由cpu控制调度，而协程由程序控制调度，协程在单个线程上主动切换来实现多任务并发

- goroutine
go在运行时会创建多个线程来执行并发任务，而且任务可以调度到其他线程并行执行，更像是多线程和协程的结合，能最大限度提升执行效率，发挥多核处理能力



###### 锁使用建议
- 对性能要求较高时，应该避免defer Unlock()
- 读写并发时，用WRMutex性能会更好
- 对单个数据读写保护，可以使用原子操作，例如 atomic 包
- 执行严格测试，尽可能打开数据竞争检查（Golang Data Race Detector）


###### Hot Path
A hot path is a sequence of instructions executed very frequently.

When accessing the first field of a structure, we can directly dereference the pointer to the structure to access the first field. To access other fields, we need to provide an offset from the first value in addition to the struct pointer.

In machine code, this offset is an additional value to pass with the instruction which makes it longer. The performance impact is that the CPU must perform an addition of the offset to the struct pointer to get the address of the value to access.

Thus machine code to access the first field of a struct is more compact and faster.

Note that this assumes that the layout of the field values in memory is the same as in the struct definition.

In 'sync.Once' struct 'done uint32' indicates whether the action has been performed.
It is first in the struct because it is used in the hot path.
The hot path is inlined at every call site.
Placing done first allows more compact instructions on some architectures (amd64/386), and fewer instructions (to calculate offset) on other architectures.

##### CAS (Compare And Swap)
原子操作，在多线程编程中是一个很常见的课题，指的是一个操作或一系列操作在被CPU调度的时候不可中断。早期的软件基本都是单核单线程，每个操作都可以视为原子操作，因此不会有并发问题，但随着现在多核多线程编程的出现，线程并发成为了多线程编程中不可回避的一个课题。

从硬件层面来实现原子操作，有两种方式：
1. 总线加锁：因为CPU和其他硬件的通信都是通过总线控制的，所以可以通过在总线加LOCK#锁的方式实现原子操作，但这样会阻塞其他硬件对CPU的访问，开销比较大。
2. 缓存锁定：频繁使用的内存会被处理器放进高速缓存中，那么原子操作就可以直接在处理器的高速缓存中进行而不需要使用总线锁，主要依靠缓存一致性来保证其原子性。

Golang提供了了一套原子操作的接口，可以在sync\atomic目录下查看，在里面我们可以看到经典的CAS函数。CAS即比较及交换，以 CompareAndSwapInt64 这个函数为例：
``` go
// CompareAndSwapInt64 executes the compare-and-swap operation for an int64 value.
func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)
```
这个函数有三个参数，第一个是目标数据的地址，第二个是目标数据的旧值，第三个则是等待更新的新值。每次CAS都会用old和addr内的数据进行比较，如果数值相等，则执行操作，用new覆盖addr内的旧值，如果数据不相等，则忽略后面的操作。在高并发的情况下，单次CAS的执行成功率会降低，因此需要配合循环语句for，形成一个for+atmoc的类似自旋乐观锁的操作。
```go
for{
    i:=atomic.LoadInt64(&Int)
    success:=atomic.CompareAndSwapInt64(&Int,i,i+1)
    if success {
        break;
    }
}
```