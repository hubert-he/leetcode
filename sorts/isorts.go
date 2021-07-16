package sorts

func choosePivotMedianOfThree(a []int, left, right int) int{
	//根据l r 计算中间位置 mid
	/* 方式1： (right - left + 1) / 2 ，总数为偶数时，得到的是偏右的那个元素下标
	   方式2: (left + right) / 2, 总数为偶数时，得到的是偏左的那个元素下标
	 */
	mid := (left + right) / 2
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
