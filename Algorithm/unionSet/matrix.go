package unionSet

/* 200. Number of Islands
** Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water),
** return the number of islands.
** An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
** You may assume all four edges of the grid are all surrounded by water.
 */
/*
知识点：1. golang 中 数组是可以比较的 2. 数组可以作为map的key
*/
func NumIslands(grid [][]byte) int {
	index2map := map[[2]int][2]int{}
	for i, row := range grid{
		for j,elem := range row{
			if elem == '1'{
				loc := [2]int{i,j}
				index2map[loc] = loc
			}
		}
	}
	var find func(m [2]int) [2]int
	find = func(m [2]int) [2]int {
		if m != index2map[m]{
			return find(index2map[m])
		}
		return m
	}
	var union func(x, y [2]int)
	union = func(x, y [2]int){
		px, py := find(x), find(y)
		// if px != px 错误点-1： 低级错误
		if px != py {
			//index2map[x] = py  错误点-2： 是类 低级错误
			index2map[px] = py
		}
	}
	rowLen := len(grid)
	for i := 0; i < rowLen; i++ {
		colLen := len(grid[i])
		for j := 0; j < colLen; j++{
			if j+1 < colLen && grid[i][j] == '1' && grid[i][j+1] == '1'{
				union([2]int{i,j}, [2]int{i, j+1})
			}
			if i+1 < rowLen && grid[i][j] == '1' && grid[i+1][j] == '1'{
				union([2]int{i,j}, [2]int{i+1, j})
			}
		}
	}
	islandSet := map[[2]int]bool{}
	//for _,elem := range index2map{ 统计的是key， value可能重复的 低级错误
	for elem,_ := range index2map{
		island := find(elem)
		if !islandSet[island]{
			islandSet[island] = true
		}
	}
	return len(islandSet)
}

// 利用按秩压缩
type UnionFind struct {
	count		int
	parent		[]int
	rank		[]int
}

func ConstructUnionFindByNumIslands(grid [][]byte)*UnionFind{
	// 初始化
	m, n, cnt := len(grid),len(grid[0]), 0
	rank, parent := make([]int, m*n), make([]int, m*n)
	for i := 0; i < m; i++{
		for j := 0; j < n; j++{
			if grid[i][j] == '1'{
				parent[i * n + j] = i * n + j // 2. 类
				cnt++ // 1. 计数
			}
			rank[i * n + j] = 0 // 3. 秩
		}
	}
	return &UnionFind{count: cnt, parent: parent, rank: rank}

}

func (ufs *UnionFind) find(i int)int{
	if ufs.parent[i] != i {
		ufs.parent[i] = ufs.find(ufs.parent[i]) // 1. 路径压缩
	}
	return ufs.parent[i]
}

func(ufs *UnionFind) union(x, y int){
	rootx, rooty := ufs.find(x), ufs.find(y)
	if rootx != rooty{ // 4. 按秩压缩： 查询高度小的 插入到 查询高度大的
		if ufs.rank[rootx] > ufs.rank[rooty]{
			ufs.parent[rooty] = rootx
		}else if ufs.rank[rootx] < ufs.rank[rooty]{
			ufs.parent[rootx] = rooty
		}else{
			ufs.parent[rooty] = rootx
			ufs.rank[rootx] += 1 // 2. 按秩压缩，查询高度记录
		}
		ufs.count-- // 3. union 后 计数递减
	}
}

func NumIslandsUFSImprove(grid [][]byte) int {
	if len(grid) <= 0{
		return 0
	}
	nr, nc := len(grid), len(grid[0])
	pUfs := ConstructUnionFindByNumIslands(grid)
	for r := 0; r < nr; r++{
		for c := 0; c < nc; c++{
			if grid[r][c] == '1' {
				grid[r][c] = '0'
				cur := r * nc + c
				if r - 1 >= 0 && grid[r-1][c] == '1'{
					pUfs.union(cur, (r-1) * nc + c)
				}
				if r + 1 < nr && grid[r+1][c] == '1'{
					pUfs.union(cur, (r+1) * nc + c)
				}
				if c - 1 >= 0 && grid[r][c-1] == '1'{
					pUfs.union(cur, r * nc + c - 1)
				}
				if c + 1 < nc && grid[r][c+1] == '1'{
					pUfs.union(cur, r * nc + c + 1)
				}
			}
		}
	}
	return pUfs.count
}

// 2022-02-24 重新刷出此题，并查集解决方法
// 按秩压缩思路
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	ufs := make([][]int, m*n) // 记录高度，按秩压缩
	for i := range ufs{
		ufs[i] = make([]int, 2)
	}
	ans := 0
	for i := range ufs{
		ufs[i][0], ufs[i][1] = i, 0
		if grid[i/n][i%n] == '1'{ // 记录个数
			ans++
		}else{
			ufs[i][0] = -1
		}
	}
	dirs := [][]int{ []int{1, 0},[]int{-1, 0},[]int{0, 1},[]int{0, -1} }
	isValid := func(x, y int)bool{
		if x < 0 || y < 0 || x >= m || y >= n{
			return false
		}
		return true
	}
	var find func(src int)int
	find = func(src int)int{
		if src != ufs[src][0]{
			ufs[src][0] = find(ufs[src][0])
		}
		return ufs[src][0]
	}
	union := func(src, dst int){
		p_src, p_dst := find(src), find(dst)
		if p_src != p_dst{
			//ufs[src] = ufs[dst] 错误 ❌
			// ufs[p_src] = p_dst // 更改为按秩压缩
			// 按秩压缩： 查询高度小的 插入到 查询高度大的
			if ufs[p_src][1] > ufs[p_dst][1]{
				ufs[p_dst][0] = p_src
			}else if ufs[p_src][1] < ufs[p_dst][1]{
				ufs[p_src][0] = p_dst
			}else{
				ufs[p_src][0] = p_dst
				ufs[p_dst][1]++
			}
			ans--

		}
	}
	for i := range grid{
		for j := range grid[i] {
			if grid[i][j] == '1'{
				src := i*n+j
				grid[i][j] = '0'
				for _, d := range dirs{
					x, y := i+d[0], j+d[1]
					if isValid(x, y) && grid[x][y] == '1'{
						dst := x*n+y
						union(src, dst)
					}
				}
			}
		}
	}
	//fmt.Println(ufs)
	return ans
}

/* 695. Max Area of Island
** You are given an m x n binary matrix grid.
** An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.)
** You may assume all four edges of the grid are surrounded by water.
** The area of an island is the number of cells with a value 1 in the island.
** Return the maximum area of an island in grid. If there is no island, return 0.
 */
// 2022-02-25 使用并查集刷出此题，这里只是记录对并查集秩的理解
func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dirs := [][]int{ []int{1, 0},[]int{-1, 0},[]int{0, 1},[]int{0, -1} }
	isValid := func(x, y int)bool{
		if x < 0 || y < 0 || x >= m || y >= n{
			return false
		}
		return true
	}
	ufs := make([][]int, m*n) // [][0]:表示所属组  [][0]:表示秩，在这里是统计个数
	for i := range ufs{
		ufs[i] = []int{i, 1}
	}
	find := func(node int)int{
		for node != ufs[node][0]{
			node = ufs[node][0]
		}
		return ufs[node][0]
	}
	union := func(x, y int){ // 增加按秩合并
		px, py := find(x), find(y)
		if px != py {
			ufs[px][0] = py
			//ufs[px][1] += ufs[py][1] // py 是主类
			ufs[py][1] += ufs[px][1] // 按秩合并，最后统计的时候 会看到 秩的 含义
		}
	}
	for i := range grid{
		for j := range grid[i]{
			src := i*n+j
			if grid[i][j] == 1{
				grid[i][j] = 0 // 这里visited， 避免回来
				for _, d := range dirs{
					x, y := i+d[0], j+d[1]
					if isValid(x, y) && grid[x][y] == 1{
						dst := x*n+y
						union(src, dst)
					}
				}
			}else{
				ufs[src][0] = -1
			}
		}
	}
	ans := 0
	/*
	   groups := map[int]int{}
	   for i := range ufs{
	       if ufs[i] == -1{ continue }
	       groups[find(i)]++
	   }
	   for i := range groups{
	       if ans < groups[i]{
	           ans = groups[i]
	       }
	   }
	*/
	for i := range ufs{
		if ufs[i][0] == -1{ continue }
		if ufs[i][0] == i{ // 最终的类别
			if ans < ufs[i][1]{
				ans = ufs[i][1]
			}
		}
	}
	return ans
}