package main

import (
	"fmt"
)

func main() {
	/*
	var testdata = []int{8,1,2,2,3,3}
	ret := unclassified.SmallerNumbersThanCurrentSolution(testdata)
	fmt.Println(ret)
	 */
	hd := generateLinkedList([]int{1,2,3,4,5,6})
	PrintLinkedList(hd)
	PrintLinkedList(reverseKGroup(hd, 2))
}

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
	curse, hd, ppre:= head,head,head
	cnt := 1
	for curse != nil {
		end := curse.Next
		if cnt % k == 0 && k != 1{
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
