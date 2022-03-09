package slideWindow

/* 76. Minimum Window Substring
** Given two strings s and t of lengths m and n respectively,
** return the minimum window substring of s such that every character in t (including duplicates) is included in the window.
** If there is no such substring, return the empty string "".
** The testcases will be generated such that the answer is unique.
** A substring is a contiguous sequence of characters within the string.
 */
// 2021-12-15 刷出此题
// 如何判断当前的窗口包含所有 t 所需的字符呢？
// 我们可以用一个哈希表表示 t 中所有的字符以及它们的个数，用一个哈希表动态维护窗口中所有的字符以及它们的个数，
// 如果这个动态表中包含 t 的哈希表中的所有字符，并且对应的个数都不小于 t 的哈希表中各个字符的个数，那么当前的窗口是「可行」的
// 考虑如何优化？ 如果 s = XX⋯XABCXXXX，t = ABC，
// 那么显然 [XX⋯XABC] 是第一个得到的「可行」区间，得到这个可行区间后，我们按照「收缩」窗口的原则更新左边界，得到最小区间。
// 我们其实做了一些无用的操作，就是更新右边界的时候「延伸」进了很多无用的 X，更新左边界的时候「收缩」扔掉了这些无用的 X，
// 做了这么多无用的操作，只是为了得到短短的 ABC。
// 没错，其实在 s 中，有的字符我们是不关心的，我们只关心 t 中出现的字符，我们可不可以先预处理 s，扔掉那些 t 中没有出现的字符，然后再做滑动窗口呢？
// 也许你会说，这样可能出现 XXABXXC 的情况，在统计长度的时候可以扔掉前两个 X，但是不扔掉中间的 X，怎样解决这个问题呢？优化后的时空复杂度又是多少
// 预处理，记录下标位置
func MinWindow(s string, t string) string {
	ns, nt := len(s), len(t)
	if ns < nt { // 易漏点-1
		return ""
	}
	m := make([]int, 128)
	mt := map[byte]int{}
	for i := range t{
		mt[t[i]]++
	}
	valid := func()bool{ // 验证方法
		for k := range mt{
			if m[k] < mt[k]{
				return false
			}
		}
		return true
	}
	i, j := 0, 0
	mins, mine := 0, ns
	for j < ns{
		m[s[j]]++
		pre := i
		for valid(){
			m[s[i]]--
			i++
		}
		j++
		if pre != i{
			i = i-1
			m[s[i]]++
			if valid() && mine - mins > j - i{ // 易漏点-2： 是否valid
				mins, mine = i, j
			}
		}
	}
	if valid(){ // 易漏点-3
		return s[mins:mine]
	}
	return ""
}


/* 239. Sliding Window Maximum
** You are given an array of integers nums,
** there is a sliding window of size k which is moving from the very left of the array to the very right.
** You can only see the k numbers in the window. Each time the sliding window moves right by one position.
** Return the max sliding window.
 */
// 分块+预处理   +++++  稀疏表
/* 除了「随着窗口的移动实时维护最大值」这种思路外，考虑其他思路
** 可以将数组 nums 从左到右按照 k 个一组进行分组，最后一组中元素的数量可能会不足 k 个
** 如果我们希望求出nums[i] 到 nums[i+k-1]的最大值，就会出现2种情况：
** 1. 如果 i 是 k 的倍数，那么nums[i] 到 nums[i+k-1]恰好是一个分组，我们只要预处理出每个分组中的最大值，即可得到答案
** 2. 如果 i 不是 k 的倍数，那么nums[i] 到 nums[i+k-1]会跨越2个分组，占有第一个分组的后缀以及第二个分组的前缀
**	假设 j 是 k 的倍数，并且满足i < j <= i+k-1 那么nums[i] 到 nums[j-1]就是第一个分组的后缀， nums[j]到nums[i+k-1]就是第二个分组的前缀
**	如果我们能够预处理出每个分组中的前缀最大值以及后缀最大值，同样可以在 O(1) 的时间得到答案
** 使用2个数组：
**	prefixMax[i] 表示下标 i 对应的分组中，以 i 结尾的前缀最大值
** 	suffixMax[i] 表示下标 i 对应的分组中，以i开始的后缀最大值
** 递推关系：
**	prefixMax[i] = nums[i]  若 i 是 k的倍数
** 	prefixMax[i] = max{ prefixMax[i-1], nums[i] } 若 i 不是 k 的倍数
** 边界条件：prefixMax[0] = nums[0] 恰好包含在递推式的第一种情况中，因此无需特殊考虑

**	suffixMax[i] = nums[i]  若 i+1 是 k 的倍数
** 	suffixMax[i] = max{ suffixMax[i+1], nums[i] } 若 i+1 不是 k 的倍数
** 边界条件： 递推suffixMax[i]需要考虑边界条件 suffixMax[n-1] = nums[n-1]

** 预处理完成后， 对于 nums[i] 到 nums[i+k-1] 的所有元素，
** 1. 如果 i 不是 k 的倍数，那么窗口中的最大值为： max{ suffixMax[i], prefixMax[i+k-1] }
** 2. 如果i是k 的倍数，那么此时窗口恰好对应一整个分组，suffixMax[i] 与 prefixMax[i+k-1] 都等于分组中的最大值
** 总和1 和 2 两种情况，得出 结果值为 max{ suffixMax[i], prefixMax[i+k-1] }
 */
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	prefixMax := make([]int, n)
	suffixMax := make([]int, n)
	for i, v := range nums{
		if i % k == 0{
			prefixMax[i] = v
		}else{
			prefixMax[i] = max(prefixMax[i-1], v)
		}
	}
	for i := n-1; i >= 0; i--{
		if i == n-1 || (i+1)%k == 0{
			suffixMax[i] = nums[i]
		}else{
			suffixMax[i] = max(suffixMax[i+1], nums[i])
		}
	}
	ans := make([]int, n-k+1)
	for i:= range ans {
		ans[i] = max(suffixMax[i], prefixMax[i+k-1])
	}
	return ans
}

/* 325. Maximum Size Subarray Sum Equals k
** Given an integer array nums and an integer k, return the maximum length of a subarray that sums to k.
** If there is not one, return 0 instead.
 */
func maxSubArrayLen(nums []int, k int) int {
	n := len(nums)
	prefixsum := make([]int, n+1)
	m := map[int]int{0: 0}
	for i := range nums{
		prefixsum[i+1] = prefixsum[i] + nums[i]
		//m[prefixsum[i+1]] = i
		if _, ok := m[prefixsum[i+1]]; !ok{
			m[prefixsum[i+1]] = i+1 // 只记录最先开始的
		}
	}
	ans := 0
	/* 平方级，可以参考 2 sum 引入hash 优化
	   for i := 1; i <= n; i++{
	       for j := 0; j < i; j++{
	           if prefixsum[i] - prefixsum[j] == k{
	               if ans < i-j{
	                   ans = i-j
	               }
	           }
	       }
	   }
	*/
	for i := 1; i <= n; i++{
		diff := prefixsum[i] - k
		if idx, ok := m[diff]; ok {
			if ans < i-idx{
				ans = i - idx
			}
		}
	}
	return ans
}
// 进一步优化
func maxSubArrayLen_Best(nums []int, k int) int {
	m := map[int]int{0: -1}
	sum, ans := 0, 0
	for i := range nums{
		sum += nums[i]
		if idx, ok := m[sum-k];ok && ans < i-idx{
			ans = i-idx
		}
		if _, ok := m[sum]; !ok{// 记录最先出现的，保证最长
			m[sum] = i
		}
	}
	return ans
}