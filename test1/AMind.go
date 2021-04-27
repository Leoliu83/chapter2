package test1

/* 一些相关与go语言编程开发的思想放在这里 */
/*
	go只实现了部分面向对象的特征，更倾向于“组合优先于继承”的思想。
	将模块分解成互相独立的更小的单元，分别处理不同方面的需求，最后以匿名的方式组合在一起，共同实现对外接口。
	而且其简短一致的调用方式，更是隐藏了内部细节。
	组合没有依赖，不会破坏封装。而且整体和局部松耦合，和任意增加来实现扩展。各个单元负责单一任务，无关联，实现更简单。
*/
/*
	Go的接口实现机制很简洁，只要目标类型方法集内包含接口声明的全部方法，就被视为实现了接口，无需做显式声明。
	换句话说，开发时可以先实现类型，再抽象接口，这种*非侵入式设计*有很多的好处：
	举例来说，在项目初期就设计出合理的接口并不容易，而在代码重构，模块拆分时，再分离出接口用以解耦就很常见。
	另外在使用第三方库时，抽象出所需接口，即可屏蔽太多不需要关注的内容，以便日后替换。
*/
/*
For an operand x of type T, the address operation &x generates a pointer of type *T to x.
The operand must be addressable, that is, either a variable, pointer indirection,
or slice indexing operation; or a field selector of an addressable struct operand;
or an array indexing operation of an addressable array.
As an exception to the addressability requirement,
x may also be a (possibly parenthesized) composite literal.
If the evaluation of x would cause a run-time panic, then the evaluation of &x does too.
对于一个对象x, 如果它的类型为T, 那么&x则会产生一个类型为*T的指针，这个指针指向x。
上面规范中的这段话规定， x必须是可寻址的， 也就是说，它只能是以下几种方式：
	·一个变量: &x
	·可寻址struct的字段: &point.X
	·可寻址数组的索引操作: &a[0]
	·slice索引操作(不管slice是否可寻址): &s[1]
	·指针引用(pointer indirection): &*x
下列情况x是不可以寻址的，你不能使用&x取得指针：
	·字符串中的字节:
	·map对象中的元素
	·接口对象的动态值(通过type assertions获得)
	·常量（包括'命名常量'和'字面值'）
	·package 级别的函数
	·方法method (用作函数值)
	·中间值(intermediate value):
		·函数调用
		·显式类型转换
		·各种类型的操作 （除了指针引用pointer dereference操作 *x):
			·channel receive operations
			·sub-string operations
			·sub-slice operations
			·加减乘除等运算符
**需要注意的是，在go中有一个语法糖，&T{}，
实际是以下代码的缩写：
tmp := T{}; (&tmp)
但是，通过 &T{} 取址是有效的，而字面值 T{} 依旧是不可寻址的。
Tapir Games在他的文章unaddressable-values中做了很好的整理。
文章地址: https://go101.org/article/unofficial-faq.html#unaddressable-values

有几个点需要解释下：
	·常数为什么不可以寻址?： 如果可以寻址的话，我们可以通过指针修改常数的值，破坏了常数的定义。
	·map的元素为什么不可以寻址？:两个原因，如果对象不存在，则返回零值，零值是不可变对象，所以不能寻址，如果对象存在，因为Go中map实现中元素的地址是变化的，这意味着寻址的结果是无意义的。
	·为什么slice不管是否可寻址，它的元素读是可以寻址的？:因为slice底层实现了一个数组，它是可以寻址的。
	·为什么字符串中的字符/字节又不能寻址呢：因为字符串是不可变的。
规范中还有几处提到了 addressable:
	·调用一个receiver为指针类型的方法时，使用一个addressable的值将自动获取这个值的指针
	·++、--语句的操作对象必须是addressable或者是map的index操作
	·赋值语句=的左边对象必须是addressable,或者是map的index操作，或者是_
	·上条同样使用for ... range语句
*/

/*
	并行不同于并发
	并行：物理上具备处理多个任务的能力
	并发：逻辑上具备处理多个任务的能力
	协程：协程与线程不同，线程由cpu控制调度，而协程由程序控制调度，协程在单个线程上主动切换来实现多任务并发
	goroutine：go在运行时会创建多个线程来执行并发任务，而且任务可以调度到其他线程并行执行，更像是多线程和协程的结合，能最大限度提升执行效率，发挥多核处理能力
*/

/*
	锁使用建议
	·对性能要求较高时，应该避免defer Unlock()
	·读写并发时，用WRMutex性能会更好
	·对单个数据读写保护，可以使用原子操作，例如 atomic 包
	·执行严格测试，尽可能打开数据竞争检查（Golang Data Race Detector）
*/
