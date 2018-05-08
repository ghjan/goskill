package main

import "fmt"

func main() {
	s := make([]int, 5)
	testSliceAppend(s)

	var s2 []int
	testSliceAppend(s2)
}

func testSliceAppend(s []int) {
	fmt.Println(s)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}
