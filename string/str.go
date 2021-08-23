package string

import (
	"strconv"
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

/* 389. Find the Difference
You are given two strings s and t.
String t is generated by random shuffling string s and then add one more letter at a random position.
Return the letter that was added to t.
 */
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
	sum := 0
	for i := range t{
		sum += int(s[i])
	}
	for i := range s{
		sum -= int(s[i])
	}
	return byte(sum)
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
	for i := 1; i*2 <= n; i++{
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