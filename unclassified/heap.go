package unclassified

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"
	"time"
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
	/*
	case1 := []int{3,2,1,5,6,4}
	case2 := []int{3,2,3,1,2,4,5,5,6}
	fmt.Println(findKthLarges_1(case1, 2))
	fmt.Println(findKthLarges_1(case2, 4))

	case3 := []int{1,1,1,2,2,3}
	fmt.Println("topKFrequent=> ", topKFrequent(case3, 2))
	fmt.Println("topKFrequent_bucket => ", topKFrequent_bucket(case3, 2))
	fmt.Println("topKFrequent_qsort => ", topKFrequent_qsort(case3, 2))
	PriorityRun()
	*/
	result := reorganizeString("aaababaacbb")
	fmt.Println("recorganizaString => ", result)
	result = reorganizeString_counting("aaababaacbb")
	fmt.Println("reorganizeString_counting => ", result)
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

type frequentInt struct {
	num int
	freq int
}
type freqHeap []frequentInt
func (fh freqHeap) Len() int{
	return len(fh)
}
func (fh freqHeap) Less(i, j int) bool {
	return fh[i].freq < fh[j].freq
}
func (fh freqHeap) Swap(i, j int) {
	fh[i], fh[j] = fh[j], fh[i]
}
func (fhp *freqHeap) Push(x interface{}){
	*fhp = append(*fhp, x.(frequentInt))
}
func (fhp *freqHeap) Pop() interface{}{
	ret := (*fhp)[len(*fhp) - 1]
	*fhp = (*fhp)[:len(*fhp) - 1]
	return ret
}
func topKFrequent(nums []int, k int) []int {
	h := &freqHeap{}
	heap.Init(h)
	freqMap := map[int]int{}
	for i := 0; i < len(nums); i++{
		freqMap[nums[i]]++
	}
	for key,value := range freqMap{
		heap.Push(h, frequentInt{num: key, freq: value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := []int{}
	for i := 0; i < h.Len(); i++{
		ret = append(ret, (*h)[i].num)
	}
	return ret
}
func topKFrequent_bucket(nums []int, k int) []int{
	freqMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		freqMap[nums[i]]++
	}
	numsBucket := make([][]int, len(nums)+1)
	for key, value := range freqMap{
		numsBucket[value] = append(numsBucket[value], key)
	}
	ret := []int{}
	for i := len(numsBucket) - 1; i > 0; i--{
		if k <= 0{
			break
		}
		if numsBucket[i] != nil{
			for j := 0; j < len(numsBucket[i]); j++{
				if k <= 0{
					break
				}
				ret = append(ret, numsBucket[i][j])
				k--
			}
		}
	}
	return ret
}

func topKFrequent_qsort(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	values := [][]int{}
	for key, value := range occurrences {
		values = append(values, []int{key, value})
	}
	ret := make([]int, k)
	qsort(values, 0, len(values) - 1, ret, 0, k)
	return ret
}

func qsort(values [][]int, start, end int, ret []int, retIndex, k int) {
	rand.Seed(time.Now().UnixNano())
	picked := rand.Int() % (end - start + 1) + start;
	values[picked], values[start] = values[start], values[picked]

	pivot := values[start][1]
	index := start
	for i := start + 1; i <= end; i++ {
		if values[i][1] >= pivot {
			values[index + 1], values[i] = values[i], values[index + 1]
			index++
		}
	}
	values[start], values[index] = values[index], values[start]
	if k <= index - start {// 左侧多
		qsort(values, start, index - 1, ret, retIndex, k)
	} else {// 不足或正好k个
		for i := start; i <= index; i++ {
			ret[retIndex] = values[i][0]
			retIndex++
		}
		if k > index - start + 1 { // 不足K个
			qsort(values, index + 1, end, ret, retIndex, k - (index - start + 1))
		}
	}
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

var cnt [26]int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return cnt[h.IntSlice[i]] > cnt[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hp) push(v int)         { heap.Push(h, v) }
func (h *hp) pop() int           { return heap.Pop(h).(int) }

func reorganizeString(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	cnt = [26]int{}
	maxCnt := 0
	for _, ch := range s {
		ch -= 'a'
		cnt[ch]++
		if cnt[ch] > maxCnt {
			maxCnt = cnt[ch]
		}
	}
	if maxCnt > (n+1)/2 {
		return ""
	}

	h := &hp{}
	for i, c := range cnt[:] {
		if c > 0 {
			h.IntSlice = append(h.IntSlice, i)
		}
	}
	heap.Init(h)

	ans := make([]byte, 0, n)
	fmt.Println(h.IntSlice)
	for len(h.IntSlice) > 1 {
		i, j := h.pop(), h.pop()
		fmt.Println(i,j)
		ans = append(ans, byte('a'+i), byte('a'+j))
		if cnt[i]--; cnt[i] > 0 {
			h.push(i)
		}
		if cnt[j]--; cnt[j] > 0 {
			h.push(j)
		}
	}
	if len(h.IntSlice) > 0 {
		ans = append(ans, byte('a'+h.IntSlice[0]))
	}
	return string(ans)
}
/*
  当 nn 是奇数且出现最多的字母的出现次数是(n+1)/2 时，出现次数最多的字母必须全部放置在偶数下标，否则一定会出现相邻的字母相同的情况。
  其余情况下，每个字母放置在偶数下标或者奇数下标都是可行的。
  维护偶数下标evenIndex和奇数下标oddIndex，初始值分别为0和1。遍历每个字母，根据每个字母的出现次数判断字母应该放置在偶数下标还是奇数下标。
  首先考虑是否可以放置在奇数下标。只要字母的出现次数不超过字符串的长度的一半（即出现次数小于或等于n/2），就可以放置在奇数下标，只有当字母的出现次数超过字符串的长度的一半时，才必须放置在偶数下标
  字母的出现次数超过字符串的长度的一半只可能发生在n是奇数的情况下，且最多只有一个字母的出现次数会超过字符串的长度的一半
 */
func reorganizeString_counting(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	cnt := [26]int{}
	maxCnt := 0
	for _, ch := range s{
		cnt[ch - 'a']++
		if maxCnt < cnt[ch - 'a']{
			maxCnt = cnt[ch - 'a']
		}
	}
	fmt.Printf("%#v\n", cnt)
	if maxCnt > (n+1) / 2 {
		return ""
	}
	answer := make([]byte, n)
	evenIdx, oddIdx, halfLen := 0, 1, n/2
	for i, c := range cnt{
		ch := byte('a' + i)
		for c > 0 && c <= halfLen && oddIdx < n{
			answer[oddIdx] = ch
			c--
			oddIdx += 2
		}
		for c > 0 { // 出现次数大于 n/2 或者奇数位置放满了
			answer[evenIdx] = ch
			c--
			evenIdx += 2
		}
	}
	return string(answer)
}