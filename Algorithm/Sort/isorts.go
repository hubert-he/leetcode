package Sort

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"unicode"
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

/*10.01. Sorted Merge LCCI   <--- 强化练习
You are given two sorted arrays, A and B, where A has a large enough buffer at the end to hold B.
Write a method to merge B into A in sorted order.
Initially the number of elements in A and B are m and n respectively.
这是归并排序中Merge实现的部分，原地Merge
在原地归并排序中最主要用到了内存反转, 即交换相邻两块内存，在《编程珠玑》中又被称为手摇算法。
内存反转：给定序列a1,a2,a3,...,am,b1,b2,b3,...,bm , 将其变为 b1,b2,b3,...,bm,a1,a2,a3,...,am
手摇算法：先对a1,a2,a3,...,am 反转， 再对 b1,b2,b3,...,bm 反转，最后对am,...,a3,a2,a1,bm,...b3,b2,b1整体反转继而得到
	b1,b2,b3,...,bm,a1,a2,a3,...,am
 */
func MergeInPlace(A []int, m int, B []int, n int)  {
	copy(A[m:], B)
	fmt.Println(A)
	reverse := func(s, e int){
		for i, j := s, e; i <= j; i, j = i+1, j-1{
			A[i], A[j] = A[j], A[i]
		}
	}
	swap := func(s, m, e int){
		reverse(s, m)
		reverse(m+1, e)
		reverse(s,e)
	}
	i, j, s := 0, m, m+n
	//for i < m && j < s{
	for i < j && j < s{
		for i < j && A[i] <= A[j]{
			i++
		}
		index := j
		for j < s && A[i] > A[j]{
			j++
		}
		// 交换[i, index) 和 [index,j)的内存块
		swap(i, index-1, j-1)
		i += j - index
	}
}
// 此解法仅练习双指针，没有其他用处
func Merge2Pointer(A []int, m int, B []int, n int)  {
	if m == 0 || n == 0{
		copy(A[m:],B)
		return
	}
	i,j := m-1, n-1
	for p := m+n-1; i >= 0 && j >= 0 && i < p ; p--{
		if A[i] <= B[j]{
			A[p] = B[j]
			j--
		}else{
			A[p] = A[i]
			i--
		}
	}
	if i < 0 && j >= 0{
		copy(A, B[:j+1])
	}
}

/* 937. Reorder Data in Log Files
You are given an array of logs. Each log is a space-delimited string of words, where the first word is the identifier.
There are two types of logs:
	Letter-logs: All words (except the identifier) consist of lowercase English letters.
	Digit-logs: All words (except the identifier) consist of digits.
Reorder these logs so that:
	The letter-logs come before all digit-logs.
	The letter-logs are sorted lexicographically by their contents.
	If their contents are the same, then sort them lexicographically by their identifiers.
	The digit-logs maintain their relative ordering.
	Return the final order of the logs.
 */
func ReorderLogFiles(logs []string) []string {
	n := len(logs)
	lessEqual := func(i, j int)bool{ // 传递索引，比直接传string要优秀
		part_i := strings.SplitN(logs[i], " ", 2)
		part_j := strings.SplitN(logs[j], " ", 2)
		isDigit_i := unicode.IsDigit(rune(part_i[1][0]))
		isDigit_j := unicode.IsDigit(rune(part_j[1][0]))
		if !isDigit_i && !isDigit_j{
			return part_i[1] < part_j[1] || (part_i[1] == part_j[1] && part_i[0] < part_j[0])
		}
		if isDigit_i && isDigit_j{
			return true
		}
		return !isDigit_i && isDigit_j
	}
	merge := func(s, m, e int){
		tmp := []string{}
		i, j := s, m+1
		for i <= m && j <= e{
			if lessEqual(i, j){
				fmt.Println(logs[i], "<", logs[j])
				tmp = append(tmp, logs[i])
				i++
			}else{
				fmt.Println(logs[i], ">=", logs[j])
				tmp = append(tmp, logs[j])
				j++
			}
		}
		for i <= m{
			tmp = append(tmp, logs[i])
			i++
		}
		for j <= s{
			tmp = append(tmp, logs[j])
			j++
		}
		copy(logs[s:], tmp)
	}
	var dfs func(start, end int)
	dfs = func(start, end int){
		if start >= end{
			return
		}
		mid := (start & end) + (start ^ end)>>1
		dfs(start, mid)
		dfs(mid+1, end)
		merge(start, mid, end)
	}
	dfs(0, n-1)
	return logs
}
func less(a, b string)bool{
	ai, bi := strings.Index(a, " "), strings.Index(b, " ")
	a_isDigit, b_isDigit := false, false
	if a[ai+1] >= '0' && a[ai+1] <= '9'{
		a_isDigit = true
	}
	if b[bi+1] >= '0' && b[bi+1] <= '9'{
		b_isDigit = true
	}
	if  a_isDigit && b_isDigit {
		return true
	}else if !a_isDigit && b_isDigit {
		return true
	}else if a_isDigit && !b_isDigit {
		return false
	}else{
		content := strings.Compare(string(a[ai+1:]), string(b[bi+1:]))
		if content == 0{
			return strings.Compare(string(a[:ai-1]), string(b[bi-1])) < 0
		}else {
			return content < 0
		}
	}
}
/*练习使用 golang 标准库
  sort.SliceStable 与 sort.Slice 一个稳定排序，一个不稳定排序
  func SliceStable(x interface{}, less func(i, j int) bool)
  SliceStable sorts the slice x using the provided less function, keeping equal elements in their original order.
  It panics if x is not a slice.
  The less function must satisfy the same requirements as the Interface type's Less method.
 */
func ReorderLogFilesLib(logs []string) []string {
	less := func(i, j int)bool{ // 传递索引，比直接传string要优秀
		part_i := strings.SplitN(logs[i], " ", 2)
		part_j := strings.SplitN(logs[j], " ", 2)
		isDigit_i := unicode.IsDigit(rune(part_i[1][0]))
		isDigit_j := unicode.IsDigit(rune(part_j[1][0]))
		if !isDigit_i && !isDigit_j{
			return part_i[1] < part_j[1] || (part_i[1] == part_j[1] && part_i[0] < part_j[0])
		}
		return !isDigit_i && isDigit_j
	}
	sort.SliceStable(logs, less)
	return logs
}