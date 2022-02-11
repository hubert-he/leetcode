package Graph

import "math"

/* 329. Longest Increasing Path in a Matrix
** Given an m x n integers matrix, return the length of the longest increasing path in matrix.
** From each cell, you can either move in four directions: left, right, up, or down.
** You may not move diagonally or move outside the boundary (i.e., wrap-around is not allowed).
*/
// 此题 记忆化DFS 2021-12-09 刷出
// 每个单元格对应的最长递增路径的结果只和相邻单元格的结果有关, 但是 根据状态转移方程 没法提前锁定，因此处理 需要一些前提
// 计算当前节点的时候要先把相邻的比它大的节点先计算出来, 而此矩阵又是无序的，因此 无法 直接 DP 递推
// 额外的处理思路：
// 1. 把所有节点排个序，值较大的排前面，这样就可以保证从值较大的开始了
// 2. 使用拓扑排序的思想，把整个矩阵转换成有向无环图
// 这里采用拓扑排序， 将矩阵看成一个有向图，计算每个单元格对应的出度
// 对于作为边界条件的单元格，该单元格的值比所有的相邻单元格的值都要大，因此作为边界条件的单元格的出度都是 0
// 基于出度的概念，可以使用拓扑排序求解。
// 从所有出度为 0 的单元格开始广度优先搜索，每一轮搜索都会遍历当前层的所有单元格，更新其余单元格的出度，并将出度变为 0 的单元格加入下一层搜索。
// 当搜索结束时，搜索的总层数即为矩阵中的最长递增路径的长度。
// 有向图 构建： 节点值小的 ----> 节点值大的
// 拓展：
// 1. 基于拓扑排序对排序后的有向无环图做了层次遍历，如果没有拓扑排序直接进行广度优先搜索会发生什么
// 	不行。拓扑排序的目的是防止出现后效性，拓扑序就是dp 中阶段的顺序。如果换成 bfs，解法会变得很麻烦。
// 2. 给定一个整数矩阵，找出符合以下条件的路径的数量：这个路径是严格递增的，且它的长度至少是 3。矩阵的边长最大为 10的立方,答案对 10^9 + 7取模
//		其他要求与本题相同
/* 通用做法是通过 **矩阵快速幂求解** 复杂度是 $O(n^3*logn)$, n 是节点个数。在这里会超时。
（通用做法可以求路径长度为 k 的数量） 发现只需要求解长度至少是 3 的路径数量。
	公式转换，$ans = 路径总数量 - 长度为 1 的数量 - 长度为 2 的数量$
	路径总数量可以通过拓扑序求出来。时间复杂度为 $O(n+m)$， 类似于dp 求方案数。
	长度为 1 的路径数量就是 边数，时间复杂度为 $O(n)$
	长度为 2 的路径数量可以通过 长度为 1 的计算出来，时间复杂度为 $O(n)$
*/
func longestIncreasingPath(matrix [][]int) int {
	dirs := [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1}}
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0{
		return 0
	}
	outdegrees := make([][]int, m)
	for i := range outdegrees{
		outdegrees[i] = make([]int, n)
	}
	// 计算 出度
	for i := 0; i < m; i++{
		for j := 0; j < n; j++{
			for _, d := range dirs{
				x, y := i + d[0], j + d[1]
				if x >= 0 && x < m && y >= 0 && y < n && matrix[x][y] > matrix[i][j]{
					outdegrees[i][j]++
				}
			}
		}
	}
	// BFS 计算最大层数
	queue := [][]int{}
	for i := 0; i < m; i++{
		for j := 0; j < n; j++{
			if outdegrees[i][j] == 0{
				queue = append(queue, []int{i, j})
			}
		}
	}
	ans := 0
	for len(queue) > 0{
		ans++
		t := [][]int{}
		for i := range queue{
			x, y := queue[i][0], queue[i][1]
			for _, d := range dirs{
				xx, yy := x + d[0], y + d[1]
				if xx >= 0 && xx < m && yy >= 0 && yy < n &&  matrix[xx][yy] < matrix[x][y]{
					outdegrees[xx][yy]--
					if outdegrees[xx][yy] == 0{
						t = append(t, []int{xx, yy})
					}
				}
			}
		}
		queue = t
	}
	return ans
}
/*
1. 207
2. 210  Course Schedule II
3. 261. Graph Valid Tree
4. 310. Minimum Height Trees
5. 630. Course Schedule III
6. 269. Alien Dictionary
7. 444. Sequence Reconstruction
 */
/* 207. Course Schedule
** There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1.
** You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first
	if you want to take course ai.
** For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
** Return true if you can finish all courses. Otherwise, return false.
 */
//2022-02-08 刷出此题 方法一： 拓扑排序
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]bool, numCourses)
	indegree := make([]int, numCourses)
	for i := range graph{
		graph[i] = make([]bool, numCourses)
	}
	for i := range prerequisites{
		y, x := prerequisites[i][0], prerequisites[i][1]
		indegree[y]++
		graph[x][y] = true
	}
	q := []int{}
	for i := range indegree{
		if indegree[i] == 0{
			q = append(q, i)
		}
	}
	for len(q) > 0{
		t := []int{}
		for _,node := range q{
			for j := range graph[node]{
				if graph[node][j] {
					indegree[j]--
					if j != node && indegree[j] == 0{
						t = append(t, j)
					}
				}
			}
		}
		q = t
	}
	for i := range indegree{
		if indegree[i] != 0{
			return false
		}
	}
	return true
}
/* 求出一种拓扑排序方法的最优时间复杂度为 O(n+m)，其中 n 和 m 分别是有向图 G 的节点数和边数，
** 而判断图 G 是否存在拓扑排序，至少也要对其进行一次完整的遍历，时间复杂度也为 O(n+m)。
** 因此不可能存在一种仅判断图是否存在拓扑排序的方法，它的时间复杂度在渐进意义上严格优于 O(n+m)。
 */
/* DFS 本质是查找环
** 本题dfs困难点在与 节点的3个状态 与 子问题的关系
** 对于图中的任意一个节点，它在搜索的过程中有三种状态：
	1.「未搜索」：我们还没有搜索到这个节点；
	2. 搜索中」：我们搜索过这个节点，但还没有回溯到该节点，即该节点还没有入栈，还有相邻的节点没有搜索完成）
	3.「已完成」：我们搜索过并且回溯过这个节点，即该节点已经入栈，并且所有该节点的相邻节点都出现在栈的更底部的位置，满足拓扑排序的要求。
 */
func CanFinish_DFS(numCourses int, prerequisites [][]int) bool {
	// dfs 判断是否有环
	graph := make([][]int, numCourses)
	visited := make([]int, numCourses)// 列出3中状态，0--未被搜索  1--正在搜索  2--搜索完成
	for i := range prerequisites{
		src, dest := prerequisites[i][1], prerequisites[i][0]
		graph[src] = append(graph[src], dest)
	}

	var dfs func(node int)bool
	dfs = func(node int)bool {
		visited[node] = 1 // 搜索中
		for _, v := range graph[node]{
			if visited[v] == 0{
				if !dfs(v){
					return false
				}
			}else if visited[v] == 1{
				return false
			}
		}
		visited[node] = 2 // 搜索完成
		return true
	}
	for i := range graph{
		if visited[i] == 0 && !dfs(i){
			return false
		}
	}
	return true
}

/* 310. Minimum Height Trees
** A tree is an undirected graph in which any two vertices are connected by exactly one path.
** In other words, any connected graph without simple cycles is a tree.
** Given a tree of n nodes labelled from 0 to n - 1,
** and an array of n - 1 edges where edges[i] = [ai, bi] indicates that
** there is an undirected edge between the two nodes ai and bi in the tree, you can choose any node of the tree as the root.
** When you select a node x as the root, the result tree has height h.
** Among all possible rooted trees, those with minimum height (i.e. min(h))  are called minimum height trees (MHTs).
** Return a list of all MHTs' root labels. You can return the answer in any order.
** The height of a rooted tree is the number of edges on the longest downward path between the root and a leaf.
 */
/* 简单BFS 超时，
** 对每个节点遍历bfs，统计下每个节点的高度
 */
func findMinHeightTrees(n int, edges [][]int) []int {
	ans := []int{}
	minH := math.MaxInt32
	g := make([][]int, n)
	for i := range edges{
		x, y := edges[i][0], edges[i][1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	vis := make([]bool, n)
	// 无向图有个陷阱，即如何判断 上一个节点是否访问过
	bfs := func(u int)(h int){
		for i := range vis {
			vis[i] = false
		}
		q := []int{u}
		vis[u] = true
		for len(q) > 0{
			t := []int{}
			h++
			if h > minH{
				break
			}
			for i := range q{
				for _, v := range g[q[i]]{
					if !vis[v]{
						t = append(t, v)
						vis[v] = true
					}
				}
			}
			q = t
		}
		return
	}
	for i := range g{
		height := bfs(i)
		if height < minH{
			ans = []int{i}
			minH = height
		}else if height == minH{
			ans = append(ans, i)
		}
	}
	return ans
}
// 拓扑排序来优化， 将一个树 写成拓扑排序，发现答案在第一层，即多个拓扑排序的首节点
// 错误解法: 无法处理：  4
//					 [[1,0],[1,2],[1,3]]
func findMinHeightTrees_bfs(n int, edges [][]int) []int {
	g := make([][]int, n)
	indegree := make([]int, n)
	for i := range edges{
		x, y := edges[i][0], edges[i][1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
		indegree[x]++
		indegree[y]++
	}
	q := []int{}
	for i := range indegree{
		if indegree[i] == 1{
			q = append(q, i)
		}
	}
	ans := q
	for len(q) > 0{
		for _, u := range q{
			for _, v := range g[u]{// 问题就出在这里，不需要始终为何indegree
				indegree[u]--
				indegree[v]--
			}
		}
		ans = q
		q = []int{}
		for i := range indegree{
			if indegree[i] == 1{
				q = append(q, i)
			}
		}
	}
	return ans
}
// 拓扑排序解法-- 注意 此题经典在于 是求排序的首节点
func FindMinHeightTrees(n int, edges [][]int) []int {
	if n == 1{// 易错点-1： 只有一个点的情况
		return []int{0}
	}
	g := make([][]int, n)
	indegree := make([]int, n)
	for i := range edges{
		x, y := edges[i][0], edges[i][1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
		indegree[x]++
		indegree[y]++
	}
	q := []int{}
	for i := range indegree{
		if indegree[i] == 1{
			q = append(q, i)
		}
	}
	ans := q
	for len(q) > 0{
		t := []int{}
		for _, u := range q{
			for _, v := range g[u]{
				indegree[v]--
				if indegree[v] == 1{ // 易错点-2： 发现度为1 就加入
					t = append(t, v)
				}
			}
		}
		ans = q
		q = t
	}
	return ans
}
// 这是树形 DP 解法， down 的高度计算很好处理， up 高度计算 需要多理解
// 需要多理解学习这个方法
func FindMinHeightTreesDP(n int, edges [][]int) []int {
	g := make([][]int, n)
	for i := range edges{
		x, y := edges[i][0], edges[i][1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	var dfs_d func(u, from int) int
	var dfs_u func(u, from int)
	up := make([]int, n) // up[i]表示结点i往上走的最大距离
	d1 := make([]int, n) // d1[i]表示结点i往下走的最大距离
	d2 := make([]int, n) // d2[i]表示结点i往下走的次大距离
	p  := make([]int, n) // p[i]表示结点i往下走的最长路经过p[i]
	for i := range p{
		p[i] = -1
	}
	dfs_d = func(u, from int) int{
		d1[u], d2[u] = 0, 0
		for _, v := range g[u]{
			if v == from { continue }
			d := dfs_d(v, u) + 1
			if d > d1[u]{
				//d1[u] = d
				//d2[u] = d1[u] 错在这
				d2[u] = d1[u]
				d1[u] = d
				p[u] = v	// 记录最长路时还需记录当前最长路是往哪个子结点走的
			}else if d > d2[u]{
				d2[u] = d
			}
		}
		return d1[u]
	}
	dfs_u = func(u, from int) {// 向上的值 通过向下的得知
		for _, v := range g[u]{
			if v == from { continue }
			if p[u] == v{
				up[v] = max(d2[u], up[u]) + 1
			}else{
				up[v] = max(d1[u], up[u]) + 1
			}
			dfs_u(v, u)
		}
	}
	dfs_d(0, -1)
	dfs_u(0, -1)
	dist := map[int][]int{}
	minH := math.MaxInt32
	for i := 0; i < n; i++{
		d := max(up[i], d1[i])
		if minH > d{
			minH = d
		}
		dist[d] = append(dist[d], i)
	}
	return dist[minH]
}

/* 444. Sequence Reconstruction
** You are given an integer array nums of length n where nums is a permutation of the integers in the range [1, n].
** You are also given a 2D integer array sequences where sequences[i] is a subsequence of nums.
** Check if nums is the shortest possible and the only supersequence.
** The shortest supersequence is a sequence with the shortest length and has all sequences[i] as subsequences.
** There could be multiple valid supersequences for the given array sequences.
** For example, for sequences = [[1,2],[1,3]], there are two shortest supersequences, [1,2,3] and [1,3,2].
** While for sequences = [[1,2],[1,3],[1,2,3]], the only shortest supersequence possible is [1,2,3].
** [1,2,3,4] is a possible supersequence but not the shortest.
** Return true if nums is the 「only」 「shortest」 supersequence for sequences, or false otherwise.
** A subsequence is a sequence that can be derived from another sequence by deleting some or no elements
** without changing the order of the remaining elements.
 */
// 2022-02-10 刷出此题
// 存在易错点
// sequences 并不是 [1,2] 这种2元素， 可以是[5,2,6,3] 转换为 [5,2] [2,6] [6,3] 这种
// 题目潜在含义：  拓扑排序结果一定唯一  要求Only
func sequenceReconstruction(nums []int, sequences [][]int) bool {
	n := len(nums)
	g := make([][]int, n)
	indegree := make([]int, n)
	/* 易错点-1： sequences 并不是 [1,2]  可以是[5,2,6,3] 转换为 [5,2] [2,6] [6,3]
	for _, u := range sequences{
		src, dest := u[0]-1, u[1]-1
		g[src] = append(g[src], dest)
		indegree[dest]++
	}
	 */
	for _, u := range sequences{
		for j := 0; j < len(u); j++{
			if j+1 >= len(u){
				break
			}
			src, dest := u[j]-1, u[j+1]-1
			g[src] = append(g[src], dest)
			indegree[dest]++
		}
	}

	//q := []int{}
	q := map[int]bool{} // 但是题目要求是 only 所以选 集合 是不合适的
	idx := 0
	for i := range indegree{
		if indegree[i] == 0 {
			//q = append(q, i)
			q[i] = true
		}
	}
	/* 易错点-2： 注意拓扑排序的含义
	for len(q) > 0{
		ok := false
		t := []int{}
		for _, u := range q {
			if u+1 == nums[idx]{
				idx++
				ok = true
			}
			for _, v := range g[u]{
				indegree[v]--
				if indegree[v] == 0{
					t = append(t, v)
				}
			}
		}
		if !ok {
			return false
		}
		q = t
	}
	 */
	for len(q) > 0{
		t := map[int]bool{}
		/* 题目要求的是 Only
		nq := len(q)
		for i := 0; i < nq; i++{
			if !q[nums[idx+i]-1]{
				return false
			}
		}
		idx += nq
		*/
		if len(q) > 1{	return false }
		idx++
		for u := range q{
			for _, v := range g[u]{
				indegree[v]--
				if indegree[v] == 0{
					t[v] = true
				}
			}
		}
		q = t
	}
	return true
}

/* 269. Alien Dictionary
** There is a new alien language that uses the English alphabet. However, the order among the letters is unknown to you.
** You are given a list of strings words from the alien language's dictionary,
** where the strings in words are sorted lexicographically by the rules of this new language.
** Return a string of the unique letters in the new alien language sorted in lexicographically increasing order
** by the new language's rules. If there is no solution, return "". If there are multiple solutions, return any of them.
** A string s is lexicographically smaller than a string t if at the first letter where they differ,
** the letter in s comes before the letter in t in the alien language.
** If the first min(s.length, t.length) letters are the same, then s is smaller if and only if s.length < t.length.
 */
// 2022-02-10  刷出此题
// 有个case 理解不清： ["vlxpwiqbsg","cpwqwqcd"]  结果："bdgilpqsvcwx" 任意顺序
// 难点2个： 想到拓扑排序解决  难点-2 异常条件判断很多，需要把控边界
func AlienOrder(words []string) string {
	n := len(words)
	ans := []byte{}
	if n == 0{
		return string(ans)
	}
	if n <= 1{
		//return words[0]
		mg := map[byte]bool{}
		for i := range words[0]{
			if !mg[words[0][i]]{
				ans = append(ans, words[0][i])
				mg[words[0][i]] = true
			}
		}
		return string(ans)
	}
	g := map[byte]map[byte]bool{}
	indegree := map[byte]int{}
	for i := range words{
		for j := i+1; j < n; j++{
			ni, nj := len(words[i]), len(words[j])
			equal := true
			k := 0
			for ; k < ni && k < nj; k++{
				src, dest := words[i][k], words[j][k]
				if g[src] == nil{
					g[src] = map[byte]bool{}
					indegree[src] = 0 // 入度表示用map 跟 slice 的区别
				}
				if g[dest] == nil{
					g[dest] = map[byte]bool{}
					indegree[dest] = 0// 入度表示用map 跟 slice 的区别
				}
				if src != dest{
					equal = false
					if !g[src][dest]{
						g[src][dest] = true
						indegree[dest]++
					}
					// g[src][dest] = true
					// indegree[dest]++ 重复+1
					break
				}
			}
			if ni > nj && equal{// 额外判断大小是否合理
				return ""
			}
			for m := k;m < ni; m++{
				src := words[i][m]
				if g[src] == nil{
					indegree[src] = 0
				}
			}
			for m := k;m < nj; m++{
				src := words[j][m]
				if g[src] == nil{
					indegree[src] = 0
				}
			}
			if ni < nj{
				for idx := ni; idx < nj; idx++{
					src := words[j][idx]
					if g[src] == nil{
						indegree[src] = 0
					}
				}
			}else if ni > nj{
				for idx := ni; idx < nj; idx++{
					src := words[i][idx]
					if g[src] == nil{
						indegree[src] = 0
					}
				}
			}
		}
	}
	q := []byte{}
	for i := range indegree{
		if indegree[i] == 0{
			q = append(q, i)
		}
	}
	ans = q
	for len(q) > 0{
		t := []byte{}
		for _, u := range q{
			for v := range g[u]{
				indegree[v]--
				if indegree[v] == 0{
					t = append(t, v)
					ans = append(ans, v)
				}
			}
		}
		q = t
	}
	// 判断环
	for i := range indegree{
		if indegree[i] != 0{
			return ""
		}
	}
	return string(ans)
}
















