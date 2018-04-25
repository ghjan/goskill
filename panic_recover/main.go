// go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理

// 在一个主进程，多个go程处理逻辑的结构中，这个很重要，如果不用recover捕获panic异常，会导致整个进程出错中断

// Golang错误和异常处理的正确姿势
//https://www.cnblogs.com/zhangboyu/p/7911190.html
package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

func main() {
	// recoverDemo1()
	recoverDemo2()
}

func funcA() (err error) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("panic recover! p: %v\n", p)
			str, ok := p.(string)
			if ok {
				err = errors.New(str)
			} else {
				err = errors.New("panic")
			}
			fmt.Println("-----stack begin--")
			debug.PrintStack()
			fmt.Println("-----stack end--")
		}
	}()
	return funcB()
}

func funcB() error {
	// simulation
	panic("foo")
	return errors.New("success")
}

func recoverDemo2() {
	err := funcA()
	if err == nil {
		fmt.Printf("err is nil\n")
	} else {
		fmt.Printf("\nerr is %v\n", err)
	}
}
func recoverDemo1() {
	defer func() { //必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println(err) //这里的err其实就是panic传入的内容，55
		}
		fmt.Println("d")
	}()
	f()
}

func f() {
	fmt.Println("a")
	panic(55)
	fmt.Println("b")

	fmt.Println("f")
}
