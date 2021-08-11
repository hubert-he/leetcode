package array

import "sort"

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
