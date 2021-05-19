package test1

import (
	"errors"
	"log"
	"reflect"

	// "runtime"
	"unsafe"
)

/*
	slice并非动态数组或者数组指针。它通过内部指针引用底层数组，设定相关属性将数据读写操作限定在指定区域内
	slice本身是一个只读对象，其工作机制类似于数组指针的一种包装
	结构如下：
	type slice struct {
		array unsafe.Printer
		len int
		cap int
	}
	切片不支持比较操作，类型相同也不行，只能与nil进行比较。
	如何区分数组和切片：
	 	·数组需要固定长度，也就是说，如果在初始化的时候，确定了长度（元素个数），那么初始化的就是一个数组
	  如果在初始化的时候没有确定长度，那就是一个切片：
	  	[2]int{} // 数组
	  	[...]int{} // 数组，虽然没有显式的写元素个数，但是...表示由编译器根据元素的个数来确定数组长度，也是定长，因此也是一个数组
	  	[]int{} // 切片
	性能方面：
		·切片可以用来代替数组传参，避免复制的开销
		·切片底层数组可能分配在堆上，而且小数组在栈拷贝上消耗未必比make代价大
	注意：
	    如果初始化时写的是一个指定的值，那么数组就会被填充为该类型的初始化值
		例如：var s = make([]int, 2, 5) 就会填充2个0，append的时候会往后扩展，即从3开始扩展

*/
func SliceAppendTest() {
	// 初始化一个slice，1表示当前数组元素个数，初始化为0，3为最大容量
	var s = make([]int, 1, 3)
	log.Printf("s: %+v,%d \n", s, unsafe.Sizeof(s))
	s[0] = 1
	var s1 = append(s, 1)
	log.Printf("s: %+v,%d,cap: %d \n", s, unsafe.Sizeof(s), cap(s))
	log.Printf("s1: %+v,%d,cap: %d \n", s1, unsafe.Sizeof(s), cap(s1))
	// s1 append之后，slice长度变为2 没有超过3，因此s1底层的数组元素和s的数组元素是共享的，因此s1[0]改变，s[0]也会改变
	s1[0] = 2
	log.Printf("s: %+v,%d,cap: %d \n", s, unsafe.Sizeof(s), cap(s))
	log.Printf("s1: %+v,%d,cap: %d \n", s1, unsafe.Sizeof(s), cap(s1))
	// 都以s为append的对象的时候，s3会覆盖s2所扩展的元素，所以s2和s3打印元素都为 3
	var s2 = append(s, 2)
	var s3 = append(s, 3)
	log.Printf("s2: %+v,%d,cap: %d \n", s2, unsafe.Sizeof(s2), cap(s2))
	log.Printf("s3: %+v,%d,cap: %d \n", s3, unsafe.Sizeof(s3), cap(s3))
	// s4 基于s3进行扩容 所以容量扩到3，s5基于s4扩容，容量扩到4
	var s4 = append(s3, 4)
	var s5 = append(s4, 5)
	log.Printf("s4: %+v,%d,cap: %d \n", s4, unsafe.Sizeof(s4), cap(s4))
	log.Printf("s5: %+v,%d,cap: %d \n", s5, unsafe.Sizeof(s5), cap(s5))
	/*
		由于初始化的时候，总容量初始化为3，因此当s4扩容完成后，刚好达到总容量3，
		所以此时，s4和s3、s2、s1、s是共享内存的，所以修改了s4[0]也就是第一个元素，
		s3、s2、s1、s的第一个元素也会同时改变
		但是当s5扩展时，容量超过了3，变为了4，因此s5扩展指向了新的切片，也就是新的内存空间，
		因此修改s5[0]不会影响s3、s2、s1、s、s4
	*/
	s4[0] = 444
	s5[0] = 555
	log.Printf(" s: %+v,%d,cap: %d \n", s, unsafe.Sizeof(s), cap(s))
	log.Printf("s1: %+v,%d,cap: %d \n", s1, unsafe.Sizeof(s1), cap(s1))
	log.Printf("s2: %+v,%d,cap: %d \n", s2, unsafe.Sizeof(s2), cap(s2))
	log.Printf("s3: %+v,%d,cap: %d \n", s3, unsafe.Sizeof(s3), cap(s3))
	log.Printf("s4: %+v,%d,cap: %d \n", s4, unsafe.Sizeof(s4), cap(s4))
	log.Printf("s5: %+v,%d,cap: %d \n", s5, unsafe.Sizeof(s5), cap(s5))
	// s4仍然和s3,s2,s1,s公用数组，且与s5互不相关
	s4[0] = 666
	log.Printf(" s: %+v,%d,cap: %d \n", s, unsafe.Sizeof(s), cap(s))
	log.Printf("s1: %+v,%d,cap: %d \n", s1, unsafe.Sizeof(s1), cap(s1))
	log.Printf("s2: %+v,%d,cap: %d \n", s2, unsafe.Sizeof(s2), cap(s2))
	log.Printf("s3: %+v,%d,cap: %d \n", s3, unsafe.Sizeof(s3), cap(s3))
	log.Printf("s4: %+v,%d,cap: %d \n", s4, unsafe.Sizeof(s4), cap(s4))
	log.Printf("s5: %+v,%d,cap: %d \n", s5, unsafe.Sizeof(s5), cap(s5))
}

/*
	可以基于数组或者数组指针创建切片，以开始和结束索引位置确定所引用的数组片段。
	不支持反向索引，实际范围是一个右半开区间（左闭右开，包左不包右）例如 a[2:5] 表示 下标[2,5)

	生成规则，以 a[2:5:3] 为例，切片指针指向a[2]
		     low         high    max
              ↓           ↓       ↓
	+---+---+---+---+---+---+---+---+---+---+----+            +-----+---+---+---+
	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 |            | ptr | 2 | 3 | 4 |
	+---+---+---+---+---+---+---+---+---+---+----+            +-----+---+---+---+
	          ↑<- len:3 ->↑       ↑                              |
	          |<--     cap:5   -->|                              |
	          |------------------------------- array point ------|

*/
func SliceCreateTest() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Printf("数组变量[a]地址为: %p, 值为: %d \n", &a, a)
	s0 := a[:]
	log.Printf("切片变量[s0](a[:])地址为: %p, 值为: %d, 长度为: %d, 容量为: %d \n", &s0, s0, len(s0), cap(s0))
	s1 := a[2:5]
	log.Printf("切片变量[s1](a[2:5])地址为: %p, 值为: %d, 长度为: %d, 容量为: %d \n", &s1, s1, len(s1), cap(s1))
	// low: 2, high: 5,max: 7 因此 len = high-low, cap = max-low
	// pointer -> a[2]
	s2 := a[2:5:7]
	log.Printf("切片变量[s2](a[2:5:7])地址为: %p, 值为: %d, 长度为: %d, 容量为: %d \n", &s2, s2, len(s2), cap(s2))
	s3 := a[4:]
	log.Printf("切片变量[s3](a[4:])地址为: %p, 值为: %d, 长度为: %d, 容量为: %d \n", &s3, s3, len(s3), cap(s3))
	s4 := a[:4]
	log.Printf("切片变量[s4](a[:4])地址为: %p, 值为: %d, 长度为: %d, 容量为: %d \n", &s4, s4, len(s4), cap(s4))
	s5 := a[:4:6]
	log.Printf("切片变量[s5](a[:4:6])地址为: %p, 值为: %d, 长度为: %d, 容量为: %d \n", &s5, s5, len(s5), cap(s5))

	/*
		下面两种方法初始化不同在于：
		s6 *没有*为 slice内部指针赋值，但是已经分配了所需的内存
		s7 *已经*为slice内部的指针赋值，指针指向 runtime.zerobase，并且分配了所需的内存
	*/
	var s6 []int
	s7 := []int{}
	// log.Printf("runtime.zerobase address: %p", runtime.GetZeroBasePtr())
	log.Printf("s6 address: %p \n", s6)
	log.Printf("s6 slice header: %#v, size: %d \n", (*reflect.SliceHeader)(unsafe.Pointer(&s6)), unsafe.Sizeof(s6))
	log.Printf("s7 address: %p \n", s7)
	log.Printf("s7 slice header: %#v, size: %d  \n", (*reflect.SliceHeader)(unsafe.Pointer(&s7)), unsafe.Sizeof(s6))

}

/*
	数组与切片比较
	可以看到，（数组）&s8 头地址就是第一个元素的地址，而（切片）&s9 头地址不是第一个元素的地址
*/
func DiffBetweenArrayAndSlice() {
	/************ 数组 *************/
	s8 := [...]int{0, 1, 2, 3, 4, 5}
	// 获取header地址
	p8 := &s8
	// 获取元素地址
	p80 := &s8[0]
	p81 := &s8[1]
	log.Printf("数组[s8]头地址为: %p, 头2个元素地址为: [%p %p]", p8, p80, p81)
	// 元素操作
	(*p8)[0] += 100
	*p81 += 100
	log.Printf("数组[s8]值: %+v", p8)

	/************ 切片 *************/
	s9 := []int{0, 1, 2, 3, 4, 5}
	// 获取header地址
	p9 := &s9
	// 获取元素地址
	p90 := &s9[0]
	p91 := &s9[1]
	log.Printf("切片[s9]头地址为: %p, 头2个元素地址为: [%p %p]", p9, p90, p91)
	// 元素操作
	(*p9)[0] += 1000
	*p91 += 1000
	log.Printf("数组[s9]值: %+v", p9)
}

func CreateSliceTest() []int {
	s := make([]int, 1024)
	for i := 0; i < len(s); i++ {
		s[i] = i
	}
	return s
}

func CreateArrayTest() [1024]int {
	var a [1024]int
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	return a
}

/*
	通过原始切片创建新切片时，不受len的限制，但是不能超过原有切片的cap
	+---+---+---+---+---+---+---+---+---+---+
	| 0 | 1 | 2 | 3 | 4 | 5 | 6 |   |   |   |      s             len: 6 cap: 10
	+---+---+---+---+---+---+---+---+---+---+
                  ↑                   ↑
	              3                   8
	            +---+---+---+---+---+---+---+
	            | 3 | 4 | 5 | 6 | 0 |   |   |      s1:=s[3:8]    len: 6 cap: 10
	            +---+---+---+---+---+---+---+
                          ↑       ↑       ↑
                          2       4       6
	                    +---+---+---+---+
	                    | 5 | 6 |   |   |          s2:=s1[2:4:6] len: 6 cap: 10
	                    +---+---+---+---+
                          ↑   ↑
						  0   1
	                    +---+---+---+---+---+
	                    | 5 |   |   |   |   |      s3:=s2[:1:5]  len: 6 cap: 10 // 超过s2的cap，产生panic
	                    +---+---+---+---+---+
	下面的实验可以看出子slice的元素地址和父slice的地址完全一致，
	也就是说，子slice的数组元素只是指向了原slice的数组元素的指针，并不会产生数据的复制
*/
func ResliceTest() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	printElementsAddress("s", &s)
	s1 := s[3:8]
	printElementsAddress("s1", &s1)
	s2 := s1[2:4:6]
	printElementsAddress("s2", &s2)
	// s3 := s2[:1:5] // slice bounds out of range [::5] with capacity 4
	log.Println(s)
	log.Println(s1)
	log.Println(s2)
	// log.Println(s3)
}

/*
	打印所有元素的地址
	@param name 参数名
	@param slice 切片指针
*/
func printElementsAddress(name string, slice *[]int) {
	log.Printf("'%s' [ ", name)
	for i := 0; i < len(*slice); i++ {
		log.Printf("slice[%d]: %p ", i, &(*slice)[i])
	}
	log.Println(" ]")
}

/*********** 利用切片生产栈式数据结构 ************/
func StackTest() {
	stack := make([]int, 0, 5)
	// 入栈
	push := func(x int) (idx int, err error) {
		err = nil
		l := len(stack)
		c := cap(stack)
		if l == c {
			idx = -1
			err = errors.New("stack is full")
			return idx, err
		}
		idx = l
		stack = stack[:l+1]
		stack[idx] = x
		return
	}
	// 出栈
	pop := func() (e interface{}, err error) {
		err = nil
		l := len(stack)
		if l == 0 {
			err = errors.New("stack is empty")
			return nil, err
		}
		idx := l - 1
		e = stack[idx]
		stack = stack[:idx]
		return e, nil
	}

	for i := 0; i < 7; i++ {
		idx, err := push(i)
		if err != nil {
			log.Println("[push]: ", err)
		}
		log.Print("[push]: ", idx)
		log.Println("[push][stack]: ", stack)
	}

	for i := 7; i > 0; i-- {
		e, err := pop()
		if err != nil {
			log.Println("[pop]: ", err)
		}
		log.Println("[pop]: ", e)
		log.Println("[pop][stack]: ", stack)
	}
}

/*
	copy函数可以在两个切片对象之间进行复制，允许指向同一底层数组，允许目标区间重叠。
	最终所复制长度以*较短的*切片长度为准
	copy(dst,src)
*/
func SliceCopyTest() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	log.Println("数组变量[s]值为: ", s)
	s1 := s[5:8]
	log.Println("数组变量[s1]值为: ", s1)
	// 将s1的所有元素复制到s，从s的第5个元素覆盖
	n := copy(s[4:], s1)
	// 打印可以看到，原始切片s，子切片s1都收到了影响，进一步说明，子切片与父切片的底层数组保持一致
	log.Println("拷贝元素数量 s1 -> s[4:] 为: ", n)
	log.Println("数组变量[s]值为: ", s)
	log.Println("数组变量[s1]值为: ", s1)
	s2 := make([]int, 6)
	log.Println("数组变量[s2]值为: ", s2)
	n = copy(s2, s)
	log.Println("拷贝元素数量 s -> s2 为: ", n)
	log.Println("数组变量[s]值为: ", s)
	log.Println("数组变量[s2]值为: ", s2)
}
