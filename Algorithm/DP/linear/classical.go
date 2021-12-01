package linear

import "math"

/* 486. Predict the Winner
You are given an integer array nums. Two players are playing a game with this array: player 1 and player 2.
Player 1 and player 2 take turns, with player 1 starting first. Both players start the game with a score of 0.
At each turn, the player takes one of the numbers from either end of the array (i.e., nums[0] or nums[nums.length - 1])
which reduces the size of the array by 1.
The player adds the chosen number to their score. The game ends when there are no more elements in the array.
Return true if Player 1 can win the game. If the scores of both players are equal, then player 1 is still the winner, and you should also return true. You may assume that both players are playing optimally.
 */
func PredictTheWinnerDFS(nums []int) bool {
	var dfs func(arr []int)int // 当前选择的玩家所能赢对方的分数
	// 返回当前做选择的玩家，基于当前区间[i,j]，赢过对手的分数
	dfs = func(arr []int)int{
		n := len(arr)
		if n == 1{ //此时只有一种选择，选的人赢对方arr[0]，且没有剩余可选
			return arr[0]
		}
		chooseHead := arr[0] - dfs(arr[1:]) // 选择首端，获得arr[0]，之后输掉dfs(i+1,j)分
		chooseTail := arr[n-1] - dfs(arr[:n-1])
		return max(chooseHead, chooseTail)
	}
	return dfs(nums) >= 0 // 基于整个数组玩这个游戏，玩家1先手，>=0就获胜
}
/* 877. Stone Game
** Alice and Bob play a game with piles of stones.
** There are an even（偶数个）number of piles arranged in a row(排成一行),
** and each pile has a positive integer number(正整数) of stones piles[i].
** The objective of the game is to end with the most stones.
** The total number(石头总数) of stones across all the piles is odd, so there are no ties(平局).
** Alice and Bob take turns, with Alice starting first.
** Each turn, a player takes the entire pile of stones either from the beginning or from the end of the row.
** This continues until there are no more piles left, at which point the person with the most stones wins.
** Assuming Alice and Bob play optimally(Alice存在赢的可能就返回true), return true if Alice wins the game, or false if Bob wins.
*/
// 尝试DFS 暴力枚举
/* 题目卡壳的点 在于： 每个节点都是其中一个玩家在选择，下一个节点变成对手在选择，交替在选。
  dfs 设置返回值需要是差值
  返回当前做选择的玩家，基于当前区间[i,j]，赢过对手的分数。
  当前选择的分数，减去，往后对手赢过自己的分数（对剩余数组递归）。因为有两端可选择，所以差值有两个，取较大的判断是否 >= 0
 */
func StoneGame(piles []int) (ans bool) {
	n := len(piles)
	const Alice, Bob int = 0, 1
	cache := [2][][]int{}
	for i := range cache{
		cache[i] = make([][]int, n)
		for j := range cache[i]{
			cache[i][j] = make([]int, n)
		}
	}
	// dfs 返回差值--> alice的分数 - bob 的分数
	var dfs func(head int, tail int, player int)int
	dfs = func(head, tail, player int) int {
		if head > tail{
			return 0
		}
		if cache[player][head][tail] != 0{
			return cache[player][head][tail]
		}
		t := 0
		if head == tail{
			if player == Alice{
				t =  dfs(head+1, tail-1, Bob) + piles[head]
			}else{
				t =  dfs(head+1, tail-1, Bob) - piles[head]
			}
			cache[player][head][tail] = t
			return t
		}
		if player == Alice{
			t = max(dfs(head+1, tail, Bob)+piles[head], dfs(head, tail-1, Bob)+piles[tail])

		}else{
			t = max(dfs(head+1, tail, Alice)-piles[head], dfs(head, tail-1, Alice)-piles[tail])
		}
		cache[player][head][tail] = t
		return t
	}
	// Alice: 0   Bob: 1
	if dfs(0, n-1, Alice) > 0{
		return true
	}
	return false
}
/*属于区间DP：
** 定义dp[i][j]为考虑区间[i,j]在双方都做最好选择的情况下，先手与后手的最大得分差值为多少
** dp[0][n-1]为考虑所有情况，先手和后手得分差值分为2中情况：
** dp[0][n-1] > 0, 先手必胜返回true
** dp[0][n-1] < 0, 先手必败返回false
** 从两端取值，分2种情况：
** 1. 本次先手从左端点取石子的话，双方差值为： piles[i] - dp[i+1][j]
** 2. 本次先手从右端点取石子的话，双方差值为:  piles[j] - dp[i][j-1]
** 双方都想赢，都会做最优决策（即使自己与对方分差最大) 故 dp[i][j] 为上述情况的最大值
** 根据状态转移方程，我们发现大区间的状态值依赖于小区间的状态值，典型的区间 DP 问题。
** 按照从小到大「枚举区间长度」和「区间左端点」的常规做法进行求解即可。
*/
func StoneGameDP(piles []int) (ans bool) {
	n := len(piles)
	dp := make([][]int, n+1)
	for i := range dp{
		dp[i] = make([]int, n+1)
	}
	// 额外初始化 区间长度为1的情况
	for i := 1; i <= n; i++{ // 枚举区间的长度，区间从1 至 n
		for l := 0; l + i - 1 < n; l++{ // 枚举左端点，注意，l + i - 1 为右端点的索引值
			r := l + i - 1 // 右端点的索引值
			chooseLeft := piles[l] - dp[l+1][r] // 选择左端点
			chooseRight := math.MinInt32
			if r > 0{
				chooseRight = piles[r] - dp[l][r-1] // 选择右端点, 需要注意 区间长度为 1 的情况，此时右端点为 -1 只能选一次
			}
			dp[l][r] = max(chooseLeft, chooseRight)
		}
	}
	return dp[0][n-1] > 0
}
func StoneGameDP0(piles []int) (ans bool) {
	n := len(piles)
	dp := make([][]int, n+1)
	for i := range dp{
		dp[i] = make([]int, n+1)
	}
	// 额外初始化 区间长度为1的情况
	dp[0][0] = piles[0]
	// 显然从 区间长度为2 开始才有意义
	for i := 2; i <= n; i++{ // 枚举区间的长度，区间从1 至 n
		for l := 0; l + i - 1 < n; l++{ // 枚举左端点，注意，l + i - 1 为右端点的索引值
			r := l + i - 1 // 右端点的索引值
			chooseLeft := piles[l] - dp[l+1][r] // 选择左端点
			chooseRight := piles[r] - dp[l][r-1] // 选择右端点, 需要注意 区间长度为 1 的情况，此时右端点为 -1 只能选一次
			dp[l][r] = max(chooseLeft, chooseRight)
		}
	}
	return dp[0][n-1] > 0
}
/* 如果只剩下一堆石子，则当前玩家只能取走这堆石子。
如果剩下多堆石子，则当前玩家可以选择从行的开始或结束处取走整堆石子，然后轮到另一个玩家在剩下的石子堆中取走石子。
这是一个递归的过程，因此可以使用递归进行求解，递归过程中维护一个总数，表示Alex和 Lee 的石子数量之差，
当游戏结束时，如果总数大于 0，则 Alex 赢得比赛，否则 Lee 赢得比赛。
dp[i][j] 表示当剩下的石子堆为下标 i 到下标 j 时,当前玩家与另一个玩家的石子数量之差的最大值，注意当前玩家不一定是先手Alex
显然，i <= j 才有意义。
当 i == j： 当前玩家只能取走最后这堆石子，一次对于所有i 都有 dp[i][j] = piles[i]
当 i < j: 当前玩家可以选择取走piles[i] 或 piles[j]， 然后轮到另一个玩家在剩余的piles中取走石子。在两种方案里，当前玩家取最优的
dp[i][j] = max( piles[i] - dp[i+1][j], piles[j] - dp[i][j-1] )
看参考官网画出 矩阵图
 */
func StoneGameDP2(piles []int) bool {
	length := len(piles)
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
		dp[i][i] = piles[i]
	}
	for i := length - 2; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			dp[i][j] = max(piles[i] - dp[i+1][j], piles[j] - dp[i][j-1])
		}
	}
	return dp[0][length-1] > 0
}

func StoneGameDP1(piles []int) (ans bool) {
	n := len(piles)
	dp := make([]int, n+1)
	copy(dp, piles)
	for i := n - 2; i >= 0; i--{
		for j := i+1; j < n; j++{
			dp[j] = max(piles[i] - dp[j], piles[j] - dp[j-1])
		}
	}
	return dp[n-1] > 0
}
/* 博弈论: 经典的博弈论问题
核心： 先手只需要在进行第一次操作前计算原序列中「奇数总和」和「偶数总和」哪个大，然后每一次决策都「限制」对方只能选择「最优奇偶性序列」的对立面即可。
为了方便，我们称「石子序列」为石子在原排序中的编号，下标从 1 开始。
由于石子的堆数为偶数，且只能从两端取石子。因此先手后手所能选择的石子序列，完全取决于先手每一次决定。
证明如下：
由于石子的堆数为偶数，对于先手而言：每一次的决策局面，都能「自由地」选择奇数还是偶数的序列，从而限制后手下一次「只能」奇数还是偶数石子。
具体的，对于本题，由于石子堆数为偶数，因此先手的最开始局面必然是 [奇数, 偶数][奇数,偶数]，即必然是「奇偶性不同的局面」；
当先手决策完之后，交到给后手的要么是 [奇数,奇数][奇数,奇数] 或者 [偶数,偶数][偶数,偶数]，即必然是「奇偶性相同的局面」；
后手决策完后，又恢复「奇偶性不同的局面」交回到先手 ...
不难归纳推理，这个边界是可以应用到每一个回合。
因此先手只需要在进行第一次操作前计算原序列中「奇数总和」和「偶数总和」哪个大，然后每一次决策都「限制」对方只能选择「最优奇偶性序列」的对立面即可。
同时又由于所有石子总和为奇数，堆数为偶数，即没有平局，所以先手必胜。
 */
func StoneGameMath(piles []int) (ans bool) {
	return true
}

/*122. Best Time to Buy and Sell Stock II
** You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
** On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time.
** However, you can buy it then immediately sell it on the same day.
** Find and return the maximum profit you can achieve.
*/
// 注意初始情况，总要开始选择买入
func maxProfit(prices []int) int {
	// 一天一共3种情况： 1. 买 2. 卖 3. 不操作
	// 买入dp[i][0] = max(dp[i-1][1]-prices[i], dp[i-1][0])  持有一支股票的情况
	// 卖出dp[i][1] = max(dp[i-1][0]+prices[i], dp[i-1][1])  不持有股票的情况
	dp := make([]int, 2)
	dp[0], dp[1] = -prices[0], 0
	n := len(prices)
	for i := 1; i < n; i++{
		t := dp[0]
		dp[0] = max(dp[1]-prices[i], dp[0])
		dp[1] = max(t+prices[i], dp[1])
	}
	return max(dp...)
}
/* 贪心的方法：
** 由于股票的购买没有限制，因此整个问题等价于寻找 x 个不相交的区间 (li, ri] 使得如下的等式最大化 SUM(a[ri]-a[li])
** 上述等价于：
** 		对于(li, ri]这一个区间贡献的价值a[ri]-a[li], 等价于 (li, li+1], (li+1, li+2], ..., (ri-1, ri] 这若干区间长度为1的区间价值和
**		a[ri]-a[li] = a[ri]-a[ri-1] + a[ri-1]-a[ri-2] + ... + a[li+1] - a[li]
**		问题简化为 找 x 个长度为1的区间 (li, li+1] 使得 SUM(a[li+1]-a[li]) 最大 i属于 [1,x]
** 		贪心的角度考虑我们每次选择贡献大于 0 的区间即能使得答案最大化
** 		需要说明的是，贪心算法只能用于计算最大利润，计算的过程并不是实际的交易过程
 */
func maxProfitII(prices []int) (ans int) {
	for i := 1; i < len(prices); i++ {
		ans += max(0, prices[i]-prices[i-1])
	}
	return
}

/* 309. Best Time to Buy and Sell Stock with Cooldown
** You are given an array prices where prices[i] is the price of a given stock on the ith day.
Find the maximum profit you can achieve.
You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times) with the following restrictions:
After you sell your stock, you cannot buy stock on the next day (i.e., cooldown one day).
Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).
*/
func maxProfitIII(prices []int) int {
	n := len(prices)
	dp := make([]int, 3)
	dp[0] = -prices[0]	// 手上持有股票的最大收益
	dp[1] = 0 			// 手上不持有股票，并且处于冷冻期中的累计最大收益
	dp[2] = 0			// 手上不持有股票，并且不在冷冻期中的累计最大收益
	for i := 1; i < n; i++{
		dp0 := dp[0]
		dp1 := dp[1]
		dp[0] = max(dp[0], dp[2]-prices[i])
		dp[1] = dp0 + prices[i]
		dp[2] = max(dp1, dp[2])
	}
	return max(dp[1], dp[2])
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
/* 每当我们选择一个元素作为摆动序列的一部分时，这个元素要么是上升的，要么是下降的，这取决于前一个元素的大小。
** 那么列出状态表达式为
** 1. up[i]: 表示以前 i 个元素中的某一个为结尾的最长的「上升摆动序列」的长度
** 2. down[i]: 表示以前 i 个元素中的某一个为结尾的最长的「下降摆动序列」的长度
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
func wiggleMaxLength(nums []int) int {

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
 */
func LongestCommonSubsequence(text1 string, text2 string) int {

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







