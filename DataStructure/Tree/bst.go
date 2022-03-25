package Tree

import (
	"fmt"
	"math"
)

type BinarySearchTree struct {
	root *BiTreeNode
}
// 173. Binary Search Tree Iterator
type BSTIterator struct {
	stack	[]*BiTreeNode
	cur		*BiTreeNode
}

func Constructor(root *BiTreeNode) BSTIterator {
	return BSTIterator{cur: root}
}

func (this *BSTIterator) Next() interface{}{
	for node := this.cur; node != nil; node = node.Left{
		this.stack = append([]*BiTreeNode{node}, this.stack...)
	}
	this.cur, this.stack = this.stack[0], this.stack[1:]
	val := this.cur.Val
	this.cur = this.cur.Right
	return val
}

func (this *BSTIterator) HasNext() bool {
	return this.cur != nil || len(this.stack) > 0
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

// BST remove node
func (tree *BinarySearchTree)DeleteNode(key interface{}) (node *BiTreeNode){
	fakeParent := &BiTreeNode{Left: tree.root}
	// 注意parent的写法，二级指针，无需判断left right 即可指定左右方向
	parent := &(fakeParent.Left)
	cur := fakeParent.Left
	for cur != nil {
		if key == cur.Val{
			node = cur
			break
		}
		if key.(int) > cur.Val.(int){
			parent = &(cur.Right)
			cur = cur.Right
		}
		if key.(int) < cur.Val.(int){
			parent = &(cur.Left)
			cur = cur.Left
		}
	}
	if cur == nil {
		// key 不存在
		return nil
	}
	successor := cur.Right
	if successor == nil {
		*parent = cur.Left
	}else{
		parent = &(cur.Right)
		for successor.Left != nil {
			parent = &(successor.Left)
			successor = successor.Left
		}
		cur.Val = successor.Val
		*parent = successor.Right
	}
	tree.root = fakeParent.Left
	return
}
// 递归
func (tree *BinarySearchTree)DeleteNode2(key interface{}) (node *BiTreeNode){
	var remove func(*BiTreeNode, interface{})*BiTreeNode
	remove = func(treeNode *BiTreeNode, key interface{})*BiTreeNode{
		if treeNode == nil {
			return nil
		}
		if treeNode.Val.(int) < key.(int){
			treeNode.Right = remove(treeNode.Right, key)
			return treeNode
		}
		if treeNode.Val.(int) > key.(int){
			treeNode.Left = remove(treeNode.Left, key)
			return treeNode
		}
		// leaf remove
		if treeNode.Right == nil && treeNode.Left == nil{
			return nil
		}
		// 右子树为空，用前驱替换
		cur := treeNode.Right
		if cur == nil {
			cur = treeNode.Left
			for cur.Right != nil {
				cur = cur.Right
			}
			treeNode.Val = cur.Val
			treeNode.Left = remove(treeNode.Left, cur.Val)
		}else{
			for cur.Left != nil {
				cur = cur.Left
			}
			treeNode.Val = cur.Val
			treeNode.Right = remove(treeNode.Right, cur.Val)
		}
		return treeNode
	}
	tree.root = remove(tree.root, key)
	return
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

//938. 二叉搜索树的范围和
func rangeSumBST(root *BiTreeNode, low int, high int) int {
	sum := 0
	var dfs func(*BiTreeNode)
	dfs = func(node *BiTreeNode){
		if node == nil{
			return
		}
		v := node.Val.(int)
		if v >= low && v <= high{
			sum += v
		}
		if v >= low{ // 剪枝
			dfs(node.Left)
		}
		if v <= high{
			dfs(node.Right)
		}
	}
	dfs(root)
	return sum
}

// 783. 二叉搜索树节点最小距离
func minDiffInBST(root *BiTreeNode) int {
	var cur *BiTreeNode
	min := math.MaxInt32
	var dfs func(*BiTreeNode)
	dfs = func(node *BiTreeNode){
		if node == nil{
			return
		}
		dfs(node.Left)
		if cur != nil{
			diff := node.Val.(int) - cur.Val.(int)
			if diff < min{
				min = diff
			}
		}
		cur = node
		dfs(node.Right)
	}
	dfs(root)
	return min
}

//270. 最接近的二叉搜索树值
/*
给定一个不为空的二叉搜索树和一个目标值 target，请在该二叉搜索树中找到最接近目标值 target 的数值。
注意：
给定的目标值 target 是一个浮点数
题目保证在该二叉搜索树中只会存在一个最接近目标值的数
示例：
输入: root = [4,2,5,1,3]，目标值 target = 3.714286

    4
   / \
  2   5
 / \
1   3

输出: 4
 */
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

// 1305 All Elements in Two Binary Search Trees
/*
Given two binary search trees root1 and root2.
Return a list containing all the integers from both trees sorted in ascending order.
*/

func getAllElements(root1 *BiTreeNode, root2 *BiTreeNode) []interface{} {
	res := []interface{}{}
	ch1, ch2 := make(chan interface{}), make(chan interface{})
	go traval(root1, ch1)
	go traval(root2, ch2)
	ok1, ok2 := true, true
	//v1, v2 := math.MaxInt32, math.MaxInt32
	var v1,v2 interface{}
	for ok1 || ok2 {
		if v1 == nil {
			v1, ok1 = <-ch1
		}
		if v2 == nil {
			v2, ok2 = <-ch2
		}

		if ok1 && ok2 {
			if v1.(int) < v2.(int) {
				res = append(res, v1)
				v1 = nil
			} else {
				res = append(res, v2)
				v2 = nil
			}

		} else if ok1 {
			res = append(res, v1)
			v1 = nil
		}else if ok2 {
			res = append(res, v2)
			v2 = nil
		}
	}

	return res
}

func traval(root *BiTreeNode, ch chan interface{}) {
	var dfs func(*BiTreeNode)
	dfs = func(node *BiTreeNode){
		if node != nil {
			dfs(node.Left)
			ch <- node.Val
			dfs(node.Right)
		}
	}
	dfs(root)
	close(ch)
}

// 04.09. 二叉搜索树序列
/* backtracking
  1. 确定回溯函数返回值以及参数
  2. 确定回溯函数终止条件
  3. 确定回溯函数遍历过程
模板：
  void backtracking(参数) {
     if 终止条件 {
 		存放结果
		return
     }
	for 选择：本层集合中元素 {
		处理结点
		backtracking(路径，选择列表)
		回溯，撤销处理结果
	}
  }
	回溯算法除了要维护一个path用来保存路径之外，还需要额外维护一个候选节点队列dq
	择使用双端队列来维护候选队列，这样每次插入和回溯的时间复杂度都可以降到O(1).
 */
func BSTSequences(tree *BinarySearchTree) (result [][]interface{}) {
	root := tree.root
	if root == nil {
		// return nil 会返回 []int{}
		return [][]interface{}{}
	}
	dq := []*BiTreeNode{root}
	var dfs func(path []interface{}, dq []*BiTreeNode)
	cnt := 0
	dfs = func(path []interface{}, dq []*BiTreeNode){
		size := len(dq)
		if size == 0{
			result = append(result, append([]interface{}{}, path...))
			//result = append(result, path)  底层数据存在共享，因此需要copy
			return
		}
		cc := cnt
		cnt++
		for i := 0; i < size; i++{
			fmt.Printf("~%d~%d: ", cc,i)
			for _,value := range dq{
				fmt.Printf("%v ", value.Val)
			}
			fmt.Println("")
			cur := dq[0]
			dq = dq[1:]
			path = append(path, cur.Val)
			if cur.Left != nil {
				dq = append(dq, cur.Left)
			}
			if cur.Right != nil {
				dq = append(dq, cur.Right)
			}
			dfs(path, dq)
			fmt.Printf("-%d-%d: ", cc,i)
			for _,value := range dq{
				fmt.Printf("%v ", value.Val)
			}
			fmt.Println(path, size)
			dq = dq[0:size - 1]
			dq = append(dq, cur)
			path = path[:len(path) - 1]
		}
	}
	dfs([]interface{}{}, dq)
	return
}

/*
  Given the root of a Binary Search Tree (BST),
  convert it to a Greater Tree(累加树) such that every key of the original BST is changed to the original key plus sum of all keys greater than the original key in BST.
  As a reminder, a binary search tree is a tree that satisfies these constraints:
The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
 */
/*
  反中序遍历-- 可以采用线索二叉树提升性能
  Morris遍历的核心思想是利用树的大量空闲指针，实现空间开销的缩减。
  1. 如果当前节点的Right为空， 处理当前节点，并遍历当前节点的Left
  2. 如果当前节点的Right不为空，找到当前节点Right的最左节点（该节点为当前节点中序遍历的前驱节点）：
 	2.a 如果最左节点的Left为空， 将最左节点的Left指向当前节点，遍历当前节点的Right
	2.b 如果最左节点的Left不为空， 将最左节点的Left重置为空（恢复树的原状），处理当前节点，并将当前节点置为其Left
  3. 重复1和2
 */
func getSuccessor(root *BiTreeNode) (succ *BiTreeNode){
	succ = root.Right
	for succ.Left != nil && succ.Left != root{ // succ.Left 与 root 关系
		succ = succ.Left
	}
	return
}
func convertBST(root *BiTreeNode) *BiTreeNode {
	sum := 0
	node := root
	for node != nil {
		if node.Right == nil{
			sum += node.Val.(int)
			node.Val = sum
			node = node.Left
		}else{
			succ := getSuccessor(node)
			if succ.Left == nil {
				succ.Left = node
				node = node.Right
			}else{// 恢复
				succ.Left = nil
				sum += node.Val.(int)
				node.Val = sum
				node = node.Left
			}
		}
	}
	return root
}
/* 04.05. 合法二叉搜索树
方法1： 将bst 看作链表来对待，即中序处理
 */
func IsValidBST(root *BiTreeNode) bool {
	var prev *BiTreeNode
	var dfs func(*BiTreeNode)bool
	dfs = func(node *BiTreeNode)bool{
		if node == nil {
			return true
		}
		if dfs(node.Left){
			if prev != nil && prev.Val.(int) >= node.Val.(int){
				return false
			}
			prev = node
			return dfs(node.Right)
		}
		return false
	}
	return dfs(root)
}
/* 非递归实现
*/
func IsValidBST2(root *BiTreeNode)bool{
	st := []*BiTreeNode{}
	var prev *BiTreeNode
	for len(st) > 0 || root != nil{
		for root != nil {
			st = append([]*BiTreeNode{root}, st...)
			root = root.Left
		}
		if prev != nil && st[0].Val.(int) <= prev.Val.(int){
			return false
		}
		prev = st[0]
		root = st[0].Right
		st = st[1:]
	}
	return true
}
/*
  递归：
  如果该二叉树的左子树不为空，则左子树上所有节点的值均小于它的根节点的值；
  若它的右子树不空，则右子树上所有节点的值均大于它的根节点的值；它的左右子树也为二叉搜索树
 */
func IsValidBST3(root *BiTreeNode)bool{
	var dfs func(*BiTreeNode, int, int)bool
	dfs = func(node *BiTreeNode, lower int, upper int)bool{
		if node == nil {
			return true
		}
		value := node.Val.(int)
		if value <= lower || value >= upper{
			return false
		}
		return dfs(node.Left, lower, value) && dfs(node.Right, value, upper)
	}
	return dfs(root, math.MinInt32, math.MaxInt32)
}

/* 95. Unique Binary Search Trees II
	Given an integer n, return all the structurally unique BST's (binary search trees),
    which has exactly n nodes of unique values from 1 to n. Return the answer in any order.
Example 1:
	Input: n = 3
	Output: [[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
Example 2:
	Input: n = 1
	Output: [[1]]
*/
func GenerateTreesBST(n int) []*BiTreeNode {
	ans := []*BiTreeNode{}
	nums := make([]int, n)
	for i := 1; i <= n; i++{
		nums[i-1] = i
	}
	var dfs func([]int)[]*BiTreeNode
	dfs = func(nodes []int)[]*BiTreeNode{
		if len(nodes) <= 0{
			return []*BiTreeNode{nil} // nil 为必须，双重
		}
		res := []*BiTreeNode{}
		for i := range nodes{
			left := dfs(nodes[:i])
			right := dfs(nodes[i+1:])
			for _, l := range left{
				for _, r := range right{
					res = append(res, &BiTreeNode{Val: nodes[i], Left: l, Right: r})
				}
			}
		}
		return res
	}
	ans = append(ans, dfs(nums)...)
	return ans
}
/*	官方题解
*/
func generateTrees(n int) []*BiTreeNode {
	if n == 0{
		return nil
	}
	var dfs func(int, int)[]*BiTreeNode
	dfs = func(start, end int)[]*BiTreeNode{
		if start > end{
			return []*BiTreeNode{nil}
		}
		allTrees := []*BiTreeNode{}
		for i := start; i <= end; i++{
			left := dfs(start, i - 1)
			right := dfs(i+1, end)
			for _, l := range left{
				for _, r := range right{
					allTrees = append(allTrees, &BiTreeNode{Val: i, Left: l, Right: r})
				}
			}
		}
		return allTrees
	}
	return dfs(1, n)
}
/*96. Unique Binary Search Trees
Given an integer n,
return the number of structurally unique BST's which has exactly n nodes of unique values from 1 to n.
Example 1:
	Input: n = 3
	Output: 5
Example 2:
	Input: n = 1
	Output: 1
	方程为： dp[i] = SUM(dp[j-1]*dp[i-j]) j属于[1,i]
 */
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++{
		for j := 1; j <= i; j++{
			dp[i] += (dp[j-1] * dp[i-j])
		}
	}
	return dp[n]
}

// 2021-11-15 重刷此题
func NumTrees2(n int) int {
	// dp[i] = sum(dp[j]*dp[i-1-j]) j 从0到i-1
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++{
		for j := 0; j < i; j++{
			dp[i] += dp[j]*dp[i-1-j]
		}
	}
	return dp[n]
}

/* 285. Inorder Successor in BST
** Given the root of a binary search tree and a node p in it, return the in-order successor of that node in the BST.
** If the given node has no in-order successor in the tree, return null.
** The successor of a node p is the node with the smallest key greater than p.val.
 */
/* 二叉搜索树的中序遍历结果是一个递增的数组，顺序后继是中序遍历中当前节点 之后 最小的节点。
** 可以分成两种情况来讨论：
	1. 如果当前节点有右孩子，顺序后继在当前节点之下
		先找到当前节点的右孩子，然后再持续往左直到左孩子为空
	2. 如果当前节点无右孩子，顺序后继在当前节点之上
		由于无法访问父亲节点，只能从根节点开始中序遍历，直接在中序遍历过程保存前一个访问的节点，判断前一个节点是否为p
		如果前一个节点是 p 如果是，则当前节点就是 p 节点顺序后继节点。
** 算法
	如果当前节点有右孩子，找到右孩子，再持续往左走直到节点左孩子为空，直接返回该节点。
	如果没有的话，就需要用到非递归的中序遍历。维持一个栈，当栈中有节点时：
		往左走直到节点的左孩子为空，并将每个访问过的节点压入栈中。
		弹出栈中节点，判断当前的前继节点是否为 p，如果是则直接返回当前节点。如果不是，将当前节点赋值给前继节点。
		往右走一步。
	如果走到这一步，说明不存在顺序后继，返回空。
 */
// 2022-03-23 刷出此题
// 注意BST 中 两种情况的分析
func inorderSuccessor(root *BiTreeNode, t *BiTreeNode) *BiTreeNode {
	found := false
	st := []*BiTreeNode{}
	p := root
	for p != nil || len(st) > 0{
		if p == t && p.Right != nil{ // 情况-1：连续往左
			p = p.Right
			for p.Left != nil{
				p = p.Left
			}
			return p
		}
		for p != nil{
			st = append(st, p)
			p = p.Left
		}
		top := st[len(st)-1]
		st = st[:len(st)-1]
		if found{  return top }
		if top == t{ found = true }
		p = top.Right
	}
	return nil
}