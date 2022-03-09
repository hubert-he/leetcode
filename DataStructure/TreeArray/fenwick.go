package TreeArray

// 本层是 树状数组
// 树状数组，或称Binary Indexed Tree, Fenwick Tree，是一种用于高效处理对一个存储数字的列表进行更新及求前缀和的数据结构

/* 307. Range Sum Query - Mutable
** Given an integer array nums, handle multiple queries of the following types:
	1. Update the value of an element in nums.
	2. Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.
** Implement the NumArray class:
	* NumArray(int[] nums) Initializes the object with the integer array nums.
	* void update(int index, int val) Updates the value of nums[index] to be val.
	* int sumRange(int left, int right) Returns the sum of the elements of nums between indices left and right inclusive
	(i.e. nums[left] + nums[left + 1] + ... + nums[right]).
*/
/* 经典题目
** 针对不同的题目，我们有不同的方案可以选择（假设我们有一个数组）:
** 1. 数组不变，求区间和: 「前缀和」、「树状数组」、「线段树」
** 2. 多次修改某个数，求区间和：「树状数组」、「线段树」
** 3. 多次整体修改某个区间，求区间和：「线段树」、「树状数组」（看修改区间的数据范围）
** 4. 多次将某个区间变成同一个数，求区间和：「线段树」、「树状数组」（看修改区间的数据范围）
** 因为「线段树」代码很长，而且常数很大，实际表现不算很好
** 只有在我们遇到第 4 类问题，不得不写「线段树」的时候，我们才考虑线段树
** 总结一下，我们应该按这样的优先级进行考虑：
	1. 简单求区间和，用「前缀和」
	2. 多次将某个区间变成同一个数，用「线段树」
	3. 其他情况，用「树状数组」
*/

// 方法一： 前缀和
/* 2022-02-25 刷出此题， 注意update 的错误
type NumArray struct {
	prefixsum []int
}
func Constructor(nums []int) NumArray {
	this := NumArray{}
	this.prefixsum = make([]int, len(nums)+1)
	for i := range nums{
		this.prefixsum[i+1] = this.prefixsum[i] + nums[i]
	}
	return this
}
func (this *NumArray) Update(index int, val int)  {
	if index > len(this.prefixsum){ return }
	diff := val - this.prefixsum[index+1] + this.prefixsum[index]  // 这边容易犯糊涂
	for j := index+1; j < len(this.prefixsum); j++{
		this.prefixsum[j] += diff
	}
}
func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixsum[right+1] - this.prefixsum[left]
}
 */

// 方法二：树状数组
/* 定义树状数组
 */
type NumArray struct {
	tree	[]int
	nums	[]int
}
// 利用负数 取反+1 特性，查找 x 的低位方向第一个为1的位置
func (this *NumArray) lowbit(x int)int{
	return x & -x
}
func (this *NumArray) query(x int)(result int) {
	for i := x; i > 0; i -= this.lowbit(i){
		result += this.tree[i]
	}
	return
}
func (this *NumArray) add(x, u int){
	for i := x; i <= len(this.nums); i += this.lowbit(i) {
		this.tree[i] += u
	}
}
func Constructor(nums []int) NumArray {
	n := len(nums)
	treeArr := NumArray{}
	treeArr.nums = nums
	treeArr.tree = make([]int, n+1)
	for i := range nums{// 初始化树状数组
		treeArr.add(i+1, nums[i])
	}
	return treeArr
}
func (this *NumArray) Update(index int, val int)  {
	this.add(index+1, val-this.nums[index])
	this.nums[index] = val
}
func (this *NumArray) SumRange(left int, right int) int {
	return this.query(right+1) - this.query(left)
}

















