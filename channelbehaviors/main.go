package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	//成本/收益
	//无缓冲 channel 提供了信号被发送就会被接收的保证，这很好，但是没有任何东西是没有代价的。
	//这个成本就是保证是未知的延迟。在等待任务场景中，员工不知道你要花费多长时间发送你的报告。
	//在等待结果场景中，你不知道员工会花费多长时间把报告发送给你。
	//在以上两个场景中，未知的延迟是我们必须面对的，因为它需要保证。没有这种保证行为，逻辑就不会起作用。
	waitForTask()
	waitForResult()

	fanOut()

}

// 场景1 - 等待任务
func waitForTask() {
	//一个带有属性的无缓冲channel被创建
	ch := make(chan string)

	go func() {
		p := <-ch

		// Employee performs work here.
		fmt.Println("get task:", p)
		// Employee is done and free to go.
	}()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	ch <- "paper"
}

//waitForResult 场景2 - 等待结果
func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

		ch <- "paper"

		// Employee is done and free to go.
	}()

	p := <-ch
	fmt.Println("get result:", p)
}

//fanOut 场景1 - 扇出（Fan Out）
func fanOut() {
	emps := 20
	ch := make(chan string, emps)

	for e := 0; e < emps; e++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "paper"
		}()
	}

	for emps > 0 {
		p := <-ch
		fmt.Println("get paper:", p)
		emps--
	}
}
