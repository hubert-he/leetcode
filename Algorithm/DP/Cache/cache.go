package Cache

import "math"

/* DFS 缓存
** 254. Factor Combinations
** Numbers can be regarded as the product of their factors.
	For example, 8 = 2 x 2 x 2 = 2 x 4.
	Given an integer n, return all possible combinations of its factors. You may return the answer in any order.
	Note that the factors should be in the range [2, n - 1].
*/
/* 组合题目： 1. 画出递归树， 2. 按顺序可消除重复组合 3. 增加cache记事本，跳过重复计算值
** 此题无法应用 cache
** 剪枝点：
	1. 为了避免重复，没必要从1开始遍历，而是从上一次的mulNum开始遍历，这样保证mulNum后续dfs的过程是递增的，所以不会出现重复。
	2. 遍历终点没必要为num， 而是num的开根号， 因此最大情况2^32的开根号结果为2^16次方=65536，是可接受范围。
*/
func getFactors(n int) [][]int {
	var dfs func(num, start int) [][]int
	dfs = func(num, start int)(ret [][]int){
		end := 1
		for i := 2; i*i <= num; i++{// 易错点-1 加等号 可以优化为2分
			end = i
		}
		for i := start; i <= end; i++{
			if num % i != 0{
				continue
			}
			ret = append(ret, []int{i, num/i})
			others := dfs(num/i, i)
			for j := range others{
				ret = append(ret, append([]int{i}, others[j]...))
			}
		}
		return ret
	}
	return dfs(n, 2)
}

/* 139. Word Break
Given a string s and a dictionary of strings wordDict,
return true if s can be segmented into a space-separated sequence of one or more dictionary words.
Note that the same word in the dictionary may be reused multiple times in the segmentation.
*/

func WordBreak_BFS(s string, wordDict []string) bool {

}

/* DFS 回溯：
** 1. 如果指针左侧在单词列表中，则对剩余的子串递归
** 2. 如果指针左侧不在单词列表中，返回，回溯
*/
// 借用map 作为 cache
func WordBreak_DFS(s string, wordDict []string) bool {
	m := map[string]bool{}
	for i := range wordDict{
		m[wordDict[i]] = true
	}
	var dfs func(str string)bool
	dfs = func(str string)bool{
		cache, ok := m[str]
		if ok {
			return cache
		}
		n := len(str)
		for i := 1; i < n; i++{
			if m[str[:i]] && dfs(str[i:]){
				m[str] = true
				return true
			}
		}
		m[str] = false
		return false
	}
	return dfs(s)
}
/* 直接DP，此题还是完全背包问题，并且是有序的
** 问题分解：将大的问题 拆解为 多个 小问题
** 1. 前 i 个字符的子串，能否分解成单词
** 2. 剩余子串，是否为单个单词
** 状态定义：dp[i]: 长度为 i 的s[:i]子串是否能拆分为单词
** 状态转移： dp[i] = dp[j] && s[j:i]能否构成单词
*/
func WordBreak_DP(s string, wordDict []string) bool {
	m := map[string]bool{}
	for i := range wordDict{
		m[wordDict[i]] = true
	}
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++{
		for j := 0; j < i; j++{
			if dp[j] && m[s[j:i]]{
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
// 依据前2个WordBreak_DFS 和 WordBreak_DP 可以使用起始字符作为标识符，进行缓存优化
func WordBreak_DFS2(s string, wordDict []string) bool {
	m := map[string]bool{}
	for i := range wordDict {
		m[wordDict[i]] = true
	}
	n := len(s)
	// cache 有 3 种可能情况， true false 未定义
	// 可以使用 数组 或 map 来处理
	// 数组： dp := make([]int, n), 0, 1, 2 表示3种情况
	// map:  dp := map[int]bool{}
	dp := map[int]bool{}
	var dfs func(start int) bool
	dfs = func(start int)bool{
		if start >= n{ // 易漏点-1
			return true
		}
		c, ok := dp[start]
		if ok {
			return c
		}
		for i := start+1; i <= n; i++{// 易漏点-2： i < n
			if m[s[start:i]] && dfs(i){ // 易漏点-3：m[s[:i]]
				dp[start] = true
				return true
			}
		}
		dp[start] = false
		return false
	}
	return dfs(0)
}

/* 787. Cheapest Flights Within K Stops
** There are n cities connected by some number of flights.
** You are given an array flights where
	flights[i] = [fromi, toi, pricei] indicates that there is a flight from city fromi to city toi with cost pricei.
** You are also given three integers src, dst, and k, return the cheapest price from src to dst with at most k stops.
** If there is no such route, return -1.
 */
// 2022-04-13 使用DFS 重刷此题，特殊例子会 TLE
// 特殊情况即 K 很大，但是 图中 src dst 是无法连通的情况，此时TLE
// 所以改进：加入了 提前判断图的两点是否会 连通的检查
func findCheapestPrice_DFS(n int, flights [][]int, src int, dst int, k int) int {
	ans := math.MaxInt32
	vis := make([]bool, n)
	g := make([][]int, n)
	for i := range g{
		g[i] = make([]int, n)
	}
	for i := range flights{
		s, d, w := flights[i][0], flights[i][1], flights[i][2]
		g[s][d] = w
	}
	//TLE改成策略 —— 连通性检查：排除不能连通的情况
	var connected func(node int)bool
	connected = func(node int)bool{
		if node == dst{ return true }
		vis[node] = true
		for d, w := range g[node]{
			if w != 0 && !vis[d]{
				if connected(d) {
					return true
				}
			}
		}
		return false
	}
	if !connected(src) {
		return -1
	}
	var dfs func(int, int, int)
	dfs = func(node, cost, stop int){
		if node == dst{
			if cost < ans{
				ans = cost
			}
			return
		}
		if stop > k { return }
		if cost >= ans{// 剪枝
			return
		}
		for d, w := range g[node]{
			if w != 0 && d != node{
				//fmt.Println(node, d, w, stop)
				dfs(d, cost+w, stop+1)
			}
		}
	}
	dfs(src, 0, 0)
	if ans == math.MaxInt32{ return -1 }
	return ans
}

/* DP：这里有2种维度的DP
** 一维：
** 二维：
 */
// 一维DP
func findCheapestPrice_DP_1(n int, flights [][]int, src int, dst int, k int) int {
	ans := math.MaxInt32
	g := make([][]int, n)
	for i := range g{
		g[i] = make([]int, n)
	}
	for i := range flights{
		s, d, w := flights[i][0], flights[i][1], flights[i][2]
		g[s][d] = w
	}

	if ans == math.MaxInt32{ return -1 }
	return ans
}
/* 2维DP: 这个 DP 思维方式很有代表性！
** 状态定义： dp[t][i] 表示 通过恰好 t 次航班，从出发成熟src 到 城市 i 需要的最小花费。
** 状态方程：枚举最后一次航班的起点 j
			dp[t][i] = min( dp[t-1][j] + cost(j,i)  )  (j, i 在给定的航班数组flights 中存在从城市 j 出发到达城市 i 的航班)
** 该状态转移方程的意义在于，枚举最后一次航班的起点 j ，
** 那么前 t - 1 次航班的最小花费为 dp[t-1][j] 加上最后一次航班的花费cost(j,i)中的最小值，即为 dp[t][i]
** 由于我们最多只能中转 k 次，也就是最多搭乘 k+1 次航班，最终的答案即为
				min( dp[1][dst], dp[2][dst], ..., dp[k+1][dst] )
** 初始值：t = 0时候，dp[t][i]表示不搭乘航班到达城市 i 的最小花费，因此有  dp[t][src] = 0, 非src 为 无穷大
 */
func findCheapestPrice_DP_2(n int, flights [][]int, src int, dst int, k int) int {
	ans := math.MaxInt32
	g := make([][]int, n)
	for i := range g{
		g[i] = make([]int, n)
	}
	for i := range flights{
		s, d, w := flights[i][0], flights[i][1], flights[i][2]
		g[s][d] = w
	}
	dp := make([][]int, k+2)
	for i := range dp{
		dp[i] = make([]int, n)
		for j := range dp[i]{
			dp[i][j] = math.MaxInt32
		}
	}
	// dp[0][0] = 0  低级错误
	dp[0][src] = 0
	for t := 1; t <= k+1; t++{
		for _, flight := range flights{
			j, i, cost := flight[0], flight[1], flight[2]
			if dp[t][i] > dp[t-1][j] + cost{
				dp[t][i] = dp[t-1][j] + cost
			}
		}
	}
	for t := 1; t <= k+1; t++{
		if ans > dp[t][dst]{
			ans = dp[t][dst]
		}
	}
	if ans == math.MaxInt32{ return -1 }
	return ans
}

func findCheapestPrice_DFS_DP_2(n int, flights [][]int, src int, dst int, k int) int {
	ans := math.MaxInt32
	g := make([][]int, n)
	for i := range g{
		g[i] = make([]int, n)
	}
	for i := range flights{
		s, d, w := flights[i][0], flights[i][1], flights[i][2]
		g[s][d] = w
	}
	dp := make([][]int, k+1)
	for i := range dp{
		dp[i] = make([]int, n)
	}
	for i := range dp[0]{
		dp[0][i] = math.MaxInt32
	}
	dp[0][0] = 0
	var dfs func(city, stop int)int
	dfs = func(city, stop int)(cost int){
		if city == dst { return 0 }
		if stop > k { return math.MaxInt32 }
		stop += 1
		if dp[city][stop] != 0{
			return dp[city][stop]
		}
		ret := math.MaxInt32

		dp[city][stop] = ret
		return ret
	}
	ans = dfs(src, 0)
	if ans == math.MaxInt32{ return -1 }
	return ans
}