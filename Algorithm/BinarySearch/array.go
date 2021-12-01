package BinarySearch

import (
	"math"
)

/* 1060. Missing Element in Sorted Array
** Given an integer array nums which is sorted in ascending order and all of its elements are unique and given also an integer k,
** return the kth missing number starting from the leftmost number of the array.
 */
// 2021-11-24 刷出此题
// 找到一个最小的区间 [l, r] 包含答案
func MissingElement(nums []int, k int) int {
	n := len(nums)
	// 处理边界
	dis := nums[n-1] - nums[0] - n + 1
	if dis < k{
		return nums[n-1] + k - dis
	}
	i, j := 0, n-1
	for j - i > 1{
		mid := int(uint(i+j)>>1)
		d := (nums[mid] - nums[i]) - (mid - i)
		if k > d{
			k -= d
			i = mid
		}else{
			j = mid
		}
	}
	return nums[i]+k
}
// 官方题解
func MissingElement2(nums []int, k int) int {
	n := len(nums)
	missing := func(idx int)int{
		return nums[idx] - nums[0] - idx
	}
	if k > missing(n-1){ // 边界，k超出nums最大可能值
		return nums[n-1] + k - missing(n-1)
	}
	left, right := 0, n-1
	// 找到left等于right，即 missing(left - 1) < k <= missing(left)
	for left < right{
		mid := int(uint(left+right)>>1)
		if missing(mid) < k{
			left = mid+1
		}else{
			right = mid
		}
	}
	// kth missing number is larger than nums[idx - 1] and smaller than nums[idx]
	return nums[left-1] + k - missing((left-1))
}

/*1901. Find a Peak Element II
** A peak element in a 2D grid is an element that is strictly greater than all of its adjacent neighbors to the left, right, top, and bottom.
** Given a 0-indexed m x n matrix mat where no two adjacent cells are equal, find any peak element mat[i][j] and return the length 2 array [i,j].
** You may assume that the entire matrix is surrounded by an outer perimeter with the value -1 in each cell.
** You must write an algorithm that runs in O(m log(n)) or O(n log(m)) time.
** Note: No two adjacent cells are equal.
 */
// 二分思路： 相邻元素各不相同
func findPeakGrid(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	maxRow := func(r int)(c int){
		if r < 0 || r >= m{
			return -1
		}
		for i := 1; i < n; i++{
			if mat[r][i] > mat[r][c]{
				c = i
			}
		}
		return
	}
	i, j := 0, m-1
	ans := []int{}
	for i <= j{
		mid := int(uint(i+j) >> 1)
		up, cur, down := maxRow(mid-1), maxRow(mid), maxRow(mid+1)
		upVal, curVal, downVal := -1, -1, -1
		if up != -1{
			upVal = mat[mid-1][up]
		}
		if cur != -1{
			curVal = mat[mid][cur]
		}
		if down != -1{
			downVal = mat[mid+1][down]
		}
		// 中间行最大，并且又是行内最大值，所以找到 直接 return
		if curVal >= upVal && curVal >= downVal{
			ans = append(ans, mid, cur)
			break
		}
		if upVal > curVal && upVal > downVal{
			j = mid - 1
		}else{
			i = mid + 1
		}
	}
	return ans
}

/* 1231. Divide Chocolate
** You have one chocolate bar that consists of some chunks. Each chunk has its own sweetness given by the array sweetness.
You want to share the chocolate with your k friends so you start cutting the chocolate bar into k + 1 pieces using k cuts, each piece consists of some consecutive chunks.
Being generous, you will eat the piece with the minimum total sweetness and give the other pieces to your friends.
Find the maximum total sweetness of the piece you can get by cutting the chocolate bar optimally.
*/
func MaximizeSweetness(sweetness []int, k int) int {
	var dfs func(arr []int, part int)[][][]int
	dfs = func(arr []int, part int)(ret [][][]int){
		n := len(arr)
		if n == 0 || part == 0{
			return
		}
		if n < part{
			return
		}
		if part == 1 {
			ret = append(ret, [][]int{arr})
			return
		}
		if n == part{
			t := [][]int{}
			for i := range arr{
				t = append(t, []int{arr[i]})
			}
			ret = append(ret, t)
			return
		}
		for i := 0; n - i >= part ; i++{
			sub := dfs(arr[i+1:], part-1)
			//fmt.Println(len(sub))
			for j := range sub{
				sub[j] = append(sub[j], arr[:i+1])
				ret = append(ret, sub[j])
			}
		}
		return
	}
	all := dfs(sweetness, k+1)
	ans := 0
	/*
	   for i := range all{
	       fmt.Println(all[i])
	   }
	*/
	for i := range all{
		t := math.MaxInt32
		for j := range all[i]{
			s := sum(all[i][j]...)
			if t > s{
				t = s
			}
		}
		ans = max(ans, t)
	}
	return ans
}
func sum(nums ...int)int{
	m := 0
	for _, c := range nums{
		m += c
	}
	return m
}

func max(nums ...int)int{
	m := nums[0]
	for _, c := range nums{
		if m < c {
			m = c
		}
	}
	return m
}

func MaximizeSweetnessBS(sweetness []int, k int) int {
	sum := 0
	minVal := math.MaxInt32
	for i := range sweetness{
		sum += sweetness[i]
		if minVal > sweetness[i]{
			minVal = sweetness[i]
		}
	}
	if k == 0{
		return sum
	}
	// 计算sweet最小甜度块时， 能够切出块的最大数量
	calccount := func(sweet int)int{
		total := 0
		cnt := 0
		for i := range sweetness{
			total += sweetness[i]
			if total >= sweet{
				cnt++
				total = 0
			}
		}
		return cnt
	}
	// 求 right_bound
	left, right := minVal, sum
	for left <= right{ // 搜索终止： [right+1, right] [left,left-1]
		mid := int(uint(left+right)>>1)
		count := calccount(mid)
		if k+1 < count{ // 证明还有变大的可能
			left = mid + 1
		}else if k+1 > count{
			right = mid - 1
		}else{
			left = mid + 1
		}
	}
	return right
}

/* 287. Find the Duplicate Number
** Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.
** There is only one repeated number in nums, return this repeated number.
** You must solve the problem without modifying the array nums and uses only constant extra space.
 */
/* 此题可以使用 floyd 快慢指针 同时也可bit 现在用二分
** 定义 cnt[i] 表示 nums 数组中小于等于 i 的数有多少个
** 假设我们重复的数是 target，那么[1,target−1]里的所有数满足[target,n] 里的所有数满足 cnt[i]>i，具有单调性
 */
func findDuplicate(nums []int) int {
	n := len(nums)
	l, r := 1, n-1
	ans := -1
	for l <= r{
		mid := int(uint(l+r)>>1)
		cnt := 0
		for i := range nums{
			if nums[i] <= mid{
				cnt++
			}
		}
		if cnt <= mid{
			l = mid + 1
		}else{// cnt > mid 结果可能在 [l, r]中
			r = mid - 1
			ans = mid
		}
	}
	return ans
}
// 另外一种写法
func findDuplicate2(nums []int) int {
	n := len(nums)
	l, r := 1, n-1
	for l < r{ // 搜索空间 [l, r) == > 终止条件 [r, r)
		mid := int(uint(l+r)>>1)
		cnt := 0
		for i := range nums{
			if nums[i] <= mid{
				cnt++
			}
		}
		if cnt <= mid{
			l = mid + 1
		}else{// cnt > mid 结果可能在 [l, r]中
			r = mid
		}
	}
	return r
}






