package backtracking

import (
	"fmt"
	"math/bits"
)

/* 37. Sudoku Solver
** Write a program to solve a Sudoku puzzle by filling the empty cells.
** A sudoku solution must satisfy all of the following rules:
	Each of the digits 1-9 must occur exactly once in each row.
	Each of the digits 1-9 must occur exactly once in each column.
	Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
** The '.' character indicates empty cells.
** Constraints:
	board.length == 9
	board[i].length == 9
	board[i][j] is a digit or '.'.
	It is guaranteed that the input board has only one solution.
 */
// 2022-01-19 刷出此题
func SolveSudoku(board [][]byte) {
	m, n := len(board), len(board[0])
	check := func(row, col, value int)bool{
		for i := 0; i < m; i++{
			if i != row && board[i][col] == byte(value)+'0'{
				return false
			}
		}
		for i := 0; i < n; i++{
			if i != col && board[row][i] == byte(value)+'0'{
				return false
			}
		}
		for i := row/3*3; i < row/3*3+3; i++{
			for j := col/3*3; j < col/3*3+3; j++{
				if i != row && j != col && board[i][j] == byte(value)+'0'{
					return false
				}
			}
		}
		fmt.Println(row, col, board[row][col]-'0', value)
		return true
	}
	var dfs func(x, y int)bool
	dfs = func(x, y int)bool{
		result := false
		var i, j int
		end := true
		for i = x; i < m; i++{
			if i == x{ j = y }else{ j = 0 }
			for ;j < n; j++{
				if board[i][j] == '.'{
					end = false
					for u := 1; u < 10; u++{
						board[i][j] = byte(u) + '0'
						if check(i, j, u){
							result = dfs(i, j)
							if result {
								return true
							}
						}
					}
					board[i][j] = '.'
					return result
				}
			}
		}
		if end {
			fmt.Println("xx")
			return true
		}
		fmt.Println(result)
		return result
	}
	dfs(0, 0)
}

func SolveSudoku_DFS(board [][]byte) {
	m, n := len(board), len(board[0])
	check := func(row, col, value int)bool{
		for i := 0; i < m; i++{
			if i != row && board[i][col] == byte(value)+'0'{
				return false
			}
		}
		for i := 0; i < n; i++{
			if i != col && board[row][i] == byte(value)+'0'{
				return false
			}
		}
		for i := row/3*3; i < row/3*3+3; i++{
			for j := col/3*3; j < col/3*3+3; j++{
				if i != row && j != col && board[i][j] == byte(value)+'0'{
					return false
				}
			}
		}
		fmt.Println(row, col, board[row][col]-'0', value)
		return true
	}
	var dfs func(x, y int)bool
	dfs = func(x, y int)bool{
		var i, j int
		for i = x; i < m; i++{
			if i == x{ j = y }else{ j = 0 }
			for ;j < n; j++{
				if board[i][j] == '.'{
					for u := 1; u < 10; u++{
						board[i][j] = byte(u) + '0'
						if check(i, j, u){
							if dfs(i, j){
								return true
							}
						}
					}
					board[i][j] = '.'
					return false
				}
			}
		}
		return true
	}
	dfs(0, 0)
}

// 合并成线性 有助于处理
func solveSudoku_DFS(board [][]byte)  {
	var line, column [9][9]bool
	var block [3][3][9]bool
	var spaces [][2]int  // 将所有空格 合并成 线性访问
	for i := range board{
		for j, c := range board[i]{
			if c == '.'{
				spaces = append(spaces, [2]int{i, j})
			}else{
				digit := c - '1'
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
			}
		}
	}
	var dfs func(int) bool
	dfs = func(pos int) bool{
		if pos == len(spaces){
			return true
		}
		i, j := spaces[pos][0], spaces[pos][1]
		for digit := byte(0); digit < 9; digit++{
			if !line[i][digit] && !column[j][digit] && !block[i/3][j/3][digit]{
				line[i][digit], column[j][digit], block[i/3][j/3][digit] = true, true, true
				board[i][j] = digit + '1'
				if dfs(pos + 1){
					return true
				}
				line[i][digit], column[j][digit], block[i/3][j/3][digit] = false, false, false
			}
		}
		return false
	}
	dfs(0)
}
/* 借助位运算，仅使用一个整数表示每个数字是否出现过
** 数 b 的二进制表示的第 i 位（从低到高，最低位为第 0 位）为 1，当且仅当数字 i+1 已经出现过。
** 例如当 b 的二进制表示为 (011000100) 时，就表示数字 3，7，8 已经出现过。
** 1. 对于第 i 行第 j 列的位置, line[i] | column[j] | block[i/3][j/3]中第 k 位为 1，表示该位置不能填入 数字 k + 1
	  如果我们对这个值进行 ~ 按位取反运算, 那么第k位为1 就表示该位置可以填入数字k+1， 我们就可以通过寻找 1 来进行枚举。
      由于在进行按位取反运算后，这个数的高位也全部变成了1，而这是我们不应当枚举到的，
	  因此还需要将这个数和 0x1FF 进行按位与运算 &，即将所有无关的位置为 0.
** 2. 使用按位异或运算，将第 i 位 从 0 变成 1，或从 1 变为 0.
** 3. b & (-b) 可得到 二进制中最低位的 1。这是因为 (−b) 在计算机中以补码的形式存储，它等于 ~b+1 取反加一
	  由于 b & ~b = 0，若把 ~b 增加 1 之后，最低位的连续都变成0，而最低位的0变成1，对应到b中即为最低位的 1
	  结论：当 b 和 ~b+1 进行按位与运算时， 只有最低位的 1 会被保留
** 4. 当我们得到这个最低位的 1 时， 可以通过golang 的bits 库函数得到这个最低位1 究竟是第几位(即i值)
** 5. 我们可以用 b 和 最低位的 1 进行按位 异或 运算，就可以将其从 b 中消除，这样就可以枚举下一个 1
	  同样 也可以 通过  b & (b-1) 来把 b 中 最低位的1消除
 */
func solveSudoku_bitset(board [][]byte)  {
	var line, column [9]int
	var block [3][3]int
	// 在flip 中为何 用的是异或 而不是 或 运算： 这是因为在下面dfs中 flip 担任2个功能：把数字置为可用 和 不可用
	// 如果只用或， 则把数字恢复为可用的时候，需要再实现一个函数。
	// flip 对同一个digit flip 一次 标记为已用， 再次对相同的digit 进行flip，则清除已用标记
	flip := func(x, y int, digit byte){ // 利用bitset 进行状态压缩
		line[x] ^= 1 << digit
		column[y] ^= 1 << digit
		block[x/3][y/3] ^= 1 << digit
	}
	var spaces [][2]int
	for i := range board{
		for j, c := range board[i]{
			if c == '.'{
				spaces = append(spaces, [2]int{i, j})
			}else{// bitset 记录
				digit := c - '1'
				flip(i, j, digit)
			}
		}
	}
	n := len(spaces)
	var dfs func(int)bool
	dfs = func(pos int)bool{
		if pos == n {
			return true
		}
		i, j := spaces[pos][0], spaces[pos][1]
		// 0x1ff 即二进制的 9 个 1, 取反 将情况反过来1 表示未用数字， 0 表示已用数字， 0x1ff 与 去掉高位无关
		// int mask = ~(line[i] | column[j] | block[i / 3][j / 3]) & 0x1ff; <== CLanguage
		mask := 0x1ff &^ uint(line[i]|column[j]|block[i/3][j/3]) // 注意Golang 的取反 运算
		for ;mask > 0; mask &= mask-1{ // mask &= mask-1 最右侧末尾1 置 0
			digit := byte(bits.TrailingZeros(mask))
			flip(i, j, digit)
			board[i][j] = digit + '1'
			if dfs(pos+1){
				return true
			}
			flip(i, j, digit)// flip 为何采用异或的原因
		}
		return false
	}
	dfs(0)
}
// 官方的枚举优化：
// 如果一个空白格只有唯一的数可以填入，也就是其对应的 b 值和 b−1 进行按位与运算后得到 0（即 b 中只有一个二进制位为 1）。
// 此时，我们就可以确定这个空白格填入的数，而不用等到递归时再去处理它。
// 这样一来，我们可以不断地对整个数独进行遍历，将可以唯一确定的空白格全部填入对应的数
func solveSudoku(board [][]byte) {
	var line, column [9]int
	var block [3][3]int
	var spaces [][2]int
	flip := func(i, j int, digit byte){
		line[i] ^= 1 << digit
		column[j] ^= 1 << digit
		block[i/3][j/3] ^= 1 << digit
	}
	for i := range board{
		for j, c := range board[i]{
			if c != '.'{
				digit := c - '1'
				flip(i, j, digit)
			}
		}
	}
	for { // 识别一个 可以产生一系列连锁反应
		modified := false
		for i := range board{
			for j, c := range board[i]{
				if c != '.'{ continue }
				// mask 不可能为 0
				mask := 0x1FF & ^uint(line[i] | column[j] | block[i/3][j/3])
				if mask & (mask-1) == 0{// 表示mask 的二进制表示中仅余 一个 1，此时可以直接填数.
					digit := byte(bits.TrailingZeros(mask))
					flip(i, j, digit)
					board[i][j] = digit + '1'
					modified = true
				}
			}
		}
		if !modified {
			break
		}
	}
	for i := range board{
		for j, c := range board[i]{
			if c == '.'{
				spaces = append(spaces, [2]int{i,j})
			}
		}
	}
	n := len(spaces)
	var dfs func(int)bool
	dfs = func(pos int)bool{
		if pos == n{ return true }
		i, j := spaces[pos][0], spaces[pos][1]
		mask := 0x1FF & ^uint(line[i] | column[j] | block[i/3][j/3])
		for ; mask > 0; mask &= (mask-1){
			digit := byte(bits.TrailingZeros(mask))
			flip(i, j, digit)
			board[i][j] = digit + '1'
			if dfs(pos+1){
				return true
			}
			flip(i, j, digit)
		}
		return false
	}
	dfs(0)
}

/* 51. N-Queens
** The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.
** Given an integer n, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.
** Each solution contains a distinct board configuration of the n-queens' placement,
** where 'Q' and '.' both indicate a queen and an empty space, respectively.
 */
// 2022-01-20 刷出此题
func solveNQueens(n int) [][]string {
	ans := [][][2]int{}
	rec := [][2]int{}
	dir := [][2]int{[2]int{1,1}, [2]int{1,-1}, [2]int{-1,1}, [2]int{-1,-1}}
	valid := func(x, y int)bool{
		for i := range rec{
			r, c := rec[i][0], rec[i][1]
			if x == r || y == c {
				return false
			}
		}
		// 斜线
		for _, d :=range dir{
			for i, j := x+d[0], y+d[1]; i >= 0 && j >= 0 && i < n && j < n; i, j = i+d[0], j+d[1]{
				for k := range rec{
					r, c := rec[k][0], rec[k][1]
					if i == r && j == c{
						return false
					}
				}
			}
		}
		return true
	}
	var dfs func(x int)bool
	dfs = func(x int)bool{
		if x == n {
			if len(rec) == n{
				ans = append(ans, append([][2]int{}, rec...))
				return true
			}
			return false
		}
		ok := false
		for i := 0; i < n; i++{
			if valid(x, i){
				rec = append(rec, [2]int{x, i})
				if dfs(x+1){
					ok = true
				}
				rec = rec[:len(rec)-1]
			}
		}
		return ok
	}
	dfs(0)
	ret := make([][][]byte, len(ans))
	for i := range ret{
		ret[i] = make([][]byte, n)
		for j := range ret[i]{
			ret[i][j] = make([]byte, n)
			for k := range ret[i][j]{
				ret[i][j][k] = '.'
			}
		}
	}
	for i := range ans{
		for _, d := range ans[i]{
			ret[i][d[0]][d[1]] = 'Q'
		}
	}
	result := make([][]string, len(ans))
	for i := range ret{
		for j := range ret[i]{
			result[i] = append(result[i], string(ret[i][j]))
		}
	}
	return result
}

/* 基于集合的回溯
** 为了判断一个位置所在的列和两条斜线上是否已经有皇后，
** 使用三个集合columns, diagonals1, diagonals2 分别记录每一列以及两个方向的每条斜线上是否有Queen
**
 */
func solveNQueens_set(n int) [][]string {

}

func solveNQueens_bit(n int) [][]string {

}














