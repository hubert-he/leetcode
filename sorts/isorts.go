package sorts

import (
	"math"
)

func choosePivotMedianOfThree(a []int, left, right int) int{
	//根据l r 计算中间位置 mid
	/* 方式1： (right - left + 1) / 2 ，总数为偶数时，得到的是偏右的那个元素下标
	   方式2: (left + right) / 2, 总数为偶数时，得到的是偏左的那个元素下标
	 */
	mid := (left + right) >> 1
	// 求中位数
	if (a[left] - a[mid]) * (a[left] - a[right]) <= 0{
		return left
	}else if (a[mid] - a[left]) * (a[mid] - a[right]) <= 0{
		return mid
	}else{
		return right
	}
}

func quickSort(a []int, left, right int){
	// base condition
	if left >= right{
		return
	}
	var swap func(i, j int)
	var partition func(start, end int)int
	swap = func(i, j int){
		a[i],a[j] = a[j], a[i]
	}
	partition = func(start, end int) int{
		// pick pivot
		pivotIndex := choosePivotMedianOfThree(a, left, right)
		// 若第一个元素不是pivot, 需要将pivot与第一个元素进行交换，这样保证代码的统一性
		swap(pivotIndex, left)
		i := left + 1
		for j := left + 1; j <= right; j++{
			if a[j] < a[left] {
				swap(j, i)
				i++
			}
		}
		swap(left, i - 1)
		return i - 1
	}
	splitPos := partition(left, right)
	quickSort(a, left, splitPos - 1)
	quickSort(a, splitPos + 1, right)
}

func QuickSort(nums []int){
	var partition func() int
	var choosePivote func() int
	if len(nums) <= 1{
		return
	}
	choosePivote = func() int {
		end := len(nums) - 1
		mid := end / 2
		if (nums[0] - nums[mid]) * (nums[0] - nums[end]) <= 0{
			return 0
		}else if (nums[mid] - nums[0]) * (nums[mid] - nums[end]) <= 0{
			return mid
		}else{
			return end
		}
	}
	partition = func() int{
		pivotIdx := choosePivote()
		// 方便处理
		nums[pivotIdx], nums[0] = nums[0], nums[pivotIdx]
		i := 1
		for j := 1; j < len(nums); j++{
			if nums[j] < nums[0]{
				nums[j], nums[i] = nums[i], nums[j]
				i++
			}
		}
		nums[i-1], nums[0] = nums[0], nums[i-1]
		return i - 1
	}
	splitPos := partition()
	QuickSort(nums[:splitPos]) // 易错点-1： 切片前闭后开
	// QuickSort(nums[:splitPos-1])
	QuickSort(nums[splitPos+1:])
}

func ConcurrentQuickSort(nums []int, chanSend chan struct{}){
	length := len(nums)
	if length <= 1{
		chanSend <- struct{}{}
		return
	}
	// 并发优化
	if length < 100000{
		QuickSort(nums)
		chanSend <- struct{}{}
		return
	}
	var partition func() int
	var choosePivote func() int
	choosePivote = func() int {
		end := len(nums) - 1
		mid := end / 2
		if (nums[0] - nums[mid]) * (nums[0] - nums[end]) <= 0{
			return 0
		}else if (nums[mid] - nums[0]) * (nums[mid] - nums[end]) <= 0{
			return mid
		}else{
			return end
		}
	}
	partition = func() int{
		pivotIdx := choosePivote()
		nums[0], nums[pivotIdx] = nums[pivotIdx], nums[0]
		i := 1
		for j := 1; j < length; j++{
			if nums[j] < nums[0]{
				nums[j], nums[i] = nums[i], nums[j]
				i++
			}
		}
		nums[i-1], nums[0] = nums[0], nums[i-1]
		return i - 1
	}
	splitPos := partition()
	chanReceive := make(chan struct{})
	go ConcurrentQuickSort(nums[:splitPos], chanReceive)
	go ConcurrentQuickSort(nums[splitPos+1:], chanReceive)
	<-chanReceive
	<-chanReceive
	chanSend <- struct{}{}
	return
}

/* 剑指 Offer II 075. 数组相对排序
给定两个数组，arr1 和 arr2，
arr2 中的元素各不相同
arr2 中的每个元素都出现在 arr1 中
对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾。
示例：
	输入：arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
	输出：[2,2,2,1,4,3,3,9,6,7,19]
 */
func RelativeSortArray(arr1 []int, arr2 []int) []int {
	minValue, maxValue := math.MaxInt32, math.MinInt32
	for i := range arr1{
		if minValue > arr1[i]{
			minValue = arr1[i]
		}
		if maxValue < arr1[i]{
			maxValue = arr1[i]
		}
	}
	bucket := make([]int, maxValue - minValue + 1)
	for _, v := range arr1{
		bucket[v-minValue]++
	}
	i := 0
	for _, v := range arr2{
		k := v - minValue
		for idx := bucket[k]; idx > 0; idx--{
			arr1[i] = v
			i++
		}
		bucket[k] = 0
	}
	for k, v := range bucket{
		kk := k + minValue
		for j := v; j > 0; j--{
			arr1[i] = kk
			i++
		}
	}
	return arr1
}