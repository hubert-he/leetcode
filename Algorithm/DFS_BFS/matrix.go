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

// 2022-03-02 使用BFS 出现思维漏洞
// 不能 仅 从两个方向走， 由于会有回流的情况，考虑下面这个case
/*
	1	2	3
	8	9	4
	7	6	5
==> toA 后
	1	2	3*
	8*	9*	4*
	7*	6*	5*
==> toP 后
	1*	2*	3*
	8*	9*	4*
	7*	6	5*
==> 漏掉了  [2,1] 6 这个元素
 */
func pacificAtlanticBFS_error(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	toA, toP := [][]int{[]int{1, 0}, []int{0,1}}, [][]int{[]int{-1,0}, []int{0,-1}}
	valid := func(x, y int)bool{
		if x < 0 || y < 0 || x >= m || y >= n{
			return false
		}
		return true
	}
	ans := [][]int{}
	q := [][]int{}
	cantoA := map[[2]int]bool{}
	for i := range heights{
		for j := range heights[i]{
			if i == n-1 || j == n-1{
				q = append(q, []int{i, j})
				cantoA[[2]int{i,j}] = true
			}
		}
	}
	for len(q) > 0{
		t := q
		q = nil
		for i := range t{
			sx, sy := t[i][0], t[i][1]
			for _, d := range toP{
				x, y := sx+d[0], sy+d[1]
				//if valid(x, y) && heights[sx][sy] <= heights[x][y] && !cantoA[[2]int{x,y}] {
				if valid(x, y) && heights[sx][sy] <= heights[x][y]{
					q = append(q, []int{x, y})
					cantoA[[2]int{x, y}] = true
				}
			}
		}
	}
	cantoP := map[[2]int]bool{}
	for i := range heights{
		for j := range heights[i]{
			if i == 0 || j == 0{
				q = append(q, []int{i, j})
				cantoP[[2]int{i,j}] = true
			}
		}
	}
	for len(q) > 0{
		t := q
		q = nil
		for i := range t{
			sx, sy := t[i][0], t[i][1]
			for _, d := range toA{
				x, y := sx+d[0], sy+d[1]
				//if valid(x, y) && heights[sx][sy] <= heights[x][y] && !cantoP[[2]int{x,y}]{
				if valid(x, y) && heights[sx][sy] <= heights[x][y]{
					q = append(q, []int{x, y})
					cantoP[[2]int{x, y}] = true
				}
			}
		}
	}
	for k := range cantoA{
		if cantoP[k]{
			ans = append(ans, []int{k[0], k[1]})
		}
	}
	return ans
}

func pacificAtlanticBFS(heights [][]int) [][]int{
	m, n := len(heights), len(heights[0])
	//toA, toP := [][]int{[]int{1, 0}, []int{0,1}}, [][]int{[]int{-1,0}, []int{0,-1}}
	dirs := [][]int{[]int{1, 0}, []int{0,1}, []int{-1,0}, []int{0,-1}}
	valid := func(x, y int)bool{
		if x < 0 || y < 0 || x >= m || y >= n{
			return false
		}
		return true
	}
	qa, qp := [][]int{}, [][]int{}
	cantoA,cantoP := map[[2]int]bool{}, map[[2]int]bool{}
	for i := range heights{
		for j := range heights[i]{
			if i == m-1 || j == n-1{
				qa = append(qa, []int{i, j})
				cantoA[[2]int{i,j}] = true
			}
			if i == 0 || j == 0{
				qp = append(qp, []int{i, j})
				cantoP[[2]int{i,j}] = true
			}
		}
	}
	for len(qa) > 0 || len(qp) > 0{
		t := qa
		qa = nil
		for i := range t{
			src := t[i]
			for _, d := range dirs{
				dst := []int{src[0]+d[0], src[1]+d[1]}
				if valid(dst[0],dst[1]) &&
					heights[src[0]][src[1]] <= heights[dst[0]][dst[1]] &&
					!cantoA[[2]int{dst[0], dst[1]}]{
					qa = append(qa, dst)
					cantoA[[2]int{dst[0], dst[1]}] = true
				}
			}
		}
		t = qp
		qp = nil
		for i := range t{
			src := t[i]
			for _, d := range dirs{
				dst := []int{src[0]+d[0], src[1]+d[1]}
				if valid(dst[0], dst[1]) &&
					heights[src[0]][src[1]] <= heights[dst[0]][dst[1]] &&
					!cantoP[[2]int{dst[0], dst[1]}]{
					qp = append(qp, dst)
					cantoP[[2]int{dst[0], dst[1]}] = true
				}
			}
		}
	}
	ans := [][]int{}
	for k := range cantoA{
		if cantoP[k]{
			ans = append(ans, []int{k[0], k[1]})
		}
	}
	return ans
}
/* 1020. Number of Enclaves(飞地)
** You are given an m x n binary matrix grid, where 0 represents a sea cell and 1 represents a land cell.
** A move consists of walking from one land cell to another adjacent (4-directionally) land cell or
** walking off the boundary of the grid.
** Return the number of land cells in grid for which we cannot walk off the boundary of the grid in any number of moves
 */
// 2022-02-22 DFS 刷出此题， 现在学习 多源BFS 来加强 对题目分析的思路
// 多源BFS收录
func numEnclaves(grid [][]int) int {
	dirs := [][]int{[]int{0,1}, []int{0,-1}, []int{1, 0}, []int{-1, 0}}
	ans := 0
	m, n := len(grid), len(grid[0])
	valid := func(x, y int)bool{
		if x >= m || y >= n || x < 0 || y < 0{
			return false
		}
		return true
	}
	vis := make([][]bool, m)
	for i := range vis{
		vis[i] = make([]bool, n)
	}
	queue := [][]int{}
	// 将所有「边缘陆地」看做与超级源点相连，起始将所有「边缘陆地」进行入队
	// 等价于只将超级源点入队，然后取出超级源点并进行拓展
	for i := range grid{
		for j := range grid[i]{
			if i == 0 || j == 0 || i == m-1 || j == n-1{
				if grid[i][j] == 1{
					vis[i][j] = true
					queue = append(queue, []int{i,j})
				}
			}
		}
	}
	for len(queue) > 0 {
		t := [][]int{}
		for i := range queue{
			for _, d := range dirs {
				x, y := queue[i][0]+d[0], queue[i][1]+d[1]
				if valid(x, y) && grid[x][y] == 1 && !vis[x][y]{
					vis[x][y] = true
					t = append(t, []int{x, y})
				}
			}
		}
		queue = t
	}
	// 将边缘的处理掉，然后统计剩余的为1 单元格即为答案
	for i := 0; i < m; i++{
		for j := 0; j < n; j++{
			if grid[i][j] == 1 && !vis[i][j]{
				ans++
			}
		}
	}
	return ans
}

/* 1091. Shortest Path in Binary Matrix
** Given an n x n binary matrix grid, return the length of the shortest clear path in the matrix.
** If there is no clear path, return -1.
** A clear path in a binary matrix is a path from the top-left cell (i.e., (0, 0)) to
** the bottom-right cell (i.e., (n - 1, n - 1)) such that:
	All the visited cells of the path are 0.
	All the adjacent cells of the path are 8-directionally connected
	(i.e.,they are different and they share an edge or a corner).
** The length of a clear path is the number of visited cells of this path.
 */
// 2022-02-23 刷出此题 双向BFS
func shortestPathBinaryMatrix_BFS(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1{
		return -1
	}
	if n == 1{	return 1  }
	valid := func(pos []int)bool{
		x, y := pos[0], pos[1]
		if x >= n || y >= n || x < 0 || y < 0{
			return false
		}
		return true
	}
	// case中考虑[3,0]和[3,1] 此时同在front队列发现彼此在已访问中，所有要分开
	visf := make([][]int, n)
	for i := range visf{
		visf[i] = make([]int, n)
	}
	visb := make([][]int, n)
	for i := range visb{
		visb[i] = make([]int, n)
	}
	dirs := [][]int{[]int{1, 0}, []int{0, 1}, []int{1,1}, []int{-1, 0}, []int{0, -1}, []int{-1,-1},[]int{1,-1}, []int{-1,1}}
	//fdirs := [][]int{[]int{1, 0}, []int{0, 1}, []int{1,1}, []int{1,-1}}
	//bdirs := [][]int{[]int{-1, 0}, []int{0, -1}, []int{-1,-1}, []int{-1,1}}
	front := [][]int{[]int{0,0}}
	bottom := [][]int{[]int{n-1,n-1}}
	visf[0][0], visb[n-1][n-1] = 1, 1
	for len(front) > 0 && len(bottom) > 0{
		tf, tb := [][]int{}, [][]int{}
		for i := range front{
			src := []int{front[i][0], front[i][1]}
			for _, d := range dirs{
				dst := []int{src[0]+d[0], src[1]+d[1]}
				if valid(dst) && grid[dst[0]][dst[1]] == 0 && visf[dst[0]][dst[1]] == 0{
					if visb[dst[0]][dst[1]] != 0{
						return visb[dst[0]][dst[1]] + visf[front[i][0]][front[i][1]]
					}
					visf[dst[0]][dst[1]] = visf[front[i][0]][front[i][1]] + 1
					tf = append(tf, dst)
				}
			}
		}
		for i := range bottom{
			src := []int{bottom[i][0], bottom[i][1]}
			for _, d := range dirs{
				dst := []int{src[0]+d[0], src[1]+d[1]}
				if valid(dst) && grid[dst[0]][dst[1]] == 0 && visb[dst[0]][dst[1]] == 0{
					if visf[dst[0]][dst[1]] != 0{
						return visf[dst[0]][dst[1]] + visb[front[i][0]][front[i][1]]
					}
					visb[dst[0]][dst[1]] = visb[bottom[i][0]][bottom[i][1]] + 1
					tb = append(tb, dst)
				}
			}
		}
		front, bottom = tf, tb
	}
	return -1
}

// 2022-03-02 双向BFS重刷此题
// 与上一个区别在于，基于gird 修改，无需2个vis 矩阵，另外 增加slen, elen 计数来汇总结果
func shortestPathBinaryMatrixBFS(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if grid[0][0] == 1 || grid[m-1][n-1] == 1{
		return -1
	}
	if n == 1{
		return 1
	}
	grid[0][0], grid[m-1][n-1] = 3, 2
	s, e := [][]int{[]int{0,0}}, [][]int{[]int{m-1, n-1}}
	dirs := [][]int{
		[]int{-1, 0}, []int{-1, 1}, []int{0,1}, []int{1,1},
		[]int{1, 0}, []int{1, -1}, []int{0,-1}, []int{-1,-1}}
	valid := func(x,y int)bool{
		if x < 0 || y <0 || x >= m || y >= n{
			return false
		}
		return true
	}
	slen, elen := 1, 1
	for len(s) > 0 && len(e) > 0{
		qs, qe := s, e
		s, e = nil, nil
		for i := range qs{
			for _, d := range dirs{
				x, y := qs[i][0]+d[0], qs[i][1]+d[1]
				if valid(x,y){
					if grid[x][y] == 0{
						s = append(s, []int{x, y})
						grid[x][y] = 3
					}
					if grid[x][y] == 2{
						return slen+elen
					}
				}
			}
		}
		slen++
		for i := range qe{
			for _, d := range dirs{
				x, y := qe[i][0] + d[0], qe[i][1] + d[1]
				if valid(x, y){
					if grid[x][y] == 0{
						e = append(e, []int{x,y})
						grid[x][y] = 2
					}
					if grid[x][y] == 3{
						return slen+elen
					}
				}
			}
		}
		elen++
	}
	return -1
}

// 此DP 解法是错误的 ❌, 收录进来 个人需要思考 为何会出现思维漏洞
/* 考虑 case：
[
	[0,1,1,0,0,0],
	[0,1,0,1,1,0],
	[0,1,1,0,1,0],
	[0,0,0,1,1,0],
	[1,1,1,1,1,0],
	[1,1,1,1,1,0]]
 */
func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1{
		return -1
	}
	valid := func(x, y int)bool{
		if x >= n || y >= n || x < 0 || y < 0{
			return false
		}
		return true
	}
	min := func(nums ...int)int{
		m := nums[0]
		for _, c := range nums{
			if m > c { m = c}
		}
		return m
	}
	dp := make([][]int, n)
	for i := range dp{
		dp[i] = make([]int, n)
		for j := range dp[i]{
			dp[i][j] = math.MaxInt32
		}
	}
	dp[n-1][n-1] = 1
	dirs := [][]int{[]int{1, 0}, []int{0, 1}, []int{1,1}}
	for i := n-1; i >= 0; i--{
		for j := n-1; j >= 0; j--{
			if i == n-1 && j == n-1 { continue }
			for _, d := range dirs{
				x, y := i+d[0], j+d[1]
				if valid(x, y) && grid[x][y] == 0{
					dp[i][j] = min(dp[i][j], dp[x][y]+1)
				}
			}
		}
	}
	return dp[0][0]
}

func shortestPathBinaryMatrix_DP(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1{
		return -1
	}
	valid := func(x, y int)bool{
		if x >= n || y >= n || x < 0 || y < 0{
			return false
		}
		return true
	}
	dp := make([][]int, n)
	for i := range dp{
		dp[i] = make([]int, n)
		for j := range dp[i]{
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 1
	// DP 的迭代路线 需要BFS 作为支撑
	q := [][]int{[]int{0,0}}
	// 需要从 8 个方向开进，这也是之前思考缺少的地方，认为从上往下走，但是参考测试例子-5，会发现路径会从下往上走的可能
	dirs := [][]int{[]int{1, 0},[]int{0, 1},[]int{1,1},[]int{-1, 0},[]int{0, -1},[]int{-1,-1},[]int{1,-1},[]int{-1,1}}
	for len(q) > 0{
		t := [][]int{}
		for i := range q{
			src := []int{q[i][0], q[i][1]}
			for _, d := range dirs {
				dst := []int{src[0]+d[0], src[1]+d[1]}
				if valid(dst[0],dst[1]) && grid[dst[0]][dst[1]] == 0{
					// 利用dp 作为visited 矩阵
					if dp[dst[0]][dst[1]] == math.MaxInt32{
						t = append(t, dst)
						dp[dst[0]][dst[1]] = dp[src[0]][src[1]] + 1
					}else{// 曾经访问过
						dp[dst[0]][dst[1]] = min(dp[dst[0]][dst[1]], dp[src[0]][src[1]] + 1)
					}
				}
			}
		}
		q = t
	}
	if dp[n-1][n-1] == math.MaxInt32{ 	return -1  }
	return dp[n-1][n-1]
}

/* 542. 01 Matrix
** Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.
** The distance between two adjacent cells is 1.
 */
// 2022-02-24 刷出此题， 多源BFS
// 复用matrix 为 visited 矩阵的写法
func updateMatrix_BFS(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	dirs := [][]int{[]int{0,1}, []int{0,-1}, []int{1, 0}, []int{-1, 0}}
	valid := func(x, y int)bool{
		if x < 0 || y < 0 || x >= m || y >= n{
			return false
		}
		return true
	}
	ans := [][]int{}
	// 将目标节点首先添加进队列
	queue := [][]int{}
	for i := range matrix{
		for j := range matrix[i]{
			if matrix[i][j] == 0{
				queue = append(queue, []int{i, j})
			}
		}
	}
	for len(queue) > 0{
		q := queue
		queue = nil
		for _, src := range q{
			for _, d := range dirs{
				x, y := src[0]+d[0], src[1]+d[1]
				if valid(x, y) && matrix[x][y] == 1{
					ans[x][y] = ans[src[0]][src[1]] + 1
					matrix[x][y] = 0 // 复用matrix 为 visited
					queue = append(queue, []int{x,y})
				}
			}
		}
	}
	return ans
}

// 方法二：多源BFS  更贴近 DP 的写法
// 结果直接在matrix 进行迭代
func updateMatrix_BFS_DP(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	dirs := [][]int{[]int{0,1}, []int{0,-1}, []int{1, 0}, []int{-1, 0}}
	valid := func(x, y int)bool{
		if x < 0 || y < 0 || x >= m || y >= n{
			return false
		}
		return true
	}
	visited := make([][]bool, m)
	for i := range visited{
		visited[i] = make([]bool, n)
	}
	// 将目标节点首先添加进队列
	queue := [][]int{}
	for i := range matrix{
		for j := range matrix[i]{
			if matrix[i][j] == 0{
				queue = append(queue, []int{i, j})
				visited[i][j] = true
			}
		}
	}
	for len(queue) > 0{
		q := queue
		queue = nil
		for _, src := range q{
			for _, d := range dirs{
				x, y := src[0]+d[0], src[1]+d[1]
				if valid(x, y) && !visited[x][y] {
					matrix[x][y] = matrix[src[0]][src[1]] + 1 // 迭代结果
					visited[x][y] = true
					queue = append(queue, []int{x,y})
				}
			}
		}
	}
	return matrix
}

/* 方法三： DP 推导： 可降低空间复杂度
** 在方法二中，可以进一步优化，
** 如果 0 在矩阵中的位置是(i0, j0), 1 在矩阵中的位置是(i1, j1)那么我们可以直接算出 此 0 和 此 1 之间的距离。
** 但是如何计算出最短，这个时候就是DP 处理的全部情况下，选出最短的距离（也即选出离此 1 最近的那个 0）
** 令 f(i, j) 表示位置 (i,j) 到最近的 0 的距离
** 从一个固定的 1 走到任意一个 0，在距离最短的前提下可能有四种方法：
	1. 只有 水平向左移动 和 竖直向上移动  ==> f(i, j) = 1 + min( f(i-1,j), f(i, j-1) )   (i,j)为1
	2. 只有 水平向左移动 和 竖直向下移动	==>	f(i, j) = 1 + min( f(i-1,j), f(i, j+1) )   (i,j)为1
	3. 只有 水平向右移动 和 竖直向上移动	==>	f(i, j) = 1 + min( f(i+1,j), f(i, j-1) )   (i,j)为1
	4. 只有 水平向右移动 和 竖直向下移动	==>	f(i, j) = 1 + min( f(i+1,j), f(i, j+1) )   (i,j)为1
** 因此，最后的f(i,j) 为这4种情况比较剩余的结果
 */
// 这个解法是错误的 ❌
func updateMatrix_DP_error(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	dirs := [][]int{[]int{0,1}, []int{0,-1}, []int{1, 0}, []int{-1, 0}}
	dp := make([][]int, m)
	for i := range dp{
		dp[i] = make([]int, n)
		for j := range dp[i]{
			if mat[i][j] == 0{
				dp[i][j] = 0
			}else{
				dp[i][j] = math.MaxInt32
			}
		}
	}
	valid := func(x, y int)bool{
		if x < 0 || y < 0 || x >= m || y >= n{
			return false
		}
		return true
	}
	for _, d := range dirs{
		for i := range mat{
			for j := range mat[i]{
				x, y := i + d[0], j + d[1]
				//if mat[i][j] == 1 && valid(x, y){
				if !valid(x, y){ continue}
				if mat[i][j] == 1{
					dp[i][j] = min( dp[i][j], 1 + dp[x][y] )
				}else{
					dp[i][j] = 0 // 这里有个遗漏，即点的4个方向都存在 才能赋值 ，否则是默认的 MaxInt32
				}
			}
		}
	}
	return dp
}
// 这个解法是错误的 ❌
func updateMatrix_DP_error2(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	//dirs := [][]int{[]int{0,1}, []int{0,-1}, []int{1, 0}, []int{-1, 0}}
	dirs := [][][]int{
		[][]int{[]int{-1, 0}, []int{0,-1}}, // 水平向左移动 和 竖直向上移动
		[][]int{[]int{-1, 0}, []int{0,1}},	//水平向左移动 和 竖直向下移动
		[][]int{[]int{1, 0}, []int{0,-1}},	// 水平向右移动 和 竖直向上移动
		[][]int{[]int{1, 0}, []int{0,1}},	// 水平向右移动 和 竖直向下移动
	}
	dp := make([][]int, m)
	for i := range dp{
		dp[i] = make([]int, n)
		for j := range dp[i]{
			if mat[i][j] == 0{
				dp[i][j] = 0
			}else{
				dp[i][j] = math.MaxInt32
			}
		}
	}
	valid := func(x, y int)bool{
		if x < 0 || y < 0 || x >= m || y >= n{
			return false
		}
		return true
	}
	for _, dir := range dirs{
		for i := range mat{
			for j := range mat[i]{
				for _, d := range dir{
					x, y := i + d[0], j + d[1]
					if !valid(x, y){ continue }
					dp[i][j] = min(dp[i][j], dp[x][y]+1)
				}
			}
		}
	}
	return dp
}
// 上述解法，忽略了 动态规划 迭代的路径方向， 导致有情况始终 无法更新
func updateMatrix_DP(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	dp := make([][]int, m)
	for i := range dp{
		dp[i] = make([]int, n)
		for j := range dp[i]{
			if mat[i][j] == 0{
				dp[i][j] = 0
			}else{
				dp[i][j] = math.MaxInt32
			}
		}
	}
	// 务必注意 递推的方向
	// 水平向左移动 和 竖直向上移动  ==> f(i, j) = 1 + min( f(i-1,j), f(i, j-1) )
	for i := 0; i < m; i++{
		for j := 0; j < n; j++{
			if i-1 >= 0 && dp[i][j] > dp[i-1][j] + 1{
				dp[i][j] = dp[i-1][j] + 1
			}
			if j-1 >= 0 && dp[i][j] > dp[i][j-1] + 1{
				dp[i][j] = dp[i][j-1] + 1
			}
		}
	}
	// 只有 水平向右移动 和 竖直向下移动，注意动态规划的计算顺序
	for i := m-1; i >= 0; i--{
		for j := n-1; j >= 0; j--{
			if i+1 < m && dp[i][j] > dp[i+1][j] + 1 {
				dp[i][j] = dp[i+1][j] + 1
			}
			if j+1 < n && dp[i][j] > dp[i][j+1] + 1 {
				dp[i][j] = dp[i][j+1] + 1
			}
		}
	}
	return dp
}

/* 1926. Nearest Exit from Entrance in Maze
** You are given an m x n matrix maze (0-indexed) with empty cells (represented as '.') and walls (represented as '+').
** You are also given the entrance of the maze,
** where entrance = [entrancerow, entrancecol] denotes the row and column of the cell you are initially standing at.
** In one step, you can move one cell up, down, left, or right.
** You cannot step into a cell with a wall, and you cannot step outside the maze.
** Your goal is to find the nearest exit from the entrance.
** An exit is defined as an empty cell that is at the border of the maze. The entrance does not count as an exit.
** Return the number of steps in the shortest path from the entrance to the nearest exit, or -1 if no such path exists.
 */
// 先序 DFS: 2020-03-03 通过DFS 刷出此题
// 通过添加 无出口情况提前处理, 超时的额外处理 才避免超时
// 先序DFS  无法 使用 记忆化 ，因为拆分的子问题计算 是  非 后向性
// 动态规划条件：
// 0. 最优子结构性质 —— 一个最优化策略的子策略总是最优的。一个问题满足最优化原理又称其具有最优子结构性质 此条目是DP 基础核心
// 1. 子问题的重叠性
// 2. 无后向性 ——
//		将各阶段按照一定的次序排列好之后，对于某个给定的阶段状态，它以前各阶段的状态无法直接影响它未来的决策，而只能通过当前的这个状态。
//		换句话说，每个状态都是过去历史的一个完整总结。这就是无后向性，又称为无后效性
func nearestExit_DFS(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	dirs := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0}}
	valid := func(x, y int)bool{
		if x < 0 || y < 0 || x >= m || y >= n {
			return false
		}
		if maze[x][y] == '+'{
			return false
		}
		return true
	}
	isExit := func(x, y int)bool{
		if x == 0 || y == 0 || x == m-1 || y == n-1{
			if maze[x][y] == '.' && !(x == entrance[0] && y == entrance[1]){
				return true
			}
		}
		return false
	}
	ans := math.MaxInt32
	// 无出口情况提前处理, 超时的额外处理
	dead := true
	for i := range maze{
		for j := range maze[i]{
			if isExit(i, j){
				dead = false
				break
			}
		}
	}
	if dead { return -1 }
	dist := make([][]int, m)
	for i := range dist{
		dist[i] = make([]int, n)
	}
	var dfs func(src [2]int, step int)
	dfs = func(src [2]int, step int) {
		if step >= ans { return }
		// 判断出口
		if isExit(src[0], src[1]){
			if ans > step{
				ans = step
			}
			return
		}
		nextStep := step+1
		for _, d := range dirs{
			dst := [2]int{src[0]+d[0], src[1]+d[1]}
			if valid(dst[0], dst[1]) && (dist[dst[0]][dst[1]] == 0 || dist[dst[0]][dst[1]] > nextStep ){
				dist[dst[0]][dst[1]] = nextStep
				dfs(dst, nextStep)
			}
		}
	}
	dfs([2]int{entrance[0], entrance[1]}, 0)
	if ans == math.MaxInt32{ return -1 }
	return ans
}

// DFS dp 错误版本-- DFS 采用是 后序DFS
// 无法应用DP， 无后效性 特性不满足
func nearestExit_DFS_Error(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	dirs := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0}}
	valid := func(x, y int)bool{
		if x < 0 || y < 0 || x >= m || y >= n {
			return false
		}
		if maze[x][y] == '+'{
			return false
		}
		return true
	}
	isExit := func(x, y int)bool{
		if x == 0 || y == 0 || x == m-1 || y == n-1{
			if maze[x][y] == '.' && !(x == entrance[0] && y == entrance[1]){
				return true
			}
		}
		return false
	}
	dead := true // 无出口情况提前处理, 超时的额外处理
	dp := make([][]int, m)
	for i := range dp{
		dp[i] = make([]int, n)
		for j := range dp[i]{
			if isExit(i, j){
				dp[i][j], dead = 0, false
			}else if i == 0 || j == 0 || i == m-1 || j == n-1{
				dp[i][j] = math.MaxInt32
			}else {
				dp[i][j] = -1
			}
		}
	}
	if dead {return -1 }
	var dfs func(x, y int)int
	dfs = func(x, y int)int{
		if dp[x][y] != -1 && (x != entrance[0] || y != entrance[1]){
			return dp[x][y]
		}
		maze[x][y] = '+'
		for _, d := range dirs{
			dx, dy := x+d[0], y+d[1]
			if valid(dx, dy){
				ret := dfs(dx, dy)
				maze[dx][dy] = '.'
				if dp[dx][dy] == -1 || dp[dx][dy] > ret{
					dp[dx][dy] = ret
				}
				if ret != math.MaxInt32 && (dp[x][y] == -1 || dp[x][y] > dp[dx][dy]+1){
					dp[x][y] = dp[dx][dy]+1
				}
			}
		}
		// 需要判断：死路 dp[x][y]=-1 需要更正
		if dp[x][y] == -1{ dp[x][y] = math.MaxInt32 }
		return dp[x][y]
	}
	ans := dfs(entrance[0], entrance[1])
	if ans == math.MaxInt32{ return -1 }
	return ans
}

/* 934. Shortest Bridge
** You are given an n x n binary matrix grid where 1 represents land and 0 represents water.
** An island is a 4-directionally connected group of 1's not connected to any other 1's.
** There are exactly two islands in grid.
** You may change 0's to 1's to connect the two islands to form one island.
** Return the smallest number of 0's you must flip to connect the two islands.
 */
func shortestBridge(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dirs := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0}}
	valid := func(x, y int)bool{
		if x < 0 || y < 0 || x >= m || y >= n {
			return false
		}
		return true
	}
	vis, vis1, vis2 := make([][]bool, n), make([][]bool, n), make([][]bool, n)
	for i := range vis1{
		vis1[i] = make([]bool, n)
		vis2[i] = make([]bool, n)
	}
	var dfs func(x, y int, visited [][]bool, q [][]int)
	dfs = func(x, y int, visited [][]bool, q [][]int){
		visited[x][y] = true
		vis[x][y] = true
		isBoarder := false
		for _, d := range dirs{
			px, py := x+d[0], y+d[1]
			if valid(px, py) && !vis[px][py] && !visited[px][py]{
				if grid[px][px] == 1 {
					dfs(px, py, visited, q)
				}else{// 岛屿的边界
					if !isBoarder{
						isBoarder = true
						q = append(q, []int{x, y})
					}
				}
			}
		}
	}
	q1, q2 := [][]int{}, [][]int{}
	for i := range grid{
		for j := range grid[i]{
			if !vis[i][j] && grid[i][j] == 1{
				dfs(i, j, vis1, q1)
			}
			if !vis[i][j] && grid[i][j] == 1{
				dfs(i, j, vis2, q2)
			}
		}
	}
	// 双向BFS 查找
	c1, c2 := 0, 0
	for len(q1) > 0 && len(q2) > 0 {
		t1, t2 := q1, q2
		q1,q2 = nil, nil
		for i := range t1{
			x, y := t1[i][0], t1[i][1]
			for _, d := range dirs{
				px, py := x+d[0], y+d[1]
				if valid(px, py) && grid[px][py] == 0 && !vis1[px][py]{
					if vis2[px][py]{
						return c1+c2
					}
					vis1[px][py] = true
					q1 = append(q1, []int{px, py})
				}
			}
		}
		c1++
		for i := range t2{
			x, y := t2[i][0], t2[i][1]
			for _, d := range dirs{
				px, py := x+d[0], y+d[1]
				if valid(px, py) && grid[px][py] == 0 && !vis2[px][py]{
					if vis1[px][py]{
						return c1+c2
					}
					vis2[px][py] = true
					q2 = append(q2, []int{px, py})
				}
			}
		}
		c2++
	}
	return -1
}






















