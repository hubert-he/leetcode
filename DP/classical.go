package DP

import (
	"sort"
)

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
func WaysToChangeDFS(amount int, coins []int) int {
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
func WaysToChangeDFSDP(amount int, coins []int) int {
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

func WaysToChangeDFSDP2(amount int, coins []int) int {
	return 0
}

func WaysToChangeDP(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins{
		for i := coin; i <= amount; i++{
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}