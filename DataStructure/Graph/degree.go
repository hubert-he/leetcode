package Graph

import "sort"

/* 1615. Maximal Network Rank
** There is an infrastructure of n cities with some number of roads connecting these cities.
** Each roads[i] = [ai, bi] indicates that there is a bidirectional road between cities ai and bi.
** The network rank of two different cities is defined as the total number of directly connected roads to either city.
** If a road is directly connected to both cities, it is only counted once.
** The maximal network rank of the infrastructure is the maximum network rank of all pairs of different cities.
** Given the integer n and the array roads, return the maximal network rank of the entire infrastructure.
 */
// 2022-03-09 刷出此题
// 此题考察点有2个：
//	一个是要想到是 图节点的 度
//  另一个要点是 排序数组下 求最大和，但是这个和 有限制
func maximalNetworkRank(n int, roads [][]int) int {
	type Node struct{
		degree  int
		id      int
	}
	g := make([][]bool, n)
	for i := range g {
		g[i] = make([]bool, n)
	}
	in := make([]Node, n)
	for _, con := range roads{
		g[con[0]][con[1]], g[con[1]][con[0]] = true, true
		in[con[0]].id = con[0]
		in[con[0]].degree++
		in[con[1]].id = con[1]
		in[con[1]].degree++
	}
	sort.Slice(in, func(i, j int)bool{return in[i].degree > in[j].degree})
	ans := in[0].degree + in[1].degree - 1
	for i := 0; i < n; i++{
		for j := i+1; j < n; j++{
			rank := in[i].degree + in[j].degree
			if g[in[i].id][in[j].id]{ rank -= 1 }
			if ans < rank { return rank }
			if ans > rank{ // 排序后，提前判断
				break
			}
		}
	}
	return ans
}
// 方法二： 直接平方级查询， 带无向图矩阵压缩
func maximalNetworkRank2(n int, roads [][]int) int {
	g := make([][]bool, n)
	for i := range g{
		//g[i] = make([]bool, n)  无向图稀疏矩阵压缩
		g[i] = make([]bool, i+1)
	}
	in := make([]int, n)
	for i := range roads{
		n1, n2 := roads[i][0], roads[i][1]
		// g[n1][n2], g[n2][n1] = true, true
		if n1 < n2{
			g[n2][n1] = true
		}else{
			g[n1][n2] = true
		}
		in[n1], in[n2] = in[n1]+1, in[n2]+1
	}
	ans := 0
	for i := 0; i < n; i++{
		for j := i+1; j < n; j++{
			rank := in[i] + in[j]
			//if g[i][j]{ rank -= 1 }
			if g[j][i] { rank -= 1 }
			if ans < rank{
				ans = rank
			}
		}
	}
	return ans
}