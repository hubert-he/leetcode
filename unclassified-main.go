package main

import (
	"./unclassified"
	"fmt"
)
func test1() {
	var testdata = []int{8,1,2,2,3,3}
	ret := unclassified.SmallerNumbersThanCurrentSolution(testdata)
	fmt.Println(ret)
	nums := []int{0,0,1,1,1,1,2,3,3}
	index := unclassified.RemoveDuplicates(nums)
	fmt.Println(index)
	nums2 := []int{1,1,1,2,2,3}
	index = unclassified.RemoveDuplicates([]int{})
	fmt.Println(index)
	index = unclassified.RemoveDuplicates(nums2)
	fmt.Println(index)

	teststr := "abbaca"
	fmt.Println(unclassified.RemoveDuplicatesString(teststr))

	nums = []int{3,5,9,10,11,89,90,1000}
	fmt.Println(unclassified.SearchArray(nums, 90))

	nums = []int{7,2,5,10,8}
	m := 2
	fmt.Println(unclassified.SplitArray(nums, m))
}
func main() {
	// test1()
	nums := []int{1,3,4,2,2}
	//nums := []int{2,2,2,2,2}
	fmt.Println(unclassified.FindDuplicate(nums))
	unclassified.Run()
}
