package heap

/* 17.14. Smallest K LCCI
	Design an algorithm to find the smallest K numbers in an array.
Example:
	Input:  arr = [1,3,5,7,2,4,6,8], k = 4
	Output:  [1,2,3,4]
 */
/* 最小堆的创建 -- O(kn) 超时
 */
func SmallestK(arr []int, k int) []int {
	shiftUp(arr)
	ans := []int{}
	for i := 0; i < k; i++{
		ans = append(ans, arr[i])
		shiftUp(arr[i+1:])
	}
	return ans
}
func shiftUp(a []int){
	n := len(a)
	for i := n/2 - 1; i >= 0; i--{
		l,r  := 2*i+1, 2*i+2
		if l < n && a[l] < a[i] {
			a[l], a[i] = a[i], a[l]
		}
		if r < n && a[r] < a[i] {
			a[r], a[i] = a[i], a[r]
		}
	}
}
// O(nlogn) shift down 堆化
func SmallestK2(arr []int, k int) []int {
	n := len(arr)
	down := func(i int)bool{
		t := i
		for {
			l := 2*i + 1
			if l >= n || l < 0{
				break
			}
			j := l
			// 先孩子比较
			if r := l + 1; r < n && arr[r] < arr[l]{
				j = r
			}
			// 与父比较
			if arr[i] < arr[j]{
				break
			}
			arr[i], arr[j] = arr[j], arr[i]
			i = j
		}
		return i > t
	}
	// heapify
	for i := n/2 - 1; i >= 0; i--{
		down(i)
	}
	ans := []int{}
	for i := 0; i < k; i++{
		ans = append(ans, arr[0])
		arr[0], arr[n-1] = arr[n-1], arr[0]
		n--
		arr = arr[:n]
		down(0)
	}
	return ans
}
/* 利用快排 可达到最优
注意到题目要求「任意顺序返回这 k 个数即可」，因此我们只需要确保前k小的数都出现在下标为 [0, k) 的位置即可。
利用「快速排序」的数组划分即可做到。
我们知道快排每次都会将小于等于基准值的值放到左边，将大于基准值的值放到右边。
因此我们可以通过判断基准点的下标 idx 与 k 的关系来确定过程是否结束：
idx < k：基准点左侧不足 k 个，递归处理右边，让基准点下标右移；
idx > k：基准点左侧超过 k 个，递归处理左边，让基准点下标左移；
idx = k：基准点左侧恰好 k 个，输出基准点左侧元素。
 */
func smallestKQS(arr []int, k int) []int {
	if k == 0{
		return nil
	}
	var qSort func(start, end int)
	qSort = func(start, end int){
		if start >= end {
			return
		}
		var choosePivote func()int
		var partition func()int
		choosePivote = func()int{
			mid := (end-start+1)>>1
			if (arr[start] - arr[end]) * (arr[start] - arr[mid]) < 0{
				return start
			}else if (arr[mid] - arr[start]) * (arr[mid] - arr[end]) < 0{
				return mid
			}else{
				return end
			}
		}
		partition = func()int{
			pivotIdx := choosePivote()
			arr[start], arr[pivotIdx] = arr[pivotIdx], arr[start]
			i := start + 1 // 双指针
			for j := start+1; j <= end; j++{
				if arr[j] < arr[start]{
					arr[j], arr[i] = arr[i], arr[j]
					i++
				}
			}
			arr[i-1], arr[start] = arr[start], arr[i-1]
			return i - 1
		}
		splitPos := partition()
		if splitPos > k{
			qSort(start, splitPos-1)
		}
		if splitPos < k{
			qSort(splitPos+1, end)
		}
	}
	qSort(0, len(arr) - 1)
	return arr[:k+1]
}

















