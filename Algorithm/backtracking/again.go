package backtracking

/* 79. Word Search
** Given an m x n grid of characters board and a string word, return true if word exists in the grid.
** The word can be constructed from letters of sequentially adjacent cells,
** where adjacent cells are horizontally or vertically neighboring.
** The same letter cell may not be used more than once. <需要访问计数>
 */
func Exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	vis := make([][]bool, m)
	for i := range vis{
		vis[i] = make([]bool, n)
	}
	dir := [][2]int{[2]int{0,1}, [2]int{0,-1}, [2]int{1,0}, [2]int{-1,0}}
	var dfs func(start, end, idx int)bool
	dfs = func(start, end, idx int)bool{
		if start < 0 || start >= m || end < 0 || end >= n{
			return false
		}
		if vis[start][end] || board[start][end] != word[idx]{
			return false
		}
		if idx == len(word) - 1{// 易错点-1： idx == len(word)
			return true
		}
		vis[start][end] = true
		for i := range dir{
			x, y := dir[i][0], dir[i][1]
			if dfs(start+x, end+y, idx+1){
				return true
			}
		}
		vis[start][end] = false // 遗漏点-1： 恢复状态
		return false
	}
	for i := range board{
		for j := range board[i]{
			if dfs(i, j, 0){
				return true
			}
		}
	}
	return false
}

/* 17. Letter Combinations of a Phone Number
** Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
** Return the answer in any order.
** A mapping of digit to letters (just like on the telephone buttons) is given below.
** Note that 1 does not map to any letters.
 */
// 2021-11-12 刷过
func LetterCombinations(digits string) []string {
	m := [8][]string{[]string{"a", "b", "c"}, []string{"d", "e", "f"},
		[]string{"g", "h", "i"},[]string{"j", "k", "l"},[]string{"m", "n", "o"},
		[]string{"p", "q", "r", "s"},[]string{"t", "u", "v"},[]string{"w", "x", "y", "z"}}
	var dfs func(s string)[]string
	dfs = func(s string)[]string{
		n := len(s)
		ret := []string{}
		if n < 1{ // 遗漏点-1： 特殊情况 ""
			return ret
		}
		if n == 1{
			return m[s[0]-'2']
		}
		t := LetterCombinations(s[1:])
		for _,c := range m[s[0]-'2']{
			for i := range t{
				ret = append(ret, c+t[i])
			}
		}
		return ret
	}
	return dfs(digits)
}
// 官方题解
func LetterCombinations2(digits string) []string {
	phoneMap := [8]string{
		"abc", "def", "ghi", "jkl",
		"mno", "pqrs","tuv", "wxyz",
	}
	n := len(digits)
	combinations := []string{}
	if n == 0{
		return combinations
	}
	var backtrack func(idx int, comb string)
	backtrack = func(idx int, comb string){
		if idx == n{
			combinations = append(combinations, comb)
			return
		}
		letters := phoneMap[digits[idx]-'2']
		for _, c := range letters{
			backtrack(idx+1, comb+string(c))
		}
	}
	backtrack(0, "")
	return combinations
}