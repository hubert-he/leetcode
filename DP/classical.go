package DP

import (
	"math"
	"sort"
)

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