package Graph

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