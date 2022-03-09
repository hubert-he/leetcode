package TwoPointer

import "math"

/* 31. Next Permutation
** Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.
** If such an arrangement is impossible, it must rearrange it to the lowest possible order (i.e., sorted in ascending order).
** The replacement must be in place and use only constant extra memory.
** Constraints:
	1 <= nums.length <= 100
	0 <= nums[i] <= 100
*/
/*
** 我们可以将该问题形式化地描述为：给定若干个数字，将其组合为一个整数。
** 如何将这些数字重新排列，以得到下一个更大的整数。如 123 下一个更大的数为 132。
** 如果没有更大的整数，则输出最小的整数。
 */
/* 将给定数字序列重新排列成字典序中下一个更大的排列
** 要求：下一个排列总是比当前排列要大，除非该排列已经是最大的排列，能够找到一个大于当前序列的新序列，且变大的幅度尽可能小
** 1. 我们希望下一个数比当前数大，这样才满足“下一个排列”的定义。
		因此只需要将后面的「大数」与前面的「小数」交换，就能得到一个更大的数。比如 123456，将 5 和 6 交换就能得到一个更大的数 123465
** 2. 同时我们要让这个「较小数」尽量靠右，而「较大数」尽可能小。当交换完成后，「较大数」右边的数需要按照升序重新排列。
**    这样可以在保证新排列大于原来排列的情况下，使变大的幅度尽可能小
	  即希望下一个数增加的幅度尽可能的小，这样才满足“下一个排列与当前排列紧邻“的要求，需要：
		2.1 在尽可能靠右的低位进行交换，需要从后向前查找
		2.2 将一个 尽可能小的「大数」 与前面的「小数」交换。比如 123465，下一个排列应该把 5 和 4 交换而不是把 6 和 4 交换
			将「大数」换到前面后，需要将「大数」后面的所有数重置为升序，升序排列就是最小的排列。
			以 123465 为例：首先按照上一步，交换 5 和 4，得到 123564；然后需要将 5 之后的数重置为升序，得到 123546。
			显然 123546 比 123564 更小，123546 就是 123465 的下一个排列
** 算法：
** 首先从后向前查找第一个顺序对 (i,i+1)，满足a[i] < a[i+1]。这样「较小数」即为 a[i]。此时[i+1, n)必然是下降序列
** 如果找到了顺序对，那么在区间 [i+1,n) 中从后向前查找第一个元素 j 满足 a[i] < a[j]。这样「较大数」即为 a[j]
** 交换a[i] 与 a[j] 此时可以证明区间[i+1, n)必为降序。可以直接使用双指针反转区间 [i+1,n)使其变为升序，而无需对该区间进行排序
 */
func nextPermutation(nums []int)  {
	n := len(nums)
	little, much := -1, -1
	for i := n-2; i >= 0; i--{
		if nums[i] < nums[i+1]{
			little = i
			break
		}
	}
	if little != -1{
		for i := n-2; i >= little+1; i--{
			if nums[i] > nums[i+1]{
				much = i
			}
		}
		nums[little], nums[much] = nums[much], nums[little]
	}
	i, j := little+1, n-1
	for i < j{
		nums[i], nums[j] = nums[j], nums[i]
	}
}


/* 283. Move Zeroes
** Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
** Note that you must do this in-place without making a copy of the array.
 */
// 2022-02-15 刷出此题， zerop 始终指向为 0 的点， 但是不够简洁
func moveZeroes(nums []int)  {
	zerop := -1
	n := len(nums)
	for i := 0; i < n; i++{
		if zerop == -1 && nums[i] == 0{
			zerop = i
		}else if nums[i] != 0{
			if zerop != -1{
				nums[i], nums[zerop] = nums[zerop], nums[i]
				zerop++
				if zerop < n && nums[zerop] != 0{
					zerop = -1
				}
			}
		}
	}
}
/* 参考官方题解：
** left: 左边均为非零数
** i ： i的左边直到left指针处均为 0
 */
func moveZeroes2(nums []int)  {
	left, n := 0, len(nums)
	for i := 0; i < n; i++{
		if nums[i] != 0{
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}
}

/* 556. Next Greater Element III
** Given a positive integer n, find the smallest integer
** which has exactly the same digits existing in the integer n and is greater in value than n.
** If no such positive integer exists, return -1.
** Note that the returned integer should fit in 32-bit integer,
** if there is a valid answer but it does not fit in 32-bit integer, return -1.
 */
// 2022-02-16 刷出此题，枚举排列，接近超时
func nextGreaterElement(n int) int {
	nums := []int{}
	for t := n; t > 0; t /= 10{
		nums = append(nums, t%10)
	}
	toInt := func(arr []int)int{
		t := 0
		for i := range arr{
			t = t*10 + arr[i]
		}
		if t > math.MaxInt32{
			return -1
		}
		return t
	}
	ans := -1 // golang int 占字节位数与cpu 有关，64位即为 int64
	flag := make([]bool, len(nums))
	var dfs func(result []int)
	dfs = func(result []int){
		if len(result) >= len(nums){
			t := toInt(result)
			if t > n && (ans == -1 || ans > t){
				ans = t
			}
			return
		}
		for i := range flag{
			if !flag[i]{
				flag[i] = true
				dfs(append(result, nums[i]))
				flag[i] = false
			}
		}
	}
	dfs([]int{})
	return ans
}
/* 双指针
** 首先我们观察到任意降序的序列，不会有更大的排列出现。例如 9 5 4 3 1
** 我们需要从右往左找到第一对连续的数字 a[i] 和 a[i-1] 满足 a[i-1] < a[i]，此时 a[i-1]右边的数字无法产生一个更大的排列
** 因为右边的数字时降序的，我们需要重新排布 a[i-1]至最右边的数字来得到下一个排列。
** 那么怎样排布能得到下一个更大的数字呢？
** 我们想得到恰好大于当前数字的下一个排列，所以我们需要用「恰好大于」 a[i-1] 的数字去替换 a[i-1], 设为 a[j]
** 将 a[i-1] 与 a[j] 交换，我们现在在下标为 i-1 的地方得到了正确的数字，但当前的结果还不是一个正确的排列
** 我们需要用从 i-1 开始到最右边数字剩下来的数字升序排列，来得到它们中的最小排列。
** 我们注意到在从右往左找到第一对 a[i-1] < a[i]的连续数字前， a[i-1]右边的数字均是降序排列的，
** 所以交换 a[i-1] 和 a[j] 不会改变下标从 i 开始到最后的顺序。所以我们交换了a[i-1] 和 a[j] 以后，
** 只需要反转下标 i 开始到最后的数字，就可以得到下一个字典序最小的排列。
** 例如： [1, 5, 8, 4, 7, 6, 5, 3, 1]
** 找到 a[i-1] = 4, a[i] = 7
** 从 i 开始向后寻找最下大于a[i-1]的值 a[j] = 5
** 然后交换 a[i-1] 与 a[j] 得到序列 [1, 5, 8, 5, 7, 6, 4, 3, 1] 确认：交换 a[i-1] 和 a[j] 不会改变下标从 i 开始到最后的顺序
** 从 i 开始反转 得到序列 [1, 5, 8, 5, 1, 3, 4, 6, 7] 构造最小的 大于 序列 的值
 */
func nextGreaterElement2(n int) int {
	reverse := func(arr []int){
		for i := 0; i < len(arr)/2; i++{
			arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
		}
	}
	toInt := func(arr []int)int{
		t := 0
		for i := range arr{
			t = t*10 + arr[i]
		}
		if t > math.MaxInt32{
			return -1
		}
		return t
	}
	nums := []int{}
	for t := n; t > 0; t /= 10{
		nums = append([]int{t%10}, nums...)
	}
	t := math.MinInt32
	i := 0
	// 测试case： n=11
	for i = len(nums)-1; i >= 0; i--{
		if nums[i] >= t{// 注意等号
			t = nums[i]
		}else{
			break
		}
	}
	if i < 0{
		return -1
	}
	j := 0
	for j = i+1; j < len(nums); j++{
		if nums[j] <= nums[i]{// 注意等号：12443322
			j -= 1
			break
		}
	}
	if j >= len(nums){
		j -= 1
	}
	nums[i], nums[j] = nums[j], nums[i]
	reverse(nums[i+1:])
	return toInt(nums)
}