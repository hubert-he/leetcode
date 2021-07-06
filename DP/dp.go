package DP

import (
	"fmt"
	"math"
)

func max(i, j int)int {
	if i > j {
		return i
	}
	return j
}
func min(i, j int) int {
	if i > j {
		return j
	}
	return i
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

// 53. Maximum Subarray
/*
  nums数组的长度是n，下标从 0 到 n-1
  dp[i]表示以第i个数结尾的 Maximum Subarray, 答案就是 max{dp[i]} i从0到n-1
  故只需要求出每个位置的dp[i] 然后返回最大的哪个即可。
  与 dp[i-1]关系即 nums[i] 要不要并入到 dp[i-1] 构成 dp[i]
  求 dp[i] 思路： 考虑nums[i]单独成为一段 还是 加入dp[i-1]对应的那一段，取决于 nums[i] 和 dp[i-1] + nums[i]大小的比较。
  因此： dp[i] = max{ dp[i-1] + nums[i], nums[i] }
  一个例子：nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4]
  dp[0] = -2   dp[1] = 1  dp[2] = -2 .... dp[6] = 6 dp[7] = 1 dp[8] = 5
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

func maxSubArraySimply(nums []int) int {
	dp := nums[0]
	ans := dp
	for i := 1; i < len(nums); i++{
		dp = max(dp + nums[i], nums[i])
		ans = max(dp, ans)
	}
	return ans
}

/* 16.17. 连续数列
	利用前缀和计算，但是复杂度在 平方级别
 */
func maxSubArrayPrefixSum(nums []int) int {
	prefixSum := []int{0}
	for i := 0; i < len(nums); i++{
		prefixSum = append(prefixSum, prefixSum[i] + nums[i])
	}
	ans := math.MinInt32
	for i := 0; i < len(prefixSum); i++{
		for j := 0; j < i; j++{
			diff := prefixSum[i] - prefixSum[j]
			if diff > ans{
				ans = diff
			}
		}
	}
	return ans
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
/* cash-不持股  stock-持股
                     cash
                rest/     \buy
                 cash     stock
                /    \   /rest \sell
                        stock cash
 */
func MaxProfitBrust(prices []int) int{
	days := len(prices)
	if days < 2 {
		return 0
	}
	var ans *[]string
	res := 0
	var dfs func(int, int, int,  []string)
	// status 0 表示不持有股票，1表示持有股票
	dfs = func(day, profit, status int, action []string){
		if day == days {
			if res < profit {
				res = profit
				ans = &action
			}
			return
		}
		dfs(day + 1, profit, status, append(action, "rest")) // 不交易
		if status == 0{ // 尝试 buy
			dfs(day+1, profit - prices[day], 1, append(action, "buy"))
		}else { // 尝试 sell
			dfs(day+1, profit + prices[day], 0, append(action, "sell"))
		}
	}
	dfs(0, 0, 0, []string{})
	fmt.Println(ans)
	return res
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
/*
  另通用思路解决
  T[i][k][0] = max{ T[i-1][k][0], T[i-1][k][1] + prices[i] }
  T[i][k][1] = max{ T[i-1][k][1], T[i-1][k-1][0] - prices[i] }
  当 k = 2 时，每天有4个未知变量：
  T[i][1][0] = max{ T[i-1][1][0], T[i-1][1][1] + prices[i] }
  T[i][1][1] = max{ T[i-1][1][1], T[i-1][0][0] - prices[i] }  = max{ T[i-1][1][1], -prices[i] }
  T[i][2][0] = max{ T[i-1][2][0], T[i-1][2][1] + prices[i] }
  T[i][2][1] = max{ T[i-1][2][1], T[i-1][1][0] - prices[i] }

 */
func StockIII_maxProfitII(prices []int) int {
	days := len(prices)
	if days <= 0{
		return 0
	}
	var dp [][3][2]int = make([][3][2]int, days)
	dp[0][1][0] = 0
	dp[0][1][1] = -prices[0]
	dp[0][2][0] = 0
	dp[0][2][1] = -prices[0]
	for i := 1; i < days; i++{
		dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1] + prices[i])
		dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][0]-prices[i])
		dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][2][1] + prices[i])
		dp[i][2][1] = max(dp[i-1][2][1], dp[i-1][1][0] - prices[i])
	}
	return dp[days -1][2][0]
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
    base condition 继而分成 记 k 为 买入操作会使用一次交易
     T[-1][k][0] = 0, T[-1][k][1] = -Infinity  -Infinity表示没有进行股票交易时，不许持有股票
     T[i][0][0] = 0, T[i][0][1] = -Infinity
    状态转移方程：-- 与 i-1 天关联
     T[i][k][0] = max{ T[i-1][k][0], T[i-1][k][1] + prices[i] }
	 T[i][k][1] = max{ T[i-1][k][1], T[i-1][k-1][0] - prices[i] }
    最优值：
     为了得到最后一天结束时的最大收益，通过遍历价格数组，根据状态转移方程计算T[i][k][0]和T[i][k][1]的值。
     因为结束时持有0份股票的收益一定大于持有1份股票的收益，所以最终答案是T[n-1][k][0]
 */
/*
  T[i][0] = max{ T[i-1][0], T[i-1][1]+prices[i] }
  T[i][1] = max{ T[i-1][1], T[i-2][0]-prices[i] }
 */
func Stock309_maxProfit(prices []int) int {
	days := len(prices)
	if days <=1 {
		return 0
	}
	dp := make([][2]int, days)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[1][0] = max(0, -prices[0]+prices[1])
	dp[1][1] = max(-prices[0], -prices[1])
	for i := 2; i < days; i++{
		dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-2][0] - prices[i])
	}
	return dp[days-1][0]
}
// 714 Best Time to Buy and Sell Stock with Transaction Fee
/*
  股票问题最通用的情况由三个特征决定：当前的天数 i、允许的最大交易次数 k 以及每天结束时持有的股票数。
  这篇文章阐述了最大利润的状态转移方程和终止条件，
  由此可以得到时间复杂度为 O(nk)O(nk) 和空间复杂度为 O(k)O(k) 的解法
  链接：https://leetcode-cn.com/circle/article/qiAgHn/
 */
func maxProfitWithFee(prices []int, fee int) int {
	days := len(prices)
	if days <= 1{
		return 0
	}
	dp := make([][2]int, 2)
	dp[0][0], dp[0][1] = 0, -prices[0]
	for i := 1; i < days; i++{
		dp[1][0] = max(dp[0][0], dp[0][1] + prices[i] - fee)
		dp[1][1] = max(dp[0][1], dp[0][0] - prices[i])
		dp[0][0] = dp[1][0]
		dp[0][1] = dp[1][1]
	}
	return max(dp[days-1][0], dp[days-1][1])
}

// 188. Best Time to Buy and Sell Stock IV
/*
  dp[i][k][0] = max{ dp[i-1][k][0], dp[i-1][k][1] + prices[i] }
  dp[i][k][1] = max{ dp[i-1][k][1], dp[i-1][k-1][0] - prices[i] }

  优化-1：
  一个有收益的交易至少需要两天（在前一天买入，在后一天卖出，前提是买入价格低于卖出价格）。
  如果股票价格数组的长度为 n，则有收益的交易的数量最多为 n / 2（整数除法）。因此 k 的临界值是 n / 2。
  如果给定的 k 不小于临界值，即 k >= n / 2，则可以将 k 扩展为正无穷，此时问题等价于情况二
 */
func StockIV_maxProfit(k int, prices []int) int {
	days := len(prices)
	if days <= 1{
		return 0
	}
	if (k >= days / 2) { // 优化-1
		return StockII_maxProfit(prices);
	}
	dp := make([][][2]int, days)
	for i := 0; i < days; i++{
		/* 设置k=0 为屏蔽墙， 读入默认0，防止数组出界 */
		dp[i] = make([][2]int, k+1)
	}
	for i := 1; i <= k; i++{
		dp[0][i][0] = 0
		dp[0][i][1] = -prices[0]
	}
	for i := 1; i < days; i++{
		for j := 1; j <= k; j++{
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1] + prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0] - prices[i])
		}
	}
	return dp[days-1][k][0]
}
// 学习空间优化： 第 i 天的最大收益只和第 i - 1 天的最大收益相关，空间复杂度可以降到 O(k)O(k)
func StockIV_maxProfitII(k int, prices []int) int {
	days := len(prices)
	if days <= 1{
		return 0
	}
	if k >= days / 2{
		return StockII_maxProfit(prices)
	}
	dp := make([][2]int, k+1)
	for i := 1; i <= k; i++{
		dp[i][0] = 0
		dp[i][1] = -prices[0]
	}
	for i := 1; i < days; i++{
		for j := 1; j <= k; j++{
			dp[j][0] = max(dp[j][0], dp[j][1] + prices[i])
			dp[j][1] = max(dp[j][1], dp[j-1][0] - prices[i])
		}
	}
	return dp[k][0]
}

// 327. Count of Range Sum  so so hard!!!!
func countRangeSum(nums []int, lower int, upper int) int {
	return 0
}

// 118. Pascal's Triangle
/*
	递推公式：
     F[i][j] = F[i-1][j-1] + F[i-1][j]
	F[i][0] = F[i][j] = 0
 */
func generate(numRows int) [][]int {
	ans := [][]int{[]int{1}}
	if numRows <= 1{
		return ans
	}
	for i := 1; i < numRows; i++{
		tmp := make([]int, i+1)
		tmp[0], tmp[i] = 1, 1
		for j := 1; j < i; j++{
			tmp[j] = ans[i-1][j-1] + ans[i-1][j]
		}
		ans = append(ans, tmp)
	}
	return ans
}
// 119. Pascal's Triangle II
func getRow(rowIndex int) []int {
	ans := []int{1}
	if rowIndex <= 0{
		return ans
	}
	for i := 1; i <= rowIndex; i++{
		tmp := make([]int, i+1)
		tmp[0], tmp[i] = 1, 1
		for j := 1; j < i; j++{
			tmp[j] = ans[j-1] + ans[j]
		}
		ans = tmp
	}
	return ans
}

/*
   17.16. 按摩师
 */
func massage(nums []int) int {
	return 0
}

/* 08.01. 三步问题
  dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
  dp[0] = 0  dp[1] = 1
  dp[2] = 2  dp[3] = 4
  1. 需要注意超过32位整数的处理
*/
func WaysToStep(n int) int {
	dp := make([]int, n+3) // 排除 1 2 3 特例
	dp[1] = 1
	dp[2] = 2
	dp[3] = 4
	for i := 4; i <= n; i++{
		dp[i] = ((dp[i-1] + dp[i-2]) % 1000000007 + dp[i-3])% 1000000007
	}
	return dp[n]
	// return dp[n] % 1000000007 在这里取模 没有意义 上面加法就已经溢出了
}
/*
   内存优化版
   循环使用dp区
 */
func WaysToStepII(n int) int{
	dp := make([]int, 4)
	dp[1], dp[2], dp[3] = 1,2,4
	for i := 4; i <= n; i++{
		dp[i%4] = ((dp[(i-1)%4] + dp[(i-2)%4]) % 1000000007 + dp[(i-3)%4] ) % 1000000007
	}
	return dp[n%4]
}

/*
  788. Rotated Digits
  1. 把 数 转换为字符串解析
  2. 递归检查 数的 最后一位数字 --- 掌握这种方法
 */
func RotatedDigits(n int) int {
	ans := 0
	var isGood func(int, bool) bool
	isGood = func(num int, flag bool) bool {
		if num == 0{
			return flag
		}
		d := num % 10
		if d == 3 || d == 4 || d == 7{
			return false
		}
		if d == 0 || d == 1 || d == 8{ // 需要判断其他数字是否也是旋转后依然是自身的数字
			return isGood(num / 10, flag)
		}
		return isGood(num / 10, true)
	}
	for i := 1; i <=n; i++{
		if isGood(i, false){
			ans++
		}
	}
	return ans
}
/*
	根据好数定义，每个好数只能包含数字 0125689，并且至少包含 2569 中的一个。因此可以逐个写出小于等于 N 的所有好数。
	状态可以表示为3个变量，即 i, equality_flag, involution_flag
	i: 表示当前正在写第 i 位数字
	equality_flag: 表示已经写出的 j 位数字是否等于 N 的 j 位前缀。
	involution_flag: 表示从最高位到比当前位高一位的这段前缀中是否含有 2 5 6 9 中的任意一个数字。
	状态转移方程：
	dp(i, equality_flag, involution_flag) 表示在特定 equality_flag 和 involution_flag的状态下，
	有多少种从 i 到末尾的后缀能组成一个好数。 最终的结果为 dp(0, true, false)
	注意：N 从最高位到最低位的索引，从 0 开始增大。 第 i 位表示索引为 i 的位置。
 */

/*
	定义dp[i]状态为：数字 i 的三种状态，即
	1. i 翻装后与原数相同 记为2
	2. i 翻装后与原数不同 记为1
	3. i 翻装后是非法数字 记为0
  i 为 1位数字时：
	dp[i] = 1 when i 为 0 1 8
	dp[i] = 2 when i 为 2 5 6 9
	dp[i] = 0 when i 为 3 4 7
  i > 10时
    dp[i]的值来自于 dp[i/10] dp[i % 10] 这2个的状态 得出
 */
func RotatedDigitsDP2(n int) int{
	ans, dp := 0, make([]int, n+1)
	for i := 0; i <= n; i++{
		if i < 10{
			if i == 0 || i == 1 || i == 8{
				dp[i] = 1
			}else if i == 2 || i == 5 || i == 6 || i == 9{
				dp[i] = 2
				ans++
			}
		}else {
			quotient, left := dp[i/10], dp[i%10]
			if quotient == 1 && left == 1 {
				// 前缀和个位数均是反转后是相同的数字
				dp[i] = 1
			} else if quotient >= 1 && left >= 1 {
				dp[i] = 2
				ans++
			}
		}
	}
	return ans
}
/* 回溯： 每步可以选当前数和不选当前数
   LCS 01. 下载插件
 */
func leastMinutes(n int) int { // n 为 n 个插件
	ans := math.MaxInt32
	var bt func(int, int, int)
	// cur: 当前下载数量   idx: 递归层数  cnt: 当前带宽
	bt = func(cur, idx, cnt int){
		if idx >= ans || cnt * 2 > math.MaxInt32{
			return
		}
		if cur >= n{
			ans = min(ans, idx)
			return
		}
		bt(cur, idx + 1, cnt * 2)   // 不选当前cnt，则加倍
		bt(cur + cnt, idx + 1, cnt) // 选择当前cnt
	}
	bt(0, 0, 1)
	return ans
}

/* 1025. Divisor Game
   n 为奇数的时候Alice（先手）必败，nn 为偶数的时候 Alice 必胜
  博弈类的问题常常让我们摸不着头脑。当我们没有解题思路的时候，不妨试着写几项试试：
	n = 1 的时候，区间 (0, 1) 中没有整数是 n 的因数，所以此时 Alice 败。
	n = 2 的时候，Alice 只能拿 1，n 变成 1，Bob 无法继续操作，故 Alice 胜。
	n = 3 的时候，Alice 只能拿 1，n 变成 2，根据 n = 2 的结论，我们知道此时 Bob 会获胜，Alice 败。
	n = 4 的时候，Alice 能拿 1 或 2，如果 Alice 拿 1，根据 n = 3的结论，Bob 会失败，Alice 会获胜。
	n = 5 的时候，Alice 只能拿 1，根据 n = 4 的结论，Alice 会失败。
	Alice 处在n = k 状态时，Alice的每一步操作，必然使得 Bob处于 n = m(m < k)的状态。因此我们只要看是否存在一个 m 是必败的状态，
    那么 Alice 直接执行对应的操作让当前的数字变成 m，Alice 就必胜了，如果没有任何一个是必败的状态的话，说明 Alice 无论怎么进行操作
	最后都会让 Bob 处于必胜的状态，此时 Alice 是必败的
    定义 dp[i] 表示当前数字 i 的时候先手是处于必胜态还是必败态，true 表示先手必胜， false表示先手必败。
    从前往后递推，枚举 i 在 （0，i) 中 i 的因数 j ，看是否存在 dp[i-j]为必败态即可。
 */
func divisorGame(n int) bool {
	dp := make([]bool, n + 2) // 多选一个数字 可以避开特殊条件，比如 n = 1 或 0 的情况
	dp[1], dp[2] = false, true
	for i := 3; i <= n; i++{
		for j := 1; j < i; j++{
			if i % j == 0 && !dp[i-j]{
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}











