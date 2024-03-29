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

/* 238. Product of Array Except Self
** Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
You must write an algorithm that runs in O(n) time and without using the division operation.
Constraints:
	2 <= nums.length <= 105
	-30 <= nums[i] <= 30
	The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
Follow up: 
	Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)
*/
// 前后缀 积 2021-12-08 刷出此题
func productExceptSelf(nums []int) []int {
	n := len(nums)
	dp := make([]int, n+1)
	dp[n] = 1
	for i := n-1; i >= 0; i--{
		dp[i] = dp[i+1]*nums[i]
	}
	prefix := make([]int, n+1)
	prefix[0] = 1
	for i := 1; i <= n; i++{
		prefix[i] = prefix[i-1]*nums[i-1]
	}
	for i := range nums{
		nums[i] = dp[i+1]*prefix[i]
	}
	return nums
}

// 算是O(1) 空间复杂度
func productExceptSelf2(nums []int) []int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++{
		dp[i] = nums[i-1]*dp[i-1]
	}
	r := 1
	for i := n-1; i >= 0; i--{
		dp[i] = dp[i]*r
		r *= nums[i]
	}
	return dp
}







