package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	test1()
	test2()
	test3()
	test4()
}

/*有数据信号 - 保证 - 无缓冲 Channels
成本/收益
无缓冲 channel 提供了信号被发送就会被接收的保证，这很好，但是没有任何东西是没有代价的。
这个成本就是保证是未知的延迟。在等待任务场景中，员工不知道你要花费多长时间发送你的报告。
在等待结果场景中，你不知道员工会花费多长时间把报告发送给你。
在以上两个场景中，未知的延迟是我们必须面对的，因为它需要保证。
没有这种保证行为，逻辑就不会起作用。
*/
func test1() {
	fmt.Println("========有数据信号 - 保证 - 无缓冲 Channels===========")
	fmt.Println("-------waitForTask----")
	waitForTask()
	fmt.Println("-------waitForResult----")
	waitForResult()
}

/*  有数据信号 - 无保证 - 缓冲 Channels > 1
成本/收益
有缓冲的 channel 缓冲大于1提供无保证发送的信号被接收到。
离开保证是有好处的，在两个goroutine之间通信可以降低或者是没有延迟。
在扇出场景，这有一个有缓冲的空间用于存放员工将被发送的报告。
在Drop场景，缓冲是测量能力的，如果容量满，工作被丢弃以便工作继续。

在两个选择中，这种缺乏保证是我们必须面对的，因为延迟降低非常重要。
0到最小延迟的要求不会给系统的整体逻辑造成问题。
*/
func test2() {
	fmt.Println("========有数据信号 - 无保证 - 缓冲 Channels > 1===========")
	fmt.Println("-------fanOut----")
	fanOut()
	fmt.Println("-------selectDrop----")
	selectDrop()
}

/*有数据信号 - 延迟保证- 缓冲1的channel
 */
func test3() {
	fmt.Println("========有数据信号 - 延迟保证- 缓冲1的channel===========")
	fmt.Println("-------waitForTasks----")
	waitForTasks()
}

/*无数据信号 - Context
 */
func test4() {
	fmt.Println("========无数据信号 - Context===========")
	fmt.Println("-------withTimeout----")
	withTimeout()

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

// 场景2 - Drop
//假设你是经理，你雇佣了单个员工来完成工作。你有一个单独的任务想员工去执行。
//当员工完成他们任务时，你不在乎知道他们已经完成了。
//最重要的是你能或不能把新工作放入盒子。
//如果你不能执行发送，这时你知道你的盒子满了并且员工是满负荷的。
//这时候，新工作需要丢弃以便让事情继续进行。
func selectDrop() {
	const cap = 5
	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("employee : received :", p)
		}
	}()

	const work = 20
	for w := 0; w < work; w++ {
		select {
		case ch <- "paper":
			fmt.Println("manager : send ack")
		default:
			fmt.Println("manager : drop")
		}
	}

	close(ch)
}

//有数据信号 - 延迟保证- 缓冲1的channel
//场景1 - 等待任务
func waitForTasks() {
	ch := make(chan string, 1)

	go func() {
		for p := range ch {
			fmt.Println("employee : working :", p)
		}
	}()

	const work = 10
	for w := 0; w < work; w++ {
		ch <- "paper"
	}

	close(ch)
}

//无数据信号 - Context
/*
你是经理，你雇佣了一个单独的员工来完成工作，
这次你不会等待员工未知的时间完成他的工作。
你分配了一个截止时间，如果你的员工没有按时完成工作，你将不会等待

特别注意ch使用了一个缓冲的chan，
如果你使用一个无缓冲channels，如果你超时以后离开，员工将一直阻塞在那尝试你给发送报告。
这会引起goroutine泄漏。因此一个缓冲的channels用来防止这个问题发生。
 */
func withTimeout() {
	duration := 50 * time.Millisecond

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		ch <- "paper"
	}()

	select {
	case p := <-ch:
		fmt.Println("work complete", p)

	case <-ctx.Done():
		fmt.Println("moving on")
	}
}
