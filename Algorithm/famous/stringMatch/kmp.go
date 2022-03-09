package stringMatch

import "strings"

/* 记录在前面
** 字符串匹配问题的形式定义：
** 1. 文本（Text）是一个长度为 n 的数组 T[1..n]；
** 2. 模式（Pattern）是一个长度为 m 且 m ≤ n 的数组 P[1..m]；
** 3. T 和 P 中的元素都属于有限的字母表 Σ 表
** 4. 如果 0 ≤ s ≤ n-m，并且 T[s+1..s+m] = P[1..m]，即对 1 ≤ j ≤ m，有 T[s+j] = P[j]，
		则说模式 P 在文本 T 中出现且位移为 s，且称 s 是一个有效位移（Valid Shift）。
** 解决字符串匹配的算法包括：：
	朴素算法（Naive Algorithm）
	Rabin-Karp 算法
	有限自动机算法（Finite Automation）
	Knuth-Morris-Pratt 算法（即 KMP Algorithm）
	Boyer-Moore 算法
	Simon 算法、Colussi 算法、Galil-Giancarlo 算法、Apostolico-Crochemore 算法、Horspool 算法和 Sunday 算法等

** 字符串匹配算法通常分为两个步骤：预处理（Preprocessing）和匹配（Matching）。
** 所以算法的总运行时间为预处理和匹配的时间的总和。
 */
/* 朴素的字符串匹配-Brute Force Algorithm
	1. 没有预处理阶段；
	2. 滑动窗口总是后移 1 位；
	3. 对模式中的字符的比较顺序不限定，可以从前到后，也可以从后到前
	4. 匹配阶段需要 O((n - m + 1)m) 的时间复杂度
	5. 需要 2n 次的字符比较
**
 */
/* 在模拟 KMP 匹配过程之前，我们先建立两个概念：
	前缀：对于字符串 abcxxxxefg，我们称 abc 属于 abcxxxxefg 的某个前缀。
	后缀：对于字符串 abcxxxxefg，我们称 efg 属于 abcxxxxefg 的某个后缀。
 */

/* 28. Implement strStr()
** Implement strStr().
** Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
** Clarification:声明
	What should we return when needle is an empty string? This is a great question to ask during an interview.
	For the purpose of this problem, we will return 0 when needle is an empty string.
	This is consistent to C's strstr() and Java's indexOf().
 */
// 2022-03-07 刷出此题，平方级方法
func strStr(haystack string, needle string) int {
	ans := 0
	i, n := 0, len(haystack)
	j := 0
	term := n - len(needle)
	for i <= term {
		idx := i
		for j < len(needle) && haystack[idx] == needle[j]{
			idx, j = idx+1, j+1
		}
		if j >= len(needle){
			return i
		}
		i, j = i+1, 0
	}
	if i > term{ return -1 }
	return ans
}
// 学习 KMP 算法
func strStr_kmp(haystack string, needle string) int {

}

/* 459. Repeated Substring Pattern
** Given a string s,
** check if it can be constructed by taking a substring of it and appending multiple copies of the substring together.
 */
// 2022-03-07 刷出此题
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	for i := 1; i <= n/2; i++{
		if strings.Repeat(s[:i], n/i) == s {
			return true
		}
	}
	return false
}
// 学习 KMP 算法
func repeatedSubstringPattern_KMP(s string) bool {

}