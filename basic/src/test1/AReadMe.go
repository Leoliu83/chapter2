package test1

/*
	go中的默认引用类型只有：
		·函数（func）
		·切片（slice）
		·字典（map）
		·通道（channel）
*/
/*
	如何将go文件编译成object file:
		go tool compile xxx.go
	如何查看汇编(object file)：
		go tool objdump xxx.o // go file无效
*/
/*
	go在runtime期间会为我们检测所有协程是否在等待着什么，例如接收channel数据 <- channel，
	如果发现有协程在等待，但是却没有其他协程在运行，就会抛出错误：
		fatal error: all goroutines are asleep - deadlock!
	我们也可以让一个协程进入死循环但不做发送数据，另一个协程等待接收channel数据 <- channel
	只有死循环的协程不结束，就不会出现 deadlock 错误。
	e.g.
		go func() { // 如果该协程结束，就会引发 deadlock 错误
			defer wg.Done()
			for {
				time.Sleep(1)
			}
		}()

		go func() {
			defer wg.Done()
			<-a
		}()
*/
