package linear

import (
	"container/heap"
	"math"
	"sort"
	"strconv"
)
/*
func max(i, j int)int {
	if i > j {
		return i
	}
	return j
}
 */
func max(nums ...int)int{
	m := nums[0]
	for _, u := range nums{
		if m < u{
			m = u
		}
	}
	return m
}

func min(nums ...int) int {
	m := nums[0]
	for _, c := range nums{
		if m > c{
			m = c
		}
	}
	return m
}
// 排除except 下标的 值
func min2(except int, nums ...int) int{
	m := math.MaxInt32
	for i := range nums{
		if m > nums[i] && i != except{
			m = nums[i]
		}
	}
	return m
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
/* 暴力 */
func maxSubArrayBurst(nums []int) int{
	ans := nums[0]
	for i := range nums{
		for j := i; j < len(nums); j++{
			tmp := 0
			for k := i; k <= j; k++{
				tmp += nums[k]
			}
			ans = max(tmp, ans)
		}
	}
	return ans
}
func maxSubArrayBurst2(nums []int) int{
	ans := nums[0]
	for i := range nums{
		acc := 0
		for j := i; j < len(nums); j++{
			acc += nums[j]
			ans = max(acc, ans) // 直接处理
		}
	}
	return ans
}
/*
	贪心： O(n)考虑全负数情况
 */
func maxSubArrayGreedy(nums []int)(ans int){
	ans = nums[0]
	sum := 0
	for _, num := range nums{
		sum = max(sum + num, num) // 当前和加上当前值，与当前值 进行比较, 不能改为 判断是否为负数， 因为有全是负数情况
		ans = max(ans, sum)
	}
	return
}
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

/* 152. Maximum Product Subarray
** Given an integer array nums, find a contiguous non-empty subarray within the array that has the largest product,
** and return the product.
** It is guaranteed that the answer will fit in a 32-bit integer.
** A subarray is a contiguous subsequence of the array.
*/
/* 考虑当前位置如果是一个负数的话，那么我们希望以它前一个位置结尾的某个段的积也是个负数，这样就可以负负得正，并且我们希望这个积尽可能「负得更多」，即尽可能小。
** 如果当前位置是一个正数的话，我们更希望以它前一个位置结尾的某个段的积也是个正数，并且希望它尽可能地大。于是这里我们可以再维护一个 fmin(i)，
** 它表示以第 i 个元素结尾的乘积最小子数组的乘积，那么我们可以得到这样的动态规划转移方程：
** 代码所示 转移方程
 */
func maxProduct(nums []int) int {
	n := len(nums)
	dp := make([]int, 2) // 0:存最小 1：存最大
	dp[0], dp[1] = nums[0], nums[0]
	//ans := math.MinInt32 注意初始值
	ans := nums[0]
	for i := 1; i < n; i++{
		t := make([]int, 2)
		// 取dp[0]*nums[i], dp[1]*nums[i], nums[i] 之间的极大 极小 值
		t[0] = min(dp[0]*nums[i], dp[1]*nums[i], nums[i])
		t[1] = max(dp[0]*nums[i], dp[1]*nums[i], nums[i])
		dp = t
		ans = max(ans, dp[1])
	}
	return ans
}

/* 487. Max Consecutive Ones II
** Given a binary array nums, return the maximum number of consecutive 1's in the array if you can flip at most one 0.
** Follow up:
	What if the input numbers come in one by one as an infinite stream? In other words,
	you can't store all numbers coming from the stream as it's too large to hold in memory.
	Could you solve it efficiently?
*/
/* 必须把 是否flip 0 的状态 反映到 动态方程中
** z这个是看到使用dp ，才想起的转移方程
*/
func FindMaxConsecutiveOnes(nums []int) int {
	dp := make([]int, 2)
	ans := 0
	for i := range nums{
		if nums[i] == 0{
			dp[0] = dp[1]+1
			dp[1] = 0
			ans = max(ans, dp[0])
		}else{
			dp[0]++
			dp[1]++
			ans = max(ans, dp[0], dp[1])
		}
	}
	return ans
}
/* 918. Maximum Sum Circular Subarray
** Given a circular integer array nums of length n, return the maximum possible sum of a non-empty subarray of nums.
** A circular array means the end of the array connects to the beginning of the array.
** Formally, the next element of nums[i] is nums[(i + 1) % n] and the previous element of nums[i] is nums[(i - 1 + n) % n].
** A subarray may only include each element of the fixed buffer nums at most once.
** Formally, for a subarray nums[i], nums[i + 1], ..., nums[j], there does not exist i <= k1, k2 <= j with k1 % n == k2 % n.
 */
/* 对于环形数组，分两种情况:
** 1. 答案在数组中间，就是最大子序和
** 2. 答案在数组两边, 例如[5,-3,5]最大的子序和就等于数组的总和SUM-最小的子序和
** 3. 一种特殊情况是数组全为负数，也就是SUM-最小子序和==0，最大子序和等于数组中最大的那个
 */
func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	sum := nums[0]
	dpmax, dpmin := nums[0], nums[0]
	ans := nums[0]
	minVal := nums[0]
	for i := 1; i < n; i++{
		sum += nums[i]
		dpmax = max(dpmax + nums[i], nums[i])
		dpmin = min(dpmin + nums[i], nums[i])
		if ans < dpmax{
			ans = dpmax
		}
		if minVal > dpmin{
			minVal = dpmin
		}
	}
	if sum - minVal != 0{
		return max(ans, sum-minVal)
	}
	return ans
}
/* Kadane 算法: 用来找到 A 的最大子段和, 基于动态规划
**
 */

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
	lSum int // [l, r] 内以 l 为左端点的最大字段和
	rSum int // [l, r] 内以 r 为右端点的最大字段和
	mSum int // [l, r] 内的最大字段和
	iSum int // [l, r] 的区间和
}
/* 定义一个操作get(a, l, r)表示查询a序列[l, r]区间内最大字段和
   分治实现：对一个区间[l,r]，取 mid = (l+r) >> 1, 然后对[l,m] 和 [m+1,r]分治求解
        当递归逐层深入直到区间长度缩小为1的时候，递归回升，然后处理 通过[l,m]区间的信息和[m+1,r]区间的信息合并成区间[l,r]的信息。
*/
func maxSubArrayII(nums []int) int {
	var get func(int, int) Status
	get = func(l, r int) Status{
		if l == r{
			return Status{nums[l],nums[l],nums[l],nums[l]}
		}
		m := (l+r) >> 1
		lSub := get(l, m)
		rSub := get(m+1, r)
		return pushUp(lSub, rSub) // 回升
	}
	return get(0, len(nums) - 1).mSum;
}
/* [l,m]: [l,r]的左子区间
   [m+1,r]: [l,r]的右子区间
	iSum: [l,r]的iSum就等于左子区间 加 右子区间的 iSum
	lSum: 存在2种可能，要么等于左子区间的 lSum, 要么等于 左子区间的 iSum + 右子区间的 lSum， 二者取最大。
	rSum: 要么等于右子区间的rSum, 要么等于右子区间的 iSum + 左子区间的 rSum， 二者取最大。
	mSum: [l,r]的mSum对应的区间是否跨越 m（也可能不跨域m) 即 [l,r]的mSum 可能是 左子区间的 mSum 和 右子区间 mSum 中的一个
          也可能是 左子区间的 rSum 和 右子区间的 lSum求和。 三者取最大。
 */
func pushUp(l, r Status) Status{
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum + r.lSum)
	rSum := max(r.rSum, r.iSum + l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum + r.lSum)
	return Status{lSum: lSum, rSum: rSum, mSum: mSum, iSum: iSum}
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
	//var ans *[]string
	res := 0
	var dfs func(int, int, int,  []string)
	// status 0 表示不持有股票，1表示持有股票
	dfs = func(day, profit, status int, action []string){
		if day == days {
			if res < profit {
				res = profit
				//ans = &action
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
	//fmt.Println(ans)
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

// LCP 07. 传递信息
/* BFS： 在图中搜索计算方案数
 */
func numWays(n int, relation [][]int, k int) int {
	q := []int{0}
	total := len(relation)
	for len(q) > 0{
		if k <= 0{
			break
		}
		tmp := []int{}
		for _, elem := range q{
			for i := 0; i < total; i++{
				if relation[i][0] == elem {
					tmp = append(tmp, relation[i][1])
				}
			}
		}
		q = tmp
		k--
	}
	ans := 0
	for _, item := range q{
		if item == n - 1{
			ans++
		}
	}
	return ans
}

/*
  DFS_BFS
 */
func numWaysDFS(n int, relation [][]int, k int) int {
	ans := 0
	var dfs func(int, int)
	dfs = func(round, curr int){
		if round == k{
			if curr == n-1{
				ans++
			}
			return
		}
		for _, r := range relation{
			src,dst := r[0], r[1]
			if src == curr{
				dfs(round+1, dst)
			}
		}
	}
	dfs(0, 0)
	return ans
}
/* 计数问题，很大可能可通过 DP 优化
 dp: 定义状态 dp[i][j] 为经过 i 轮 传递到编号 j 的玩家的方案数，其中 0 <= i <= k, 0 <= j < n
	 由于 从编号 0 的玩家开始传递，当 i = 0时，一定位于编号0的玩家，不会传递到其他玩家
	 dp[0][j] = 1(j = 0) 0(j != 0)
	对于信息传递关系[src,dst], 如果第i轮传递到编号src的玩家，则第 i + 1 轮可以从编号 src 的玩家传递到编号 dst的玩家。
	在计算dp[i+1][dst]时，需要考虑可以传递到编号dst的所在玩家。状态转移方程：
	dp[i+1][dst] = SUM{dp[i][src]} [src,dst]属于relation
	最终得到 dp[k][n-1]的方案数
 */
func numWaysDP(n int, relation [][]int, k int) int {
	 dp := make([][]int, k+1)
	 for i := range dp{
	 	dp[i] = make([]int, n)
	 }
	 dp[0][0] = 1
	 for i := 0; i < k; i++{
	 	for _, r := range relation{
	 		src, dst := r[0], r[1]
	 		dp[i+1][dst] += dp[i][src]
		}
	 }
	 return dp[k][n-1]
}

/* 740. Delete and Earn
** You are given an integer array nums.
** You want to maximize the number of points you get by performing the following operation any number of times:
** Pick any nums[i] and delete it to earn nums[i] points.
** Afterwards, you must delete every element equal to nums[i] - 1 and every element equal to nums[i] + 1.
** Return the maximum number of points you can earn by applying the above operation some number of times.
 */
/* 不要被 nums[i] + 1 干扰掉*/
func DeleteAndEarnDFS(nums []int)int{
	memo := make([]int, 1e4+1)
	cnt := make([]int, 1e4+1)
	for _, num := range nums{
		cnt[num]++
	}
	var dfs func(num int)int
	dfs = func(num int)int{
		if num <= 0{
			return 0
		}
		if cnt[num] == 0{
			return dfs(num-1)
		}
		if memo[num] != 0{
			return memo[num]
		}
		res := cnt[num]*num
		if cnt[num-1] != 0{
			res = max(dfs(num-1), dfs(num-2)+res)
		}else{
			res += dfs(num-2)
		}
		memo[num] = res
		return res
	}
	for i := int(1e4); i >= 1; i--{
		if cnt[i] != 0{
			return dfs(i)
		}
	}
	return 0
}
/*官方题解-1: 转换为 打家劫舍 */
func DeleteAndEarn1(nums []int) int {
	maxVal := max(nums...)
	sum := make([]int, maxVal+1)
	// 避免排序，hash
	for _, u := range nums{
		sum[u] += u
	}
	// sum 就转换为 打家劫舍的 序列
	rob := func() int{
		first, second := sum[0], max(sum[0], sum[1])
		for i := 2; i < maxVal+1; i++{
			first, second = second, max(first+sum[i], second)
		}
		return second
	}
	return rob()
}

// 2021-11-08 重刷此题
// 分为2维，同时将情况枚举为3类， 然后得出每种情况的转移方程
func DeleteAndEarnDP(nums []int)int{
	n := len(nums)
	sort.Ints(nums)
	dp := make([][2]int, n+1)
	dp[0][0], dp[0][1] = 0, 0
	dp[1][0], dp[1][1] = 0, nums[0]
	for i := 2; i <= n; i++{
		diff := nums[i-1] - nums[i-2]
		if diff == 0{ // 相等的情况, 一选都选
			dp[i][0] = dp[i-1][0]
			dp[i][1] = dp[i-1][1] + nums[i-1]
		}else if diff == 1{ // 相差1的情况
			dp[i][0] = max(dp[i-1][0], dp[i-1][1])
			// 查找上一个
			j := i-2
			for j >= 0 && nums[j] == nums[i-2]{
				j--
			}
			dp[i][1] = nums[i-1] + max(dp[j+1][0], dp[j+1][1])
		}else{
			dp[i][0] = max(dp[i-1][0], dp[i-1][1])
			dp[i][1] = dp[i][0] + nums[i-1]
		}
	}
	return max(dp[n][0], dp[n][1])
}

/* 55. Jump Game
** You are given an integer array nums.
** You are initially positioned at the array's first index,
** and each element in the array represents your maximum jump length at that position.
** Return true if you can reach the last index, or false otherwise.
 */
/* 暴力 超时*/
func canJump(nums []int) bool {
	n := len(nums)
	var dfs func(int)bool
	dfs = func(idx int)bool{
		if idx == n-1{
			return true
		}
		if nums[idx] == 0{
			return false
		}
		for i := 1; i <= nums[idx]; i++{
			if dfs(idx+i){
				return true
			}
		}
		return false
	}
	return dfs(0)
}
/* 根据暴力 画出递归树
** 发现重复计算的节点
*/
func canJumpDFSD(nums []int) bool {
	n := len(nums)
	cache := make([]int, n)
	for i := range cache{
		cache[i] = -1
	}
	var dfs func(int)bool
	dfs = func(idx int)bool{
		if idx == n-1{
			return true
		}
		if nums[idx] == 0{
			return false
		}
		if cache[idx] == 0{
			return false
		}
		if cache[idx] == 1{
			return true
		}
		for i := 1; i <= nums[idx]; i++{
			if dfs(idx+i){
				cache[idx+i] = 1
				return true
			}
			cache[idx+i] = 0
		}
		return false
	}
	return dfs(0)
}
/* dp[i] 表示从前面0到i-1个元素是否可以跳到第i个元素上，如果可以为true，否则false
** dp[i] 有0到i-1位置的 dp[j] 决定，如果 dp[j] 为true 并且 j+nums[i] >= i 则可以跳到第 i 个位置
** 初始条件： dp[0] = true
*/
func CanJumpDP(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n)
	dp[0] = true
	for i := 1; i < n; i++{
		for j := 0; j < i; j++{
			if dp[j] && i <= nums[j] + j {
				dp[i] = true
				break
			}
		}
	}
	return dp[n-1]
}

/* 2021-11-12 重刷
** 向前更新状态
*/
func canJump12(nums []int) bool {
	// dp[i] = dp[i+j] j [1,nums[i]]
	n := len(nums)
	dp := make([]bool, n)
	dp[0] = true
	for i := range nums{
		if dp[i]{
			for j := 1; j < n-i && j <= nums[i]; j++{
				dp[i+j] = true
			}
		}
	}
	return dp[n-1]
}

/* 贪心: 如果一个位置能够到达，那么这个位置左侧所有位置都能到达 */
func CanJumpGreedy(nums []int) bool {
	rightMost := 0 // 为当前能够到达的最大位置
	for i := range nums{
		if(i > rightMost) {//【关键】遍历元素位置下标大于当前能够到达的最大位置下标，不能到达
			return false
		}
		//能够到达当前位置，看是否更新能够到达的最大位置righMost
		rightMost = max(rightMost, i+nums[i])
	}
	//跳出则表明能够到达最大位置
	return true
}
/* 45. Jump Game II
** Given an array of non-negative integers nums, you are initially positioned at the first index of the array.
Each element in the array represents your maximum jump length at that position.
Your goal is to reach the last index in the minimum number of jumps.
You can assume that you can always reach the last index.
 */
func JumpDFS(nums []int) int {
	n := len(nums)
	var dfs func(idx int)int
	dfs = func(idx int)int{
		if idx >= n-1{
			return 0
		}
		step := math.MaxInt32
		if  nums[idx] == 0{
			return step
		}
		for i := 1; i <= nums[idx]; i++{
			step = min(step, dfs(idx+i))
		}
		return step+1
	}
	return dfs(0)
}
func JumpDFSDP(nums []int) int {
	n := len(nums)
	cache := make([]int, n)
	var dfs func(idx int)int
	dfs = func(idx int)int{
		if idx >= n-1{
			return 0
		}
		step := math.MaxInt32
		if  nums[idx] == 0{
			return step
		}
		if cache[idx] != 0{
			return cache[idx]
		}
		for i := 1; i <= nums[idx]; i++{
			step = min(step, dfs(idx+i))
		}
		cache[idx] = step+1
		return step+1
	}
	return dfs(0)
}

func JumpDP(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := range dp{
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i < n; i--{
		for j := 0; j < i; j++{
			if (j + nums[j] >= i) {
				dp[i] = min(dp[i], dp[j])
			}
		}
		dp[i] += 1
	}
	return dp[n-1]
}

/* 264. Ugly Number II
** An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.
** Given an integer n, return the nth ugly number.
*/
// 方法一：直接求，方法超时
func nthUglyNumber(n int) int {
	dp := map[int]bool{1: true, 2:true, 3:true}
	if n < 4{
		return n
	}
	k := 4
	for i := 4; i <= n;{
		for {
			if k % 2 == 0{
				if dp[k/2]{
					dp[k] = true
					i++
					k++
					break
				}
			}
			if k % 3 == 0{
				if dp[k/3]{
					dp[k] = true
					i++
					k++
					break
				}
			}
			if k % 5 == 0{
				if dp[k/5]{
					dp[k] = true
					i++
					k++
					break
				}
			}
			k++
		}
	}
	return k-1
}
/* DP2
** 由上面除的方式，改成乘的方式，凑出下一个
 */
func NthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	i, j, k := 1, 1, 1
	for idx := 1; idx <= n; idx++{
		num2, num3, num5 := dp[i]*2, dp[j]*3, dp[k]*5
		dp[idx] = min(num2, num3, num5)
		if dp[idx] == num2{
			i++
		}
		if dp[idx] == num3{
			j++
		}
		if dp[idx] == num5{
			k++
		}
	}
	return dp[n]
}
/* 使用堆
**
*/
type hp struct {sort.IntSlice}
func(h *hp) Push(v interface{}){
	h.IntSlice = append(h.IntSlice, v.(int))
}
func(h *hp) Pop()interface{}{
	n := len(h.IntSlice)
	v := h.IntSlice[n-1]
	h.IntSlice = h.IntSlice[:n-1]
	return v
}
func NthUglyNumberHeap(n int) int {
	h := &hp{sort.IntSlice{1}}
	vis := map[int]bool{1: true}
	ans := 0
	for i := 1;i <=n; i++{
		ans = heap.Pop(h).(int)
		for _, f := range []int{2,3,5}{
			next := ans*f
			// 去重
			if !vis[next]{
				heap.Push(h, next)
				vis[next] = true
			}
		}
	}
	return ans
}
/* 91. Decode Ways
** A message containing letters from A-Z can be encoded into numbers using the following mapping:
'A' -> "1"
'B' -> "2"
...
'Z' -> "26"
To decode an encoded message, all the digits must be grouped then mapped back into letters using the reverse of the mapping above (there may be multiple ways).
For example, "11106" can be mapped into:
"AAJF" with the grouping (1 1 10 6)
"KJF" with the grouping (11 10 6)
Note that the grouping (1 11 06) is invalid because "06" cannot be mapped into 'F' since "6" is different from "06".
Given a string s containing only digits, return the number of ways to decode it.

The answer is guaranteed to fit in a 32-bit integer.
 */
// 2021-11-16 刷出此题，
func numDecodingsI(s string) int {
	n := len(s)
	if s[0] == '0'{
		return 0
	}
	// dp[i] = dp[i-1] + dp[i-2],根据是否与前一位结合 还是不结合
	dp := [2]int{1, 1}
	for i := 2; i <= n; i++{
		// 优化
		//x, _ := strconv.Atoi(s[i-2:i])
		x := (s[i-2]-'0')*10 + s[i-1]-'0'
		if s[i-1] == '0'{// 只能是: 与前一位结合
			if s[i-2] == '0'{ //连续2个0 非法
				return 0
			}else{// 只能是: 与前一位结合
				if x <= 26{
					dp[i%2] = dp[(i-2)%2]
				}else{// 非法
					return 0
				}
			}
		}else{// 可以单独 也可以结合，结合的情况要注意不能超过26
			if s[i-2] == '0'{// 此情况只能 单独 ，因为不能出现前缀 0 的情况
				dp[i%2] = dp[(i-1)%2]
			}else{
				if x <= 26{// 此情况 可单独 亦可 结合
					dp[i%2] = dp[(i-1)%2] + dp[(i-2)%2]
				}else{// 此情况 只可单独
					dp[i%2] = dp[(i-1)%2]
				}
			}
		}
	}
	return dp[n%2]
}

/* 413. Arithmetic Slices
** An integer array is called arithmetic if it consists of at least three elements and if the difference between any two consecutive elements is the same.
For example, [1,3,5,7,9], [7,7,7,7], and [3,-1,-5,-9] are arithmetic sequences.
Given an integer array nums, return the number of arithmetic subarrays of nums.
A subarray is a contiguous subsequence of the array.
 */
/*暴力： 判断所有的子数组是不是等差数列，如果是的话就累加次数
** 复杂度 O(n^3)
*/
func numberOfArithmeticSlicesBrust(nums []int) int {
	sn := len(nums)
	isArithmetic := func(arr []int)bool{
		n := len(arr)
		if n < 3{
			return false
		}
		for i := 1; i < n-1; i++{
			if arr[i-1] - arr[i] != arr[i] - arr[i+1]{
				return false
			}
		}
		return true
	}
	ans := 0
	for i := 0; i < sn-2; i++{
		for j := i+2; j < sn; j++{
			if isArithmetic(nums[i:j+1]){
				ans++
			}
		}
	}
	return ans
}
/*双指针（滑动窗口） 的解法
** 在方法一的暴力枚举中， 我们对每个长度大于等于3的序列都进行了是否为等差数列的判断
** 其实，如果我们已经知道一个子数组的前面部分不是等差数列以后，那么后面部分就不用判断了
** 等差数列的所有的相邻数字的差是固定的
** 因此，对于每个起始位置，我们只需要向后进行一遍扫描，直到不再构成等差数列为止，此时已经没有必要再向后扫描
** 即 固定起点，判断后面的等差数列有多少个
*/
func numberOfArithmeticSlicesTwoPointer(nums []int) int {
	n := len(nums)
	ans := 0
	for i := 0; i < n-2; i++{
		d := nums[i+1] - nums[i]
		for j := i+1; j < n-1; j++{
			if nums[j+1] - nums[j] == d{
				ans++
			}else{
				break
			}
		}
	}
	return ans
}
/* 固定起点，判断后面的等差数列有多少个
** 递归解法: 定义递归函数 slices(A, end)的含义是区间 A[0, end] 中，以 end 作为终点的，等差数列的个数
** A[0, end]内的以 end 作为终点的等差数列的个数，相当于在 A[0, end - 1]的基础上，增加了 A[end]
** 有两种情况：
** 1. A[end] - A[end - 1] == A[end - 1] - A[end - 2]时，说明增加的A[end]能和前面构成等差数列，那么 slices(A, end) = slices(A, end - 1) + 1；
** 2. A[end] - A[end - 1] != A[end - 1] - A[end - 2]时， 说明增加的 A[end]不能和前面构成等差数列，所以slices(A, end) = 0；
** 最后，我们要求的是整个数组中的等差数列的数目，所以需要把 0 <= end <= len(A - 1)0<=end<=len(A−1) 的所有递归函数的结果累加起来。
 */
func numberOfArithmeticSlicesDFS(nums []int) int {
	n := len(nums)
	if n < 2{
		return 0
	}
	ans := 0
	end := n-1
	if nums[end] - nums[end-1] == nums[end-1] - nums[end-2]{
		ans = 1 + numberOfArithmeticSlicesDFS(nums[:end-1])
	}else{
		ans = numberOfArithmeticSlicesDFS(nums[:end-1])
	}
	return ans
}

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	ans := 0
	if n == 1{
		return ans
	}
	d, t := nums[0] - nums[1], 0
	for i := 2; i < n; i++{
		if nums[i-1] - nums[i] == d{
			t++
		}else{
			d, t = nums[i-1] - nums[i], 0
		}
		ans += t
	}
	return ans
}

/*221. Maximal Square
** Given an m x n binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.
*/
// 此题可以用二分法，复杂度在 O(log(min(m,n)*m*n*m*n))
/* dp(i,j) 表示以(i,j) 为右下角, 且只包含 1 的正方形的边长最大值.
** 如果我们能计算出所有 dp(i,j) 的值，那么其中的最大值即为矩阵中只包含 1 的正方形的边长最大值，其平方即为最大正方形的面积
** 如何计算 dp 中的每个元素值呢？对于每个位置 (i,j)，检查在矩阵中该位置的值：
** 1. 如果该位置的值是 0，则 dp(i,j)=0，因为当前位置不可能在由 1 组成的正方形中；
** 2. 如果该位置的值是 1，则 dp(i,j) 的值由其上方、左方和左上方的三个相邻位置的 dp 值决定，即
** 	dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
** 初始以及边界条件：
**  如果i和j中至少有一个为 0， 则以位置(i,j)为右下角的最大正方形的边长只能为1， 即 dp[i][j] = 1
*/
func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := [2][]int{}
	for i := range dp{
		dp[i] = make([]int, n)
	}
	maxSide := 0
	for i := 0; i < m; i++{
		for j := 0; j < n; j++{
			if matrix[i][j] == '0'{
				dp[i%2][j] = 0
			}else{
				if i == 0 || j == 0 {
					dp[i%2][j] = 1
				}else{
					dp[i%2][j] = min(dp[(i-1)%2][j], dp[i%2][j-1], dp[(i-1)%2][j-1]) + 1
				}
				if maxSide < dp[i%2][j]{
					maxSide = dp[i%2][j]
				}
			}
		}
	}
	return maxSide * maxSide
}
/* 746. Min Cost Climbing Stairs
** You are given an integer array cost where cost[i] is the cost of ith step on a staircase.
** Once you pay the cost, you can either climb one or two steps.
** You can either start from the step with index 0, or the step with index 1.
** Return the minimum cost to reach the top of the floor.
*/
// 2021-11-24 重刷此题
func minCostClimbingStairs(cost []int) int {
	// dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	dp := [2]int{}
	n := len(cost)+1 // 跳出去的
	for i := 3; i <= n; i++{
		dp[i%2] = min(dp[(i-1)%2]+cost[i-2], dp[(i-2)%2]+cost[i-3])
	}
	return dp[n%2]
}

/* 1182. Shortest Distance to Target Color
** You are given an array colors, in which there are three colors: 1, 2 and 3.
** You are also given some queries. Each query consists of two integers i and c,
** return the shortest distance between the given index i and the target color c. If there is no solution return -1.
Constraints:
	1 <= colors.length <= 5*10^4
	1 <= colors[i] <= 3
	1 <= queries.length <= 5*10^4
	queries[i].length == 2
	0 <= queries[i][0] < colors.length
	1 <= queries[i][1] <= 3
*/
/* 通用做法是二分，但是这里使用DP思路
** 题目限制颜色数量最大为 3
*/
func shortestDistanceColor(colors []int, queries [][]int) []int {
	n := len(colors)
	dp_left := make([][3]int, n)
	dp_right := make([][3]int, n)
	// 初始化
	for i := range dp_left{
		for j := range dp_left[i]{
			dp_left[i][j] = -1
		}
	}
	for i := range dp_right{
		for j := range dp_right[i]{
			dp_right[i][j] = -1
		}
	}
	dp_left[0][colors[0]-1] = 0
	dp_right[n-1][colors[n-1]-1] = 0
	for i := 1; i < n; i++{
		// dp[i] = dp[i-1]+1
		for j := range dp_left[i]{
			if dp_left[i-1][j] == -1{
				dp_left[i][j] = -1
			}else{
				dp_left[i][j] = dp_left[i-1][j] + 1
			}
		}
		dp_left[i][colors[i]-1] = 0
	}
	for i := n-2; i >= 0; i--{
		// dp[i] = dp[i+1]+1
		for j := range dp_right[i]{
			if dp_right[i+1][j] == -1{
				dp_right[i][j] = -1
			}else{
				dp_right[i][j] = dp_right[i+1][j] + 1
			}
		}
		dp_right[i][colors[i]-1] = 0
	}
	ans := []int{}
	for i := range queries{
		idx, color := queries[i][0], queries[i][1]
		// 忽略了 初始值为-1情况
		//ans = append(ans, min(dp_left[idx][color-1], dp_right[idx][color-1]))
		if dp_left[idx][color-1] == -1{
			ans = append(ans, dp_right[idx][color-1])
		}else if dp_right[idx][color-1] == -1{
			ans = append(ans, dp_left[idx][color-1])
		}else {
			ans = append(ans, min(dp_left[idx][color-1], dp_right[idx][color-1]))
		}
	}
	return ans
}
// 官方的写法，有 2个 编码 可以学习借鉴
func shortestDistanceColorDP(colors []int, queries [][]int) []int {
	n := len(colors)
	dp_left := make([][3]int, n)
	dp_right := make([][3]int, n)
	// 初始化
	for i := range dp_left{
		for j := range dp_left[i]{
			dp_left[i][j] = -1
		}
	}
	for i := range dp_right{
		for j := range dp_right[i]{
			dp_right[i][j] = -1
		}
	}
	dp_left[0][colors[0]-1] = 0
	dp_right[n-1][colors[n-1]-1] = 0
	for i := 1; i < n; i++{
		// dp[i] = dp[i-1]+1
		// 编码优化-1: 默认值已经是 -1
		for j := range dp_left[i]{
			if dp_left[i-1][j] != -1{
				dp_left[i][j] = dp_left[i-1][j] + 1
			}
		}
		dp_left[i][colors[i]-1] = 0
	}
	for i := n-2; i >= 0; i--{
		// dp[i] = dp[i+1]+1
		// 编码优化-1: 默认值已经是 -1
		for j := range dp_right[i]{
			if dp_right[i+1][j] != -1{
				dp_right[i][j] = dp_right[i+1][j] + 1
			}
		}
		dp_right[i][colors[i]-1] = 0
	}
	ans := []int{}
	for i := range queries{
		idx, color := queries[i][0], queries[i][1]
		// 忽略了 初始值为-1情况
		//ans = append(ans, min(dp_left[idx][color-1], dp_right[idx][color-1]))
		/*
		if dp_left[idx][color-1] == -1{
			ans = append(ans, dp_right[idx][color-1])
		}else if dp_right[idx][color-1] == -1{
			ans = append(ans, dp_left[idx][color-1])
		}else {
			ans = append(ans, min(dp_left[idx][color-1], dp_right[idx][color-1]))
		}
		 */
		// 编码优化-2:
		d := math.MaxInt32
		if dp_left[idx][color-1] != -1{
			d = min(d, dp_left[idx][color-1])
		}
		if dp_right[idx][color-1] != -1{
			d = min(d, dp_right[idx][color-1])
		}
		if d == math.MaxInt32{
			ans = append(ans, -1)
		}else{
			ans = append(ans, d)
		}
	}
	return ans
}

/* 1277. Count Square Submatrices with All Ones
**Given a m * n matrix of ones and zeros, return how many square submatrices have all ones.
 */
// dp[i][j]表示(i,j)为右下角的正方形的最大边长, 也表示以(i,j)为右下角的正方形的数目--- 这一点没有想到
// 在计算出所有的dp[i][j]后，将其累加，即为矩阵中正方形的数目
func countSquares(matrix [][]int) int {

}











