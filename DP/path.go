package DP

import "fmt"

/* 62. Unique Paths
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
The robot can only move either down or right at any point in time.
The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
How many possible unique paths are there?
Example 1:
	Input: m = 3, n = 7
	Output: 28
Example 2:
	Input: m = 3, n = 2
	Output: 3
	Explanation:
	From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
	1. Right -> Down -> Down
	2. Down -> Down -> Right
	3. Down -> Right -> Down
Example 3:
	Input: m = 7, n = 3
	Output: 28
Example 4:
	Input: m = 3, n = 3
	Output: 6
 */
/*
  dp[i][j] = dp[i-1][j] + dp[i][j-1]
 */
func UniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++{
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++{
		dp[0][i] = 1
	}
	for i := 1; i < m; i++{
		for j := 1; j < n; j++{
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

/* 63. Unique Paths II
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
The robot can only move either down or right at any point in time.
The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
Now consider if some obstacles are added to the grids. How many unique paths would there be?
An obstacle and space is marked as 1 and 0 respectively in the grid.
Example 1:
	Input: obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
	Output: 2
	Explanation: There is one obstacle in the middle of the 3x3 grid above.
	There are two ways to reach the bottom-right corner:
	1. Right -> Right -> Down -> Down
	2. Down -> Down -> Right -> Right
Example 2:
	Input: obstacleGrid = [[0,1],[0,0]]
	Output: 1
 */
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1{
		return 0
	}
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++{
		dp[i] = make([]int, n+1)
	}
	dp[1][1] = 1
	for i := 1; i <= m; i++{
		for j := 1; j <= n; j++{
			if i == 1 && j == 1{
				continue
			}
			if obstacleGrid[i-1][j-1] == 1 {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	//fmt.Println(dp)
	return dp[m][n]
}

// 参考官方： 转 一维
func UniquePathsWithObstaclesI(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([]int, n)
	if obstacleGrid[0][0] == 0 {
		dp[0] = 1
	}
	for i := 0; i < m; i++{
		for j := 0; j < n; j++{ // 单行计算
			if obstacleGrid[i][j] == 1{
				dp[j] = 0
				continue
			}
			if j - 1 >= 0 && obstacleGrid[i][j-1] == 0{
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[len(dp) - 1]
}

/* 980. Unique Paths III
You are given an m x n integer array grid where grid[i][j] could be:
1 representing the starting square. There is exactly one starting square.
2 representing the ending square. There is exactly one ending square.
0 representing empty squares we can walk over.
-1 representing obstacles that we cannot walk over.
Return the number of 4-directional walks from the starting square to the ending square,
that walk over every non-obstacle square exactly once.
Example 1:
	Input: grid = [[1,0,0,0],[0,0,0,0],[0,0,2,-1]]
	Output: 2
	Explanation: We have the following two paths:
	1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
	2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2)
Example 2:
	Input: grid = [[1,0,0,0],[0,0,0,0],[0,0,0,2]]
	Output: 4
	Explanation: We have the following four paths:
	1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2),(2,3)
	2. (0,0),(0,1),(1,1),(1,0),(2,0),(2,1),(2,2),(1,2),(0,2),(0,3),(1,3),(2,3)
	3. (0,0),(1,0),(2,0),(2,1),(2,2),(1,2),(1,1),(0,1),(0,2),(0,3),(1,3),(2,3)
	4. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2),(2,3)
Example 3:
	Input: grid = [[0,1],[2,0]]
	Output: 0
	Explanation: There is no path that walks over every empty square exactly once.
	Note that the starting and ending square can be anywhere in the grid.
 */
func UniquePathsIII(grid [][]int) int {
	row := len(grid)
	ans := 0
	cnt := 1
	if row == 0{
		return ans
	}
	col := len(grid[0])
	vis := make([][]bool, row)
	for i := range vis{
		vis[i] = make([]bool, col)
	}
	var dfs func(i, j int)
	dfs = func(i, j int){
		if i < 0 || j < 0 || i >= row || j >= col{
			return
		}
		if grid[i][j] == 2 && cnt == row*col{
			fmt.Println(vis)
			ans++
			return
		}
		if j > 0 && !vis[i][j-1]{
			if grid[i][j-1] != -1{
				vis[i][j-1] = true
				dfs(i, j-1)
				vis[i][j-1] = false
			}
			cnt++
		}
		if j < col-1 && !vis[i][j+1]{
			if grid[i][j+1] != -1 {
				vis[i][j+1] = true
				dfs(i, j+1)
				vis[i][j+1] = false
			}
			cnt++
		}
		if i > 0 && !vis[i-1][j]{
			if grid[i-1][j] != -1{
				vis[i-1][j] = true
				dfs(i-1, j)
				vis[i-1][j] = false
			}
			cnt++
		}
		if i < row - 1 && !vis[i+1][j]{
			if grid[i+1][j] != -1{
				vis[i+1][j] = true
				dfs(i+1, j)
				vis[i+1][j] = false
			}
			cnt++
		}
	}
	vis[0][0] = true
	dfs(0,0)
	return ans
}
