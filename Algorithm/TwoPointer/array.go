package TwoPointer

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
