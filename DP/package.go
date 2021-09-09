package DP

import (
	"math"
	"sort"
)

/*类属组合优化-NP完全问题-无法直接求解-通过穷举+验证求解
泛指一类「给定价值与成本」，同时「限定决策规则」，在这样的条件下，如何实现价值最大化的问题
 */
/* 0 1 背包问题
「01背包」是指给定物品价值与体积（对应了「给定价值与成本」），在规定容量下（对应了「限定决策规则」）如何使得所选物品的总价值最大
示例：
	有N件物品和一个容量是V的背包。每件物品有且只有一件,其中第i件物品的体积是v[i],价值是w[i]
	求解将哪些物品装入背包，可使这些物品的总体积不超过背包容量，且总价值最大.
枚举求解：
	func dfs(i, c) int  <== i 和 c 分别代表「当前枚举到哪件物品」和「现在的剩余容量」， 返回最大价值
DP求解:
	dp数组：二维数组，其中一维代表当前「当前枚举到哪件物品」，另外一维「现在的剩余容量」，数组装的是「最大价值」
	状态定义： 考虑前 i 件物品，使用容量不超过 C 的条件下的背包最大价值
	状态转移方程：只需要考虑第 i 件物品如何选择即可，对于第 i 件物品，我们有「选」和「不选」两种决策
		决策-1：「不选」其实就是 ，等效于我们只考虑前 i-1 件物品，当前容量不超过 C 的情况下的最大价值
		决策-2：「选」第 i 件物品，代表消耗了 v[i]的背包容量，获取了 w[i]的价值，那么留给前 i-1件物品的背包容量只剩 c-v[i]，即 dp[i-1][c-v[i]+w[i]
        在「选」和「不选」之间取最大值，就是我们「考虑前 i 件物品，使用容量不超过「C」的条件下的「背包最大价值」
		dp[i][c] = max(dp[i-1][c], dp[i-1][c-v[i]] + w[i])
	额外条件：剪枝
		选第 i 件有一个前提：「当前剩余的背包容量」>=「物品的体积」
 */
// dp[N][C+1] 解法
func example_1(N int, C int, v []int, w []int)int{
	//构建 dp[N][C+1] 数组
	dp := make([][]int, N)
	for i := range dp{
		dp[i] = make([]int, C+1)
	}
	// 状态初始化：考虑第一件物品的情况
	for i := 0; i <= C; i++{
		if i >= v[0]{
			dp[0][i] = w[0]
		}
	}
	// dp: dp[i][c] = max(dp[i-1][c], dp[i-1][c-v[i]] + w[i])
	for i := 1; i < N; i++{
		for j := 0; j <= C; j++{
			// 额外条件：「当前剩余的背包容量」>=「物品的体积」
			if j >= v[i]{
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-v[i]] + w[i])
			}else{
				dp[i][j] = max(dp[i-1][j], 0)
			}
		}
	}
	return dp[N-1][C]
}
// dp[2][C+1] 解法 滚动数组
/* 只依赖上一层，故可用滚动数组优化空间复杂度
** 小技巧： i % 2 可以写为  i & 1
 */
func example_2(N int, C int, v []int, w []int)int{
	dp := [2][]int{}
	for i := range dp{
		dp[i] = make([]int, C+1)
	}
	for i := 0; i <= C; i++{
		if i >= v[0]{
			dp[0][i] = w[0]
		}
	}
	for i := 1; i < N; i++{
		for j := 0; j <= C; j++{
			if j >= v[i]{
				dp[i&1][j] = max(dp[(i-1)&1][j], dp[i-1][j-v[i]] + w[i])
			}else{
				dp[i&1][j] = max(dp[(i-1)&1][j], 0)
			}
		}
	}
	return dp[(N-1)&1][C]
}
// dp[C+1] 解法  降维度
/* 继续分析 我们的 转移方程，dp[i][c] = max(dp[i-1][c], dp[i-1][c-v[i]] + w[i])
  不难发现，当求解第 i 行格子的值时，不仅是只依赖第i-1行，还明确只依赖第i-1行的第c个格子和第c-v[i]个格子（也就是对应着第 i 个物品不选和选的两种情况）
  只依赖于「上一个格子的位置」以及「上一个格子的左边位置」
  因此，只要我们将求解第 i 行格子的顺序「从 0 到 c 」改为「从 c 到 0 」，就可以将原本两行的二维数组压缩到一行（转换为一维数组）
 */
func example_3(N int, C int, v []int, w []int)int{
	dp := make([]int, C+1)
	for i := 0; i < N; i++{
		for j := C; j >= v[i]; j--{
			no := dp[j] // 不选择该 i 物品
			yes := dp[j-v[i]] + w[i] // 选择该 i 物品
			dp[j] = max(no, yes)
		}
	}
	return dp[C]
}
/* 416. 分割等和子集-Partition Equal Subset Sum
给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
Given a non-empty array nums containing only positive integers,
find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.
 */
/* 能否将一个数组分成两个「等和」子集 ==> 能否从数组中挑选若干个元素，使得元素总和等于所有元素总和的一半
** ==> 背包容量为target=sum/2，每个数组元素的「价值」与「成本」都是其数值大小，求我们能否装满背包。
** dp[i][j]表示前 i 个数值，其选择数字值的总和不超过j的最大价值
** dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]] + nums[i]
** dp[n-1][target] == target =>  如果最大价值等于 target，说明可以拆分成两个「等和子集」
 */
func CanPartition(nums []int) bool {
	sum := 0
	for i := range nums{
		sum += nums[i]
	}
	// 对应了总和为奇数的情况，注定不能被分为两个「等和子集」
	if sum & 1 == 1{
		return false
	}
	target := sum >> 1
	dp := make([]int, target+1)
	for _, t := range nums{
		for j := target; j >= 0; j--{
			// 不选 第 i 个数
			no := dp[j]
			// 选第 i 个数
			yes := 0
			if j >= t{
				yes = dp[j-t] + t
			}
			dp[j] = max(yes, no)
		}
	}
	// 如果最大价值等于 target，说明可以拆分成两个「等和子集」
	return dp[target] == target
}
/* 直接思路
  状态定义：
  	dp[i][j]定义为 前 i 个数字，其选择的数字总和是否为 j
  状态转移方程：还是对应 nums[i] 选 与 不选
	dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
  ==> 转1维度
	dp[j] = dp[j] || dp[j-nums[i]]
 */
func CanPartitionDP(nums []int) bool {
	sum := 0
	for i := range nums{
		sum += nums[i]
	}
	// 对应了总和为奇数的情况，注定不能被分为两个「等和子集」
	if sum & 1 == 1{
		return false
	}
	target := sum >> 1
	// 取消「物品维度」
	dp := make([]bool, target+1)
	dp[0] = true // 初始化状态，对应不选任何
	for i := 1; i <= len(nums); i++{
		for j := target; j >= 0; j--{
			no := dp[j]
			yes := false
			if j >= nums[i-1]{
				yes = dp[j-nums[i-1]]
			}
			dp[j] = no || yes
		}
	}
	return dp[target]
}
/* 完全背包问题
约束条件：每种物品的数量为无限个，你可以选择任意数量的物品。
 */
/*

 */
/* 322. Coin Change
BST: 源于 可构成一个树 或者图，找寻最短路径
*/
func CoinChangeBFS(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	// 优化1： 避开重复计算的情况
	visited := map[int]bool{amount: true}
	// 优化2： 对coins 升序
	sort.Ints(coins)
	q := []int{amount}
	ans := 0
	for len(q) > 0 {
		tmp := []int{}
		ans++
		for _, item := range q{
			for _, c := range coins{
				if item - c == 0{
					return ans
				}
				if item - c > 0{
					if !visited[item-c]{
						tmp = append(tmp, item - c)
						visited[item-c] = true
					}
				}else{ // 由于 coins 升序排序，后面的面值会越来越大，剪枝
					break
				}
			}
		}
		q = tmp
	}
	return -1 // 不可直接返回ans，无法换零 情况
}

func CoinChange(coins []int, amount int) int {
	cache := make([]int, amount+1)
	for i := range cache{
		cache[i] = math.MaxInt32
	}
	cache[0] = 0
	var dfs func(num int)int
	dfs = func(num int)int{
		if num < 0{
			return -1
		}
		if num == 0{
			return 0
		}
		if cache[num] != math.MaxInt32{
			return cache[num]
		}
		ret := math.MaxInt32
		for i := range coins{
			t := dfs(num - coins[i])
			//fmt.Println(num, coins[i], "-->", t)
			if t >= 0 {
				ret = min(ret, t+1)
			}
		}
		if ret == math.MaxInt32{
			cache[num] = -1
			return -1
		}
		cache[num] = ret
		return ret
	}
	return dfs(amount)
}
/*
	dp[i][j] = min(dp[i-1][j], dp[i][j-nums[i]] + 1)
	dp[j] = min(dp[j], dp[j-nums[i]] + 1)
 */
func CoinChangeDP(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp{
		dp[i] = math.MaxInt32
	}
	for i := 0; coins[0]*i <= amount; i++{
		dp[coins[0]*i] = i
	}
	for i := range coins{
		for j := 0; j <= amount; j++{
			if j >= coins[i]{
				dp[j] = min(dp[j], dp[j-coins[i]] + 1)
			}
		}
	}
	if dp[amount] == math.MaxInt32{
		return -1
	}
	return dp[amount]
}

/*  08.11. Coin LCCI
Given an infinite number of quarters (25 cents), dimes (10 cents), nickels (5 cents), and pennies (1 cent),
write code to calculate the number of ways of representing n cents. 
(The result may be large, so you should return it modulo 1000000007)
 */
func waysToChange(n int) int {
	mod := 1000000007
	unit := []int{25,10,5,1}
	sizeUnit := len(unit)
	cache := map[[2]int]int{}
	var dfs func(unitIdx, left int)int
	dfs = func(unitIdx, left int)int{
		if left < 0{
			return 0
		}
		if left == 0{
			return 1
		}
		if cache[[2]int{unitIdx,left}] != 0{
			return cache[[2]int{unitIdx,left}]
		}
		if unitIdx >= sizeUnit-1{
			cache[[2]int{unitIdx, left}] = 1
			return 1
		}
		sumCnt := 0
		for i := unitIdx; i < sizeUnit; i++{
			sumCnt = (sumCnt + dfs(i, left - unit[i]))%mod
		}
		return sumCnt
	}
	return dfs(0, n)
}

/* 多重背包问题

 */

/* 混合背包问题

 */

/* 二维费用背包问题

 */

/* 分组背包问题

 */

/* 背包问题求方案数

 */

/* 求背包问题的方案

 */

/* 有依赖的背包问题

 */

/* 0 1 背包问题

 */

/* 0 1 背包问题

 */