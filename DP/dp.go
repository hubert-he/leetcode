package DP

func max(i, j int)int {
	if i > j {
		return i
	}
	return j
}

//
func CanWinNim(n int) bool {
	if n <= 3{
		return true
	}
	if n == 4 {
		return false
	}
	dp := make([]bool, n+1)
	dp[1],dp[2],dp[3] = true,true,true
	for i := 4; i <= n; i++{
		dp[i] = !(dp[i-1] && dp[i-2] && dp[i-3])
	}
	return dp[n]
}
// 121. Best Time to Buy and Sell Stock
/*
  DP一般分为1D 2D 3+D(使用状态压缩),对应的形式为dp(i) dp(i)(j), 二进制dp(i)(j)
  1. DP步骤
    明确dp(i)/ dp(i)(j) 应该表示什么
    根据dp(i)和dp(i-1)的关系得出状态转移方程
    确定初始条件，如dp(0)
  2. 本题思路 --- DP-1D思想
    dp[i] 表示前 i 天的最大利润，因为我们始终要使得利润最大化，则：
 		dp[i] = max{ dp[i-1], prices[i] - minprice }
 */
func MaxProfit(prices []int)int{
	var dfs func(day int)(int, int)
	dfs = func(day int)(minValue, profit int){
		if day == 0{
			return prices[day], 0
		}
		minPrice, profit := dfs(day-1)
		if minPrice > prices[day]{
			minPrice = prices[day]
		}
		return minPrice, max(profit, prices[day] - minPrice)
	}
	_, max_profite := dfs(len(prices) - 1)
	return max_profite
}
// 53. Maximum Subarray
/*
  dp[i] = max{ dp[i-1] + nums[i], nums[i] }
 */
func maxSubArray(nums []int) int {
	total := len(nums)
	if total <= 0{
		return 0
	}
	//dp := make([]int, total)
	dp := make([]int, 2) // 只记录前一个和当前值
	dp[0] = nums[0]
	maxium := dp[0]
	for i := 1; i < total; i++{
		dp[i%2] = max(dp[(i-1)%2]+nums[i], nums[i])
		maxium = max(dp[i%2], maxium)
		/* 优化空间
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		maxium = max(dp[i], maxium)
		 */
	}
	return maxium
}
/*
  线段树
 */
type Status struct {
	lSum, rSum, mSum, iSum int
}
func maxSubArrayII(nums []int) int {
	return get(nums, 0, len(nums) - 1).mSum;
}
func get(nums []int, l, r int) Status{
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l+r) >> 1
	lSub := get(nums, l, m)
	rSub := get(nums, m+1, r)
	return pushUp(lSub, rSub)
}
func pushUp(l, r Status) Status{
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum + r.lSum)
	rSum := max(r.rSum, r.iSum + l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum + r.lSum)
	return Status{lSum, rSum, mSum, iSum}
}
// 122 Best Time to Buy and Sell Stock II
/*
  dp[i][0]: 第i天交易完成后未持有股票的最大利润
  dp[i][1]: 第i天交易完成后持有一只股票的最大利润
  dp[i][0] = max{ dp[i-1][0], dp[i-1][1] + prices[i] }
  dp[i][1] = max{ dp[i-1][1], dp[i-1][0] - prices[i] }
  初始状态：dp[0][0] = 0, dp[0][1] = -prices[0]
 */
func StockII_maxProfit(prices []int) int {
	days := len(prices)
	if days == 0{
		return 0
	}
	pre := [2]int{0, -prices[0]}
	dp := [2]int{}
	for i := 1; i < days; i++{
		dp[0] = max(pre[0], pre[1] + prices[i])
		dp[1] = max(pre[1], pre[0] - prices[i])
		pre = dp
	}
	return dp[0]
}
/*
  贪心: 由于股票的购买没有限制，因此整个问题等价于寻找 xx 个不相交的区间 (li,ri]使得如下的等式最大化
   贪心的角度考虑我们每次选择贡献大于 0 的区间即能使得答案最大化
 */
func StockII_maxProfitII(prices []int) int {
	result := 0
	for i := 1; i < len(prices); i++{
		result += max(0, prices[i] - prices[i-1])
	}
	return result
}
//123. Best Time to Buy and Sell Stock III
/*
 最多可以完成两笔交易，因此在任意一天结束之后，我们会处于以下五个状态中的一种：
	未进行过任何操作；								利润为0，不用记录
	只进行过一次买操作；  							最大利润 buy1
	进行了一次买操作和一次卖操作，即完成了一笔交易；		最大利润 sell1
	在完成了一笔交易的前提下，进行了第二次买操作；		最大利润 buy2
	完成了全部两笔交易。							最大利润 sell2
    buy1[i] = max{ buy1[i-1], -prices[i] }
	sell1[i] = max{ sell1[i-1],  buy[i-1]+prices[i]}
	buy2[i] = max { buy2[i-1], sell1[i-1]-prices[i] }
	sell2[i] = max { sell2[i-1], buy2[i-1]+prices[i] }
	初始条件：
	buy1[0] = -prices[0]
	sell1[0] = 0
	buy2[0] = -prices[0]
	sell2[0] = 0
 */
func StockIII_maxProfit(prices []int) int {
	if len(prices) <= 0{
		return 0
	}
	pre := [4]int{-prices[0], 0, -prices[0], 0}
	dp := [4]int{}
	for i := 1; i < len(prices); i++{
		dp[0] = max(pre[0], -prices[i])
		dp[1] = max(pre[1], pre[0]+prices[i])
		dp[2] = max(pre[2], pre[1] - prices[i])
		dp[3] = max( pre[3], pre[2] + prices[i])
		pre = dp // 拷贝动作 耗费
	}
	return dp[3]
}

// 309. Best Time to Buy and Sell Stock with Cooldown
/*
  n 表示股票价格数组的长度
  i 表示第 i([0,n-1]) 天
  k 表示允许的最大交易次数
  T[i][k] 表示在第 i 天 结束时，最多进行 k 次交易的情况下可以获得的最大收益
  base condition：
   T[-1][k] = T[i][0] = 0,表示没有进行股票交易时没有收益
  子问题关联：
   如果可以将T[i][k]关联到子问题，例如 T[i-1][k]、T[i][k-1]、T[i-1][k-1]等子问题，就可以得到状态转移方程
  状态转移方程：
    看第 i 天 可能有的操作。即 buy  sell  rest。
    无法直接得知哪个操作是最优的，需要通过计算得到选择每个操作可以得到的最大收益。
    假设没有限制条件，则可以尝试每一种操作，并选择可以最大化收益的一种操作。
    但是题目中确实有条件限制，规定不能同时进行多次交易, 即 如果决定在第 i 天 buy, 在buy之前必须持有 0 份股票，如果决定在第 i 天卖出
    在 sell 之前必须恰好持有 1 份 股票。持有股票的数量是隐藏因素，该因素影响第 i 天 可以进行的操作，进而影响最大收益。
    需要对 T[i][k] 的定义需要分成2项：
    <1> T[i][k][0] 表示在第 i 天结束时，最多进行 k 次交易且在进行操作后持有 0 份股票的情况下可以获得的最大收益。
    <2> T[i][k][1] 表示在第 i 天结束时，最多进行 k 次交易且在进行操作后持有 1 份股票的情况下可以获得的最大收益。
    base condition 继而分成
     T[-1][k][0] = 0, T[-1][k][1] = -Infinity  -Infinity表示没有进行股票交易时，不许持有股票
     T[i][0][0] = 0, T[i][0][1] = -Infinity
    状态转移方程：-- 与 i-1 天关联
     T[i][k][0] = max{ T[i-1][k][0], T[i-1][k][1] + prices[i] }
	 T[i][k][1] = max{ T[i-1][k][1], T[i-1][k-1][0] - prices[i] }
 */

















