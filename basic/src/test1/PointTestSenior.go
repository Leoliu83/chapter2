package test1

import (
	"log"
	// "sync"
)

type VertexType = string
type EdgeType = int

func init() {

}

type MGraph struct {
	vexs                  []VertexType
	arc                   []EdgeType
	numVertexes, numEdges int
	// lock                  sync.RWMutex
}

/*
	创建图，该函数返回两个值，用于测试
	1. 指针类型返回的是：存放结构体所在的真实内存地址的变量，也就是指针
	2. 变量类型返回的是：将原指针所指向的结构体的数据拷贝到新的内存地址后，新地址的别名，也就是普通变量
	@Return *MGraph Mgraph类型指针
	@Return Mgraph Mgraph类型变量
*/
func CreateMGraph(v []VertexType, vr []EdgeType, es int) (*MGraph, MGraph) {
	if v == nil || vr == nil {
		log.Fatalf("Error!")
	}
	vlen := len(v)
	vrlen := len(vr)
	if vlen > es || vrlen > es {
		log.Fatalf("Expect len: %d, actural: v -> %d,vr -> %d", es, len(v), len(vr))
	}
	mGraph := new(MGraph)
	mGraph.arc = make([]EdgeType, 0, es)
	mGraph.arc = append(mGraph.arc, vr...)
	// for _, vr1 := range vr {
	// 	tmp := vr1
	// 	mGraph.arc = append(mGraph.arc, tmp)
	// }
	mGraph.vexs = make([]VertexType, 0, es)
	mGraph.vexs = append(mGraph.vexs, v...)
	// for _, v1 := range v { //v1 地址不会变，值会不断变，因此需要中间变量tmp
	// 	tmp := v1
	// 	mGraph.vexs = append(mGraph.vexs, tmp)
	// }
	mGraph.numEdges = vlen
	mGraph.numVertexes = vrlen
	log.Printf("===>变量[mGraph]的类型: %T \n", mGraph)
	log.Printf("===>指针变量[mGraph]的内存地址: 0x%x \n", &mGraph)
	log.Printf("===>指针变量[mGraph]的内存地址: %p \n", &mGraph)
	log.Printf("===>结构体[mGraph]的真实内存地址: %p \n", mGraph)
	log.Printf("===>变量[mGraph]的值: %+v \n", mGraph)
	log.Printf("===>变量[mGraph]的真实值: %+v \n", *mGraph)
	return mGraph, *mGraph
}

/*

 */
func AddNode(mg MGraph, node VertexType) {
	log.Printf(">>>>[AddNode1]变量[mg]的类型: %T \n", mg)
	log.Printf(">>>>[AddNode1]指针变量[mg]的内存地址: %p \n", &mg)
	// log.Printf(">>>>[AddNode1]结构体[mg]的真实内存地址: %p \n", mg) -- warning 说明mg是一个普通变量，只是一个别名
	log.Printf(">>>>[AddNode1]变量[mg]的值: %+v \n", mg)
	// log.Printf(">>>>[AddNode1]变量[mg]的真实值: %+v \n", *mg) -- error
	mg.numEdges++
	// log.Printf("[AddNode1]: %p,%+v", mg.vexs, mg.vexs)
	mg.vexs = append(mg.vexs, node)
	log.Printf(">>>>[AddNode1]: %p,%+v", mg.vexs, mg.vexs)
}

func AddNode2(mg *MGraph, node VertexType) {
	log.Printf(">>>>[AddNode2]变量[mg]的类型: %T \n", mg)
	log.Printf(">>>>[AddNode2]指针变量[mg]的内存地址: %p \n", &mg)
	log.Printf(">>>>[AddNode2]结构体[mg]的真实内存地址: %p \n", mg)
	log.Printf(">>>>[AddNode2]变量[mg]的值: %+v \n", mg)
	log.Printf(">>>>[AddNode2]变量[mg]的真实值: %+v \n", *mg)
	mg.numEdges++
	// log.Printf("[AddNode2]: %p,%+v", mg.vexs, mg.vexs)
	(*mg).vexs = append(mg.vexs, node)
	log.Printf("[AddNode2]: %p,%+v", mg.vexs, mg.vexs)
}

func PointTestSenior() {
	// algsort.BubbleSort()
	// algsort.SelectSort()
	edges := []EdgeType{1, 2}
	vexes := []string{"a", "b"}
	// mGraph1 指针
	// 值
	mGraph1, mGraph2 := CreateMGraph(vexes, edges, 10)
	log.Printf("--->变量[mGraph1]的类型: type: %T \n", mGraph1)
	log.Printf("--->指针变量[mGraph1]的内存地址: %p \n", &mGraph1)
	log.Printf("--->指针变量[mGraph1]的真实内存地址: %p \n", mGraph1)
	log.Printf("--->变量[mGraph1]的值: %+v \n", mGraph1)
	log.Printf("--->变量[mGraph1]的真实值: %+v \n", *mGraph1)
	AddNode(*mGraph1, "1")
	log.Printf("--->(AddNode()执行后)变量[mGraph1]的类型: type: %T \n", mGraph1)
	log.Printf("--->(AddNode()执行后)指针变量[mGraph1]的内存地址: %p \n", &mGraph1)
	log.Printf("--->(AddNode()执行后)指针变量[mGraph1]的真实内存地址: %p \n", mGraph1)
	log.Printf("--->(AddNode()执行后)变量[mGraph1]的值: %+v \n", mGraph1)
	log.Printf("--->(AddNode()执行后)变量[mGraph1]的真实值: %+v \n", *mGraph1)
	AddNode2(mGraph1, "2")
	log.Printf("--->(AddNode2()执行后)变量[mGraph1]的类型: type: %T \n", mGraph1)
	log.Printf("--->(AddNode2()执行后)指针变量[mGraph1]的内存地址: %p \n", &mGraph1)
	log.Printf("--->(AddNode2()执行后)指针变量[mGraph1]的真实内存地址: %p \n", mGraph1)
	log.Printf("--->(AddNode2()执行后)变量[mGraph1]的值: %+v \n", mGraph1)
	log.Printf("--->(AddNode2()执行后)变量[mGraph1]的真实值: %+v \n", *mGraph1)
	log.Println("======================================================================================")
	log.Println("======================================================================================")
	log.Printf("--->变量[mGraph2]的类型: type: %T \n", mGraph2)
	log.Printf("--->指针变量[mGraph2]的内存地址: %p \n", &mGraph2)
	// log.Printf("--->指针变量[mGraph2]的真实内存地址: %p \n", mGraph2) // warning 说明mg是不可寻址的，只是一个别名
	log.Printf("--->变量[mGraph2]的值: %+v \n", mGraph2)
	// log.Printf("--->变量[mGraph2]的真实值: %+v \n", *mGraph2) // error
	AddNode(mGraph2, "3")
	log.Printf("--->(AddNode()执行后)变量[mGraph2]的类型: type: %T \n", mGraph2)
	log.Printf("--->(AddNode()执行后)指针变量[mGraph2]的内存地址: %p \n", &mGraph2)
	// log.Printf("--->(AddNode()执行后)指针变量[mGraph2]的真实内存地址: %p \n", mGraph2) // warning 说明mg是不可寻址的，只是一个别名
	log.Printf("--->(AddNode()执行后)变量[mGraph2]的值: %+v \n", mGraph2)
	// log.Printf("--->(AddNode()执行后)变量[mGraph2]的真实值: %+v \n", *mGraph2) // error
	AddNode2(&mGraph2, "4")
	log.Printf("--->(AddNode2()执行后)变量[mGraph2]的类型: type: %T \n", mGraph2)
	log.Printf("--->(AddNode2()执行后)指针变量[mGraph2]的内存地址: %p \n", &mGraph2)
	log.Printf("--->(AddNode2()执行后)变量[mGraph2]的值: %+v \n", mGraph2)
}
