package main

import "fmt"

type Showman interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() Showman {
	var stu *Student
	return stu
}

/*
1.普通的 struct（非指针类型）的对象不能赋值为 nil，也不能和 nil 进行判等（==），即如下代码，不能判断 *s == nil（编译错误），也不能写：var s Student = nil。
 */

func testNil() {
	man := live()
	if man == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}

func main() {
	testNil()
}
