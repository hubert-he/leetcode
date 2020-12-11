package unclassified

import "fmt"

/**
给定一个链表，判断链表中是否有环。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
如果链表中存在环，则返回 true 。 否则，返回 false 。
进阶：你能用 O(1)（即，常量）内存解决此问题吗？
示例 1：
输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。
示例 2：
输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。
示例 3：
输入：head = [1], pos = -1
输出：false
解释：链表中没有环。
提示：
链表中节点的数目范围是 [0, 104]
-105 <= Node.val <= 105
pos 为 -1 或者链表中的一个 有效索引 。

链接：https://leetcode-cn.com/problems/linked-list-cycle
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}
func Run() {
	t1 := []int{1,2,3,4,5,6}
	//t2 := []int{1}
	headt1 := GenerateLinkedList(t1)
	//headt2 := GenerateLinkedList(t2)
	//PrintLinkedList(headt1)
	//PrintLinkedList(reverseBetween(headt1, 2, 4))

	PrintLinkedList(reverseKGroup(headt1,2))

//	PrintLinkedList(headt2)
//	PrintLinkedList(reverseBetween(headt2, 1, 4))
}
func GenerateLinkedList(nums []int) *ListNode {
	if len(nums) <= 0 {
		return nil
	}
	head := ListNode{nums[0], nil}
	pre := &head
	for i := 1; i < len(nums); i++ {
		tmp := ListNode{nums[i], nil}
		pre.Next = &tmp
		pre = pre.Next
	}
	return &head
}

func PrintLinkedList(head *ListNode){
	curr := head
	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
	}
	fmt.Printf("\n---\n")
}

func reverseBetweenII(head *ListNode, m int, n int) *ListNode {
	var pre,post *ListNode
	hd,tl := head,head
	for n > 1 {
		if m > 1 {
			pre = hd
			hd = hd.Next
			m--
		}
		tl = tl.Next
		n--
	}
	end := tl.Next
	post = tl.Next
	cur := hd
	for cur != end {
		temp := cur.Next
		cur.Next = post
		post = cur
		cur = temp
	}
	if pre != nil {
		pre.Next = post
	}else {
		head = post
	}

	return head
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// 3. 查找hd tl prehd posttl节点
	var prehd,posttl,hd,tl *ListNode
	hd,tl = head,head
	for ;n > 1; n-- { // m n 从大于1开始计数
		if m > 1 {
			prehd = hd
			hd = hd.Next
			m--  // 容易漏掉
		}
		tl = tl.Next
	}
	posttl = tl.Next
	// 1. 先写反转
	// hd tl 标记为要反转的首尾
	// prehd 标记为hd的pre节点，无则为nil
	// posttl 标记为tl的post节点， 无则为nil
	pre := posttl
	cur := hd
	for cur != posttl {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	// 2. 根据prehd 情况 确定head首节点
	if prehd != nil {
		prehd.Next = pre
	}else{
		head = pre
	}
	return head
}
/*
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
示例：
给你这个链表：1->2->3->4->5
当 k = 2 时，应当返回: 2->1->4->3->5
当 k = 3 时，应当返回: 3->2->1->4->5
说明：
你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	curse,hd,ppre := head,head,head
	cnt := 1
	for curse != nil {
		end := curse.Next // 必须预保存
		if cnt % k == 0{
			pre := end
			lastTop := hd
			for hd != end {
				tmp := hd.Next
				hd.Next = pre
				pre = hd
				hd = tmp
			}
			if ppre == head {
				head = pre
			} else {
				ppre.Next = pre // 上一组尾节点连当前组首节点
				ppre = lastTop // 更新为当前组尾节点
			}
		}
		curse = end
		//curse = curse.Next 此时变换掉了
		cnt++
	}
	return head
}

func hasCycle(head *ListNode) bool {
	return false
}