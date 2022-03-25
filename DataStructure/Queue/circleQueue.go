package Queue

import "sync"

/* 622. Design Circular Queue
** Design your implementation of the circular queue.
** The circular queue is a linear data structure in which the operations are performed based on
** FIFO (First In First Out) principle and the last position is connected back to the first position to make a circle.
** It is also called "Ring Buffer".
** One of the benefits of the circular queue is that we can make use of the spaces in front of the queue.
** In a normal queue, once the queue becomes full,
** we cannot insert the next element even if there is a space in front of the queue.
** But using the circular queue, we can use the space to store new values.
** Implementation the MyCircularQueue class:
	MyCircularQueue(k) Initializes the object with the size of the queue to be k.
	int Front() Gets the front item from the queue. If the queue is empty, return -1.
	int Rear() Gets the last item from the queue. If the queue is empty, return -1.
	boolean enQueue(int value) Inserts an element into the circular queue. Return true if the operation is successful.
	boolean deQueue() Deletes an element from the circular queue. Return true if the operation is successful.
	boolean isEmpty() Checks whether the circular queue is empty or not.
	boolean isFull() Checks whether the circular queue is full or not.
 */
/* 双指针head和tail
** head和tail之间的位置关系有0至k-1种情况总共k种情况
** 但是 数组队列元素有空队列，有1至k个元素总共k+1种情况
** 以上两种情况需要一一对应，所以需要浪费一个数组空间，即数组元素达到k-1个则认为队列已满，所以该方法不能提交完成所有测试用例。
** 此情况：head指针指在队首元素前一个位置，tail则指向队尾元素
	head == tail 表示空
	(head+1)%size == tail 表示满
	head与tail之间距离与队列中元素个数一一对应
** ========= 修改方法
** 鉴于 head和tail之间位置关系只有k种情况， 所以可以引入一种情况来表示空情况， 此时就可以与「队列元素存在的k+1种情况一一对应」
** 1. head == tail == -1则表示队列为空
** 2. head == tail且不等于-1 则表示队列只有一个元素
** 3. head到tail之间距离+1等于队列中元素个数
** 4. (tail+1)%size == head 则队列满，即head走到tail距离k-1，对应队列中有k个元素
 */
type CircularQueue struct {
	data    []int
	fp, bp  int
}

func ConstructorCircularQueue(k int) CircularQueue {
	this := CircularQueue{}
	this.data = make([]int, k)
	this.fp, this.bp = -1, -1 // 增加修改-1
	return this
}

func (this *CircularQueue) EnQueue(value int) bool {
	if this.IsFull() { return false }
	if this.IsEmpty() { this.fp = 0 }  // 增加修改-2
	k := len(this.data)
	this.bp = (this.bp+1)%k
	this.data[this.bp] = value
	return true
}

func (this *CircularQueue) DeQueue() bool {
	if this.IsEmpty() { return false }
	if this.fp == this.bp { // 增加修改-3
		this.fp, this.bp = -1, -1
		return true
	}
	k := len(this.data)
	this.fp = (this.fp+1)%k
	return true
}

func (this *CircularQueue) Front() int {
	if this.IsEmpty() { return -1 }
	return this.data[this.fp]
}

func (this *CircularQueue) Rear() int {
	if this.IsEmpty() { return -1 }
	return this.data[this.bp]
}

func (this *CircularQueue) IsEmpty() bool {
	return this.fp == -1 && this.bp == -1 // 增加修改-4
}

func (this *CircularQueue) IsFull() bool {
	k := len(this.data)
	return (this.bp + 1)%k == this.fp
}

/*方法二：利用计数
** 对于一个固定大小的数组，任何位置都可以是队首，只要知道队列长度，就可以根据下面公式计算出队尾位置
	tail = (head + count - 1) % capacity
 */
type MyCircularQueue struct {
	lock		sync.Mutex
	data		[]int
	head		int
	count		int
	capacity	int
}


func Constructor(k int) MyCircularQueue {
	this := MyCircularQueue{}
	this.capacity = k
	this.data = make([]int, k)
	return this
}


func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull(){ return false }
	this.lock.Lock()
	// 应该是tail
	// 求 tail 的下一个插入位置
	tail := (this.head + this.count) % this.capacity
	this.count++
	this.data[tail] = value
	this.lock.Unlock()
	return true
}


func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() { return false }
	this.lock.Lock()
	this.count--
	this.head = (this.head + 1) % this.capacity
	this.lock.Unlock()
	return true
}


func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() { return -1 }
	return this.data[this.head]
}


func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() { return  -1 }
	tail := (this.head + this.count - 1) % this.capacity
	return this.data[tail]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.count == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.count == this.capacity
}
/* 方法三：单链表实现，起始也可以使用带伪首的单链表。失去了 循环队列的考察意义
** capacity：循环队列可容纳的最大元素数量。
** head：队首元素索引。
** count：当前队列长度。该属性很重要，可以用来做边界检查。
** tail：队尾元素索引。与数组实现方式相比，如果不保存队尾索引，则需要花费 O(N) 时间找到队尾元素。
 */

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */