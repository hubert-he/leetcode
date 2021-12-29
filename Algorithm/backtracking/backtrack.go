package backtracking

import "sort"

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

/* 39. Combination Sum
** Given an array of distinct integers candidates and a target integer target,
** return a list of all unique combinations of candidates where the chosen numbers sum to target.
** You may return the combinations in any order.
** The same number may be chosen from candidates an unlimited number of times.
** Two combinations are unique if the frequency of at least one of the chosen numbers is different.
** It is guaranteed that the number of unique combinations that sum up to target is less than 150 combinations for the given input.
 */
// 针对每个元素，都有选 或 不选择两个
func CombinationSum(candidates []int, target int) [][]int {
	length := len(candidates)
	result, path := [][]int{}, []int{} // 此时， result的值实际为nil
	var backtrack func(int,int)
	backtrack = func(sum, start int){
		if sum == 0{
			path_copy := make([]int, len(path))
			copy(path_copy, path)
			result = append(result, path_copy)
			return
		}
		if sum < 0 || start >= length{
			return
		}
		// 跳过，不选择candidates[start]
		backtrack(sum, start + 1)
		// 选择candidates[start]
		path = append(path, candidates[start])
		backtrack(sum - candidates[start], start)
		path = path[:len(path) - 1]
	}
	backtrack(target, 0)
	return result
}
func CombinationSumII(candidates []int, target int) (result [][]int){
	sort.Ints(candidates) // 剪枝
	length := len(candidates)
	path := []int{}
	//var backtrack func(int)
	var backtrack func(int, int) // 增加start变量，去重
	//backtrack = func(sum int){
	backtrack = func(sum, start int){
		if sum == target{
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			result = append(result, pathCopy)
			return
		}
		for i := start; i < length; i++{ // 控制决策树中每层的元素选取
			num := candidates[i]
			//fmt.Printf("befor start=> %v, sum => %d, add => %d\n", path, sum, num)
			// 剪枝，序列有序前提下
			if sum + num > target{
				//fmt.Printf("cut all start=> %v, sum => %d\n", path, sum)
				break
			}
			if i > 0 && num == candidates[i-1]{
				//fmt.Printf("cut start=> %v, sum => %d\n", path, sum)
				continue
			}
			path = append(path, num)
			//backtrack(sum + num, start+1)
			backtrack(sum+num, i) // 指定当层只能往后选择
			path = path[:len(path) - 1]
			//fmt.Printf("after start=> %v, sum => %d\n", path, sum)
		}
	}
	backtrack(0, 0)
	return
}
/* 2021-12-14 重刷此题 直接DFS*/
func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	for i, c := range candidates{
		if target < c {
			continue
		}else if target > c{
			r := combinationSum(candidates[i:], target-c)
			for j := range r{
				ans = append(ans, append([]int{c}, r[j]...))
			}
		}else{
			ans = append(ans, []int{c})
		}
	}
	return ans
}

// 40. Combination Sum II
/*
由于我们需要求出所有和为target的组合，并且每个数只能使用一次，因此我们可以使用递归 + 回溯的方法来解决这个问题：
1. 我们用dfs(pos,rest)表示递归的函数，其中 pos 表示我们当前递归到了数组 candidates 中的第 pos 个数，
   而 rest 表示我们还需要选择和为 rest 的数放入列表作为一个组合；
2. 对于当前的第 pos 个数，我们有两种方法：选或者不选。如果我们选了这个数，那么我们调用 dfs(pos+1,rest−candidates[pos]) 进行递归，
   注意这里必须满足 rest ≥ candidates[pos]。如果我们不选这个数，那么我们调用 dfs(pos+1,rest) 进行递归；
3. 在某次递归开始前，如果 rest 的值为 0，说明我们找到了一个和为 target 的组合，将其放入答案中。
   每次调用递归函数前，如果我们选了那个数，就需要将其放入列表的末尾，该列表中存储了我们选的所有数。
   在回溯时，如果我们选了那个数，就要将其从列表的末尾删除。
标准的递归+回溯，但是存在重复组合
下面解法会出现重复解
*/
func CombinationSum2II(candidates []int, target int) [][]int {
	result,path := [][]int{}, []int{}
	length := len(candidates)
	used := make([]bool, length)
	var backtrack func(int, int)
	backtrack = func(sum, start int){
		if sum == 0{
			path_copy := make([]int, len(path))
			copy(path_copy, path)
			result = append(result, path_copy)
			return
		}
		if sum < 0 || start >= length{
			return
		}
		backtrack(sum, start + 1)
		if !used[start]{
			path = append(path, candidates[start])
			used[start] = true
			backtrack(sum - candidates[start], start)
			path = path[:len(path) - 1]
			used[start] = false
		}
	}
	backtrack(target, 0)
	return result
}
/*
这个方法最重要的作用是，可以让同一层级，不出现相同的元素。即
                  1
                 / \
                2   2  这种情况不会发生 但是却允许了不同层级之间的重复即：
               /     \
              5       5
                例2
                  1
                 /
                2      这种情况确是允许的
               /
              2
*/
func CombinationSum2(candidates []int, target int) [][]int {
	result,path := [][]int{}, []int{}
	length := len(candidates)
	if length == 0 || target <= 0{ //处理掉特殊情况
		return result
	}
	sort.Ints(candidates) // sort inplace，方便去重

	var backtrack func(int, int)
	backtrack = func(sum, start int){
		if sum == 0{
			path_copy := make([]int, len(path))
			copy(path_copy, path)
			result = append(result, path_copy)
			return
		}
		if sum < 0 || start >= length{
			return
		}
		for i := start; i < length; i++{
			if i != start && candidates[i] == candidates[i-1]{
				continue // 跳过和本轮循环中已访问过元素数值相同的元素们。
			}
			path = append(path, candidates[i])
			backtrack(sum - candidates[i], i+1)
			path = path[:len(path) - 1] // 回溯
		}

	}
	backtrack(target, 0)
	return result
}












