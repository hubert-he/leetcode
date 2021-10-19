package DP

import (
	"math"
)

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
	cnt := 0
	total := 1 // 额外增加 起始节点
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
		if grid[i][j] == 2 && cnt == total{
			ans++
			return
		}
		cnt++
		if j > 0 && !vis[i][j-1]{
			if grid[i][j-1] != -1{
				vis[i][j-1] = true
				dfs(i, j-1)
				vis[i][j-1] = false
			}
		}
		if j < col-1 && !vis[i][j+1]{
			if grid[i][j+1] != -1 {
				vis[i][j+1] = true
				dfs(i, j+1)
				vis[i][j+1] = false
			}
		}
		if i > 0 && !vis[i-1][j]{
			if grid[i-1][j] != -1{
				vis[i-1][j] = true
				dfs(i-1, j)
				vis[i-1][j] = false
			}
		}
		if i < row - 1 && !vis[i+1][j]{
			if grid[i+1][j] != -1{
				vis[i+1][j] = true
				dfs(i+1, j)
				vis[i+1][j] = false
			}
		}
		cnt--
	}
	start := [2]int{}
	for i := range grid{
		for j := range grid[i]{
			switch grid[i][j] {
			case 1:
				vis[i][j] = true
				start[0], start[1] = i, j
			case 0:
				total++
			}
		}
	}
	dfs(start[0],start[1])
	return ans
}
/*
  注意下面的coding技巧：
	1 是 提前用数组表示 方向
    2 bit运算标记是否已访问
 */
func UniquePathsIII1(grid [][]int) int {
	row := len(grid)
	ans := 0
	if row == 0{
		return ans
	}
	col := len(grid[0])
/*  优化-2
	cnt := 0
	total := 1 // 额外增加 起始节点
	vis := make([][]bool, row)
	for i := range vis{
		vis[i] = make([]bool, col)
	}
 */
	target := 0
	// 优化下面的代码行数-1
	dr := [4]int{0,-1,0,1}
	dc := [4]int{1,0,-1,0}
	code := func(r, c int)int{
		return 1 << uint(r * col + c)
	}
	var dfs func(i, j, todo int)
	dfs = func(i, j, todo int){
		if i < 0 || j < 0 || i >= row || j >= col{
			return
		}
		if grid[i][j] == 2 {
			if todo == 0{
				ans++
			}
			return
		}
		for k := 0; k < 4; k++{ // 代码行数较之前变少
			nr, nc := i + dr[k], j + dc[k]
			if nr >= 0 && nr < row && nc >= 0 && nc < col{
				if (code(nr, nc) & todo) != 0 { // 未访问
					dfs(nr, nc, todo ^ code(nr, nc)) // 直接计算，避免回溯 todo 恢复现场
				}
				/*
				if grid[nr][nc] != -1 && !vis[nr][nc]{
					vis[nr][nc] = true
					dfs(nr, nc)
					vis[nr][nc] = false
				}
				*/
			}
		}
	}
	start := [2]int{}
	for i := range grid{
		for j := range grid[i]{
			v := grid[i][j]
			if v % 2 == 0{
				target |= code(i, j)
			}else if v == 1{
				start[0], start[1] = i, j
			}
		}
	}
	dfs(start[0],start[1], target)
	return ans
}
/*
	定义 dp(r, c, todo) 为从 (r, c) 开始行走，还没有遍历的无障碍方格集合为 todo 的好路径的数量
	过记忆化状态 (r, c, todo) 的答案来避免重复搜索
 */
func UniquePathsIIIDP(grid [][]int) (ans int) {
	row := len(grid)
	if row == 0{
		return ans
	}
	col := len(grid[0])
	target := 0
	// 定义两个方向矩阵
	dr := [4]int{0,-1,0,1}
	dc := [4]int{1,0,-1,0}
	memo := make([][][]int, row)
	for i := range memo{
		memo[i] = make([][]int, col)
		for j := range memo[i]{
			memo[i][j] = make([]int, 1<< uint(row*col))
		}
	}
	code := func(r, c int)int{
		return 1 << uint(r * col + c)
		//return 1 << (r * row + c) 很容易出错
	}
	var dp func(i, j, todo int)int
	dp = func(r, c, todo int) int{
		//fmt.Printf("%b\n", todo)
		if memo[r][c][todo] != 0{
			return memo[r][c][todo]
		}
		if grid[r][c] == 2{
			if todo == 0{
				return 1
			}
			return 0
		}
		result := 0
		for k := 0; k < 4; k++{
			nr, nc := r + dr[k], c + dc[k]
			if nr >= 0 && nr < row && nc >= 0 && nc < col{
				if (code(nr, nc) & todo) != 0{ // 如果标0 则表示应被访问
					result += dp(nr, nc, todo ^ code(nr, nc))
				}
			}
		}
		memo[r][c][todo] = result
		return result
	}
	start := [2]int{}
	for i := range grid{
		for j := range grid[i]{
			v := grid[i][j]
			if v % 2 == 0{// 即grid[i][j]为0 或 2
				target |= code(i,j)
			} else if v == 1{
				start[0], start[1] = i, j
			}
		}
	}
	//fmt.Printf("==>%b %d %d\n", target, target, x)
	return dp(start[0],start[1], target)
}

/* 1575. Count All Possible Routes
You are given an array of distinct positive integers locations where locations[i] represents the position of city i.
You are also given integers start, finish and fuel representing the starting city, ending city, and the initial amount of fuel you have, respectively.
At each step, if you are at city i, you can pick any city j such that j != i and 0 <= j < locations.length and move to city j. 
Moving from city i to city j reduces the amount of fuel you have by |locations[i] - locations[j]|. Please notice that |x| denotes the absolute value of x.
Notice that fuel cannot become negative at any point in time, and that you are allowed to visit any city more than once (including start and finish).
Return the count of all possible routes from start to finish.
Since the answer may be too large, return it modulo 10^9 + 7.
Example 1:
	Input: locations = [2,3,6,8,4], start = 1, finish = 3, fuel = 5
	Output: 4
	Explanation: The following are all possible routes, each uses 5 units of fuel:
	1 -> 3
	1 -> 2 -> 3
	1 -> 4 -> 3
	1 -> 4 -> 2 -> 3
Example 2:
	Input: locations = [4,3,1], start = 1, finish = 0, fuel = 6
	Output: 5
	Explanation: The following are all possible routes:
	1 -> 0, used fuel = 1
	1 -> 2 -> 0, used fuel = 5
	1 -> 2 -> 1 -> 0, used fuel = 5
	1 -> 0 -> 1 -> 0, used fuel = 3
	1 -> 0 -> 1 -> 0 -> 1 -> 0, used fuel = 5
Example 3:
	Input: locations = [5,2,1], start = 0, finish = 2, fuel = 3
	Output: 0
	Explanation: It's impossible to get from 0 to 2 using only 3 units of fuel since the shortest route needs 4 units of fuel.
Example 4:
	Input: locations = [2,1,5], start = 0, finish = 0, fuel = 3
	Output: 2
	Explanation: There are two possible routes, 0 and 0 -> 1 -> 0.
Example 5:
	Input: locations = [1,2,3], start = 0, finish = 2, fuel = 40
	Output: 615088286
	Explanation: The total number of possible routes is 2615088300. Taking this number modulo 10^9 + 7 gives us 615088286.
Constraints:
	2 <= locations.length <= 100
	1 <= locations[i] <= 10^9
	All integers in locations are distinct.
	0 <= start, finish < locations.length
	1 <= fuel <= 200
 */
func CountRoutesDFS(locations []int, start int, finish int, fuel int) int {
	ans := 0
	memo := make([][]int, len(locations))
	for i := range memo{
		memo[i] = make([]int, fuel)
	}
	var dfs func(int, int, int)
	dfs = func(s, e, left int){
		if left < 0{
			return
		}
		if e == finish {
			ans++
		//	return
		}
		for i := range locations{
			if i != e{
				if locations[i] > locations[e]{
					dfs(e, i, left - locations[i] + locations[e])
				}else{
					dfs(e, i, left - locations[e] + locations[i])
				}
			}
		}
	}
	for i := range locations{
		if i != start{
			if locations[i] > locations[start]{
				dfs(start, i, fuel - locations[i] + locations[start])
			}else{
				dfs(start, i, fuel - locations[start] + locations[i])
			}
		}
	}
	// 特殊情况, 但不允许 start -> start -> start -> ... -> finish
	if start == finish{
		ans++
	}
	return ans
}

func CountRoutesDFSDP(locations []int, start int, finish int, fuel int) int {
	ans := 0
	// 特殊情况, 但不允许 start -> start -> start -> ... -> finish
	if start == finish{
		ans++
	}
	memo := make([][]int, len(locations))
	for i := range memo{
		memo[i] = make([]int, fuel+1)
		for j := range memo[i]{
			memo[i][j] = -1
		}
	}
	var dfs func(begin int, left int)int
	dfs = func(begin int, left int) int{
		result := 0
		if left < 0{
			return result
		}
		if memo[begin][left] != -1{
			return memo[begin][left]
		}
		// 剪枝：从某个位置出发直达目的地耗费的油量 如果大于 可用油量，则不会出现其他路径可达目的地
		need := locations[finish] - locations[begin]
		if need < 0{
			need = -need
		}
		if need > left{
			memo[begin][left] = 0
			return 0
		}
		if begin == finish{
			result = (result+1) % 1000000007
		}
		for i := range locations{
			if i != begin{
				if locations[i] > locations[begin]{
					result += dfs(i, left - locations[i] + locations[begin])
				}else{
					result += dfs(i, left - locations[begin] + locations[i])
				}
			}
		}
		memo[begin][left] = result % 1000000007
		return  memo[begin][left]
	}
	dfs(start, fuel)
	for i := range locations{
		if i != start{
			if locations[i] > locations[start]{
				ans += dfs(i, fuel - locations[i] + locations[start])
			}else{
				ans += dfs(i, fuel - locations[start] + locations[i])
			}
		}
	}
	return ans % 1000000007
	//return ans%(10e9+7)
}
/*
	dp[i][j]:表示从地点i出发，当前剩余fuel为j的情况下，到达目的地的路径数量
1. 从 DFS 方法签名出发。分析哪些入参是可变的，将其作为 DP 数组的维度；将返回值作为 DP 数组的存储值。 对应了 状态定义
2. 从 DFS 的主逻辑可以抽象中单个状态的计算方法。	对应了 状态方程转移
	DFS主逻辑：枚举所有的位置，看从当前位置i 出发，可以到达的位置有哪些。从而得出状态转移方程：
	dp[i][j] += dp[k][j-need]
	k 表示计算位置 i 油量的状态时枚举的下一个位置， need代表从 i 到达 k 需要的油量
	其中 i 与 k 无大小限制关系，只要求 i != k
	j 与 j - need 有严格限制，要求 j >= j-need
结果为 dp[start][fuel]; start 起始location 和 总油量
 */
func CountRoutesDP(locations []int, start int, finish int, fuel int) int {
	mod := 1000000007
	localLen := len(locations)
	dp := make([][]int, localLen)
	for i := range dp{
		dp[i] = make([]int, fuel+1)
	}
	// 对于本身位置就在目的地的状态，路径数为 1
	for i := 0; i <= fuel; i++{
		dp[finish][i] = 1
	}
	// dp[i][j] += dp[k][j-need], j 与 j-need 存在大小关系，因此需要先从小到大枚举油量
	for j := 0; j <= fuel; j++{
		for i := 0; i < localLen; i++{
			for k := 0; k < localLen; k++{
				if i != k{
					need := int(math.Abs(float64(locations[i] - locations[k])))
					if j >= need{
						dp[i][j] += dp[k][j - need]
						dp[i][j] %= mod
					}
				}
			}
		}
	}
	return dp[start][fuel]
}

/* 120. Triangle
Given a triangle array, return the minimum path sum from top to bottom.
For each step, you may move to an adjacent number of the row below. More formally,
if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.
 */
func MinimumTotal(triangle [][]int) int {
	// dp[i+1][j] = num[j] + min(dp[i][j-1], dp[i][j])
	row := len(triangle)
	if row <=0 {
		return 0
	}
	dp := [2][]int{}
	dp[0] = make([]int, len(triangle[0]))
	copy(dp[0], triangle[0])
	for i := 1; i < row; i++{
		pre, cur := (i-1)%2, i % 2
		tmp := []int{}
		for j := range triangle[i]{
			v := triangle[i][j]
			if j > 0 && j < len(dp[pre]){
				v += min(dp[pre][j-1], dp[pre][j])
			}else if j <= 0{
				v += dp[pre][j]
			}else {
				v += dp[pre][j-1]
			}
			tmp = append(tmp, v)
			// dp[cur] = append(dp[cur], v)  一直追加，导致历史计算数据在里面
		}
		dp[cur] = tmp
	}
	return min(dp[(row-1)%2]...)
}
// 从底往上计算的 dfs dp
func MinimumTotalDFS(triangle [][]int) int {
	// dp[i+1][j] = triangle[i][j] + min(dp[i][j-1], dp[i][j])
	// dp[i][j] = triangle[i][j] + min(dp[i+1][j], dp[i+1][j+1])
	n := len(triangle)
	dp := make([][]int, n)
	for i := range dp{
		dp[i] = make([]int, n)
		for j := range dp[i]{
			dp[i][j] = math.MaxInt32
		}
	}
	var dfs func(int, int)int
	dfs = func(row, col int)int{
		if row >= n{
			return 0
		}
		if dp[row][col] != math.MaxInt32{
			return dp[row][col]
		}
		dp[row][col] = min(dfs(row+1, col), dfs(row+1, col+1)) + triangle[row][col]
		return dp[row][col]
	}
	return dfs(0,0)
}

/* 931. Minimum Falling Path Sum
Given an n x n array of integers matrix, return the minimum sum of any falling path through matrix.
A falling path starts at any element in the first row and chooses the element in the next row that is either directly below or diagonally left/right.
Specifically, the next element from position (row, col) will be (row + 1, col - 1), (row + 1, col), or (row + 1, col + 1).
方程：dp[i][j] = matrix[i][j] + min(dp[i+1][j-1], dp[i+1][j], dp[i+1][j+1])
结果：min(dp[0]...)
*/
func MinFallingPathSumDFS(matrix [][]int) int {
	n := len(matrix)
	dp := [2][]int{} // 滚动数组
	for i := range dp{
		dp[i] = make([]int, len(matrix[0]))
		for j := range dp[i]{
			dp[i][j] = math.MaxInt32
		}
	}
	var dfs func(row int)
	dfs = func(row int){
		cur,next := row % 2, (row+1)%2
		if row >= n {
			for i := range dp[cur]{
				dp[cur][i] = 0
			}
			return
		}
		if dp[cur][0] == math.MaxInt32 {
			dfs(row+1)
			v := dp[next%2]
			for col := range matrix[row] {
				//fmt.Println(col, n, matrix)
				if col <= 0 {
					if len(v) == 1 {
						dp[cur][col] = matrix[row][col] + v[col]
					}else {
						dp[cur][col] = matrix[row][col] + min(v[col], v[col+1])
					}
				}else if col >= n-1 {
					dp[cur][col] = matrix[row][col] + min(v[col-1], v[col])
				}else {
					dp[cur][col] = matrix[row][col] + min(v[col-1], v[col], v[col+1])
				}
			}
		}
	}
	dfs(0)
	ans := math.MaxInt32
	for i := range dp[0]{
		if ans > dp[0][i]{
			ans = dp[0][i]
		}
	}
	return ans
}
func MinFallingPathSum(matrix [][]int) int {
//	方程：dp[i][j] = matrix[i][j] + min(dp[i+1][j-1], dp[i+1][j], dp[i+1][j+1])
//	结果：min(dp[0]...)
	ans := math.MaxInt32
	c := len(matrix[0])
	dp := make([]int, len(matrix[0]))
	for i := range matrix{
		tmp := make([]int, len(dp))
		for j := range matrix[i]{
			if j <= 0{
				if len(dp) == 1 {
					tmp[j] = matrix[i][j] + dp[j]
				}else{
					tmp[j] = matrix[i][j] + min(dp[j], dp[j+1])
				}
			}else if j >= c-1{
				tmp[j] = matrix[i][j] + min(dp[j-1], dp[j])
			}else{
				tmp[j] = matrix[i][j] + min(dp[j-1], dp[j], dp[j+1])
			}
		}
		dp = tmp
	}
	for i := range dp{
		if ans > dp[i]{
			ans = dp[i]
		}
	}
	return ans
}

/* 1289. Minimum Falling Path Sum II
Given an n x n integer matrix grid, return the minimum sum of a falling path with non-zero shifts.
A falling path with non-zero shifts is a choice of exactly one element from each row of grid such that no two elements chosen in adjacent rows are in the same column.
Example 1:
	Input: arr = [[1,2,3],[4,5,6],[7,8,9]]
	Output: 13
	Explanation:
	The possible falling paths are:
	[1,5,9], [1,5,7], [1,6,7], [1,6,8],
	[2,4,8], [2,4,9], [2,6,7], [2,6,8],
	[3,4,8], [3,4,9], [3,5,7], [3,5,9]
	The falling path with the smallest sum is [1,5,7], so the answer is 13.
*/
func MinFallingPathSumII(grid [][]int) int {
	// dp[i][j] = grid[i][j] + min(dp[i-1][col] col != j)
	n := len(grid)
	dp := make([]int, n)
	for i := range grid{
		tmp := make([]int, n)
		for j := range grid[i]{
			tmp[j] = grid[i][j]
			// t := min(append(dp[:j], dp[j+1:]...)...)  slice 底层数组 浅copy
			t := min2(j, dp...)
			if t != math.MaxInt32{
				tmp[j] += t
			}
		}
		dp = tmp
	}
	return min(dp...)
}
/* 576. Out of Boundary Paths
  There is an m x n grid with a ball. The ball is initially at the position [startRow, startColumn].
  You are allowed to move the ball to one of the four adjacent cells in the grid (possibly out of the grid crossing the grid boundary).
  You can apply at most maxMove moves to the ball.
  Given the five integers m, n, maxMove, startRow, startColumn, return the number of paths to move the ball out of the grid boundary.
  Since the answer can be very large, return it modulo 109 + 7.
Example 1:
	Input: m = 2, n = 2, maxMove = 2, startRow = 0, startColumn = 0
	Output: 6
Example 2:
	Input: m = 1, n = 3, maxMove = 3, startRow = 0, startColumn = 1
	Output: 12
 */
func FindPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	ans := 0
	//mod := 100000007
	// 方向数组： [i][j-1]  [i-1][j]  [i][j+1]  [i+1][j]
	dir := [2][]int{[]int{0, -1, 0, 1}, []int{-1, 0, 1, 0}}
	var dfs func(row, col, step int)
	dfs = func(row, col, step int){
		if step > maxMove{
			return
		}
		if row >=m || row < 0 || col >= n || col < 0{
			if step <= maxMove{ // 对应 at most， 如果是equal 则为 step == maxMove
				//ans = (ans+1)%mod
				ans++
			}
			return
		}
		for i := 0; i < 4; i++{
			dfs(row+dir[0][i], col+dir[1][i], step+1)
		}
	}
	dfs(startRow, startColumn, 0)
	return ans
}

func FindPathsDFSDP(m int, n int, maxMove int, startRow int, startColumn int) int {
	// dp[i][j][k] = dp[i-1][j][k-1] + dp[i][j-1][k-1] + dp[i][j+1][k-1] + dp[i+1][j][k-1]
	mod := 1000000007
	dp := make([][][]int, m+1)
	for i := range dp{
		dp[i] = make([][]int, n+1)
		for j := range dp[i]{
			dp[i][j] = make([]int, maxMove+1)
			for k := range dp[i][j]{
				dp[i][j][k] = -1
			}
		}
	}
	dp[startRow][startColumn][0] = 0
	// 方向数组： [i][j-1]  [i-1][j]  [i][j+1]  [i+1][j]
	dir := [2][]int{[]int{0, -1, 0, 1}, []int{-1, 0, 1, 0}}
	var dfs func(row, col, left int) int
	dfs = func(row, col, left int) int{
		if left < 0{
			return 0
		}
		if row >=m || row < 0 || col >= n || col < 0{
			if left >= 0{ // 对应 at most， 如果是equal 则为 step == maxMove
				return 1
			}
			return 0
		}
		//fmt.Println(row, col, left, dp[row][col][left])
		if dp[row][col][left] != -1{
			return dp[row][col][left]
		}
		cnt := 0
		for i := 0; i < 4; i++{
			cnt += dfs(row+dir[0][i], col+dir[1][i], left-1)
		}
		//fmt.Println(cnt)
		dp[row][col][left] = cnt % mod
		return cnt
	}
	ans := dfs(startRow, startColumn, maxMove) % mod
	//fmt.Println(dp)
	return ans
}
/* 需要修改维度表示
	dp[k][i][j]表示球移动 k 次之后位于坐标(i, j)的路径数量
	当 k = 0时，球一定位于(startRow,startColumn),因此边界条件：
	<1> dp[0][startRow][startColumn] = 1
	<2> dp[0][i][j] = 0, (i,j) != (startRow,startColumn)
	如果球移动i+1次之后，球一定位于(i,j)相邻坐标，记为(ii, jj) 此时分2类情况
	<1> (ii,jj)在正常范围内
		dp[k+1][ii][jj] += dp[k][i][j]
	<2> 出界
		统计路径值

 */
func FindPathsDP(m int, n int, maxMove int, startRow int, startColumn int) int {
	const mod int = 1e9+7
	ans := 0
	var dirs = []struct{x, y int}{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	dp := make([][][]int, maxMove+1)
	for k := range dp{
		dp[k] = make([][]int, m)
		for i := range dp[k]{
			dp[k][i] = make([]int, n)
		}
	}
	dp[0][startRow][startColumn] = 1
	for k := 0; k < maxMove; k++{
		for i := 0; i < m; i++{
			for j := 0; j < n; j++{
				if dp[k][i][j] > 0{
					for _, dir := range dirs{
						ii, jj := i+dir.x, j+dir.y
						if ii >= 0 && ii < m && jj >= 0 && jj < n{ // 正常范围内
							dp[k+1][ii][jj] = (dp[k][i][j] + dp[k+1][ii][jj]) % mod
						}else{ // 出界 开始统计
							ans = (ans + dp[k][i][j]) % mod
						}
					}
				}
			}
		}
	}
	return ans
}
// 优化: 注意到 dp[k][][] 只在计算 dp[k+1][][] 时会用到，因此可以将 dp 中的移动次数的维度省略
func FindPathsDPBest(m int, n int, maxMove int, startRow int, startColumn int) int {
	const mod int = 1e9+7
	ans := 0
	var dirs = []struct{x, y int}{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	/*
	dp := make([][][]int, maxMove+1)
	for k := range dp{
		dp[k] = make([][]int, m)
		for i := range dp[k]{
			dp[k][i] = make([]int, n)
		}
	}
	dp[0][startRow][startColumn] = 1
	 */
	dp := make([][]int, m)
	for i := range dp{
		dp[i] = make([]int, n)
	}
	dp[startRow][startColumn] = 1
	for k := 0; k < maxMove; k++{
		// 存储cur计算结果，后面 dp = tmp
		tmp := make([][]int, m)
		for o := range tmp{
			tmp[o] = make([]int, n)
		}
		for i := 0; i < m; i++{
			for j := 0; j < n; j++{
				/*
				if dp[k][i][j] > 0{
					for _, dir := range dirs{
						ii, jj := i+dir.x, j+dir.y
						if ii >= 0 && ii < m && jj >= 0 && jj < n{ // 正常范围内
							dp[k+1][ii][jj] = (dp[k][i][j] + dp[k+1][ii][jj]) % mod
						}else{ // 出界 开始统计
							ans = (ans + dp[k][i][j]) % mod
						}
					}
				} */
				if dp[i][j] > 0{
					for _, dir := range dirs{
						ii, jj := i+dir.x, j+dir.y
						if ii >= 0 && ii < m && jj >= 0 && jj < n { // 正常范围内
							tmp[ii][jj] = (tmp[ii][jj]+dp[i][j]) % mod
						}else{// 出界 开始统计
							ans = (ans + dp[i][j]) % mod
						}
					}
				}
			}
		}
		dp = tmp
	}
	return ans
}

/* 1301. Number of Paths with Max Score
	You are given a square board of characters. You can move on the board starting at the bottom right square marked with the character 'S'.
You need to reach the top left square marked with the character 'E'.
The rest of the squares are labeled either with a numeric character 1, 2, ..., 9 or with an obstacle 'X'.
In one move you can go up, left or up-left (diagonally) only if there is no obstacle there.
Return a list of two integers: the first integer is the maximum sum of numeric characters you can collect,
and the second is the number of such paths that you can take to get that maximum sum, taken modulo 10^9 + 7.
In case there is no path, return [0, 0].
 */
/* 三维处理 未完成  思路有误区，为什么首先想到的是三维，而不是携带2个状态参数
 */
/*
func PathsWithMaxScore(board []string) []int {
	maxScore := 0
	for i := range board{
		for j := range board[i]{
			if board[i][j] != 'E' && board[i][j] != 'S' && board[i][j] != 'X'{
				maxScore += int(board[i][j] - '0')
			}
		}
	}
	r, c := len(board), len(board[0])
	dp := make([][][]int, r)
	for i := range dp{
		dp[i] = make([][]int, c)
		for j := range dp[i]{
			dp[i][j] = make([]int, maxScore)
		}
	}
	dp[r-1][c-1][0] = 1
	dir := [][]int{[]int{1,0}, []int{1,1}, []int{0,1}}
	for i := r-1; i >= 0; i--{
		for j := c-1; j >= 0; j--{
			t := 0
			if board[i][j] == 'S'{
				t = 0
			}else if board[i][j] == 'X'{
				t = -1
			}else {
				t = int(board[i][j] - '0')
			}
			for k := 0; k <= maxScore; k++{
				for _, o := range dir{
					if k >= t && dp[i+o[0]][j+o[1]][k-t] != -1{
						dp[i][j][k] += dp[i+o[0]][j+o[1]][k-t]
					}
				}
			}
		}
	}
	ms,mc := 0,0
	return []int{}
}
 */
func PathsWithMaxScore(board []string) []int {
	const mod int = 1e9+7
	dir := [][]int{[]int{1,0}, []int{1,1}, []int{0,1}}
	n := len(board)
	dp := make([][][]int, n+1)
	for i := range dp{
		dp[i] = make([][]int, n+1)
		for j := range dp[i]{
			dp[i][j] = make([]int, 2)
		}
	}
	dp[n-1][n-1] = []int{0, 1}
	update := func(x,y int){
		u := board[x][y]
		if u == 'S'{
			return
		}
		if u == 'X'{
			dp[x][y][0] = -1
			dp[x][y][1] = 0
			return
		}
		w := 0
		if u != 'E'{
			w = int(board[x][y] - '0')
		}
		maxW, maxCnt := 0, 0
		for i := range dir{
			r, c := dir[i][0],dir[i][1]
			if x+r < n && y+c < n{
				if maxW == dp[x+r][y+c][0] {
					maxCnt = (maxCnt + dp[x+r][y+c][1]) % mod
				}else if maxW < dp[x+r][y+c][0]{
					maxCnt = dp[x+r][y+c][1]
					maxW = dp[x+r][y+c][0]
				}
			}
		}
		// 遗漏点-1 如果maxcnt== 0 即没有可达路径，那么最大权值也没有意义
		if maxCnt == 0{
			dp[x][y][0], dp[x][y][1] = maxW, maxCnt
		}else{
			dp[x][y][0], dp[x][y][1] = maxW+w, maxCnt
		}
	}
	for i := n-1; i >= 0; i--{
		for j := n-1; j >= 0; j--{
			update(i,j)
		}
	}
	return dp[0][0]
}
// 改自 leetcode 花费时间最小的解答
func PathsWithMaxScore2(board []string) []int {
	const mod int = 1e9+7
	dir := [][]int{[]int{1,0}, []int{1,1}, []int{0,1}}
	n := len(board)
	dp := make([][][2]int, n+1)
	for i := range dp{
		dp[i] = make([][2]int, n+1)
	}
	dp[n-1][n-1] = [2]int{0, 1}
	update := func(x,y int){
		u := board[x][y]
		if u == 'S'{
			return
		}
		if u == 'X'{
			dp[x][y][0] = -1
			dp[x][y][1] = 0
			return
		}
		w := 0
		if u != 'E'{
			w = int(board[x][y] - '0')
		}
		maxW, maxCnt := 0, 0
		for i := range dir{
			r, c := dir[i][0],dir[i][1]
			if x+r < n && y+c < n{
				if maxW == dp[x+r][y+c][0] {
					maxCnt = (maxCnt + dp[x+r][y+c][1]) % mod
				}else if maxW < dp[x+r][y+c][0]{
					maxCnt = dp[x+r][y+c][1]
					maxW = dp[x+r][y+c][0]
				}
			}
		}
		// 遗漏点-1 如果maxcnt== 0 即没有可达路径，那么最大权值也没有意义
		if maxCnt == 0{
			dp[x][y][0], dp[x][y][1] = maxW, maxCnt
		}else{
			dp[x][y][0], dp[x][y][1] = maxW+w, maxCnt
		}
	}
	for i := n-1; i >= 0; i--{
		for j := n-1; j >= 0; j--{
			update(i,j)
		}
	}
	return dp[0][0][:]
}
