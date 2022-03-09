package string

import (
	"fmt"
	"strconv"
	"strings"
)

/* 415. Add Strings
Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 as a string.
You must solve the problem without using any built-in library for handling large integers (such as BigInteger).
You must also not convert the inputs to integers directly.
Example 1:
	Input: num1 = "11", num2 = "123"
	Output: "134"
Example 2:
	Input: num1 = "456", num2 = "77"
	Output: "533"
Example 3:
	Input: num1 = "0", num2 = "0"
	Output: "0"
 */
func addStrings(num1 string, num2 string) string {
	carry := 0
	ans := ""
	for i,j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0|| carry > 0; i,j = i-1,j-1{
		x,y := 0,0
		if i >= 0{
			x = int(num1[i] - '0')
		}
		if j >= 0{
			y = int(num2[j] - '0')
		}
		result := x + y + carry
		ans = strconv.Itoa(result%10) + ans
		carry = result / 10
	}
	return ans
}
func addStrings2(num1 string, num2 string) string {
	var carry byte
	ans := []byte{}
	for i,j :=len(num1)-1,len(num2)-1; i>=0 || j>=0 || carry>0; i,j=i-1,j-1{
		var x,y byte
		if i >= 0{
			x = num1[i] - '0'
		}
		if j >=0 {
			y = num2[j] - '0'
		}
		sum := x + y + carry
		ans = append([]byte{(sum%10)+'0'}, ans...)
		carry = sum / 10
	}
	return string(ans)
}
/* 387. First Unique Character in a String
	Given a string s, find the first non-repeating character in it and return its index.
	If it does not exist, return -1.
Example 1:
	Input: s = "leetcode"
	Output: 0
Example 2:
	Input: s = "loveleetcode"
	Output: 2
Example 3:
	Input: s = "aabb"
	Output: -1
可以通过map 统计词频计算， 但不是最优的， 因为会遍历2次
下面通过队列，遍历一遍
 */
func firstUniqChar(s string) int {
	m := map[byte]int{}
	type pair struct{
		ch byte
		pos int
	}
	q := []pair{}
	for i := range s{
		if _, ok := m[s[i]]; ok {
			m[s[i]] = -1
			for len(q) > 0 && m[q[0].ch] == -1{
				q = q[1:]
			}
		}else{
			m[s[i]] = i
			q = append(q, pair{s[i], i})
		}
	}
	if len(q) > 0{
		return q[0].pos
	}
	return -1
}
// 官方版本
func firstUniqChar2(s string) int {
	m := [26]int{}
	length := len(s)
	for i := range m{
		m[i] = length
	}
	type pair struct {
		ch  byte
		pos int
	}
	q := []pair{} // 记录
	for i := range s {
		v := s[i] - 'a' // 遗漏
		if m[v] == length{
			m[v] = i
			q = append(q, pair{v, i})
		}else{
			m[v] = length + 1
			/* 延迟删除:即使队列中有一些字符出现了超过一次，但它只要不位于队首，那么就不会对答案造成影响，
				我们也就可以不用去删除它。只有当它前面的所有字符被移出队列，它成为队首时，我们才需要将它移除
			 */
			for len(q) > 0 && m[q[0].ch] == length + 1{
				q = q[1:]
			}
		}
	}
	if len(q) > 0 {
		return q[0].pos
	}
	return -1
}
// 2021-11-08 重刷此题
// 队列具有「先进先出」的性质，因此很适合用来找出第一个满足某个条件的元素
func FirstUniqChar3(s string) int {
	m := [26]int{}
	q := []int{}
	for i := range s{
		c := s[i] - 'a'
		m[c]++
		q = append(q, i)
		for len(q) > 0 {
			x := s[q[0]] - 'a'
			if m[x] > 1{
				q = q[1:]
			}else{
				break
			}
		}
	}
	if len(q) > 0{
		return q[0]
	}
	return -1
}

/* 389. Find the Difference
You are given two strings s and t.
String t is generated by random shuffling string s and then add one more letter at a random position.
Return the letter that was added to t.
 */
// 2022-02-17 刷出此题， 需要学习此题3中方法思想
// 计数
func FindTheDifference(s string, t string) byte {
	m := [26]int{}
	for i := range t{
		v := t[i] - 'a'
		m[v]++
	}
	for i := range s{
		v := s[i] - 'a'
		m[v]--
	}
	for i := range m{
		if m[i] > 0{
			return byte(i) + 'a'
		}
	}
	return 0
}

// 求和
func FindTheDifference2(s string, t string) byte {
	ans := 0
	n := len(t)
	for i := 0; i < n; i++{
		ans += int(t[i])
		if i != n-1{
			ans -= int(s[i])
		}
	}
	return byte(ans)
}
/* bit
如果将两个字符串拼接成一个字符串，则问题转换成求字符串中出现奇数次的字符
与 136. Single Number 雷同
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
You must implement a solution with a linear runtime complexity and use only constant extra space.
*/
func FindTheDifference3(s string, t string) byte {
	var diff byte
	for i := range s{
		diff ^= s[i] ^ t[i]
	}
	return diff ^ t[len(t) - 1]
}
// 位运算： 两个相同的字节 异或 为 0  0 异或任何值 均是原值
func findTheDifference4(s string, t string) byte {
	var ans byte
	n := len(t)
	for i := 0; i < n; i++{
		ans ^= t[i]
		if i != n-1{
			ans ^= s[i]
		}
	}
	return ans
}
func singleNumber(nums []int) int {
	n := len(nums)
	ans := nums[0]
	for i := 1; i < n; i++{
		ans ^= nums[i]
	}
	return ans
}
/* 459. Repeated Substring Pattern
  Given a string s, check if it can be constructed by taking a substring of it and appending multiple copies of the substring together.
 */
func RepeatedSubstringPattern(s string) bool {
	n := len(s)
	if n <= 1{
		return false
	}
	/* 参考测试用例， 与奇偶性无关
	if n%2 != 0{ // 奇数个
		t := s[0]
		for i := range s{
			if s[i] != t{
				return false
			}
		}
		return true
	}else{*/
		for i := 0; i < n/2; i++{
			for j := i; j < n/2; j++{
				subStr := s[i:j+1]
				tmp := ""
				//fmt.Println(i, j, n, subStr)
				for k := 0; k < n/(j-i+1); k++{
					tmp += subStr
				}
				if mathNative(tmp, s){
					return true
				}
			}
		}
		return false
}
func mathNative(s1, s2 string)bool{
	n1, n2 := len(s1), len(s2)
	if n1 != n2{
		return false
	}
	for i := range s1{
		if s1[i] != s2[i]{
			return false
		}
	}
	return true
}
/*
	如果长度为 n 的字符串 s 可以由它的一个长度为n1的子串s1 重复多次构成：
    <1> n 一定是 n1 的倍数
	<2> s1 一定是 s 的前缀
	<3> 对于任意 i i属于[n1,n)， 有 s[i] = s[i-n1]
	s中长度为n1的前缀就是s1 并且在这之后的每一个位置上的字符 s[i] 都需要与它之前的第 n1个字符s[i-n1]相同
	子串至少需要重复一次，所以 n1 的大小不会大于n的一半，只需要枚举 [1, n/2]内即可
 */
func RepeatedSubstringPattern2(s string) bool {
	n := len(s)
	for i := 1; i*2 <= n; i++{ // 这里 i表示可能的子串长度
		if n % i == 0{ // n 一定是 n1 的倍数
			match := true
			for j := i; j < n; j++{
				if s[j] != s[j-i]{
					match = false
					break
				}
			}
			if match{
				return true
			}
		}
	}
	return false
}
func RepeatedSubstringPatternTest(s string)bool{
	n := len(s)
	for i := 1; i < n; i++{
		if n % i != 0{
			continue
		}
		sub := s[:i]
		repeat := n / i
		if strings.Index(s, strings.Repeat(sub, repeat)) == 0{
			return true
		}
	}
	return false
}
/*1668. Maximum Repeating Substring
For a string sequence, a string word is k-repeating if word concatenated k times is a substring of sequence.
The word's maximum k-repeating value is the highest value k where word is k-repeating in sequence.
If word is not a substring of sequence, word's maximum k-repeating value is 0.
Given strings sequence and word, return the maximum k-repeating value of word in sequence.
Example 1:
	Input: sequence = "ababc", word = "ab"
	Output: 2
	Explanation: "abab" is a substring in "ababc".
Example 2:
	Input: sequence = "ababc", word = "ba"
	Output: 1
	Explanation: "ba" is a substring in "ababc". "baba" is not a substring in "ababc".
Example 3:
	Input: sequence = "ababc", word = "ac"
	Output: 0
	Explanation: "ac" is not a substring in "ababc".
 */
// 注意关注二分查找中： mid 的计算
// mid := i + (j - i) / 2  以及  mid := (i+j) / 2
func MaxRepeating(sequence string, word string) int {
	i, j := 0, len(sequence)/len(word)
	ans := 0
	for i <= j{
		mid := i + (j - i) / 2
		//mid := (i+j) / 2
		fmt.Println(mid)
		if strings.Index(sequence, strings.Repeat(word, mid)) == -1{
			j = mid - 1
		}else{
			ans = mid
			i = mid + 1
		}
	}
	ans = i-1
	return ans
}

/* 17.17. Multi Search LCCI
Given a string band an array of smaller strings T, design a method to search b for each small string in T.
Output positions of all strings in smalls that appear in big, where positions[i] is all positions of smalls[i].
Example:
	Input:
	big = "mississippi"
	smalls = ["is","ppi","hi","sis","i","ssippi"]
	Output:  [[1,4],[8],[],[3],[1,4,7,10],[5]]
 */
func MultiSearch(big string, smalls []string) [][]int {
	sn := len(smalls)
	//bn := len(big)
	ans := make([][]int, sn)
	for i := range smalls{
		/* strings.Index 特殊点
		   subString = "" 空字符串，Index 返回 0
		   与subString = 首字符， Index 也返回 0
		*/
		if len(smalls[i]) == 0{
			continue
		}
		s := big
		pre := 0
		p := strings.Index(s, smalls[i])
		for p != -1{
			s = s[p+1:]
			ans[i] = append(ans[i], p+pre)
			pre += p+1
			p = strings.Index(s, smalls[i])
		}
	}
	return ans
}