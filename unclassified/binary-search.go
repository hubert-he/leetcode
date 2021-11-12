package unclassified

import (
	"fmt"
	"math"
	"sort"
)

/* Binary Search 两大基本原则
	1. 每次迭代都要缩减搜索区域  	Shrink the search scope in every iteration/recursion
	2. 每次缩减都不能排除掉潜在答案 	Can NOT exclude potential answers during each shrinking
*/
func SearchArray(a []int, key int) int {
	return binarySearchII(a, key)
}
// 模板一：查找准确值
// 循环条件： 	l <= r
// 缩减搜索空间：	l = mid + 1  &&  r = mid - 1
func binarySearch(a []int, key int) int {
	mid := -1
	i,j := 0, len(a)-1
	for i <= j { // 等号容易忘记，匹配i=j a[i] == key这种情况
		//mid = (j - i)/2 +i   // +i 容易忘记，注意除2只是偏移量
		mid = (i + j) >> 1  // 优化
		if key < a[mid] {
			j = mid - 1
		} else if key == a[mid]{
			break;
		} else {
			i = mid + 1
		}
	}
	if (i <= j) { // 等号容易忘记
		return mid
	} else {
		return -1
	}
}

// 模板二： 查找模糊值
// 循环条件：		l < r
// 缩减搜索空间：	l = mid, r = mid - 1(最后出现的) 或者 l = mid + 1, r = mid（最先出现的）


// 万用型模板：
// 循环条件：		l < r - 1
// 缩减搜索空间： 	l = mid, r = mid
func binarySearchII(a []int, key int) int {
	mid := -1
	l,r := 0, len(a)-1
	for l < r - 1 { // 剩余2个 l r 指向
		//mid = (r - l)/2 + l   // +i 容易忘记，注意除2只是偏移量
		mid = (r + l) >> 1  // 优化
		fmt.Printf("1=>%d 2=>%d\n", mid, (r+l)>>1)
		if key == a[mid]{
			return mid
		} else if key < a[mid] {
			r = mid
		} else {
			l = mid
		}
	}
	if a[l] == key {
		return l
	}
	if a[r] == key {
		return r
	}
	return -1
}

/* GoLang sort 库中 Search 代码 优秀Binary Search 实现
	arr := []int{1,3,5,7}
	index := sort.Search(len(arr), func(i int)bool{
		return arr[i] >= 4 // 返回2 标识 插入的下标位置
	})
	优势：
	1. 当有多个元素都等于target时，实际上可以找到下标最小的那个元素
	2. 当target不存在时，返回的下标标识了如果要将target插入，插入的位置（插入后依然保持数组有序）
	3. 基于这种编程范式，上层只用提供一个与自身数据类型相关的比较函数
   golang sort 提供的基础类型查找函数
	// - 只需传入整型切片和target
	// - 注意，该函数只能对升序数组做查找
	func SearchInts(a []int, x int) int {
		return Search(len(a), func(i int) bool { return a[i] >= x })
	}
	func SearchFloat64s(a []float64, x float64) int {
		return Search(len(a), func(i int) bool { return a[i] >= x })
	}
	func SearchStrings(a []string, x string) int {
		return Search(len(a), func(i int) bool { return a[i] >= x })
	}
 */
func Search(n int, f func(int)bool) int{
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // 防止下标 相加 溢出
		if !f(h){ // 如果序列是升序，比较函数应该用 >=   降序 <=
			i = h + 1 // 转到右半边去处理
		} else {
			j = h // 并不退出循环，而是继续查找，锁定最小的下标（如果有重复元素）
		}
	}
	return i
}

/* 74. Search a 2D Matrix
** Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:
** Integers in each row are sorted from left to right.
** The first integer of each row is greater than the last integer of the previous row.
*/
func SearchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	search := func(t int)int{
		i, j := 0, m*n
		for i < j {
			mid := int(uint(i+j)>>1)
			if matrix[mid/n][mid%n] < t{
				i = mid+1
			}else{
				j = mid
			}
		}
		return i
	}
	ans := search(target)
	return ans < m*n && matrix[ans/n][ans%n] == target
}
func SearchMatrix2(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	moreEqual := func(t int)bool{
		return matrix[t/n][t%n] >= target
	}
	// 配置为 moreEqual 说明 查找的是第一个大于等于target 的
	i := sort.Search(m*n, moreEqual)
	return i < m*n && matrix[i/n][i%n] == target
}
// 思路2：两次二分查找，应对 二维数组中的一维数组的元素个数不一
func SearchMatrix3(matrix [][]int, target int) bool {
	more := func(t int)bool{
		return matrix[t][0] > target
	}
	// 配置为 more 说明 查找的是第一个大于target 的
	row := sort.Search(len(matrix), more) - 1
	if row < 0{
		return false
	}
	col := sort.SearchInts(matrix[row], target)
	return col < len(matrix[row]) && matrix[row][col] == target
}

/* 162. Find Peak Element
** A peak element is an element that is strictly greater than its neighbors.
Given an integer array nums, find a peak element, and return its index.
If the array contains multiple peaks, return the index to any of the peaks.
You may imagine that nums[-1] = nums[n] = -∞.
 */
/* 注意是找其中一个峰值，而非最大的峰值
** 根据 nums[i−1],nums[i],nums[i+1] 三者的关系决定
** 1. nums[i−1] < nums[i] > nums[i+1]   找到一个峰值，返回 i
** 2. nums[i−1] < nums[i] < nums[i+1]	i 上坡 i = i+1
** 3. nums[i−1] > nums[i] > nums[i+1]   i 下坡 i = i-1
** 4. nums[i−1] > nums[i] < nums[i+1]	任意方向
** 根据第4个情况，缩减情况分支，规定第4情况 i = i+1
** nums[i] < nums[i+1] i = i+1
** nums[i] > nums[i+1] i = i-1
** 另外， nums[i] < nums[i+1], i = i+1 ，那么位置左侧的所有位置都不可能在后续迭代找那个遍历到。
** 「二段性」其实是指：在以 mid 为分割点的数组上，根据 nums[mid] 与 nums[mid±1] 的大小关系，
** 可以确定其中一段满足「必然有解」，另外一段不满足「必然有解」（可能有解，可能无解）
 */
func FindPeakElement(nums []int) int {
	n := len(nums)
	get := func(i int)int{
		if i == -1 || i >= n{
			return math.MinInt64 // 必须是64 32的不满足最小值
		}
		return nums[i]
	}
	i, j := 0, n
	for i < j{
		mid := int(uint(i+j)>>1)
		if get(mid-1) < get(mid) && get(mid) > get(mid+1){
			return mid
		}
		if get(mid) > get(mid+1){
			j = mid-1
		}else{
			i = mid+1
		}
	}
	return i
}
/* 始终选择大于边界一端进行二分，可以确保选择的区间一定存在峰值，并随着二分过程不断逼近峰值位置。*/
func FindPeakElement2(nums []int) int {
	n := len(nums)
	get := func(i int)int{
		if i == -1 || i >= n{
			return math.MinInt64 // 必须是64 32的不满足最小值
		}
		return nums[i]
	}
	i, j := 0, n
	for i < j{
		mid := int(uint(i+j)>>1)
		if get(mid) > get(mid+1){
			j = mid
		}else{
			i = mid+1
		}
	}
	return i
}