package Array

import (
	"../../utils"
)

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

/* 48. Rotate Image
** You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
** You have to rotate the image in-place, which means you have to modify the input 2D matrix directly.
** DO NOT allocate another 2D matrix and do the rotation.
 */
// 借助辅助数组很简单，找规律
/* 旋转度	原坐标		转换坐标
	90°		[x,y]	==>	[y, n-1-x]
	180°	[x,y]	==> [n-1-x, n-1-y]
	270°	[x,y]	==>	[n-1-y, x]
 */
/* 原地旋转
** 旋转90度的公式：
	row = col
	col = n-1 - row
不断带入 上面公式 得出递推公式
	[row, col] => [col, n-1-row] => [n-1-row, n-1-col] => [n-1-col, row] => [row, col]
** 当我们知道了如何原地旋转矩阵之后, 还有一个重要的问题在于：我们应该枚举哪些位置[row, col]上进行上述的原地交换操作？
** 由于每一次原地交换 4 个位置，因此：
** 当 n 为偶数时，需要枚举 n^2 / 4 = (n/2)*(n/2) 个位置，保证不重复，不遗漏
** 当 n 为奇数时候，由于中心的位置经过旋转后位置不变， 需要枚举 (n^2-1)/4 = (n-1)/2 * (n+1)/2 个位置
** 需要换另外一种划分方法. 可画图理解 4 ✖️ 4  5 ✖️ 5
 */
func rotate(matrix [][]int)  {
	t, n := 0, len(matrix)
	for i := 0; i < n/2; i++{
		for j := 0; j < (n+1)/2; j++{
			t = matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = t
		}
	}
}
/* 方法三：翻转方法
** 先水平翻转，然后再沿 主对角线翻转
** 对于水平轴翻转而言，我们只需要枚举矩阵上半部分的元素 与 下半部分的元素进行交换
	matrix[row][col] ==+水平轴翻转+==> matrix[n-1-row][col]
** 对于主对角线翻转而言，我们只需要枚举对角线左侧的元素, 和 右侧的元素进行交换
	matrix[row][col] ==+主对角线翻转+==> matrix[col][row]
** 把他们联合就是 matrix[row][col]  ===> matrix[col][n-1-row]
 */
func rotate_flip(matrix [][]int)  {
	n := len(matrix)
	for i := 0; i < n/2; i++{ // 水平翻转
		for j := 0; j < n; j++{
			matrix[i][j], matrix[n-1-i][j] = matrix[n-1-i][j], matrix[i][j]
		}
	}
	for i := 0; i < n; i++{
		for j := 0; j < i; j++{
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

/* 54. Spiral Matrix(螺旋矩阵)
Given an m x n matrix, return all elements of the matrix in spiral order.
*/
/* 按层访问，由外向里层
** 对于每层，从左上方开始以顺时针的顺序遍历所有元素。假设当前层的左上角位于(top,left)，右下角位于 (bottom,right)，按照如下顺序遍历当前层的元素
** 1. 从左到右遍历上侧元素，依次为(top, left) 到 (top, right)
** 2. 从上到下遍历右侧元素，依次为(top+1, right) 到 (bottom, right)
** 3. ....
** 遍历完当前层的元素之后，将left 和 top 分别增加 1，将 right 和 bottom 分别减少 1，进入下一层继续遍历，直到遍历完所有元素为止
 */
// 2021-09-29 刷出此题
func SpiralOrder(matrix [][]int) (ans []int) {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, n-1
	top, bottom := 1, m-1 // 易错点-1 注意 upper 的初始值
	i, j := 0, 0
	cnt := 0
	dir := 0
	for cnt < m*n{
		ans = append(ans, matrix[i][j])
		switch dir {
		case 0:
			if j == right{
				dir = 1
				right--
				i++
			}else{
				j++
			}
		case 1:
			if i == bottom{
				bottom--
				dir = 2
				j--
			}else{
				i++
			}
		case 2:
			if j == left{
				left++
				dir = 3
				i--
			}else{
				j--
			}
		case 3:
			if i == top{
				dir = 0
				top++
				j++
			}else{
				i--
			}
		}
		cnt++
	}
	return ans
}

// 2022-03-11 刷出此题
/* 思路：是按层模拟访问
** 从外向里 递推
** 遗漏点-1： 重复访问的元素， 下一个要加一
** 易漏点-2： 处理只有一排的元素
 */
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	x0, xn := 0, n-1
	y0, yn := 0, m-1
	ans := []int{}
	for x0 <= xn && y0 <= yn{
		for i := x0; i <= xn; i++{
			ans = append(ans, matrix[x0][i])
		}
		//for i := y0; i <= yn; i++{ 遗漏点-1
		for i := y0+1; i <= yn; i++{
			//ans = append(ans, matrix[i][yn])
			ans = append(ans, matrix[i][xn])
		}
		//for i := xn; i >= x0; i--{  遗漏点-1
		for i := xn-1; y0 < yn && i >= x0; i--{ // 遗漏点-2   y0 < yn
			ans = append(ans, matrix[yn][i])
		}
		//for i := yn; i >= 0; i--{ 遗漏点-1
		for i := yn-1; x0 < xn && i > y0; i--{ // // 遗漏点-2  x0 < xn
			ans = append(ans, matrix[i][x0])
		}
		x0, y0, xn, yn = x0+1, y0+1, xn-1, yn-1
	}
	return ans
}

/* 59. Spiral Matrix II
Given a positive integer n, generate an n x n matrix filled with elements from 1 to n2 in spiral order.
*/
/* 给出与54 不同的方式
** 初始位置设为矩阵的左上角，初始方向设为向右，若下一步的位置超出矩阵边界，或者是之前访问过的位置，则顺时针旋转，进入下一个方向。
** 如此反复直至填入n平方个元素
 */
func GenerateMatrix(n int) [][]int {
	var dirs = []struct{x, y int}{{0,1}, {1, 0}, {0, -1}, {-1, 0}}// 右下左上
	matrix := make([][]int, n)
	for i := range matrix{
		matrix[i] = make([]int, n)
	}
	row, col, dirIdx := 0, 0, 0
	for i := 1; i <= n*n; i++{
		matrix[row][col] = i
		dir := dirs[dirIdx]
		// 注意matrix[r][c] > 0  表示是否已经访问过了，如果访问过拐弯
		if r, c := row+dir.x, col+dir.y; r < 0 || r >= n || c >= n || c < 0 || matrix[r][c] > 0{
			dirIdx = (dirIdx+1)%4
			dir = dirs[dirIdx]
		}
		row += dir.x
		col += dir.y
	}
	return matrix
}

/*LCP 29. 乐团站位
** 某乐团的演出场地可视作 num * num 的二维矩阵 grid（左上角坐标为 [0,0])，每个位置站有一位成员。
** 乐团共有 9 种乐器，乐器编号为 1~9，每位成员持有 1 个乐器。
** 为保证声乐混合效果，成员站位规则为：自 grid 左上角开始顺时针螺旋形向内循环以 1，2，...，9 循环重复排列。
** 请返回位于场地坐标 [Xpos,Ypos] 的成员所持乐器编号
 */
// 超时
func orchestraLayout(num int, xPos int, yPos int) int {
	dir := 0
	i, j := 0, 0
	ans := 0
	// 设置边界
	left, right, upper, down := 0, num-1, 0, num-1
	upper = 1
	for i != xPos || j != yPos{
		switch dir {
		case 0:
			if j != right{
				j++
			}else{
				dir = 1
				i++
				right--
			}
		case 1:
			if i != down{
				i++
			}else{
				dir = 2
				j--
				down--
			}
		case 2:
			if j != left{
				j--
			}else{
				dir = 3
				i--
				left++
			}
		case 3:
			if i != upper{
				i--
			}else{
				dir = 0
				upper++
				j++
			}
		}
		ans++
	}
	//return (ans+1) % 9
	return ans % 9 + 1 // 乐队是从1开始计的
}
/* 通过公式计算 思路
** 参考： https://leetcode-cn.com/problems/SNJvJP/solution/shu-xue-yi-ge-gong-shi-ji-ke-by-ivon_shi-mo6a/
** 解决公式比较优雅
** 1. 计算出不包含此位置的 外层 层数 并统计方块个数
** 2. 在当前圈层中，统计方块个数
** 3. 1 和 2 中就地取mod
**
** 1. 求所在的层数k ==> k=min(x,n-1-x,y,n-1-y)
** 记：k(0,1,2...⌈n/2⌉ (往上取整))， C(k)=第k层的方块数目,如C(0)=16,C(1)=8,C(2)=1;
** T(k)=第k层以外（不包括第k层的方块数目）T(0)=0,T(1)=C(0)=16,T(2)=C(0)+C(1)=16+8=24,T(3)=16+8+1=25;
** 补集的思想计算T(k) => T(k)=n*n-(n-2k)*(n-2k)=4*k*(n-k)
**
** 2.求所求点(x,y)相对该层左上角点的相对路径长度-dl
** 画出对角线
** 2.1 对角线以上(包含对角线) x <= y
	dl = (x-k)+1 + (y-k)+1 - 1 = (x-k) + (y-k) + 1
 	绝对路径 = T(k) + dl
** 2.2 对角线以下  x > y
	如果直接类似2.1计算会导致(i,j) 与 (j,i)相对路径值dl相同
	采用补集方式计算，即 在计算层数k的时候，多计算一层，然后再从下一层的入口处回退 dl 个路径
	dl = (x - k) + (y - k) - 1
	绝对路径 = T(k+1) - dl
*/
// go的取模运算不同于python，而是和c++相同，如果是正数就正常操作,
// 如果是负数的取模运算，则需要特别注意，必须(k%n+n)%n  额外加 n ！！！
func OrchestraLayout(num int, xPos int, yPos int) int {
	mod := 9
	ans := 0
	k := utils.Min(xPos, yPos, num-1-xPos, num-1-yPos)
	dl := 0
	if xPos > yPos{
		ans = ((num * num)%mod - ((num-2*(k+1))*(num-2*(k+1)))%mod + mod)%mod
		dl = ((xPos-k)%mod + ((yPos -k) - 1)%mod)%mod
		if ans == 0{ ans = mod }
		if dl == 0{ dl = mod }
		ans = (ans - dl + mod) % mod
		if ans == 0{ return mod }
		return ans%mod
	}
	ans = ((num * num)%mod  - ((num-2*k)*(num-2*k))%mod + mod)%mod  // 补集 计算 外圈 方块个数
	dl = ((xPos-k) + (yPos -k) + 1)%mod
	if ans == 0{ ans = mod }
	if dl == 0{ dl = mod }
	ans = (ans + dl) % mod
	if ans == 0{ return mod }
	return ans%mod
}

/* 885. Spiral Matrix III
** You start at the cell (rStart, cStart) of an rows x cols grid facing east.
** The northwest corner is at the first row and column in the grid, and the southeast corner is at the last row and column.
** You will walk in a clockwise spiral shape to visit every position in this grid.
** Whenever you move outside the grid's boundary,
** we continue our walk outside the grid (but may return to the grid boundary later.).
** Eventually, we reach all rows * cols spaces of the grid.
** Return an array of coordinates representing the positions of the grid in the order you visited them.
 */
/* 螺旋形行走: 找规律
** 检查我们在每个方向的行走长度，我们发现如下模式：1, 1, 2, 2, 3, 3, 4, 4, ....
** 即我们先向东走 1 单位，然后向南走 1 单位，再向西走 2 单位，再向北走 2 单位，再向东走 3 单位，等等
** 按照我们访问的顺序执行遍历并记录网格的位置
 */
func SpiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	dirs := [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
	valid := func(x, y int)bool{
		if x < 0 || y < 0 || x >= rows || y >= cols{
			return false
		}
		return true
	}
	ans := [][]int{[]int{rStart, cStart}}
	n := rows * cols
	if n == 1{ return ans }
	// 1, 1, 2, 2 规律
	for k := 1; k < 2*(rows+cols); k += 2{
		for i, d := range dirs{
			dk := k + (i/2) // 此方向的步数
			for j := 0; j < dk; j++{
				rStart += d[0]
				cStart += d[1]
				if valid(rStart, cStart){
					ans = append(ans, []int{rStart, cStart})
				}
			}
		}
	}
	return ans
}

