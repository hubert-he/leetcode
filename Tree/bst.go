package Tree

import (
	"math"
)

type BinarySearchTree struct {
	root *BiTreeNode
}

func (tree * BinarySearchTree)GetRoot() *BiTreeNode{
	return tree.root
}

func NewBST() *BinarySearchTree{
	return &BinarySearchTree{root: nil}
}

func NewBSTFromSortedList(array []interface{}) *BinarySearchTree{
	return &BinarySearchTree{root: sortedArrayToBST(array)}
}

func sortedArrayToBST(nums []interface{}) *BiTreeNode {
	right := len(nums) - 1
	left := 0
	if right >= left{
		/*
		mid := (left + right) / 2 		<== 总是选择中间位置左边的数字作为根节点
		mid := (left + right + 1) / 2	<== 总是选择中间位置右边的数字作为根节点
		mid := (left + right + rand.Intn(2)) / 2 <== 选择任意一个中间位置数字作为根节点
		 */
		mid := (right - left + 1) / 2
		root := &BiTreeNode{Val: nums[mid]}
		root.Left = sortedArrayToBST(nums[:mid])
		root.Right = sortedArrayToBST(nums[mid+1:])
		return root
	}else{
		return nil
	}
}

func NewBSTFromPlainList(array []interface{}) *BinarySearchTree{
	return &BinarySearchTree{root: GenerateBiTree(array)}
}

//501: 官方优秀解法
func (tree *BinarySearchTree)FindMode() (modes []int) {
	curElem, curCnt, maxCnt := 0, 0, 0
	update := func(x int){
		if x == curElem{
			 curCnt++
		}else{
			curElem, curCnt = x, 1
		}
		if maxCnt == curCnt{ // 不受 x 是否为新元素干扰
			modes = append([]int{curElem}, modes...)
		}else if maxCnt < curCnt{
			maxCnt = curCnt
			modes = []int{curElem}
		}
	}
	var dfs func(*BiTreeNode)
	dfs = func(root *BiTreeNode){
		if root == nil{
			return
		}
		dfs(root.Left)
		update(root.Val.(int))
		dfs(root.Right)
	}
	dfs(tree.root)
	return
}

func (tree *BinarySearchTree)FindModeMorris() (modes []int){
	var maxCnt, curElem, curCnt int
	update := func(x int){
		if x == curElem{
			curCnt++
		}else{
			curCnt, curElem = 1, x
		}
		if curCnt == maxCnt{
			modes = append([]int{curElem}, modes...)
		}else if curCnt > maxCnt {
			maxCnt = curCnt
			modes = []int{curElem}
		}
	}
	cur := tree.root
	for cur != nil {
		if cur.Left == nil{
			update(cur.Val.(int))
			cur = cur.Right
			continue
		}
		pre := cur.Left
		//如果当前节点的左孩子不为空，在当前节点的左子树中找到当前节点在中序遍历下的前驱节点
		for pre.Right != nil && pre.Right != cur{ // 通过pre.Right 是否等于 cur 来判断是否处理过了（加线索使得前驱节点指向后继节点）
			pre = pre.Right
		}
		if pre.Right == nil{ // 如果前驱节点的右孩子为空
			pre.Right = cur
			cur = cur.Left
		}else{ // 前驱节点的右孩子为当前节点 pre.Right == cur
			pre.Right = nil // 将它的右孩子重新设为空（恢复树的形状）
			update(cur.Val.(int)) // 输出当前节点
			cur = cur.Right // 当前节点更新为当前节点的右孩子
		}
	}
	return
}

//501: 求众数： 找出 BST 中的所有众数（出现频率最高的元素）
func (tree *BinarySearchTree)FindMode2() (modes []int){
	if tree.root == nil {
		return nil
	}
	maxCnt, curElem, curCnt := 0, 0, 0
	var dfs func(node *BiTreeNode)
	dfs = func(root *BiTreeNode){
		if root == nil {
			return
		}
		v := root.Val.(int)
		dfs(root.Left)
		if curCnt == 0{
			curCnt = 1
			curElem = v
		}else{
			if curElem != v{
				if curCnt == maxCnt{
					modes = append([]int{curElem}, modes...)
				}else{
					if curCnt > maxCnt{
						modes = []int{curElem}
						maxCnt = curCnt
					}
				}
				curCnt = 1
				curElem = v
			}else{
				curCnt++
			}
		}
		dfs(root.Right)
	}
	dfs(tree.root)
	if curCnt == maxCnt{
		modes = append([]int{curElem}, modes...)
	}else{
		if curCnt > maxCnt{
			modes = []int{curElem}
		}
	}
	return
}

//530: 二叉搜索树的最小绝对差
func (tree BinarySearchTree)GetMinimumDifference2() (difference int){
	pre := -1
	getMinimumDifference(tree.root, &pre, &difference)
	return
}
//530的优秀解法， 利用闭包，避免传参数
func (tree BinarySearchTree)GetMinimumDifference() (difference int){
	difference, pre := math.MaxInt64, -1
	var dfs func(*BiTreeNode)
	dfs = func(root *BiTreeNode){
		if root == nil {
			return
		}
		v := root.Val.(int)
		dfs(root.Left)
		if pre != -1{
			d := v - pre
			if d < difference{
				difference = d
			}
		}
		pre = v
		dfs(root.Right)
	}
	dfs(tree.root)
	if difference == math.MaxInt64{
		difference = 0
	}
	return
}

//530: 二叉搜索树的最小绝对差
func getMinimumDifference(root *BiTreeNode, pre *int, diff *int)  {
	if root == nil{
		return
	}
	v := root.Val.(int)
	getMinimumDifference(root.Left, pre, diff)
	if *pre != -1{
		d := v - *pre
		if d < *diff{
			*diff = d
		}
	}
	*pre = v
	getMinimumDifference(root.Right, pre, diff)
	return
}
// 17.12. BiNode
/*
二叉树数据结构TreeNode可用来表示单向链表（其中left置空，right为下一个链表节点）。
实现一个方法，把二叉搜索树转换为单向链表，要求依然符合二叉搜索树的性质，转换操作应是原址的
题眼：就是找到prev BST里的前驱节点
 */
func (tree *BinarySearchTree)ConvertBiNode() {
	var prev, head *BiTreeNode
	var dfs func(*BiTreeNode)
	dfs = func(node *BiTreeNode){
		if node == nil {
			return
		}
		dfs(node.Left)
		if prev == nil {
			head = node
		}else{
			prev.Right = node
		}
		prev = node
		node.Left = nil
		dfs(node.Right)
	}
	dfs(tree.root)
	tree.root = head
}

//235: 二叉搜索树种的 LCA 问题
/*
方法一：
  当我们分别得到了从根节点到p和q的路径之后，我们就可以很方便地找到它们的最近公共祖先了。
  显然，p和q的最近公共祖先就是从根节点到它们路径上的「分岔点」，也就是最后一个相同的节点。
  因此，如果我们设从根节点到p的路径为数组path_p，从根节点到q的路径为数组path_q，那么只要找出最大的编号i，其满足
   path_p[i]=path_q[i]
  那么对应的节点就是「分岔点」，即p和q的最近公共祖先就是path_p[i]（或path_q[i]）

方法二
我们从根节点开始遍历；
如果当前节点的值大于p和q的值，说明p和q应该在当前节点的左子树，因此将当前节点移动到它的左子节点；
如果当前节点的值小于p和q的值，说明p和q应该在当前节点的右子树，因此将当前节点移动到它的右子节点；
如果当前节点的值不满足上述两条要求，那么说明当前节点就是「分岔点」。此时，p和q要么在当前节点的不同的子树中，要么其中一个就是当前节点。
 */
func (tree *BinarySearchTree) LowestCommonAncestor(p, q *BiTreeNode) *BiTreeNode {
	return lowestCommonAncestor(tree.root, p, q)
}
// 用迭代实现 空间复杂O(1)
func lowestCommonAncestor(root, p, q *BiTreeNode) (ancestor *BiTreeNode){
	ancestor = root
	rv, pv, qv := ancestor.Val.(int), p.Val.(int), q.Val.(int)
	for {
		rv = ancestor.Val.(int)
		switch {
		case rv > pv && rv > qv:
			ancestor = ancestor.Left
		case rv < pv && rv < qv:
			ancestor = ancestor.Right
		default:
			return
		}
	}
}
// 递归实现，代价高
func lowestCommonAncestor2(root, p, q *BiTreeNode) *BiTreeNode {
	switch {
	case root.Val.(int) > p.Val.(int) && root.Val.(int) > q.Val.(int):
		return lowestCommonAncestor2(root.Left, p, q)
	case root.Val.(int) < p.Val.(int) && root.Val.(int) < q.Val.(int):
		return lowestCommonAncestor2(root.Right, p, q)
	default:
		return root
	}
}
//剑指 Offer 54  二叉搜索树的第k大节点  关联问题为 逆序打印BST
func (tree *BinarySearchTree) KthLargest(k int) interface{} {
	var dfs func(*BiTreeNode)
	cnt := 0
	var target interface{}
	dfs = func(node *BiTreeNode){
		if node == nil{
			return
		}
		dfs(node.Right)
		cnt++
		if cnt == k{
			target = node.Val
			return
		}
		dfs(node.Left)
	}
	dfs(tree.root)
	return target
}

//270. 最接近的二叉搜索树值
func (tree *BinarySearchTree) ClosestValue(target float64) (it *BiTreeNode){
	var dfs func(node *BiTreeNode)
	var diff float64 = math.MaxFloat64
	dfs = func(node *BiTreeNode){
		if node == nil {
			return
		}
		v := float64(node.Val.(int))
		if v < target{
			if (target - v) < diff{
				diff = target - v
				it = node
			}
			dfs(node.Right)
		}else{
			if (v - target) < diff{
				diff = v - target
				it = node
			}
			dfs(node.Left)
		}
	}
	dfs(tree.root)
	return
}









