package test1

import (
	"errors"
	"log"
)

// 自定义error信息
var errorDivByZero = errors.New("division by zero")

// 自定义error结构
type DivError struct {
	x, y int
}

/*
	error的字符串内容全部小写，没有结束标点，以便嵌入到其他格式化字符串中输出
	*不建议使用全局错误变量，因为可以被用户重新赋值，这就可能导致结果不匹配，
	go暂时没有只读变量功能，只能靠自觉
*/
func (DivError) Error() string {
	return "division by zero"
}

func ErrorTest() {
	d, err := DivisionErrorTest(1, 0)
	if err == errorDivByZero {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatalf("undefine error [%s]", err)
	}

	log.Printf("x / y = %f", d)
}

func DivisionErrorTest(x, y int) (float32, error) {
	if y == 0 {
		return 0, errorDivByZero
	}

	return float32(x) / float32(y), nil
}

/*
	在方法名中写入Struct不是一个好主意，但是这里为了
	学习时区分方法的用处
	在真正开发过程中不要这么做
*/
func DivisionErrorStructRunTest() {
	d, err := DivisionErrorStructTest(1, 2)
	log.Printf("type of err is: %T", err)
	// 如果err是nil则err.(type)也返回nil
	switch e := err.(type) {
	case nil:
		log.Printf("error?[%v]: x / y = %f", e, d)
	case DivError:
		log.Fatalf("%s,[%d ÷ %d]", err, e.x, e.y)
	default:
		log.Fatalln("undefine error")
	}

}

func DivisionErrorStructTest(x, y int) (float32, error) {
	if y == 0 {
		return 0, DivError{x, y}
	}

	return float32(x) / float32(y), nil
}

/*
	与error相比，panic/recover在使用方式上更接近try/catch结构化异常
	连续调用panic，只有最后一个会被recover捕获
	*建议：除非是不可恢复性错误，或是导致系统无法正常工作的错误，否则不建议使用panic
		例如：文件没有权限，服务端口被占用，数据库未启动 等等都可以使用panic
*/
func PanicTest() {
	defer func() {
		for {
			if err := recover(); err != nil {
				log.Println(err)
			} else {
				log.Fatalln("no error")
			}
		}
	}()

	defer func() {
		panic("defer dead")
	}()

	panic("main dead")
}
