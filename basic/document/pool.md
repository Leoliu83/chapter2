#### golang对象池

##### 说明
在构建大规模并发下，golang的GC会成为性能瓶颈，为了减少GC，从而提高并发性能，golang提供了sync.pool 也就是对象池。golang的sync.pool是**可伸缩**的，**线程安全**的。
这个池主要用来放一些需要共享的临时对象，避免gc

**官方说明:**
```doc
A Pool is a set of temporary objects that may be individually saved and retrieved.
池是一组临时对象，这些对象可以被单独保存和检索。

Any item stored in the Pool may be removed automatically at any time without notification. If the Pool holds the only reference when this happens, the item might be deallocated.
存储在Pool中的任何项目都可能在任何时候被自动删除，而无需通知。如果发生这种情况时，Pool持有唯一的引用，那么该项目可能会被取消分配。

A Pool is safe for use by multiple goroutines simultaneously.
一个Pool对于多个goroutine同时使用是安全的。

Pool's purpose is to cache allocated but unused items for later reuse, relieving pressure on the garbage collector. That is, it makes it easy to build efficient, thread-safe free lists. However, it is not suitable for all free lists.
Pool的目的是缓存已分配但未使用的项目，以便以后再使用，减轻垃圾收集器的压力。也就是说，它使建立高效、线程安全的空闲列表变得容易。然而，它并不适合于所有的自由列表。

An appropriate use of a Pool is to manage a group of temporary items silently shared among and potentially reused by concurrent independent clients of a package. Pool provides a way to amortize allocation overhead across many clients.
池的一个合适的用途是管理一组临时项目，这些临时项目在包的独立客户端之间默默地共享，并有可能被重复使用。Pool提供了一种在许多客户端之间分摊分配开销的方法。

An example of good use of a Pool is in the fmt package, which maintains a dynamically-sized store of temporary output buffers. The store scales under load (when many goroutines are actively printing) and shrinks when quiescent.
一个很好的使用Pool的例子是在fmt包中，它维护一个动态大小的临时输出缓冲区的存储。存储器在负载下（当许多goroutines积极打印时）会扩大，而在静止时则会缩小。

On the other hand, a free list maintained as part of a short-lived object is not a suitable use for a Pool, since the overhead does not amortize well in that scenario. It is more efficient to have such objects implement their own free list.
另一方面，作为生命周期较短的对象的一部分而维护的自由列表并不是Pool的合适用途，因为在这种情况下，开销并不能很好地摊销。让这些对象实现自己的自由列表会更有效率。

A Pool must not be copied after first use.
一个Pool在第一次使用后不能被复制。
```

##### Pool不适合哪些场景
```
sync.Pool 不适合用于像 socket 长连接或数据库连接池
因为，我们不能对 sync.Pool 中保存的元素做任何假设，以下事情是都可以发生的：
1. Pool 池里的元素随时可能释放掉，释放策略完全由 runtime 内部管理；
2. Get 获取到的元素对象可能是刚创建的，也可能是之前创建好 cache 住的。使用者无法区分；
3. Pool 池里面的元素个数你无法知道；

所以，只有的你的场景满足以上的假定，才能正确的使用 Pool 。sync.Pool 本质用途是增加“临时对象”的重用率，减少 GC 负担。所以说，像 socket 这种带状态的，长期有效的资源是不适合 Pool 的。

其实可以理解为，每次pool中的数据都会在垃圾回收时被集中回收，第一次GC会将池中的对象放入victim
第二次GC会将victim中的对象销毁，因此
```


##### 源码解读
``` go
type Pool struct {
    noCopy noCopy // noCopy 是一个空结构，用来防止值传递，它可以在我们使用 go vet 工具的时候生效；
    local     unsafe.Pointer // local 字段存储的是一个 poolLocal 数组的指针，poolLocal 数组大小是 goroutine 中 P 的数量,访问时,P 的 id 对应 poolLocal 数组下标索引，所以 Pool 的最大个数 runtime.GOMAXPROCS(x)所设置的数量。
    localSize uintptr        // local 的大小,也就是[P]poolLocal的大小
    victim     unsafe.Pointer // victim 是一个poolLocal数组的指针，每次垃圾回收的时候，Pool 会把 victim 中的对象移除，然后把 local 的数据给 victim
	victimSize uintptr        // victim 数组的大小
    // New 函数是在创建 pool 的时候设置的，当 pool 没有缓存对象的时候，会调用New方法生成一个新的对象。
    New func() interface{}
}

type poolLocal struct {
	// poolLocalInternal 在这里是一个匿名属性，poolLocal 可以直接调用 poolLocalInternal 的属性，可以说，poolLocal 其实就是 poolLocalInternal，只是多了一个pad来防止伪共享
	poolLocalInternal
	// 这个字段是为了防止“false sharing/伪共享”。
	pad [128 - unsafe.Sizeof(poolLocalInternal{})%128]byte
}

/*
    在1.13开始就没有了 Mutex，使用了 atomic cas
*/
type poolLocalInternal struct {
    // 私有缓存区
    // private 代表缓存的一个元素，只能由相应的一个 P 存取。因为一个 P 同时只能执行一个 goroutine，所以不会有并发的问题；
	private interface{}
	// 公共缓存区
    // shared则可以由任意的 P 访问，但是只有本地的 P 才能 pushHead/popHead，其它 P 可以 popTail。
	shared  []interface{}
}

/*
    pin将当前的goroutine固定在P上，禁止抢占，并返回P的poolLocal pool和P的id。
    调用者在做完与pool相关的操作后必须调用runtime_procUnpin()。

    这里procPin函数实际上就是先获取当前goroutine，然后对当前协程绑定的线程（即为m）加锁，即mp.locks++，然后返回m目前绑定的p的id。这个所谓的加锁有什么用呢？这个理就涉及到goroutine的调度了，系统线程在对协程调度的时候，有时候会抢占当前正在执行的协程的所属p，原因是不能让某个协程一直占用计算资源，那么在进行抢占的时候会判断m是否适合抢占，其中有一个条件就是判断m.locks==0，ok，看起来这个procPin的含义就是禁止当前P被抢占。相应的，procUnpin就是解锁了呗，取消禁止抢占。

    那么我们来看下，为何要对m设置禁止抢占呢？其实所谓抢占，就是把m绑定的P给剥夺了，其实我们后面获取本地的poolLocal就是根据P获取的，如果这个过程中P突然被抢走了，后面就乱套了，我们继续看是如何获取本地的poolLocal的。
    获取pool的localSize大小，这里加了一个原子操作atomic.LoadUintptr来获取，为什么呢？核心来说其实就是这个localSize有可能会存在并发读写的情况，而且我们的赋值语句并非一个原子操作，有可能会读取到中间状态的值，这是不可接受的。
    pool的poolLocal数组地址赋给了l
    然后比较pid与s的大小，如果小于s则直接indexLocal访问即可，否则直接调p.pinSlow函数。我们知道，s代表poolLocal数组大小，并且每个P拥有其中一个元素，看着代码我们知道pid的取值范围就是0~N
    我们先来看看indexLocal是如何获取本地的poolLocal的
*/
func (p *Pool) pin() (*poolLocal, int) {
	pid := runtime_procPin()
	// 在pinSlow中，我们先存储到本地，然后再存储到localSize，这里我们以相反的顺序加载。
	// 由于我们已经禁用了抢占，GC不能在这中间发生。
	// 因此在这里我们必须观察到至少和 localSize 一样大的 local。
    // 我们可以观察到一个更新/更大的local，这很好(我们必须观察其零初始化(zero-initialized-ness))。
	s := runtime_LoadAcquintptr(&p.localSize) // load-acquire
	l := p.local                              // load-consume
	if uintptr(pid) < s {
		return indexLocal(l, pid), pid
	}
	return p.pinSlow()
}

func (p *Pool) Get() interface{} {
	if race.Enabled {
		race.Disable()
	}
    // 阻止 m 被抢占，并拿出当前 P 对应的池子 (这里如果不阻止抢占的话，如果当前 g 换了一个 P，就可能会出现同一个 g 操作多个 P 对应的池子，前一半代码操作 A 池子后一半代码操作 B 池子，导致问题)
	l, pid := p.pin()
	x := l.private
	l.private = nil
	if x == nil {
		// Try to pop the head of the local shard. We prefer
		// the head over the tail for temporal locality of
		// reuse.
		x, _ = l.shared.popHead()
		if x == nil {
			x = p.getSlow(pid)
		}
	}
	runtime_procUnpin()
	if race.Enabled {
		race.Enable()
		if x != nil {
			race.Acquire(poolRaceAddr(x))
		}
	}
	if x == nil && p.New != nil {
		x = p.New()
	}
	return x
}
```
```go
//go:linkname sync_runtime_procPin sync.runtime_procPin
//go:nosplit
func sync_runtime_procPin() int {
    return procPin()
}
//go:linkname sync_runtime_procUnpin sync.runtime_procUnpin
//go:nosplit
func sync_runtime_procUnpin() {
    procUnpin()
}
//go:nosplit
func procPin() int {
    _g_ := getg()
    mp := _g_.m
    mp.locks++
    return int(mp.p.ptr().id)
}
//go:nosplit
func procUnpin() {
    _g_ := getg()
    _g_.m.locks--
}
```