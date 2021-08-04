package Tree

import "fmt"

/* 112. Path Sum
	Given the root of a binary tree and an integer targetSum,
	return true if the tree has a root-to-leaf path such that adding up all the values along the path equals targetSum.
	A leaf is a node with no children.
 */
func HasPathSum(root *BiTreeNode, targetSum int) bool {
	if root == nil {
		return false  // 每次开始先解决特殊情况
	}
	return hasPathSum(root, targetSum, 0)
}
// 注意，题目要求，是根到叶子节点所有和
func hasPathSum(root *BiTreeNode, targetSum, pathSum int) bool {
	pathSum += root.Val.(int)
	if root.Left == nil && root.Right == nil{
		if targetSum == pathSum{
			return true
		}else{
			return false
		}
	}
	if root.Left != nil && hasPathSum(root.Left, targetSum, pathSum) {
		return true
	}
	if root.Right != nil && hasPathSum(root.Right, targetSum, pathSum) {
		return true
	}
	return false
}
/* 最优题解 */
func hasPathSumBest(root *BiTreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	return hasPathSumBest(root.Left, targetSum - root.Val.(int)) || hasPathSumBest(root.Right, targetSum - root.Val.(int))
}
/* 113. Path Sum II
	Given the root of a binary tree and an integer targetSum,
	return all root-to-leaf paths where each path's sum equals targetSum.
	A leaf is a node with no children.
 */
func pathSumII(root *BiTreeNode, targetSum int) [][]int {
	var dfs func(*BiTreeNode, []int)
	path := [][]int{}
	dfs = func(node *BiTreeNode, log []int){
		if node == nil{
			return
		}
		if node.Left == nil && node.Right == nil{
			if targetSum == node.Val{
				log = append(log, node.Val.(int))
				// path = append(path, log) golang的坑点：log 是引用，后面结果可能会被覆盖
				tmp := make([]int, len(log))
				copy(tmp, log)
				path = append(path, tmp)
				return
			}else{
				return
			}
		}
		// 回溯
		tmp := log
		log = append(log, node.Val.(int))
		dfs(node.Left, log)
		dfs(node.Right, log)
		log = tmp
	}
	dfs(root, []int{})
	return path
}
// 官方题解： 使用 defer 恢复现场
func pathSumIILeetCode(root *BiTreeNode, targetSum int) (path [][]int) {
	log := []int{}
	var dfs func(*BiTreeNode, int)
	dfs = func(node *BiTreeNode, sum int){
		if node == nil {
			return
		}
		log = append(log, node.Val.(int))
		defer func() { log = log[:len(log) - 1] }() // 闭包递归函数中使用defer来恢复现场
		if node.Left == nil && node.Right == nil && sum == node.Val{
			path = append(path, append([]int{}, log...)) // 可以规避使用 copy 但是效率不高
			return
		}
		dfs(node.Left, sum - node.Val.(int))
		dfs(node.Right, sum - node.Val.(int))
	}
	dfs(root, targetSum)
	return
}
/* 437. Path Sum III
	Given the root of a binary tree and an integer targetSum,
	return the number of paths where the sum of the values along the path equals targetSum.
	The path does not need to start or end at the root or a leaf,
	but it must go downwards (i.e., traveling only from parent nodes to child nodes).
	前缀和 + 回溯 + map
 */
/*
	关联题目：
 */
func PathSumIII(root *BiTreeNode, targetSum int) int {
	cnt := 0
	// var dfs func(*BiTreeNode, map[int]bool, int)
	//dfs = func(node *BiTreeNode, prefix map[int]bool, prefixSum int){
	var dfs func(*BiTreeNode, map[int]int, int)
	dfs = func(node *BiTreeNode, prefix map[int]int, prefixSum int){
		if node == nil {
			return
		}
		prefixSum += node.Val.(int)
		// defer func() { delete(prefix, prefixSum) }()
		defer func() {
			prefix[prefixSum]--
			if prefix[prefixSum] == 0{
				delete(prefix, prefixSum)
			}
		}()
		if _, ok := prefix[prefixSum - targetSum]; ok {
			// cnt++
			cnt += prefix[prefixSum - targetSum]
		}
		/*
		if node.Val == targetSum                                                                                                                                                                                                                                                                        {
			cnt++
		}
		prefix[prefixSum] = true
		 */
		if prefixSum == targetSum{
			cnt++
		}
		prefix[prefixSum] += 1
		dfs(node.Left, prefix, prefixSum)
		dfs(node.Right, prefix, prefixSum)
	}
	// dfs(root, map[int]bool{0: true}, 0)
	dfs(root, map[int]int{}, 0)
	return cnt
}

/* 666 Path Sum IV
If the depth of a tree is smaller than 5, then this tree can be represented by a list of three-digits integers.
For each integer in this list:
The hundreds digit represents the depth D of this node, 1 <= D <= 4.
The tens digit represents the position P of this node in the level it belongs to, 1 <= P <= 8.
The position is the same as that in a full binary tree.
The units digit represents the value V of this node, 0 <= V <= 9.
  Given a list of ascending three-digits integers representing a binary with the depth smaller than 5.
You need to return the sum of all paths from the root towards the leaves.

Example 1:
Input: [113, 215, 221]
Output: 12
Explanation:
The tree that the list represents is:
    3
   / \
  5   1

The path sum is (3 + 5) + (3 + 1) = 12.

Example 2:
Input: [113, 221]
Output: 4
Explanation:
The tree that the list represents is:
    3
     \
      1

The path sum is (3 + 1) = 4.
 */
/*
  可以将每个结点的位置信息和结点值分离开，然后建立两者之间的映射
  数组是有序的，所以首元素就是根结点，然后我们进行先序遍历即可
  在递归函数中，我们先将深度和位置拆分出来，然后算出左右子结点的深度和位置的两位数，
  我们还要维护一个变量cur，用来保存当前路径之和。如果当前结点的左右子结点不存在，说明此时cur已经是一条完整的路径之和了，加到结果res中，直接返回。
  否则就是对存在的左右子结点调用递归函数即可
 */
func PathSumIV(nums []int) (ans int) {
	if len(nums) <= 0 {
		return
	}
	m := map[int]int{}
	for _, n := range nums{ // 分离位置信息 和 节点值信息
		m[n / 10] = n % 10
	}
	fmt.Println(m)
	var dfs func(int, int)
	dfs = func(num int, cur int){
		level, pos := num / 10, num % 10
		left := (level + 1) * 10 + 2 * pos - 1
		right := left + 1
		cur += m[num]
		_, hasLeft := m[left]
		_, hasRight := m[right]
		if hasLeft == false && hasRight == false{
			ans += cur
		}
		if hasLeft {
			dfs(left, cur)
		}
		if hasRight{
			dfs(right, cur)
		}
	}
	dfs(nums[0] / 10, 0)
	return
}
/* 关键点： 加上父节点的值
  BFS: 与先序遍历不同的是，我们不能维护一个当前路径之和的变量cur，这样会重复计算结点值，
       而是在遍历每一层的结点时，加上其父结点的值，如果某一个结点没有子结点了，才将累加起来的结点值加到结果ans
 */
func PathSumIVBFS(nums []int)(ans int){
	if len(nums) <= 0{
		return
	}
	root := nums[0] / 10
	m := map[int]int{}
	for _, n := range nums{
		m[n/10] = n % 10
	}
	q := []int{root}
	for len(q) > 0{
		node := q[0]
		q = q[1:]
		level, pos := node / 10, node % 10
		left := (level + 1) * 10 + 2 * pos -1
		right := left + 1
		_, hasLeft := m[left]
		_, hasRigh := m[right]
		if !hasLeft && !hasRigh {
			ans += m[node]
		}
		if hasLeft{
			m[left] += m[node]
			q = append(q, left)
		}
		if hasRigh{
			m[right] += m[node]
			q = append(q, right)
		}
	}
	return
}

/* 124. Binary Tree Maximum Path Sum
A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them.
A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.
The path sum of a path is the sum of the node's values in the path.
Given the root of a binary tree, return the maximum path sum of any path.
Example 1:
	Input: root = [1,2,3]
	Output: 6
	Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.
Example 2:
	Input: root = [-10,9,20,null,null,15,7]
	Output: 42
	Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.
 */
func maxPathSum(root *TreeNode) int {
	return 0
}