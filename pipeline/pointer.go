package pipeline

import "fmt"

func main() {
	testPointer1()
	a, b := 3, 4
	//pass by value
	swapBad(a, b)
	fmt.Println(a, b)
	//pass by reference
	swapGood(&a, &b)
	fmt.Println(a, b)

	a, b = swap(a, b)
	fmt.Println(a, b)

}

func testPointer1() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)
}

func swapBad(a, b int) {
	a, b = b, a
}

func swapGood(a, b *int) {
	*a, *b = *b, *a
}

func swap(a, b int) (int, int) {
	return b, a
}
