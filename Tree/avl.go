package Tree

type Interface interface {
	Less(target interface{}) bool
	Swap(target interface{})
}

type avlTreeNode struct {
	Key		Interface
	Val		interface{}
	Left	*avlTreeNode
	Right	*avlTreeNode
	Height	int
}

type AVLTree struct{
	size uint64
	root *avlTreeNode
}

func newNode(value interface{}) *avlTreeNode{
	return &avlTreeNode{Val: value}
}

func NewAVLTree(value interface{}) *AVLTree{
	return &AVLTree{root: newNode(value), size: 1}
}

func(tree *AVLTree) GetSize() uint64{
	return tree.size
}

func(tree *AVLTree) GetHeigth() int{
	return 0
}

func(tree *AVLTree) Insert(key Interface, value interface{}){
	added, newRoot := insert(tree.root, key, value)
	if added{
		tree.size++
	}
	tree.root = newRoot // 容易遗漏??
}
func insert(root *avlTreeNode, key Interface, value interface{})(added bool, newRoot *avlTreeNode){
	if root == nil {
		return true, &avlTreeNode{Key: key, Val: value}
	}
	if key == root.Key{
		// update value
		root.Val = value
		return false, root
	}else if root.Key.Less(key){
		added, root.Right = insert(root.Right, key, value)
	}else{
		added, root.Left = insert(root.Left, key, value)
	}
	if added {
		root.Height = max(root.Left.Height, root.Right.Height) + 1
		root = rebalance(root)
	}
	return added, root
}
func rebalance(root *avlTreeNode) *avlTreeNode{
	balanceFactor := getBalanceFactor(root)
	switch {
	case balanceFactor > 1:
		if getBalanceFactor(root.Left) < 0{
			root.Left = leftRotation(root.Left)
		}
		root = rightRotation(root)
	case balanceFactor < -1:
		if getBalanceFactor(root.Right) > 0{
			root.Right = rightRotation(root.Right)
		}
		root = leftRotation(root)
		leftRotation(root)
	}
	return root
}

func getBalanceFactor(root *avlTreeNode) (balanceFactor int) {
	if root == nil {
		return 0
	}
	balanceFactor = root.Left.Height - root.Right.Height
	return
}

func leftRotation(root *avlTreeNode) *avlTreeNode{
	newRoot := root.Right
	root.Right = newRoot.Left
	newRoot.Left = root
	// Must update the height of node,that is newRoot & root
	newRoot.Height = 1 + max(newRoot.Left.Height, newRoot.Right.Height)
	root.Height = 1 + max(root.Left.Height, root.Right.Height)
	return newRoot
}

func rightRotation(root *avlTreeNode) *avlTreeNode{
	newRoot := root.Left
	root.Left = newRoot.Right
	newRoot.Right = root
	// Must update the height of node,that is newRoot & root
	newRoot.Height = 1 + max(newRoot.Left.Height, newRoot.Right.Height)
	root.Height = 1 + max(root.Left.Height, root.Right.Height)
	return newRoot
}

func max(i, j int) int{
	if i > j {
		return i
	}else{
		return j
	}
}

func(tree *AVLTree) Remove(key interface{}) (value interface{}) {
	return nil
}

func(tree *AVLTree) Search(key interface{}) (value interface{}) {
	return nil
}

func(tree *AVLTree) IsBST() bool{
	return true
}

func(tree *AVLTree) IsBalanced() bool{
	return true
}

func(tree AVLTree) Serialization(){

}

func(tree AVLTree) Deserialization(){

}



