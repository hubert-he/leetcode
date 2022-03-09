package graph

// 547. Number of Provinces
/*  统计连通分量 方法有dfs 和 并查集
	连通的：无向图中每一对不同的顶点之间都有路径。如果这个条件在有向图里也成立，那么就是强连通的。
	连通分量：不连通的图是由2个或者2个以上的连通子图组成的。这些不相交的连通子图称为图的连通分量
	在深度优先搜索的递归调用期间，只要是某个顶点的可达顶点都能在一次递归调用期间访问到。
	有向图的连通分量：如果某个有向图不是强连通的，而将它的方向忽略后，任何两个顶点之间总是存在路径，则该有向图是弱连通的。
 			 	  若有向图的子图是强连通的，且不包含在更大的连通子图中，则可以称为有向图的强连通分量。
 */
func findCircleNum(isConnected [][]int) (connectedComponentCount int) {
	visited := make([]bool, len(isConnected))
	var dfs func(node int)
	dfs = func(node int){
		visited[node] = true
		for target, conn := range isConnected[node]{
			if conn == 1 && !visited[target]{
				dfs(target)
			}
		}
	}
	for i, v := range visited {
		if !v {
			connectedComponentCount++
			dfs(i)
		}
	}
	return
}
// 1319. Number of Operations to Make Network Connected
/*
使用深度优先搜索来得到图中的连通分量数。
具体地，初始时所有节点的状态均为「待搜索」。我们每次选择一个「待搜索」的节点，
从该节点开始进行深度优先搜索，并将所有搜索到的节点的状态更改为「已搜索」，这样我们就找到了一个连通分量
 */
func makeConnected(n int, connections [][]int) (ans int) {
	if len(connections) < n - 1{
		return -1
	}
	// 邻接矩阵: 压缩过
	graph := make([][]int, n)
	for _, c := range connections{
		v1, v2 := c[0], c[1]
		graph[v1] = append(graph[v1], v2)
		graph[v2] = append(graph[v2], v1)
	}
	// visited 状态标记
	visited := make([]bool, n)
	var dfs func(int)
	dfs = func(from int){
		visited[from] = true
		for _, to := range graph[from]{
			if !visited[to]{
				dfs(to)
			}
		}
	}
	for i, v := range visited{
		if !v{
			ans++
			dfs(i)
		}
	}
	return ans - 1
}

/* 1557. Minimum Number of Vertices to Reach All Nodes
** Given a directed acyclic graph（有向无环图）, with n vertices numbered from 0 to n-1, 
** and an array edges where edges[i] = [fromi, toi] represents a directed edge from node fromi to node toi.
** Find the smallest set of vertices from which all nodes in the graph are reachable.
** It's guaranteed that a unique solution exists.
** Notice that you can return the vertices in any order.
 */
// 2022-02-22 刷出此题
/* 图思维的运用
** 对于任意节点 x 如果其入度不为零，则一定存在节点 y 指向节点 x
** 从节点 y 出发即可到达节点 y 和 x 因此如果从节点 y 出发， 节点 x 和 节点 y 都可以到达，
** 且从节点 y 出发可以到达的节点比从节点 x 出发可以到达的节点更多。
** 由于给定的图是有向无环图，基于上述分析可知，
** 对于任意入度不为零的节点 x，一定存在另一个节点 z，使得从节点 z 出发可以到达节点 x。
** 为了获得最小的点集，只有入度为零的节点才应该加入最小的点集。
	由于入度为零的节点必须从其自身出发才能到达该节点，从别的节点出发都无法到达该节点，因此最小的点集必须包含所有入度为零的节点。
	因为入度不为零的点总可以由某个入度为零的点到达，所以这些点不包括在最小的合法点集当中。
	因此，我们得到「最小的点集使得从这些点出发能到达图中所有点」就是入度为零的所有点的集合。
*/
func findSmallestSetOfVertices(n int, edges [][]int) []int {
	in := make([]bool, n)
	for i := range edges{
		in[edges[i][1]]=true
	}
	ans := []int{}
	for i := range in{
		if !in[i]{
			ans = append(ans, i)
		}
	}
	return ans
}