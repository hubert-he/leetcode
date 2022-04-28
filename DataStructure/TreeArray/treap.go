package TreeArray

import (
	"math"
	"math/rand"
)

/* 本次为 Treap 树堆 属于 AVL 平衡树系列
** treap让平衡树的每一个节点存放2个信息：
** 	1. 值（满足BST特性）
**  2. 一个随机索引（满足heap的性质）
** 即结合二叉搜索树和二叉堆的性质来使树平衡
** 解决的问题是BST 如果插入值是顺序插入的，BST 会退回为 链表 达不到 O(logn) 复杂度查询
** Treap使用二叉堆来维护随机索引，起始就相当于把插入次序随机化，插入一数值后必然满足索引满足堆的特性，但因为是随机索引，必然导致旋转
** 继而相当于次序打乱。
 */
/* Treap 的题目：
**
** 1649. Create Sorted Array through Instructions
** 998. Maximum Binary Tree II
** 327. Count of Range Sum
** 面试题 10.10. Rank from Stream LCCI
 */
type BalancedNode struct{
	key 	int // BST
	seed 	int // heap 值
	count 	int // 记录相同的数个数
	size 	int // Treap中节点子树的记录数量
	left, right *BalancedNode
}

func(this *BalancedNode)LeftRotate()*BalancedNode{
	prevSize := this.size
	currSize := this.count
	if this.left != nil{
		currSize += this.left.size
	}
	if this.right != nil{
		currSize += this.right.size
	}
	node := this.right
	this.right = node.left
	node.left = this

	node.size = prevSize
	this.size = currSize
	return node // 返回调整后的
}
func(this *BalancedNode)RightRotate()*BalancedNode{
	prevSize := this.size
	currSize := this.count
	if this.right != nil{
		currSize += this.right.size
	}
	if this.left.right != nil{
		currSize += this.left.right.size
	}
	node := this.left
	this.left = node.right
	node.right = this
	node.size = prevSize
	this.size = currSize
	return node
}

type BalancedTree struct{
	root 	*BalancedNode
	size 	int
}

func(this *BalancedTree)_put(node *BalancedNode, val int) *BalancedNode{
	if node == nil {
		// return &BalancedNode{key: val, count: 1, size: 1} 更新增加树堆支持，由BST 转换为 Treap
		return &BalancedNode{key: val, seed: rand.Int(), count: 1, size: 1}
	}
	node.size++
	if val < node.key {
		node.left = this._put(node.left, val)
		if node.left.seed > node.seed { // 不满足大根堆性质
			node = node.RightRotate()
		}
	}else if val > node.key{
		node.right = this._put(node.right, val)
		if node.right.seed > node.seed{ // node 与 node的右节点比较 判定是否 左旋
			node = node.LeftRotate()
		}
	}else{
		node.count++ // 插入重复数
	}
	return node
}

func(this *BalancedTree) Put(val int){
	this.size++
	this.root = this._put(this.root, val)
}

func(this *BalancedTree) _del(node *BalancedNode, val int)(*BalancedNode, bool){
	var ok bool
	if node.key == val{
		if node.count > 1{
			node.count--
		}else{
			// 情况-1： 若当前节点没有左儿子与右儿子，则直接删除该节点
			// 情况-2： 若当前节点只有左儿子/右儿子，则用左儿子/右儿子替代该节点
			if node.left == nil{
				return node.right, true
			}else if node.right == nil{
				return node.left, true
			// 	情况-3： 若当前节点有左儿子与右儿子，则不断旋转当前节点，并走到当前节点新的对应位置，直到当前节点只有左儿子/右儿子为止。
			}else{
				if node.left.seed < node.right.seed{ // 满足大根堆
					node = node.LeftRotate()
					node.left, ok = this._del(node.left, val)
				}else{
					node = node.RightRotate()
					node.right, ok = this._del(node.right, val)
				}
			}
		}
	}else if node.key > val{
		node.left, ok = this._del(node.left, val)
	}else{
		node.right, ok = this._del(node.right, val)
	}
	if ok{
		node.size--
	}
	return node, ok
}
func(this *BalancedTree) Del(val int){
	if this.size <= 0 { return }
	var ok bool
	if this.root, ok = this._del(this.root, val); ok{
		this.size--
	}
}

func(this *BalancedTree)LowerBound(val int)*BalancedNode{
	root := this.root
	var ans *BalancedNode
	for root != nil {
		if val == root.key{
			return root
		}
		if val < root.key{
			ans = root
			root = root.left
		}else{
			root = root.right
		}
	}
	return ans
}

func(this *BalancedTree)UpperBound(val int)*BalancedNode{
	root := this.root
	var ans *BalancedNode
	for root != nil {
		if val < root.key{
			ans = root
			root = root.left
		}else{
			root = root.right
		}
	}
	return ans
}

func(this *BalancedTree)Rank(val int)(int, int){
	root := this.root
	ans := 0
	for root != nil{
		if val < root.key{
			root = root.left
		}else{
			if root.left != nil{
				ans += root.left.size
			}
			ans += root.count
			if val == root.key{
				return ans - root.count + 1, ans
			}
			root = root.right
		}
	}
	return math.MinInt32, math.MaxInt32
}

////////////////////另外一种实现：FHQ Treap 不选择 而是分裂 //////////////
type FHQTreapNode struct{
	key 		int // BST
	rnd 		int // heap 值
	count 		int // 记录相同的数个数
	size 		int // Treap中节点子树的记录数量
	left, right int   // 左右子树的编号
}

func newFHQTreepNode(val int){

}
/////////////////// FHQ 数组写法
type FhqNode struct {
	left, right 	int
	key, rnd 		int // key: BST的值   rnd: Heap的值
	size, count		int
}

func NewFhqNode(val int)FhqNode{
	this := FhqNode{key: val, rnd: rand.Int(), count: 1, size: 1}
	return this
}

type FhqTreapArray struct{
	fhq 	[]FhqNode
	root 	int
	idx 	int
}

func ConstructFhqTreapArray()*FhqTreapArray{
	// 新建立内存池
	const Maxn = 1e5+5
	this := FhqTreapArray{}
	this.fhq = make([]FhqNode, Maxn)
	return &this
}

func(this FhqTreapArray)updateSize(now int){
	left, right := this.fhq[now].left, this.fhq[now].right
	this.fhq[now].size = this.fhq[left].size + this.fhq[right].size
}
// 分裂有2种：按值分裂 和 按 大小分裂
// 按值分裂： 把树拆成课树，拆出来的一颗树的值 全部小于等于给定的值，另外一部分的值全部大于给定的值
//		一般来说，当正常的平衡树用的时候，使用按值分裂
// 大小分裂： 把树拆成两颗树，拆出来的一颗树的大小 等于 给定的大小，剩余部分在另外一颗树里
// 		在维护区间信息的时候，使用按大小分裂，例如文艺平衡树
// split： 返回的 x 表示 全部小于等于 val 的节点根；  y 表示 大于 val 的节点根  都满足 BST 性质
func(this *FhqTreapArray)Split(now, val int)(x, y int){
	if now == 0{ // 空节点
		return 0, 0
	}
	fhq := this.fhq
	if this.fhq[now].key <= val{ // BST：向右查找
		x = now
		fhq[now].right, y= this.Split(fhq[now].right, val)
	}else{
		y = now
		fhq[now].left, x = this.Split(fhq[now].left, val)
	}
	this.updateSize(now)
	return
}
// merge：把两颗树 x, y 合并为一颗树，其中 x树上的所有值都 小于等于 y 树上的所有值。
// 并且新合并出来的树依然满足Treap 性质
func(this *FhqTreapArray)Merge(x, y int)(root int){// 前提是 x <= y
	if x == 0 || y == 0{
		return x+y
	}
	fhq := this.fhq
	// 首先确定满足Heap的性质，因此此处 > < >= <= 均可以，不在乎是否是大根堆 还是小根堆
	// 这里只要求 合并的随机性
	if fhq[x].rnd < fhq[y].rnd{// 这是小根堆的方式, 小根堆：返回的root是 x  大根堆的话 root返回的是 y
		fhq[x].right = this.Merge(fhq[x].right, y)
		this.updateSize(x)
		return x
	}else{
		fhq[y].left = this.Merge(x, fhq[y].left)
		this.updateSize(y)
		return y
	}
}

// 插入：设插入的值为val
// 1. 按val值分裂，得到 x 和 y 两棵树
// 2. 按值分裂的定义，我们分裂出来的 x 树上的所有值一定都小于等于 val
// 	  y树上的所有值一定都大于val
// 3. 直接合并 x，用值val 新建的新节点 y
func(this *FhqTreapArray)Insert(val int){
	node := NewFhqNode(val)
	this.idx++
	this.fhq[this.idx] = node
	left, right := this.Split(this.root, val)
	this.root = this.Merge(this.Merge(left, this.idx), right)
}
// 设删除 val 值
// 按值 val 把树分裂成 x 和 z
// 再按值 val-1 把子树 x 分裂成 x 和 y
// 那么此时 y 树上的所有值 都是等于 val 的，我们去掉它的根节点：让 y 等于合并y的左子树和右子树
// 注意 此 Treap 中 相同的val 的节点 可以有多个，并没有使用count 合并， 所有 y 是一颗子树 而非一个节点
// 最后合并 x y z
func(this *FhqTreapArray)Delete(val int){
	fhq := this.fhq
	x, z := this.Split(this.root, val)
	x, y := this.Split(x, val-1)
	// y 就为一个子树 或 节点
	y = this.Merge(fhq[y].left, fhq[y].right)
	this.root = this.Merge(this.Merge(x, y), z)
}
// 查找 val 的排名
// 1. 按val-1 分裂为 x 和 y
// 2. 则 x.size + 1 就是 val的排名
// 3. 最后把x 和 y 合并起来
func(this *FhqTreapArray)GetRank(val int) int{
	fhq := this.fhq
	x, y := this.Split(this.root, val-1)
	ans := fhq[x].size + 1
	this.root = this.Merge(x, y)
	return ans
}
// 根据Rank 获取对应的数字
// 直接循环处理, 返回的是节点索引
func(this *FhqTreapArray)GetNumByRank(rank int) int{
	if rank > this.idx{
		return 0 // nil 节点
	}
	fhq := this.fhq
	now := this.root
	for now != 0{
		left, right := fhq[now].left, fhq[now].right
		if fhq[left].size + 1 == rank{
			return now
		}else if fhq[left].size >= rank{
			now = left
		}else{// fhq[left].size < rank
			// 注意此处需要减掉left 个数
			rank -= fhq[left].size + 1
			now = right
		}
	}
	return 0
}
// val 的 前驱：按值 val-1 分裂成 x 和 y, 则 x 里面最右的数就是 val 的前驱
func (this *FhqTreapArray)Predecessor(val int)int{

}
// val 的后继：按值 val 分裂成 x 和 y，则 y 里面最左的数就是 val 的后继
func (this *FhqTreapArray)Predecessor(val int)int{

}

func FhqTreap() {
	// 新建立内存池
	const maxn = 1e5+5
	var fhq = make([]FhqNode, maxn)
	cnt, root := 0, 0
	newFhqNode := func(val int) int{
		cnt++ // 避开 索引 0
		fhq[cnt].key = val
		fhq[cnt].rnd = rand.Int()
		fhq[cnt].size, fhq[cnt].count = 1, 1
		return cnt
	}
	updateSize := func(now int){
		fhq[now].size = fhq[fhq[now].left].size + fhq[fhq[now].right].size + fhq[now].count
	}
	// 分裂有2种：按值分裂 和 按 大小分裂
	// 按值分裂： 把树拆成课树，拆出来的一颗树的值 全部小于等于给定的值，另外一部分的值全部大于给定的值
	//		一般来说，当正常的平衡树用的时候，使用按值分裂
	// 大小分裂： 把树拆成两颗树，拆出来的一颗树的大小 等于 给定的大小，剩余部分在另外一颗树里
	// 		在维护区间信息的时候，使用按大小分裂，例如文艺平衡树
	// split： 返回的 x 表示 全部小于等于 val 的节点根；  y 表示 大于 val 的节点根  都满足 BST 性质
	var split func(now, val int)(x, y int)
	split := func(now, val int)(x, y int){
		if now == 0{ return 0, 0}
		if fhq[now].key <= val{
			x = now
			// 查看右边递归分裂
			left, right := split(fhq[now].right, val)
			fhq[x].right = left
			y = right
			// fhq[x].right, y = split(fhq[now].right, val)
		}else{
			y = now
			left, right := split(fhq[now].left, val)
			fhq[y].left = right
			x = left
			// x, fhq[y].left = split(fhq[now].left, val)
		}
		updateSize(now)
		return
	}
	// merge：把两颗树 x, y 合并为一颗树，其中 x树上的所有值都 小于等于 y 树上的所有值。并且新合并出来的树依然满足Treap 性质
	var merge func(x, y int)int
	merge := func(x, y int) int{// 记住前提===>要求 x <= y
		if x == 0 || y == 0 { return x + y }
		// 首先确定满足Heap的性质，因此此处 > < >= <= 均可以，不在乎是否是大根堆 还是小根堆
		if fhq[x].key < fhq[y].key{ // 这是小根堆的方式, 小根堆：返回的root是 x  大根堆的话 root返回的是 y
			fhq[x].right = merge(fhq[x].right, y)
			updateSize(x)
			return x
		}else{// 这是小根堆的方式, 小根堆：返回的root是 y
			fhq[y].left = merge(x, fhq[y].left)
			updateSize(y)
			return y
		}
	}
	// 插入：设插入的值为val
	// 1. 按val值分裂，得到 x 和 y 两棵树
	// 2. 按值分裂的定义，我们分裂出来的 x 树上的所有值一定都小于等于 val
	// 	  y树上的所有值一定都大于val
	// 3. 直接合并 x，用值val 新建的新节点 y
	insert := func(val int){
		left, right := split(root, val)
		root = merge(merge(left, newFhqNode(val)), right)
	}
	// 设删除 val 值
	// 按值 val 把树分裂成 x 和 z
	// 再按值 val-1 把子树 x 分裂成 x 和 y
	// 那么此时 y 树上的所有值 都是等于 val 的，我们去掉它的根节点：让 y 等于合并y的左子树和右子树
	// 注意 此 Treap 中 相同的val 的节点 可以有多个，并没有使用count 合并， 所有 y 是一颗子树 而非一个节点
	// 最后合并 x y z
	delete := func(val int){
		x, z := split(root, val)
		x, y := split(x, val-1) // y 为 val 的子树
		y = merge(fhq[y].left, fhq[y].right)
		root = merge(merge(x, y), z)
	}
	// 查找 val 的排名
	// 1. 按val-1 分裂为 x 和 y
	// 2. 则 x.size + 1 就是 val的排名
	// 3. 最后把x 和 y 合并起来
	getRank := func(val int)int{
		left, right := split(root, val-1)
		ans := fhq[left].size + 1
		root = merge(left, right)
		return ans
	}
	// 直接循环处理
	getRankNum := func(rank int)int{
		now := root
		for now != 0{
			if fhq[fhq[now].left].size + 1 == rank{// 找到了
				break
			}else if fhq[fhq[now].left].size >= rank{
				now = fhq[now].left
			}else{
				rank -= fhq[fhq[now].left].size + 1
				now = fhq[now].right
			}
		}
		return fhq[now].val
	}
	// val 的 前驱：按值 val-1 分裂成 x 和 y, 则 x 里面最右的数就是 val 的前驱
	predecessor := func(val int)int{
		left, right := split(root, val-1)
		now := left
		ans := now // 此题是数组实现， 可以不需要ans
		for now != 0{
			ans = now
			now = fhq[now].right
		}
		root = merge(left, right)
		return ans
	}
	// val 的后继：按值 val 分裂成 x 和 y，则 y 里面最左的数就是 val 的后继
	successor := func(val int)int{
		left, right := split(root, val)
		now := right
		for fhq[now].left != 0{
			now = fhq[now].left
		}
		root = merge(left, right)
		return now
	}
}


/* 1649. Create Sorted Array through Instructions
** Given an integer array instructions, you are asked to create a sorted array from the elements in instructions.
** You start with an empty container nums. For each element from left to right in instructions, insert it into nums.
** The cost of each insertion is the minimum of the following:
	1. The number of elements currently in nums that are strictly less than instructions[i].
	2. The number of elements currently in nums that are strictly greater than instructions[i].
** For example, if inserting element 3 into nums = [1,2,3,5],
** the cost of insertion is min(2, 1) (elements 1 and 2 are less than 3, element 5 is greater than 3) and
** nums will become [1,2,3,3,5].
** Return the total cost to insert all elements from instructions into nums.
** Since the answer may be large, return it modulo 10^9 + 7
 */
// 平方级不允许
func createSortedArray_tle(instructions []int) int {
	ans := 0
	n := len(instructions)
	for i := 1; i < n; i++{
		min, max := 0, 0
		for j := 0; j < i; j++{
			if instructions[j] > instructions[i]{
				max++
			}else if instructions[j] < instructions[i]{
				min++
			}
		}
		if max > min{
			ans += min
		}else{
			ans += max
		}
	}
	return ans
}

func createSortedArray_treap(instructions []int) int {
	const Mod = 1e9+7
	treap := BalancedTree{}
	ans := 0
	for i, cmd := range instructions{
		lb := treap.LowerBound(cmd)
		smaller := i
		if lb != nil {
			fst, _ := treap.Rank(lb.key)
			smaller = fst - 1
		}
		rb := treap.UpperBound(cmd)
		larger := 0
		if rb != nil{
			fst, _ := treap.Rank(rb.key)
			//larger = fst + 1
			larger = i - (fst - 1) // 根据题意 i 为当前总数
		}
		if larger > smaller{
			ans += smaller
		} else{
			ans += larger
		}
		treap.Put(cmd)
	}
	return ans % Mod
}

/* 998. Maximum Binary Tree II
** 考虑Treap的建树方式，给定一个数组A，
** 该Treap的堆性质由 A[i] 来维护，其BST性质由下标维护。
** 那么Treap的建树方式的递归的，树根是全数组最大值 A[k]，然后其左子树和右子树分别是 A[0:k−1] 和A[k+1:]，接着就可以递归建树了。
** 给定一棵已经建好的Treap的树根，其是以某个数组 A 建立的（该数组事实上是唯一的），
** 要求返回A ∪ { x } 建立的Treap的树根，直接在给出的Treap里把 x 加进去，不允许新建树。
 */
// 题目把 Treap 的插入一个值 给抽象出来了
// 并且这个Treap 中 Heap 为大根堆
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	node := TreeNode{Val: val}
	if root.Val < val{
		node.Left = root // 维护BST 性质：新加的 下标肯定 大于 全部已有的
		return &node
	}
	// 否则 为了维护BST性质，只能往右边插入
	root.Right = insertIntoMaxTree(root.Right, val)
	return root
}

















