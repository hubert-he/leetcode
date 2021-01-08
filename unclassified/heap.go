package unclassified

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i]}

func (h *IntHeap) Push(x interface{}){
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	ret := (*h)[len(*h) - 1]
	*h = (*h)[:len(*h) - 1]
	return ret
}
func (h *IntHeap) Peek() interface{} {
	return (*h)[0]
}

func HeapRun() {
	case1 := []int{3,2,1,5,6,4}
	case2 := []int{3,2,3,1,2,4,5,5,6}
	fmt.Println(findKthLarges_1(case1, 2))
	fmt.Println(findKthLarges_1(case2, 4))

	PriorityRun()
}

func findKthLarges_1(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)
	for i := 0; i < len(nums); i++ {
		heap.Push(h, nums[i])
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return h.Peek().(int) // 类型断言
}


// Priority Queue

/*
 An Item is sth we manage in a priority queue
 */
type Item struct {
	value 		string
	priority 	int		// The Priority of the item in the queue
	index		int		// The index of the item in the heap, which is needed by update and is maintained by the heap. Interface Method
}

// A PriorityQueue implements heap.Interface and holds Items
type PriorityQueue []*Item
func (pq PriorityQueue) Len() int { return len(pq)}
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i // 注意index
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
func (pq PriorityQueue) Peek() interface{} {
	return pq[0]
}
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func PriorityRun() {
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items{
		pq[i] = &Item{
			value: value,
			priority: priority,
			index: i,
		}
		i++
	}
	heap.Init(&pq)

	item := &Item{value: "orange", priority: 1}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
}