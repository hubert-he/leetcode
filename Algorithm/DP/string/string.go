package string

/* 392. Is Subsequence
** Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
** A subsequence of a string is a new string
** that is formed from the original string by deleting some (can be none) of the characters
** without disturbing the relative positions of the remaining characters.
** (i.e., "ace" is a subsequence of "abcde" while "aec" is not).
** Follow up: Suppose there are lots of incoming s, say s1, s2, ..., sk where k >= 109,
** 		and you want to check one by one to see if t has its subsequence.
** 		In this scenario, how would you change your code?
 */
/* 此题可以用双指针，大量操作用于在t 中找下一个匹配的字符
** 后续进阶 DP
** 预处理出对于t 的每一个位置，从该位置开始往后每一个字符第一次出现的位置
** 动态规划实现预处理
** dp[i][j] 表示字符串 t 从 i 开始往后字符 j 第一次出现的位置
** dp[i][j] = i  			若 t[i] == j
**          = dp[i+1][j] 	若 t[i] != j
** 每次 O(1) 找出 t 中下一个位置
 */
func isSubsequence(s string, t string) bool {
	ns, nt := len(s), len(t)
	if nt == 0{
		if ns == 0{
			return true
		}
		return false
	}
	//dp[i][j] 表示字符串 t 从 i 开始往后字符 j 第一次出现的位置
	dp := make([][26]int, nt)
	for i := 0; i < 26; i++{
		dp[nt-1][i] = nt
	}
	dp[nt-1][t[nt-1]-'a'] = nt-1
	for i := nt-2; i >= 0; i--{
		for j := 0; j < 26; j++{
			c := int(t[i] - 'a')
			if c == j{
				dp[i][j] = i
			}else{
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	j := 0
	for i := range s{
		loc := int(s[i] - 'a')
		if j >= nt || dp[j][loc] == nt{
			return false
		}
		j = dp[j][loc] + 1
	}
	return true
}
// 针对特殊情况的处理
func isSubsequence2(s string, t string) bool {
	nt := len(t)
	dp := make([][26]int, nt+1) // 额外申请一个，方便特视情况处理
	for i := range dp[nt]{ // 设置保护边界
		dp[nt][i] = nt // 表示不存在
	}
	for i := nt-1; i >= 0; i--{
		for j := range dp[i]{
			c := int(t[i] - 'a')
			if c == j{
				dp[i][j] = i
			}else{
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	j := 0
	for i := range s{
		c := int(s[i] - 'a')
		if dp[j][c] == nt{
			return false
		}
		j = dp[j][c] + 1
	}
	return true
}
