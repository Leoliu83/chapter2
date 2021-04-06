package test1

import (
	"fmt"
	"unsafe"
)

// MapTest 方法用来测试map
func MapTest() {
	// defer 关键词在return前执行，执行顺序按照defer的相反顺序进行
	defer fmt.Println("执行完成.")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("发现异常：%s \n", r)
		}
	}()

	var stuScoreMap map[string]int32
	// 这里必须使用make进行初始化，因为默认是nil，不执行初始化，会报：panic: assignment to entry in nil map
	stuScoreMap = make(map[string]int32)
	stuScoreMap["Leo"] = 90
	stuScoreMap["Liu"] = 95
	fmt.Printf("%+v,%d \n", stuScoreMap, unsafe.Sizeof(stuScoreMap))
	for k, v := range stuScoreMap {
		fmt.Printf("%s -> %d \n", k, v)
	}
	// new 返回指针,只是分配了map类型本身所需要的内存，但不分配键值对存放的内存区域
	m := new(map[string]int32)
	stuScoreMapNew := *m
	// 由于没有键值对存放的内存区域,因此在给键赋值的时候回抛出异常:panic: assignment to entry in nil map
	stuScoreMapNew["Leo"] = 100
	fmt.Println(stuScoreMap)

}
