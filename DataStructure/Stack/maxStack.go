package Stack

import (
	"container/heap"
	"container/list"
)

/*
** 716. Max Stack
** Design a max stack data structure that supports the stack operations and supports finding the stack's maximum element.
** Implement the MaxStack class:
	MaxStack() Initializes the stack object.
	void push(int x) Pushes element x onto the stack.
	int pop() Removes the element on top of the stack and returns it.
	int top() Gets the element on the top of the stack without removing it.
	int peekMax() Retrieves the maximum element in the stack without removing it.
	int popMax() Retrieves the maximum element in the stack and removes it.
		If there is more than one maximum element, only remove the top-most one.
** Follow up:
	Could you come up with a solution that supports O(1) for each top call and O(logn) for each other call? 
 */
type bigHeap	[]int
func(h bigHeap) Len()int{ return len(h) }
func(h bigHeap) Less(i, j int)bool { return h[i] > h[j] }
func(h bigHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func(h *bigHeap) Push(v interface{}){
	(*h) = append((*h), v.(int))
}
func(h *bigHeap) Pop()(v interface{}){
	ret := (*h)[h.Len()-1]
	(*h) = (*h)[:h.Len()-1]
	return ret
}
/* 方法一：类似于 155. Min Stack
** 双栈类似思路，或者栈元素存放值和当前的最小值/最大值 这类思路
** 但是此方法在 PopMax 存在O(n) 时间复杂度
** 无法提升到 O(logN) 复杂度级别
 */
/* 方法二：双向链表 + 平衡树
** 使用线性的数据结构（例如数组和栈）无法在较短的时间复杂度内完成 popMax() 操作，
** 因此我们可以考虑使用双向链表 + 平衡树，其中双向链表用来表示栈，平衡树中的每一个节点存储一个键值对，
** 其中"键" 表示某个在栈中出现的值，"值"为一个列表，表示这个值在双向链表中出现的位置，存储方式为指针
** 平衡树的插入，删除和查找操作都是 O(logn) 的，而通过平衡树定位到双向链表中的某个节点后，对该节点进行删除也是 O(1) 的，
** 因此所有操作的时间复杂度都不会超过 O(logn)。
** 算法：
** 使用双向链表存储栈，使用带键值对的平衡树（Java TreeMap），存储栈中出现的值以及这个值在双向链表中出现的位置
** GoLang中不存在TreeMap，因此引入了 大顶堆
 */
type MaxStack struct {
	h 	bigHeap
	kv	map[int][]*list.Element
	dll *list.List
}


func Constructor() MaxStack {
	this := MaxStack{}
	this.kv = map[int][]*list.Element{}
	this.dll = list.New()
	heap.Init(&this.h)
	return this
}


func (this *MaxStack) Push(x int)  {
	t := this.dll.PushBack(x) // 现在知道为何PushBack 为何设置返回值了 *list.Element 类型
	this.kv[x] = append(this.kv[x], t)
	heap.Push(&this.h, x)
}


func (this *MaxStack) Pop() int {
	val := this.dll.Remove(this.dll.Back())
	l := this.kv[val.(int)]
	l = l[:len(l)-1]  // 需要把 l 赋值回去
	this.kv[val.(int)] = l
	/*if len(l) == 0{
		//heap.Pop(&this.h)
		// 此处删除的不一定Pop的最大值
		//heap.Remove(&this.h, val.(int)) //需要fix， 第二个参数时下标
	}*/
	return val.(int)
}


func (this *MaxStack) Top() int {
	val := this.dll.Back().Value
	return val.(int)
}


func (this *MaxStack) PeekMax() int {
	top := this.h[0]
	for len(this.kv[top]) <= 0{
		//top = heap.Pop(&this.h).(int)  ❌ top值不对 逻辑错误
		heap.Pop(&this.h)
		top = this.h[0]
	}
	return this.h[0]
}


func (this *MaxStack) PopMax() int {
	max := this.PeekMax()
	l := this.kv[max]
	target := l[len(l)-1]
	l = l[:len(l)-1] // 需要把 l 赋值回去
	this.kv[max] = l
	if len(l) == 0{
		heap.Pop(&this.h)
	}
	this.dll.Remove(target)
	return max
}


/**
 * Your MaxStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.PeekMax();
 * param_5 := obj.PopMax();
 */