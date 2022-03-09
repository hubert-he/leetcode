package Stack

import "math"

/* 496. Next Greater Element I
** The next greater element of some element x in an array is the first greater element
	that is to the right of x in the same array.
** You are given two distinct 0-indexed integer arrays nums1 and nums2, where nums1 is a subset of nums2.
** For each 0 <= i < nums1.length, find the index j such that nums1[i] == nums2[j] and
	determine the next greater element of nums2[j] in nums2. If there is no next greater element,
	then the answer for this query is -1.
** Return an array ans of length nums1.length such that ans[i] is the next greater element as described above.
** Follow up: Could you find an O(nums1.length + nums2.length) solution?
 */
// 2022-02-15 刷出此题， 借助单调栈
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n := len(nums2)
	m := map[int]int{nums2[n-1]: -1}
	st := []int{nums2[n-1]} // 单调栈
	for i := n-2; i >= 0; i--{
		top := st[len(st)-1]
		if nums2[i] < top{
			m[nums2[i]] = top
		}else if nums2[i] >= top{
			st = st[:len(st)-1]
			m[nums2[i]] = -1
			for len(st) > 0{
				if nums2[i] > st[len(st)-1]{
					st = st[:len(st)-1]
				}else{
					m[nums2[i]] = st[len(st)-1]
					break
				}
			}
		}
		st = append(st, nums2[i])
	}
	ans := []int{}
	for _, c := range nums1{
		ans = append(ans, m[c])
	}
	return ans
}
// 官方题解
func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	stack := []int{}
	for i := len(nums2) - 1; i >= 0; i--{
		num := nums2[i]
		for len(stack) > 0 && num >= stack[len(stack)-1]{
			stack = stack[:len(stack) - 1]
		}
		if len(stack) > 0{
			m[num] = stack[len(stack)-1]
		}else{
			m[num] = -1
		}
		stack = append(stack, num)
	}
	ans := make([]int, len(nums1))
	for i, num := range nums1{
		ans[i] = m[num]
	}
	return ans
}

/* 503. Next Greater Element II
** Given a circular integer array nums (i.e., the next element of nums[nums.length - 1] is nums[0]),
** return the next greater number for every element in nums.
** The next greater number of a number x is the first greater number to its traversing-order next in the array,
** which means you could search circularly to find its next greater number.
** If it doesn't exist, return -1 for this number.
** Constraints:
	1 <= nums.length <= 10^4
	-10^9 <= nums[i] <= 10^9
 */
//2022-02-16 失败5次后 刷出，测试案例 前5个
func NextGreaterElements(nums []int) []int {
	n := len(nums)
	st := []int{} // 单调栈
	ans := make([]int, n)
	for i := n-1; i >= 0; i--{
		for len(st) > 0{
			if nums[i] >= st[len(st)-1]{
				st = st[:len(st)-1]
			}else{
				ans[i] = st[len(st)-1]
				break
			}
		}
		if len(st) == 0{
			ans[i] = math.MinInt32
		}
		st = append(st, nums[i])
	}
	for i := range nums{
		if ans[i] == math.MinInt32{
			for j := 0; j < i; j++{
				if nums[j] > nums[i]{
					ans[i] = nums[j]
					break
				}
			}
			if ans[i] == math.MinInt32{
				ans[i] = -1
			}
		}
	}
	return ans
}
// 官方单调栈
/* 单调栈
** 单调栈中保存数组下标，从栈底到栈顶的下标在数组 nums 中对应的值是单调不升的
** 每次我们移动到数组中的一个新的位置 i , 我们就将当前单调栈中所有对应值小于 nums[i]的下标弹出，
** 这些值的下一个更大元素即为 nums[i] （证明很简单：如果有更靠前的更大元素，那么这些位置将被提前弹出栈）
** 随后我们将位置 i 入栈
** 但是注意到只遍历一次序列是不够的， 例如序列 [2,3,1]，最后单调栈中将剩余[3,1] 其中元素 [1] 的下一个更大未知
** 此时利用循环数组处理
 */
/* 如何实现循环数组
** 循环数组的最后一个元素下一个元素是数组的第一个元素，形状类似于「环」
** 一种实现方式是，把数组复制一份到数组末尾，这样虽然不是严格的循环数组，但是对于本题已经足够了，因为本题对数组最多遍历两次
** 另一个常见的实现方式是，使用取模运算 % 可以把下标 i 映射到数组 nums 长度的 0 - N 内
 */
func NextGreaterElements2(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	st := []int{}
	// 把两个相同数组粘在一起
	for i := 0; i < n*2-1; i++{
		for len(st) > 0 && nums[st[len(st)-1]] < nums[i%n]{
			ans[st[len(st)-1]] = nums[i%n]
			st = st[:len(st)-1]
		}
		st = append(st, i%n)
	}
	return ans
}

/* 739. Daily Temperatures
** Given an array of integers temperatures represents the daily temperatures,
** return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature.
** If there is no future day for which this is possible, keep answer[i] == 0 instead.
 */
// 2022-02-17 刷出此题，解法使用的是逆向遍历， 可学习 正向遍历思路
// 此题与 上面的 题目的 不同点是， 此题必须存下标，计算距离
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	st := []int{}// 记录下标
	ans := make([]int, n)
	for i := n-1; i >= 0; i--{
		if len(st) == 0{
			ans[i] = 0
			st = append(st, i)
		}else{
			for len(st) > 0{
				if temperatures[st[len(st)-1]] > temperatures[i]{
					ans[i] = st[len(st)-1] - i
					break
				}else{
					st = st[:len(st)-1]
				}
			}
			if len(st)==0{
				ans[i] = 0
			}
			st = append(st, i)
		}
	}
	return ans
}
// 正向遍历, 此外可复用 temperature 数组
func dailyTemperatures2(temperatures []int) []int {
	st:= []int{}// 必须存下标
	for i := range temperatures {
		for len(st) > 0 &&  temperatures[i] > temperatures[st[len(st)-1]]{
			temperatures[st[len(st)-1]] = i - st[len(st)-1]
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	for len(st) > 0{
		temperatures[st[len(st)-1]] = 0
		st = st[:len(st)-1]
	}
	return temperatures
}