package main

import (
	"./unclassified"
	"fmt"
	"strconv"
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
	a := monotoneIncreasingDigitsII(2331)
	fmt.Println(a)
}

func monotoneIncreasingDigits(n int) int {
	s := []byte(strconv.Itoa(n))
	i := 1
	for i < len(s) && s[i] >= s[i-1] {
		i++
	}
	if i < len(s) {
		for i > 0 && s[i] < s[i-1] {
			s[i-1]--
			i--
		}
		for i++; i < len(s); i++ {
			s[i] = '9'
		}
	}
	ans, _ := strconv.Atoi(string(s))
	return ans
}

func monotoneIncreasingDigitsII(n int) int {
	num := []byte(strconv.Itoa(n)) // 注意字符串是只读类型
	length := len(num)
	flag := length
	for i := length - 1; i > 0; i--{
		if num[i-1] > num[i] {
			flag = i
			num[i-1]--
		}
	}
	for i := flag; i < length; i++{
		num[i] = '9'
	}
	x, _ := strconv.Atoi(string(num))
	return x
}
