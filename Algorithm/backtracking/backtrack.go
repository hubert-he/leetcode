package backtracking

import (
	"fmt"
	"sort"
	"strconv"
)
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

/* 489. Robot Room Cleaner
** You are controlling a robot that is located somewhere in a room.
** The room is modeled as an m x n binary grid where 0 represents a wall and 1 represents an empty slot.
** The robot starts at an unknown location in the root that is guaranteed to be empty,
** and you do not have access to the grid, but you can move the robot using the given API Robot.
** You are tasked to use the robot to clean the entire room (i.e., clean every empty cell in the room).
** The robot with the four given APIs can move forward, turn left, or turn right. Each turn is 90 degrees.
** When the robot tries to move into a wall cell, its bumper sensor detects the obstacle, and it stays on the current cell.
** Design an algorithm to clean the entire room using the following APIs:
	interface Robot {
	  // returns true if next cell is open and robot moves into the cell.
	  // returns false if next cell is obstacle and robot stays on the current cell.
	  boolean move();

	  // Robot will stay on the same cell after calling turnLeft/turnRight.
	  // Each turn will be 90 degrees.
	  void turnLeft();
	  void turnRight();

	  // Clean the current cell.
	  void clean();
	}
** Note that the initial direction of the robot will be facing up.
** You can assume all four edges of the grid are all surrounded by a wall.
 */
type Robot struct {}
func (robot *Robot) Move() bool {return true}
func (robot *Robot) TurnLeft() {}
func (robot *Robot) TurnRight() {}
func (robot *Robot) Clean() {}
func cleanRoom(robot *Robot) {
	visited := map[[2]int]bool{}
	// 思维误区：考虑方向，顺时针=> 0:up  1:right  2: down, 3:left
	dirs := [][]int{[]int{-1,0}, []int{0,1}, []int{1,0}, []int{0,-1}}
	var dfs func(cur [2]int, direction int)
	dfs = func(cur [2]int, direction int){
		visited[cur] = true
		robot.Clean()
		// 思维误区：考虑方向，顺时针: 0: 'up', 1: 'right', 2: 'down', 3: 'left'
		for i := range dirs{
			next_dir := (direction+i) % 4
			next := [2]int{cur[0]+dirs[next_dir][0], cur[1]+dirs[next_dir][1]}
			if !visited[next] && robot.Move(){
				dfs(next, next_dir)
				// 此处递归回来后，robot 要回到原位置 并恢复朝向
				robot.TurnLeft()
				robot.TurnLeft()
				robot.Move()
				robot.TurnLeft()
				robot.TurnLeft()
			}
			// turn the robot following chosen direction : clockwise
			// 所以前提要求 robot 方向要恢复
			robot.TurnRight()
		}
	}
	dfs([2]int{0,0}, 0)// 0 向前
}

/* 282. Expression Add Operators
** Given a string num that contains only digits and an integer target,
** return all possibilities to insert the binary operators '+', '-', and/or '*' between the digits of num
** so that the resultant expression evaluates to the target value.
Note that operands in the returned expressions should not contain leading zeros.
 */
/* 往 num 中间的 n−1 个空隙添加 + 号、- 号或 * 号，或者不添加符号
** 隐含的要求：
**	1. 有乘法，所以需要考虑 运算符优先级的问题
**	2. 前导0 情况
**	3. + 和 - 不作为一元运算符， 也即 运算符不能出现在表达式的首部
 */
// 2022-01-18 刷出此题
func AddOperators(num string, target int) []string {
	exp := [][]byte{}
	op := []byte{'+', '-', '*'}
	for i := range num{
		/*
		for j := range exp{
			if exp[j] == nil { continue }
			//t := exp[j]  shallow copy 会污染
			t := make([]byte, len(exp[j]))
			copy(t, exp[j])
			// 消除前缀0的情况
			if i > 0 && num[i-1] != '0'{
				exp[j] = append(exp[j], num[i]) // 情况1：不加符号的情况
			}else{// 清理到残缺的
				// exp[j] = nil 方式1 置 nil
				exp = append(exp[:j], exp[j+1:]...)// 方式2 直接删
			}
			for k := range op{
				exp = append(exp, append(t, op[k], num[i]))
			}
		} */
		tmp := [][]byte{}
		for j := range exp{
			n := len(exp[j])
			t := make([]byte, n)
			copy(t, exp[j])
			// 消除前缀0的情况
			//if i > 0 && num[i-1] != '0'{ 这个条件会漏情况， 100*0 此100 这种情况
			var o int
			for o = n-1; o >= 0 && t[o] != '+' && t[o] != '-' && t[o] != '*'; o--{ }
			if t[o+1] != '0'{
				//tmp = append(tmp, append(exp[j], num[i]))
				tmp = append(tmp, append(t, num[i]))
			}
			for k := range op{
				//tmp = append(tmp, append(exp[j], ops[k], num[i]))
				tmp = append(tmp, append(t, op[k], num[i]))
			}
		}
		exp = tmp
		if len(exp) == 0{
			exp = append(exp, []byte{num[i]})
		}
	}
	// 此compute 实现方式 可参见 面试题 16.26. Calculator LCCI（实现在题目下方）
	compute := func(exp string)int{
		tmp, result := 0, 0
		preop, sign := '+', 1
		st := []int{}
		for i := range exp{
			c := exp[i]
			switch c {
			case '+':
				if preop == '*' {
					st[len(st)-1] *= tmp * sign
				} else {
					st = append(st, tmp*sign)
				}
				preop, sign = '+', 1
				tmp = 0
			case '-':
				if preop == '*' {
					st[len(st)-1] *= tmp * sign
				} else {
					st = append(st, tmp*sign)
				}
				preop, sign = '-', -1
				tmp = 0
			case '*':
				if preop == '*' {
					st[len(st)-1] *= tmp * sign
				} else {
					st = append(st, tmp*sign)
				}
				preop, sign = '*', 1
				tmp = 0
			default:
				tmp = tmp*10 + int(c-'0')
			}
		}
		//st = append(st, tmp) 需要分情况
		if preop == '*'{
			st[len(st)-1] *= tmp * sign
		}else{
			st = append(st, tmp*sign)
		}
		for i := range st{
			result += st[i]
		}
		return result
	}
	ans := []string{}
	for i := range exp{
		fmt.Println(string(exp[i]), compute(string(exp[i])))
		if compute(string(exp[i])) == target{
			ans = append(ans, string(exp[i]))
		}
	}
	return ans
}

func AddOperators_DFS(num string, target int) []string {
	n := len(num)
	ans := []string{}
	var dfs func(idx int, prev int, cur int, s string)
	dfs = func(idx int, prev int, cur int, s string) {
		if idx == n{
			if cur == target { // 符合情况的表达式
				ans = append(ans, s)
			}
			return
		}
		for i := idx; i < n; i++{
			if i != idx && num[idx] == '0'{
				break
			}
			next, _ := strconv.Atoi(num[idx:i+1])
			if idx == 0{
				dfs(i+1, next, next, fmt.Sprintf("%d", next))
			}else{
				dfs(i+1, next, cur + next, fmt.Sprintf("%s+%d", s, next))
				dfs(i+1, -next, cur - next, fmt.Sprintf("%s-%d", s, next))
				x := prev * next // 运算符的优先级问题, 先➖ 再 ➕
				dfs(i+1, x, cur - prev + x, fmt.Sprintf("%s*%d", s, next))
			}
		}
	}
	dfs(0, 0, 0, "")
	return ans
}
// 官方题解：当前最优
func AddOperators_backtrack(num string, target int) []string {
	n := len(num)
	ans := []string{}
	var backtrack func(expr []byte, i, res, mul int)
	// expr: 为当前构建出的表达式  res: 当前表达式的计算结果  mul: 表达式最后一个连乘串的计算结果
	backtrack = func(expr []byte, i, res, mul int) {
		if i == n {
			if res == target {
				ans = append(ans, string(expr))
			}
			return
		}
		signIndex := len(expr)
		if i > 0 {
			expr = append(expr, 0)// 占位，下面填充符号
		}
		// 枚举截取的数字长度（取多少位），注意数字可以是单个 0 但不能有前导零
		for j, val := i, 0; j < n && (j == i || num[i] != 0); j++{
			val = val * 10 + int(num[j] - '0') // 边递归 边计算
			expr = append(expr, num[j])
			if i == 0{// 表达式开头不能添加符号
				backtrack(expr, j+1, val, val)
			}else{// 枚举符号
				expr[signIndex] = '+'; backtrack(expr, j+1, res + val, val) // val单独组成表达式最后一个连乘串；
				expr[signIndex] = '-'; backtrack(expr, j+1, res - val, -val) // -val 单独组成表达式最后一个连乘串；
				// 由于乘法运算优先级高于加法和减法运算，我们需要对res 撤销之前 mul 的计算结果，并添加新的连乘结果 mul∗val，
				// 也就是将 res 减少 mul 并增加 mul∗val
				expr[signIndex] = '*'; backtrack(expr, j+1, res - mul + mul * val, mul * val)
			}
		}
	}
	backtrack(make([]byte, 0, n*2-1), 0, 0, 0)
	return ans
}

/* 面试题 16.26. Calculator LCCI
** Given an arithmetic equation consisting of positive integers, +, -, * and / (no paren­theses), compute the result.
** The expression string contains only non-negative integers, +, -, *, / operators and empty spaces .
** The integer division should truncate toward zero.
** 1. 哨兵来消除 最后尾巴num 的处理
** 2. 运算符优先级的处理： 往前看一个运算符和数
 */
func calculate_lcci(s string) int {
	ans, prenum, num := 0, 0, 0
	var preop byte = '+'
	sb := []byte(s)
	sb = append(sb, 'x') // 哨兵，注意体会此方法 哨兵的用意
	for _, c := range sb{
		if c == ' '{ continue }
		if c >= '0' && c <= '9'{
			num = num * 10 + int(c-'0')
		}else{ // 哨兵作用：处理最后一个数字
			if preop == '+'{
				ans += prenum
				prenum = num
			}
			if preop == '-'{
				ans += prenum
				prenum = -num
			}
			if preop == '*'{
				prenum = prenum * num
			}
			if preop == '/'{
				prenum = prenum / num
			}
			num, preop = 0, c
		}
	}
	return ans + prenum
}





