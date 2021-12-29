package Cache
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



