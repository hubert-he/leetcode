package main

import (
	"./unclassified"
	"fmt"
	"strconv"
)

type ListNode struct {
	Val int
	Next *ListNode
}
func generateLinkedList(nums []int) *ListNode{
	var head *ListNode
	for i := len(nums) - 1; i >= 0; i--{
		item := ListNode{nums[i], head}
		head = &item
	}
	return head
}

func PrintLinkedList(head *ListNode){
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println("\n---\n")
}
func reverseKGroup(head *ListNode, k int) *ListNode {
	curse, hd, ppre := head, head, head
	cnt := 1
	for curse != nil {
		end := curse.Next
		if cnt%k == 0 && k != 1 {
			pre := end
			top := hd // 为什么不是 ppre = hd 用top 暂存
			for hd != end {
				tmp := hd.Next
				hd.Next = pre
				pre = hd
				hd = tmp
			}
			if ppre == head {
				fmt.Println(ppre)
				head = pre
			} else {
				ppre.Next = pre
				ppre = top
				//fmt.Println(ppre)
			}
		}
		fmt.Println(cnt)
		curse = end
		cnt++
	}

	return head
}
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
	//nums := []int{1,3,4,2,2}
	//nums := []int{2,2,2,2,2}
	//fmt.Println("FindDuplicate-->", unclassified.FindDuplicate(nums))
	fmt.Println("FindDuplicate-->", unclassified.FindDuplicate([]int{2,6,4,1,3,1,5}))
	unclassified.Run()
	a := monotoneIncreasingDigitsII(2331)
	fmt.Println(a)

	unclassified.RunArray()

	unclassified.HeapRun()

	/*
		var testdata = []int{8,1,2,2,3,3}
		ret := unclassified.SmallerNumbersThanCurrentSolution(testdata)
		fmt.Println(ret)
	*/
	hd := generateLinkedList([]int{1,2,3,4,5,6})
	PrintLinkedList(hd)
	PrintLinkedList(reverseKGroup(hd, 2))
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
