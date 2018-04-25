package main

// https://www.zhihu.com/question/61200249
import "fmt"

func main() {
	// github.com/jmcvetta/randutil
	sum := useChan()
	fmt.Println(sum)
}

func useChan() int {
	ch := make(chan int, 1000)
	for count := 1; count <= 1000; count++ {
		go func(n int) {
			ch <- n * n
		}(count)
	}
	var sum int
	for i := 0; i < 1000; i++ {
		sum += <-ch
	}
	return sum
}
