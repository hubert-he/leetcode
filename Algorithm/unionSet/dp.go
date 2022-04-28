package unionSet

/* 305. Number of Islands II
** You are given an empty 2D binary grid grid of size m x n.
** The grid represents a map where 0's represent water and 1's represent land.
** Initially, all the cells of grid are water cells (i.e., all the cells are 0's).
** We may perform an add land operation which turns the water at position into a land.
** You are given an array positions where positions[i] = [ri, ci] is the position (ri, ci) at which we should operate the ith operation.
** Return an array of integers answer where answer[i] is the number of islands after turning the cell (ri, ci) into a land.
** An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
** You may assume all four edges of the grid are all surrounded by water.
 */
/*  2022-04-12 刷出此题
** 考虑这个case： 相同位置多次添加
	3
	3
	[[0,0],[0,1],[1,2],[1,2]]
 */
func numIslands2(m int, n int, positions [][]int) []int {
	dirs := [][]int{[]int{1, 0}, []int{-1, 0}, []int{0,1}, []int{0,-1}}
	ufs := make([]int, m*n)
	for i := range ufs{
		ufs[i] = -1
	}
	find := func(x int)int{
		px := ufs[x]
		for px != x {
			ufs[x] = ufs[px]
			x = px
			px = ufs[x]
		}
		return px
	}
	union := func(x, y int){
		px, py := find(x), find(y)
		if px != py{
			ufs[px] = py
		}
	}
	isvalid := func(x, y int)bool{
		if x < 0 || y < 0 || x >= m || y >= n{
			return false
		}
		return true
	}
	ans := make([]int, len(positions)+1)
	for i := range positions{
		x, y := positions[i][0], positions[i][1]
		idx := x*n + y
		if ufs[idx] == -1{
			ufs[idx] = idx
			ans[i+1] = ans[i] + 1
		}else{
			ans[i+1] = ans[i]
			//continue
		}
		// 提前预加1 对特别case  [0,0],[0,0] 情况不适用
		//ans[i+1] = ans[i] + 1
		for _, d := range dirs{
			if !isvalid(x+d[0], y+d[1]){ continue }
			neigh := (x+d[0])*n + y+d[1]
			if ufs[neigh] == -1{ continue }
			if find(neigh) != find(idx){
				union(idx, neigh)
				ans[i+1]--
			}
		}
	}
	return ans[1:]
}
