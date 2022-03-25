package slideWindow

import (
	"fmt"
	"math"
	"sort"
)

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
// 通过前缀和计算，利用two sum hash表的使用，可降低复杂度，
// 但是仔细看上面解法程序的主 for 循环，发现prefixsum 只跟当前计算值，和先前计算值有关，
// 因此可以不必事先计算前缀和和
func maxSubArrayLen_imporve(nums []int, k int) int {
	n := len(nums)
	m := map[int]int{0: 0}
	ans := 0
	prefixSum := 0 // 不必事先计算
	for i := 1; i <= n; i++{
		prefixSum += nums[i-1]
		if _, ok := m[prefixSum]; !ok{
			m[prefixSum] = i
		}
		d := prefixSum - k
		if idx, ok := m[d]; ok {
			if ans < i - idx{
				ans = i-idx
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
/* 2022-03-24 脑子忽然想用二分
** 这是因为数组中有负数，导致和不一定是单调的情况。
** 忘记了 2分的 特性
** 1. 目标函数单调性（单调递增或者递减） 此问题不满足这条， 不是单调的
** 2. 存在上下界，0 或 n
** 3. 能够通过索引访问， 满足因为从 0 到 n 查找
*/
func maxSubArrayLen_error(nums []int, k int) int {
	n := len(nums)
	check := func(step int)bool{
		sum := 0
		for i := 0; i < step; i++{
			sum += nums[i]
		}
		if sum == k { return true}
		for i := step; i < n; i++{
			sum += nums[i] - nums[i-step]
			if sum == k { return true }
		}
		return false
	}
	i, j := 0, n
	ans := 0
	for i < j {
		mid := int(uint(i+j) >> 1)
		//fmt.Println(mid, check(mid))
		if check(mid){
			ans = mid
			i = mid+1
		}else{
			j = mid-1
		}
	}
	return ans
}

/* 713. Subarray Product Less Than K
** Given an array of integers nums and an integer k,
** return the number of contiguous subarrays where the product of all the elements in the subarray is strictly less than k.
 */
// 题意未理解清楚：不能排序，计算原序列上
func numSubarrayProductLessThanK_error(nums []int, k int) int {
	m := map[int]map[int][][]int{}
	equal := func(a, b []int)bool{
		fmt.Println("-->", a, b)
		if len(a) != len(b){ return false }
		for i := range a{
			if a[i] != b[i]{
				return false
			}
		}
		return true
	}
	less := func(a []int)bool{
		n := len(a)
		prod := 1
		for _, c := range a{
			prod *= c
		}
		if prod >= k{
			return false
		}
		if m[prod] == nil{
			m[prod] = map[int][][]int{n: [][]int{a}}
		}else{
			for _, list := range m[prod][n]{
				if equal(a, list){
					fmt.Println("equal=>", a, list)
					return false
				}
			}
			m[prod][n] = append(m[prod][n], a)
			fmt.Println(a, m[prod][n])
		}
		return true
	}
	n := len(nums)
	ans := 0
	sort.Ints(nums)
	for i := 0; i < n; i++{
		for step := 1; step <= n; step++{
			if i+step > n{
				continue
			}
			if less(nums[i:i+step]){
				fmt.Println(nums[i:i+step])
				ans++
			}else{ // 排序后，发现大于直接退
				break
			}
		}
	}
	fmt.Println(m)
	return ans
}

// 2022-03-15 刷出此题 暴力累积
func numSubarrayProductLessThanK(nums []int, k int) int {
	ans, n := 0, len(nums)
	for i := 0; i < n; i++{
		prod := 1
		for j := i; j < n; j++{
			prod *= nums[j]
			if prod < k{
				ans++
			}else{ // 此处必须加break，否则数值溢出，导致 <k 统计值ans 偏大
				break
			}
		}
	}
	return ans
}
/* 滑动窗口思路分析
** 对于每个right，需要找到最小的left, 是的 nums[i] i从left到right 的累积 < k。
** 由于当 left 增加时，这个乘积是单调不增的，因此可以使用双指针控制窗口。
 */

func numSubarrayProductLessThanK_slide(nums []int, k int) int {
	if k <= 1{ return 0 } // 特殊情况
	prod,ans, left := 1, 0, 0
	for right, val := range nums{
		prod *= val
		for prod >= k{// 收缩left
			prod /= nums[left]
			left++
		}
		ans += right - left + 1
	}
	return ans
}

/* 正向思路： 二分查找的思路，取对数 乘积转和
** 即对于固定的 left, 二分查找出 最大的 right 满足 nums[left] 到 nums[right] 的乘积 小于 k
** 但由于乘积可能会非常大（在最坏情况下会达到1000^50000） 会导致溢出
** 因此我们需要对 nums 数组取对数，将乘法转换为加法，这样就不会出现数值溢出的问题了
 */
func numSubarrayProductLessThanK_BinarySearch(nums []int, k int) int {
	ans := 0
	if k <= 1{ return ans } // 特殊情况
	n := len(nums)
	logk := math.Log(float64(k))
	// 对 nums 中的每个数取对数后，我们存储它的前缀和 prefix
	prefix := make([]float64, n+1)
	for i,c := range nums{
		prefix[i+1] = prefix[i] + math.Log(float64(c))
	}
	// 对于 i 和 j 我们可以用 prefix[j+1] - prefix[i] 得到nums[i] 到 nums[j]的乘积的对数
	// 对于固定的 i 当找到最大的满足条件的 j 后，它会包含 j−i+1 个乘积小于 k 的连续子数组。
	for i := 0; i <= n; i++{
		lo, hi := i+1, n+1
		for lo < hi{
			mid := int(uint(lo+hi)>>1)
			if prefix[mid] < prefix[i] + logk - 1e-9{ // 浮点数的运算精度有限存在误差，因此必须使用误差小于某个预先给定值的方法来做。
				lo = mid + 1
			}else{
				hi = mid
			}
		}
		ans += lo - i - 1
	}
	return ans
}

/* 1151. Minimum Swaps to Group All 1's Together
** Given a binary array data,
** return the minimum number of swaps required to group all 1’s present in the array together in any place in the array.
 */
func minSwaps(data []int) int {
	n := len(data)
	oneCnt := 0
	for i := range data{
		if data[i] == 1{
			oneCnt++
		}
	}
	// 窗口大小设定为 oneCnt
	ones := 0
	for i := 0; i < oneCnt; i++{
		if data[i] == 1{ ones++ }
	}
	ans := oneCnt - ones
	for i := 1; i <= n-oneCnt; i++{
		if data[i-1] == 1{
			ones--
		}
		if data[i+oneCnt-1] == 1{
			ones++
		}
		t := oneCnt - ones
		if ans > t{ ans = t }
	}
	return ans
}

// 官方代码： 例如 只有 0 和 1 两种值 来优化代码, 学习coding方式 特别是滑动窗口的
func minSwaps_official(data []int) int {
	n := len(data)
	oneCnt := 0
	for i := range data{
		oneCnt += data[i]
	}
	ones := 0
	for i := 0; i < oneCnt; i++{
		ones += data[i]
	}
	max := ones
	for i := oneCnt; i < n; i++{
		ones += data[i] - data[i-oneCnt]
		if max < ones{
			max = ones
		}
	}
	return oneCnt - max
}












