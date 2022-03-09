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
// 2022-03-01 刷出此题  但是代码方式仍然不如官方解答
func checkInclusion2(s1 string, s2 string) bool {
	n1, n := len(s1), len(s2)
	if n1 > n{ return false }
	m := [26]byte{}
	cnt := 0
	for _, c := range s1{
		m[c-'a']++
		cnt++
	}
	isPerm := func(x string)bool{
		t := [26]byte{}
		for _, c := range x {
			t[c-'a']++
		}
		return t == m
	}
	i, j := 0, 0
	for i < n{
		if m[s2[i]-'a'] == 0{
			i++
			j = i
			continue
		}
		i++
		if i - j == cnt {
			if isPerm(s2[j:i]){
				return true
			}else{
				j++
			}
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
/* 优化：每次窗口滑动时，只统计了一进一出两个字符，却比较了整个cnt1和cnt2数组
** 从这个角度出发，我们可以用一个变量diff 来记录 cnt1 和 cnt2 的不同值的个数，就这样判断cnt1 和 cnt2是否相等转换为 diff 是否为 0
** 每次窗口滑动，记一进一出两个字符为 x 和 y
** 若 x == y 则对cnt2无影响，直接跳过
** 若 x != y 则对于字符x， 在修改 cnt2之前若有 cnt2[x] == cnt1[x], 则将diff 加 1；在修改cnt2之后若有cnt2[x]==cnt1[x], 则diff 减 1
** 	字符y同理
** 可以只用 一个数组 cnt， 其中 cnt[x] = cnt2[x] - cnt1[x], 将 cnt1[x] 与 cnt2[x]的比较替换成 cnt[x] 与 0 的比较
*/
func CheckInclusionOp(s1, s2 string) bool {
	n1, n := len(s1), len(s2)
	if n1 > n{ return false }
	diff := 0
	cnt := [26]int{}
	for i := 0; i < n1; i++{// 区分正负值
		cnt[s1[i]-'a']--
		cnt[s2[i]-'a']++
	}
	for _, c := range cnt{
		if c != 0{
			diff++
		}
	}
	if diff == 0{ return true }
	for i := n1; i < n; i++{
		in, out := s2[i]-'a', s2[i-n1]-'a'
		if in == out{ continue }
		if cnt[in] == 0{ diff++ } //
		cnt[in]++
		if cnt[in] == 0{ diff-- } //
		if cnt[out] == 0 { diff++ }
		cnt[out]--
		if cnt[out] == 0 { diff-- }
		if diff == 0{ return true }
	}
	return false
}
/* 引申出双指针解法
** 方法二中 在保证区间长度为 n 的情况下，去考察是否存在一个区间使得 cnt 的值全为 0
** 反过来，还可以在保证 cnt 的值不为正的情况下，去考察是否存在一个区间，其长度恰好为 n
** 初始时，仅统计 s1 中的字符，则 cnt 的值均不为正，且元素值之和为 -n。
** 然后用两个指针 out 和 in 表示考察的区间 [out, in]。 in 每向右移动一次，就统计一次进入区间的字符x
** 为保证cnt的值不为正，若此时cnt[x] > 0, 则向右移动 out， 即减少离开区间的字符的cnt值直到cnt[x] <= 0
** [out, in] 的长度每增加 1， cnt的元素值之和 就会增加 1， 当[out, in] 恰好为 n 时候， 就意味着 cnt 元素全为 0
** 注意 cnt 的元素 保证 不为正，如果 元素之和 为 0 就只有一种可能，即 cnt 全为 0，这就找到了一个目标子串。
*/
func CheckInclusion2Pointer(s1, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 > n2{
		return false
	}
	cnt := [26]int{}
	for _, c := range s1{
		cnt[c-'a']--
	}
	out := 0
	for in := range s2{
		cnt[s2[in]-'a']++
		for cnt[s2[in]-'a'] > 0{
			cnt[s2[out]-'a']--
			out++
		}
		if in-out+1 == n1{
			return true
		}
	}
	return false
}














