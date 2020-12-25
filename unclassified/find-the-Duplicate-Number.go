/*287. 寻找重复数
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

import (
	"sort"
)

func FindDuplicate(nums []int) int {
	//return binarySearchFind(nums)
	return findDuplicate_bit(nums)
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

/* 方法一：
Space complexity : O(1) or O(n)：如果可以用原本的数列（sort in place），即为O(1)，但是因为有第一条限制，所以这里是O(n)（其实违背了题目要求）。
Time complexity : O(nlog(n))：因为sort
链接：https://leetcode-cn.com/problems/find-the-duplicate-number/solution/287-6chong-jie-fa-si-lu-xiang-xi-fen-xi-zong-jie-b/
*/
func findDuplicate_Sort(nums []int) int {
	sort.Sort(sort.IntSlice(nums))
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1]{
			return nums[i]
		}
	}
	return -1
}
/* 方法二 位运算
思路： 利用了Integer有32 Bits来进行对比：
a：在整数1->n里，有多少数的第i个bit是1？
b：在数列nums里，有多少数的第i个bit是1？
如果b > a，那么结果res的第i个bit是1。
这个方法我们来将所有数二进制展开按位考虑如何找出重复的数，如果我们能确定重复数每一位是 11 还是 00 就可以按位还原出重复的数是什么。
考虑到第i位，我们记 \textit{nums}[]nums[] 数组中二进制展开后第 ii 位为 11 的数有 xx 个，
数字 [1,n][1,n] 这 nn 个数二进制展开后第 ii 位为 11 的数有 yy 个，那么重复的数第 ii 位为 11 当且仅当 x>yx>y。
仍然以示例 1 为例，如下的表格列出了每个数字二进制下每一位是 11 还是 00 以及对应位的 xx 和 yy 是多少：
 		1	3	4	2	2	x	y
第 0 位	1	1	0	0	0	2	2
第 1 位	0	1	0	1	1	3	2
第 2 位	0	0	1	0	0	1	1
那么按之前说的我们发现只有第 11 位 x>yx>y ，所以按位还原后 \textit{target}=(010)_2=(2)_{10}target=(010)
2
 =(2)
10
 ，符合答案。
正确性的证明其实和方法一类似，我们可以按方法一的方法，考虑不同示例数组中第 ii 位 11 的个数 xx 的变化：
如果测试用例的数组中 \textit{target}target 出现了两次，其余的数各出现了一次，且 \textit{target}target 的第 ii 位为 11，那么 \textit{nums}[]nums[] 数组中第 ii 位 11 的个数 xx 恰好比 yy 大一。如果\textit{target}target 的第 ii 位为 00，那么两者相等。
如果测试用例的数组中 \textit{target}target 出现了三次及以上，那么必然有一些数不在 \textit{nums}[]nums[] 数组中了，这个时候相当于我们用 \textit{target}target 去替换了这些数，我们考虑替换的时候对 xx 的影响：
如果被替换的数第 ii 位为 11，且 \textit{target}target 第 ii 位为 11：xx 不变，满足 x>yx>y。
如果被替换的数第 ii 位为 00，且 \textit{target}target 第 ii 位为 11：xx 加一，满足 x>yx>y。
如果被替换的数第 ii 位为 11，且 \textit{target}target 第 ii 位为 00：xx 减一，满足 x\le yx≤y。
如果被替换的数第 ii 位为 00，且 \textit{target}target 第 ii 位为 00：xx 不变，满足 x\le yx≤y。
也就是说如果 \textit{target}target 第 ii 位为 11，那么每次替换后只会使 xx 不变或增大，如果为 00，只会使 xx 不变或减小，
始终满足 x>yx>y 时 \textit{target}target 第 ii 位为 11，否则为 00，因此我们只要按位还原这个重复的数即可。
链接：https://leetcode-cn.com/problems/find-the-duplicate-number/solution/xun-zhao-zhong-fu-shu-by-leetcode-solution/
*/
func findDuplicate_bit(nums []int) int{
	n := len(nums) - 1
	mask := 31
	result := 0
	// 确定n的最高为1的bit位，方便后续处理
	for (n >> mask) == 0 {
		mask--
	}
	// 加b<=mask 防止{1,1}情况，此时mask为0
	for b := 0; b <= mask; b++{
		x,y := 0,0
		for i := 0; i <= n; i++ {
			if i & (1 << b) > 0 {
				x++
			}
			if nums[i] &(1 << b) > 0{
				y++
			}
		}
		if y > x{
			result |= (1 << b)
		}
	}
	return result
}

/* 方法三：抽象成环问题，查找环入口
链接：https://leetcode-cn.com/problems/find-the-duplicate-number/solution/287-6chong-jie-fa-si-lu-xiang-xi-fen-xi-zong-jie-b/
 */
func findDuplicate_flydCircle(nums []int) int{
	return 0
}
