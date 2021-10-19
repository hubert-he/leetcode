package heap

import (
	"container/heap"
	"fmt"
)

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

/* 239. Sliding Window Maximum
You are given an array of integers nums,
there is a sliding window of size k which is moving from the very left of the array to the very right.
You can only see the k numbers in the window. Each time the sliding window moves right by one position.
Return the max sliding window.
 */
// 朴素解法，竟然可以AC
func MaxSlidingWindow(nums []int, k int) []int {
	start, end := 0, k-1
	m := 0
	ans := []int{}
	max := func (s int, e int) int {
		m := s
		for i := s; i <= e; i++{
			if nums[m] <= nums[i]{// 必须带上= 使得相同的尽量在里面
				m = i
			}
		}
		return m
	}
	m = max(start, end)
	for end < len(nums){
		if m >= start && m <= end{
			if nums[m] <= nums[end]{ // 必须带上= 使得相同的尽量在里面
				m = end
			}
		}else{
			m = max(start, end)
		}
		ans = append(ans, nums[m])
		start++
		end++
	}
	return ans
}
// 解法2： 使用heap  O(nLogn)
type hp struct {
	nums    []int // 参与比较
	data    []int // 存放nums索引
}
func (h hp)Len()int{ return len(h.data) }
func (h hp)Less(i, j int)bool { return h.nums[h.data[i]] > h.nums[h.data[j]] } // 大顶堆
func (h hp)Get() int { return h.data[0] }
func (h *hp)Swap(i, j int){ h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *hp)Push(v interface{}) { h.data = append(h.data, v.(int)) }
func (h *hp)Pop()interface{} { ret := h.data[h.Len()-1]; h.data = h.data[:h.Len()-1]; return ret}

func MaxSlidingWindowHeap(nums []int, k int) []int {
	n := len(nums)
	h := hp{nums: nums}
	for i := 0; i < k; i++{
		h.Push(i)
	}
	heap.Init(&h) // 需要定义一个struct，其实现了heap.Interface接口，包含Len() Swap() Less() Push() Pop()方法
	ans := []int{nums[h.Get()]}
	for i := k; i < n; i++{
		heap.Push(&h, i)
		//if h.Get() <= i-k{// 最大的 超出范围了
		for h.Get() <= i - k{ // 这边要用 for 循环Pop 超出范围的
			fmt.Println(nums[h.Get()], h.Len())
			heap.Pop(&h)
			fmt.Println(nums[h.Get()], h.Len())
		}
		ans = append(ans, nums[h.Get()])
	}
	return ans
}
/* 方法三： 单调队列  O(n)
双端队列： 保证 队尾始终为新加的元素，其他小于队尾的 统统弹出
        队首保证始终在 有效范围内，不在有效范围的 统统弹出
 */
func MaxSlidWindowDQ(nums []int, k int) []int{
	q := []int{}
	push := func(i int){
		// 不断的弹出队尾元素，直至 新加入的元素 变为队列中最小或者是队首
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]]{
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	for i := 0; i < k; i++{
		push(i)
	}

	n := len(nums)
	ans := []int{q[0]}
	for i := k; i < n; i++{
		push(i)
		for q[0] <= i-k{
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}











