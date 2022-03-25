package Array

import "math"

/* 1588. Sum of All Odd Length Subarrays
** Given an array of positive integers arr, calculate the sum of all possible odd-length subarrays.
** A subarray is a contiguous subsequence of the array.
** Return the sum of all odd-length subarrays of arr.
 */
//2022-02-15 刷出 暴力枚举
func sumOddLengthSubarrays(arr []int) int {
	n := len(arr)
	ans := 0
	for i := 1; i <= n; i = i+2{
		for j := range arr{
			// 本质上就是求连续子数组的和，所以可以考虑前缀和解决
			for k := j; j+i <= n && k < j+i; k++{
				ans += arr[k]
			}
		}
	}
	return ans
}
//2022-02-15 刷出 前缀和
func sumOddLengthSubarrays_prefixsum(arr []int) int {
	n := len(arr)
	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++{
		prefixSum[i] = prefixSum[i-1] + arr[i-1]
	}
	ans := prefixSum[n]
	for i := 3; i <= n; i = i + 2{
		for j := i; j <= n; j++{
			ans += prefixSum[j] - prefixSum[j-i]
		}
	}
	return ans
}
// 官方题解优化 到 O(n)
/* 就是遍历一遍所有的元素，然后查看这个元素会在多少个长度为奇数的数组中出现过
** 测试用例 [1, 4, 2, 5, 3]
** 	1 在 3 个长度为奇数的数组中出现过：[1], [1, 4, 2], [1, 4, 2, 5, 3]；所以最终的和，要加上 1 * 3；
	4 在 4 个长度为奇数的数组中出现过：[4], [4, 2, 5], [1, 4, 2], [1, 4, 2, 5, 3]；所以最终和，要加上 4 * 4；
	2 在 5 个长度为奇数的数组中出现过：[2], [2, 5, 3], [4, 2, 5], [1, 4, 2], [1, 4, 2, 5, 3]；所以最终和，要加上 5 * 2；
	...
** 下面的关键就是，如何计算一个数字在多少个奇数长度的数组中出现过？
** 对于一个数字，它所在的数组，可以在它前面再选择 0, 1, 2, ... 个数字，一共有 left = i + 1 个选择；
** 可以在它后面再选择 0, 1, 2, ... 个数字，一共有 right = n - i 个选择
** 有2中组合情况：
	1. 如果在前面选择了偶数个数字，那么在后面，也必须选择偶数个数字，这样加上它自身，才构成奇数长度的数组
	2. 如果在前面选择了奇数个数字，那么在后面，也必须选择奇数个数字，这样加上它自身，才构成奇数长度的数组；
** 数字前面共有 left 个选择， 其中偶数个数字的选择方案有 left_even = (left+1)/2 个，奇数个数字的选择方案有 left_odd = left / 2 个
** 数字后面共有 right 个选择，其中偶数个数字的选择方案有 right_even = (right+1) / 2 个， 奇数个数字的选择方案有 right_odd = right / 2个
** 所以，每个数字一共在 left_even * right_even + left_odd * right_odd 个奇数长度的数组中国出现过
*/
/* 2022-03-24 重刷此题，发现一个 潜在的思维问题 */
func sumOddLengthSubarrays_math(arr []int) int {
	ans, n := 0, len(arr)
	for i := range arr{
		left , right := i + 1, n - i
		left_even, right_even := (left+1)/2, (right+1)/2
		left_odd, right_odd := left/2, right/2
		ans += (left_even * right_even + left_odd * right_odd) * arr[i]
	}
	return ans
}

func sumOddLendthSubarray_error(arr []int) int{
	ans, n := 0, len(arr)
	for i := range arr{
		left := i+1 // 注意不是 i-1，也不是 i， 因为 i 是从 0 开始计数，并且 0个的情况也算一种可能
		right := n-i
		times := (left+1)/2 * (right+1)/2 + left/2*right/2 // 此行问题
		// 如果直接运算，(left+1)/2 * (right+1)/2 误差的问题 会导致结果偏大。left/2*right/2 也会导致
		// 首先思路是 可能的情况数，left+1 也是考虑的情况数，floor
		ans += times * arr[i]
	}
	return ans
}
// 去数字的Floor值
func sumOddLengthSubarrays_error_fixed(arr []int) int {
	n := len(arr)
	sum := 0
	for i := range arr{
		//left, right := i-1, n-i
		left, right := i+1, n-i // 0 也算一种选择， i 是从 0 开始计数
		even_times := math.Floor(float64(left+1)/2) * math.Floor(float64((right+1)/2))
		odd_times := math.Floor(float64(left/2))*math.Floor(float64(right/2))
		sum += arr[i] * int(even_times + odd_times)
	}
	return sum
}
