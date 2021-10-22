package array

import "../utils"
import "math"

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

/* dfs 如果用 dfs[i][j] = min(四个方向dfs） 会发生递归死循环
** 从每一个周围有0的1开始dfs， 破除dfs 死循环
*/
func UpdateMatrixDFS(matrix [][]int)[][]int{
	m, n := len(matrix), 0
	if m == 0{
		return matrix
	}
	n = len(matrix[0])
	dir := [][2]int{[2]int{0,1}, [2]int{0,-1}, [2]int{1,0}, [2]int{-1,0}}
	valid := func(x, y int)bool{
		if x >= 0 && y >= 0 && x < m && y < n{
			return true
		}
		return false
	}
	nearZero := func(x, y int)bool{
		for i := range dir{
			dx, dy := x + dir[i][0], y + dir[i][1]
			if valid(dx,dy) && matrix[dx][dy] == 0{
				return true
			}
		}
		return false
	}
	var dfs func(x, y int)
	dfs = func(x, y int){
		for i := range dir{
			dx, dy := x + dir[i][0], y + dir[i][1]
			if valid(dx,dy) && matrix[dx][dy] > matrix[x][y]+1{
				matrix[dx][dy] = matrix[x][y]+1
				dfs(dx, dy)
			}
		}
	}
	/* 将周围没有0，且值为1的位置 设置为一个极大值*/
	for i := range matrix{
		for j := range matrix[i]{
			if matrix[i][j] == 1 && !nearZero(i, j){
				matrix[i][j] = math.MaxInt32
			}
		}
	}
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 1{
				dfs(i, j)
			}
		}
	}
	return matrix
}
/*转换为图算法
** 我们可以从 0 的位置开始进行 广度优先搜索。
** 广度优先搜索可以找到从起点到其余所有点的 最短距离，因此如果我们从 0 开始搜索，每次搜索到一个 1，就可以得到 0 到这个 1 的最短距离，
** 也就离这个 1 最近的 0 的距离了（因为矩阵中只有一个 0）
**
** 我们需要对于每一个 1 找到离它最近的 0。如果只有一个 0 的话，我们从这个 0 开始广度优先搜索就可以完成任务了；
** 但在实际的题目中，我们会有不止一个 0。我们会想，要是我们可以把这些 0 看成一个整体好了。
** 有了这样的想法，我们可以添加一个「超级零」，它与矩阵中所有的 0 相连，这样的话，任意一个 1 到它最近的 0 的距离，会等于这个 1 到「超级零」的距离减去一。
** 由于我们只有一个「超级零」，我们就以它为起点进行广度优先搜索。
** 这个「超级零」只和矩阵中的 0 相连，所以在广度优先搜索的第一步中，「超级零」会被弹出队列，而所有的 0 会被加入队列，它们到「超级零」的距离为 1。
** 这就等价于：一开始我们就将所有的 0 加入队列，它们的初始距离为 0。
** 这样以来，在广度优先搜索的过程中，我们每遇到一个 1，就得到了它到「超级零」的距离减去一，也就是 这个 1 到最近的 0 的距离
** 一些图的原理：
** 熟悉「最短路」的读者应该知道，我们所说的「超级零」实际上就是一个「超级源点」。
** 在最短路问题中，如果我们要求多个源点出发的最短路时，一般我们都会建立一个「超级源点」连向所有的源点，
** 用「超级源点」到终点的最短路等价多个源点到终点的最短路
** 这类题型属于 多源BFS
** 对于 「Tree 的 BFS」 （典型的「单源 BFS」）
** 对于 「图 的 BFS」 （「多源 BFS」）
** Tree 只有 1 个 root，而图可以有多个源点，所以首先需要把多个源点都入队；
** Tree 是有向的因此不需要标识是否访问过，而对于无向图来说，必须得标志是否访问过哦！
** 并且为了防止某个节点多次入队，需要在其入队之前就将其设置成已访问！
*/
func UpdateMatrixBFS(matrix [][]int)[][]int{
	m, n := len(matrix), 0
	if m == 0{
		return matrix
	}
	n = len(matrix[0])
	visited := make([][]bool, m)
	for i := range visited{
		visited[i] = make([]bool, n)
	}
	dir := [][2]int{[2]int{0,1}, [2]int{0,-1}, [2]int{1,0}, [2]int{-1,0}}
	valid := func(x, y int)bool{
		if x >= 0 && y >= 0 && x < m && y < n{
			return true
		}
		return false
	}
	q := [][2]int{}
	// 将所有的 0 添加进队列
	for i := range matrix{
		for j := range matrix[i]{
			if matrix[i][j] == 0{
				q = append(q, [2]int{i, j})
				visited[i][j] = true
			}
		}
	}
	// BFS 查找最短路
	for len(q) > 0{
		for i := range dir{
			dx, dy := q[0][0]+dir[i][0], q[0][1]+dir[i][1]
			if valid(dx, dy) && !visited[dx][dy]{
				matrix[dx][dy] = matrix[q[0][0]][q[0][1]] + 1
				q = append(q, [2]int{dx,dy})
				visited[dx][dy] = true
			}
		}
		q = q[1:]
	}
	return matrix
}

/* 如果 0 在矩阵中的位置是(i0, j0), 1在矩阵中的位置是(i1,j1)
** 那么我们可以直接算出 0 和 1 之间的距离。因为我们从 1 到 0 需要在水平方向走|i0-i1|, 竖直方向走 |j0-j1|
** 得出距离：|i0-i1| + |j0-j1|
** 对于矩阵中的任意一个 1 以及一个 0，我们如何从这个 1 到达 0 并且距离最短呢
** 我们可以从 1 开始，先在水平方向移动，直到与 0 在同一列，随后再在竖直方向上移动，直到到达 0 的位置。
** 这样一来，从一个固定的 1 走到任意一个 0，在距离最短的前提下可能有四种方法：
**	只有 水平向左移动 和 竖直向上移动；
**	只有 水平向左移动 和 竖直向下移动；
**	只有 水平向右移动 和 竖直向上移动；
**	只有 水平向右移动 和 竖直向下移动。
** 例如一个矩阵包含四个 0。从中心位置的 1 移动到这四个 0，就需要使用四种不同的方法
** 这样以来，我们就可以使用动态规划解决这个问题了
** dp[i][j]表示位置(i,j)到最近的0的距离。
** 如果我们只能「水平向左移动」和「竖直向上移动」，那么
** 	  <1> 可以向上移动一步，再移动dp[i-1][j]步到达某个0
**    <2> 可以向左移动一步，再移动dp[i][j-1]步到达某个0
** 推出转移方程：
** dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1]), 	位置(i,j)的元素为1
** dp[i][j] = 0, 								位置(i,j)的元素为0
** 对于另外三种移动方法，我们也可以写出类似的状态转移方程，得到四个 dp[i][j] 的值，那么其中最小的值就表示位置(i,j) 到最近的 0 的距离。
*/
func UpdateMatrixDP(matrix [][]int)[][]int{
	m, n := len(matrix), 0
	if m == 0{
		return matrix
	}
	n = len(matrix[0])
	// 初始化动态规划的数组，所有的距离值都设置为一个很大的数
	ans := make([][]int, m)
	for i := range ans{
		ans[i] = make([]int, n)
		for j := range ans[i]{
			ans[i][j] = math.MaxInt32
			if matrix[i][j] == 0{
				ans[i][j] = 0
			}
		}
	}
	// 只有 水平向左移动 和 竖直向上移动，注意动态规划的计算顺序
	for i := range ans{
		for j := range ans[i]{
			if i > 0{
				ans[i][j] = utils.Min(ans[i][j], ans[i-1][j] + 1)
			}
			if j > 0{
				ans[i][j] = utils.Min(ans[i][j], ans[i][j-1] + 1)
			}
		}
	}
	// 只有 水平向左移动 和 竖直向下移动，注意动态规划的计算顺序
	for i := m-1; i >= 0; i--{ // 注意动态规划的计算顺序
		for j := 0; j < n; j++{
			if i+1 < m{
				ans[i][j] = utils.Min(ans[i][j], ans[i+1][j] + 1)
			}
			if j-1 >= 0{
				ans[i][j] = utils.Min(ans[i][j], ans[i][j-1] + 1)
			}
		}
	}
	// 只有 水平向右移动 和 竖直向上移动，注意动态规划的计算顺序
	for i := 0; i < m; i++ {
		for j := n - 1; j >= 0; j-- {
			if i-1 >= 0{
				ans[i][j] = utils.Min(ans[i][j], ans[i-1][j] + 1)
			}
			if j+1 <= n{
				ans[i][j] = utils.Min(ans[i][j], ans[i][j+1] + 1)
			}
		}
	}
	// 只有 水平向右移动 和 竖直向下移动，注意动态规划的计算顺序
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i + 1 < m{
				ans[i][j] = utils.Min(ans[i][j], ans[i + 1][j] + 1);
			}
			if j + 1 < n {
				ans[i][j] = utils.Min(ans[i][j], ans[i][j + 1] + 1);
			}
		}
	}
	return ans
}
/* 动态规划的常数优化
** 我们发现方法二中的代码有一些重复计算的地方。实际上，我们只需要保留
	只有 水平向左移动 和 竖直向上移动；
	只有 水平向右移动 和 竖直向下移动。
**  对于任一点(i,j),距离0的距离为：<== 这一步是可以想到的
**  dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i+1][j], dp[i][j+1])  if matrix[i][j] == 1
**           = 0 														if matrix[i][j] == 0
**  发现 dp[i][j]dp[i][j] 是由其上下左右四个状态来决定，无法从一个方向开始递推！
** 于是我们尝试将问题分解： trick： 分解问题，再合并解决
** 1. 距离 (i,j) 最近的 0 的位置，是在其 「左上，右上，左下，右下」4个方向之一；
** 2. 因此我们分别从四个角开始递推，就分别得到了位于「左上方、右上方、左下方、右下方」距离(i,j)的最近的 0 的距离，取 min 即可；
** 3. 从四个角开始的 4 次递推，其实还可以优化成从任一组对角开始的 2 次递推；原因是：
** 	首先从左上角开始递推 dp[i][j] 是由其 「左方」和 「左上方」的最优子状态决定的；
**  然后从右下角开始递推 dp[i][j] 是由其 「右方」和 「右下方」的最优子状态决定的；
**  看起来第一次递推把「右上方」的最优子状态给漏掉了，其实不是的，因为第二次递推的时候「右方」的状态在第一次递推时已经包含了「右上方」的最优子状态了；
**  看起来第二次递推把「左下方」的最优子状态给漏掉了，其实不是的，因为第二次递推的时候「右下方」的状态在第一次递推时已经包含了「左下方」的最优子状态了。
*/
func UpdateMatrixDP2(matrix [][]int)[][]int{
	m, n := len(matrix), 0
	if m == 0{
		return matrix
	}
	n = len(matrix[0])
	// 初始化动态规划的数组，所有的距离值都设置为一个很大的数
	ans := make([][]int, m)
	for i := range ans{
		ans[i] = make([]int, n)
		for j := range ans[i]{
			ans[i][j] = math.MaxInt32
			if matrix[i][j] == 0{
				ans[i][j] = 0
			}
		}
	}
	// 只有 水平向左移动 和 竖直向上移动，注意动态规划的计算顺序
	for i := range ans{
		for j := range ans[i]{
			if i > 0{
				ans[i][j] = utils.Min(ans[i][j], ans[i-1][j] + 1)
			}
			if j > 0{
				ans[i][j] = utils.Min(ans[i][j], ans[i][j-1] + 1)
			}
		}
	}
	// 只有 水平向右移动 和 竖直向下移动，注意动态规划的计算顺序
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i + 1 < m{
				ans[i][j] = utils.Min(ans[i][j], ans[i + 1][j] + 1);
			}
			if j + 1 < n {
				ans[i][j] = utils.Min(ans[i][j], ans[i][j + 1] + 1);
			}
		}
	}
	return ans
}

/* 73. Set Matrix Zeroes
** Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's, and return the matrix.
You must do it in place.
 */
/* 方法一：使用标记数组，两个标记数组分别记录每一行和每一列是否有零出现
** 方法二：使用两个标记变量
**   	可以用矩阵的第一行和第一列代替方法一中的两个标记数组，以达到 O(1) 的额外空间。
**		但这样会导致原数组的第一行和第一列被修改，无法记录它们是否原本包含 0。
**		因此我们需要额外使用两个标记变量分别记录第一行和第一列是否原本包含 0
** 方法三：使用一个标记变量
**		只使用一个标记变量记录第一列是否原本存在 0
**		第一列的第一个元素即可以标记第一行是否出现 0，
**		为了防止每一列的第一个元素被提前更新，我们需要从最后一行开始，倒序地处理矩阵元素
*/
func SetZeroesTwoFlag(matrix [][]int){
	m, n := len(matrix), len(matrix[0])
	r0, c0 := false, false
	for i := 0; i < n; i++{
		if matrix[0][i] == 0{
			r0 = true
			break
		}
	}
	for i := 0; i < m; i++{
		if matrix[i][0] == 0{
			c0 = true
			break
		}
	}
	for i := 1; i < m; i++{
		for j := 1; j < n; j++{
			if matrix[i][j] == 0{
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	// 注意 覆写 第0行 与 第 0列的情况
	//for i := 0;  i < m; i++{
	for i := 1;  i < m; i++{
		if matrix[i][0] == 0{
			for j := 0; j < n; j++{
				matrix[i][j] = 0
			}
		}
	}
	// 注意 覆写 第0行 与 第 0列的情况
	//for i := 0;  i < n; i++{
	for i := 1;  i < n; i++{
		if matrix[0][i] == 0{
			for j := 0; j < m; j++{
				matrix[j][i] = 0
			}
		}
	}
	if r0 {
		for i := 0; i < n; i++{
			matrix[0][i] = 0
		}
	}
	if c0 {
		for i := 0; i < m; i++{
			matrix[i][0] = 0
		}
	}
}
// 方法三的 最优
func SetZeroes(matrix [][]int){
	m, n := len(matrix), len(matrix[0])
	c0 := false
	for i := range matrix{
		// 标记第0列是否含有原始0
		if matrix[i][0] == 0{
			c0 = true
		}
		for j := 1; j < n; j++{
			if matrix[i][j] == 0{
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	// 防止覆写第0行，逆序执行，最后执行第0行
	for i := m-1; i >= 0; i--{
		for j := 1; j < n; j++{
			if matrix[i][0] == 0 || matrix[0][j] == 0{
				matrix[i][j]= 0
			}
		}
		if c0{
			matrix[i][0] = 0
		}
	}
}

/*36. Valid Sudoku
Determine if a 9 x 9 Sudoku board is valid. 
Only the filled cells need to be validated according to the following rules:
Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
Note:
	A Sudoku board (partially filled) could be valid but is not necessarily solvable.
	Only the filled cells need to be validated according to the mentioned rules.
*/
func IsValidSudoku(board [][]byte) bool {
	var row, col [9][9]int
	var subboxes [3][3][9]int
	for i := range board{
		for j, c := range board[i]{
			if c == '.'{
				continue
			}
			index := c - '1'
			row[i][index]++
			col[j][index]++
			subboxes[i/3][j/3][index]++
			if row[i][index] > 1 || col[j][index] > 1 || subboxes[i/3][j/3][index] > 1{
				return false
			}
		}
	}
	return true
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
func SpiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {

}

