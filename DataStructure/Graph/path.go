package Graph

import (
	"container/heap"
	"math"
)

/* 此部分 关注 最短路径问题
* 1. 单源最短路径
	1. Dijkstra
	2. Bellman-Ford
* 2. 多源最短路径
	1. Floyd
*/
/* 最短路径题目汇总：
** leetcode-743
** leetcode-1368
*/
/* 743. Network Delay Time
** You are given a network of n nodes, labeled from 1 to n.
** You are also given times, a list of travel times as directed edges times[i] = (ui, vi, wi),
** where ui is the source node, vi is the target node, and wi is the time it takes for a signal to travel from source to target.
** We will send a signal from a given node k.
** Return the time it takes for all the n nodes to receive the signal.
** If it is impossible for all the n nodes to receive the signal, return -1.
 */
/* 题目主要是 应用单源最短路径算法
** 这里采用 Dijkstra， 其主要思想是贪心
** 将所有节点分成两类：已确定从起点到当前点的最短路长度的节点，以及未确定从起点到当前点的最短路长度的节点（下面简称「未确定节点」和「已确定节点」）
** 每次从「未确定节点」中取一个与起点距离最短的点，将它归类为「已确定节点」，并用它「更新」从起点到其他所有「未确定节点」的距离。
** 直到所有点都被归类为「已确定节点」
** 用节点 A「更新」节点 B 的意思是，用起点到节点 A 的最短路长度加上从节点 A 到节点 B 的边的长度，
** 去比较起点到节点 B 的最短路长度，如果前者小于后者，就用前者更新后者。这种操作也被叫做「松弛」。
** 每次选择「未确定节点」时，起点到它的最短路径的长度可以被确定
** 因为我们已经用了每一个「已确定节点」更新过了当前节点，无需再次更新（因为一个点不能多次到达）。
** 而当前节点已经是所有「未确定节点」中与起点距离最短的点，不可能被其它「未确定节点」更新。所以当前节点可以被归类为「已确定节点」。
*/
/* 从节点 k 发出的信号，到达节点 x 的时间就是节点 k 到节点 x 的最短路的长度。因此我们需要求出节点 k 到其余所有点的最短路，其中的最大值就是答案。
** 若存在从 k 出发无法到达的点，则返回 −1。
 */
func networkDelayTime(times [][]int, n int, k int) int {
	// 构建图
	graph := make([][]int, n)
	for i := range graph{
		graph[i] = make([]int, n)
		for j := range graph[i]{
			graph[i][j] = math.MaxInt32
		}
	}
	for _, t := range times{
		x, y := t[0]-1, t[1]-1 // 将节点编号减小了 1，从而使节点编号位于 [0,n-1] 范围
		graph[x][y] = t[2]
	}
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	// 源点初始化
	dist[k-1] = 0
	visited := make([]bool, n)
	for i := 0; i < n; i++{
		x := -1
		// 每次找到「最短距离最小」且「未被更新」的点 t
		for y, v := range visited{
			if !v && (x == -1 || dist[y] < dist[x]){
				x = y
			}
		}
		visited[x] = true
		// 用节点 A「更新」节点 B 的意思是，用起点到节点 A 的最短路长度加上从节点 A 到节点 B 的边的长度
		for y, time := range graph[x]{
			dist[y] = min(dist[y], dist[x] + time)
		}
	}
	ans := max(dist...)
	if ans == math.MaxInt32{
		return -1
	}
	return ans
}
// 使用一个小根堆来寻找「未确定节点」中与起点距离最近的点。
func networkDelayTime_Heap(times [][]int, n int, k int) int {
	type edge struct { to, time int}
	graph := make([][]edge, n)
	for _, t := range times{
		x, y := t[0] - 1, t[1] -1
		graph[x] = append(graph[x], edge{y, t[2]})
	}
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[k-1] = 0
	h := &hp{{0, k-1}}
	for h.Len() > 0{
		p := heap.Pop(h).(pair)
		x := p.x
		if dist[x] < p.d{
			continue
		}
		for _, e := range graph[x]{
			y := e.to
			if d := dist[x] + e.time; d < dist[y]{
				dist[y] = d
				heap.Push(h, pair{d, y})
			}
		}
	}
	ans := max(dist...)
	if ans == math.MaxInt32{
		return -1
	}
	return ans
}
/* 使用邻接矩阵 来是实现 Dijkstra
*/
func networkDelayTime_matrix_dij(times [][]int, n int, k int) int {
	graph := make([][]int, n+1) // 注意题目是从 1 开始的
	for i := range graph {
		graph[i] = make([]int, n+1)
		for j := range graph[i] {
			if i != j {
				graph[i][j] = math.MaxInt32
			}
		}
	}
	dist := make([]int, n+1)
	// 构建图
	for _, t := range times{
		u, v, w := t[0], t[1], t[2]
		graph[u][v] = w
	}
	dijkstra := func(){
		// 起始先将所有的点标记为「未更新」和「距离为正无穷」
		visited := make([]bool, n+1)
		for i := range dist{
			dist[i] = math.MaxInt32
		}
		// 只有起点最短距离为 0
		dist[k] = 0
		// 迭代 n 次
		for p := 1; p <= n; p++{
			// 每次找到「最短距离最小」且「未被更新」的点 t
			t := -1
			for i := 1; i <= n; i++{
				if !visited[i] && (t == -1 || dist[i] < dist[t]){
					t = i
				}
			}
			// 标记点 t 为已更新
			visited[t] = true
			// 用点 t 的「最小距离」更新其他点
			for i := 1; i <= n; i++{
				dist[i] = min(dist[i], dist[t] + graph[t][i])
			}
		}
	}
	dijkstra()
	ans := max(dist[1:]...) // 注意题目是从 1 开始的
	if ans == math.MaxInt32{
		return -1
	}
	return ans
}
/*Bellman Ford: 在「负权图中求最短路」的 Bellman Ford 进行求解，该算法也是「单源最短路」算法，复杂度为 O(n * m)
** 通常为了确保 O(n * m)，可以单独建一个类代表边，将所有边存入集合中，在 n 次松弛操作中直接对边集合进行遍历
** 另外对于 Bellman ford 的优化版本，暂时记下，以后学习
** SPFA（邻接表）SPFA 是对 Bellman Ford 的优化实现，可以使用队列进行优化，也可以使用栈进行优化。
*/
func networkDelayTime_bellmanFord(times [][]int, n int, k int) int {
	type edge struct {
		u, v int // 边的端点
		w int	// 权值
	}
	dist := make([]int, n+1)
	es := []edge{}
	//m := len(times) // 边的个数
	// 构建图
	for _, t := range times{
		u, v, w := t[0], t[1], t[2]
		es = append(es, edge{u, v, w})
	}
	bellmanFord := func(){
		// 起始先将所有的点标记为「距离为正无穷」
		for i := range dist{
			dist[i] = math.MaxInt32
		}
		dist[k] = 0
		tmpDist := make([]int, n+1)
		for p := 1; p <= n; p++{
			copy(tmpDist, dist) // 每次都使用上一次迭代的结果，执行松弛操作
			for _, e := range es{
				dist[e.v] = min(dist[e.v], tmpDist[e.u] + e.w)
			}
		}
	}
	bellmanFord()
	ans := max(dist[1:]...) // 注意题目是从 1 开始的
	if ans == math.MaxInt32{
		return -1
	}
	return ans
}
/* Floyd
** 根据「基本分析」，我们可以使用复杂度为 立方级 的「多源汇最短路」算法 Floyd 算法进行求解，同时使用「邻接矩阵」来进行存图。
** 跑一遍 Floyd，可以得到「从任意起点出发，到达任意起点的最短距离」。
** 然后从所有 w[k][x] 中取 max 即是「从 k 点出发，到其他点 x 的最短距离的最大值」。
*/
func networkDelayTime_floyd(times [][]int, n int, k int) int {
	// 使用邻接矩阵
	//graph := make([][]int, n) 注意题目是从 1 开始的
	graph := make([][]int, n+1)
	for i := range graph{
		graph[i] = make([]int, n + 1)
		for j := range graph[i]{
			if i != j{
				graph[i][j] = math.MaxInt32
			}
		}
	}
	// 构建图
	for _, t := range times{
		u, v, w := t[0], t[1], t[2]
		graph[u][v] = w
	}
	floyd := func(){ // floyd 基本流程为三层循环：枚举中转点 - 枚举起点 - 枚举终点 - 松弛操作
		for p := 1; p <= n; p++{
			for i := 1; i <= n; i++{
				for j := 1; j <= n; j++{
					graph[i][j] = min(graph[i][j], graph[i][p] + graph[p][j])
				}
			}
		}
	}
	floyd()
	//ans := max(graph[k]...) 注意题目是从 1 开始的
	ans := max(graph[k][1:]...)
	if ans == math.MaxInt32{
		return -1
	}
	return ans
}

/* 1368. Minimum Cost to Make at Least One Valid Path in a Grid
** Given an m x n grid. Each cell of the grid has a sign pointing to the next cell you should visit if you are currently in this cell.
** The sign of grid[i][j] can be:
	1 which means go to the cell to the right. (i.e go from grid[i][j] to grid[i][j + 1])
	2 which means go to the cell to the left. (i.e go from grid[i][j] to grid[i][j - 1])
	3 which means go to the lower cell. (i.e go from grid[i][j] to grid[i + 1][j])
	4 which means go to the upper cell. (i.e go from grid[i][j] to grid[i - 1][j])
** Notice that there could be some signs on the cells of the grid that point outside the grid.
** You will initially start at the upper left cell (0, 0).
** A valid path in the grid is a path that starts from the upper left cell (0, 0) and ends at the bottom-right cell (m - 1, n - 1) following the signs on the grid.
** The valid path does not have to be the shortest.
** You can modify the sign on a cell with cost = 1. You can modify the sign on a cell one time only.
** Return the minimum cost to make the grid have at least one valid path.
*/
/* 如果没有「每个格子中的数字只能修改一次」这个条件，我们可以很轻松地将本题建模成一个求最短路径的问题：
	1. 我们希望求出从左上角 (0, 0) 到右下角 (m - 1, n - 1) 的「最小代价」；
	2. 当我们在任意位置 (i, j) 时，我们都可以向上、下、左、右移动一个位置，但不能走出边界。
	3. 如果移动的位置与 (i, j) 处的箭头方向一致，那么移动的代价为 0，否则为 1； 这个思路
		这样以来，我们可以将数组 grid 建模成一个包含 mn 个节点和不超过 4mn 条边的有向图 G。
	4. 图 G 中的每一个节点表示数组 grid 中的一个位置，它会向不超过 4 个相邻的节点各连出一条边，
		边的权值要么为 0（移动方向与箭头方向一致），要么为 1（移动方向与箭头方向不一致）；
我们在图 GG 上使用任意一种最短路算法，求出从 (0,0) 到 (m−1,n−1) 的最短路，即可得到答案
 */
func minCost(grid [][]int) int {
	dirs := [][2]int{ [2]int{0, 1}, [2]int{0, -1}, [2]int{1, 0}, [2]int{-1, 0}}
	m, n := len(grid), len(grid[0])
	// 注意此题图的建模方式
	dist, visited := make([]int, m*n), make([]bool, m*n)
	// Dijkstra 算法适合用来求出无负权边图中的单源最短路径
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	valid := func(x, y int)bool{
		if x < 0 || y < 0 || x >= m || y >= n{
			return false
		}
		return true
	}
	dist[0] = 0
	// 迭代 m*n 次
	for i := range dist{
		/*
		if visited[i] {
			continue
		}
		 */
		// 找全网最小的
		t := -1
		for p := 0; p < m*n; p++{
			if !visited[p] && (t == -1 || dist[t] > dist[p]){
				t = p
			}
		}
		visited[t] = true
		//x, y := i/n, i%n
		x, y := t / n, t % n
		for j, d := range dirs{
			xx, yy := x + d[0], y + d[1]
			if valid(xx, yy){
				newPos := xx * n + yy
				newDisc := dist[i]
				if grid[x][y] != j+1{
					newDisc += 1 // 切换方向代价为1
				}
				if newDisc < dist[newPos]{
					dist[newPos] = newDisc
				}
			}
		}
	}
	return dist[m*n-1]
}
// Dijkstra
func minCost_Heap(grid [][]int) int {
	dirs := [][]int{ []int{}, []int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0} }
	m, n := len(grid), len(grid[0])
	total := m*n
	dist := make([]int, total)
	for i := range dist{
		dist[i] = math.MaxInt32
	}
	valid := func(x, y int)bool{
		if x < 0 || y < 0 || x >= m || y >= n{
			return false
		}
		return true
	}
	dijkstra := func(){
		dist[0] = 0 // 初始化起点
		h := &hp{{0, 0}}
		for h.Len() > 0{ // 默认是最小堆
			p := heap.Pop(h).(pair) // p.x 即是全网未访问的到起点距离最小的节点
			// 下面利用 p.x 更新其他点
			x, y := p.x/n, p.x%n
			for j, d := range dirs{
				if j == 0{ continue }
				xx, yy := x + d[0], y + d[1]
				if valid(xx, yy){
					newPos, newDist := xx*n+yy, dist[p.x]
					if grid[x][y] != j{
						newDist += 1
					}
					if dist[newPos] > newDist{
						dist[newPos] = newDist
						heap.Push(h, pair{newDist, newPos})
					}
				}
			}
		}
	}
	dijkstra()
	return dist[total-1]
}

/* 1129. Shortest Path with Alternating Colors
** You are given an integer n, the number of nodes in a directed graph where the nodes are labeled from 0 to n - 1.
** Each edge is red or blue in this graph, and there could be self-edges and parallel edges.
** You are given two arrays redEdges and blueEdges where:
	redEdges[i] = [ai, bi] indicates that there is a directed red edge from node ai to node bi in the graph, and
	blueEdges[j] = [uj, vj] indicates that there is a directed blue edge from node uj to node vj in the graph.
** Return an array answer of length n,
** where each answer[x] is the length of the shortest path from node 0 to node x such that
** the edge colors alternate along the path, or -1 if such a path does not exist.
 */
func shortestAlternatingPaths_BFS(n int, redEdges [][]int, blueEdges [][]int) []int {
	const Red, Blue = 0, 1
	g := make([][][]int, n)
	for i := range g{
		g[i] = make([][]int, 2)
	}
	for i := range redEdges{
		src, dst := redEdges[i][0], redEdges[i][1]
		g[src][Red] = append(g[src][Red], dst)
	}
	for i := range blueEdges{
		src, dst := blueEdges[i][0], blueEdges[i][1]
		g[src][Blue] = append(g[src][Blue], dst)
	}
	distance := make([][]int, n)
	for i := range distance{
		distance[i] = make([]int, 2)
		distance[i][0], distance[i][1] = math.MaxInt32, math.MaxInt32
	}
	distance[0][0], distance[0][1] = 0, 0
	qr, qb := []int{0}, []int{0}
	step := 0
	for len(qr) > 0 || len(qb) > 0{
		step++
		tr, tb := qr, qb
		qr, qb = nil, nil
		for i := range tr{
			for _, node := range g[tr[i]][Red]{
				if distance[node][Red] > step{
					distance[node][Red] = step
					qb = append(qb, node)
				}
			}
		}
		for i := range tb{
			for _, node := range g[tb[i]][Blue]{
				if distance[node][Blue] > step{
					distance[node][Blue] = step
					qr = append(qr, node)
				}
			}
		}
	}
	ans := make([]int, n)
	for i, d := range distance{
		ans[i] = min(d...)
		if ans[i] == math.MaxInt32{
			ans[i] = -1
		}
	}
	return ans
}
/* 在使用dfs 解决的时候，dfs 退出条件犯迷糊
** dfs 退出的2种思维：
** 1. visited 矩阵的使用，但是如果环存在合理的情况下，即节点可多次访问，因此无法使用visited，如此题情况
** 2. 根据距离的计算，如果距离大于已知的，证明兜圈子了，可退出
此题 想到用第 2 种情况，但是局限在一维，导致思维 dead road
此题有 2种边， 也即有2种距离，如果计算两种距离来判断是否已经兜圈子了，可完美解决
 */
func shortestAlternatingPaths_DFS_error(n int, redEdges [][]int, blueEdges [][]int) []int {
	const Red, Blue = 0, 1
	g := make([][][]int, n)
	for i := range g{
		g[i] = make([][]int, 2)
	}
	// red:0, blue: 1
	for i := range redEdges{
		src, dst := redEdges[i][0], redEdges[i][1]
		g[src][Red] = append(g[src][0], dst)
	}
	for i := range blueEdges{
		src, dst := blueEdges[i][0], blueEdges[i][1]
		g[src][Blue] = append(g[src][1], dst)
	}
	ans := make([]int, n)
	for i := range ans{
		ans[i] = -1
	}
	var dfs func(node, dist, color int)
	dfs = func(node, dist, color int){
		if ans[node] == -1 || ans[node] > dist{
			ans[node] = dist
		}
		// visited[node] = true
		for _, neigh := range g[node][color]{
			// 不能直接设置vis，因为环的存在是合理的，case-2 就会有节点无法到达的通路
			//if vis[neigh] { continue } // 针对case-1 图存在环
			if color == Red{
				dfs(neigh, dist+1, Blue)
			}else{
				dfs(neigh, dist+1, Red)
			}
			//visited[neigh] = false
		}
	}
	dfs(0, 0, Red)
	dfs(0, 0, Blue)
	return ans
}

func shortestAlternatingPaths_DFS(n int, redEdges [][]int, blueEdges [][]int) []int {
	const Red, Blue = 0, 1
	g := make([][2][]int, n)
	for _, path := range redEdges{
		g[path[0]][Red] = append(g[path[0]][Red], path[1])
	}
	for _, path := range blueEdges{
		g[path[0]][Blue] = append(g[path[0]][Blue], path[1])
	}
	distance := make([][2]int, n)
	for i := range distance{
		distance[i][0], distance[i][1] = math.MaxInt32, math.MaxInt32
	}
	var dfs func(node, dist, color int)
	dfs = func(node, dist, color int){
		if distance[node][color] < dist{
			return
		}
		if distance[node][color] > dist{
			distance[node][color] = dist
		}
		for _, neigh := range g[node][color]{
			// 不能直接设置vis，因为环的存在是合理的，case-2 就会有节点无法到达的通路
			//if vis[neigh] { continue } // 针对case-1 图存在环
			/*
			if color == Red{
				dfs(neigh, dist+1, Blue)
			}else{
				dfs(neigh, dist+1, Red)
			}*/
			// 异或应用
			dfs(neigh, dist+1, color ^ 1)
		}
	}
	dfs(0, 0, Red)
	dfs(0, 0, Blue)
	ans := make([]int, n)
	for i := range ans{
		ans[i] = min(distance[i][Red], distance[i][Blue])
		if ans[i] == math.MaxInt32{
			ans[i] = -1
		}
	}
	return ans
}


















