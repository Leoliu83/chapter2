#### 栈的链式存储结构
栈的链式存储结构，简称**链栈**
由于栈顶就在链表的头部，因此链栈不需要头结点
空栈其实就是 top=nil
1. 进栈操作
``` go
// 将新元素的next指向原来的top元素
newElement.next = stack.top
// 将新元素赋值给top，也就是新的top
stack.top = newElement
```
2. 出栈操作
``` go
// 将next赋值给top
oldTop := stack.top
stack.top = stack.top.next
```