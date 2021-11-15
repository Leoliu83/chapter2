package leetcode

import "testing"

var root *BiThrTreeNode

/*
	初始化树，结构如下：
                   1
         ┌─────────┴─────────┐
         2                   3
    ┌────┴────┐         ┌────┴────┐
    4         5         6         7
┌───┴───┐ ┌───┴───┐ ┌───┴───┐ ┌───┴───┐
nil   nil 10     11 12     13 14     15
*/
func init() {
	// 初始化树
	root := &BiThrTreeNode{data: 1}
	CreateBiThrTreeNode(root, 1)
}

func BenchmarkPrintNonQueueRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		root.PrintNonQueueRecursion()
	}
}

func BenchmarkPrintNonQueueLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		root.PrintNonQueueLoop()
	}
}
