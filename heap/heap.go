package heap

/*
最小堆实现
 */

import "sort"

/*
  此Interface是heap的接口类型，所有基于此heap的类型，必实现此接口
  实现的类型是最小堆，并满足下面的条件
  !h.Less(j, i) for 0 <= i < h.Len() and 2*i+1 <= j <= 2*i+2 and j < h.Len()
 */
/* sort.Interface
type Interface interface {
    Len() int
    Less(i int, j int) bool
    Swap(i int, j int)
}
 */
type Interface interface {
	sort.Interface // 携带3个接口函数，Len() Swap() Less()
	Push(x interface{}) // add x as element {Len()}
	Pop() interface{} // remove and return element {Len() - 1}
}

/*
  Init establishes the heap invariants(不变式) required by the other routines in this knapsack.
  Init is idempotent(幂等性) with respect to the heap invariants and may be called whenever the heap
  invariants may have been invalidated. The Complexity is O(n) where n = h.Len()
 */

func Init(h Interface){
	// heapify 堆化
	n := h.Len()
	for i := n/2 - 1; i >= 0; i--{
		down(h, i, n)
	}
}
/*
  Push pushes the element x onto the heap
  The Complexity is O(logn) where n = h.Len()
 */
func Push(h Interface, x interface{}) {
	h.Push(x) // 调用heap具体实现的push操作
	up(h, h.Len() - 1)
}

/*
  Pop removes and returens the minimum element(according th Less)from the heap.
  The Complexity is O(logn) where n = h.Len().
  Pop is equivalent to Remove(h, 0).
 */
func Pop(h Interface) interface{}{
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop() // 调整好后 调用真正的Pop实现
}

/*
  Remove removes and returns the element at index i from the heap
  The Complexity is O(logn) where n = h.Len()
 */
func Remove(h Interface, i int) interface{} {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		if !down(h, i, n){
			up(h, i)
		}
	}
	return h.Pop()
}

/*
  Fix re-establishes the heap ordering after the element at index i has changed its value.
  Changing the value of the element at index i and then calling Fix is equivalent to,
  but less expensive than, calling Remove(h,i) followed by a Push of the new value.
  The Complexity is O(logn) where n = h.Len()
 */
func Fix(h Interface, i int){
	if !down(h, i, h.Len()) {
		up(h, i)
	}
}

func up(h Interface, j int){
	for {
		i := (j - 1) / 2 // parent
		if i == j || h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}
func down(h Interface, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0{
			break
		}
		j := j1 // left child
		if j2 := j1 +1; j2 < n && h.Less(j2, j1) {
			j = j2 // => 2*i + 2 rigth child
		}
		if !h.Less(j, i){
			break // j > i 满足最小堆, 无需调整
		}
		h.Swap(i, j) // i > j
		i = j
	}
	return i > i0  // 判断是否有调整，如果有调整，返回true，通知调用down的父函数继续检查调整
}

