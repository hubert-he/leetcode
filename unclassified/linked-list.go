package unclassified

import (
	"fmt"
	"math"
)

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

//	PrintLinkedList(reverseKGroup(headt1,2))

//	PrintLinkedList(headt2)
//	PrintLinkedList(reverseBetween(headt2, 1, 4))

//	PrintLinkedList(oddEvenList(headt1))
	PrintLinkedList(reorderList(headt1))
	t2 := []int{1,2,1}
	headt2 := GenerateLinkedList(t2)
	fmt.Println(isPalindrome(headt2))

	t3 := []int {1,0,0,1,0,0,1,1,1,0,0,0,0,0,0}
	headt3 := GenerateLinkedList(t3)
	fmt.Println(getDecimalValue(headt3))
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

func oddEvenList(head *ListNode) *ListNode {
	var odd,even,tmpeven *ListNode
	cur := head
	isodd := true
	for cur != nil {
		t := cur.Next  // 此位置设置cur.Next = nil 可避免最后结果出现环 忘记清理尾节点Next指针导致
		if isodd {
			if odd == nil {
				odd = cur
			} else {
				odd.Next = cur
				odd = odd.Next
			}
			isodd = false
		}else {
			if even == nil {
				even = cur
				tmpeven = even
			} else {
				even.Next = cur
				even = even.Next
			}
			isodd = true
		}
		cur.Next = nil
		cur = t
	}
	if odd != nil { // 遗漏点： [] 空链表情况处理
		odd.Next = tmpeven
	}
	return head
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true  // 快慢指针操作，先排除特殊情况，简化代码
	}
	dummyhead := &ListNode{Next: nil}
	slow,fast := head,head.Next.Next
	odd := false
	for fast != nil {
		tmp := slow.Next
		if fast.Next != nil { // 不可放到最后，因为slow 要修改位置
			fast = fast.Next.Next
			odd = false
		}else {
			odd = true
			fast = nil
		}
		slow.Next = dummyhead.Next
		dummyhead.Next = slow
		slow = tmp
	}
	tmp := slow.Next
	if !odd { // 考虑奇 偶性
		slow.Next = dummyhead.Next
		dummyhead.Next = slow
	}
	slow = tmp
	cur := dummyhead.Next
	for slow != nil {
		if slow.Val != cur.Val{
			return false
		}
		slow = slow.Next
		cur = cur.Next
	}
	if cur == nil && slow == nil {
		return true
	}
	return false
}

func reverseList(head *ListNode) *ListNode { // 反转链表标准操作
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

func endOfFirstHalf(head *ListNode) *ListNode { // 找中点标准操作
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func isPalindromeII(head *ListNode) bool {
	if head == nil {
		return true
	}

	// 找到前半部分链表的尾节点并反转后半部分链表
	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverseList(firstHalfEnd.Next)

	// 判断是否回文
	p1 := head
	p2 := secondHalfStart
	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// 还原链表并返回结果
	firstHalfEnd.Next = reverseList(secondHalfStart)
	return result
}

func reorderList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head,head
	var pre *ListNode
	for fast != nil {
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = nil
		}
		pre = slow
		slow = slow.Next
	}
	cur  := slow
	pre.Next = nil
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre.Next
		pre.Next = cur
		cur = tmp
	}
	post := pre.Next
	pre.Next = nil
	pre = head
	for pre != nil && post != nil {
		t1 := pre.Next
		t2 := post.Next

		pre.Next = post
		post.Next = t1
		pre = t1
		post = t2
	}
	return head
}


func getDecimalValueII(head *ListNode) int {
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}
	cur = head
	sum := 0
	for cur != nil {
		sum += cur.Val*int(math.Pow(2,float64(length-1)))
		length--
		cur = cur.Next
	}
	return sum
}
func getDecimalValue(head *ListNode) int {
	// 2进制转10进制 归纳求和方式
	// abcd ==> a*2^3 + b*2^2 + c*2^1 + d*2^0
	//      ==> [(a*2+b)*2+c]*2+d ==> {[(0*2+a)*2+b]*2+c}*2+d
	cur := head
	sum := 0
	for cur != nil {
		sum = sum*2+cur.Val
		cur = cur.Next
	}
	return sum
}

func hasCycle(head *ListNode) bool {
	return false
}