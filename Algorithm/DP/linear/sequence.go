package linear

import (
	"math"
	"sort"
)

// 本层 属于 序列DP
/* 题目列表：
** 1.
** 2.
** 3.
** 4.
** 5.
 */

/* 以下部分是 经典题目系列
** Longest Common SubString -- 最长公共子串
** Longest Common Subsequence -- 最长公共子序列
 */

/* 1143. Longest Common Subsequence
** Given two strings text1 and text2, return the length of their longest common subsequence.
** If there is no common subsequence, return 0.
** A subsequence of a string is a new string generated from the original string with some characters (can be none)
** deleted without changing the relative order of the remaining characters.
** For example, "ace" is a subsequence of "abcde".
** A common subsequence of two strings is a subsequence that is common to both strings.
** 变种题目：
** 	712. Minimum ASCII Delete Sum for Two Strings
**  1035. Uncrossed Lines
 */
// dp[i][j] 表示 text1[0:i] 和 text2[0:j] 的最长公共子序列的长度
func LongestCommonSubsequence(text1 string, text2 string) int {
	n1, n2 := len(text1), len(text2)
	dp := make([][]int, n1+1)
	for i := range dp{
		dp[i] = make([]int, n2+1)
	}
	for i := 1; i <= n1; i++{
		for j := 1; j <= n2; j++{
			if text1[i-1] == text2[j-1]{
				dp[i][j] = dp[i-1][j-1] + 1
			}else{
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n1][n2]
}
// 不设置 边界墙
func LongestCommonSubsequence2(text1 string, text2 string) int {
	n1, n2 := len(text1), len(text2)
	dp := make([][]int, n1)
	for i := range dp{
		dp[i] = make([]int, n2)
	}
	for i := range text1{
		for j := range text2{
			if text1[i] == text2[j]{
				if i > 0 && j > 0{
					dp[i][j] = dp[i-1][j-1] + 1
					continue
				}
				if i == 0 || j == 0{
					dp[i][j] = 1
				}
			}else{
				if i > 0 && j > 0{
					dp[i][j] = max(dp[i-1][j], dp[i][j-1])
				}else if i > 0{
					dp[i][j] = dp[i-1][j]
				}else if j > 0{
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return dp[n1-1][n2-1]
}

/* 5. Longest Palindromic Substring
** Given a string s, return the longest palindromic substring in s.
 */
/* 暴力求出所有子串，然后逐个判断
 */
func longestPalindrome(s string) string{
	var isPalindrome func([]byte) bool
	isPalindrome = func(ss []byte) bool {
		for i, j := 0, len(ss)-1; i < j; i,j = i+1, j-1{
			if ss[i] != ss[j]{
				return false
			}
		}
		return true
	}
	length := len(s)
	ans := []byte{}
	for i := 0; i < length; i++{
		for j := i; j < length; j++{
			if isPalindrome([]byte(s[i:j+1])){
				if len(ans) < (j-i+1){
					ans = []byte(s[i:j+1])
				}
			}
		}
	}
	return string(ans)
}
/* 把原来的字符串倒置了，然后找他们俩的最长的公共子串就可以
** 题目转换为 求最长公共子串问题
** 定义dp[i][j]为公共子串的长度
** dp[i][j] = dp[i-1][j-1] + 1
 */
func longestPalindromeDP(s string) string{

}

func longestPalindromeDP2(s string) string{
	n := len(s)
	if n < 2{
		return s
	}
	maxLen, begin := 1, 0
	// dp[i][j]表示s[i:j+1]是否为回文串
	dp := make([][]bool, n)
	for i := range dp{
		dp[i] = make([]bool, n)
		// 初始化: 所有长度为1的子串都是回文串
		dp[i][i] = true
	}
	// 子串长度 从小 到 大 开始递推
	// 先枚举子串长度
	for l := 2; l <= n; l++{
		// 枚举左边界，左边界的上限设置可以宽松些
		for i := 0; i < n; i++{
			// 由于 l 和 i 可以确定右边界， 即 j-i+1得
			j := i + l - 1
			//若右边界越界，退出当前循环
			if j >= n{
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
			}else{
				if j - i < 3{ // 特殊情况
					dp[i][j] = true
				}else{
					dp[i][j] = dp[i+1][j-1]
				}
			}
			// 只要dp[i][l]=true 成立，就表示s[i:l+1]是回文，此时记录回文的长度和起始索引
			if dp[i][j] && j - i + 1 > maxLen{
				maxLen = j-i+1
				begin = i
			}
		}
	}
	return s[begin:begin+maxLen]
}

/* 516. Longest Palindromic Subsequence
** Given a string s, find the longest palindromic subsequence's length in s.
** A subsequence is a sequence that can be derived from another sequence by deleting some or no elements without changing the order of the remaining elements.
 */
/* 基本规律：
** 对于一个子序列而言，如果它是回文子序列，并且长度大于 2，那么将它首尾的两个字符去除之后，它仍然是个回文子序列。
** 因此可以用动态规划的方法计算给定字符串的最长回文子序列
** dp[i][j] 表示字符串 s 的下标范围 [i,j] 内的最长回文子序列的长度
** 边界：
** 1. 任何长度为 1 的子序列都是回文子序列： dp[i][i] = 1
** 2. 0 <= i <= j < n, 非此条件下 dp[i][j] = 0
** i < j 情况：
** 1. s[i] = s[j]
	则首先得到 s 的下标范围 [i+1,j−1] 内的最长回文子序列，然后在该子序列的首尾分别添加 s[i] 和 s[j]，即可得到 s 的下标范围 [i,j] 内的最长回文子序列，
	因此 dp[i][j]=dp[i+1][j−1]+2；
** 2. s[i] != s[j]
	则 s[i] 和 s[j] 不可能同时作为同一个回文子序列的首尾，因此 dp[i][j]=max(dp[i+1][j],dp[i][j−1])。
	其实此处dp[i][j]=max(dp[i+1][j],dp[i][j−1], dp[i+1][j-1]), 但是dp[i+1][j-1] 状态合并掉了
** 由于状态转移方程都是从长度较短的子序列向长度较长的子序列转移，因此需要注意动态规划的循环顺序。
** 最终得到 dp[0][n−1] 即为字符串 ss 的最长回文子序列的长度
*/

func LongestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp{
		dp[i] = make([]int, n)
	}
	for i := n-1; i >= 0; i--{
		dp[i][i] = 1
		for j := i+1; j < n; j++{
			if s[i] == s[j]{
				dp[i][j] = dp[i+1][j-1] + 2
			}else{
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}
/* 归类为区间DP
** 之所以可以使用区间 DP 进行求解，是因为在给定一个回文串的基础上，如果在回文串的边缘分别添加两个新的字符，可以通过判断两字符是否相等来得知新串是否回文
** 使用小区间的回文状态可以推导出大区间的回文状态值
** 从图论意义出发就是，任何一个长度为 len 的回文串，必然由「长度为 len−1」或「长度为 len−2」的回文串转移而来。
** 两个具有公共回文部分的回文串之间存在拓扑序（存在由「长度较小」回文串指向「长度较大」回文串的有向边）。
** 通常区间 DP 问题都是，常见的基本流程为：
** 1. 从小到大枚举区间大小 len
** 2. 枚举区间左端点 l，同时根据区间大小 len 和左端点计算出区间右端点 r = l + len - 1
** 通过状态转移方程求 f[l][r] 的值
 */
func LongestPalindromeSubseqDP(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp{
		dp[i] = make([]int, n)
	}
	for len := 1; len <= n; len++{
		for l := 0; l+len-1 < n; l++{
			r := l+len-1
			if len == 1{
				dp[l][r] = 1
			}else if len == 2{
				if s[l] == s[r]{
					dp[l][r] = 2
				}else {
					dp[l][r] = 1
				}
			}else{
				if s[l] == s[r]{
					dp[l][r] = max(dp[l][r], dp[l+1][r], dp[l][r-1], dp[l+1][r-1]+2)
				}else{
					dp[l][r] = max(dp[l][r], dp[l+1][r], dp[l][r-1], dp[l+1][r-1])
				}
			}
		}
	}
	return dp[0][n-1]
}

/* 300. Longest Increasing Subsequence
** Given an integer array nums, return the length of the longest strictly increasing subsequence.
** A subsequence is a sequence that can be derived from an array by deleting some or no elements without changing the order of the remaining elements.
** For example, [3,6,2,7] is a subsequence of the array [0,3,1,6,2,2,7].
 */
// 2021-12-02 重刷此题，思路还是：dp 定义 dp[i] 为考虑前 i 个元素，以第 i 个数字结尾的最长上升子序列的长度，注意 nums[i] 必须被选取
func LengthOfLIS(nums []int) int {
	// dp[i] = max{dp[j]+1} j <= i and nums[j] < nums[i]
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	ans := 1
	for i := 1; i < n; i++{
		for j := 0; j < i; j++{
			if nums[i] > nums[j]{// strictly increasing
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] += 1
		ans = max(ans, dp[i])
	}
	return ans
}
/* 方法二： 贪心+ 二分 降低时间复杂度
** 如果我们要使上升子序列尽可能的长，则我们需要让序列上升得尽可能慢，因此我们希望每次在上升子序列最后加上的那个数尽可能的小
** 基于上面的贪心思路，我们维护一个数组 d[i] ，表示长度为 i 的最长上升子序列的末尾元素的最小值，用 len 记录目前最长上升子序列的长度，
** 起始时 len 为 1，d[1]=nums[0]
** d[i] 是关于 i 单调递增的。因为如果d[j] >= d[i] 且 j < i
** 算法：
** 依次遍历数组nums 中的每个元素，并更新数组 d 和 len 的值。如果 nums[i] > d[len] 则更新len = len+1
** 否则 在 d[1...len]中找满足 d[i-1] < nums[j] < d[i] 的下标 i， 并更新d[i] = nums[j]
** 根据 d 数组的单调性， 可以使用二分查找寻找下标 i
**
*/
func LengthOfLIS_BS(nums []int) int {
	n := len(nums)
	d := []int{nums[0]}
	ans := 1
	for i := 1; i < n; i++{
		if nums[i] > d[len(d)-1]{
			ans++
			d = append(d, nums[i])
		}else{// 二分查找 找 第一个比nums[i]小的数d[k],并更新d[k+1]=nums[i]
			first := func(t int) bool {
				// return nums[t] >= nums[i] 低级❌
				return d[t] >= nums[i]
			}
			insert := sort.Search(len(d), first)
			d[insert] = nums[i]
		}
	}
	return ans
}
/* 673. Number of Longest Increasing Subsequence
** Given an integer array nums, return the number of longest increasing subsequences.
** Notice that the sequence has to be strictly increasing.
*/
/* DP
** 状态定义：
** 	dp[i]：  以第 i 个数字结尾的最长上升子序列的长度
**  cnt[i]: 以第 i 个数字结尾的最长上升子序列的个数  注意定义！！！
** 设 nums 的最长上升子序列的长度为 maxLen，那么答案为所有满足dp[i]=maxLen 的 i 所对应的 cnt[i] 之和
** 对于 cnt[i]，其等于所有满足 dp[j]+1=dp[i] 的 cnt[j] 之和。在代码实现时，我们可以在计算 dp[i] 的同时统计 cnt[i] 的值
** 由于每个数都能独自一个成为子序列，因此起始必然有cnt[i] = 1
** 枚举[0,i) 的所有数 nums[j] 如果满足 nums[j] < nums[i], 说明nums[i] 可以接在 nums[j] 后面形成上升子序列。
** 这时 需要对 dp[i] 和 dp[j] + 1 的大小分情况讨论： <== 就是这个地方 没有思考出
** 1. 满足 dp[i] < dp[j] + 1: 说明 dp[i] 会被 dp[j]+1 直接更新， 此时直接同步更新cnt[i] = cnt[j] 即可
** 2. 满足 dp[i] == dp[j] + 1：说明找到了一个新的符合条件的前驱，此时将值累加到方案数中，即 cnt[i] += cnt[j]
 */
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	cnt := make([]int, n)
	ans := 0
	maxLen := 0
	for i := 0; i < n; i++{
		dp[i] = 1
		cnt[i] = 1
		for j := 0; j < i; j++{
			if nums[i] > nums[j]{
				if dp[j] + 1 > dp[i]{
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j]
				}else if dp[j] + 1 == dp[i]{
					cnt[i] += cnt[j] // 情况合并
				}
			}
		}
		if maxLen < dp[i]{
			maxLen = dp[i]
			ans = cnt[i]
		}else if maxLen == dp[i]{
			ans += cnt[i]
		}
	}
	return ans
}
/* 方法二： 贪心 前缀和  二分
** 继上题贪心思路， 将数组d 扩展为一个二维数组，其中d[i]数组表示所有能成为长度为 i 的最长上升子序列的末尾元素的值
** 即，将更新d[i]=nums[j]这一操作 替换成 将 nums[j] 置于 d[i]数组末尾。这样d[i]中就保留了历史信息，且 d[i]中的元素是单调非递增
** 同样，定义 二维数组 cnt， 其中 cnt[i][j]记录了以d[i][j]为结尾的最长上升子序列的个数。
** 为了计算cnt[i][j] 考察d[i-1] 和 cnt[i-1], 将所有满足 d[i-1][k] < d[i][j]的 cnt[i-1][k] 累加到 cnt[i][j]，
** 这样最终答案就是cnt[maxLen]的所有元素之和
*/
func findNumberOfLIS_BS(nums []int) int {
	d := [][]int{}
	cnt := [][]int{}
	for _, v := range nums{
		i := sort.Search(len(d), func(i int)bool{ return d[i][len(d[i])-1] >= v })
		c := 1
		if i > 0{ // 为了计算cnt[i][j]
			k := sort.Search(len(d[i-1]), func(k int)bool{ return d[i-1][k] < v })
			c = cnt[i-1][len(cnt[i-1])-1] - cnt[i-1][k]
		}
		if i == len(d){
			d = append(d, []int{v})
			cnt = append(cnt, []int{0, c}) // 前缀0： 方便前缀和 优化
		}else{
			d[i] = append(d[i], v)
			cnt[i] = append(cnt[i], cnt[i][len(cnt[i])-1] + c)
		}
	}
	c := cnt[len(cnt) - 1]
	return c[len(c) - 1]
}

/* 376. Wiggle Subsequence
** A wiggle sequence is a sequence where the differences between successive numbers strictly alternate between positive and negative.
** The first difference (if one exists) may be either positive or negative.
** A sequence with one element and a sequence with two non-equal elements are trivially wiggle sequences.
** For example,
** [1, 7, 4, 9, 2, 5] is a wiggle sequence because the differences (6, -3, 5, -7, 3) alternate between positive and negative.
** In contrast, [1, 4, 7, 2, 5] and [1, 7, 4, 5, 5] are not wiggle sequences.
** The first is not because its first two differences are positive, and the second is not because its last difference is zero.
** A subsequence is obtained by deleting some elements (possibly zero) from the original sequence,
** leaving the remaining elements in their original order.
** Given an integer array nums, return the length of the longest wiggle subsequence of nums.
 */
/*  定义：难度在 这个「上升摆动序列」 和 「下降摆动序列」的定义上
	1. 某个序列被称为「上升摆动序列」，当且仅当该序列是摆动序列，且最后一个元素呈上升趋势。如序列 [1,3,2,4] 即为「上升摆动序列」
	2. 某个序列被称为「下降摆动序列」，当且仅当该序列是摆动序列，且最后一个元素呈下降趋势。如序列 [4,2,3,1] 即为「下降摆动序列」
每当我们选择一个元素作为摆动序列的一部分时，这个元素要么是上升的，要么是下降的，这取决于前一个元素的大小。
** 那么列出状态表达式为
** 1. up[i]: 表示以前 i 个元素中的某一个为结尾的最长的「上升摆动序列」的长度
** 2. down[i]: 表示以前 i 个元素中的某一个为结尾的最长的「下降摆动序列」的长度
** 3. 只有一个数字时默认为 1
** 状态转移规则：
** 1. 当 nums[i] <= nums[i-1]时， 我们无法选出更长的「上升摆动序列」的方案，因为对于任何以 nums[i] 结尾的「上升摆动序列」，
		我们都可以将nums[i] 替换为 nums[i−1]，使其成为以nums[i−1] 结尾的「上升摆动序列」
** 2. 当 nums[i] > nums[i-1]时，我们既可以从 up[i-1] 进行转移， 也可从 down[i-1]转移。
** 状态方程：
	up[i] 	= up[i-1] 						if nums[i] <= nums[i-1]
		  	= max(up[i-1], down[i-1] + 1) 	if nums[i] > nums[i-1]
	down[i] = down[i-1]						if nums[i] >= nums[i-1]
			= max(up[i-1] + 1, down[i-1])	if nums[i] < nums[i-1]
*/
/* 求解的过程类似最长上升子序列, 不过是需要判断两个序列， 分情况讨论
** 1. nums[i] > nums[i-1]
	1.1 假设 down[i-1] 表示的最长摆动序列的最远末尾元素下标正好为 i-1, 遇到新的上升元素后, up[i] = down[i-1] + 1
	1.2 假设 down[i-1] 表示的最长摆动序列的最远末尾元素下标小于 i，设为 j, 那么 nums[j:i] 一定是递增的，
		因为若完全递减，最远元素下标等于 i，若波动，那么down[i] > down[j]。
		由于 nums[j:i] 递增，down[j:i] 一直等于 down[j]， 依然满足 up[i] = down[i-1] + 1
** 2. nums[i] < nums[i-1]
** 3. nums[i] = nums[i-1]
	新的元素不能用于任何序列，保持不变
 */
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	up, down := 1, 1
	for i := 1; i < n; i++{
		tup, tdown := 0, 0
		if nums[i] > nums[i-1]{
			tup = max(up, down+1)
			tdown = down
		}else if nums[i] < nums[i-1]{
			tdown = max(up+1, down)
			tup = up
		}else{
			tup = up
			tdown = down
		}
		up = tup
		down = tdown
	}
	return max(up, down)
}

/* 如果是求 subarray 的话，即连续的子序列 */
func wiggleMaxLength_subArray(nums []int) int {
	dp := make([]int, 2)
	ans := 1
	n := len(nums)
	dp[0], dp[1] = 1, 1
	for i := 1; i < n; i++{
		d := nums[i] - nums[i-1]
		if d > 0{
			dp[0] = dp[1] + 1
			dp[1] = 1
			ans = max(dp[0], ans)
		}else{
			dp[1] = dp[0] + 1
			dp[0] = 1
			ans = max(ans, dp[1])
		}
	}
	return ans
}

/* 1048. Longest String Chain
** You are given an array of words where each word consists of lowercase English letters.
** wordA is a predecessor of wordB if and only if we can insert exactly one letter anywhere in wordA without changing the order of the other characters to make it equal to wordB.
** For example, "abc" is a predecessor of "abac", while "cba" is not a predecessor of "bcad".
** A word chain is a sequence of words [word1, word2, ..., wordk] with k >= 1, where word1 is a predecessor of word2, word2 is a predecessor of word3, and so on.
** A single word is trivially a word chain with k == 1.
** Return the length of the longest possible word chain with words chosen from the given list of words.
*/
// 2021-12-03 刷出此题
//动态规划解决：
// 遗漏的失误点有 2个：
// 1. 刚开始 未意识到 要排序， 原因未仔细体会题目意图，单词可以不按 sequence 顺序。 但是为了借助 sequence DP 解决，必须先排好序
// 2. predecessor的计算 没有理解清除
// 另外，还要注意 sort 包的使用
func LongestStrChain(words []string) int {
	sort.Slice(words, func(p, q int)bool{ return len(words[p]) < len(words[q]) })
	ispre := func(a string, b string)bool{
		na, nb := len(a), len(b)
		if na != nb + 1{
			return false
		}
		/* 优化掉 cnt
		cnt := 0
		for i,j := 0,0; i < na && j < nb;{
			if a[i] != b[j]{
				cnt++
				i++
			}else{
				i++
				j++
			}
			if cnt > 1{ // 必须放后面，考虑case：["a","b","ab","bac"]
				return false
			}
		}
		 */
		// i 的值 是比 j 大的
		for i, j := 0, 0; i < na && j < nb;i++{
			if a[i] == b[j]{
				j++
			}
			if i-j >= 2{
				return false
			}
		}
		return true
	}
	n := len(words)
	dp := make([]int, n)
	ans := 1
	dp[0] = 1
	// dp[i] = max(dp[j]) + 1
	for i := 1; i < n; i++{
		dp[i] = 1
		for j := 0; j < i; j++{
			if ispre(words[i], words[j]){
				if dp[i] < dp[j] + 1{
					dp[i] = dp[j] + 1
				}
			}
		}
		if ans < dp[i]{
			ans = dp[i]
		}
	}
	return ans
}
// 方法二： 直接DP，排除 使用 ispre 函数,  效率比方法一差些
//可以学习这种思考方式， 使用map 以及 采用扣掉一个字符 穷举可能的情况
func LongestStrChain2(words []string) int {
	ans, dp := 0, map[string]int{}
	sort.Slice(words, func(i, j int)bool{ return len(words[i]) < len(words[j]) } )
	for i := range words{
		w := words[i]
		for j := range words[i]{
			// 采用直接DP 方式, 穷举扣掉一个字符的情况
			cnt := dp[w[:j]+w[j+1:]] + 1
			if cnt > dp[w]{
				dp[w] = cnt
			}
		}
		if dp[w] > ans{
			ans = dp[w]
		}
	}
	return ans
}
// 方法三： 图算法应用-- 图的DFS
// 思路：
//	1. 将 words按照长度进行分组
//	2. 遍历长度相差 1 的 2组之间是否有predecessor关系，若有，则从前身到单词之间有条路径，将路径加入到有向图中。
//	3. 遍历图，找到最长路径
func LongestStrChain_graph(words []string) int {
	// m := map[int][]string{} 必须存索引值，构造图使用索引 表示每个图节点
	m := map[int][]int{}
	minlen, maxlen := math.MaxInt32, 0
	for i := range words{
		n := len(words[i])
		m[n] = append(m[n], i)
		minlen = min(minlen, n)
		maxlen = max(maxlen, n)
	}
	ispre := func(a, b string)bool{
		na, nb := len(a), len(b)
		if na != nb + 1{
			return false
		}
		ia, ib := 0, 0
		for ia < na && ib < nb{
			if a[ia] == b[ib]{
				ib++
			}
			ia++
			if ia - ib >= 2{
				return false
			}
		}
		if ia - ib >= 2{ // 尾部判断：比较两数组时候务必检查
			return false
		}
		return true
	}
	// 构造图 邻接表
	graph := make([][]int, len(words)) // 每个word 对应一个图节点
	for length := minlen; length <= maxlen; length++{
		shortGroup := m[length]
		longGroup := m[length+1]
		for sk := range shortGroup{
			for lk := range longGroup{
				if ispre(words[longGroup[lk]], words[shortGroup[sk]]){
					//graph[sk] = append(graph[sk], longGroup[lk]) 糊涂了
					graph[shortGroup[sk]] = append(graph[shortGroup[sk]], longGroup[lk])
				}
			}
		}
	}
	// 遍历有向图，dfs找到最长路径
	maxLen := make([]int, len(words)) // 类似于visited 矩阵
	var dfs func(node int)int
	dfs = func(node int)int{
		if maxLen[node] != 0{
			return maxLen[node]
		}
		maxSubLen := 0
		for nextNode := range graph[node]{
			maxSubLen = max(maxSubLen, dfs(graph[node][nextNode]))
			//maxSubLen = max(maxSubLen, dfs(nextNode)) 糊涂了
		}
		maxLen[node] = maxSubLen + 1
		return maxLen[node]
	}
	ans := 0
	for i := 0; i < len(words); i++{
		ans = max(ans, dfs(i))
	}
	return ans
}

/* 1035. Uncrossed Lines
** You are given two integer arrays nums1 and nums2. We write the integers of nums1 and nums2 (in the order they are given) on two separate horizontal lines.
We may draw connecting lines: a straight line connecting two numbers nums1[i] and nums2[j] such that:
nums1[i] == nums2[j], and
the line we draw does not intersect any other connecting (non-horizontal) line.
Note that a connecting line cannot intersect even at the endpoints (i.e., each number can only belong to one connecting line).
Return the maximum number of connecting lines we can draw in this way.
*/
/*lcs 问题  悟性差呀呀呀*/
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	dp := make([][]int, n1+1)
	for i := range dp{
		dp[i] = make([]int, n2+1)
	}
	for i := 1; i <= n1; i++{
		for j := 1; j <= n2; j++{
			if nums1[i-1] == nums2[j-1]{
				dp[i][j] = dp[i-1][j-1]+1
			}else{
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n1][n2]
}