package knapsack

import (
	"math"
	"sort"
	"strings"
	"../../../tree"
	"../../../utils"
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
func KnapsackZeroONe1(N int, C int, v []int, w []int)int{
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
				dp[i][j] = utils.Max(dp[i-1][j], dp[i-1][j-v[i]] + w[i])
			}else{
				dp[i][j] = utils.Max(dp[i-1][j], 0)
			}
		}
	}
	return dp[N-1][C]
}
// dp[2][C+1] 解法 滚动数组
/* 只依赖上一层，故可用滚动数组优化空间复杂度
** 小技巧： i % 2 可以写为  i & 1
 */
func KnapsackZeroONe2(N int, C int, v []int, w []int)int{
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
				dp[i&1][j] = utils.Max(dp[(i-1)&1][j], dp[i-1][j-v[i]] + w[i])
			}else{
				dp[i&1][j] = utils.Max(dp[(i-1)&1][j], 0)
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
func KnapsackZeroONe3(N int, C int, v []int, w []int)int{
	dp := make([]int, C+1)
	for i := 0; i < N; i++{
		for j := C; j >= v[i]; j--{
			no := dp[j] // 不选择该 i 物品
			yes := dp[j-v[i]] + w[i] // 选择该 i 物品
			dp[j] = utils.Max(no, yes)
		}
	}
	return dp[C]
}

/* 1230. Toss Strange Coins
** You have some coins.  The i-th coin has a probability prob[i] of facing heads when tossed.
** Return the probability that the number of coins facing heads equals target if you toss every coin exactly once.
 */
/* 01背包变形问题
** 状态定义：
** 集合：从0到第 i 枚硬币，正面朝上个数等于 j 的概率
** 属性： 概率

** 状态转移：
** 抛第 i 枚硬币，要么朝上，要么朝下两种状态
** 上： dp[i][j] = dp[i-1][j-1]*p[i]
** 下： dp[i][j] = dp[i-1][j]*(1-p[i])
** 综合： dp[i][j] = dp[i-1][j-1]*p[i] + dp[i-1][j]*(1-p[i])
 */
func probabilityOfHeads(prob []float64, target int) float64 {
	n := len(prob)
	dp := [2][]float64{}
	for i := range dp{
		dp[i] = make([]float64, target+1)
	}
	// 初始化
	dp[0][0] = 1
	// dp[1][0], dp[1][1] = 1-prob[0], prob[0] <== 这个初始化方式不好，参考[0.5,0.5,0.5,0.5,0.5] target = 0 情况，dp[1][1]会初始化溢出
	for i := 1; i <= n; i++{
		dp[i%2][0] = dp[(i-1)%2][0]*(1-prob[i-1]) // 单独处理
		for j := 1; j <= i && j <= target; j++{
		//for j := 1; j <= target; j++{
			dp[i%2][j] = dp[(i-1)%2][j-1]*prob[i-1] + dp[(i-1)%2][j]*(1-prob[i-1])
		}
	}
	return dp[n%2][target]
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
			dp[j] = utils.Max(yes, no)
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

/* 1049. Last Stone Weight II
You are given an array of integers stones where stones[i] is the weight of the ith stone.
We are playing a game with the stones. On each turn, we choose any two stones（注意是任意2个） and smash them together.
Suppose the stones have weights x and y with x <= y. The result of this smash is:
If x == y, both stones are destroyed, and
If x != y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.
At the end of the game, there is at most one stone left.
Return the smallest possible weight of the left stone. If there are no stones left, return 0.
 */
/*题解参考链接：
	https://leetcode-cn.com/problems/last-stone-weight-ii/solution/gong-shui-san-xie-xiang-jie-wei-he-neng-jgxik/
	与leetcode-494 类似，需要考虑正负号两边时，其实只需要考虑一边就可以了，使用总和 sum 减去决策出来的结果，就能得到另外一边的结果。
	同时，由于想要「计算表达式」结果绝对值，因此我们需要将石子划分为差值最小的两个堆。
	其实就是对「计算表达式」中带 − 的数值提取公因数 −1，进一步转换为两堆石子相减总和，绝对值最小。
	这就将问题彻底切换为 01 背包问题：从 stones 数组中选择，凑成总和不超过 sum/2 的最大价值。
本质是 将stones 划分为2堆，然后相减所获得的差绝对值最小，即为合理分配
令石头的总重量为sum, ki=-1 的石头的重量之和为neg, 则其余 ki=1 的石头的重量之和为 sum - neg。则有：
	SUM{k[i] * stones[i]} = (sum - neg) - neg = sum - 2*neg
所以要使最后一块石头的重量尽可能小， neg 需要在不超过 sum/2的前提下尽可能的大。从而题目转换为01背包问题。
 */
func LastStoneWeightII(stones []int) int {
	sum := 0
	for i := range stones{
		sum += stones[i]
	}
	target := sum >> 1
	dp := make([]int, target+1)
	for i := range stones{
		for j := target; j >= stones[i]; j--{
			dp[j] = utils.Max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	return sum - 2 * dp[target]
}
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
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
还有一个特例：即背包中元素置放是要求有顺序的问题，后面的leetcode-139单词分割 DP 解法
	1. 如果是完全背包且不考虑元素之间顺序，物品个数放置在外循环（保证了物品按顺序遍历），容量要求target放在内循环，且内循环是正序
	2. 如果组合问题需要考虑元素物品放置的顺序，需要将要求target放在外循环，将物品放在内循环，且内循环正序。
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
			dp[j] = utils.Max(no, yes)
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
// 完全背包解决
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
			dp[j] = utils.Min(no, yes)
		}
	}
	return dp[n]
}
/* 2021-11-23 重刷此题
** 注题目要求 n = 0 为非法情况
*/
func numSquaresDP(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++{
		m := math.MaxInt32 // 避免初始化，优化时间
		for j := 1; j*j <= i; j++{
			m = utils.Min(m, dp[i-j*j])
		}
		dp[i] = m + 1
	}
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
				ret = utils.Min(ret, t+1)
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
				dp[j] = utils.Min(dp[j], dp[j-coins[i]] + 1)
			}
		}
	}
	if dp[amount] == math.MaxInt32{
		return -1
	}
	return dp[amount]
}
// 2021-11-2 重刷
func CoinChangeDP2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// dp[0] = 0
	for i := 1; i <= amount; i++{
		dp[i] = -1
	}
	for i := 1; i <= amount; i++{
		dp[i] = -1
		for _, c := range coins{
			if i < c{
				continue
			}
			if dp[i-c] != -1 && (dp[i] == -1 || dp[i] > dp[i-c]){
				dp[i] = dp[i-c]
			}
		}
		if dp[i] != -1{
			dp[i] += 1
		}
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
/* 139. Word Break
Given a string s and a dictionary of strings wordDict,
return true if s can be segmented into a space-separated sequence of one or more dictionary words.
Note that the same word in the dictionary may be reused multiple times in the segmentation.
*/
// 朴素 DFS_BFS 超时，参见testcase-1，故加DP cache
// 需要画出dfs 树图，找到重复计算的地方：一次DFS之后可能会计算出很多的结果
func WordBreakDFS(s string, wordDict []string) bool {
	n := len(s)
	m := map[string]bool{}
	for i := range wordDict{
		m[wordDict[i]] = true
	}
	cache := make([]int, n+1)
	for i := range cache{
		cache[i] = -1
	}
	var dfs func(i int)bool
	dfs = func(i int)bool{
		if i >= n{
			return true
		}
		if cache[i] == 0{
			return false
		}
		if cache[i] == 1{
			return true
		}
		for j := i; j < n; j++{
			if m[s[i:j+1]]{
				if dfs(j+1) {
					cache[i] = 1
					return true
				}
			}
		}
		cache[i] = 0
		return false
	}
	return dfs(0)
}
/* 转化为是否可以用 wordDict 中的词组合成 s，完全背包问题，并且为“考虑排列顺序的完全背包问题”，外层循环为 target，内层循环为选择池 wordDict
** dp[i] 表示以 i 结尾的字符串是否可以被 wordDict 中组合而成
 */
func WordBreakDP(s string, wordDict []string) bool {
	n := len(s)
	m := map[string]bool{}
	for i := range wordDict {
		m[wordDict[i]] = true
	}
	dp := make([]bool, n+1)
	dp[0] = true // 对于边界条件，我们定义 dp[0] = true 表示空串且合法
	// dp[i] = dp[i] || dp[i - word]  <== 还是对应选与不选这个单词物品
	for i := 1; i <= n; i++{
		for j := range wordDict{
			ws := len(wordDict[j])
			if i >= ws && m[s[i-ws:i]] {
				if dp[i]{
					break
				}
				dp[i] = dp[i] || dp[i-ws]
			}
		}
	}
	return dp[n]
}
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/* 多重背包问题-bounded knapsack problem
描述： 有 N 种物品和一个容量为 C 的背包，每种物品「数量有限」
      第i件物品的成本是 v[i]， 价值是 w[i]， 数量为 s[i]
	求解将哪些物品装入背包可使这些物品的耗费的成本总和不超过背包容量C，且价值总和最大
区别： 0-1背包基础上，增加了每件物品可以选择「有限次数」的特点（成本总和不超过背包容量C）
状态方程：因为对第 i 种物品 有 s[i]+1 种策略：即取0件，取1件... 取 s[i]件。
	设定定义： dp[i][j] 表示 前 i 种物品恰好放入一个容量为 j 的背包的最大价值，则由状态方程：
	dp[i][j] = max( dp[i-1][j], dp[i-1][j - 1*v[i]] + 1*w[i], .... ) =>
	dp[i][j] = max( dp[i-1][j - k * v[i]] + k * w[i] ) 0 <= k <= s[i]
	复杂度为 O(N*C*S) => N*S = SUM{s[i]} 0<=i<=N-1
 */
/*朴素解法*/
func MultiPackage(N int, C int, s []int, v []int, w []int)int{
	dp := [2][]int{}
	dp[0], dp[1] = make([]int, C+1), make([]int, C+1)
	// 初始化
	for i := 0; i <= C; i++{
		k := utils.Min(i / v[0], s[0])
		dp[0][i] = k * w[0]
	}
	for i := 1; i < N; i++{
		for j := 0; j <= C; j++{
			no := dp[(i-1)&1][j] // 不选 第 i 物品
			yes := 0 // 选择 第 i 个物品
			for k := 1; k <= s[i]; k++{
				if j < k*v[i]{
					break
				}
				y := dp[(i-1)&1][j-k*v[i]] + k*w[i]
				if yes < y{
					yes = y
				}
			}
			dp[i&1][j] = utils.Max(no, yes)
		}
	}
	return dp[(N-1)&1][C]
}
/* 一维优化
** dp[i][j] = max( dp[i-1][j - k * v[i]] + k * w[i] ) 0 <= k <= s[i]
** ==>
** dp[j] = max( dp[j - k * v[i]] + k * w[i] ) 0 <= k <= s[i]
*/
func MultiPackage1(N int, C int, s []int, v []int, w []int)int{
	dp := make([]int, C+1)
	dp[0] = 0
	for i := 1; i < N; i++{
		for j := C; j >= v[i]; j--{
			for k := 0; k <= s[i] && j >= k*v[i]; k++{
				t := dp[j - k*v[i]] + k*w[i]
				if dp[j] < t{
					dp[j] = t
				}
			}
		}
	}
	return dp[C]
}

/* bound knapsack 转换为0-1问题的二进制优化
** 扁平化方式是直接展开，一个数量为10的物品等效于 [1,1,1,1,1,1,1,1,1,1]
** 并没有减少运算量，但是如果我们能将 10 变成小于 10 个数，那么这样的「扁平化」就是有意义的
** 借鉴思路： 将原本数量为 n 的物品用 ceil(logn) 个数来代替，从而降低复杂度
** 类似例子：r：4 w:2 x: 1 linux 文件权限表示 通过3个数字 可表示8种情况  压缩了长度
** 伪码：
	// 源自背包九讲：F 为dp数组， C 为成本， W 为 价值， M 为物品数量, V 为背包容量
	def BoundKnapsackBinary(F, C, W, M)
		if C * M >= V
			CompletePack(F, C, W)  // 转换成完全背包
		k = 1
		while k < M
			ZeroOnePack(kC, kW)
			M = M - k
			k = 2*k
		ZeroOnePack(C*M, W*M)  // 转换成0-1背包
二进制优化的本质，是对「物品」做分类，使得总数量为 M 的物品能够用更小的 logM 个数所组合表示出来
*/
func BoundKnapsackBinary(N int, C int, s []int, v []int, w []int)int{
	// 重拆成多个物品
	worth, cost := []int{}, []int{}
	// 进行扁平化：如果一件物品规定的使用次数为 7 次，我们将其扁平化为三件物品：1*重量&1*价值、2*重量&2*价值、4*重量&4*价值
	// 三件物品都不选对应了我们使用该物品 0 次的情况、只选择第一件扁平物品对应使用该物品 1 次的情况、只选择第二件扁平物品对应使用该物品 2 次的情况，只选择第一件和第二件扁平物品对应了使用该物品 3 次的情况 ...
	for i := 0; i < N; i++{
		times := s[i]
		for k := 1; k <= times; k *= 2{
			times -= k
			worth = append(worth, k * w[i])
			cost = append(cost, k * v[i])
		}
		if times > 0{
			worth = append(worth, times * w[i])
			cost = append(cost, times * v[i])
		}
	}
	// 直接套用0-1背包
	dp := make([]int, C+1)
	for i := 0; i < len(worth); i++{
		for j := C; j >= cost[i]; j--{
			dp[j] = utils.Max(dp[j], dp[j - cost[i]] + worth[i])
		}
	}
	return dp[C]
}
/* 单调队列优化
单调队列优化，某种程度上也是利用「分类」实现优化。只不过不再是针对「物品」做分类，而是针对「状态」做分类。
在朴素的bound-Knapsack的解决方案中， 当我们在处理第 i 个物品 从 dp[0] 到 dp[C]的状态时，每次都是通过遍历当前容量 x 能够装多少件该物品，
然后从所有遍历结果中取最优。
但事实上，转移只会发生在「对当前物品体积取余相同」的状态之间
假设当前我们处理到的物品体积和价值均为 2， 数量s[i] = 3, 背包容量V为 10

 */
func BoundKnapsackQueue(N int, C int, s []int, v []int, w []int)int{
	dp := make([]int, C+1)
	g, q := make([]int, C+1), make([]int, 0, C+1)
	//
	for i := 0; i < N; i++{
		copy(g, dp) // 上次dp
		// 枚举余数
		for j := 0; j < v[i]; j++{
			// 枚举同一余数下，有多少种方案
			// 例如余数为1的情况下有： 1、v[i]+1、2*v[i]+1、3*v[i]+1
			for k := j; k <= C; k += v[i]{
				dp[k] = g[k]
				if len(q) != 0 && q[0] < k - s[i]*v[i]{
					// 将不在窗口范围内的 popup
					q = q[1:]
				}
				if len(q) != 0 && dp[k] < g[q[0]] + (k - q[0]){
					// 如果队列不为空 直接使用队头来更新
					dp[k] = g[q[0]] + (k - q[0])
				}
				// 当前值比队尾值更优，队尾元素直接出队
				for len(q) != 0 && g[q[len(q)-1]] - (q[len(q)-1] - j)/v[i]*w[i] <= g[k] - (k-j)/v[i]*w[i]{
					q = q[:len(q)-1]
				}
				q = append(q, k)
			}
		}
	}
	return dp[C]
}



////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/* 混合背包问题

 */

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/* 分组背包问题
** 给定N个物品组，和容量为C的背包
** 第i个物品组有 S[i]件物品，其中第 i 组的第 j 件物品的成本为 v[i][j]， 价值为 w[i][j]
** 每组有若干个物品，同一组内的物品最多只能选一个。
** 求 将哪些物品装入背包可使得这些物品的费用总和不超过背包的容量，且价值总和最大。
例题：
输入：N = 2, C = 9, S = [2, 3], v = [[1,2,-1],[1,2,3]], w = [[2,4,-1],[1,3,6]]
输出：10
 */
/* 解法：
** 定义dp[i][j] 为考虑前i个物品组，背包容量不超过j的最大价值
** 考虑条件设置：
**	由于每组有若干个物品，且每组「最多」选择一件物品
** 推导：即对于第 i 组而言，可决策的方案如下
** 不选择该组的任何物品： dp[i][j] = dp[i-1][j]
** 选该组第一个物品: dp[i][j] = dp[i-1][j-v[i][0]] + w[i][0]
** 选该组第 k 个物品：dp[i][j] = dp[i-1][j-v[i][k]] + w[i][k]
** 最终状态转移方程：
	dp[i][j] = max( dp[i-1][j], dp[i-1][j-v[i][k]] + w[i][k] ) 0 <= k < S[i]
 */
func GroupingKnapsack(N int, C int, S []int, v [][]int, w [][]int) int{
	dp := make([][]int, N+1)
	for i := range dp{
		dp[i] = make([]int, C+1)
	}
	// 初始化全为0，1. 不选任何物品情况下， 所有价值为 0
	// 2. 容量为0情况下， 所有价值为 0
	for i := 1; i <= N; i++{
		vi, wi := v[i-1], w[i-1] // 数组是从0开始计
		for j := 0; j <= C; j++{
			no := dp[i-1][j]
			yes := 0
			//for k := 0; j >= vi[k] && k < S[i-1]; k++{ // 注意 逻辑判断，k 溢出的问题
			for k := 0; k < S[i-1] && j >= vi[k]; k++{
				if yes < dp[i-1][j-vi[k]]+wi[k]{
					yes = dp[i-1][j-vi[k]]+wi[k]
				}
			}
			dp[i][j] = utils.Max(no, yes)
		}
	}
	//fmt.Println(dp)
	return dp[N][C]
}
// 1维 优化
func GroupingKnapsack1(N int, C int, S []int, v [][]int, w [][]int) int{
	dp := make([]int, C+1)
	dp[0] = 0
	for i := 1; i <= N; i++{
		vi, wi := v[i-1], w[i-1] // 数组是从0开始计
		for j := C; j >= 0; j--{
			// no := dp[i-1][j] 被优化在之前的dp[j]保存的数值中
			for k := 0; k < S[i-1] && j >= vi[k]; k++{
				dp[j] = utils.Max(dp[j], dp[j-vi[k]] + wi[k])
			}
		}
	}
	return dp[C]
}


/* 1155. Number of Dice Rolls With Target Sum
** You have d dice and each die has f faces numbered 1, 2, ..., f. You are given three integers d, f, and target.
** Return the number of possible ways (out of fd total ways) modulo 109 + 7 to roll the dice so the sum of the face-up numbers equals target.
 */
func NumRollsToTarget(d int, f int, target int) int {
	if d * f < target{
		return 0
	}
	const mod int = 1e9+7
	// 此处状态数组定义从 0 开始计算，即第0个表示题意中的第一个骰子
	dp := make([][]int, d)
	for i := range dp{
		dp[i] = make([]int, target+1)
		// 因为从 0 计算，所以初始化，必须从这行开始
		if i == 0{
			//for j := 1; j <= target; j++{ // 防止target > f 污染初始化值，超过的应该初始化为0 不存在这样的骰子方案
			for j := 1; j <= f && j <= target; j++{
				dp[i][j] = 1
			}
		}
	}
	for i := 1; i < d; i++ {
		for j := 0; j <= target ;j++{
			// 因为所有骰子都必须要选一个，因此不可能有不选这个状态，与01背包不同
			for k := 1; j >= k && k <= f; k++{
				dp[i][j] = (dp[i][j] + dp[i-1][j-k]) % mod
			}
		}
	}
	//fmt.Println(dp)
	return dp[d-1][target]%mod
}

func NumRollsToTarget2(d int, f int, target int) int {
	if d * f < target{
		return 0
	}
	const mod int = 1e9+7
	// 注意和上面解法的 状态的定义，从0开始计数，还是从 1 开始计算，导致初始化方向不同
	dp := make([][]int, d+1)
	for i := range dp{
		dp[i] = make([]int, target+1)
	}
	/* 根据定义，此方法是dp 大小 d+1 额外多了一行， 即
	** dp[0] 表示 任何骰子都不选，target = 0 的方案数为 1
	 */
	dp[0][0] = 1
	for i := 1; i <= d; i++ { // 枚举物品组（每个骰子）
		for j := 0; j <= target ;j++{ // 枚举背包容量（所掷得的总点数）
			for k := 1; j >= k && k <= f; k++{
				dp[i][j] = (dp[i][j] + dp[i-1][j-k]) % mod
			}
		}
	}
	//fmt.Println(dp)
	return dp[d][target]
}

func NumRollsToTarget1(d int, f int, target int) int {
	if d * f < target{
		return 0
	}
	const mod int = 1e9+7
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= d; i++ {
		for j := target; j >= 0 ;j--{
			dp[j] = 0 // 易错点：忽略初始化数据，导致历史数据叠加
			for k := 1; j >= k && k <= f; k++{
				dp[j] = (dp[j] + dp[j-k])%mod
			}
		}
	}
	//fmt.Println(dp)
	return dp[target]
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/* 多维背包问题
** 在背包九讲里，也指二维费用的背包问题，即
** 对于每件物品，具有2种不同的费用，选择这件物品必须同时付出这两笔费用。对于每种费用都有一个可付出的最大值（对应2种背包容量），
** 求如何选择物品可以得到最大的价值
** 设第i件物品所需的两种费用分别为 v[i] 和 u[i]。两种费用可付出的最大值（也即两种背包容量）分别为 V 和 U。物品价值 w[i]
** dp[i][v][u]: 表示前 i 件物品 付出的两种费用分别为 u 和 u 时可获得的最大价值
** dp[i][v][u] = max( dp[i-1][u][u], dp[i-1][v-v[i]][u-u[i]]+w[i] )
** 空间可以优化：降为2维的处理方法：
	1. 当每件物品只可取一次时，变量 u 和 u 采用逆序的循环，
	2. 当物品有如完全背包问题时采用顺序的循环，
	3. 当物品有如多重背包问题时拆分物品。
 */
//TODO: 增加一个代码
/* 还有一个变种： 物品总个数的限制
	二维费用的条件是以一种隐含的方式给出： 最多只能取 U 件物品。这相当于每件物品多了一种"件数"的费用，每个物品的件数费用均为 1，可以付出的
	的最大件数费用为 U。也就是说，设 F[v, u] 表示付出费用 v 最多选 u件时可得到的最大价值，则根据物品的类型（01 完全 多重）用不同的方法循环更新，
	最后在 F[0->V, 0->U]的范围内寻找答案。
 */

/* 474. Ones and Zeroes
You are given an array of binary strings strs and two integers m and n.
Return the size of the largest subset of strs such that there are at most m 0's and n 1's in the subset.
A set x is a subset of a set y if all elements of x are also elements of y.
*/
/* 与0 1 背包问题不同的是 这里的容量有2种，即选取的字符串子集中的 0 和 1 的数量上限。
** 经典的背包问题可以使用二维动态规划求解，两个维度分别是物品和容量。
** 这道题有两种容量，因此需要使用三维动态规划求解，三个维度分别是字符串、0 的容量和 1 的容量。
** dp[i][j][k]表示前i件物品，在数字 1 容量不超过 j， 数字 0 的容量不超过 k 的条件下的 「最大价值」（每个字符串的价值均为 1）
** dp[i][j][k] = max( dp[i-1][j][k], dp[i-1][i - cnt[k][0]][j - cnt[k][1]] + 1 )
** cnt数组记录的是字符串中出现的 0 1 数量。
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
				dp[i&1][j][k] = utils.Max(no, yes)
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
				dp[j][k] = utils.Max(dp[j][k], dp[j-zeros][k-ones]+1)
			}
		}
	}
	return dp[m][n]
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
  2. 选，首先需要满足人数达到要求，即 j >= groups[i-1]，
	但是需要考虑「至少利润」负值问题：
	如果直接令「利润维度」为 k - profit[i - 1]可能会出现负值，那么负值是否为合法状态呢？
	这需要结合「状态定义」来看，由于是「利润至少为k」，因此属于「合法状态」，需要参与转移.
	解决： 由于我们没有设计动规数组存储「利润至少为负权」状态，我们需要根据「状态定义」做一个等价替换,
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
					yes = dp[i&1][j-group[i]][utils.Max(0, k-profit[i])]
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
				dp[j][k] = (dp[j][k] + dp[j-members][utils.Max(0, k-earn)]) % mod
			}
		}
	}
	// 注意结果求解与 3维的区别
	return dp[n][minProfit]
}
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/* 有依赖的背包问题 又称 树形背包
问题描述：有 N 个物品和一个容量为 C 的背包，物品编号为 0 ... N-1
		物品之间具有依赖关系，且依赖关系组成一棵树的形状, 如果选择一个物品，则必须选择它的父节点.
		第 i 件物品的体积为 v[i]，价值为 w[i]，其父节点物品编号为 p[i]，其中根节点 p[i]=-1。
		求解将哪些物品装入背包，可使这些物品的总体积不超过背包容量，且总价值最大。
分组背包：dp[i][j]为考虑前 i 个物品组，背包容量不超过 j 的最大价值
	从状态定义我们发现，常规的分组背包问题对物品组的考虑是“线性“的（从前往后考虑每个物品组）。
	然后在状态转移时，由于物品组之间没有依赖关系，限制只发生在”组内“（每组「最多」选择一件物品）
	所以常规分组背包问题只需要采取「枚举物品组 - 枚举背包容量 - 枚举组内物品（决策）」 的方式进行求解即可
树形背包：在树形背包问题中，每个物品的决策与其父节点存在依赖关系，因此我们将”线性“的状态定义调整为”树形“的
	dp[u][j]为考虑以 u 为根的子树，背包容量不超过 j 的最大价值。
	限制条件： 首先，根据树形背包的题目限制，对于以 u 为根的子树，无论选择哪个节点，都需要先把父节点选上
			如果从选择节点 u 的哪些子树入手的话，我们发现节点 u 最坏情况下会有 100 个子节点，
			而每个子节点都有选和不选两种决策，因此总的方案数为 2的100次方，这显然是不可枚举的
转移方程：这时候可以从”已有维度“的层面上对方案进行划分，而不是单纯的处理每一个具体方案。
		最终这2的100次方的方案的最大价值会用于更新dp[u][j]，根据「容量」这个维度对这2的100次方个方案进行划分
		* 消耗容量为 0 的方案数的最大价值
		* 消耗容量为 1 的方案数的最大价值
		* 消耗容量为 j-v[u] 的方案数的最大价值
		消耗的容量的范围为[0, j-v[u]] ，是因为需要预留 v[u] 的容量选择当前的根节点u
	dp[u][j] = max( dp[u][j], dp[u][j-k]+dp[x][k] ) 0<= k <= j-v[u]
	x节点为 u 的子节点
	从状态转移方式发现，在计算 dp[u][j] 时需要用到 dp[x][k]，因此我们需要先递归处理节点 u 的子节点 x 的状态值
*/
func TreeKnapsack(N int, C int, p []int, v []int, w []int)int{
	//root := -1
	//dp := make([][]int, N)
	return 0

}

/*LCP 34. 二叉树染色
小扣有一个根结点为 root 的二叉树模型，初始所有结点均为白色，可以用蓝色染料给模型结点染色，模型的每个结点有一个 val 价值。
小扣出于美观考虑，希望最后二叉树上每个蓝色相连部分的结点个数不能超过 k 个，求所有染成蓝色的结点价值总和最大是多少？
*/
/* 思路不对
func PaintBinaryTree(root *Tree.BiTreeNode, k int) int {
	if root == nil {
		return 0
	}
	dp := map[*Tree.BiTreeNode][]int{}
	dp[nil] = make([]int, k+1)
	var dfs func(node *Tree.BiTreeNode)
	dfs = func(node *Tree.BiTreeNode) {
		if node == nil{
			return
		}
		dp[node] = make([]int, k+1)
		dfs(node.Left)
		dfs(node.Right)
		for i := k; i >= 2; i--{
			yes, no := 0, 0
			yes = max(dp[node.Left][i-1] + dp[node.Right][0], dp[node.Left][0] + dp[node.Right][i-1]) + node.Val.(int)
			for j := 0; j <= k; j++{
				no = max(no, dp[node.Left][j]+dp[node.Right][k-j])
			}
			dp[node][i] = max(yes, no)
		}
	}
	dfs(root)
	fmt.Println(dp)
	return dp[root][k]
}
 */
/* 自底向上dp，保存子节点的状态并计算当前是否染色的状态  一维树形DP
状态定义： dp[i]表示以该节点为根，相邻的子节点为蓝色的个数为 i 的情况下(包括自身), 节点价值总和的最大值。
状态转移： 当前节点为root，dp逻辑为
	1. root不染色，那么只要返回 dp[0]，其值为左、右子树染色或不染色的最大值之和
		dp[node][i] = max(dp[node.Left][j]) + max(dp[node.Right][j])  0<=j<=k
	2. root染色，那么就分左子树染色 j 个，右子树染色 i - 1 - j 个时，加上 root.val 的和。
		dp[node][i] = max(node.Val + dp[node.Left][j] + dp[node.Right][i-1 - j]), 0 <i<=k && 0 < j < i
 */
func PaintBinaryTree(root *Tree.BiTreeNode, k int) int {
	if root == nil {
		return 0
	}
	dp := map[*Tree.BiTreeNode][]int{}
	dp[nil] = make([]int, k+1)
	var dfs func(node *Tree.BiTreeNode)
	dfs = func(node *Tree.BiTreeNode) {
		if node == nil{
			return
		}
		dp[node] = make([]int, k+1)
		dfs(node.Left)
		dfs(node.Right)

		// root 不染色
	/* 换种写法
		maxLeft, maxRight := math.MinInt32, math.MinInt32
		for i := 0; i <= k; i++{
			maxLeft = max(maxLeft, dp[node.Left][i]) 	// 求出左子树可能的最大值
			maxRight = max(maxRight, dp[node.Right][i])	// 求出右子树可能的最大值
		}
		dp[node][0] = maxLeft + maxRight // root 不染色
	 */
		dp[node][0] = utils.Max(dp[node.Left]...) + utils.Max(dp[node.Right]...)
		// root 染色：左子树染色 j 个，右子树染色 i - 1 - j 个时，加上 root.val 的和
		// 需要求出 dp[root][i]的值
		for i := 1; i <=k; i++{
			for j := 0; j < i; j++{
				dp[node][i] = utils.Max(dp[node][i], dp[node.Left][j]+dp[node.Right][i-j-1]+node.Val.(int))
			}
		}

	}
	dfs(root)
	return utils.Max(dp[root]...) // 不超过
}
// 递归返回数组的方式
func PaintBinaryTree2(root *Tree.BiTreeNode, k int) int {
	nildp := make([]int, k+1)
	var dfs func(node *Tree.BiTreeNode) []int
	dfs = func(node *Tree.BiTreeNode)[]int{
		if node == nil{
			return nildp
		}
		left,right := dfs(node.Left), dfs(node.Right)
		dp := make([]int, k+1)
		// 不选root
		dp[0] = utils.Max(left...) + utils.Max(right...)
		// 选root
		for i := 1; i <= k; i++{
			for j := 0; j < i; j++{
				dp[i] = utils.Max(dp[i], left[j] + right[i-1 - j] + node.Val.(int))
			}
		}
		return dp
	}
	return utils.Max(dfs(root)...)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/* 背包问题求方案数

 */

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/* 求背包问题的方案

 */

/* 0 1 背包问题

 */

/* 343. Integer Break
Given an integer n, break it into the sum of k positive integers, where k >= 2, and maximize the product of those integers.

Return the maximum product you can get.
*/
/* 对于的正整数 n，当 n≥2 时，可以拆分成至少两个正整数的和
** 令 k 是拆分出的第一个正整数，则剩下的部分是 n−k，n−k 可以不继续拆分，或者继续拆分成至少两个正整数的和。
** 由于每个正整数对应的最大乘积取决于比它小的正整数对应的最大乘积，因此可以使用动态规划求解
** 状态方程：
** 创建数组 dp，其中 dp[i] 表示将正整数 i 拆分成至少两个正整数的和之后，这些正整数的最大乘积。
** 特别地，0 不是正整数，1 是最小的正整数，0 和 1 都不能拆分，因此 dp[0]=dp[1]=0
** 当 i >= 2 时， 假设对正整数 i 拆分出的第一个正整数是 j（1≤j<i），则有以下两种方案：
** 将 i 拆分成 j 和 i−j 的和，且 i−j 不再拆分成多个正整数，此时的乘积是 j×(i−j)；
** 将 i 拆分成 j 和 i−j 的和，且 i−j 继续拆分成多个正整数，此时的乘积是 j×dp[i−j]
** 状态转移：dp[i]=max{max(j×(i−j),j×dp[i−j])} j 属于[1,i)
 */
func IntegerBreak(n int) int {
	dp := make([]int, n+1)
	dp[0],dp[1] = 0, 0
	for i := 2; i <= n; i++{
		for j := 1; j < i; j++{
			dp[i] = utils.Max(dp[i], j*(i-j), j*dp[i-j])
		}
	}
	return dp[n]
}
/* dp[i]=max{max(j×(i−j),j×dp[i−j])} j 属于[1,i)
** 计算 dp[i] 时，j 的值遍历了从 1 到 i−1 的所有值，因此总时间复杂度是 O(n^2)
** 转移方程包含两项,当 j固定时， dp[i]的值由 j×(i−j),j×dp[i−j]较大值决定，因此需要分2项考虑
** j×dp[i−j] => 计算 dp[i] 的值只需要考虑 j=2 和 j=3 的情况,不需要遍历从 11 到 i-1i−1 的所有值
**
** j×(i−j) =>
** 如果j >= 4 则 dp[j] >= j 当且仅当 j = 4 时等号成立；
** 因此 i - j >= 4的情况下，有 dp[i-j] >= i-j 和 dp[i] >= j * dp[i-j] >= j *(i-j), 此时计算中不需要考虑 j*(i-j)的值
** 如果 i-j < 4, 计算dp[i]的值就需要考虑j*(i-j)的值；
** 1. j=1 则1*（i-1) = i-1
** 		当i = 2 或 i = 3 时有 dp[i] = i-1
**		当 i >= 4 有 dp[i] >= i >= i-1 显然 当 i>=4时 j = 1 不可能取到最大乘积，故 j=1 不需要考虑
** 2. j >= 4 dp[i] 是否可能等于 j * (i-j) ?
**	 当 i 固定，要是的j * (i-j)最大，j 的值应该取 j = i/2
** 	 j >= 4 若要满足 j = i/2, 则 i >= 8, 此时i-j >= 4 ==> dp[i-j] >= i-j ==> j*dp[i-j] >= j*(i-j)
** 	 由此可见，当 j >= 4 计算dp[i] 只需要考虑 j * dp[i-j], 不需要考虑 j*(i-j)
** 3. 另外，在使用 j * dp[i-j] 计算 dp[i] 时， j = 2 和 j = 3的情况 一定优于 j >= 4 的情况，
**	 因此，无论考虑 j * dp[i-j] 还是考虑 j * (i-j)， 都只需要考虑 j = 2 和 j = 3 的情况
** 特殊边界： n = 2  唯一拆分，最大乘积 1
** 当 n >= 3 转移方程为： dp[i] = max(2*(i-2), 2 * dp[i-2], 3 * dp[i-3], 3 * dp[i-3])
 */
func IntegerBreakII(n int) int {
	if n < 4{
		return n-1
	}
	dp := make([]int, n + 1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = utils.Max(utils.Max(2 * (i - 2), 2 * dp[i - 2]), utils.Max(3 * (i - 3), 3 * dp[i - 3]))
	}
	return dp[n]
}
/* 数学推论
** 1. 将数字 n 尽可能以因子 3 等分时，乘积最大
** 2. 若拆分的数量 a 确定， 则 各拆分数字相等时 ，乘积最大
** 拆分规则：
	最优： 3 。把数字 n 可能拆为多个因子 3 ，余数可能为 0,1,2 三种情况。
	次优： 2 。若余数为 2 ；则保留，不再拆为 1+1 。
	最差： 1 。若余数为 1 ；则应把一份 3+1 替换为 2+2，因为 2×2>3×1。
 */
func integerBreak3(n int) int {
	if n <= 3 {
		return n - 1
	}
	quotient := n / 3
	remainder := n % 3
	if remainder == 0 {
		return int(math.Pow(3, float64(quotient)))
	} else if remainder == 1 {
		return int(math.Pow(3, float64(quotient - 1))) * 4
	}
	return int(math.Pow(3, float64(quotient))) * 2
}