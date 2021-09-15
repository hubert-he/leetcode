package DP

import (
	"math"
	"sort"
	"strings"
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
/* 474. Ones and Zeroes
You are given an array of binary strings strs and two integers m and n.
Return the size of the largest subset of strs such that there are at most m 0's and n 1's in the subset.
A set x is a subset of a set y if all elements of x are also elements of y.
 */
/* 与0 1 背包问题不同的是 这里的容量有2种，即选取的字符串子集中的 0 和 1 的数量上限。
** 经典的背包问题可以使用二维动态规划求解，两个维度分别是物品和容量。
** 这道题有两种容量，因此需要使用三维动态规划求解，三个维度分别是字符串、0 的容量和 1 的容量。
** i 为当前字符串，
** 记 zeros 为 strs[i]中 0 的个数；ones 为 strs[i]中 1 的个数
** 当0和1的容量分别为j和k时，考虑2种情况：
** 如果 j < zeros || k < ones， 则不能选择第 i 个字符串，此时有 dp[i][j][k] = dp[i-1][j][k]
** 如果 j >= zeros && k >= ones, 则取下面2项的最大值。
**		<1> 如果不选择第 i 个字符串，有 dp[i][j][k] = dp[i-1][j][k]
**		<2> 如果选择第 i 个字符串， 有 dp[i][j][k] = dp[i-1][j-zeros][k-ones] + 1
 */
func FindMaxForm(strs []string, m int, n int) int {
	total := len(strs)
	cnt := map[string][2]int{}
	for i := range strs{
		zero, one := 0, 0
		for j := range strs[i]{
			if strs[i][j] == '0'{
				zero++
			}else{
				one++
			}
		}
		cnt[strs[i]] = [2]int{zero, one}
	}
	dp := make([][][]int, 2)
	for i := range dp{
		dp[i] = make([][]int, m+1)
		for j := range dp[i]{
			dp[i][j] = make([]int, n+1)
		}
	}
	t0, t1 := cnt[strs[0]][0], cnt[strs[0]][1]
	// 对应上面的2种情况，如果只有1个字符串，如果 i < zeros 或者 j < ones 则无法选择此字符串，
	// 因为无法满足题目要求的至多i个 0 和 至多 j 个 1的要求
	for i := 0; i <= m; i++{// 初始化，
		for j := 0; j <= n; j++{
			if i >= t0 && j >= t1{
				dp[0][i][j] = 1
			}
		}
	}
	for i := 1; i < total; i++{
		t0, t1 := cnt[strs[i]][0], cnt[strs[i]][1]
		for j := 0; j <= m; j++{
			for k := 0; k <= n; k++{
				no := dp[(i-1)&1][j][k]
				yes := 0
				if j >= t0 && k >= t1{
					yes = dp[(i-1)&1][j-t0][k-t1] + 1
				}
				dp[i&1][j][k] = max(no, yes)
			}
		}
	}
	return dp[(total-1)&1][m][n]
}
// 优化：空间降维度，另使用strings.Count 来统计0 和 1 的个数
func FindMaxForm2(strs []string, m int, n int) int {
	total := len(strs)
	dp := make([][]int, m+1)
	for i := range dp{
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < total; i++{
		// zeros, ones := strings.Count(strs[i],"0"), strings.Count(strs[i], "1")
		zeros := strings.Count(strs[i], "0")
		ones := len(strs[i]) - zeros // 优化1
		/* 优化-2： 将if j >= zeros && k >= ones 放到 for 循环中  缩减代码行数
		for j := m; j >= 0; j--{
			for k := n; k >= 0; k--{
				no := dp[j][k]
				yes := 0
				if j >= zeros && k >= ones{
					yes = dp[j-zeros][k-ones] + 1
				}
				dp[j][k] = max(yes, no)
			}
		}
		 */
		for j := m; j >= zeros; j-- {
			for k := n; k >= ones; k-- {
				dp[j][k] = max(dp[j][k], dp[j-zeros][k-ones]+1)
			}
		}
	}
	return dp[m][n]
}

/* 494. Target Sum
You are given an integer array nums and an integer target.
You want to build an expression out of nums by adding one of the symbols '+' and '-' before each integer in nums and then concatenate all the integers.
For example, if nums = [2, 1], you can add a '+' before 2 and a '-' before 1 and concatenate them to build the expression "+2-1".
Return the number of different expressions that you can build, which evaluates to target.
 */
/* 失败原因 公式写错： target = sum - neg 错误
** 记数组的元素和为 sum，添加 - 号的元素之和为 neg，则其余添加 + 的元素之和为 sum−neg，
** 得到的表达式的结果为： (sum - neg) - neg = target ==> neg = (sum - target) / 2
** 由于数组 nums 中的元素都是非负整数，neg 也必须是非负整数。
** 所以上式成立的前提是 sum−target 是非负偶数。若不符合该条件可直接返回 0
** 若上式成立，问题转化成在数组 nums 中选取若干元素，使得这些元素之和等于 neg，计算选取元素的方案数。
** 使用动态规划的方法求解
 */
/* ❌
func FindTargetSumWays(nums []int, target int) int {
	total := 0
	for i := range nums{
		total += nums[i]
	}
	dp := make([]int, total+1)
	dp[0], dp[nums[0]] = 1,1
	for i := 1; i < len(nums); i++{
		for j := total; j >= nums[i]; j--{
			no := dp[j]
			yes := dp[j - nums[i]] + 1
			dp[j] = no + yes
		}
	}
	return dp[total-target]
}
 */
func FindTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := range nums{
		sum += nums[i]
	}
	diff := sum - target // 有负数可能
	if diff < 0 || diff & 1 == 1{ // 奇数无效
		return 0
	}
	neg := diff >> 1
	dp := make([]int, neg+1)
	dp[0] = 1 // 不选nums[0]
	for i := range nums{
		for j := neg; j >= nums[i]; j--{
			no := dp[j] // 不选择nums[i]
			yes := dp[j-nums[i]] //选 nums[i]
			dp[j] = no + yes
		}
	}
	return dp[neg]
}

/* 879. Profitable Schemes
There is a group of n members, and a list of various crimes they could commit.
The ith crime generates a profit[i] and requires group[i] members to participate in it.
If a member participates in one crime, that member can't participate in another crime.
Let's call a profitable scheme any subset of these crimes that generates at least minProfit profit,
and the total number of members participating in that subset of crimes is at most n.
Return the number of schemes that can be chosen. Since the answer may be very large, return it modulo 1000000007.
 */
/*
func ProfitableSchemes(n int, minProfit int, group []int, profit []int) int {
	sum := 0
	for i := range profit{
		sum += profit[i]
	}
	dp := make([][2]int, sum+1)
	dp[0][0], dp[0][1] = 1, 0
	for i := range profit{
		for j := sum; j >= profit[i]; j--{
			no := dp[j]
			cur := dp[j-profit[i]]
			yes := [2]int{}
			if (cur[1]+group[i]) <= n {
				yes[0] = cur[0]
				yes[1] = cur[1]+group[i]
			}
			dp[j][0], dp[j][1] = no[0]+yes[0], max(no[1], cur[1]+group[i])
		}
		fmt.Println(dp)
	}
	ans := 0
	for i := minProfit; i < sum+1; i++{
		ans += dp[i][0]
	}
	return ans
}
*/
/*
  将每个任务看作一个「物品」，完成任务所需要的人数看作「成本」，完成任务得到的利润看作「价值」。
  其特殊在于存在一维容量维度需要满足「不低于」，而不是常规的「不超过」。这需要我们对于某些状态作等价变换。
  定义 dp[i][j][k] 为考虑前 i 件物品，使用人数不超过 j，所得利润至少为 k 的方案数。
  对于每件物品（令下标从 1 开始），我们有「选」和「不选」两种决策：
  1. 不选，dp[i][j][k] = dp[i-1][j][k]
  2. 选，首先需要满足人数达到要求，即 j >= groups[i-1]，但是需要考虑「至少利润」负值问题：
	如果直接令「利润维度」为 k - profit[i - 1]可能会出现负值，那么负值是否为合法状态呢？
	这需要结合「状态定义」来看，由于是「利润至少为k」，因此属于「合法状态」，需要参与转移.
	由于我们没有设计动规数组存储「利润至少为负权」状态，我们需要根据「状态定义」做一个等价替换,
	将这个「状态」映射到 dp[i][j][0]。这主要是利用所有的任务利润都为「非负数」，所以不可能出现利润为负的情况，
	这时候「利润至少为某个负数 k」的方案数其实是完全等价于「利润至少为 0」的方案数。
  初始化：
	当不存在任何物品（任务）时，所得利用利润必然为 0（满足至少为 0），同时对人数限制没有要求
	dp[0][x][0] = 1
 */
func ProfitableSchemes(n int, minProfit int, group []int, profit []int) int {
	const mod int = 1e9+7 // 必须这样声明 否则默认float64
	dp := make([][][]int, 2)
	for i := range dp{
		dp[i] = make([][]int, n+1)
		for j := range dp[i]{
			dp[i][j] = make([]int, minProfit+1)
		}
	}
	dp[0][0][0] = 1
//	length := len(profit)
//	for i := 1; i < length; i++{
	for i := range profit{
		for j := 0; j <= n; j++{
			for k := 0; k <= minProfit; k++{
				no := dp[i&1][j][k]
				yes := 0
				if j >= group[i]{
					yes = dp[i&1][j-group[i]][max(0, k-profit[i])]
				}
				dp[(i+1)&1][j][k] = (yes + no) % mod
			}
		}
	}
	ans := 0
	for _, u := range dp[len(profit)&1]{
		ans = (ans + u[minProfit])%mod
	}
	return ans
}
/*
dp[i][j][k]仅与i-1有关，因此可降维优化
注意初始化情况：
对于最小工作利润为0的情况，无论当前在工作的员工有多少人，总能给出一种方案，即初始化为dp[i][0] = 1
此外，降维后 dp 数组的遍历顺序应为逆序，这样才能保证求dp[j][k]时，用到的dp[j-groups[i]][max(0, k-profit[i])] 是上一次的值。
正序遍历会改写。
 */
func ProfitableSchemesDP(n int, minProfit int, group []int, profit []int) int {
	const mod int = 1e9+7
	dp := make([][]int, n+1)
	// 注意降维度后的 初始化
	for i := range dp{
		dp[i] = make([]int, minProfit+1)
		dp[i][0] = 1
	}
	for i := range group{
		members, earn := group[i], profit[i]
		for j := n; j >= members; j-- {
			for k := minProfit; k >= 0; k-- {
				dp[j][k] = (dp[j][k] + dp[j-members][max(0, k-earn)]) % mod
			}
		}
	}
	// 注意结果求解与 3维的区别
	return dp[n][minProfit]
}


/**************  完全背包问题   **********************************************************************
题目:
	有 N 种物品和一个容量为 V 的背包，每种物品都有无限件可用。放入第 i 种物品 的费用（可以是体积也可以是重量）是 Ci，价值是 Wi。
	求解:将哪些物品装入背包，可使这些物品的耗费的费用总 和不超过背包容量，且价值总和最大。
约束条件：每种物品的数量为无限个，你可以选择任意数量的物品。
与01的区别：从每种物品的角度考虑，与它相关的策略已非选与不选两种情况，而是取 0件，1件，... 直至取 V/Ci 件等多种情况。
解决方法：套用01
	dp[i][j] 代表考虑前 i 件物品，放入一个容量为 j 的背包可以获得的最大价值
	dp[i][j] = max( dp[i-1][j], dp[i-1][j - k*Ci] + Wi * k ) 0 < k*Ci <= j
	复杂度与01一样有 O(VN)个状态需要求解，但求解每个状态的时间 已经不是常数了，求解的单个状态 dp[i][j]的时间是 O(V/Ci)
	因此总的复杂度 O(V*N*SUM(V/Ci))
优化方式-1：基于一个思路：任何情况下都可将价值小 而 费用高的物品 替换成 大价值低费用的 物品。
	若两件物品 i，j 满足 Ci < Cj && Wi >= Wj, 则可以将 物品j 直接去掉，不用考虑
	此方法不会改善最差情况的时间复杂度
	这个优化可以简单的O(N*N)实现，但是也可以借助计数排序降低到 O(V+N)
	首先将费用大于V 的 物品去掉， 然后使用类似计数排序的做法，计算出费用相同的物品中价值最高的哪个。
优化方式-2： 二进制思想，物品拆分成01背包问题
	考虑到第 i 种物品 最多选择 V / Ci 件，于是可以把第 i 种物品转化为 V / Ci 件 费用和价值均不变的物品，然后求解这个01背包问题
	此方式不会降低复杂度，但是提供了 完全背包转01背包问题思路，即 将一件物品拆成多件只能选0件或1件的01背包中的物品
	更高效的转化方法：把第 i 种物品拆成费用为 Ci*2^k 价值为 Wi * 2^k 的若干件物品，其中 k 取遍 满足 Ci*2^k <= V 的非负整数
	二进制思想： 不管 最优策略选 几件 第 i 种物品， 其件数写成二进制后， 总可以表示 若干个 2^k 件物品的和。
	这样一来就把每种物品拆成O(logV/Ci)件物品
优化方式-3： 最优解
	dp[i][j] = max( dp[i-1][j], dp[i-1][j-Ci]+Wi,  dp[i-1][j-2*Ci] + 2*Wi, ..., dp[i-1][j - k * Ci] + k*Wi ), 0 <= k * Ci <= j
	dp[i][j-Ci] = max( dp[i-1][j-Ci], dp[i-1][j-2*Ci]+Wi, dp[i-1][j-3*Ci]+2Wi, ..., dp[i-1][j - k * Ci] + (k-1)Wi ), 0 <= k * Ci <= j
    对比上面2个公式看出 dp[i][j] 与 dp[i][j-Ci] 部分情况总是相差 Wi， 将 dp[i][j-Ci] 等式 两边 加 Wi 即可规约dp[i][j]等式为下面情况：
	dp[i][j] = max( dp[i-1][j], dp[i][j-Ci]+Wi )
	最后通过降维消除 i 物品维度得出最后的方程：
	dp[j] = max( dp[j], dp[j - Ci] + Wi ), 注意与 01 背包区别 dp[i][j-Ci]+Wi
根本区别：
	1. 0-1 背包问题的状态转换方程是：dp[i][j] = max( dp[i-1][j], dp[i-1][j-Ci]+Wi )
		计算dp[i][j]依赖dp[i-1][j-Ci]，故在降维优化的时候，需要确保dp[j-Ci]存储的是上一行的值
		即确保dp[j-Ci]还没有被更新，所以比遍历方向需要时从大到下
	2. 完全背包问题的状态转移方程是： dp[i][j] = max( dp[i-1][j], dp[i][j-Ci]+Wi )
		计算dp[i][j]依赖dp[i][j-Ci]，故在降维时，需要确保 dp[j-Ci]存储的是当前行的值，
		即确保dp[j-Ci]已经被更新，所以遍历方向是从小到大。

 */

func UnboundedKnapsackProblem(N int, V int, c []int, w []int) int {
	dp := make([]int, V+1)
	for i := 0; i < N; i++{
		for j := 0; j <= V; j++{
			no := dp[j] // 不选此类物品， 选上一次的值
			yes := 0
			if j >= c[i] { // 可以选此类物品
				yes = dp[j - c[i]] + w[i]
			}
			dp[j] = max(no, yes)
		}
	}
	return dp[V]
}

/* 279. Perfect Squares
Given an integer n, return the least number of perfect square numbers that sum to n.
A perfect square is an integer that is the square of an integer; in other words, it is the product of some integer with itself.
For example, 1, 4, 9, and 16 are perfect squares while 3 and 11 are not.
限制条件： 1<= n <= 10000
贪心算法不适合
 */
func NumSquaresDP(n int) int {
	perfectNums := []int{}
	for i := 1; i * i <= n; i++{
		perfectNums = append(perfectNums, i*i)
	}
	dp := make([]int, n+1)
	/* 易错点-1 这里求的是最小， 因此dp初始化需要需要变更为最大值 */
	for i := range perfectNums{
		for j := 1; j <= n; j++{
			no := dp[j]
			//yes := 0 易错点-2 初始化值需要变更为最大值
			yes := math.MaxInt32
			if j >= perfectNums[i]{
				yes = dp[j-perfectNums[i]] + 1
			}
			dp[j] = min(no, yes)
		}
	}
	return dp[n]
}
/* 数学 -- 四平方和定理
	1. -- 任意一个正整数都可以被表示为至多四个正整数的平方和
	2. -- 当且仅当 n != 4^k * (8m + 7) 时， n可以被表示为至多三个正整数的平方和。
	      因此当 n == 4^k * (8m + 7)时， n只能被表示成 4 个正整数的平方和，可以直接返回4
当  n != 4^k * (8m + 7)， 情况可能为 1， 2， 3
answer=1： 则必有n 为完全平方数
answer=2： 则有 n = a^2 + b^2， 只需要枚举所有的 a ([1, 根号n]), 判断 n - a^2 是否为完全平方数即可
answer=3： 排除法。
*/

func NumSquaresMath(n int) int {
	checkAnswer4 := func(x int)bool{
		for x % 4 == 0{
			x /= 4
		}
		return x % 8 == 7
	}
	isPerfectSquare := func(x int)bool {
		y := int(math.Sqrt(float64(x)))
		return y*y == x
	}
	if isPerfectSquare(n) {
		return 1
	}
	if checkAnswer4(n) {
		return 4
	}
	for i := 1; i * i <= n; i++{
		j := n - i*i
		if isPerfectSquare(j) {
			return 2
		}
	}
	return 3
}

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
func WaysToChange(n int) int {
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

/* 518. Coin Change 2
	You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.
Return the number of combinations that make up that amount. If that amount of money cannot be made up by any combination of the coins, return 0.
You may assume that you have an infinite number of each kind of coin.
The answer is guaranteed to fit into a signed 32-bit integer.
Example 1:
	Input: amount = 5, coins = [1,2,5]
	Output: 4
	Explanation: there are four ways to make up the amount:
	5=5
	5=2+2+1
	5=2+1+1+1
	5=1+1+1+1+1
*/
//DFS会超时
func ChangeDFS(amount int, coins []int) int {
	sort.Ints(coins)
	ans := 0
	var dfs func(end int, left int)
	dfs = func(end int, left int){
		if left <= 0{
			if left == 0{
				ans++
			}
			return
		}
		for i := 0; i < end; i++{
			dfs(i+1, left - coins[i])
		}
	}
	dfs(len(coins), amount)
	return ans
}
// 2维
func ChangeDFSDP(amount int, coins []int) int {
	//sort.Ints(coins)
	ans, n := 0, len(coins)
	cache := make([][]int, n+1)
	for i := range cache{
		cache[i] = make([]int, amount+1)
		for j := range cache[i]{
			cache[i][j] = -1
		}
		cache[i][0] = 1
	}
	var dfs func(end int, left int)int
	dfs = func(end int, left int)int{
		if left < 0{
			return 0
		}
		if cache[end-1][left] != -1{
			return cache[end-1][left]
		}
		sum := 0
		for i := 0; i < end; i++{
			sum += dfs(i+1, left - coins[i])
		}
		cache[end-1][left] = sum
		return sum
	}
	ans = dfs(n, amount)
	return ans
}

func ChangeDP(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins{
		for i := coin; i <= amount; i++{
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
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