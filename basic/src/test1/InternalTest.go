package test1

import (
	"gostudy/basic/src/test1/innerpkg"
	// "chapter2/test1/innerpkg/internal"
)

/*
	在外部调用 chapter2/test1/innerpkg/internal 是不允许的
	test1\InternalTest.go:4:2: use of internal package chapter2/test1/innerpkg/internal not allowed
*/
// func InternalTest() {
// ib := internal.GetInternalObj()
// println(ib)
// }

func InternalTest1() {
	innerpkg.InternalTest()
}
