package unclassified

import (
	"sort"
)

/*
给你一个数组 nums，对于其中每个元素 nums[i]，请你统计数组中比它小的所有数字的数目。
换而言之，对于每个 nums[i] 你必须计算出有效的 j 的数量，其中 j 满足 j != i 且 nums[j] < nums[i] 。
以数组形式返回答案。
示例 1：
输入：nums = [8,1,2,2,3]
输出：[4,0,1,1,3]
解释：
对于 nums[0]=8 存在四个比它小的数字：（1，2，2 和 3）。
对于 nums[1]=1 不存在比它小的数字。
对于 nums[2]=2 存在一个比它小的数字：（1）。
对于 nums[3]=2 存在一个比它小的数字：（1）。
对于 nums[4]=3 存在三个比它小的数字：（1，2 和 2）。
示例 2：
输入：nums = [6,5,4,8]
输出：[2,1,0,3]
示例 3：
输入：nums = [7,7,7,7]
输出：[0,0,0,0]
提示：
2 <= nums.length <= 500
0 <= nums[i] <= 100
链接：https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number
 */
/*除了n平方时间复杂度的方法外，提供排序-映射 方法*/

type numsmap struct {
	position int
	num	int
}
func SmallerNumbersThanCurrentSolution(nums []int) []int {
	return sortMapSolution(nums)
}
func sortMapSolution(nums []int) []int {
	/* 方式一 make 创建
	numsMaped := make([]numsmap, len(nums))
	for index,item := range(nums) {
		numsMaped[index].position = index
		numsMaped[index].num = item
	}
	 */
	// 方法二 字面量
	numsMaped := []numsmap{}
	for index,item := range(nums){
		numsMaped = append(numsMaped, numsmap{position: index, num: item}) // z注意不要漏掉numsmap 类型信息来初始化
	}
	sort.Slice(numsMaped, func (i, j int) bool {return numsMaped[i].num < numsMaped[j].num}) // 排序函数使用，工具函数 func (i, j int) bool 返回值bool不要漏掉
	serial := make([]int, len(nums))
	prev := -1 // 处理重复元素
	for index,item := range(numsMaped) {
		if prev == -1 || item.num != numsMaped[index-1].num {
			prev = index
		}
		serial[item.position] = prev
	}
	return serial
}

/*
方法三：计数排序 k+n
注意到数组元素的值域为 [0,100][0,100]，所以可以考虑建立一个频次数组 cntcnt ，cnt[i]cnt[i] 表示数字 ii 出现的次数。那么对于数字 ii 而言，小于它的数目就为 cnt[0...i-1]cnt[0...i−1] 的总和。
*/
func countingSortSolution(nums []int) []int {
	var totalCnt = [101]int{}
	for _,item := range(nums) {
		totalCnt[item]++
	}
	// 从0开始往前迭代计数和
	for i := 0; i < 100; i++ {
		totalCnt[i+1] += totalCnt[i]
	}
	var serial = []int{}
	for _,i := range(nums) {
		serial = append(serial, totalCnt[i-1])
	}
	return serial
}
