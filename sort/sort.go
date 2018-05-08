package main

// https://blog.csdn.net/llzk_/article/category/6489779
// https://blog.csdn.net/llzk_/article/details/51547923
// https://www.cnblogs.com/red-code/p/6645081.html

// O(log(n))的算法比较 最快-》最慢 ： 基数排序 快速排序 希尔排序 堆排序
// O(n **2)的几个算法比较， 一般来说，插入最快，选择其次，冒泡最慢。

//RadixSort 基数排序
/**
* 名称：基数排序
* 描述：在描述基数排序前先描述一下“桶排序”：
*      桶排序：
*      假设待排数组{A1，A1，A3·····}中所有的数都小于“M”，
*      则建立一个长度为M的数组count[M]，初始化为全0。
*      当读取A1时，将count[A1]增1（初始为0，现在为1），当读取A2时，将count[A2]增1····
*      之后count[M]中的每一个非0项的顺序就是排序结果。
*      基数排序：
*      对于数组中的所有项的“每一位数”都进行桶排序。
*      比如先对所有项的“个位”进行桶排序，根据个位的桶排序的结果，对各个项进行一次排序。
*      之后再对十位进行桶排序··········
* 时间复杂度（d代表长度，n代表关键字个数，r代表关键字的基数）：平均O(d(n+rd))，最坏O(d(n+rd))
* @param array 待排数组
* @param len_max 待排数组项的最高位位数，如待排数组={2,23,233,2333}，则len_max为4（2333的位数）
* 稳定性：稳定
 */
func RadixSort(array []int, lenMax int) {
	/**
	* 一般的桶排序的count[]只是一维数组，
	* 里面的每项（0-9）的具体值代表了该数出现的次数，
	* 如count[2]=3，代表2这个“项”，出现了3次。
	*
	* 但现在是“从个位开始”，每一位都要做桶排序。
	* 如果还是使用count[]，
	* 那么count[2]=3只能代表“在所有项的第n位桶排序中”2这个数“作为第n位数”出现了3次。
	*
	* 显然，正常的流程是：
	* 先查找“个位”的count[0]=n，将这些“个位为0”的数从第0位开始放入数组。
	* 再看“个位”的count[1]=m，将这些“个位为1”的数紧挨着刚刚插入的数插进来。·····
	* “个位”的第一轮排序完后，原数组相当于“依据个位大小，进行了一次排序”，
	* 接下来就要“依据十位大小再进行一次排序了”····。
	*
	* 综上流程能够发现，第n位中count[2]=3，这个2代表了“3个n位为2的数”，
	* 我们要排序就必须知道“这3个数具体是什么”，然后把这些数按序放入原数组。
	* 所以引入二维数组count[][]，
	* 第一维的下标代表了该位数“具体是几”，所以范围是0-9。
	* 第二维的下标代表了该位数相同的值“第几次出现”，
	* （如个位桶排序时，count[2][34]就代表了第34个“个位为2的数”。）
	* 第二维中存储的是具体的某个数
	 */
	var count [10][]int

	//该数组frequency[n]=m，用来计算“某位的桶排序”中“n这个数第m次出现”。
	// 所以n的范围只能是0-9，而m最多可能是原数组的长度（当该数组某一位的值都是同一个数时）
	var frequency [10]int

	nowDigit := 1 //当前排序的是各项的第几位数（从第一位（个位）开始排）
	n := 1        //用来计算当前位的具体值

	//从个位开始排，然后再排十位·····
	for nowDigit <= lenMax {
		//根据原数组中各项的“now_digit位”，进行桶排序。
		for i := 0; i < len(array); i++ {
			//找到具体某位的值。如n=1时，找到的就是个位的值。n=10时，找到的就是十位的值
			digit := ((array[i] / n) % 10)
			count[digit][frequency[digit]] = array[i]
			frequency[digit]++
		}

		/**
		* 现在所有的项已经根据桶排序规则存入count[][]中，现在需要按序再存回原数组。思路如下：
		* count[][]中第一个下标意味着“当前位”的具体值，为0-9.
		* 所以应该将count[0][n]中的各个数排在前面，count[1][m]中的各项跟在后面·····
		* count[0][]中存了多个“当前位为0的数”，
		* 而count[][]的第二个下标表示“被存储的数”是“第几个下标为0的数”。
		* 如:当前位为“个位”排序时，count[0][1]=21表示21是第一个“个位为0的数”，
		*    count[0][6]=341表示341是第6个“个位为0的数”
		 */
		//把数据存在原数组的什么位置（起始的存储位置自然是0，然后每存一个数后移一位）
		k := 0
		//从count[i=0][]开始找，然后找count[i++][]····
		for i := 0; i < 10; i++ {
			//frequency[i]代表了“i这个数第几次出现”，所以为0就表示没出现过，也就不用排了
			if frequency[i] != 0 {
				//从第0次出现开始找，每次都装入原数组
				for j := 0; j < frequency[i]; j++ {
					//j代表了“位值为i”的数是第几个，count[i][j]代表了该数
					array[k] = count[i][j]
					k++
				}
			}
			//“当前位数”为i的数已经存完，需要初始化，否则下一“位数为i”的数存时会出错。
			frequency[i] = 0
		}
		//每循一次，用来计算当前位的具体值的n做+10处理
		n *= 10
		//每循一次，当前位数+1
		nowDigit++
	}
}

//BucketSort 一个简单的桶排序
func BucketSort() {
	array_demo := []int{2, 5, 7, 3, 1, 6, 8, 4, 2, 1, 3, 7, 9, 4, 2, 5, 0}
	var count []int
	m := 0

	for i := range array_demo {
		count[i]++
	}

	for i := 0; i < 10; i++ {
		if count[i] != 0 {
			for k := 0; k < count[i]; k++ {
				array_demo[m] = i
				m++
			}
		}
	}
}

//QuickSort 交换排序-快速排序
/*
 * 名称：交换排序-快速排序
 * 描述：选第一个数作为“枢轴”，
 *      将枢轴与序列另一端的数比较，枢轴大于它，就换位，小于就再和另一端的倒数第二个数比较··
 *      第一次换位完了后依旧和另一边的比，但判断标准得颠倒，变成“如果枢轴小于它，就换位”
 *      一轮比完了，枢轴就到了中间，左边比它小，右边比它大。
 *      之后枢轴两边的序列继续进行快排。
 * 时间复杂度：平均O(nlogn)，最坏O(n^2)
 * 稳定性：不稳定
 * @param array 待排数组
 * @param low 开始位置（初始为0，因为一开始选[0]作为枢轴）
 * @param high 结束位置（初始为数组最后一个数）
 */
func QuickSort(array []int, low, high int) {
	start := low      //开始位置（前端）
	end := high       //结束位置（后端）
	key := array[low] //关键值，也就是枢轴。第一次从位置0开始取，一轮排完会后排到中间。
	println(low, "~", high)
	for end > start { //
		//现在关键值在“前端”，从后往前比较，要找到小于关键值的值
		//如果比关键值大，则比较下一个，直到有比关键值小的交换位置，然后又从前往后比较
		for end > start && array[end] >= key {
			end-- //如果最后一个值大于关键值，则end往前移一位，拿倒数第二个比···
		}

		//由于之前的end--，现在是往前移了一位了，
		// 如果这时候刚好比关键值小，则将小的值和关键值交换位置。
		if array[end] <= key {
			array[start], array[end] = array[end], array[start]
		}
		//现在关键值在后端，从前往后比较，要找到大于关键值的值
		//如果比关键值小，则比较下一个，直到有比关键值大的交换位置
		for end > start && array[start] <= key {
			//从前端开始找，如果前端的值比目前处在后端的关键值小，
			// 则start++，将前端位置往后移一位
			start++
		}
		//由于前端往后移了一位，就再比一次，
		// 如果此时前端值刚好比关键值大，则交换位置，把关键值交换到前端。
		if array[start] >= key {
			array[start], array[end] = array[end], array[start]
		}
		//此时第一次循环比较结束，关键值（枢轴）的位置已经确定了。
		// 左边的值都比关键值小，右边的值都比关键值大，
		// 但是两边的顺序还有可能是不一样的，进行下面的递归调用
		printSlice(array)
		println(start, "-", end)
	}
	//递归，此时分别对枢轴两边进行快排
	//此时low是初始时的开始位置，start则++了好几次，low位至start位构成了左边序列。
	// low作为左边序列的起始位，start其实是枢轴的位置，所以start-1就是左边序列的结束位
	if start > low {
		QuickSort(array, low, start-1)
	}
	//此时end是枢轴的位置，所以end+1是右边序列的起始位，high是最初的结束为
	// （也就是最后一个数，期间改变的是end，high没变），所以high就是右边序列的结束位
	if end < high {
		QuickSort(array, end+1, high)
	}
}

//SelectionSort 选择排序
func SelectionSort(a []int) {
	// 它的工作原理是每一次从无序组的数据元素中选出最小（或最大）的一个元素，
	// 存放在无序组的起始位置，无序组元素减少，有序组元素增加，直到全部待排序的数据元素排完。
	for i := 0; i < len(a)-1; i++ {
		min := i                      //每次将min置成无序组起始位置元素下标
		for j := i; j < len(a); j++ { //遍历无序组，找到最小元素的位置
			if a[j] < a[min] {
				min = j
			}
		}
		if i != min { //如果最小元素不是无序组起始位置元素，则与起始元素交换位置
			a[i], a[min] = a[min], a[i]
		}
		printSlice(a)
	}

}

func printSlice(a []int) {
	for _, v := range a {
		print(v, ",")
	}
	println()
}

//InsertSort 插入排序
func InsertSort(a []int) {
	// 为了排序方便，我们一般将数据第一个元素视为有序组，其他均为待插入组。
	for j := 1; j < len(a); j++ {
		var key = a[j] //后面的带插入组的第一个
		var i = j - 1  //依次和前面的有序组比较
		for ; i >= 0 && a[i] > key; i-- {
			//a[i]>key是按照升序排列，a[i]<key是按照降序排列
			a[i+1] = a[i]
		}
		a[i+1] = key
		printSlice(a)
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
		if !run { //说明已经排序好了
			break
		}
		printSlice(a)
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
		if !run { //说明已经排序好了
			break
		}
		printSlice(p)
	}
}

//BubbleSort3 冒泡排序3
func BubbleSort3(p []int) {
	k := len(p) - 1
	for i := 0; i < len(p)-1; i++ {
		/*每排序一趟，则至少有一个元素已经有序，
		用 j<len-i-1 可以缩小排序范围 */
		run := false
		m := 0
		for j := 0; j < Min(k, len(p)-1-i); j++ {
			/*当前面的元素大于后面的元素时，交换位置*/
			if p[j] > p[j+1] {
				p[j], p[j+1] = p[j+1], p[j]
				run = true
				//记住最后一次交换的位置
				m = j
			}
		}
		if !run { //说明已经排序好了
			break
		}
		printSlice(p)
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
	println("\n---------BubleSort1--------")
	BubleSort1(temp)

	copy(temp, a)
	println("\n---------BubbleSort2--------")
	BubbleSort2(temp)

	copy(temp, a)
	println("\n---------BubbleSort3--------")
	BubbleSort3(temp)

	copy(temp, a)
	println("\n---------InsertSort--------")
	InsertSort(temp)

	copy(temp, a)
	println("\n---------SelectionSort--------")
	SelectionSort(temp)

	copy(temp, a)
	println("\n---------QuickSort--------")
	QuickSort(temp, 0, len(temp)-1)

	copy(temp, a)
	println("\n---------RadixSort--------")
	RadixSort(temp, len(temp))
}

//Min 求整数最小值
func Min(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}
