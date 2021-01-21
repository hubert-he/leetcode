package design
import "container/heap"
/*703. 数据流中的第 K 大元素
  设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。
请实现 KthLargest 类：
KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。
示例：
输入：
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
输出：
[null, 4, 5, 5, 8, 8]
解释：
KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
kthLargest.add(3);   // return 4
kthLargest.add(5);   // return 5
kthLargest.add(10);  // return 5
kthLargest.add(9);   // return 8
kthLargest.add(4);   // return 8
链接：https://leetcode-cn.com/problems/kth-largest-element-in-a-stream
 */
type IntHeap []int
type KthLargest struct {
	index int
	nums []int
	miniHeap *IntHeap
}
func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (hp *IntHeap) Push(x interface{}){
	*hp = append(*hp, x.(int))
}

func (hp *IntHeap) Pop() interface{} {
	ret := (*hp)[len(*hp) - 1]
	*hp = (*hp)[:len(*hp)-1]
	return ret
}
func (h IntHeap) Peek() interface{}{
	return h[0]
}

func Constructor(k int, nums []int) KthLargest {
	int_heap := &IntHeap{}
	for i := 0; i < len(nums); i++{
		heap.Push(int_heap, nums[i])
		if int_heap.Len() > k{
			heap.Pop(int_heap)
		}
	}
	return KthLargest{index: k, nums: nums, miniHeap: int_heap}
}


func (this *KthLargest) Add(val int) int {
	this.nums = append(this.nums, val)
	heap.Push(this.miniHeap, val)
	if this.miniHeap.Len() > this.index{
		heap.Pop(this.miniHeap)
	}
	return this.miniHeap.Peek().(int)
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */