package BinarySearch

import "sort"

/* 240. Search a 2D Matrix II
** Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix.
** This matrix has the following properties:
	Integers in each row are sorted in ascending from left to right.
	Integers in each column are sorted in ascending from top to bottom.
 */
// 2022-04-25 刷出此题
// 注意golang search 函数用法
func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix[0])
	for i := range matrix{
		//idx := sort.Search(matrix[i], target)
		idx := sort.Search(n, func(j int)bool{
			return target <= matrix[i][j]
		})
		if idx < n && matrix[i][idx] == target{
			return true
		}
	}
	return false
}
/* 方法二：Z字形查找
** 我们可以从矩阵 matrix 的右上角(0,n-1) 进行搜索.
** 在每一步的搜索过程中，如果我们位于位置(x,y) 那么我们希望在以 matrix 的左下角为左下角、以（x,y) 为右上角的矩阵中进行搜索，即行的范围为
** [x, m-1], 列的范围为[0,y]:
	1. matrix[x][y] = target 搜索完成
	2. matrix[x][y] > target:
		由于每一列的元素都是升序排列的，那么在当前的搜索矩阵中，所有位于第 y 列的元素都是严格大于 target 的
		因此我们可以将它们全部忽略，即将 y - 1
	3. matrix[x][y] < target:
		由于每一行的元素都是升序排列的,那么在当前的搜索矩阵中，所有位于第 x 行的元素都是严格小于 target 的
		因此我们可以将它们全部忽略，即将 x + 1
 */
func searchMatrix_z(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0{ return false }
	n := len(matrix[0])
	if m == 1{
		t := sort.Search(n, func(i int)bool{
			return matrix[m-1][i] >= target
		})
		if t < n && matrix[m-1][t] == target{
			return true
		}
		return false
	}
	if n == 1{
		t := sort.Search(m, func(i int)bool{
			return matrix[i][n-1] >= target
		})
		if t < m && matrix[t][n-1] == target{
			return true
		}
		return false
	}
	for x, y := 0, n-1; x < m && y >= 0;{
		if target == matrix[x][y]{
			return true
		}
		if target > matrix[x][y]{
			x++
		}else{
			y--
		}
	}
	return false
}




















