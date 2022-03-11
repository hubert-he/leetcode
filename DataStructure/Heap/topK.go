package Heap

import (
	"container/heap"
	"math/rand"
)
/* Top K 问题 除了可以用堆
** 还可以应用 快速排序 来定位前 k 个
 */

/* 973. K Closest Points to Origin
** Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane and an integer k,
** return the k closest points to the origin (0, 0).
** The distance between two points on the X-Y plane is the Euclidean distance (i.e., √(x1 - x2)2 + (y1 - y2)2).
** You may return the answer in any order. The answer is guaranteed to be unique (except for the order that it is in).
 */

/* 关于Golang 类型定义复用的问题
**  一个例子
	type MySlice sort.IntSlice
	func (ms MySlice) ReverseSort() {
		sort.Sort(sort.Reverse(ms))
	}
	func main() {
		t2 := MySlice{5, 4, 3, 1}
		t2.ReverseSort()
		fmt.Println(t2)
	}
	报错：cannot use ms (type MySlice) as type sort.Interface in argument to sort.Reverse:
        MySlice does not implement sort.Interface (missing Len method)
** You "cannot define new methods on non-local type[s]," by design.
   The best practice is to embed the non-local type into your own own local type, and extend it.
   Type-aliasing (type MyFoo Foo) creates a type that is (more-or-less) completely distinct from the original.
   I'm not aware of a straightforward/best-practice way to use type assertions to get around that.
** A type's methods are in the package that defines it.
	This is a logical coherence（逻辑连贯）. It is a compilation virtue（编译优化？）.
	It is seen as an important benefit for large-scale maintenance and multi-person development projects.
	The power you speak of is not lost, though,
	because you can embed the base type in a new type as described above and add whatever you want to it,
	in the kind of functional "is a" that you seek, with the only caveat（警告） that your new type must have a new name,
	and all of its fields and methods must be in its new package.

** type new-type-name  old-type-name 这个语义在Golang里是 重命名，他的方法属于定义的package
** 正确的做法-1： 定义一个新类型，如此可继承获得方法
** type MySlice struct {
    	sort.IntSlice
	}
	func (ms MySlice) ReverseSort() {
    	sort.Sort(sort.Reverse(ms))
	}

** 正确做法-2: 直接在 上面补上 sort.IntSlice 方法的声明实现
 */
/*
type distance sort.IntSlice  <== 不能直接使用
func (h *distance)Push(x interface{}){
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}
// 直接复用类型定义也是不允许的
// 编译错误：Cannot define new methods on non-local type 'sort.IntSlice'
func (h *sort.IntSlice)Pop()interface{}{
	ret := h[h.Len()-1]
	*h = *h[:h.Len()-1]
	return ret
}

func kClosest(points [][]int, k int) [][]int {
	ans := [][]int{}
	h := &distance{}
	heap.Init(h)
	for _, p := range points{
		dist := p[0]*p[0] + p[1]*p[1]
		heap.Push(h, dist)
	}
	for h.Len() > 0 && k > 0{
		ans = append(ans, heap.Pop(h))
		k--
	}
	return ans
}
*/
type distance [][2]int
func (h distance)Len()int{ return len(h) }
func (h distance)Less(i, j int)bool {
	return h[i][0]*h[i][0] + h[i][1]*h[i][1] < h[j][0]*h[j][0] + h[j][1]*h[j][1]
}
func (h distance)Swap(i, j int){ h[i], h[j] = h[j], h[i] }
func (h *distance)Push(x interface{}){
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([2]int))
}
func (h *distance)Pop()interface{}{
	ret := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return ret
}
// Golang 使用 Heap 此方式是 全堆
func kClosest(points [][]int, k int) [][]int {
	ans := [][]int{}
	h := &distance{}
	heap.Init(h)
	for _, p := range points{
		heap.Push(h, [2]int{p[0], p[1]})
	}
	for h.Len() > 0 && k > 0{
		item := heap.Pop(h).([2]int)
		ans = append(ans, []int{item[0], item[1]})
		k--
	}
	return ans
}

type distSlice [][2]int
func (h distSlice)Len()int{ return len(h) }
func (h distSlice)Less(i, j int)bool { //大根堆
	return h[i][0]*h[i][0] + h[i][1]*h[i][1] > h[j][0]*h[j][0] + h[j][1]*h[j][1]
	//return h[i][0]*h[i][0] + h[i][1]*h[i][1] < h[j][0]*h[j][0] + h[j][1]*h[j][1]
}
func (h distSlice)Swap(i, j int){ h[i], h[j] = h[j], h[i] }
func (h *distSlice)Push(x interface{}){
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([2]int))
}
func (h *distSlice)Pop()interface{}{
	ret := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return ret
}
// 用大根堆 来优化的 策略  head.Fix 重点学习
// 固定 k 大根堆，然后 后面fou 循环预判
func kClosest2(points [][]int, k int) [][]int {
	h := make(distSlice, k)
	// 初始化
	for i, p := range points[:k]{
		h[i] = [2]int{p[0], p[1]}
	}
	heap.Init(&h) // O(k) 初始化 heap
	for _, p := range points[k:]{
		// 提前预判，因为已经是 k-大根堆
		if p[0]*p[0] + p[1]*p[1] < h[0][0]*h[0][0] + h[0][1]*h[0][1]{
			h[0] = [2]int{p[0], p[1]}
			heap.Fix(&h, 0)
		}
	}
	ans := [][]int{}
	// 题目要求不计顺序
	for _, p := range h{
		ans = append(ans, []int{p[0], p[1]})
	}
	return ans
}

// 方法三：借助快排
func kClosest_fastSort(points [][]int, k int) (ans [][]int) {
	n := len(points)
	rand.Shuffle(n, func(i, j int){
		points[i], points[j] = points[j], points[i]
	})
	less := func(p, q []int)bool{
		return p[0]*p[0]+p[1]*p[1] < q[0]*q[0]+q[1]*q[1]
	}
	var quickSelect func(left, right int)
	quickSelect = func(left, right int){
		pivot := points[right] // 取当前区间 [left,right] 最右侧元素作为切分参照
		lessCount := left
		for i := left; i < right; i++{
			if less(points[i], pivot){
				points[i], points[lessCount] = points[lessCount], points[i]
				lessCount++
			}
		}
		// 循环结束后，有 lessCount 个元素比 pivot 小, 把 pivot 交换到 points[lessCount] 的位置
		// 交换之后，points[lessCount] 左侧的元素均「小于」pivot，points[lessCount]右侧的元素均「不小于」pivot
		points[right], points[lessCount] = points[lessCount], points[right]
		if lessCount + 1 == k{ // 说明pivot 就是第 k 个距离最小的点
			return
		}else if lessCount+1 < k{
			quickSelect(lessCount+1, right) // 继续排右边的
		}else{
			quickSelect(left, lessCount-1) // 排左边的
		}
	}
	// 快排出 前 k 个
	quickSelect(0, n-1)
	return points[:k]
}














