package Stack

import (
	"math"
	"math/rand"
)

/* 496. Next Greater Element I
** The next greater element of some element x in an array is the first greater element
	that is to the right of x in the same array.
** You are given two distinct 0-indexed integer arrays nums1 and nums2, where nums1 is a subset of nums2.
** For each 0 <= i < nums1.length, find the index j such that nums1[i] == nums2[j] and
	determine the next greater element of nums2[j] in nums2. If there is no next greater element,
	then the answer for this query is -1.
** Return an array ans of length nums1.length such that ans[i] is the next greater element as described above.
** Follow up: Could you find an O(nums1.length + nums2.length) solution?
 */
// 2022-02-15 刷出此题， 借助单调栈
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n := len(nums2)
	m := map[int]int{nums2[n-1]: -1}
	st := []int{nums2[n-1]} // 单调栈
	for i := n-2; i >= 0; i--{
		top := st[len(st)-1]
		if nums2[i] < top{
			m[nums2[i]] = top
		}else if nums2[i] >= top{
			st = st[:len(st)-1]
			m[nums2[i]] = -1
			for len(st) > 0{
				if nums2[i] > st[len(st)-1]{
					st = st[:len(st)-1]
				}else{
					m[nums2[i]] = st[len(st)-1]
					break
				}
			}
		}
		st = append(st, nums2[i])
	}
	ans := []int{}
	for _, c := range nums1{
		ans = append(ans, m[c])
	}
	return ans
}
// 官方题解
func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	stack := []int{}
	for i := len(nums2) - 1; i >= 0; i--{
		num := nums2[i]
		for len(stack) > 0 && num >= stack[len(stack)-1]{
			stack = stack[:len(stack) - 1]
		}
		if len(stack) > 0{
			m[num] = stack[len(stack)-1]
		}else{
			m[num] = -1
		}
		stack = append(stack, num)
	}
	ans := make([]int, len(nums1))
	for i, num := range nums1{
		ans[i] = m[num]
	}
	return ans
}

/* 503. Next Greater Element II
** Given a circular integer array nums (i.e., the next element of nums[nums.length - 1] is nums[0]),
** return the next greater number for every element in nums.
** The next greater number of a number x is the first greater number to its traversing-order next in the array,
** which means you could search circularly to find its next greater number.
** If it doesn't exist, return -1 for this number.
** Constraints:
	1 <= nums.length <= 10^4
	-10^9 <= nums[i] <= 10^9
 */
//2022-02-16 失败5次后 刷出，测试案例 前5个
func NextGreaterElements(nums []int) []int {
	n := len(nums)
	st := []int{} // 单调栈
	ans := make([]int, n)
	for i := n-1; i >= 0; i--{
		for len(st) > 0{
			if nums[i] >= st[len(st)-1]{
				st = st[:len(st)-1]
			}else{
				ans[i] = st[len(st)-1]
				break
			}
		}
		if len(st) == 0{
			ans[i] = math.MinInt32
		}
		st = append(st, nums[i])
	}
	for i := range nums{
		if ans[i] == math.MinInt32{
			for j := 0; j < i; j++{
				if nums[j] > nums[i]{
					ans[i] = nums[j]
					break
				}
			}
			if ans[i] == math.MinInt32{
				ans[i] = -1
			}
		}
	}
	return ans
}
// 官方单调栈
/* 单调栈
** 单调栈中保存数组下标，从栈底到栈顶的下标在数组 nums 中对应的值是单调不升的
** 每次我们移动到数组中的一个新的位置 i , 我们就将当前单调栈中所有对应值小于 nums[i]的下标弹出，
** 这些值的下一个更大元素即为 nums[i] （证明很简单：如果有更靠前的更大元素，那么这些位置将被提前弹出栈）
** 随后我们将位置 i 入栈
** 但是注意到只遍历一次序列是不够的， 例如序列 [2,3,1]，最后单调栈中将剩余[3,1] 其中元素 [1] 的下一个更大未知
** 此时利用循环数组处理
 */
/* 如何实现循环数组
** 循环数组的最后一个元素下一个元素是数组的第一个元素，形状类似于「环」
** 一种实现方式是，把数组复制一份到数组末尾，这样虽然不是严格的循环数组，但是对于本题已经足够了，因为本题对数组最多遍历两次
** 另一个常见的实现方式是，使用取模运算 % 可以把下标 i 映射到数组 nums 长度的 0 - N 内
 */
func NextGreaterElements2(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	st := []int{}
	// 把两个相同数组粘在一起
	for i := 0; i < n*2-1; i++{
		for len(st) > 0 && nums[st[len(st)-1]] < nums[i%n]{
			ans[st[len(st)-1]] = nums[i%n]
			st = st[:len(st)-1]
		}
		st = append(st, i%n)
	}
	return ans
}

/* 739. Daily Temperatures
** Given an array of integers temperatures represents the daily temperatures,
** return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature.
** If there is no future day for which this is possible, keep answer[i] == 0 instead.
 */
// 2022-02-17 刷出此题，解法使用的是逆向遍历， 可学习 正向遍历思路
// 此题与 上面的 题目的 不同点是， 此题必须存下标，计算距离
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	st := []int{}// 记录下标
	ans := make([]int, n)
	for i := n-1; i >= 0; i--{
		if len(st) == 0{
			ans[i] = 0
			st = append(st, i)
		}else{
			for len(st) > 0{
				if temperatures[st[len(st)-1]] > temperatures[i]{
					ans[i] = st[len(st)-1] - i
					break
				}else{
					st = st[:len(st)-1]
				}
			}
			if len(st)==0{
				ans[i] = 0
			}
			st = append(st, i)
		}
	}
	return ans
}
// 正向遍历, 此外可复用 temperature 数组
func dailyTemperatures2(temperatures []int) []int {
	st:= []int{}// 必须存下标
	for i := range temperatures {
		for len(st) > 0 &&  temperatures[i] > temperatures[st[len(st)-1]]{
			temperatures[st[len(st)-1]] = i - st[len(st)-1]
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	for len(st) > 0{
		temperatures[st[len(st)-1]] = 0
		st = st[:len(st)-1]
	}
	return temperatures
}

/* 84. Largest Rectangle in Histogram
** Given an array of integers heights representing the histogram's bar height where the width of each bar is 1,
** return the area of the largest rectangle in the histogram.
 */
// 这个是枚举宽的方式
func largestRectangleArea_burst_width(heights []int) int {
	n := len(heights)
	ans := math.MinInt32
	for i := 0; i < n; i++{
		h := math.MaxInt32
		for j := i; j < n; j++{
			if h > heights[j]{
				h = heights[j]
			}
			if ans < (j-i+1)*h {
				ans = (j-i+1)*h
			}
		}
	}
	return ans
}
/* 枚举高的思路
** 使用一重循环枚举某一根柱子， 将其固定为矩形的高度 h
** 随后我们从这跟柱子开始向两侧延伸，直到遇到高度小于h的柱子，就确定了矩形的左右边界
** 从而得出答案
 */
func largestRectangleArea_burst_height(heights []int) int {
	n := len(heights)
	ans := 0
	for i, h := range heights{
		left, right := i, i
		for left > 0 && heights[left-1] >= h{
			left--
		}
		for right < n-1 && heights[right+1] >= h{
			right++
		}
		t := (right-left+1) * h
		if ans < t {
			ans = t
		}
	}
	return ans
}
// 由枚举高开始思路优化， 就可以发现 可借用单调栈 来快速返回 left  right， 达到 O(n) 思路
func largestRectangleArea(heights []int) int {
	const Left, Right = 0, 1
	n := len(heights)
	location := make([][2]int, n)
	st := []int{} // 存索引
	for i, h := range heights{
		for len(st) > 0 && heights[len(st)-1] >= h{
			st = st[:len(st)-1]
		}
		if len(st) == 0{
			location[i][Left] = -1
		}else{
			location[i][Left] = st[len(st)-1]
		}
		st = append(st, i)
	}
	for i := n-1; i >= 0; i--{
		for len(st) > 0 && heights[len(st)-1] >= heights[i]{
			st = st[:len(st)-1]
		}
		if len(st) == 0{
			location[i][Right] = n
		}else{
			location[i][Right] = st[len(st)-1]
		}
		st = append(st, i)
	}
	ans := 0
	for i, h := range heights{
		left, right := location[i][Left], location[i][Right]
		t := h * (right-left)
		if ans < t {
			ans = t
		}
	}
	return ans
}

/* 456. 132 Pattern
** Given an array of n integers nums,
** a 132 pattern is a subsequence of three integers nums[i], nums[j] and nums[k]
** such that i < j < k and nums[i] < nums[k] < nums[j].
** Return true if there is a 132 pattern in nums, otherwise, return false.
 */

/* 此题的最优解释 单调栈
** 枚举 3 是容易想到并且也是最容易实现的，太辣鸡，竟然没想到
** 注意下面的 2 点优化
** 由于3是模式中的最大值，并且其出现在 1 和 2 之间，因此只需从左到右枚举 3 的下标 j ：
	1. 由于 1 是模式中的最小值，因此我们在枚举 j 的同时，维护数组 a 中左侧元素 a[0:j-1]的最小值 <== 这个最小值没想到，题目实际上是找是否true
		即为 1 对应的元素 a[i]。只有 a[i] < a[j]时，a[i] 才能作为 1 对应的元素
	2. 由于 2 是 模式中的次小值，因此我们可以使用一个有序集合(平衡树)维护数组 a 中 右侧元素 a[j+1:] 中所有值
		当我们确定了 a[i] 与 a[j] 关系后，只要在有序集合中查询严格比 a[i] 大的 那个最小的元素，即为 a[k]
		需要注意的是，只有 a[k] < a[j] 时， a[k] 才能作为 3 对应的元素。
 */
// 立方级别
func find132pattern_tle(nums []int) bool {
	n := len(nums)
	for i := range nums{
		left, right := []int{}, []int{}
		for j := 0; j < i; j++{
			if nums[j] < nums[i]{
				left = append(left, nums[j])
			}
		}
		for j := i+1; j < n; j++{
			if nums[j] < nums[i]{
				right = append(right, nums[j])
			}
		}
		if len(left) == 0 || len(right) == 0{
			continue
		}
		for _, l := range left{
			for _, r := range right{
				if l < r{ return true }
			}
		}
	}
	return false
}
// 平方级别
func find132pattern_improve_tle(nums []int) bool {
	n := len(nums)
	left := math.MaxInt32
	for i := range nums{
		// 因为只需要查找是否存在，因此left 不需要保留那么多，只需要保证最下的元素下标即可
		/*left, right := []int{}, []int{}
		for j := 0; j < i; j++{
			if nums[j] < nums[i]{
				left = append(left, nums[j])
			}
		}
		 */
		if i > 0 && nums[i-1] < left{
			left = nums[i-1]
		}

		for j := i+1; j < n; j++{
			/*
			if nums[j] < nums[i]{
				right = append(right, nums[j])
			}*/
			if nums[j] < nums[i] && nums[j] > left {
				return true
			}
		}
		/*
		if len(left) == 0 || len(right) == 0{
			continue
		}
		for _, l := range left{
			for _, r := range right{
				if l < r{ return true }
			}
		}
		 */
	}
	return false
}
// 上面是 平方级别，必须要优化内循环 由O(n) 变为 O(logn) 才行，
// 因此使用 平衡树来存储 right
func find132pattern_improve_treap(nums []int) bool {
	n := len(nums)
	if n < 3 {  return false }
	left := nums[0]
	// 构造 treap
	rights := &treap{}
	for _, v := range nums[2:] {
		rights.put(v)
	}
	for j := 1; j < n-1; j++ {
		if nums[j] > left{
			ub := rights.upperBound(left)
			if ub != nil && ub.val < nums[j]{
				return true
			}
		}else{// 维护left 最小
			left = nums[j]
		}
		rights.delete(nums[j+1])// 删掉马上要访问的
	}
	return false
}
const Left, Right = 0, 1

type  node struct{
	ch 			[2]*node
	priority	int // 要满足堆性质
	val 		int // 要满足二叉搜索树的性质
	cnt 		int
}

func(o *node) cmp(b int)int{
	switch {
	case b < o.val:
		return Left
	case b > o.val:
		return Right
	default: // 不能相等
		return -1
	}
}

func(o *node)rotate(d int)*node{ // 左右旋合并到一起了
	x := o.ch[d^1]
	o.ch[d^1] = x.ch[d]
	x.ch[d] = o
	return x
}

type treap struct {
	root	*node
}

func (t *treap) _put(o *node, val int) *node{
	if o == nil {
		return &node{priority: rand.Int(), val: val, cnt: 1}
	}
	if d := o.cmp(val); d >= 0{
		o.ch[d] = t._put(o.ch[d], val)
		if o.ch[d].priority > o.priority{
			o = o.rotate(d ^ 1)
		}
	}else{ // 重复的key
		o.cnt++
	}
	return o
}

func(t *treap) put(val int){
	t.root = t._put(t.root, val) // 从根节点开始查询插入
}
// BST 的删除操作
func(t *treap)_delete(o *node, val int)*node{
	if o == nil{ return nil }
	if d := o.cmp(val); d >= 0 {
		o.ch[d] = t._delete(o.ch[d], val)
		return o
	}
	// d == -1 相等的
	if o.cnt > 1{
		o.cnt--
		return o
	}
	if o.ch[Right] == nil{
		return o.ch[Left]
	}
	if o.ch[Left] == nil{
		return o.ch[Right]
	}
	d := Left
	if o.ch[Left].priority > o.ch[Right].priority{
		d = Right
	}
	o = o.rotate(d)
	o.ch[d] = t._delete(o.ch[d], val)
	return o
}

func (t *treap) delete(val int){
	t.root = t._delete(t.root, val)
}

func(t *treap) upperBound(val int)(ub *node){
	for o := t.root; o != nil;{
		if o.cmp(val) == 0{ // o.val > val
			ub = o
			o = o.ch[0]
		}else{
			o = o.ch[1]
		}
	}
	return
}
