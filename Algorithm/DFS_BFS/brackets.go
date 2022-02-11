package DFS_BFS

import (
	"fmt"
	"math/bits"
	"sort"
	"strconv"
	"strings"
	"text/scanner"
)

// 括号类题目
/* 394. Decode String
** Given an encoded string, return its decoded string.
** The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times.
** Note that k is guaranteed to be a positive integer.
** You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc.
** Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k.
** For example, there will not be input like 3a or 2[4].
** Constraints:
	1 <= s.length <= 30
	s consists of lowercase English letters, digits, and square brackets '[]'.
	s is guaranteed to be a valid input.
	All the integers in s are in the range [1, 300].
同类题目：
0. 471. Encode String with Shortest Length
1. 726. Number of Atoms
2. 784. Letter Case Permutation
3. 1087. Brace Expansion
4. 1096. Brace Expansion II
*/
/* 2022-01-04 刷出此题
** 此题学习下test case
** 	Input: s = "3[a]2[bc]"
	Output: "aaabcbc"
** 	Input: s = "3[a2[c]]"
	Output: "accaccacc"
**	Input: s = "2[abc]3[cd]ef"
	Output: "abcabccdcdcdef"
*/
func decodeString(s string) string {
	if !strings.Contains(s, "["){
		return s
	}
	var ans string
	i, n := 0, len(s)
	for i < n{
		j := i
		for i < n && s[i] >= 'a' && s[i] <= 'z'{
			i++
		}
		ans += s[j:i]
		if i >= n || (s[i] < '0' && s[i] > '9'){
			break
		}
		j = i
		for i < n && s[i] >= '0' && s[i] <= '9'{
			i++
		}
		repeat, _ := strconv.Atoi(s[j:i])
		if repeat <= 0{
			break
		}
		// 寻找 [ 和 匹配的 ]
		i++
		j = i
		cnt := 1
		for cnt > 0{
			if s[i] == '['{
				cnt++
			}
			if s[i] == ']'{
				cnt--
			}
			i++
		}
		ret := decodeString(s[j:i-1])
		ans += strings.Repeat(ret, repeat)
	}
	return ans
}
/* 转换为栈
** 本题中可能出现括号嵌套的情况，比如 2[a2[bc]]，这种情况下我们可以先转化成 2[abcbc]，在转化成 abcbcabcbc。
** 我们可以把字母、数字和括号看成是独立的 TOKEN，并用栈来维护这些 TOKEN。具体的做法是，遍历这个栈：
	1. 如果当前的字符为数位，解析出一个数字（连续的多个数位）并进栈
	2. 如果当前的字符为字母或者左括号，直接进栈
	3. 如果当前的字符为右括号，开始出栈，一直到左括号出栈，出栈序列反转后拼接成一个字符串，
		此时取出栈顶的数字（此时栈顶一定是数字，想想为什么？），就是这个字符串应该出现的次数，
		我们根据这个次数和字符串构造出新的字符串并进栈
** 重复如上操作，最终将栈中的元素按照从栈底到栈顶的顺序拼接起来，就得到了答案。
** 注意：这里可以用不定长数组来模拟栈操作，方便从栈底向栈顶遍历
 */
func decodeString_stack(s string) string {
	st := []string{}
	i, n := 0, len(s)
	for i < n{
		if s[i] >= '0' && s[i] <= '9'{
			j := i
			i++
			for s[i] >= '0' && s[i] <= '9'{
				i++
			}
			st = append(st, s[j:i])
		}else if s[i] >= 'a' && s[i] <= 'z' || s[i] == '['{
			st = append(st, string(s[i]))
			i++
		}else{ // 当前的字符为右括号
			i++
			sub := []string{}
			for st[len(st)-1] != "["{ // 一直到左括号出栈
				sub = append(sub, st[len(st)-1])
				st = st[:len(st)-1]
			}
			// 出栈序列反转
			for j := 0; j < len(sub)/2; j++{
				sub[j], sub[len(sub)-j-1] = sub[len(sub)-j-1], sub[j]
			}
			st = st[:len(st)-1]
			repeat, _ := strconv.Atoi(st[len(st)-1])
			st = st[:len(st)-1]
			t := strings.Repeat(strings.Join(sub, ""), repeat)
			st = append(st, t)
		}
	}
	return strings.Join(st, "")
}

/* 471. Encode String with Shortest Length -- 此题也收录进DP
** Given a string s, encode the string such that its encoded length is the shortest.
** The encoding rule is: k[encoded_string],
** where the encoded_string inside the square brackets is being repeated exactly k times.
** k should be a positive integer.
** If an encoding process does not make the string shorter, then do not encode it.
** If there are several solutions, return any of them.
** 此题目也属于区间DP 问题
** 459.重复的子字符串  --- 找到连续重复的子字符串，我们才能进行编码(压缩)。
 */
/* DP 部分：
** 设s(i,j)表示子串s[i,…,j]。串的长度为len=j-i+1
** 用d[i][j]表示s(i,j)的最短编码串。当s(i,j)有连续重复子串时，s(i,j)可编码为”k[重复子串]”的形式.d[i][j]= "k[重复子串]"
** 当len < 5时，s(i,j)不用编码。d[i][j]=s(i,j)
** 当len > 5时，s(i,j)的编码串有两种可能。
** 现将s(i,j)分成两段s(i,k)和s(k+1,j)，(i <= k <j) 推导出状态方程
** d[i][j] = d[i][k] + d[k+1][j]  当d[i][k].length + d[k+1][j].length < d[i][j].length时
**
** 题目难点部分：快速求出字符串中连续的重复子串
** 枚举子串逐个查找，用上kmp，lcp类的算法，进行加速
** 另有一个方法： 对字符串s，s与s拼接成t=s+s
** 在 t 中 从索引位置 1 开始查找 s 如果查找到，即 在位置 p 处开始， t 中出现了 s
** 注意： t 中肯定可以查找到 s （从索引位置1开始搜索的前提下）
** 当 p >= len(s) 时，说明s中没有连续的重复子串, 不能压缩
** 当 p < len(s) 时，说明s中有连续重复子串并且 连续重复子串是 s[0:p], 重复个数为 len(s) / p
**
*/
func encode(s string) (ans string) {
	n := len(s)
	dp := make([][]string, n)
	for i := range dp{
		dp[i] = make([]string, n)
	}
	var dfs func(start, end int)string
	dfs = func(start, end int) string{
		if start > end { return ""}
		if len(dp[start][end]) > 0{
			return dp[start][end]
		}
		length := end - start + 1
		ss := s[start:end+1]
		if length < 5 { return ss }
		ret := ss // 初始最大
		p := strings.Index((ss+ss)[1:], ss) + 1 // 从索引1开始查找是否有重复子串
		if p > 0 && p < length { // ss 存在重复子串
			ret = fmt.Sprintf("%d[%s]", length/p, dfs(start, start+p-1))
			dp[start][end] = ret
			return ret
		}
		// 动态规划部分
		for mid := start; mid < end; mid++{
			s1 := dfs(start, mid)
			s2 := dfs(mid+1, end)
			if len(s1) + len(s2) < len(ret){
				ret = s1+s2
			}
		}
		dp[start][end] = ret
		return ret
	}
	return dfs(0, n-1)
}
// 区间DP
// dp[i][j] 来自 1. 存在连续的重复子串 2. 分成2段dp[i][k] 和 dp[k+1][j]
func encodeDP(s string) (ans string) {
	n := len(s)
	dp := make([][]string, n)
	for i := range dp{
		dp[i] = make([]string, n)
	}
	for length := 1; length <= n; length++{ // 从长度开始枚举
		for i := 0; i + length <= n; i++{
			j := i + length - 1
			ss := s[i:j+1]
			dp[i][j] = ss
			//if length > 5{
			if length > 4{
				p := strings.Index((ss+ss)[1:], ss) + 1
				if p > 0 && p < len(ss){
					//dp[i][j] = fmt.Sprintf("%d[%s]", len(ss)/p, ss[:p]) 不能直接ss[:p] 可能源串有压缩情况
					dp[i][j] = fmt.Sprintf("%d[%s]", len(ss)/p, dp[i][i+p-1])
				}else{
					for k := i; k < j; k++{ // 注意不要与 切片操作搞混
						//if len(dp[i][k+1]) + len(dp[k+1][j+1]) < len(dp[i][j]){
						if len(dp[i][k]) + len(dp[k+1][j]) < len(dp[i][j]){
							dp[i][j] = dp[i][k] + dp[k+1][j]
						}
					}
				}
			}
		}
	}
	return dp[0][n-1]
}

/* 1087. Brace Expansion
** You are given a string s representing a list of words. Each letter in the word has one or more options.
** If there is one option, the letter is represented as is.
** If there is more than one option, then curly braces delimit the options.
** For example, "{a,b,c}" represents options ["a", "b", "c"].
** For example, if s = "a{b,c}", the first character is always 'a', but the second character can be 'b' or 'c'.
** The original list is ["ab", "ac"].
** Return all words that can be formed in this manner, sorted in lexicographical order.
** Constraints:
	1 <= s.length <= 50
	s consists of curly brackets '{}', commas ',', and lowercase English letters.
	s is guaranteed to be a valid input.
	There are no nested curly brackets.  注意这条，题目中规定： 不存在嵌套的花括号！！ "{a,b{c,d}}" 非法输入
	All characters inside a pair of consecutive opening and ending curly brackets are different.
 */
func expand(s string) []string {
	n := len(s)
	if !strings.Contains(s, "{"){
		return []string{s}
	}
	ans := []string{}
	var backTrace func(sb []byte, start int)
	backTrace = func(sb []byte, start int){
		if start == n{
			ans = append(ans, string(sb))
			return
		}
		if s[start] == '{'{
			cnt := 0
			//计算大括号内容，下次跳转位置为 start + cnt + 2
			for j := start+1; s[j] != '}'; j++{
				cnt++
			}
			for j := start + 1; s[j] != '}'; j++{
				c := s[j]
				if c != ','{
					sb = append(sb, c)
					backTrace(sb, start+cnt+2)
					sb = sb[:len(sb)-1] // 回溯
				}
			}
		}else{
			sb = append(sb, s[start])
			backTrace(sb, start+1)
			sb = sb[:len(sb)-1] // 回溯
		}
	}
	backTrace([]byte{}, 0)
	sort.Strings(ans)
	return ans
}
// 2022-01-04 刷出此题
func expand2(s string) []string {
	n := len(s)
	if !strings.Contains(s, "{"){
		return []string{s}
	}
	// 题目要求没有 嵌套的括号，逗号无需额外处理
	ans := []string{""} // 加空字符串，否则后面循环会空
	for i := 0; i < n; i++{
		if s[i] == '{'{
			i++
			j := i
			for s[i] != '}'{ i++ }
			ret := strings.Split(s[j:i], ",")
			nt := []string{}
			for o := range ans{
				for p := range ret{
					nt = append(nt, strings.Join([]string{ans[o], ret[p]}, ""))
				}
			}
			ans = nt
		}else{
			for o := range ans{
				ans[o] = strings.Join([]string{ans[o], string(s[i])}, "")
			}
		}
	}
	sort.Strings(ans)
	return ans
}

/* 1096. Brace Expansion II
** Under the grammar given below, strings can represent a set of lowercase words.
** Let R(expr) denote the set of words the expression represents.
** The grammar can best be understood through simple examples:
** 1. Single letters represent a singleton set containing that word:
		R("a") = {"a"}
		R("w") = {"w"}
** 2. When we take a comma-delimited list of two or more expressions, we take the union of possibilities.
		R("{a,b,c}") = {"a","b","c"}
		R("{{a,b},{b,c}}") = {"a","b","c"} (notice the final set only contains each word at most once)
** 3. When we concatenate two expressions,
      we take the set of possible concatenations between two words
	  where the first word comes from the first expression and the second word comes from the second expression.
		R("{a,b}{c,d}") = {"ac","ad","bc","bd"}
		R("a{b,c}{d,e}f{g,h}") = {"abdfg", "abdfh", "abefg", "abefh", "acdfg", "acdfh", "acefg", "acefh"}
** Formally, the three rules for our grammar:
	1. For every lowercase letter x, we have R(x) = {x}.
	2. For expressions e1, e2, ... , ek with k >= 2, we have R({e1, e2, ...}) = R(e1) ∪ R(e2) ∪ ...
	3. For expressions e1 and e2, we have R(e1 + e2) = {a + b for (a, b) in R(e1) × R(e2)},
		where + denotes concatenation, and × denotes the cartesian product.
** Given an expression representing a set of words under the given grammar,
** return the sorted list of words that the expression represents.
Constraints:
	1 <= expression.length <= 60
	expression[i] consists of '{', '}', ','or lowercase English letters.
	The given expression represents a set of words based on the grammar given in the description.
题目意义：在Shell编程中 {} 花括号展开运算，可以用来生成任意字符，表达式之间允许嵌套，单一元素与表达式的连接也是允许的
例如bash中：
	# echo file{1,2}
	  file1 file2
	# echo file{1,a{b,c}}
      file1 fileab fileac
 */
// 注意此题 花括号 就是可以嵌套的了
func braceExpansionII(expression string) []string {
	n := len(expression)
	string2add := func(set map[string]bool, str string)map[string]bool{
		if len(str) == 0{return set}
		if len(set) == 0{
			set[str] = true
			return set
		}
		ret := map[string]bool{}
		for k := range set{
			ret[k+str] = true
		}
		return ret
	}
	add := func(set, set1 map[string]bool)map[string]bool{
		if len(set) == 0{
			return set1
		}
		for k := range set1{
			set[k] = true
		}
		return set
	}
	mul := func(set, set1 map[string]bool)map[string]bool{
		if len(set1) == 0{ return set }
		//if len(set) == 0 { set = set1 }
		if len(set) == 0{
			set = add(set, set1)
			return set
		}
		ret := map[string]bool{}
		for k := range set{
			for t := range set1{
				ret[k+t] = true
			}
		}
		return ret
	}
	var dfs func(start, end int)map[string]bool // 去重
	dfs = func(start, end int)(res map[string]bool){
		//fmt.Println(expression[start:end+1])
		// res 表示当前层的集合
		tmp := []byte{} // 缓存当前层当前段处理的字符，注 , 为分割符号
		s := map[string]bool{} // 缓存当前层当前段的集合 注 ,为分割符号
		i := start
		for i <= end{
			switch expression[i]{
			case ',':
				res = add(res, string2add(s, string(tmp)))
				s, tmp = map[string]bool{}, []byte{} //重置
			case '{':
				i++
				t, cnt := i, 1
				for cnt > 0 && i <= end{ // 找 } 注意花括号可嵌套
					if expression[i] == '{' { cnt++ }
					if expression[i] == '}' { cnt-- }
					i++
				}
				s = string2add(s, string(tmp))
				tmp = []byte{}
				s = mul(s, dfs(t, i-2)) // i-2 排除掉 }
				i--
			default:
				tmp = append(tmp, expression[i])
			}
			i++
		}
		res = add(res, string2add(s, string(tmp)))
		return res
	}
	ans := []string{}
	for k := range dfs(0, n-1){
		ans = append(ans, k)
	}
	sort.Strings(ans) // 题目要求排序
	return ans
}
/* 先进行词法分析，然后进行语法分析，然后编写递归下降程序
** 词法分析 通过正则表达式实现
** 语法分析 实现语法对应的函数
 */
/*
gramar
    expr -> item | item ',' expr
    item -> factor | factor item
    factor -> WORD | '{' expr '}'
*/
func BraceExpansionII_2(expression string) []string {
	var s scanner.Scanner
	s.Init(strings.NewReader(expression))
	tok := s.Scan()
	match := func(token rune)string{
		if tok == token{
			val := s.TokenText()
			tok = s.Scan()
			return val
		}else{
			panic("unmatch")
		}
	}
	var expr func() map[string]bool
	var item func()map[string]bool
	var factor func()map[string]bool
	item = func()map[string]bool{
		ret := factor()
		//if tok != scanner.EOF && (tok == scanner.Ident || tok == '{'){
		for tok != scanner.EOF && (tok == scanner.Ident || tok == '{'){
			sufs := factor()
			new := map[string]bool{}
			for pre := range ret{
				for suf := range sufs{
					new[pre+suf] = true
				}
			}
			ret = new
		}
		return ret
	}
	factor = func()map[string]bool{
		ret := map[string]bool{}
		if tok == '{'{
			match('{')
			ret = expr()
			match('}')
			return ret
		}
		ret[match(scanner.Ident)] = true
		return ret
	}
	expr = func() map[string]bool{
		ret := item()
		for tok != scanner.EOF && tok == ','{
			match(',')
			sufs := item()
			for suf := range sufs{
				ret[suf] = true
			}
		}
		return ret
	}
	ret := expr()
	fmt.Println(ret)
	ans := []string{}
	for e := range ret{
		ans = append(ans, e)
	}
	sort.Strings(ans)
	return ans
}
// 此题BFS 思路值得学习
func BraceExpansionII_BFS(expression string) []string {
	q := []string{expression}
	res := map[string]bool{}
	for len(q) > 0{
		exp := q[0]
		q = q[1:]
		// 如果表达式中没有 {，则将这个表达式加入结果中
		if strings.Contains(exp, "{"){
			res[exp] = true
			continue
		}
		// 找到表达式中嵌套最里面的 {}
		left, right := 0, 0 // 花括号开闭位置
		for right = 0; exp[right] != '}'; right++{
			if exp[right] == '{' { left = right }
		}
		fmt.Println(exp[left:right+1])
		before, after := exp[:left], exp[right+1:]
		// 处理嵌套最深的花括号内容， 这是迭代情况下，花括号嵌套的递归处理方式
		strs := strings.Split(exp[left+1:right], ",")

		//BFS 核心思路： 将 before 、 strs 中的每个元素以及 after 拼接成字符串放入到队列中
		for _, s := range strs{
			q = append(q, before+s+after)
		}
	}
	ans := []string{}
	for k := range res{
		ans = append(ans, k)
	}
	sort.Strings(ans)
	return ans
}


/* 22. Generate Parentheses
** Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
[
	"(((())))","((()()))","((())())","((()))()","(()(()))","(()()())","(()())()","(())(())",
 	"(())()()","()((()))","()(()())","()(())()","()()(())","()()()()"
]
[
	"(((())))","((()()))","((())())","((()))()","(()(()))","(()()())","(()())()",
	"(())()()","()((()))","()(()())","()(())()","()()(())","()()()()"
]
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
/* 可以在序列依然保持有效的时候才添加 ( 或者 ), 而不要每次添加
** 通过跟踪 当前为止放置的左括号 和 右括号的数目来做到这一点，即
** 如果左括号数量不大于 n 则可以放一个左括号。
** 若右括号梳理小于左括号数量，可以放一个右括号
** 递归树：https://leetcode-cn.com/problems/generate-parentheses/solution/hui-su-suan-fa-by-liweiwei1419/
** 可以看作 2*n 个空 去填充 ( 和 )
 */
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
/* 方法三：按括号序列的长度递归  动态规划
** 任何一个括号序列都一定是由 ( 开头，并且第一个 ( 一定有一个唯一与之对应的 )
** 这样一来，每一个括号序列可以用 (a)b 来表示，其中 a 与 b 分别是一个合法的括号序列（可以为空）
** 那么，要生成所有长度为 2 * n 的括号序列,定义一个函数 generate(n) 来返回所有可能的括号序列
** 那么在函数 generate(n) 的过程中：
	1. 需要枚举与第一个 ( 对应的 ) 的位置 2 * i + 1
	2. 递归调用 generate(i) 即可计算 a 的所有可能性
	3. 递归调用 generate(n - i - 1) 即可计算 b 的所有可能性
	4. 遍历 a 与 b 的所有可能性并拼接，即可得到所有长度为 2 * n 的括号序列
 */
func GenerateParenthesis_DFS(n int) []string {
	dp := make([][]string, 100)
	var dfs func(cnt int) []string
	dfs = func(cnt int) []string{
		if dp[cnt] != nil{
			return dp[cnt]
		}
		res := []string{}
		if cnt == 0{
			res = append(res, "")
		}else{
			for i := 0; i < cnt; i++{
				for _, left := range dfs(i){
					for _, right := range dfs(cnt-1-i){
						res = append(res, "("+left+")"+right)
					}
				}
			}
		}
		dp[cnt] = res
		return res
	}
	return dfs(n)
}
/* 剩下的括号要么在这一组新增的括号内部，要么在这一组新增括号的外部（右侧）
** 既然知道了 i<n 的情况，那我们就可以对所有情况进行遍历：
	"(" + 「i=p时所有括号的排列组合」 + ")" + 「i=q时所有括号的排列组合」 其中 p + q = n-1，且 p q 均为非负整数。
事实上，当上述 p 从 0 取到 n-1，q 从 n-1 取到 0 后，所有情况就遍历完了。
注：上述遍历是没有重复情况出现的，即当 (p1,q1)≠(p2,q2) 时，按上述方式取的括号组合一定不同。
*/
func GenerateParenthesis_DP(n int) []string {
	dp := make([][]string, n+1)
	dp[0] = []string{""}
	for i := 1; i <= n; i++{
		elem := []string{}
		for p := 0; p <= i; p++{
			q := i - p - 1
			list_p, list_q := dp[p], dp[q]
			for _, pp := range list_p{
				for _, qq := range list_q{
					elem = append(elem, "(" + pp + ")" + qq)
				}
			}
		}
		dp[i] = elem
	}
	return dp[n]
}

/* 301. Remove Invalid Parentheses
** Given a string s that contains parentheses and letters,
** remove the minimum number of invalid parentheses to make the input string valid.
** Return all the possible results. You may return the answer in any order.
** Constraints:
	1 <= s.length <= 25
	s consists of lowercase English letters and parentheses '(' and ')'.
	There will be at most 20 parentheses in s.
 */
/* 1. 如果当前遍历到的「左括号」的数目严格小于「右括号」的数目则表达式无效
** 2. 可以一次遍历计算出多余的「左括号」和「右括号」：统计「左括号」和「右括号」出现的次数
		当遍历到左括号，「左括号」数量加 1
		当遍历到右括号， 如果此时「左括号」的数量不为 0，因为「右括号」可以与之前遍历到的「左括号」匹配，「左括号」出现的次数 -1
		如果此时「左括号」的数量为 0，「右括号」数量加 1
	通过这样的计数规则，得到的「左括号」和「右括号」的数量就是各自最少应该删除的数量
** 当知道最少删除数后，就可以使用枚举搜索
** 首先我们利用括号匹配的规则求出该字符串 s 中最少需要去掉的左括号的数目 lremove 和右括号的数目 rremove，
** 然后我们尝试在原字符串 s 中去掉 lremove 个左括号和 rremove 个右括号，然后检测剩余的字符串是否合法匹配，
** 如果合法匹配则我们则认为该字符串为可能的结果，我们利用回溯算法来尝试搜索所有可能的去除括号的方案

** 优化：剪枝
**	1. 我们从字符串中每去掉一个括号，则更新 lremove 或者 rremove，当我们发现剩余未尝试的字符串的长度小于 lremove+rremove 时，则停止本次搜索
**	2. 当lremove和rremove同时为 0 时，则我们检测当前的字符串是否合法匹配，如果合法匹配则我们将其记录下来
 */
func removeInvalidParentheses(s string) []string {
	ans := []string{}
	isValid := func(str string)bool{
		cnt := 0
		for _, c := range str{
			if c == '('{
				cnt++
			}else if c == ')'{
				cnt--
				if cnt < 0{ return false } // 提前终止
			}
		}
		return cnt == 0
	}
	// start 表示当前str里要扫描的位置
	var dfs func(str string, start, l, r int)
	dfs = func(str string, start, l, r int){
		n := len(str)
		if l == 0 && r == 0{ // 删完了
			if isValid(str){
				ans = append(ans, str)
			}
			return
		}
		//for i := range str{   <==
		for i := start; i < n; i++{
			//if i != 0 && str[i] == str[i-1]{
			if i != start && str[i] == str[i-1]{
				continue
			}
			if l + r > n - i{// 如果剩余的字符无法满足去掉的数量要求，直接返回
				return
			}
			// 尝试去掉一个左括号
			if l > 0 && str[i] == '('{
				dfs(str[:i]+str[i+1:], i, l-1, r)
			}
			// 尝试去掉一个右括号
			if r > 0 && str[i] == ')'{
				dfs(str[:i]+str[i+1:], i, l, r-1)
			}
		}
	}
	lremove, rremove := 0, 0
	for _, c := range s {
		if c == '('{
			lremove++
		}else if c == ')'{
			if lremove == 0{
				rremove++
			}else{
				lremove--
			}
		}
	}
	fmt.Println(lremove, rremove)
	// dfs(s, lremove, rremove) 易错点-1 加入 start
	// start 表示当前str里要扫描的位置
	// 加入 start 产生顺序，从而 控制 重复
	dfs(s, 0, lremove, rremove)
	return ans
}

/* BFS 处理：
** 题目中要求最少删除，这样的描述正是广度优先搜索算法应用的场景，并且题目也要求我们输出所有的结果。
** 我们在进行广度优先搜索时每一轮删除字符串中的 1 个括号，直到出现合法匹配的字符串为止，此时进行轮转的次数即为最少需要删除括号的个数
** 我们进行广度优先搜索时，每次保存上一轮搜索的结果，然后对上一轮已经保存的结果中的每一个字符串尝试所有可能的删除一个括号的方法，
** 然后将保存的结果进行下一轮搜索。在保存结果时，我们可以利用哈希表对上一轮生成的结果去重，从而提高效率
 */

func removeInvalidParentheses_BFS(s string) []string {
	isValid := func(str string)bool{
		cnt := 0
		for _, c := range str{
			if c == '('{
				cnt++
			}
			if c == ')'{
				cnt--
				if cnt < 0{ // 易漏点-1： 务必注意此种情况
					return false
				}
			}
		}
		return cnt==0
	}
	if isValid(s){
		return []string{s}
	}
	tmp := map[string]bool{}
	q := []string{s}
	for len(q) > 0{
		// t := []string{} 利用哈希表对上一轮生成的结果去重，从而提高效率, 测试时间由4.1s 提升到0.55s
		t := map[string]bool{}
		for i := range q{
			for j := range q[i]{
				tt := q[i][:j]+q[i][j+1:]
				if isValid(tt){
					tmp[tt] = true
				}else if len(tmp) == 0{
					t[tt] = true
				}
			}
		}
		if len(tmp) > 0{
			q = nil
		}else{
			for k := range t{
				q = append(q, k)
			}
		}
	}
	ans := []string{}
	for k := range tmp{
		ans = append(ans, k)
	}
	return ans
}

/* 方法三：枚举状态子集
** 首先我们利用括号匹配的规则求出该字符串 s 中最少需要去掉的左括号的数目 lremove 和右括号的数目 rremove，
** 然后我们利用状态子集求出字符串 s 中所有的左括号去掉 lremove 的左括号的子集，和所有的右括号去掉 rremove 个右括号的子集，
** 依次枚举这两种子集的组合，检测组合后的字符串是否合法匹配，如果字符串合法则记录，最后我们利用哈希表对结果进行去重。
 */
func RemoveInvalidParentheses(s string) []string {
	left, right := []int{}, []int{}
	lremove, rremove := 0, 0
	for i, c := range s{
		if c == '('{
			left = append(left, i)
			lremove++
		}
		if c == ')'{
			right = append(right, i)
			if lremove == 0{
				rremove++
			}else{
				lremove--
			}
		}
	}
	isValid := func (lmask, rmask int)bool{
		cnt := 0
		pos1, pos2 := 0, 0
		for i := range s{
			if pos1 < len(left) && i == left[pos1]{
				if lmask>>pos1&1 == 0{  cnt++  }
				pos1++
			}else if pos2 < len(right) && i == right[pos2]{
				if rmask>>pos2&1 == 0{
					cnt--
					if cnt < 0{  return false  }
				}
				pos2++
			}
		}
		return cnt == 0
	}
	recoverStr := func (lmask, rmask int)string{
		res := []rune{}
		pos1, pos2 := 0, 0
		for i, ch := range s {
			if pos1 < len(left) && i == left[pos1] {
				if lmask>>pos1&1 == 0 {
					res = append(res, ch)
				}
				pos1++
			} else if pos2 < len(right) && i == right[pos2] {
				if rmask>>pos2&1 == 0 {
					res = append(res, ch)
				}
				pos2++
			} else {
				res = append(res, ch)
			}
		}
		return string(res)
	}
	/*left 数组记录了左括号的下标位置， right数据记录了右括号的下标位置*/
	var maskArr1, maskArr2 []int
	// maskArr 二进制构造子集个数, 二进制中 1 的个数 表示有多少个括号字符被删除, 为1 表示选中
	for i := 0; i < 1<<len(left); i++{
		if bits.OnesCount(uint(i)) == lremove{
			maskArr1 = append(maskArr1, i)
		}
	}
	for i := 0; i < 1<<len(right); i++{
		if bits.OnesCount(uint(i)) == rremove{
			maskArr2 = append(maskArr2, i)
		}
	}
	res := map[string]struct{}{}
	for _, mask1 := range maskArr1{
		for _, mask2 := range maskArr2{
			if isValid(mask1, mask2){
				res[recoverStr(mask1, mask2)] = struct{}{}
			}
		}
	}
	ans := []string{}
	for str := range res{
		ans = append(ans, str)
	}
	return ans
}

/*241. Different Ways to Add Parentheses
** Given a string expression of numbers and operators,
** return all possible results from computing all the different possible ways to group numbers and operators.
** You may return the answer in any order.
 */
// 2022-01-14 刷出此题
func diffWaysToCompute(expression string) []int {
	nums,ops := []int{}, []byte{}
	t := []byte{}
	for i := range expression{
		c := expression[i]
		if c == '+' || c == '*' || c == '/' || c == '-'{
			ops = append(ops, c)
			num, _ := strconv.Atoi(string(t))
			t = []byte{}
			nums = append(nums, num)
		}else{
			t = append(t, c)
		}
	}
	//易疏漏-1： 下面这2行 忽略了导致 最后一位数字 未添加进 nums 数组
	num, _ := strconv.Atoi(string(t))
	nums = append(nums, num)
	compute := func(f, s int, op byte)int{
		switch op{
		case '+':
			return f + s
		case '-':
			return f - s
		case '*':
			return f * s
		case '/':
			return f / s
		}
		return 0
	}
	var dfs func(start, end int)[]int
	dfs = func(start, end int)[]int{
		if start == end{
			return []int{nums[start]}
		}
		ans := []int{}
		//易疏漏-2：在分组时，最后end 不能包含在里面 否则 又是一个全量 造成 dead loop
		//for i := start; i <= end; i++{
		for i := start; i < end; i++{
			left := dfs(start, i)
			right := dfs(i+1, end)
			for j := range left{
				for k := range right{
					ans = append(ans, compute(left[j], right[k], ops[i]))
				}
			}
		}
		return ans
	}
	return dfs(0, len(nums)-1)
}

/* 224. Basic Calculator
** Given a string s representing a valid expression, implement a basic calculator to evaluate it,
** and return the result of the evaluation.
** Note: You are not allowed to use any built-in function which evaluates strings as mathematical expressions,
** such as eval().
Constraints:
	1 <= s.length <= 3 * 105
	s consists of digits, '+', '-', '(', ')', and ' '.
	s represents a valid expression.
	'+' is not used as a unary operation (i.e., "+1" and "+(2 + 3)" is invalid).
	'-' could be used as a unary operation (i.e., "-1" and "-(2 + 3)" is valid).
	There will be no two consecutive operators in the input.
	Every number and running calculation will fit in a signed 32-bit integer.
 */
/* 1. 本题目只有 "+", "-" 运算，没有 "*" , "/" 运算，因此少了不同运算符优先级的比较
** 2. 有括号
 */
func calculate(s string) int {
	i, n := 0, len(s)
	ans, num, sign := 0, 0, 1
	st := []int{}
	for i < n {
		switch s[i] {
		case '+':
			ans += sign * num
			num = 0
			sign = 1
		case '-':
			ans += sign * num
			num = 0
			sign = -1
		case '(':
			st = append(st, ans)
			st = append(st, sign) // 对符号的处理
			ans = 0
			sign = 1
		case ')':
			ans += sign * num
			num = 0
			ans *= st[len(st)-1] // 处理负号
			st = st[:len(st)-1]
			ans += st[len(st)-1]
			st = st[:len(st)-1]
		case ' ':
			break
		default: // 数字
			num = 10 * num + int(s[i] - '0')
		}
	}
	ans += sign * num
	return ans
}

/* 227. Basic Calculator II
** Given a string s which represents an expression, evaluate this expression and return its value. 
** The integer division should truncate toward zero.
** You may assume that the given expression is always valid.
** All intermediate results will be in the range of [-231, 231 - 1].
** Note: You are not allowed to use any built-in function which evaluates strings as mathematical expressions, such as eval().
** Constraints:
	1 <= s.length <= 3 * 105
	s consists of integers and operators ('+', '-', '*', '/') separated by some number of spaces.
	s represents a valid expression.
	All the integers in the expression are non-negative integers in the range [0, 231 - 1].
	The answer is guaranteed to fit in a 32-bit integer.
 */
// 与上一题比较，此题就需要考虑 运算符优先级， 但是 不需要考虑 负数的问题
// 先计算乘除
func calculateII(s string) int {
	nums := []int{}
	num, sign := 0, 1
	i, op := 0, '+' // 前一次的运算符
	for i < len(s){
		if s[i] == '('{
			j, cnt := i+1, 1
			for cnt != 0{
				if s[j] == '(' { cnt++ }
				if s[j] == ')' { cnt-- }
				j++
			}
			/* 易错点-1 注意 j 的值
			num = sign*calculateII(s[i+1:j])
			i = j+1
			 */
			num = sign*calculate(s[i+1:j-1])
			i = j
			continue
		}
		switch s[i]{
		case '+':
			if op == '*' {
				nums[len(nums)-1] *= num*sign
			}else if op == '/'{
				nums[len(nums)-1] /= num*sign
			}else{
				nums = append(nums, num*sign)
			}
			op = '+' // 遗漏-1： 忘记置位
			num, sign = 0, 1
		case '-':
			if op == '*' {
				nums[len(nums)-1] *= num*sign
			}else if op == '/'{
				nums[len(nums)-1] /= num*sign
			}else{
				nums = append(nums, num*sign)
			}
			op = '-' // 遗漏-1： 忘记置位
			num, sign = 0, -1
		case '*':
			if op == '*' {
				nums[len(nums)-1] *= num*sign
			}else if op == '/'{
				nums[len(nums)-1] /= num*sign
			}else {
				nums = append(nums, num*sign)
			}
			op = '*'
			num, sign = 0, 1
		case '/':
			if op == '*' {
				nums[len(nums)-1] *= num*sign
			}else if op == '/'{
				nums[len(nums)-1] /= num*sign
			}else {
				nums = append(nums, num*sign)
			}
			op = '/'
			num, sign = 0, 1
		case ' ':
			break
		default :
			num = num * 10 + int(s[i]-'0')
		}
		i++
	}
	// 遗漏-2： 最后一位num
	ans := 0
	if op == '*'{
		nums[len(nums)-1] *= num*sign
	}else if op == '/'{
		nums[len(nums)-1] /= num*sign
	}else{
		nums = append(nums, num*sign)
	}
	for _, c := range nums{
		ans += c
	}
	return ans
}

/* 772. Basic Calculator III
** Implement a basic calculator to evaluate a simple expression string.
** The expression string contains only non-negative integers, '+', '-', '*', '/' operators,
** and open '(' and closing parentheses ')'. The integer division should truncate toward zero.
** You may assume that the given expression is always valid.
** All intermediate results will be in the range of [-231, 231 - 1].
 */
func calculateIII(s string) int {
	nums := []int{}
	num, sign := 0, 1
	i, op := 0, '+' // 前一次的运算符
	for i < len(s){
		if s[i] == '('{
			j, cnt := i+1, 1
			for cnt != 0{
				if s[j] == '(' { cnt++ }
				if s[j] == ')' { cnt-- }
				j++
			}
			//num = sign*calculate(s[i+1:j-1])  易错点-2 括号计算后不应该算符号，入栈的值计算符号
			num = calculateIII(s[i+1:j-1])
			i = j
			continue
		}
		switch s[i]{
		case '+':
			if op == '*' {
				nums[len(nums)-1] *= num*sign
			}else if op == '/'{
				nums[len(nums)-1] /= num*sign
			}else{
				nums = append(nums, num*sign)
			}
			op = '+' // 遗漏-1： 忘记置位
			num, sign = 0, 1
		case '-':
			if op == '*' {
				nums[len(nums)-1] *= num*sign
			}else if op == '/'{
				nums[len(nums)-1] /= num*sign
			}else{
				nums = append(nums, num*sign)
			}
			op = '-' // 遗漏-1： 忘记置位
			num, sign = 0, -1
		case '*':
			if op == '*' {
				nums[len(nums)-1] *= num*sign
			}else if op == '/'{
				nums[len(nums)-1] /= num*sign
			}else {
				nums = append(nums, num*sign)
			}
			op, num, sign = '*', 0, 1
		case '/':
			if op == '*' {
				nums[len(nums)-1] *= num*sign
			}else if op == '/'{
				nums[len(nums)-1] /= num*sign
			}else {
				nums = append(nums, num*sign)
			}
			op, num, sign = '/', 0, 1
		case ' ':
			break
		default :
			num = num * 10 + int(s[i]-'0')
		}
		i++
	}
	// 遗漏-2： 最后一位num
	//fmt.Println(nums, sign, num)
	ans := 0
	if op == '*'{
		nums[len(nums)-1] *= num*sign
	}else if op == '/'{
		nums[len(nums)-1] /= num*sign
	}else{
		nums = append(nums, num*sign)
	}
	for _, c := range nums{
		ans += c
	}
	//fmt.Println(s, ans)
	return ans
}
/* 770. Basic Calculator IV
** Given an expression such as expression = "e + 8 - a + 5" and an evaluation map such as {"e": 1}
** (given in terms of evalvars = ["e"] and evalints = [1]),
** return a list of tokens representing the simplified expression, such as ["-1*a","14"]
An expression alternates chunks and symbols, with a space separating each chunk and symbol.
A chunk is either an expression in parentheses, a variable, or a non-negative integer.
A variable is a string of lowercase letters (not including digits.) Note that variables can be multiple letters,
and note that variables never have a leading coefficient or unary operator like "2x" or "-x".
Expressions are evaluated in the usual order: brackets first, then multiplication, then addition and subtraction.
For example, expression = "1 + 2 * 3" has an answer of ["7"].
The format of the output is as follows:
	For each term of free variables with a non-zero coefficient,
		we write the free variables within a term in sorted order lexicographically.
	For example, we would never write a term like "b*a*c", only "a*b*c".
	Terms have degrees equal to the number of free variables being multiplied,
		counting multiplicity. We write the largest degree terms of our answer first,
		breaking ties by lexicographic order ignoring the leading coefficient of the term.
	For example, "a*a*b*c" has degree 4.
	The leading coefficient of the term is placed directly to the left with an asterisk separating it from the variables
		(if they exist.) A leading coefficient of 1 is still printed.
	An example of a well-formatted answer is ["-2*a*a*a", "3*a*a*b", "3*b*b", "4*a", "5*c", "-6"].
	Terms (including constant terms) with coefficient 0 are not included.
	For example, an expression of "0" has an output of [].
 */
func basicCalculatorIV(expression string, evalvars []string, evalints []int) []string {
	evalmap := map[string]int{}
	for i := range evalvars{
		evalmap[evalvars[i]] = evalints[i]
	}
	var dfs func(exp string) []string
	dfs = func(exp string) []string{
		return nil
	}
	return dfs(expression)
}

/* 1190. Reverse Substrings Between Each Pair of Parentheses
** You are given a string s that consists of lower case English letters and brackets.
** Reverse the strings in each pair of matching parentheses, starting from the innermost one.
** Your result should not contain any brackets.
 */
// 错误方法： 遗漏了"t(x)x(y)" 并行有2个括号的情况
func reverseParentheses(s string) string {
	ans := []byte{}
	if !strings.ContainsRune(s, '('){
		return s
	}else{
		// 找最外层
		fst, last := strings.Index(s, "("), strings.LastIndex(s, ")")
		ans = append(ans, s[:fst]...)
		sub := reverseParentheses(s[fst+1:last])
		for i := len(sub)-1; i >= 0; i--{
			ans = append(ans, sub[i])
		}
		ans = append(ans, s[last+1:]...)
	}
	return string(ans)
}
// 此方法会 递归死循环， 缺少跳出条件
func reverseParentheses_DFS(s string) string {
	sb := []byte{}
	n, reverse := len(s), false
	fmt.Println(s)
	/* 还有有问题：(a)(b) 此时脱括号出现问题，还是需要一个方法解决：不在括号不反转的问题
	if s[0] == '(' && s[n-1] == ')'{
		s = s[1:n-1]
		n -= 2
		reverse = true
	}	 */
	for i := 0; i < n;{
		if s[i] == '('{
			reverse = true
			cnt := 1
			j := i+1
			for ; j < n && cnt > 0; j++{
				if s[j] == '('{  cnt++ }
				if s[j] == ')'{ cnt-- }
			}
			if cnt != 0{
				panic("Err: unmatch } ")
			}
			sx := reverseParentheses(string(s[i:j]))
			/* 下面语句i = j 要用到 j， 不可更改j
			for j = range sx{
				sb = append([]byte{sx[j]}, sb...)
			}*/
			for k := range sx{
				sb = append([]byte{sx[k]}, sb...)
			}
			i = j
		}else if s[i] == ')' {
			// 无法确认是否关闭 reverse
		}else{
			if reverse{
				sb = append([]byte{s[i]}, sb...)
			}else{
				sb = append(sb, s[i])
			}
			i++
		}
	}
	return string(sb)
}
// 2022-02-07 刷出此题
// 此题难点： 1. dfs 需要脱前置括号， 递归死循环问题   2. 区别注意，如果没有括号 不需要反转
func ReverseParentheses_DFS(s string) string {
	var dfs func(str string)[]byte
	dfs = func(str string)[]byte{// 前置() 已脱 必定要反转
		n, sa := len(str), []byte{}
		for i := 0; i < n; {
		//for i := range str{
			if str[i] == '('{
				j, cnt := i + 1, 1
				for j < len(str) && cnt > 0{
					if str[j] == '('	{	cnt++ }
					if str[j] == ')'	{	cnt-- }
					j++
				}
				if cnt != 0{
					panic("Err: unmatch ()")
				}
				sx := dfs(str[i+1:j-1])
				for k := range sx{
					sa = append([]byte{sx[k]}, sa...)
				}
				i = j
			}else{
				sa = append([]byte{str[i]}, sa...)
				i++
			}
		}
		return sa
	}
	n, sb := len(s), []byte{}
	for i := 0; i < n;{
		if s[i] == '('{
			j, cnt := i + 1, 1
			for j < len(s) && cnt > 0{
				if s[j] == '('	{	cnt++ }
				if s[j] == ')'	{	cnt-- }
				j++
			}
			if cnt != 0{
				panic("Err: unmatch ()")
			}
			sx := dfs(s[i+1:j-1])
			sb = append(sb, sx...)
			i = j
		}else{
			sb = append(sb, s[i])
			i++
		}
	}
	return string(sb)
}

func reverseParentheses_stack(s string) string {
	st := []byte{}
	i, n, ans := 0, len(s), []byte{}
	for i := 0; i < n; i++{
		if len(st) == 0{
			if s[i] == '('{
				st = append(st, '(')
			}else{
				ans = append(ans, s[i])
			}
		}else{
			if s[i] == ')'{

			}else{
				st = append(st, s[i])
			}
		}
	}
	return string(ans)
}
// 官方题解-1
/* 从左到右遍历该字符串, 使用字符串 str 记录当前层所遍历到的小写英文字母.
** 对于当前遍历的字符：
** 1. 左括号，将 str 插入栈中，并将 str 置空，进入下一层
** 2. 右括号, 则说明遍历完了当前层, 反转 str, 返回给上一层.
	  即将栈顶字符串弹出，然后将反转后的 str 拼接到 栈顶字符串末尾，然后将结果赋值给 str
** 3. 如果是小写英文字母，将其加到 str 末尾
 */
func ReverseParentheses_stack(s string) string {
	st := [][]byte{} // 二维栈
	str := []byte{}
	for i := range s {
		if s[i] == '('{
			st = append(st, str)
			str = []byte{} // 易错点-1 清空
		}else if s[i] == ')'{// 主处理
			for j, n := 0, len(str); j < n/2; j++{ // 反转算法- 学习-1
				str[j], str[n-1-j] = str[n-1-j], str[j]
			}
			str = append(st[len(st)-1], str...) // 合并
			st = st[:len(st)-1]
		}else{
			str = append(str, s[i])
		}
	}
	return string(str)
}
/* 官方最优题解：预处理括号
** 可以将括号的反转理解为逆序地遍历括号
** 沿着某个方向移动，此时遇到了括号, 那么我们只需要首先跳跃到该括号对应的另一个括号所在处，
** 然后改变我们的移动方向即可.这个方案同时适用于遍历时进入更深一层，以及完成当前层的遍历后返回到上一层的方案
** 实际代码中，我们需要预处理出每一个括号对应的另一个括号所在的位置，这一部分我们可以使用栈解决。
** 当我们预处理完成后，即可在线性时间内完成遍历，遍历的字符串顺序即为反转后的字符串
 */
func ReverseParentheses(s string) string {
	n := len(s)
	pair := make([]int, n)
	st := []int{}
	// 预处理，记录括号位置
	for i := range s{
		if s[i] == '('{
			st = append(st, i) // 记录位置
		}else if s[i] == ')'{
			j := st[len(st)-1]
			st = st[:len(st)-1]
			pair[i], pair[j] = j, i
		}
	}
	// 反转操作：反转理解为逆序地遍历
	ans := []byte{}
	for i, step := 0, 1; i < n; i += step{
		if s[i] == '(' || s[i] == ')'{
			i = pair[i]
			step = -step // 逆序遍历控制
		}else{
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}





