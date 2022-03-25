package LinkedList

import "fmt"

/* 强化练习双向循环链表
** 707. Design Linked List
** Design your implementation of the linked list. You can choose to use a singly or doubly linked list.
** A node in a singly linked list should have two attributes: val and next.
** val is the value of the current node, and next is a pointer/reference to the next node.
** If you want to use the doubly linked list,
** you will need one more attribute prev to indicate the previous node in the linked list.
** Assume all nodes in the linked list are 0-indexed.
** Implement the MyLinkedList class:
	MyLinkedList() Initializes the MyLinkedList object.
	int get(int index) Get the value of the indexth node in the linked list. If the index is invalid, return -1.
	void addAtHead(int val) Add a node of value val before the first element of the linked list.
		After the insertion, the new node will be the first node of the linked list.
	void addAtTail(int val) Append a node of value val as the last element of the linked list.
	void addAtIndex(int index, int val) Add a node of value val before the indexth node in the linked list.
		If index equals the length of the linked list, the node will be appended to the end of the linked list.
		If index is greater than the length, the node will not be inserted.
	void deleteAtIndex(int index) Delete the indexth node in the linked list, if the index is valid.
 */

type Node struct{
	val         int
	prev, next   *Node
}

type MyLinkedList struct {
	head    *Node
	size 	int // 增加size 方便处理index 情况
}
// 带伪首节点，方便插入删除
func Constructor() MyLinkedList {
	dumb := &Node{val: -1}
	dumb.prev, dumb.next = dumb, dumb
	return MyLinkedList{head: dumb}
}
// index 从 0 开始计数
func (this *MyLinkedList) Index(index int) *Node{
	if index >= this.size{
		return nil
	}
	p := this.head.next
	for index > 0 && p != this.head{
		p = p.next
		index--
	}
	return p
}

func (this *MyLinkedList) Get(index int) int {
	p := this.Index(index)
	if p != nil {
		return p.val
	}
	return -1
}
// 双向循环链表：注意修改的个数，涉及4个端点指针的修改
func (this *MyLinkedList) AddAtHead(val int)  {
	newNode := &Node{val: val, prev: this.head, next: this.head.next}
	// 对于this.head.prev 的修改，一般是不需修改，但是如果size = 0情况下，需要增加prev修改
	/*
	if this.head.next == this.head{ // 插入第一个点
		this.head.prev = newNode
	}else{ // 容易漏掉，插入非第一个点时
		this.head.next.prev = newNode
	}*/
	this.head.next.prev = newNode
	this.head.next = newNode
	this.size++
}
// 双向循环链表：注意修改的个数，涉及4个端点指针的修改
func (this *MyLinkedList) AddAtTail(val int)  {
	newNode := &Node{val: val, prev: this.head.prev, next: this.head}
	this.head.prev.next = newNode
	this.head.prev = newNode
	this.size++
}
// 双向循环链表：注意修改的个数，涉及4个端点指针的修改
// 题目限制当index == size 的时候，就是追加
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index == this.size{
		this.AddAtTail(val)
		return
	}
	if index > this.size{ return }
	p := this.Index(index)
	newNode := &Node{val: val, prev: p.prev, next: p}
	p.prev.next = newNode
	p.prev = newNode
	this.size++
}
// 双向循环链表：注意修改的个数，涉及4个端点指针的修改
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index >= this.size{ return }
	p := this.Index(index)
	p.next.prev = p.prev
	p.prev.next = p.next
	p.prev, p.next = nil, nil
	this.size--
}

func (this *MyLinkedList) Print(){
	p := this.head.next
	for p != this.head{
		fmt.Printf("%d-%d-%d ", p.prev.val, p.val, p.next.val)
		p = p.next
	}
	fmt.Println(" size ", this.size)
}
/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
