package Tree

func max(nums ...int)int{
	ans := nums[0]
	for i := 1; i < len(nums); i++{
		if nums[i] > ans {
			ans = nums[i]
		}
	}
	return ans
}

type Interface interface {
	Equal(target interface{}) bool
	Less(target interface{}) bool
//	Swap(target interface{})
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
	return &avlTreeNode{Key: value.(Interface), Val: value, Height: 1}
}

func NewAVLTree(value interface{}) *AVLTree{
	return &AVLTree{root: newNode(value), size: 1}
}

func(tree *AVLTree) GetSize() uint64{
	return tree.size
}

func(tree *AVLTree) GetHeigth() int{
	return getHeight(tree.root)
}
func getHeight(node *avlTreeNode) int{
	if node == nil {
		return 0
	}else{
		return node.Height
	}
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
		return true, newNode(key)
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
		root.Height = max(getHeight(root.Left), getHeight(root.Right)) + 1
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
	}
	return root
}

func getBalanceFactor(root *avlTreeNode) (balanceFactor int) {
	if root == nil {
		return 0
	}
	balanceFactor = getHeight(root.Left) - getHeight(root.Right)
	return
}

func leftRotation(root *avlTreeNode) *avlTreeNode{
	newRoot := root.Right
	root.Right = newRoot.Left
	newRoot.Left = root
	// Must update the height of node,that is newRoot & root
	root.Height = 1 + max(getHeight(root.Left), getHeight(root.Right)) // 先更新root，后更新newroot
	newRoot.Height = 1 + max(getHeight(newRoot.Left), getHeight(newRoot.Right))

	return newRoot
}

func rightRotation(root *avlTreeNode) *avlTreeNode{
	newRoot := root.Left
	root.Left = newRoot.Right
	newRoot.Right = root
	// Must update the height of node,that is newRoot & root
	root.Height = 1 + max(getHeight(root.Left), getHeight(root.Right)) // 必须先更新root，后更新newroot
	newRoot.Height = 1 + max(getHeight(newRoot.Left), getHeight(newRoot.Right))
	return newRoot
}

func(tree *AVLTree) Remove(key Interface) (target *avlTreeNode) {
	if tree.root == nil { // 创建avl 可以要求不能空创建，但是因为有删除操作，所以还是有avl是空树的可能
		return nil
	}
	/*短变量声明不支持如下语句 tree.root
	  Unlike regular variable declarations, a short variable declaration may redeclare variables provided they were originally declared earlier in the same block...with the same type,
	  and at least one of the non-blank variables is new.
	  It's not really a type inference issue, it's just that the left-hand-side of := must be a list of identifiers,
	  and tree.root is not an identifier, so it can't be declared — not even with :='s slightly-more-permissive rules for what it can declare.
	*/
	//isRemoved, tree.root, targetNode := remove(tree.root, key)
	isRemoved, newRoot, targetNode := remove(tree.root, key)
	tree.root = newRoot
	if isRemoved {
		tree.size--
	}
	if targetNode != nil {
		return targetNode
	}else{
		return nil
	}
}
/*
  isRemoved: 标记是否有节点被删除
  newRoot：返回给上一层平衡完成的新root节点
  target: 待删除的节点，如果无此key的节点，则返回nil
*/
func remove(root *avlTreeNode, key Interface) (isRemoved bool, newRoot *avlTreeNode, target *avlTreeNode){
	if root == nil {
		return
	}
	var retNode *avlTreeNode = nil
	if root.Key.Less(key){
		isRemoved, root.Right, target = remove(root.Right, key)
		retNode = root
	}else if key.Less(root.Key){
		isRemoved, root.Left, target= remove(root.Left, key)
		retNode = root
	}else{// found it<= Binary Sort Tree 形式删除
		isRemoved = true
		target = newNode(root.Val)
		*target = *root
		switch {
		case root.Left == nil:
			if root.Right == nil{ // 被删结点为叶子，直接返回
				newRoot = retNode
				return
			}
			retNode = root.Right
		case root.Right == nil:
			if root.Left == nil{
				newRoot = retNode
				return
			}
			retNode = root.Left
		default:
			// 左右子树不为nil,选右子树最小节点作为父节点，同样也可以选左子树最大节点
			retNode = getMinNodeInRight(root) // 肯定的一点 retNode没有左子树，即Left == nil，简化为上面2个情况
			root.Key = retNode.Key
			root.Val = retNode.Val
			_, root.Right, _ = remove(root.Right, retNode.Key)
			retNode = root
		}
	}
	root.Height = 1 + max(getHeight(root.Left), getHeight(root.Right))
	if isRemoved{
		newRoot = rebalance(retNode)
	} else{
		newRoot = retNode
	}
	return
	/*
	从已删除节点 w 开始，向上回溯，找到第一个不平衡点（平衡因子非-1 0 1）z；
	y 为 z 的高度最高的孩子节点； y 不再是从 w 回溯到z的路径上 z 的孩子
	x 为 y 的高度最高的孩子节点； x 不再是 z 的孙子
	对 z 为根节点的子树进行rebalance操作
	1. y 是 z 的左孩子， x 是 y 的左孩子 Left Left LL
	2. y 是 z 的左孩子， x 是 y 的右孩子 Left Right LR
	3. y 是 z 的右孩子， x 是 y 的右孩子 Right Right RR
	4. y 是 z 的右孩子， x 是 y 的左孩子 Right Left RL
	插入操作仅需要对以 z 为根的子树进行平衡操作；
	而平衡二叉树的删除操作就不一样，先对以 z 为根的子树进行平衡操作，之后可能需要对 z 的祖先结点进行平衡操作，向上回溯直到根结点
	 */
}

func getMinNodeInRight(root *avlTreeNode) (target *avlTreeNode){
	if root == nil {
		return nil
	}
	node := root.Right
	for node != nil{
		target = node
		node = node.Left
	}
	return
}

func getMaxNodeInLeft(root *avlTreeNode) (target *avlTreeNode){
	if root == nil{
		return nil
	}
	node := root.Left
	for node != nil {
		target = node
		node = node.Right
	}
	return
}

func(tree *AVLTree) Search(key Interface) (target *avlTreeNode) {
	if tree == nil || tree.root == nil {
		return nil
	}
	target = search(tree.root, key)
	return
}
func search(root *avlTreeNode, key Interface) *avlTreeNode{
	if root == nil {
		return nil
	}
	if root.Key.Equal(key){
		return root
	}
	if root.Key.Less(key){
		return search(root.Right, key)
	}else {
		return search(root.Left, key)
	}
}

func(tree *AVLTree) IsBST() bool{
	return true
}
// 剑指 Offer 55 - II
func(tree *AVLTree) IsBalanced() bool{
	if tree == nil {
		return true
	}
	var dfs func(*avlTreeNode)(int, bool)
	dfs = func(node *avlTreeNode)(height int, isbalanced bool){
		if node == nil {
			return 0,true
		}
		l, isbalanced := dfs(node.Left)
		if isbalanced == false{
			return l+1, false
		}
		r, isbalanced := dfs(node.Right)
		if isbalanced == false{
			return r+1, false
		}
		diff := l - r
		height = max(l, r) + 1
		switch {
		case diff < -1 || diff > 1:
			return height, false
		default:
			return height, true
		}
	}
	_, isB := dfs(tree.root)
	return isB
}
func isBalanced(root *avlTreeNode) (height int, balanced bool) {
	if root == nil{
		height = 0
		balanced = true
		return
	}
	var lh int
	var rh int
	var is_balanced bool
	lh, is_balanced = isBalanced(root.Left)
	if is_balanced{
		rh, is_balanced = isBalanced(root.Right)
	}else{
		return lh, false
	}
	height = max(lh, rh) + 1
	if is_balanced {
		diff := lh - rh
		switch {
		case diff < -1 || diff > 1:
			return height, false
		default:
			return height, true
		}
	}else{
		return height, false
	}
}

func(tree AVLTree) Serialization(){

}

func(tree AVLTree) Deserialization(){

}



