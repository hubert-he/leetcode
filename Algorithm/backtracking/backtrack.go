package backtracking

/* 22. Generate Parentheses
** Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
["(((())))","((()()))","((())())","((()))()","(()(()))","(()()())","(()())()","(())(())","(())()()","()((()))","()(()())","()(())()","()()(())","()()()()"]
["(((())))","((()()))","((())())","((()))()","(()(()))","(()()())","(()())()",           "(())()()","()((()))","()(()())","()(())()","()()(())","()()()()"]
 */
/* 这种计算方式不对，对漏掉 (())(()) 情况
func generateParenthesis(n int) []string {
	m := map[string]bool{"()":true}
	for i := 1; i < n; i++{
		t := map[string]bool{}
		for k := range m{
			t["("+k+")"] = true
			t["()"+k] = true
			t[k+"()"] = true
		}
		m = t
	}
	ans := []string{}
	for k := range m{
		ans = append(ans, k)
	}
	return ans
}
*/
/* 暴力: 生成所有2的2n次方个序列，然后检查每一个序列是否有效
** 长度为 n 的序列就是在长度为 n-1 的序列前加一个 '(' 或 ')'
** 序列有效性检查：
**    遍历这个序列，并使用一个变量 balance 表示左括号的数量减去右括号的数量。
**    如果在遍历过程中 balance 的值小于零，或者结束时 balance 的值不为零，那么该序列就是无效的，否则它是有效的
 */

func generateParenthesis(n int) []string {
	ans := []string{}
	valid := func(s []byte)bool{
		balance := 0
		for _, c := range s{
			if c == '('{
				balance++
			}
			if c == ')'{
				balance--
			}
			if balance < 0 {
				return false
			}
		}
		return balance == 0
	}
	var dfs func(current []byte, pos int)
	dfs = func(current []byte, pos int){
		if pos == len(current){
			if valid(current){
				ans = append(ans, string(current))
			}
			return
		}
		current[pos] = '('
		dfs(current, pos+1)
		current[pos] = ')'
		dfs(current, pos+1)
	}
	dfs(make([]byte, n*2), 0)
	return ans
}

func GenerateParenthesisBt(n int) []string {
	ans := []string{}
	var backtrack func(cur []byte, open int, close int)
	backtrack = func(cur []byte, open int, close int){
		if len(cur) == 2*n{
			ans = append(ans, string(cur))
			return
		}
		if open < n{
			cur = append(cur, '(')
			backtrack(cur, open+1, close)
			cur = cur[:len(cur)-1]
		}
		if close < open{
			cur = append(cur, ')')
			backtrack(cur, open, close+1)
			cur = cur[:len(cur)-1]
		}
	}
	backtrack([]byte{}, 0,0)
	return ans
}












