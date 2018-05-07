package main

//InsertSort 插入排序
func InsertSort(a []int) {
	for j := 1; j < len(a); j++ {
		var key = a[j]
		var i = j - 1
		for ; i >= 0 && a[i] > key; i-- {
			//a[i]>key是按照升序排列，a[i]<key是按照降序排列
			a[i+1] = a[i]
		}
		a[i+1] = key
		for _, v := range a {
			print(v, ",")
		}
		println()
	}
}

//BubleSort1 冒泡排序1
func BubleSort1(a []int) {
	for j := len(a); j > 1; j-- {
		run := false
		for i := 0; i < j-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				run = true
			}
		}
		for _, v := range a {
			print(v, ",")
		}
		println()
		if !run { //说明已经排序好了
			break
		}
	}
}

//BubbleSort2 冒泡排序2
func BubbleSort2(p []int) {
	for i := 0; i < len(p)-1; i++ {
		/*每排序一趟，则至少有一个元素已经有序，
		用 j<len-i-1 可以缩小排序范围 */
		run := false
		for j := 0; j < len(p)-1-i; j++ {
			/*当前面的元素大于后面的元素时，交换位置*/
			if p[j] > p[j+1] {
				p[j], p[j+1] = p[j+1], p[j]
				run = true
			}
		}
		for _, v := range p {
			print(v, ",")
		}
		println()
		if !run { //说明已经排序好了
			break
		}
	}
}

//BubbleSort3 冒泡排序3
func BubbleSort3(p []int) {
	k := len(p) - 1
	for i := 0; i < k; i++ {
		/*每排序一趟，则至少有一个元素已经有序，
		用 j<len-i-1 可以缩小排序范围 */
		run := false
		m := 0
		for j := 0; j < len(p)-1-i; j++ {
			/*当前面的元素大于后面的元素时，交换位置*/
			if p[j] > p[j+1] {
				p[j], p[j+1] = p[j+1], p[j]
				run = true
				//记住最后一次交换的位置
				m = j
			}
		}
		for _, v := range p {
			print(v, ",")
		}
		println()
		if !run { //说明已经排序好了
			break
		}
		//将新的长度赋值给k
		k = m
	}
}

func main() {
	println("==========")
	a := []int{5, 2, 4, 6, 1, 3}
	testSort(a)
	// println("==========")
	// a2 := []int{1, 2, 7, 4, 6, 3, 5, 8, 9}
	// testSort(a2)
}

func testSort(a []int) {
	var temp = make([]int, len(a), cap(a))
	for _, v := range a {
		print(v, ",")
	}
	println()
	copy(temp, a)
	for _, v := range temp {
		print(v, ",")
	}
	// println("\n---------BubleSort1--------")
	// BubleSort1(temp)

	// copy(temp, a)
	// println("\n---------BubbleSort2--------")
	// BubbleSort2(temp)

	copy(temp, a)
	println("\n---------BubbleSort3--------")
	BubbleSort3(temp)

	// copy(temp, a)
	// println("\n---------InsertSort--------")
	// InsertSort(temp)
}
