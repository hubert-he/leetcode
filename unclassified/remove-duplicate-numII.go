package unclassified

import "fmt"

func RemoveDuplicates(nums []int) int {
	return removeDuplicatesIV(nums)
}

func removeDuplicatesI(nums []int) int {
	i := 0
	for ; i < len(nums); {
		if i >=2 && nums[i] == nums[i-1] && nums[i] == nums[i-2]{
			nums = append(nums[:i], nums[i+1:]...)
		}else {
			i++
		}
	}
	return i
}

func removeDuplicatesII(nums []int) int {
	var i, j int
	for ;j < len(nums);j++ {
		 if i > 1 && nums[j] == nums[i-1] && nums[j] == nums[i-2] {
		 	fmt.Println(nums[i])
		 } else {
		 	nums[i] = nums[j] // 这句不好理解
		 	i++
		 }
	}
	return i
}

func removeDuplicatesIII(nums []int) int {
	var i int
	for j := 0; j < len(nums);j++ {
		if i <= 1 || (nums[j] != nums[i-1] || nums[j] != nums[i-2]) { // 可以合并，因为是有序序列
			nums[i] = nums[j] // 这句不好理解
			i++
		}
	}
	return i
}

func removeDuplicatesIV(nums []int) int {
	var i,k int
	k = 2
	for j := 0; j < len(nums);j++ {
		if i <= k-1 || nums[j] != nums[i-k] {
			nums[i] = nums[j] // 这句不好理解, 有序序列
			i++
		}
	}
	return i
}