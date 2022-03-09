package Array

/* 1001. Grid Illumination 网格照明
** There is a 2D grid of size n x n where each cell of this grid has a lamp that is initially turned off.
** You are given a 2D array of lamp positions lamps,
** where lamps[i] = [rowi, coli] indicates that the lamp at grid[rowi][coli] is turned on.
** Even if the same lamp is listed more than once, it is turned on.
** When a lamp is turned on, it illuminates its cell and all other cells in the same row, column, or diagonal.
** You are also given another 2D array queries, where queries[j] = [rowj, colj].
** For the jth query, determine whether grid[rowj][colj] is illuminated or not.
** After answering the jth query, turn off the lamp at grid[rowj][colj] and its 8 adjacent lamps if they exist.
** A lamp is adjacent if its cell shares either a side or corner with grid[rowj][colj].
** Return an array of integers ans, where ans[j] should be 1 if the cell in the jth query was illuminated, or 0 if the lamp was not.
Constraints:
	1 <= n <= 10的9次方
	0 <= lamps.length <= 20000
	0 <= queries.length <= 20000
	lamps[i].length == 2
	0 <= rowi, coli < n
	queries[j].length == 2
	0 <= rowj, colj < n
*/
// 2022-01-21 刷出此题，但是OOM
func gridIllumination(n int, lamps [][]int, queries [][]int) []int {
	//status := make([][]bool, n) bool 不能表示多个光源叠加的
	status := make([][]int, n)
	light := map[[2]int]bool{}
	for i := range status{
		status[i] = make([]int, n)
	}
	clear := func(r, c, stat int){
		dirs := [][2]int{[2]int{1,1}, [2]int{-1,1}, [2]int{1,-1}, [2]int{-1,-1}}
		status[r][c] += stat
		for i := 0; i < n; i++{
			if i != c{
				status[r][i] += stat
			}
			if i != r{
				status[i][c] += stat
			}
			if status[r][i] < 0{ status[r][i] = 0}
			if status[i][c] < 0{ status[i][c] = 0}
		}
		for _, d := range dirs{
			for x, y := r+d[0],c+d[1]; x >= 0 && y >= 0 && x < n && y < n; x, y = x+d[0], y+d[1]{
				status[x][y] += stat
				if status[x][y] < 0 { status[x][y] = 0 }
			}
		}
		//fmt.Println(status)
	}
	close := func(x, y int){
		dirs := [][2]int{
			[2]int{0,0},
			[2]int{1,0}, [2]int{-1,0}, [2]int{0,1}, [2]int{0,-1},
			[2]int{1,1}, [2]int{-1,1}, [2]int{1,-1}, [2]int{-1,-1},
		}
		for _, d := range dirs{
			pos := [2]int{x+d[0], y+d[1]}
			if light[pos]{
				light[pos] = false // 漏掉关掉
				clear(pos[0], pos[1], -1)
			}
		}
	}
	for _, pos := range lamps{
		if !light[[2]int{pos[0],pos[1]}]{ // lams 重复打开
			light[[2]int{pos[0],pos[1]}] = true
			clear(pos[0], pos[1], 1)
		}
	}
	//fmt.Println(status)
	ans := []int{}
	for _, q := range queries{
		r, c := q[0], q[1]
		if status[r][c] > 0{
			ans = append(ans, 1)
		}else{
			ans = append(ans, 0)
		}
		close(r, c)
		//fmt.Println(status)
	}
	return ans
}
/* 题目中给出的N的大小最大是10的九次方, 数据量大
** 设计4个集合， 每点亮一个灯，就必然会点亮“一行”，“一列”，“一个正对角线”，“一个反对角线”
** 不再需要去记录每一个格子的“光源亮度”
** 我们查看一个格子的光源亮度，也只需要去看“这一行是否被点亮”，“这一列是否被点亮”，“这一个正对角线是否被点亮”
** 只要这些条件中，任一一个满足，就说明这个格子是被点亮的
** 值得学习的地方，就是矩阵 对角线 的映射
** 还是 因为 n 太大，当 n = 10的9次方 依然 OOM
 */
func gridIllumination_2(n int, lamps [][]int, queries [][]int) []int {
	//1. n*n矩阵 对角线总个数为 2*n-1个
	//2. diagonal1表示的是反手对角线，坐标[x,y] 映射线性关系为：y-x+(n-1)
	//3. diagonal2表示的是正手对角线，坐标[x,y] 映射线性关系为：x+y        注意：均是从 0 开始计
	line, column, diagonal1, diagonal2 := make([]int, n), make([]int, n), make([]int, 2*n-1), make([]int, 2*n-1)
	light := map[[2]int]bool{}
	clear := func(x, y, stat int){
		line[x] += stat
		column[y] += stat
		diag := y - x
		if diag <= n-2 && diag >= 2-n{
			diagonal1[diag+n-1] += stat
		}
		diag = x + y
		diagonal2[diag] += stat
	}
	open := func(x, y int){
		light[[2]int{x, y}] = true
		clear(x, y, 1)
	}
	close := func(x, y int){
		light[[2]int{x, y}] = false
		clear(x, y, -1)
	}
	for _, pos := range lamps{
		if !light[[2]int{pos[0],pos[1]}]{ //遗漏点-1： lamps 有 重复打开的情况
			open(pos[0],pos[1])
		}
	}
	diagonal1[0], diagonal1[2*n-2], diagonal2[0], diagonal2[2*n-2] = 0,0,0,0
	dirs := [][2]int{
		[2]int{0,0},
		[2]int{1,0}, [2]int{-1,0}, [2]int{0,1}, [2]int{0,-1},
		[2]int{1,1}, [2]int{-1,1}, [2]int{1,-1}, [2]int{-1,-1},
	}
	ans := []int{}
	for _, q := range queries{
		r, c := q[0], q[1]
		if line[r] > 0 || column[c] > 0 || diagonal1[c-r+n-1] > 0 || diagonal2[r+c] > 0{
			//fmt.Println(r, c)
			//fmt.Println(line[r] >0, column[c], diagonal1[c-r+n-2], diagonal2[r+c])
			ans = append(ans, 1)
		}else{
			ans = append(ans, 0)
		}
		for _, d := range dirs{
			x, y := r + d[0], c + d[1]
			if light[[2]int{x,y}]{
				close(x, y)
			}
		}
	}
	return ans
}
// 利用map 状态压缩
func gridIllumination_3(n int, lamps [][]int, queries [][]int) []int {
	line, column, diagonal1, diagonal2 := map[int]int{},map[int]int{},map[int]int{},map[int]int{}
	light := map[[2]int]bool{}
	clear := func(x, y, stat int){
		line[x] += stat
		column[y] += stat
		diag := y - x
		if diag <= n-2 && diag >= 2-n{
			diagonal1[diag+n-1] += stat
		}
		diag = x + y
		diagonal2[diag] += stat
	}
	open := func(x, y int){
		light[[2]int{x, y}] = true
		clear(x, y, 1)
	}
	close := func(x, y int){
		light[[2]int{x, y}] = false
		clear(x, y, -1)
	}
	for _, pos := range lamps{
		if !light[[2]int{pos[0],pos[1]}]{ //遗漏点-1： lamps 有 重复打开的情况
			open(pos[0],pos[1])
		}
	}
	diagonal1[0], diagonal1[2*n-2], diagonal2[0], diagonal2[2*n-2] = 0,0,0,0
	dirs := [][2]int{
		[2]int{0,0},
		[2]int{1,0}, [2]int{-1,0}, [2]int{0,1}, [2]int{0,-1},
		[2]int{1,1}, [2]int{-1,1}, [2]int{1,-1}, [2]int{-1,-1},
	}
	ans := []int{}
	for _, q := range queries{
		r, c := q[0], q[1]
		if line[r] > 0 || column[c] > 0 || diagonal1[c-r+n-1] > 0 || diagonal2[r+c] > 0{
			//fmt.Println(r, c)
			//fmt.Println(line[r] >0, column[c], diagonal1[c-r+n-2], diagonal2[r+c])
			ans = append(ans, 1)
		}else{
			ans = append(ans, 0)
		}
		for _, d := range dirs{
			x, y := r + d[0], c + d[1]
			if light[[2]int{x,y}]{
				close(x, y)
			}
		}
	}
	return ans
}

/* 1905. Count Sub Islands
** You are given two m x n binary matrices grid1 and grid2 containing only 0's (representing water) and 1's (representing land).
** An island is a group of 1's connected 4-directionally (horizontal or vertical).
** Any cells outside of the grid are considered water cells.
** An island in grid2 is considered a sub-island if there is an island in grid1 that contains all the cells that make up this island in grid2.
** Return the number of islands in grid2 that are considered sub-islands.
 */
// 2022-02-18 刷出此题，收录此题的目的是 有个坑点，就是 当发现此岛屿不是grid1的子岛屿后，不能直接放弃遍历，否则会影响总结果
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m, n := len(grid1), len(grid1[0])
	ans := 0
	dirs := [][]int{ []int{1, 0},[]int{-1, 0},[]int{0, 1},[]int{0, -1} }
	isValid := func(x, y int)bool{
		if x < 0 || y < 0 || x >= m || y >= n{
			return false
		}
		return true
	}
	var dfs func(x, y int)bool
	dfs = func(x, y int)bool{
		ret := true
		if !isValid(x,y){ return true }
		grid2[x][y] = 0
		if grid1[x][y] == 0{
			ret = false // 不能直接返回，需要把情况全部遍历完
		}
		for _, d := range dirs{
			r, c := x+d[0], y+d[1]
			if isValid(r, c) && grid2[r][c] == 1{
				if !dfs(r,c ){
					ret = false
				}
			}
		}
		return ret
	}
	for i := range grid2{
		for j := range grid2[i]{
			if grid2[i][j] == 1{
				if dfs(i, j){
					ans++
				}
			}
		}
	}
	return ans
}

func countSubIslandsBFS(grid1 [][]int, grid2 [][]int) int {
	m, n := len(grid1), len(grid1[0])
	ans := 0
	dirs := [][]int{ []int{1, 0},[]int{-1, 0},[]int{0, 1},[]int{0, -1} }
	isValid := func(x, y int)bool{
		if x < 0 || y < 0 || x >= m || y >= n{
			return false
		}
		return true
	}
	bfs := func(x, y int)bool{
		q := [][]int{[]int{x, y}}
		grid2[x][y] = 0
		ret := true
		if grid1[x][y] == 0{
			ret = false // 不能立即返回， 需要把此岛屿清空为已访问
		}
		for len(q) > 0{
			t := [][]int{}
			for _, pos := range q{
				for _, d := range dirs{
					r, c := pos[0] + d[0], pos[1] + d[1]
					if isValid(r, c) && grid2[r][c] == 1{
						t = append(t, []int{r,c})
						grid2[r][c] = 0
						if grid1[r][c] == 0{
							ret = false
						}
					}
				}
			}
			q = t
		}
		return ret
	}
	for i := range grid2{
		for j := range grid2[i]{
			if grid2[i][j] == 1{
				if bfs(i, j){
					ans++
				}
			}
		}
	}
	return ans
}
// 方法三 并查集
func countSubIslandsUFS(grid1 [][]int, grid2 [][]int) int {
	m, n := len(grid2), len(grid2[0])
	ufs := make([]int, m*n)
	for i := range ufs{
		ufs[i] = i
	}
	var find func(t int)int
	find = func(t int)int{
		if ufs[t] != t {
			ufs[t] = find(ufs[t])
		}
		return ufs[t]
	}
	union := func(x, y int){
		ufs[find(x)] = find(y)
	}
	dirs := [][]int{[]int{0, -1}, []int{-1, 0}} // 分2个方向，或右下，或左上, 这里选左上
	for i := range grid2{
		for j := range grid2[i]{
			if grid2[i][j] == 0 { continue }
			idx := i * n + j
			for _, d := range dirs{
				r, c := i+d[0], j+d[1]
				if r > 0 && c > 0 {
					rcIdx := r*n+c
					if grid2[r][c] == 1 { union(rcIdx, idx) }
					if grid2[r][c] == 0 { ufs[rcIdx] = -1   }
				}
			}
		}
	}
	set := map[int]bool{}
	for i := range ufs{
		if ufs[i] != -1{
			r, c := i/n, i%n
			if grid1[r][c] == 0{
				set[find(ufs[i])] = false
			}
		}
	}
	ans := 0
	for i := range set{
		if set[i]{ ans++ }
	}
	return ans
}






