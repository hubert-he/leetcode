package Graph

/*本部分为 图 连通性问题
** 连通性问题分为：
	1. 强连通分量(strongly Connected components)
	2. 双连通分量
	3. 割点 与 桥
	4. 圆方树
*/

/*强连通分量(strongly Connected components) SCC
** 强连通的定义是：有向图 G 强连通是指，G 中任意两个结点连通。
** 强连通分量（Strongly Connected Components，SCC）的定义是：极大的强连通子图。
** 线性复杂度的 Tarjan 算法--- https://oi-wiki.org/graph/scc/
** 将图 DFS 成为一颗 树，有向图的 DFS 生成树主要可能有4 种边出现
** 	1. 树边（tree edge）：	每次搜索找到一个还没有访问过的结点的时候就形成了一条树边。
	2. 反祖边（back edge）：	也被叫做回边，即指向祖先结点的边。
	3. 横叉边（cross edge）：	它主要是在搜索的时候遇到了一个已经访问过的结点，但是这个结点 并不是 当前结点的祖先。
	4. 前向边（forward edge）：它是在搜索的时候遇到子树中的结点的时候形成的。
** 考虑 DFS 生成树与强连通分量之间的关系。
** 如果结点 u 是某个强连通分量在搜索树中遇到的第一个结点，那么这个强连通分量的其余结点肯定是在搜索树中以 u 为根的子树中。
** 结点 u 被称为这个强连通分量的根。
** 反证法：假设有个结点 v 在该强连通分量中但是不在以 u 为根的子树中，那么 u 到 v 的路径中肯定有一条离开子树的边。
** 但是这样的边只可能是横叉边或者反祖边，然而这两条边都要求指向的结点已经被访问过了，这就和 u 是第一个访问的结点矛盾了。得证。
*/
/* Tarjan 算法求强连通分量
** 在 Tarjan 算法中为每个结点 u 维护了以下几个变量：
	1. dfn[u]: 深度优先搜索遍历时结点 u 被搜索的次序, 在其他参考里也成为 时间戳， 规范了次序
	2. low[u]: 能够回溯到的最早的已经在栈中的结点。设以 u 为根的子树为SubTree-u 。 定义为以下结点的 dfn 的最小值：SubTree-u 中的结点；
				从SubTree-u通过一条不在搜索树上的边能到达的结点。
** 一个结点的子树内结点的 dfn 都大于该结点的 dfn。
** 从根开始的一条路径上的 dfn 严格递增，low 严格非降。
** 按照深度优先搜索算法搜索的次序对图中所有的结点进行搜索。在搜索过程中，对于结点 u 和与其相邻的结点 v（ v 不是 u 的父节点）考虑 3 种情况：
	未被访问：继续对 v 进行深度搜索。在回溯过程中，用 low[v] 更新 low[u] 。
			因为存在从 u 到 v 的直接路径，所以 v 能够回溯到的已经在栈中的结点，u 也一定能够回溯到。
	被访问过，已经在栈中：根据 low 值的定义，用 dfn[v] 更新 low[u]。
	被访问过，已不在栈中：说明 v 已搜索完毕，其所在连通分量已被处理，所以不用对其做操作。
** 对于一个连通分量图，我们很容易想到，在该连通图中有且仅有一个 u 使得 dfn[u] == low[u]。
** 该结点一定是在深度遍历的过程中，该连通分量中第一个被访问过的结点，因为它的 dfn 和 low 值最小，不会被该连通分量中的其他结点所影响。
** 因此，在回溯的过程中，判定 dfn[u] == low[u] 是否成立，如果成立，则栈中 u 及其上方的结点构成一个 SCC。
TARJAN_SEARCH(int u)
    vis[u]=true
    low[u]=dfn[u]=++dfncnt
    push u to the stack
    for each (u,v) then do
        if v hasn't been searched then
            TARJAN_SEARCH(v) // 搜索
            low[u]=min(low[u],low[v]) // 回溯
        else if v has been in the stack then
            low[u]=min(low[u],dfn[v])
 */



/* 割点 与 桥
** 割点： 对于一个无向图，如果把一个点删除后这个图的极大连通分量数增加了，那么这个点就是这个图的割点（又称割顶）。
** 如果我们尝试删除每个点，并且判断这个图的连通性，那么复杂度会特别的高。所以要介绍一个常用的算法：Tarjan。
** 在设置好 dfn 和 low 后，开始 DFS，
** 判断某个点是否是割点的根据是：对于某个顶点 u ，如果存在至少一个顶点 v（ u 的儿子），使得 low[v] >= dfn[u]，即不能回到祖先，那么 u 点为割点

** 割边/桥： 对于一个无向图，如果删掉一条边后图中的连通分量数增加了，则称这条边为桥或者割边。
		严谨来说，就是：假设有连通图 G = {V, E}， e 是其中一条边，如果 G - e 是不连通的，则边 e 是图 G 的一条割边（桥）。
** 和割点差不多，只要改一处：
	在一张无向图中，判断边 e 其对应的两个节点分别为 u 与 v 是否为桥，需要其满足如下条件即low[v] > dfn[u] 就可以了，而且不需要考虑根节点的问题。
** 割边是和根节点没关系的，原来我们求割点的时候是指点 v 是不可能不经过父节点 u 为回到祖先节点（包括父节点），所以顶点 u 是割点。
** 如果 low[v] == dfn[u] 表示还可以回到父节点，如果顶点 v 不能回到祖先也没有另外一条回到父亲的路，那么 u - v 这条边就是割边。
 */
/* 1192. Critical Connections in a Network
** There are n servers numbered from 0 to n - 1 connected by undirected server-to-server connections forming a network
** where connections[i] = [ai, bi] represents a connection between servers ai and bi.
** Any server can reach other servers directly or indirectly through the network.
** A critical connection is a connection that, if removed, will make some servers unable to reach some other server.
** Return all critical connections in the network in any order.
** 题目为求 桥
 */
func criticalConnections(n int, connections [][]int) [][]int {
	graph := make([][]int, n)
	dfn, low, visited := make([]int, n), make([]int, n), make([]bool, n)
	ts := 0
	for i := range connections{
		u, v := connections[i][0], connections[i][1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	ans := [][]int{}
	var tarjanSearch func(u, from int)
	tarjanSearch = func(u, from int){
		visited[u] = true
		dfn[u], low[u] = ts, ts
		ts++
		for _, v := range graph[u]{
			if v == from{
				continue
			}
			if visited[v]{
				//low[u] = min(low[u], low[v])
				low[u] = min(low[u], dfn[v])
				continue
			}
			// 没访问过的节点处理
			tarjanSearch(v, u)
			low[u] = min(low[u], low[v])
			if low[v] > dfn[u] {
				ans = append(ans, []int{u,v})
			}
		}
	}
	for u := range graph{
		if visited[u] {
			continue
		}
		tarjanSearch(u, -1)
	}
	return ans
}
// 可以进一步小优化，visited 与 dfn 可以合二为一
// dfn 为 0时候, 表示 未被访问
func criticalConnections2(n int, connections [][]int) [][]int {
	graph := make([][]int, n)
	dfn, low := make([]int, n), make([]int, n)
	ts := 1
	for i := range connections{
		u, v := connections[i][0], connections[i][1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	ans := [][]int{}
	var tarjanSearch func(u, from int)
	tarjanSearch = func(u, from int){
		dfn[u], low[u] = ts, ts
		ts++
		for _, v := range graph[u]{
			if v == from{
				continue
			}
			if dfn[v] != 0{
				//low[u] = min(low[u], low[v])
				low[u] = min(low[u], dfn[v])
				continue
			}
			// 没访问过的节点处理
			tarjanSearch(v, u)
			low[u] = min(low[u], low[v])
			if low[v] > dfn[u] {
				ans = append(ans, []int{u,v})
			}
		}
	}
	for u := range graph{
		if dfn[u] != 0 {
			continue
		}
		tarjanSearch(u, -1)
	}
	return ans
}

/* 1568. Minimum Number of Days to Disconnect Island
** You are given an m x n binary grid 「grid」 consisting where 1 represents land and 0 represents water.
** An island is a maximal 4-directionally (horizontal or vertical) connected group of 1's.
** The grid is said to be connected if we have exactly one island, otherwise is said disconnected.
** In one day, we are allowed to change any single land cell (1) into a water cell (0).
** Return the minimum number of days to disconnect the grid.
** Constraints:
	m == grid.length
	n == grid[i].length
	1 <= m, n <= 30
	grid[i][j] is either 0 or 1.
** 求 割点
 */
/* 仔细思考我们会发现最终的答案不可能超过 2。因为对于 n×m 的岛屿（n,m≥2），我们总是可以将某个角落相邻的两个陆地单元变成水单元，
** 从而使这个角落的陆地单元与原岛屿分离。而对于 1×n 类型的岛屿，我们也可以选择一个中间的陆地单元变成水单元使得陆地分离。
** 因此最终的答案只可能是 0,1,2。
** 转换为图 问题，就是
** 1. 如果 连通分量是 大于1， 则 直接返回 0 不需要操作
** 2. 如果 连通分量是 0，直接返回 0
** 3. 如果 连通分量 等于 1， 进行tarjan 计算割点，
	3.1 若 割点数 > 1, 则返回 1， 只需一步即可分开图
	3.2 若无割点， 则 直接返回 2
 */
/* 求割点
** 首先选定一个根节点，从该根节点开始遍历整个图（使用DFS）。
** 对于根节点，判断是不是割点很简单——计算其子树数量，如果有2棵即以上的子树，就是割点。因为如果去掉这个点，这两棵子树就不能互相到达。
	首先，“根节点有n棵子树”这句话，是说这n棵子树是独立的，没有根节点不能互相到达。因此n不一定等于与根节点相邻的顶点数。
	因此加入了vis[v]为false的条件，因为如果(u, v1)和(u, v2)在一棵子树里，
	对v1进行DFS，一定能去到v2，vis[v2]就会为true，此时就不会children++了。
** 对于非根节点，我们维护两个数组dfn[]和low[]，对于边(u, v)，如果low[v]>=dfn[u]，此时u就是割点。
	对于边(u, v)，如果low[v]>=dfn[u]，即v即其子树能够（通过非父子边）回溯到的最早的点，最早也只能是u，要到u前面就需要u的回边或u的父子边。
	也就是说这时如果把u去掉，u的回边和父子边都会消失，那么v最早能够回溯到的最早的点，已经到了u后面，无法到达u前面的顶点了，此时u就是割点。
*/
func minDays(grid [][]int) int {
	dirs := [][]int{[]int{0,1}, []int{0,-1}, []int{-1,0}, []int{1, 0}}
	graph := map[[2]int][][2]int{}
	for i := range grid{
		for j := range grid[i]{
			if grid[i][j] == 1{
				// 必须加上这个 建图, 否则会漏掉 上下左右均是 0 的点
				if graph[[2]int{i,j}] == nil{
					graph[[2]int{i,j}] = [][2]int{}
				}
				for _, d := range dirs{
					x, y := i + d[0], j + d[1]
					// 一定要判断有效性，否则 graph map里添加乱七八糟的
					if x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0]) && grid[x][y] == 1{
						//graph[[2]int{x,y}] = append(graph[[2]int{x,y}], [2]int{i,j}) 注意矩阵构建图，此题目中 重复了
						graph[[2]int{i,j}] = append(graph[[2]int{i,j}], [2]int{x,y})
					}
				}
			}
		}
	}
	// 注意： 特殊情况处理！！！！！！！！
	if len(graph) == 1{
		return 1
	}
	if len(graph) == 0{
		return 0
	}
	dfn, low := map[[2]int]int{}, map[[2]int]int{}
	for u := range graph{
		dfn[u] = 0
		low[u] = 0
	}
	ts := 1
	cuts := map[[2]int]bool{} // 求割点
	var trajan func(u, parent [2]int)
	trajan = func(u, parent [2]int){
		dfn[u], low[u] = ts, ts
		ts++
		children := 0
		//for v := range graph[u]{ 注意 这个地方 索引 和  值 经常搞混
		for _, v := range graph[u]{
			if v == parent{
				continue
			}
			if dfn[v] != 0{ // 处理回边
				low[u] = min(low[u], dfn[v])
			}else{
				children++
				trajan(v, u)
				low[u] = min(low[u], low[v])
				// 判断 割点
				if parent == [2]int{-1,-1} && children >= 2 {
					cuts[u] = true
				}else if parent != [2]int{-1,-1} && low[v] >= dfn[u]{
					cuts[u] = true
				}
			}
		}
	}
	// trajan 算法
	cnt := 0
	for u := range graph{
		if dfn[u] != 0{
			continue
		}
		cnt++
		if cnt > 1{
			return 0
		}
		trajan(u, [2]int{-1, -1})
	}
	if len(cuts) > 0{
		return 1
	}
	return 2
}

/* 1254. Number of Closed Islands
** Given a 2D grid consists of 0s (land) and 1s (water). 
** An island is a maximal 4-directionally connected group of 0s and a closed island is an island totally 
** (all left, top, right, bottom) surrounded by 1s.
** Return the number of closed islands.
 */
// 此方法 对 开放的边界的区域 也误计算进去了
func closedIsland(grid [][]int) int {
	g := map[int][]int{}
	m, n := len(grid), len(grid[0])
	dirs := [][]int{[]int{0,1}, []int{-1, 0}, []int{1, 0}, []int{0,-1}}
	valid := func(x, y int)bool{
		for _, d := range dirs{
			nx, ny := x+d[0], y+d[1]
			if nx < 0 || ny < 0 || nx >= m || ny >= n{
				return false
			}
		}
		return true
	}
	visited := map[int]bool{}
	for i := range grid{
		for j := range grid[i]{
			idx := i*n+j
			if grid[i][j] == 0 && g[idx] == nil{
				if valid(i, j){
					g[idx] = []int{}
					visited[idx] = false
				}
			}
		}
	}
	// 构造邻接表
	for node := range g{
		x, y := node/n, node%n
		for _, d := range dirs{
			nx, ny := x+d[0], y+d[1]
			if grid[nx][ny] == 0{
				g[node] = append(g[node], nx*n+ny)
			}
		}
	}
	var dfs func(node int)
	dfs = func(node int){
		visited[node] = true
		for _, v := range g[node]{
			if !visited[v]{
				dfs(v)
			}
		}
	}
	// dfs 求 连通分量
	ans := 0
	for node := range g{
		if !visited[node]{
			ans++
			dfs(node)
		}
	}
	return ans
}

func closedIsland_DFS(grid [][]int) int {
	g := map[int][]int{}
	m, n := len(grid), len(grid[0])
	dirs := [][]int{[]int{0,1}, []int{-1, 0}, []int{1, 0}, []int{0,-1}}
	valid := func(x, y int)bool{
		if x < 0 || y < 0 || x >= m || y >= n {
			return false
		}
		return true
	}
	visited := map[int]bool{}
	for i := range grid{
		for j := range grid[i]{
			idx := i*n+j
			if grid[i][j] == 0 && g[idx] == nil{
				g[idx] = []int{}
				visited[idx] = false
			}
		}
	}
	// 构造邻接表
	for node := range g{
		x, y := node/n, node%n
		for _, d := range dirs{
			nx, ny := x+d[0], y+d[1]
			if valid(nx, ny) && grid[nx][ny] == 0{
				g[node] = append(g[node], nx*n+ny)
			}
		}
	}
	var dfs func(node int)bool
	dfs = func(node int)bool{
		ret := true
		visited[node] = true
		for _, d := range dirs{
			x, y := node/n + d[0], node%n + d[1]
			if !valid(x, y){ // 不能直接返回false, 还需把其他节点遍历完
				ret = false
			}
		}
		for _, v := range g[node]{
			if !visited[v]{
				if !dfs(v) && ret{// 不能直接返回false, 还需把其他节点遍历完
					ret = false
				}
			}
		}
		return ret
	}
	// dfs 求 连通分量
	ans := 0
	for node := range g{
		if !visited[node]{
			if dfs(node){
				ans++
			}
		}
	}
	return ans
}

func closedIsland_matrix_dfs(grid [][]int) int {
	ans := 0
	dirs := [][]int{[]int{0,1}, []int{-1, 0}, []int{1, 0}, []int{0,-1}}
	m, n := len(grid), len(grid[0])
	valid := func(x, y int)bool{
		if x < 0 || y < 0 || x >= m || y >= n{
			return false
		}
		return true
	}
	var dfs func(x, y int)bool
	dfs = func(x, y int)bool{
		ret := true // 增加状态
		if !valid(x, y){
			return false
		}
		if grid[x][y] == 1{
			return true
		}
		grid[x][y] = 1 // 省了 visited
		for _, d := range dirs{
			/* 不能直接返回， 需要把 连通的 0 的字段 访问完毕
			            题目是 从连通分量中选 闭合的 连通分量子集
						if !dfs(x+d[0], y+d[1]){
			                return false
			            }*/
			if !dfs(x+d[0], y+d[1]) && ret{
				ret = false
			}
		}
		return ret
	}
	for i := 0; i < m; i++{
		for j := 0; j < n; j++{
			if grid[i][j] == 0{
				if dfs(i, j){
					ans++
				}
			}
		}
	}
	return ans
}














