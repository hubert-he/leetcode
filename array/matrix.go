package array

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