package Find_Median_from_Data_Stream_295

import (
	"container/heap"
	"fmt"
	"sort"
)
// 只有内嵌的匿名结构体，子结构体可以继承其方法和类型， 注意其引用方式，直接通过类型操作
type intHeap struct {
	sort.IntSlice
}
func (h *intHeap) Push(i interface{}){
	h.IntSlice = append(h.IntSlice, i.(int))
}

func (h *intHeap) Pop() interface{}{
	n := h.IntSlice.Len()
	tmp := h.IntSlice[n - 1]
	h.IntSlice = h.IntSlice[:n-1]
	return tmp
}
func (h *intHeap) push(i interface{}){
	heap.Push(h, i)
}
func (h *intHeap) pop() int{
	return heap.Pop(h).(int)
}
type BigHeap struct {
	intHeap
}
func (h *BigHeap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

type MedianFinder struct {
	data 	[]int
	low 	BigHeap
	high 	intHeap
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}


func (this *MedianFinder) AddNum(num int) {
	//this.addNum_bs(num)
	this.addNum_dheap(num)
}

func (this *MedianFinder) addNum_dheap(num int) {
	this.low.push(num)
	for this.low.Len() > this.high.Len() {
		x := this.low.pop()
		this.high.push(x)
	}
}

func (this *MedianFinder) addNum_bs(num int) {
	n := len(this.data)
	left,right,loc := 0,n-1,0
	for left <= right{
		mid := left + (right - left) >> 1 // 优化1：>> 替换
		fmt.Println("mid => ", mid)
		if this.data[mid] < num{  // Mid < num
			loc = mid + 1
			left = mid + 1  // Mid < num的时候 loc 不赋值，target大于数组所有元素情况 直接尾部追加
		}else{	// Mid >= num
			loc = mid
			right = mid - 1
		}
	}
	fmt.Println(left, right, loc, this.data)
	this.data = append(this.data[:loc], append([]int{num}, this.data[loc:]...)...)
}

func (this *MedianFinder) FindMedian() float64 {
	return  this.FindMedian_dheap()
}
func (this *MedianFinder) FindMedian_dheap() float64 {
	if this.high.Len() == 0 {
		return 0
	}
	fmt.Println("len ", this.low.Len(), this.high.Len(), this.low.IntSlice, this.high.IntSlice)
	if this.high.Len() > this.low.Len(){
		return (float64)(this.high.IntSlice[0])
	}else{
		l := this.low.IntSlice[0]
		h := this.high.IntSlice[0]
		return (float64)(l + h) / (float64)(2)
	}
}
func (this *MedianFinder) FindMedian_bs() float64 {
	n := len(this.data)
	fmt.Println(this.data)
	if n < 1{
		return 0
	}
	if n % 2 == 0{
		r := n/2
		l := n/2 - 1
		return (float64)(this.data[r] + this.data[l]) / (float64)(2)
	}else {
		return (float64)(this.data[(n-1)/2])
	}
}


func searchInsert(nums []int, target int) int {
	n := len(nums)
	// loc 初始为0，应对nums为[]的情况
	left, right, loc := 0, n-1,0
	for left <= right{
		mid := left + (right - left) >> 1
		if nums[mid] < target{
			loc = mid + 1
			left = mid + 1
		}else {
			loc = mid
			right = mid - 1
		}
	}
	return loc
}

func searchInsert_advance(nums []int, target int)(index int){
	n := len(nums)
	index = n // 优化：初值设置为数组长度可以省略边界条件的判断，因为存在一种情况是target大于数组所有元素，包含[] 和 确实比数组中的都大 这2中情况
	left, right := 0, n-1
	for left <= right{
		mid := left + (right - left) >> 1
		if nums[mid] < target{
			index = mid + 1 // 此语句可省略，当默认index初始化为数组长度的时候
			left = mid + 1
		}else {//有序数组中找第一个大于等于target的下标
			index = mid
			right = mid - 1
		}
	}
	return
}

