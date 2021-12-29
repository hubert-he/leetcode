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
/* 优化：每次窗口滑动时，只统计了一进一出两个字符，却比较了整个cnt1和cnt2数组
** 从这个角度出发，我们可以用一个变量diff 来记录 cnt1 和 cnt2 的不同值的个数，就这样判断cnt1 和 cnt2是否相等转换为 diff 是否为 0
** 每次窗口滑动，记一进一出两个字符为 x 和 y
** 若 x == y 则对cnt2无影响，直接跳过
** 若 x != y 则对于字符x， 在修改 cnt2之前若有 cnt2[x] == cnt1[x], 则将diff 加 1；在修改cnt2之后若有cnt2[x]==cnt1[x], 则diff 减 1
** 	字符y同理
** 可以只用 一个数组 cnt， 其中 cnt[x] = cnt2[x] - cnt1[x], 将 cnt1[x] 与 cnt2[x]的比较替换成 cnt[x] 与 0 的比较
*/
func CheckInclusionOp(s1, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 > n2{
		return false
	}
	cnt := [26]int{}
	for i, c := range s1{
		cnt[c-'a']--
		cnt[s2[i]-'a']++
	}
	diff := 0
	for _, c := range cnt{
		if c != 0{
			diff++
		}
	}
	if diff == 0{
		return true
	}
	for i := n1; i < n2; i++{
		x, y := s2[i]-'a', s2[i-n1]-'a'
		if x == y{
			continue
		}
		if cnt[x] == 0{
			diff++
		}
		cnt[x]++
		if cnt[x] == 0{
			diff--
		}
		if cnt[y] == 0{
			diff++
		}
		cnt[y]--
		if cnt[y] == 0{
			diff--
		}
		if diff == 0{
			return true
		}
	}
	return false
}
/* 引申出双指针解法
** 方法二中 在保证区间长度为 n 的情况下，去考察是否存在一个区间使得 cnt 的值全为 0
** 反过来，还可以在保证 cnt 的值不为正的情况下，去考察是否存在一个区间，其长度恰好为 n
** 初始时，仅统计 s1 中的字符，则 cnt 的值均不为正，且元素值之和为 -n。
** 然后用两个指针left 和 right 表示考察的区间 [left, right]。 right每向右移动一次，就统计一次进入区间的字符x
** 为保证cnt的值不为正，若此时cnt[x] > 0, 则向右移动左指针，减少离开区间的字符的cnt值直到cnt[x] <= 0
**  难理解！！
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
	left := 0
	for right := range s2{
		ch := s2[right] - 'a'
		cnt[ch]++
		for cnt[ch] > 0{
			cnt[s2[left]-'a']--
			left++
		}
		if right-left+1 == n1{
			return true
		}
	}
	return false
}














