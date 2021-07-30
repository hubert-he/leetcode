package Tree
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

func PathSumIII(root *BiTreeNode, targetSum int) int {
	cnt := 0
	var dfs func(*BiTreeNode, map[int]bool, int)
	dfs = func(node *BiTreeNode, prefix map[int]bool, prefixSum int){
		if node == nil {
			return
		}
		prefixSum += node.Val.(int)
		prefix[prefixSum] = true
		defer func() { delete(prefix, prefixSum) }()
		if _, ok := prefix[prefixSum - targetSum]; ok {
			cnt++
		}
		dfs(node.Left, prefix, prefixSum)
		dfs(node.Right, prefix, prefixSum)
	}
	dfs(root, map[int]bool{0: true}, 0)
	return cnt
}

/* 666 Path Sum IV
If the depth of a tree is smaller than 5, then this tree can be represented by a list of three-digits integers.
For each integer in this list:
The hundreds digit represents the depth D of this node, 1 <= D <= 4.
The tens digit represents the position P of this node in the level it belongs to, 1 <= P <= 8. The position is the same as that in a full binary tree.
The units digit represents the value V of this node, 0 <= V <= 9.
Given a list of ascending three-digits integers representing a binary with the depth smaller than 5. You need to return the sum of all paths from the root towards the leaves.

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
func pathSumIV(nums []int) int {
	return 0
}