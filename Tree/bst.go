package Tree

import "math"

type BinarySearchTree struct {
	root *BiTreeNode
}

func (tree * BinarySearchTree)GetRoot() *BiTreeNode{
	return tree.root
}

//501: 求众数： 找出 BST 中的所有众数（出现频率最高的元素）
func (tree *BinarySearchTree)FindMode() *BiTreeNode{
	return tree.root
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
	return
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

//501: 求众数： 找出 BST 中的所有众数（出现频率最高的元素）
func findMode(root *BiTreeNode) []int {
	return nil
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









