package array

import (
	"fmt"
	"math"
	"sort"
)

func twoSum(nums []int, target int) []int {
	cache := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := cache[target - nums[i]]; !ok {
			cache[nums[i]] = i
		}else {
			return []int{i,cache[target - nums[i]]}
		}
	}
	return nil
}
// hashMap 解法
func twoThing(nums []int, target int, thing func(int,int) int) []int {
	cache := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := cache[thing(nums[i], target)]; ok {
			return []int{i, cache[thing(nums[i], target)]}
		} else {
			cache[nums[i]] = i
		}
	}
	return nil
}

func difference(a int, t int) int {
	return t - a
}

// 使用双指针，必须保证是有序序列，因此先排序
func twoSumTwoPointer(nums []int, target int) []int {
	result := []int{}
	// 务必注意，golang的sort函数是就地排序，因此需要额外分配空间
	origin := make([]int, len(nums))
	copy(origin, nums)
	sort.Sort(sort.IntSlice(origin))
	for first,second := 0,len(origin) - 1; first < len(origin); first++{
		if first == 0 || origin[first] != origin[first - 1] {
			for first < second && origin[first] + origin[second] > target {
				second--
			}
			if first == second {
				break;
			}
			if origin[first] + origin[second] == target{
				// two sum 要求的是原始数组的索引值
				// result = append(result, []int{first,second}...)
				index := []int{first, second}
				for k := 0; k < len(nums) && len(index) > 0; k++ {
					for t := 0; t < len(index); t++ {
						if origin[index[t]] == nums[k] {
							result = append(result, k)
							index = append(index[:t], index[t+1:]...)
						}
					}
				}
			}
		}
	}
	return result
}

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。
示例：
给定数组 nums = [-1, 0, 1, 2, -1, -4]，
满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
链接：https://leetcode-cn.com/problems/3sum
*/
func threeSum(nums []int, summary int) [][]int {
	sort.Sort(sort.IntSlice(nums)) // 规避重复解
	result := [][]int{}
	for first := 0; first < len(nums); first++ {
		// 需要和上一次枚举的数不相同
		if first == 0 || nums[first] != nums[first - 1] {
			third := len(nums) - 1
			target := summary - nums[first]
			for second := first + 1; second < len(nums); second++ {
				// 双指针 双向跳动，可以使用双指针，是因为
				// 1. 数组经过排序，变为有序数组序列
				// 2. first + second +third = target，固定first后，second 和 third 是有限制关系的，second值变得大，third值就得小
				if second == first + 1 || nums[second] != nums[second - 1] {
					// 需要保证 second 的指针在 third 的指针的左侧
					for second < third && nums[second] + nums[third] > target {
						third--
					}
					if second == third {
						// 如果指针重合，随着 second 后续的增加
						// 就不会有满足 first + second +third = 0 并且 second < third 的third了，可以退出循环
						break
					}
					if nums[second] + nums[third] == target {
						//result = append(result, []int{first, second, third})
						result = append(result, []int{nums[first], nums[second], nums[third]})
					}
				}
			}
		}
	}
	return result
}
// 2021-11-03 重刷此题
func ThreeSum(nums []int) [][]int {
	n := len(nums)
	ans := [][]int{}
	if n < 3{
		return ans
	}
	sort.Ints(nums)
	for i := 0; i < n; i++{
		// 去重-1
		if i != 0 && nums[i] == nums[i-1]{
			continue
		}
		// 问题转换为求 0 - nums[i] 的 2sum 问题
		target := 0 - nums[i]
		j, k := i+1, n-1
		for j < k {
			sum := nums[j] + nums[k]
			if sum == target{
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				// 去重-2 勿忘 j < k
				for j < k && nums[j] == nums[j-1]{
					j++
				}
				for j < k && nums[k] == nums[k+1]{
					k--
				}
			}else if sum < target{
				j++
			}else{
				k--
			}
		}
	}
	return ans
}

/*
考虑一般情况： Sum(N, X) 从给定数组中找N个和为X的数问题（要求不出现重复解）。
构建一个递归过程
1、排序，需要从算法上规避重复解，因此需要排个序。
2、从左往右遍历，先挑选出第一个数x1，然后在x1右侧剩下的数中递归调用Sum(N-1, X-x1)
3、直到Sum(2, X)问题
因此，解决了Sum(2, X)和Sum(N, X)-->Sum(N-1, X-x1)的递归调用，NSum的问题都解决了。
链接：https://leetcode-cn.com/problems/3sum/solution/shuang-wai-wai-yi-ge-di-gui-tao-lu-jie-jue-2sumyi-/
*/
func commonSum(nums []int, target int, k int) [][]int {
	if len(nums) < k {
		return nil
	}
	sort.Sort(sort.IntSlice(nums)) // 排序+或再增加元素去重
	return kSum(nums, target, k)

}
func kSum(nums []int, target int, k int) [][]int{
	length := len(nums)
	//fmt.Println(nums, length, k)
	if length < k {
		return nil
	}
	tmpList := [][]int{}
	if k == 2 { // 递归结束条件
		for m,n := 0,length - 1; m < n; m++{
			if m == 0 || nums[m] != nums[m-1]{ // 针对有序的序列，通过判断来消除重复结果
				for n>m && nums[n] + nums[m] > target {
					n--
				}
				if m == n{
					break
				}
				if nums[n] + nums[m] == target{
					tmpList = append(tmpList, []int{nums[m],nums[n]})
				}
			}
		}
	}else {
		for i := 0; i < len(nums); i++ {

			if i == 0 || nums[i] != nums[i-1]{ // 针对有序的序列，通过判断来消除重复结果
				ret := kSum(nums[i+1:], target - nums[i], k - 1)
				if ret != nil{
					for j := 0; j < len(ret); j++ {
						ret[j] = append(ret[j], nums[i])
					}
					tmpList = append(tmpList, ret...)
				}
			}
		}
	}
	return tmpList
}
//287. 寻找重复数


/* LCP 18. 早餐组合
小扣在秋日市集选择了一家早餐摊位，一维整型数组 staple 中记录了每种主食的价格，一维整型数组 drinks 中记录了每种饮料的价格。小扣的计划选择一份主食和一款饮料，且花费不超过 x 元。请返回小扣共有多少种购买方案。
注意：答案需要以 1e9 + 7
 */
func BreakfastNumber(staple []int, drinks []int, x int) int {
	const mod int = 1e9+7
	ans := 0
	bucket := make([]int, x+1)
	for i := range staple{ // 统计频次
		if staple[i] < x{
			bucket[staple[i]]++
		}
	}
	for i := 2; i <= x; i++{ // 计算前缀和
		bucket[i] += bucket[i-1]
	}
	for i := range drinks{
		diff := x - drinks[i]
		if diff < 0{
			continue
		}
		ans = (ans + bucket[diff])%mod
	}
	return ans
}

func BreakfastNumber2(staple []int, drinks []int, x int) int {
	const mod int = 1e9+7
	sort.Ints(staple)
	sort.Ints(drinks)
	ans := 0
	n := len(drinks)
	for i := range staple{
		target := x - staple[i]
		i, j := 0, n-1
		mid := 0
		for i < j{
			mid = (i+j) >> 1
			if drinks[mid] <= target{
				i = mid + 1
			}else{
				j = mid - 1
			}
		}
		//fmt.Println(mid,i)
		if drinks[i] <= target{
			ans = (ans + i + 1)%mod
		}else{
			ans = (ans + i)%mod
		}
	}
	return ans
}

/*LCP 40. 心算挑战
「力扣挑战赛」心算项目的挑战比赛中，要求选手从 N 张卡牌中选出 cnt 张卡牌，
 若这 cnt 张卡牌数字总和为偶数，则选手成绩「有效」且得分为 cnt 张卡牌数字总和。
 给定数组 cards 和 cnt，其中 cards[i] 表示第 i 张卡牌上的数字。
 请帮参赛选手计算最大的有效得分。若不存在获取有效得分的卡牌方案，则返回 0。
 另外： 这道题，不能应用背包DP， 原因是 dp[i] dp[i-1] 无法直接推导出，因为奇偶性的原因
 */
func MaxmiumScore(cards []int, cnt int) int {
	n := len(cards)
	if cnt > n {
		return 0
	}
	sort.Sort(sort.Reverse(sort.IntSlice(cards)))
	sum := 0
	for i := 0; i < cnt; i++{
		sum += cards[i]
	}
	// 1. 若 sum 是偶数则直接返回
	if sum & 1 == 0{
		return sum
	}
	//2. sum 是奇数
	/* 需要从前 cnt 个元素中选一个元素 x，并从后面找一个最大的且奇偶性和 x 不同的元素替换 x，这样就可以使 sum 为偶数
	注意，这里存在2种情况：
	情况-1：直接用cards[cnt:]最大并且相反奇偶性的数替换
	情况-2：替换前 cnt 个元素中，最小的且奇偶性和 card[cnt−1] 不同的元素 <== 这个比较难想到
	 */
	ans := 0
	replace := func(x int){
		for _, v := range cards[cnt:]{
			if x & 1 != v & 1{
				t := sum - x + v
				if ans < t{
					ans = t
				}
				break // 易漏点-1： 找最近的
			}
		}
	}
	replace(cards[cnt-1])// 情况-1： 替换当前最小的数，保证总和满足偶数情况下最大
	//情况-2： 找一个最小的且奇偶性不同于 cards[cnt-1] 的元素，将其替换掉
	for i := cnt - 2; i >= 0; i--{
		if cards[i] & 1 != cards[cnt-1] & 1{
			replace(cards[i])
		}
	}
	return ans
}
/* 考虑构造cnt个数和为偶数，很明显有 k个奇数 和 cnt-k个偶数，其中 k 为偶数[重要]
** 让结果最大：偶数之和和奇数之和尽可能大
*/
func MaxmiumScorePrefixSum(cards []int, cnt int) int {
	n := len(cards)
	sort.Sort(sort.Reverse(sort.IntSlice(cards)))
	//1. 构造奇偶两个的前缀和数组
	odd, even := []int{0}, []int{0}
	for i := 0; i < n; i++{
		if cards[i] & 1 == 0{
			//even = append(even, even[i] + cards[i]) // 易错-1 前缀和计算
			even = append(even, even[len(even)-1] + cards[i])
		}else{
			//odd = append(odd, odd[i] + cards[i])
			odd = append(odd, odd[len(odd)-1] + cards[i])
		}
	}
	//fmt.Println(odd, even)
	ans := 0
	//2. 枚举所有组合中奇数的个数 k(k必须是偶)和 cnt-k（需判断是否足够）个偶数，它们都取最大则该轮组合结果最大。因此更新所有组合最大值就是答案
	for i := 0; i < len(odd); i += 2{
		// 遗漏点-1 判断数量是否满足
		if cnt >= i && cnt < len(even)+i && ans < odd[i] + even[cnt-i]{
			ans = odd[i] + even[cnt-i]
		}
	}
	return ans
}

func MaxmiumScoreDP(cards []int, cnt int) int {
	return 0
}

/* 2006. Count Number of Pairs With Absolute Difference K
** Given an integer array nums and an integer k, return the number of pairs (i, j) where i < j such that |nums[i] - nums[j]| == k.
** The value of |x| is defined as: x if x >= 0.  -x if x < 0.
 */
/* 思考不够深入，❌
func CountKDifference(nums []int, k int) int {
	m := map[int]int{}
	for _, u := range nums{
		m[u]++
	}
	ans := 0
	fmt.Println(m)
	for u := range m{
		ans += m[u]*m[u-k] + m[k-u]*m[u]
		if u-k == u || k-u == u{
			ans -= 1 // 排除自身
		}
	}
	return ans
}
 */
func CountKDifferenceDetail(nums []int, k int) int {
	m := map[int][]int{}
	for i, u := range nums{
		m[u] = append(m[u], i)
	}
	ans := map[[2]int]struct{}{}
	for i := range m{
		// nums[j] = nums[i]+k 或 nums[i] - k
		for _, ov := range m[i]{
			for _, pv := range m[i+k]{
				if ov > pv{
					ans[[2]int{pv, ov}] = struct{}{}
				}else{
					ans[[2]int{ov, pv}] = struct{}{}
				}
			}
			for _, pv := range m[i-k]{
				if ov > pv{
					ans[[2]int{pv, ov}] = struct{}{}
				}else{
					ans[[2]int{ov, pv}] = struct{}{}
				}
			}
		}
	}
	fmt.Println(ans)
	return len(ans)
}
func CountKDifference(nums []int, k int) int {
	m := map[int][]int{}
	for i, u := range nums{
		m[u] = append(m[u], i)
	}
	ans := 0
	for i := range m{
		// nums[j] = nums[i]+k 或 nums[i] - k
		x, y, z := len(m[i]), len(m[i+k]), len(m[i-k])
		//fmt.Println(x*y, x*z)
		ans += (x*y) + (x * z)
	}
	return ans/2 // 结果肯定是成对的
}
func countKDifference2(nums []int, k int) (ans int) {
	cnt := map[int]int{}
	for _, v := range nums {
		cnt[v]++
	}
	for _, v := range nums {
		ans += cnt[v-k] // 直接统计target
	}
	return
}

/*1413. Minimum Value to Get Positive Step by Step Sum
** Given an array of integers nums, you start with an initial positive value startValue.
** In each iteration, you calculate the step by step sum of startValue plus elements in nums (from left to right).
** Return the minimum positive value of startValue such that the step by step sum is never less than 1.
 */
/* 已经想到用Prefix Sum 处理
** 但是没有想到一个情况： 只需要找前缀和数组中最小的数即可 <== 这个点没有考虑到
** 找最小的正数，即只需要找前缀和数组中最小的数即可，当最小的数加上startValue，都能满足>=1的条件时，其余前缀和都能满足；
** 方法一： 故使用快排将前缀和（去除第一个0）按照从小到大排序，取第一个数进行分析即可。
 */
func MinStartValue(nums []int) int {
	sum,minSum := 0, math.MaxInt32
	for i := range nums{
		sum += nums[i]
		if minSum > sum{
			minSum = sum
		}
	}
	ans := 1
	if minSum < 0{
		ans = 1 - minSum
	}
	return ans
}