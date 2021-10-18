package slideWindow
func max(nums ...int)int{
	m := nums[0]
	for i := range nums{
		if m < nums[i]{
			m = nums[i]
		}
	}
	return m
}
/* 滑动窗口算法（Sliding Window Algorithm）
** Sliding window algorithm is used to perform required operation on specific window size of given large buffer or array.
** This technique shows how a nested for loop in few problems can be converted to single for loop and hence reducing the time complexity.
** 滑动窗口算法是在给定特定窗口大小的数组或字符串上执行要求的操作。该技术可以将一部分问题中的嵌套循环转变为一个单循环，因此它可以减少时间复杂度。
** 滑动窗口算法在一个特定大小的字符串或数组上进行操作，而不在整个字符串和数组上操作，这样就降低了问题的复杂度，从而也达到降低了循环的嵌套深度。
** 可以看出来滑动窗口主要应用在数组和字符串上
 */
/* 3. Longest Substring Without Repeating Characters
** Given a string s, find the length of the longest substring without repeating characters.
 */
func LengthOfLongestSubstring(s string) int {
	n := len(s)
	//m := map[byte]bool{} // 优化
	m := make([]bool, 128) // 替换到map，底层红黑树，需要logn时间
	left, right := 0, 0
	ans := 0
	for left < n{
		for right < n && !m[s[right]]{
			m[s[right]] = true
			right++
		}
		//delete(m, s[left-1])
		m[s[left]] = true
		ans = max(ans, right-left)
		left++
	}
	return ans
}

/* 567. Permutation in String
** Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.
** In other words, return true if one of s1's permutations is the substring of s2.
 */
/* 方法一：ACCEPT 模拟的方式
** 由于排列不会改变字符串中每个字符的个数，所以只有当两个字符串每个字符的个数均相等时，一个字符串才是另一个字符串的排列
** O(S1*S2)
 */
func checkInclusion(s1 string, s2 string) bool {
	n, n2 := len(s1), len(s2)
	m := make([]int, 26)
	equal := func(x, y []int)bool{
		nx, ny := len(x), len(y)
		if nx != ny{
			return false
		}
		for i := 0; i < nx; i++{
			if x[i] != y[i]{
				return false
			}
		}
		return true
	}
	for i := range s1{
		m[s1[i]-'a']++
	}
	for i := 0; i < n2; i++{
		t := make([]int, 26)
		for j := i; j < n2 && j-i < n; j++{
			t[s2[j]-'a']++
		}
		if equal(m, t){
			return true
		}
	}
	return false
}
/*
** 通过方法一可以发现 可以使用固定滑动窗口来比较
 */
func CheckInclusion(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 > n2{
		return false
	}
	cnt1, cnt2 := [26]int{}, [26]int{} // 用数组 方便使用golang 比较语法
	for i, c := range s1{
		cnt1[c-'a']++
		cnt2[s2[i]-'a']++
	}
	if cnt1 == cnt2{
		return true
	}
	// 固定滑动窗口
	for i := n1; i < n2; i++{
		cnt2[s2[i]-'a']++
		cnt2[s2[i-n1]-'a']--
		if cnt1 == cnt2 {
			return true
		}
	}
	return false
}