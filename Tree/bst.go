package Tree

type BinarySearchTree struct {
	root *BiTreeNode
}

func (tree * BinarySearchTree)GetRoot() *BiTreeNode{
	return tree.root
}

//501: 求众数： 找出 BST 中的所有众数（出现频率最高的元素）
func (tree * BinarySearchTree)FindMode() *BiTreeNode{
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

//501: 求众数： 找出 BST 中的所有众数（出现频率最高的元素）
func findMode(root *BiTreeNode) []int {
	return nil
}