package queue
/* 232. Implement Queue using Stacks
** Implement a first in first out (FIFO) queue using only two stacks.
** The implemented queue should support all the functions of a normal queue (push, peek, pop, and empty).
** Implement the MyQueue class:
	void push(int x) Pushes element x to the back of the queue.
	int pop() Removes the element from the front of the queue and returns it.
	int peek() Returns the element at the front of the queue.
	boolean empty() Returns true if the queue is empty, false otherwise.
Notes:
	You must use only standard operations of a stack,
	which means only push to top, peek/pop from top, size, and is empty operations are valid.
	Depending on your language, the stack may not be supported natively.
	You may simulate a stack using a list or deque (double-ended queue) as long as you use only a stack's standard operations.
 */
type MyQueue struct {
	st1 	[]int
	st2 	[]int
	front 	int
}


func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int)  {
	if len(this.st1) == 0{
		this.front = x
	}
	this.st1 = append(this.st1, x)
}

func (this *MyQueue) Pop() int {
	if len(this.st2) == 0{
		for i := len(this.st1)-1; i >= 0; i--{
			this.st2 = append(this.st2, this.st1[i])
		}
		this.st1 = []int{}
	}
	return this.st2[len(this.st2)-1]
}

func (this *MyQueue) Peek() int {
	if len(this.st2) > 0{
		return this.st2[len(this.st2)-1]
	}
	return this.front
}

func (this *MyQueue) Empty() bool {
	return len(this.st1) <= 0 && len(this.st2) <= 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */