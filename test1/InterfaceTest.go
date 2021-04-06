package test1

import (
	"fmt"
)

// Subscriber 是一个接口类型，定义了订阅者需要实现的方法,其中包含一个需要实现的方法“通知”
type Subscriber interface {
	notice() bool
}

// SubscriberOne 表示一个订阅者
type SubscriberOne struct{}

func (s SubscriberOne) notice() bool {
	fmt.Println("I am SubscriberOne. I receive the notice.")
	return true
}

// InterfaceTest 测试接口类型
func InterfaceTest() {
	var o SubscriberOne
	o.notice()
}
