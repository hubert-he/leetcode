package linear

import "math"

/* 区间 DP */
/* 1130. Minimum Cost Tree From Leaf Values
Given an array arr of positive integers, consider all binary trees such that:
	Each node has either 0 or 2 children;
	The values of arr correspond to the values of each leaf in an in-order traversal of the tree.
	The value of each non-leaf node is equal to the product of the largest leaf value in its left and right subtree, respectively.
Among all possible binary trees considered, return the smallest possible sum of the values of each non-leaf node.
It is guaranteed this sum fits into a 32-bit integer.
A node is a leaf if and only if it has zero children.
*/
func mctFromLeafValues_DFS(arr []int) int {
	n := len(arr)
	dp := map[[2]int][]int{}
	var dfs func(start, end int)[]int
	dfs = func(start, end int)[]int{
		length := end - start
		if length == 0{
			return []int{0,0}
		}
		if length == 1{
			return []int{0, arr[start]}
		}
		if dp[[2]int{start, end}] != nil {
			return dp[[2]int{start, end}]
		}
		ret := []int{math.MaxInt32, math.MinInt32}
		for i := 1; i < length; i++{
			l, r := dfs(start, start+i), dfs(start+i, end)
			t := l[0] + r[0] + l[1]*r[1]
			ret[0] = min(ret[0], t)
			ret[1] = l[1]
			if ret[1] < r[1]{
				ret[1] = r[1]
			}
		}
		dp[[2]int{start, end}] = ret
		return ret
	}
	ans := dfs(0, n)
	return ans[0]
}
/* 将dfs的自顶向下变为自底向上， 从小枚举区间长度
** 定义状态：
	dp[i][j] 表示将 arr[i...j] 合并之后所得的最小乘积之和。
	dp[0][len] 就是最终答案
** 状态转移方程：
	dp[i][j] = Math.min(dp[i][j], dp[i][k] + dp[k+1][j] + maxVal[i][k] * maxVal[k+1][j])
** 注意这里：
	我们可以将 arr[i...j] 分开，分为 arr[i...k] 和 arr[k+1...j]，
	arr[i...k] 和 arr[k+1...j] 这两个区间最终可以合并为两个数，再将这两个数合并。
	我们需要不断枚举 k，使得得到的 dp[i][j] 最小
*/
func mctFromLeafValues_DP(arr []int) int {
	// 这边改用 二维数组来替代map
	n := len(arr)
	dp := make([][][2]int, n)
	for i := range dp{
		dp[i] = make([][2]int, n)
		// 注意： 初始化
		//for j := range dp[i]{
		for j := i; j < n; j++{
			//dp[i][j] = [2]int{math.MaxInt32, 0}
			dp[i][j] = [2]int{0, max(arr[i:j+1]...)}
		}
	}
	for i := 1; i < n; i++{// 步长
		for l := 0; l + i < n; l++{
			r := l+i // 右侧
			t := [2]int{math.MaxInt32, math.MinInt32}
			for k := l; k < r; k++{
				t[0] = min(t[0], dp[l][k][0] + dp[k+1][r][0] + dp[l][k][1] * dp[k+1][r][1])
				t[1] = max(t[1], dp[l][k][1], dp[k+1][r][1])
			}
			dp[l][r] = t
		}
	}
	return dp[0][n-1][0]
}