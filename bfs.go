package main

import "fmt"

func main() {
	var testcase = []int{0,1}
	result := permute(testcase)
	fmt.Println(result)
}

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums} // 切片的字面量初始化
	}
	var ret = &([][]int{})
	bfs(nums, ret, 0, len(nums)-1)
	return *ret
}

func bfs(nums []int, ret *[][]int, first, maxIndex int) {
	fmt.Printf("%#v %p\n", ret, ret)
	if first == maxIndex {
		//var temp = []int{}
		temp := make([]int, len(nums))
		copy(temp, nums)
		t := append(*ret, temp)
		*ret = t

		return
	}
	for i := first; i <= maxIndex;i++ {
	//for i,j := range(nums[first:]){ // 此处i 的值是从0开始计的
		nums[first], nums[i] = nums[i], nums[first]
		bfs(nums, ret, first+1, maxIndex)
		nums[first], nums[i] = nums[i], nums[first]
	}
}