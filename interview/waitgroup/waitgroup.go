package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)
/*
https://studygolang.com/articles/10065
https://studygolang.com/articles/4738
 */
func main() {
	fmt.Println("------testWait1---------")
	testWait1()
	fmt.Println("------testWait2---------")
	testWait2()
}

func testWait1() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		//golang闭包的坑：循环中调用一个闭包的时候，闭包使用了外面的变量，多数情况下gotoutine只会处理for循环的最后一个值。
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	//time.Sleep(time.Second * 2)
	wg.Wait()
}

func testWait2() {
	runtime.GOMAXPROCS(2)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		//
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	time.Sleep(time.Second * 2)
	wg.Wait()
}

// go run -race 发现有两个数据冲突 data race

