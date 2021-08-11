package LinkedList
type ListNode struct {
	Val int
	Next *ListNode
}
/*剑指 Offer II 077. 链表排序
给定链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
 */
func SortListQuickSort(head *ListNode) *ListNode {
	var partition func(start, end *ListNode)
	partition = func(start, end *ListNode){
		if start == nil || start == end{
			return
		}
		// p1 始终指向待插入的位置
		pp1, p1, p2 := start, start.Next, start.Next
		for p2 != end {
			if p2.Val < start.Val{
				p1.Val, p2.Val = p2.Val, p1.Val
				pp1 = p1
				p1 = p1.Next
			}
			p2 = p2.Next
		}
		// 找到正确位置
		start.Val, pp1.Val = pp1.Val, start.Val
		partition(start, pp1)
		partition(p1, end)
	}
	partition(head, nil)
	return head
}

func SortListMergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	half := func(l *ListNode)*ListNode{// fast slow pointer
		preslow, slow, fast := l, l, l
		for fast != nil {
			fast = fast.Next
			if fast == nil {
				break
			}
			fast = fast.Next
			preslow = slow
			slow = slow.Next
		}
		preslow.Next = nil
		return slow
	}
	merge := func(l *ListNode, r *ListNode) *ListNode{
		dumb := &ListNode{}
		cur := dumb
		for {
			if l != nil && (r == nil || l.Val < r.Val){
				cur.Next = l
				l = l.Next
				cur = cur.Next
			}else if r != nil {
				cur.Next = r
				r = r.Next
				cur = cur.Next
			}else{
				break
			}
		}
		return dumb.Next
	}
	// mid := length/2
	h := half(head)
	return merge(SortListMergeSort(head), SortListMergeSort(h))
}



