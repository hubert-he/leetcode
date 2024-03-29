package array

import "../utils"
import "math"



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
// 2021-10-25的处理方法：模拟方式
// 模拟方式在于困难点在于 小矩阵的处理，这个在矩阵分割有借鉴
func IsValidSudoku1(board [][]byte) bool {
	for i := range board{ // 处理行
		cell := make([]bool, 9)
		for _, c := range board[i]{
			if c == '.'{
				continue
			}
			if cell[c-'1']{
				return false
			}
			cell[c-'1'] = true
		}
	}
	for i := range board{// 处理列
		cell := make([]bool, 9)
		for j := 0; j < 9; j++{
			c := board[j][i]
			if c == '.'{
				continue
			}
			if cell[c-'1']{
				return false
			}
			cell[c-'1'] = true
		}
	}
	for i := 0; i < 3; i++{ // 处理小矩阵
		for j := 0; j < 3; j++{
			cell := make([]bool, 9)
			for p := 3*i; p < 3*i+3; p++{
				for q := 3*j; q < 3*j+3; q++{
					c := board[p][q]
					if c != '.' && cell[c-'1']{
						return false
					}
					if c != '.'{
						cell[c-'1'] = true
					}
				}
			}
		}
	}
	return true
}
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

/* 1292. Maximum Side Length of a Square with Sum Less than or Equal to Threshold
** Given a m x n matrix mat and an integer threshold.
** Return the maximum side-length of a square with a sum less than or equal to threshold or return 0 if there is no such square.
** 注意额外增加为0的首行首列，方便处理
	1 <= m, n <= 300
	m == mat.length
	n == mat[i].length
	0 <= mat[i][j] <= 10000
	0 <= threshold <= 10^5
 */
/* 二维数组前缀和：dp[i][j] = mat[i][j] + dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1]
** 设二维数组 A 的大小为 m * n，行下标的范围为 [1, m]，列下标的范围为 [1, n]
** 数组 P 是 A 的前缀和数组，等价于 P 中的每个元素 P[i][j]：
** 1. 如果 i 和 j 均大于 0，那么 P[i][j] 表示 A 中以 (1, 1) 为左上角，(i, j) 为右下角的矩形区域的元素之和；
** 2. 如果 i 和 j 中至少有一个等于 0，那么 P[i][j] 也等于 0。
** 数组 P 可以帮助我们在 O(1) 的时间内求出任意一个矩形区域的元素之和。
** 具体地，设我们需要求和的矩形区域的左上角为 (x1, y1)，右下角为 (x2, y2)，则该矩形区域的元素之和可以表示为：
**  sum = A[x1...x2][y1...y2] = P[x2][y2] - P[x1 - 1][y2] - P[x2][y1 - 1] + P[x1 - 1][y1 - 1]
*/
func maxSideLength(mat [][]int, threshold int) int {
	m, n := len(mat), len(mat[0])
	// 前缀和 第一元素为 0，引入前缀 0 方便条件处理
	dp := make([][]int, m+1)
	for i := range dp{
		dp[i] = make([]int, n+1)
	}
	/* 此处理解错误，二维前缀和，在边界位置 其实为一维前缀和
	// 初始化
	copy(dp[0], mat[0])
	for i := 0; i < m; i++{
		dp[i][0] = mat[i][0]
	}
	// 前缀和 第一元素为 0，引入前缀 0 方便条件处理
	for i := 0; i < m; i++{
		for j := 0; j < n; j++{
			dp[i][j] = mat[i][j]
			if i > 0{
				dp[i][j] += dp[i-1][j]
			}
			if j > 0{
				dp[i][j] += dp[i][j-1]
			}
			if i > 0 && j > 0{
				dp[i][j] -= dp[i-1][j-1]
			}
		}
	}
	*/
	for i := 1; i <= m; i++{
		for j := 1; j <= n; j++{
			dp[i][j] = mat[i-1][j-1] + dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1]
		}
	}
	ans := 0
	/* 题意是正方形区域，并且是任意位置开始的
	for i := 1; i < m; i++{
		for j := 1; j < n; j++{
			if dp[i][j] <= threshold{
				ans = max(ans, i, j)
			}
		}
	}
	*/
	// 区域和(右下角坐标i，j， 长宽为 k) = dp[i][j] - dp[i-k][j] - dp[i][j-k] + dp[i-k][j-k]
	for k := 1; k <= m && k <= n; k++{
		for i := m; i > 0; i--{
			for j := n; j > 0; j--{
				if i - k < 0 || j - k < 0{
					continue
				}
				t := dp[i][j] - dp[i-k][j] - dp[i][j-k] + dp[i-k][j-k]
				if t <= threshold{
					if ans < k {
						ans = k
					}
				}
			}
		}
	}
	return ans
}
/*二分： 注意题目限制条件： 全是正数或0
** 查找的正方形的边长越长，其计算出来的元素总和越大。
** 我们可以二分正方形的边长，在满足阈值条件下尽可能地扩大正方形的边长，其等价于在升序数组中查找一个小于等于 k 的最大元素。
** 二分的具体思路：
	控制 l 到 h 都是可能可能的值
	如果 mid 满足阈值条件，则 l = mid，l 可能是答案，不能直接舍去。
	如果 mid 不满足阈值条件，则 h = mid - 1。
	当 l = h 或 l + 1 = h 时跳出循环（小提示：l = mid 可能造成死循环，通过 l + 1 == h 条件跳出），判断 l 和 h 两个是最优解。
*/
func maxSideLength2(mat [][]int, threshold int) int {
	m, n := len(mat), len(mat[0])
	sn := m
	if sn > n{
		sn = n
	}
	// 前缀和 第一元素为 0，引入前缀 0 方便条件处理
	dp := make([][]int, m+1)
	for i := range dp{
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++{
		for j := 1; j <= n; j++{
			dp[i][j] = mat[i-1][j-1] + dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1]
		}
	}
	ans := 0
	// 二分查找
	start, end := 0, sn
	for start <= end{
		mid := int(uint(start+end)>>1)
		target := 0
		for i := m; i > 0; i--{
			for j := n; j > 0; j--{
				if i - mid < 0 || j - mid < 0{
					continue
				}
				t := dp[i][j] - dp[i-mid][j] - dp[i][j-mid] + dp[i-mid][j-mid]
				if t <= threshold{
					target = t
					if ans < mid{
						ans = mid
					}
				}
			}
		}
		if target == 0{
			end = mid-1
		}else{
			start = mid +1
		}
	}
	return ans
}
// 优化3：回到暴力枚举的方法上，枚举算法中包括三重循环，其中前两重循环枚举正方形的左上角位置，而第三重循环枚举的是正方形的边长
// 优化思路-1： 如果边长为 c 的正方形的元素之和已经超过阈值，那么我们就没有必要枚举更大的边长了。
//		因为数组 mat 中的所有元素均为非负整数，如果固定了左上角的位置 (i, j)（即前两重循环），那么随着边长的增大，正方形的元素之和也会增大
// 优化思路-2： 由于我们的目标是找到边长最大的正方形，那么如果我们在前两重循环枚举到 (i, j) 之前已经找到了一个边长为 c' 的正方形，
//		那么在枚举以 (i, j) 为左上角的正方形时，我们可以忽略所有边长小于等于 c' 的正方形，直接从 c' + 1 开始枚举
func maxSideLength3(mat [][]int, threshold int) int {
	ans := 0
	m, n := len(mat), len(mat[0])
	sn := m
	if sn > n{
		sn = n
	}
	prefixSum := make([][]int, m+1)
	for i := range prefixSum{
		prefixSum[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++{
		for j := 1; j <= n; j++{
			prefixSum[i][j] = mat[i-1][j-1] + prefixSum[i-1][j] + prefixSum[i][j-1] - prefixSum[i-1][j-1]
		}
	}
	sum := func(i, j, c int)int{
		return prefixSum[i+c][j+c] - prefixSum[i+c][j] - prefixSum[i][j+c] + prefixSum[i][j]
	}
	for i := 0; i < m; i++{
		for j := 0; j < n; j++{
			for c := ans + 1; c <= sn; c++{//优化思路-2
				if i+c <= m && j+c <= n && sum(i,j,c) <= threshold{
					ans++
				}else{ // 优化思路-1
					break
				}
			}
		}
	}
	return ans
}
/* 与二维前缀和 相似的题目 可作为训练使用
** 1314. Matrix Block Sum
** Given a m x n matrix mat and an integer k, return a matrix answer where each answer[i][j] is the sum of all elements mat[r][c] for:
	i - k <= r <= i + k,
	j - k <= c <= j + k, and
	(r, c) is a valid position in the matrix.
Example 1:
	Input: mat = [[1,2,3],[4,5,6],[7,8,9]], k = 1
	Output: [[12,21,16],[27,45,33],[24,39,28]]
Example 2:
	Input: mat = [[1,2,3],[4,5,6],[7,8,9]], k = 2
	Output: [[45,45,45],[45,45,45],[45,45,45]]
*/
func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	dp := make([][]int, m+1)
	for i := range dp{
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++{
		for j := 1; j <= n; j++{
			dp[i][j] = mat[i-1][j-1] + dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1]
		}
	}
	for i := range ans {
		for j := range ans[i]{
			r1, c1 := i-k, j-k
			r2, c2 := i+k, j+k
			if r1 < 0{
				r1 = 0
			}
			if c1 < 0 {
				c1 = 0
			}
			if r2 >= m{
				r2 = m-1
			}
			if c2 >= n {
				c2 = n-1
			}
			// 留意 索引值 不同
			ans[i][j] = dp[r2+1][c2+1] - dp[r2+1][c1] - dp[r1][c2+1] + dp[r1][c1]
		}
	}
	return ans
}





























