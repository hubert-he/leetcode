package Tree

import (
	"fmt"
	"math"
	"sort"
)

type SegTreeNode struct {
	start	int
	end		int
	mark	int
	data	interface{}
}

func (pn *SegTreeNode) AddMark(value int){
	pn.mark += value
}

func (pn *SegTreeNode) ClearMark(){
	pn.mark = 0
}

func (n SegTreeNode) String() string{
	return fmt.Sprintf("%d - %d\n", n.start, n.end)
}

func NewSegTreeNode(left, right int) *SegTreeNode{
	return &SegTreeNode{start: left, end: right}
}

type SegTree struct {
	element []SegTreeNode
}
func NewSegTree(base []int) *SegTree{
	length := len(base)
	if length <= 0{
		return nil
	}
	tree := SegTree{}
	var segTreeBuild func(index, l, r int)
	segTreeBuild = func(index, l, r int){
		if l == r {
			tree.element[index] = SegTreeNode{start: l, end: r, data: base[l]}
		}else{
			mid := (l + r) >> 1
			segTreeBuild((index << 1) + 1, l, mid) // 构造左子树
			segTreeBuild((index << 1) + 2, mid + 1, r) // 构造右子树
			// 计算 线段树 节点值，这里采用求最小值
			data := tree.element[(index << 1) + 1].data // 默认选左边 最小值
			if tree.element[(index << 1) + 1].data.(int) > tree.element[((index << 1) + 2)].data.(int){
				data = tree.element[((index << 1) + 2)].data
			}
			tree.element[index] = SegTreeNode{start: l, end: r, data: data}
		}
	}
	tree.element = make([]SegTreeNode, len(base)<<1 + 2) // 大于 2*n + 1
	segTreeBuild(0,0, length-1)
	return &tree
}
/*
  区间查询的核心思想就是找到交集运算可以构成待查询区间的所有的子区间，并且使找到的子区间的大小尽量大。
  简单的说就是，找到一些区间，使其连接起来之后正好可以涵盖整个待查询的区间。我们只需要找到代表这些区间的节点的最小值即可。
  通过二分的思想，把查询的复杂度降到O(logn)，我们在寻找这些子区间的时候，
  对于当前搜索到的子区间来说，有四种情况：当前区间和被查询区间无交集、待查询区间包含当前区间、当前区间包含待查询区间和当前区间和待查询区间有交集但互不包含，后两种情况可看成是同一种情况。
  1. 当前区间肯定不是待查询区间的子集，所以这时候应该返回一个极大值。表示取不到这个区间
  2. 待查询区间包含当前区间，这时候当前区间肯定是待查询区间的子集，所以应当返回当前区间的权值
  3&4. 当前区间包含待查询区间和当前区间和待查询区间有交集但互不包含，这时候当前区间的一部分是带查询区间的子集，另一部分不是，所以应当递归的去查询当前节点的左右子树，返回左右子树中较小的那个
 */
// curr: 当前区间的下标； left：待查询区间的左端点； right： 待查询区间的右端点
// @return: 返回当前区间在待查询区间中的部分的最小值
func (pt *SegTree) Query(curr, left, rigth int) interface{}{
	node := pt.element[curr]
	if node.start > rigth || node.end < left { // 情况1： 当前区间与待查区间无关
		return math.MaxInt32  // 因为这里维护的是最小值，因此此情况返回极大值
	}
	if node.start >= left && node.end <= rigth {
		return node.data // 第2种情况 当前的区间被包含在待查询的区间内则返回当前区间的最小值
	}
	pt.pushDown(curr)// 在返回左右子树的最小值之前，进行扩展操作！延迟更新，就是在查询的时候,进行更新mark
	// 余下的2种情况，递归查询左右子树，并比较
	l := pt.Query((curr<<1)+1, left, rigth)
	r := pt.Query((curr<<1)+2, left, rigth)
	data := l
	if l.(int) > r.(int){
		data = r
	}
	return data
}
// 更新一个点，curr：当前节点下标； update： 需要被更新的节点下标； data： 最新值
func (pt *SegTree) UpdateOne(curr, update int, data interface{}){
	node := pt.element[curr]
	if node.start == node.end && node.start == update{// 叶子节点 并且是目标节点
		node.data = data
		return
	}
	mid := (node.start + node.end) >> 1
	if update <= mid {
		pt.UpdateOne((curr<<1)+1, update, data)
	}else{
		pt.UpdateOne((curr << 1) + 2, update, data)
	}
	// 需要更新祖先或父节点的 最小值
	node.data = pt.element[curr<<1+1].data
	if pt.element[curr<<1+1].data.(int) > pt.element[curr<<1+2].data.(int){
		node.data = pt.element[curr<<1+2].data
	}
}

// 线段树的区间更新
/*
  延迟更新，延迟更新就是，更新的时候，我不进行操作，只是标记一下这个点需要更新，在我真正使用的时候我才去更新，
  这在我们进行一些数据库的业务的时候，也是很重要的一个思想。我们在封装节点的时候，有一个成员变量我们前面一直没有使用，那就是mark，
  现在就是使用这个成员变量的时候了。我们在进行区间修改的时候，我们把这个组成这个待修改区间的所有子区间都标记上。
  查找组成当前待修改区间的所有子区间的方法和查询方法是一样的，也是分三种情况
 */
// curr: 当前节点下标  left rigin ：待更新节点的左右端点   data: 增量值
func (pt *SegTree) Update(curr, left, right int, data interface{}){
	node := pt.element[curr]
	if node.start > right || node.end < left {
		return //如果当前的节点代表的区间和待更新的区间毫无交集，则返回不处理。
	}
	if node.start >= left && node.end <= right{ //如果当前的区间被包含在待查询的区间之内，则当前区间需要被标记上
		node.data = node.data.(int) + data.(int)
		node.AddMark(data.(int))
		return
	}
	pt.pushDown(curr) // 在更新左右子树之前进行扩展操作！
	pt.Update((curr<<1)+1, left, right, data)
	pt.Update(curr<<1 + 2, left, right, data)
	node.data = pt.element[curr<<1+1].data
	if pt.element[curr<<1+1].data.(int) > pt.element[curr<<1+2].data.(int){
		node.data = pt.element[curr<<1+2].data
	}
}

func (pt *SegTree) pushDown(curr int){//把当前节点的标志值传给子节点
	node := pt.element[curr]
	if node.mark != 0{
		pt.element[curr<<1+1].AddMark(node.mark)
		pt.element[curr<<1+2].AddMark(node.mark)
		pt.element[curr<<1+1].data = pt.element[curr<<1+1].data.(int) + node.mark
		pt.element[curr<<1+2].data = pt.element[curr<<1+2].data.(int) + node.mark
		node.ClearMark()
	}
}

// 327. Count of Range Sum
/*
  1. 前缀和数组 preSum，为了找出prefixSum[1] -- prefixSum[j]之间有多少满足以下条件的 x :
      整数x满足 lower <= prefixSum[j] - x <= upper x属于prefixSum[0...i-1]
      ==>  prefixSum[j] - lower  <=  x <= prefixSum[j] - upper
  1.1 为了找到x的数量，利用不等式等价转换，引申出2种优化方式
  2. 对于每个下标 j，以 j 为右端点的下标对的数量，就等于数组preSum[0..j-1]中的所有整数，
     出现在区间[preSum[j]−upper,preSum[j]−lower] 的次数。故很容易想到基于线段树的解法
  3. 从左至右扫描prefixSum 每遇到一个数prefixSum[j]，我们就在线段树中查询区间[preSum[j] - upper, preSum[j] - lower]内的整数数量。
     随后，将prefixSum[j]插入到线段树当中。
  注意：整数范围可能很大，故需要利用hash表将所有可能出现的整数，映射到连续的整数区间内
 */
type segTree []struct{
	l, r, val int
}
func(t segTree) build(o, l, r int){
	t[o].l, t[o].r  = l, r
	if l == r{
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1 | 1, m+1, r)
}
func (t segTree) inc(o, i int){
	if t[o].l == t[o].r {
		t[o].val++
		return
	}
	if i <= (t[o].l + t[o].r) >> 1{
		t.inc(o << 1, i) // 走左子树
	}else{
		t.inc(o << 1|1, i)
	}
	t[o].val = t[o << 1].val + t[o << 1 | 1].val
}
func(t segTree)query(o, l, r int)(res int){
	if l <= t[o].l && t[o].r <= r{
		return t[o].val
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o << 1, l, r)
	}
	if l > m {
		return t.query(o << 1 | 1, l, r)
	}
	return t.query(o<<1, l, r) + t.query(o<<1 | 1, l, r)
}
func countRangeSum(nums []int, lower int, upper int) (cnt int) {
	n := len(nums)
	allNums := make([]int, 1, 3*n+1)
	preSum := make([]int, n+1)
	for i, v := range nums{
		preSum[i+1] = preSum[i] + v
		allNums = append(allNums, preSum[i+1], preSum[i+1]-lower, preSum[i+1]-upper)
	}

	// 将 allNums 离散化
	sort.Ints(allNums)
	k := 1
	kth := map[int]int{allNums[0]: k}
	for i := 1; i <= 3*n; i++{
		if allNums[i] != allNums[i-1]{
			k++
			kth[allNums[i]] = k
		}
	}
	t := make(segTree, 4*k)
	t.build(1, 1, k)
	t.inc(1, kth[0])
	for _, sum := range preSum[1:] {
		left, right := kth[sum-upper], kth[sum - lower]
		cnt += t.query(1, left, right)
		t.inc(1, kth[sum])
	}
	return
}









