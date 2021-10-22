package utils

import "math"

func Max(nums ...int)int{
	ans := nums[0]
	for i := 1; i < len(nums); i++{
		if nums[i] > ans {
			ans = nums[i]
		}
	}
	return ans
}

func Min(nums ...int) int {
	m := math.MaxInt32
	for _, n := range nums{
		if m > n{
			m = n
		}
	}
	return m
}
// 排除except 下标的 值
func Min2(except int, nums ...int) int{
	m := math.MaxInt32
	for i := range nums{
		if m > nums[i] && i != except{
			m = nums[i]
		}
	}
	return m
}
