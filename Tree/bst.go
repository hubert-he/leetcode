package Tree

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
	j := len(nums) - 1
	i := 0
	if j >= i{
		mid := (j - i + 1) / 2
		root := &BiTreeNode{Val: nums[mid]}
		root.Left = sortedArrayToBST(nums[:mid])
		root.Right = sortedArrayToBST(nums[mid+1:])
		return root
	}else{
		return nil
	}
}