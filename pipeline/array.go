package pipeline

import "fmt"

func main() {

	testArray()
}
func testArray() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5] int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
	maxIndex := -1
	maxValue := -1
	for i, v := range arr3 {
		if v > maxValue {
			maxIndex = i
			maxValue = v
		}
		fmt.Println(i, v)
	}
	fmt.Println(maxIndex, maxValue)
	sum := 0
	for _, v := range arr3 {
		sum += v
	}
	fmt.Println(sum)

	fmt.Println("-----printArray(arr1)-----")
	printArray(&arr1)
	fmt.Println("------printArray(arr3)----")
	printArray(&arr3)
	//printArray(arr2)
	fmt.Println("------arr1 ----")
	fmt.Println(arr1)
	fmt.Println("------arr3 ----")
	fmt.Println(arr3)

}

func printArray(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
