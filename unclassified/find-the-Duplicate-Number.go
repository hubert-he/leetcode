/*
给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。
示例 1:
输入: [1,3,4,2,2]
输出: 2
示例 2:
输入: [3,1,3,4,2]
输出: 3
说明：
不能更改原数组（假设数组是只读的）。
只能使用额外的 O(1) 的空间。
时间复杂度小于 O(n2) 。
数组中只有一个重复的数字，但它可能不止重复出现一次。
链接：https://leetcode-cn.com/problems/find-the-duplicate-number
note：
题眼： 包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n）
 */
package unclassified

func FindDuplicate(nums []int) int {
	return binarySearchFind(nums)
}

func normalFind(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return nums[i]
			}
		}
	}
	return -1
}

func binarySearchFind(nums []int) int {
	l := 1
	r := len(nums) - 1
	for l < r {
		mid := (l + r) >> 1
		if(check(nums, mid)) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
func check(nums []int, mid int) bool {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <= mid {
			cnt++
		}
	}
	if cnt <= mid{
		return true
	}else {
		return false
	}

}
