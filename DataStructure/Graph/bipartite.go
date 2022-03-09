package Graph

/* 二分图 部分
** 二分图定义：
	二分图又称作二部图，是图论中的一种特殊模型。
	设G=(V,E)是一个无向图，如果顶点V可分割为两个互不相交的子集(A,B)，
	并且图中的每条边（i，j）所关联的两个顶点i和j分别属于这两个不同的顶点集(i in A, j in B)，则称图G为一个二分图。
	简而言之，就是顶点集V可分割为两个互不相交的子集，并且图中每条边依附的两个顶点都分属于这两个互不相交的子集，两个子集内的顶点不相邻。
** 二分图匹配：
	给定一个二分图G，在G的一个子图M中，M的边集{E}中的任意两条边都不依附于同一个顶点，则称M是一个匹配。
	极大匹配(Maximal Matching)是指在当前已完成的匹配下,无法再通过增加未完成匹配的边的方式来增加匹配的边数。
	最大匹配(maximum matching)是所有极大匹配当中边数最大的一个匹配。选择这样的边数最大的子集称为图的最大匹配问题。
	如果一个匹配中，图中的每个顶点都和图中某条边相关联，则称此匹配为完全匹配，也称作完备匹配。
	求二分图匹配可以用最大流(Maximal Flow)或者匈牙利算法(Hungarian Algorithm)。
*/
/* 二分图常见应用
** 在二分图上，最常见的问题是求 二分图最大匹配/最大独立集。二分图的最大匹配，简单来说就是找到尽可能多的两两匹配的关系。
** 比如说在某场相亲现场，主持人先让所有男生女生都写好自己感兴趣的若干个异性，主持人需要尽可能多的凑一对对男女去约会。
 */

/* 785. Is Graph Bipartite?
** 判断二分图
**
 */
// 2022-03-09 刷出此题
// 有个Golang 注意事项，即 slice 只能与 nil 做相等与否比较
// 所以 这里使用了指针来比较是否属于哪些集合，这种方式 只能在刷题中使用
// 因为slice 的底层结构 在golang中是自动改变的，所以如果再项目中使用此方式，可能有程序崩溃的隐患，属于不安全编程
func isBipartite_DFS(graph [][]int) bool {
	n := len(graph)
	vis := make([]bool, n)
	s1, s2 := make([]bool, n), make([]bool, n)
	var dfs func(node int, s *[]bool)bool
	dfs = func(node int, s *[]bool)bool{
		vis[node] = true
		(*s)[node] = true
		for _, x := range graph[node]{
			if (*s)[x] { return false }
			t := &s1
			if s == &s1{ t = &s2 }
			if !vis[x] && !dfs(x, t){
				return false
			}
		}
		return true
	}
	for i := range graph{
		if !vis[i] && !dfs(i, &s1){
			return false
		}
	}
	return true
}
// 官方解答：染色法
/* 使用图搜索算法从各个连通域的任一顶点开始遍历整个连通域，遍历的过程中用两种不同的颜色对顶点进行染色，相邻顶点染成相反的颜色。
** 这个过程中倘若发现相邻的顶点被染成了相同的颜色，说明它不是二分图；
** 反之，如果所有的连通域都染色成功，说明它是二分图
 */
func isBipartite(graph [][]int) bool {
	const UNCOLORED, RED, BLUE = 0, 1, 2
	n := len(graph)
	color := make([]int, n)
	var dfs func(node, set int)bool
	dfs = func(node, set int)bool{
		color[node] = set
		setCopy := set
		if set == RED{
			set = BLUE
		}else{
			set = RED
		}
		for _, neighbor := range graph[node]{
			if color[neighbor] == setCopy{
				return false
			}
			if color[neighbor] == UNCOLORED{
				if !dfs(neighbor, set){
					return false
				}
			}
		}
		return true
	}
	for i := 0; i < n; i++{
		if color[i] == UNCOLORED{
			if !dfs(i, RED){
				return false
			}
		}
	}
	return true
}
/* Golang for-range 语法糖
** for-range其实是语法糖，内部调用还是for循环，初始化会拷贝带遍历的列表（如array，slice，map），
** 然后每次遍历的v都是对同一个元素的遍历赋值。 也就是说如果直接对v取地址，最终只会拿到一个地址，而对应的值就是最后遍历的那个元素所附给v的值
 */
func isBipartite_BFS(graph [][]int) bool {
	const UNCOLORED, RED, BLUE = 0, 1, 2
	n := len(graph)
	color := make([]int, n)
	for i := 0; i < n; i++{
		if color[i] == UNCOLORED{
			q := []int{i}
			color[i] = RED
			// for j := range q{
			// for _, guy := range q{   以上2种写法是错误的， golang 的 for-range 会直接把q 拷贝出来
			for j := 0; j < len(q); j++{
				targetColor := RED
				if color[q[j]] == RED{
					targetColor = BLUE
				}
				for _, neigh := range graph[q[j]]{
					if color[neigh] == color[q[j]] {
						return false
					}
					if color[neigh] == UNCOLORED{
						q = append(q, neigh)
						color[neigh] = targetColor
					}
				}
			}
		}
	}
	return true
}
// 并查集
func isBipartite_ufs(graph [][]int) bool {
	n := len(graph)
	ufs := make([]int, n)
	for i := range ufs{
		ufs[i] = i
	}
	find := func(x int)int{
		px := ufs[x]
		for px != x{
			ufs[x] = ufs[px] // 路径压缩
			x = px
			px = ufs[x]
		}
		return px
	}
	union := func(x, y int){
		px, py := find(x), find(y)
		ufs[px] = py
	}
	for i := range graph{
		length := len(graph[i])
		if  length <= 0 { continue }
		pi, p0 := find(i), find(graph[i][0])
		for j := 1; j < length; j++{
			// 如果相邻节点已经在一个集合中个，则表示非二分图
			pj := find(graph[i][j])
			if pi == pj { return false }
			union(p0, pj)
		}
	}
	return true
}