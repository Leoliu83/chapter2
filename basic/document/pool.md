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

**源码解读:**
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

```

