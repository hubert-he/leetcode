package DFS_BFS

import (
	"math"
)

/* 286. Walls and Gates
** You are given an m x n grid rooms initialized with these three possible values.
	-1 A wall or an obstacle.
	0 A gate.
	INF Infinity means an empty room.
We use the value 2^31 - 1 = 2147483647 to represent INF as you may assume that the distance to a gate is less than 2147483647.
Fill each empty room with the distance to its nearest gate.
If it is impossible to reach a gate, it should be filled with INF.
 */
// 这个解法 ❌ 每次思考此类题目 DFS 总是, 直接四周递归存在问题，就是有些情况 无法获得最小值
func wallsAndGates(rooms [][]int)  {
	m, n := len(rooms), len(rooms[0])
	vis := make([][]bool, m)
	for i := range vis{
		vis[i] = make([]bool, n)
	}
	dirs := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0}}
	valid := func(x, y int)bool{
		if x < 0 || y < 0 || x >= m || y >= n || rooms[x][y] == -1{
			return false
		}
		return true
	}
	var dfs func(x, y int)int
	dfs = func(x, y int)int{
		if rooms[x][y] == 0{
			return 0
		}

		vis[x][y] = true
		for _, d := range dirs{
			xx, yy := x + d[0], y + d[1]
			if valid(xx,yy) {
				// 漏了这个，导致缺值，不能因为访问过，就忽略距离比较
				if vis[xx][yy]{ // 问题就出在这里，无法保证现在使用的xx yy 的值是最小值，因为还未完全递归完毕。
					rooms[x][y] = min(rooms[x][y], rooms[xx][yy]+1)
				}else{
					rooms[x][y] = min(rooms[x][y], dfs(xx, yy)+1)
				}
			}
		}
		return rooms[x][y]
	}
	for i := range rooms{
		for j := range rooms[i]{
			dfs(i, j)
		}
	}
}

/* 从门的方向出发 */
func WallsAndGates_DFS(rooms [][]int)  {
	m, n := len(rooms), len(rooms[0])
	dirs := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0}}
	/* 找单元格中的所有的门，然后从每个门的上下左右4个方向出发，去搜索所有单元格，把能覆盖到的单元格赋值。
	   ** 遇到3种情况停止搜索：
	       1.墙，-1；
	       2.新的门，0；
	       3.其他门能通到这里且距离比当前距离小；
	       这3种情况归纳为一个判断条件rooms[x][y] <= distance
	*/
	valid := func(x, y int)bool{
		if x < 0 || y < 0 || x >= m || y >= n || rooms[x][y] == -1 || rooms[x][y] == 0{
			return false
		}
		return true
	}
	var dfs func(x, y, disc int)
	dfs = func(x, y, disc int){
		if !valid(x, y){
			return
		}

	}
	for i := range rooms{
		for j := range rooms[i]{
			if rooms[i][j] == 0{
				for _, d := range dirs {
					dfs(i+d[0], j+d[1], 0)
				}
			}
		}
	}
}
/* 另一个思路： 遍历所有节点， 检查它 周围节点通过该节点到达门的距离会不会更短， 若会，则更新节点值，并递归判断修改后的节点值。*/
func WallsAndGates_DFS2(rooms [][]int)  {
	m, n := len(rooms), len(rooms[0])
	dirs := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0}}
	valid := func(x, y int)bool{
		if x < 0 || y < 0 || x >= m || y >= n || rooms[x][y] == -1 {
			return false
		}
		return true
	}
	var dfs func(x, y int)
	dfs = func(x, y int){
		if !valid(x, y) || rooms[x][y] == math.MaxInt32{
			return
		}
		for _, d := range dirs{
			xx, yy := x + d[0], y + d[1]
			// 检查它 周围节点通过该节点到达门的距离会不会更短
			if valid(xx, yy) && rooms[xx][yy] > rooms[x][y] + 1{
				rooms[xx][yy] = rooms[x][y] + 1
				dfs(xx, yy)
			}
		}
	}
	for i := range rooms{
		for j := range rooms[i]{
			dfs(i, j)
		}
	}
}
/* 另外一个知识点： 多源BFS
** 几个起点同时开始BFS，其实不是严格意义的同时，如果把他们看成树的话，同时的含义是指在同一层，然后由这一层同时向下遍历
** 这里有一个隐藏的思考点：
	多源BFS，那么距离某一个源最近的点的值肯定会先被修改为距离，
	后面的源再遍历到这里的时候就已经有距离值了，而且再遍历到这个点时候的距离肯定大于等于该点已有的距离值。
	所以不需要进行比较值大小
** 图中求最短距离，BFS是经典套路
* 相似题目：
** 994. Rotting Oranges
** 1765. Map of Highest Peak
*/
func WallsAndGates_BFS(rooms [][]int)  {
	m, n := len(rooms), len(rooms[0])
	dirs := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0}}
	valid := func(x, y int)bool{
		if x < 0 || y < 0 || x >= m || y >= n || rooms[x][y] == -1 {
			return false
		}
		return true
	}
	q := [][3]int{} // 下标 以及 距离
	for i := range rooms{
		for j := range rooms[i]{
			if rooms[i][j] == 0{
				q = append(q, [3]int{i, j, 0})
			}
		}
	}
	for len(q) > 0{
		t := [][3]int{}
		for idx := range q{
			x, y, disc := q[idx][0], q[idx][1], q[idx][2]
			for _, d := range dirs{
				xx, yy := x + d[0], y + d[1]
				if valid(xx, yy) && rooms[xx][yy] == math.MaxInt32{// 空房间
					rooms[xx][yy] = disc + 1
					t = append(t, [3]int{xx, yy, disc+1})
				}
			}
		}
		q = t
	}
}
/* 994. Rotting Oranges
** You are given an m x n grid where each cell can have one of three values:
	0 representing an empty cell,
	1 representing a fresh orange, or
	2 representing a rotten orange.
** Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.
** Return the minimum number of minutes that must elapse until no cell has a fresh orange.
** If this is impossible, return -1.
 */
/* 最短路径问题还是推荐用广度优先遍历BFS
** 这里给出DFS 作为练习， 注意图思想的转换
** 网络结构中的格子有 上、下、左、右 4 个相邻节点，
** 对于格子 (row, col) 来说，其相邻的 4 个节点为 (row + 1, col)、(row - 1, col)、(row, col + 1)、(row, col - 1)相当于是一棵四叉树。
	对于网格中边缘的格子，相当于是二叉树中的叶子节点
	对于超出网格范围的格子，相当于是二叉树的 root == null
*/
func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dirs := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0}}
	valid := func(x, y int)bool{
		if x < 0 || y < 0 || x >= m || y >= n {
			return false
		}
		return true
	}
	var dfs func(x, y, time int)
	dfs = func(x, y, time int){
		if !valid(x, y) || (grid[x][y] != 1 && grid[x][y] < time ){
			return
		}
		grid[x][y] = time
		for _, d := range dirs{
			dfs(x+d[0], y+d[1], time+1)
		}
	}
	for i := range grid{
		for j := range grid[i]{
			if grid[i][j] == 2{
				dfs(i, j, 2) // 借用grid中的值来记录时间（2 表示时间 0 ），最后计算结果减去2
			}
		}
	}
	ans := 0
	for i := range grid{
		for j := range grid[i] {
			if grid[i][j] == 1{
				return -1
			} else {
				ans = max(ans, grid[i][j]-2)
			}
		}
	}
	return ans
}

/* 1765. Map of Highest Peak
** 注意 问题描述的措辞
** You are given an integer matrix isWater of size m x n that represents a map of land and water cells.
	If isWater[i][j] == 0, cell (i, j) is a land cell.
	If isWater[i][j] == 1, cell (i, j) is a water cell.
** You must assign each cell a height in a way that follows these rules:
	The height of each cell must be non-negative.
	If the cell is a water cell, its height must be 0.
	Any two adjacent cells must have an absolute height difference of 「at most」 1. // 注意 距离至多 1
	A cell is adjacent to another cell if the former is directly north, east, south, or west of the latter (i.e., their sides are touching).
	Find an assignment of heights such that the maximum height in the matrix is maximized.
	Return an integer matrix height of size m x n where height[i][j] is cell (i, j)'s height.
	If there are multiple solutions, return any of them.
*/

/* 417. Pacific Atlantic Water Flow
** There is an m x n rectangular island that borders both the Pacific Ocean and Atlantic Ocean.
** The Pacific Ocean touches the island's left and top edges, and the Atlantic Ocean touches the island's right and bottom edges.
** The island is partitioned into a grid of square cells.
** You are given an m x n integer matrix heights where heights[r][c] represents the height above sea level of the cell at coordinate (r, c).
** The island receives a lot of rain, and the rain water can flow to neighboring cells directly north, south, east, and west
** if the neighboring cell's height is less than or equal to the current cell's height.
** Water can flow from any cell adjacent to an ocean into the ocean.
** Return a 2D list of grid coordinates result where result[i] = [ri, ci] denotes
** that rain water can flow from cell (ri, ci) to both the Pacific and Atlantic oceans.
 */
// 2021-12-16 写到dfs 写不下去了
// 是因为思维出现误区， 如果一个路径 是永远走不到的， 这种情况，也是属于 访问到的节点，思考下 水流
func pacificAtlantic(heights [][]int) [][]int {
	r, c := len(heights), len(heights[0])
	m := map[[2]int]int{}
	dirs := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0} }
	toP := func(x, y int)bool{
		if x < 0 || y < 0{
			return true
		}
		return false
	}
	toA := func(x,y int)bool{
		if x >= r || y >= c{
			return true
		}
		return false
	}
	var dfs func(x int, y int, end func(int,int)bool, invalid func(int,int)bool)[][][2]int
	dfs = func(x int, y int, end func(int,int)bool, invalid func(int,int)bool)[][][2]int{
		ans := [][][2]int{}
		for _, d := range dirs{
			xx, yy := x + d[0], y + d[1]
			if invalid(xx, yy){
				continue
			}
			if end(xx, yy){

			}
			if heights[xx][yy] < heights[x][y]{
				ans = append(ans, dfs(xx, yy, end, invalid)...)
			}
		}
		return ans
	}
	// from P to A
	res := [][][2]int{}
	for i := 0; i < r; i++{
		res = dfs(i, 0, toA, toP)
		for i := range res{
			for j := range res[i] {
				m[res[i][j]] = 1
			}
		}
	}
	for i := 0; i < c; i++{
		res = dfs(0, i, toA, toP)
		for i := range res{
			for j := range res[i] {
				m[res[i][j]] = 1
			}
		}
	}
	for i := 0; i < r; i++{
		res = dfs(i, c-1, toP, toA)
		for j := range res{
			m[res[i][j]]++
		}
	}
	for i := 0; i < c; i++{
		res = dfs(r-1, i, toP, toA)
		for j := range res{
			m[res[i][j]] += 1
		}
	}
	ans := [][]int{}
	for k := range m{
		if m[k] > 1{
			ans = append(ans, []int{k[0], k[1]})
		}
	}
	return ans
}

// 本质还是求两个 Visit 数组，保存公共的元素
func pacificAtlanticDFS(heights [][]int) [][]int {
	r, c := len(heights), len(heights[0])
	P, A := make([][]bool, r), make([][]bool, r)
	for i := 0; i < r; i++{
		P[i], A[i] = make([]bool, c), make([]bool, c)
	}
	dirs := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0} }
	ans := [][]int{}
	vaild := func(x, y int)bool{
		if x < 0 || y < 0 || x >= r || y >= c{
			return false
		}
		return true
	}
	var dfs func(visited [][]bool, x, y int)
	dfs = func(visited [][]bool, x, y int){
		if visited[x][y]{
			return
		}
		visited[x][y] = true
		if P[x][y] && A[x][y]{
			ans = append(ans, []int{x, y})
		}
		for _, d := range dirs{
			xx, yy := x + d[0], y + d[1]
			if vaild(xx, yy) && heights[x][y] >= heights[xx][yy]{
				dfs(visited, xx, yy)
			}
		}
	}
	for i := 0; i < r; i++{
		dfs(P, i, 0)
		dfs(A, i, c-1)
	}
	for i := 0; i < c; i++{
		dfs(P, 0, i)
		dfs(A, r-1, i)
	}
	return ans
}