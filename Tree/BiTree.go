package Tree

import (
	"fmt"
	"time"
)

type BiTreeNode struct {
	Val   interface{}
	Left  *BiTreeNode
	Right *BiTreeNode
}

type ListNode struct {
    Val int
    Next *ListNode
}

const (
	PreOrder = iota
	PreOrderIter
	MidOrder
	MidOrderIter
	PostOrder
	PostOrderIter
	PostOrderIterII
	PostOrderIterIII
	LayerOrder
)

func GenerateBiTree(values []interface{}) *BiTreeNode {
	var cursor, root *BiTreeNode
	if len(values) == 0 {
		return nil
	}
	var Queue = []*BiTreeNode{}
	root = &BiTreeNode{values[0], nil, nil}
	Queue = append(Queue, root)
	index := 1
	for len(Queue) != 0 {
		var left, right interface{}
		if index < len(values) {
			left = values[index]
			index += 1
		}
		if index < len(values) {
			right = values[index]
			index += 1
		}
		cursor = Queue[0]
		Queue = Queue[1:]
		if left != nil {
			cursor.Left = &BiTreeNode{left, nil, nil}
			Queue = append(Queue, cursor.Left)
		}
		if right != nil {
			cursor.Right = &BiTreeNode{right, nil, nil}
			Queue = append(Queue, cursor.Right)
		}

	}
	return root
}

func PrintBiTree(root *BiTreeNode, t int) []interface{} {
	switch t {
	case PreOrder:
		return preOrder(root)
	case PreOrderIter:
		return preOrderIter(root)
	case PostOrder:
		return postOrder(root)
	case PostOrderIter:
		return postOrderIter(root)
	case PostOrderIterII:
		return postOrderIterII(root)
	case PostOrderIterIII:
		return postOrderIterIII(root)
	case MidOrder:
		return midOrder(root)
	case MidOrderIter:
		return midOrderIter(root)
	case LayerOrder:
		return layerOrderDFS(root)
	default:
		return nil
	}
}

func preOrder(root *BiTreeNode) []interface{} {
	var serial = []interface{}{}
	if root == nil {
		return serial
	}
	serial = append(serial, root.Val)
	serial = append(serial, preOrder(root.Left)...)
	serial = append(serial, preOrder(root.Right)...)
	return serial
}

func preOrderIter(root *BiTreeNode) []interface{} {
	var serial = []interface{}{}
	var stack = []*BiTreeNode{}
	if root == nil {
		return serial
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		// pop out
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		serial = append(serial, item.Val)
		// push in
		if item.Right != nil {
			stack = append(stack, item.Right)
		}
		if item.Left != nil {
			stack = append(stack, item.Left)
		}
	}
	return serial
}

func midOrder(root *BiTreeNode) []interface{} {
	var serial = []interface{}{}
	if root == nil {
		return serial
	}
	serial = append(serial, midOrder(root.Left)...)
	serial = append(serial, root.Val)
	serial = append(serial, midOrder(root.Right)...)
	return serial
}

/*
对于任一结点P，
  1)若其左孩子不为空，则将P入栈并将P的左孩子置为当前的P，然后对当前结点P再进行相同的处理；
  2)若其左孩子为空，则取栈顶元素并进行出栈操作，访问该栈顶结点，然后将当前的P置为栈顶结点的右孩子；
  3)直到P为NULL并且栈为空则遍历结束
*/
func midOrderIter(root *BiTreeNode) []interface{} {
	var serial = []interface{}{}
	stack := []*BiTreeNode{}
	var curr *BiTreeNode = root
	for curr != nil || len(stack) > 0 {
		for curr != nil { // curr 是否为nil 标志左子树是否已遍历,即如果右为空，栈要回退，不能再重复走左
			stack = append(stack, curr)
			curr = curr.Left
		}
		if len(stack) > 0 {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			serial = append(serial, curr.Val)
			curr = curr.Right // 注意
		}
	}
	return serial
}

func postOrder(root *BiTreeNode) []interface{} {
	var serial = []interface{}{}
	if root == nil {
		return serial
	}
	serial = append(serial, postOrder(root.Left)...)
	serial = append(serial, postOrder(root.Right)...)
	serial = append(serial, root.Val)
	return serial
}

/*
   要保证根结点在左孩子和右孩子访问之后才能访问，因此对于任一结点P，先将其入栈。
   如果P不存在左孩子和右孩子，则可以直接访问它；
   或者P存在左孩子或者右孩子，但是其左孩子和右孩子都已被访问过了，则同样可以直接访问该结点。
   若非上述两种情况，则将P的右孩子和左孩子依次入栈，这样就保证了每次取栈顶元素的时候，左孩子在右孩子前面被访问，左孩子和右孩子都在根结点前面被访问。
*/
func postOrderIter(root *BiTreeNode) []interface{} {
	var prev *BiTreeNode
	serial := []interface{}{}
	stack := []*BiTreeNode{root}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		if curr.Right == nil && curr.Left == nil {
			serial = append(serial, curr.Val)
			stack = stack[:len(stack)-1]
			prev = curr
		} else if prev != nil && (prev == curr.Left || prev == curr.Right) { //向上回溯
			serial = append(serial, curr.Val)
			stack = stack[:len(stack)-1]
			prev = curr
		} else {
			// 必须先右后左, 匹配栈的先进先出
			if curr.Right != nil {
				stack = append(stack, curr.Right)
			}
			if curr.Left != nil {
				stack = append(stack, curr.Left)
			}
		}
	}
	return serial
}

/*
此方法是 仿照前序遍历方式，然后再逆转。注意此与前序不同点 在于先left进栈， 后right进栈,即先访问右 再左
*/
func postOrderIterII(root *BiTreeNode) []interface{} {
	serial := []interface{}{}
	stack := []*BiTreeNode{root}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		serial = append(serial, top.Val)
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
	}
	// reverse in place
	for i, j := 0, len(serial)-1; i < j; i, j = i+1, j-1 {
		serial[i], serial[j] = serial[j], serial[i]
	}
	return serial
}

/*
   对于任一结点P，将其入栈，然后沿其左子树一直往下搜索，直到搜索到没有左孩子的结点，此时该结点出现在栈顶，但是此时不能将其出栈并访问，因此其右孩子还为被访问。
   所以接下来按照相同的规则对其右子树进行相同的处理，当访问完其右孩子时，该结点又出现在栈顶，此时可以将其出栈并访问。这样就保证了正确的访问顺序。
    可以看出，在这个过程中，每个结点都两次出现在栈顶，只有在第二次出现在栈顶时，才能访问它。因此需要多设置一个变量标识该结点是否是第一次出现在栈顶。
*/
func postOrderIterIII(root *BiTreeNode) []interface{} {
	serial := []interface{}{}
	type BiTreeNodeWithFlag struct {
		node     *BiTreeNode
		accessed bool
	}
	stack := []BiTreeNodeWithFlag{}
	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, BiTreeNodeWithFlag{curr, true})
			curr = curr.Left
		}
		if len(stack) > 0 {
			top := &stack[len(stack)-1] // 需要使用指针 改变slice中的值
			if top.accessed {
				// 第二次
				curr = top.node.Right
				top.accessed = false // 注意指向不同，值传递，需要使用指针
				//stack[len(stack) - 1].accessed = false
			} else {
				stack = stack[:len(stack)-1]
				serial = append(serial, top.node.Val)
				curr = nil // 必须为nil 结束子树访问
			}
		}
	}
	return serial
}

func layerOrder(root *BiTreeNode) []interface{} {
	t1 := time.Now()             // get current time
	var serial = []interface{}{} // []interface{}类型 是一个切片，切片元素的类型恰好是interface{}
	// var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
	if root == nil {
		return serial
	}
	var Queue = []*BiTreeNode{}
	Queue = append(Queue, root)
	for len(Queue) != 0 {
		if Queue[0] != nil {
			serial = append(serial, Queue[0].Val)
			Queue = append(Queue, Queue[0].Left)
			Queue = append(Queue, Queue[0].Right)
		} else {
			serial = append(serial, nil)
		}
		Queue = Queue[1:]
	}
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)
	return serial
}

func layerOrder2(root *BiTreeNode) []interface{} {
	t1 := time.Now()             // get current time
	var serial = []interface{}{} // []interface{}类型 是一个切片，切片元素的类型恰好是interface{}
	// var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
	if root == nil {
		return serial
	}
	var Queue = []*BiTreeNode{}
	//var Queue = make([]*BiTreeNode, 4,16)
	Queue = append(Queue, root)
	for len(Queue) != 0 {
		size := len(Queue)
		for i := 0; i < size; i++ {
			if Queue[i] != nil {
				serial = append(serial, Queue[i].Val)
				Queue = append(Queue, Queue[i].Left)
				Queue = append(Queue, Queue[i].Right)
			} else {
				serial = append(serial, nil)
			}
		}
		/*
			for _, item := range(Queue){
				if item != nil {
					serial = append(serial, item.Val)
					Queue = append(Queue, item.Left)
					Queue = append(Queue, item.Right)
				} else {
					serial = append(serial, nil)
				}
			}*/
		Queue = Queue[size:]
	}
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)
	return serial
}

func  layerOrderDFS(root *BiTreeNode) []interface{} {
	list := [][]interface{}{}

	layerDFS(&list, root, 0)
	serial := []interface{}{}
	// 转换二维slice为一维
	for _, item := range list {
		serial = append(serial, item...) // 通过初始化转换
	}
	return serial
}

func layerDFS(list *[][]interface{}, root *BiTreeNode, height int) {
	if root == nil {
		return
	}
	if height >= len(*list) {
		// new slice if nil
		*list = append(*list, []interface{}{root.Val})
	} else {
		(*list)[height] = append((*list)[height], root.Val)
	}
	layerDFS(list, root.Left, height+1)
	layerDFS(list, root.Right, height+1)
}

func Serialization(root *BiTreeNode) []interface{} {
	if root == nil {
		return nil
	}
	result := []interface{}{}
	q := []*BiTreeNode{root}
	for len(q) != 0{
		tmp := []*BiTreeNode{}
		var megerd bool = false
		for _, elem := range q{
			if elem != nil{
				result = append(result, elem.Val)
				if elem.Left != nil || elem.Right != nil{
					megerd = true
				}
				tmp = append(tmp, elem.Left, elem.Right)
			}else{
				result = append(result, nil)
			}
		}
		if megerd{
			q = tmp
		}else{
			q = nil
		}
	}
	return result
}

// leetcode-114  二叉树原地转换为单链表, 均用Left 连接
func flatten_left(root *BiTreeNode)  {
	if root == nil || root.Right == nil {
		return
	}
	flatten_left(root.Right)
	flatten_left(root.Left)
	pre,cur := root,root
	for cur != nil {
		pre = cur
		cur = cur.Left
	}
	pre.Left = root.Right
	root.Right = nil // 遗漏点： 把原来的链断开
}
// 都在Left 和 都在Right，注意区分，先序遍历
// 如果是flatten_left 这与先序遍历一致
// 若是flatten_right 与原先序遍历不一致，需要调整
func flatten_right(root *BiTreeNode) {
	if root == nil {
		return
	}
	tmp := root.Right
	flatten_right(root.Left)
	root.Right = root.Left
	flatten_right(tmp)
	pre,cur := root,root
	for cur != nil {
		pre = cur
		cur = cur.Right
	}
	pre.Right = tmp
	root.Left = nil
}
// leetcode-1367 Linked List in Binary Tree
func isSubPath(head *ListNode, root *BiTreeNode) bool {
	if root == nil {
		return false
	}
	return isSubDFS(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func isSubDFS(head *ListNode, root *BiTreeNode) bool{
	if head == nil { // 1. 链表已经全部匹配完，匹配成功，返回True
		return true
	}
	// 2. 二叉树访问到了空节点，匹配失败，返回False
	// 3. 当前匹配的二叉树上节点的值与链表节点的值不相等，匹配失败，返回 False
	if root == nil || root.Val != head.Val{
		return false
	}
	// 4. 前三种情况都不满足，说明匹配成功了一部分，我们需要继续递归匹配
	return isSubDFS(head.Next, root.Left) || isSubDFS(head.Next, root.Right)
}

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

func InvertBiTree(root *BiTreeNode) *BiTreeNode{
	if root == nil {
		return nil
	}
	tmp := root.Right
	root.Right = InvertBiTree(root.Left)
	root.Left  = InvertBiTree(tmp)
	return root
}

//257:  Binary Tree 的所有路径: root 到leaf的所有路径
func BinaryTreePaths(root *BiTreeNode) []string {
	if root == nil {
		return nil
	}
	paths := []string{}
	binaryTreePaths(root, "", &paths)
	return paths
}

func binaryTreePaths(root *BiTreeNode, subPath string, paths *[]string){
	if subPath == ""{
		subPath = fmt.Sprintf("%d", root.Val)
	}else{
		subPath = fmt.Sprintf("%s->%d", subPath, root.Val)
	}
	if root.Left == nil && root.Right == nil {
		*paths = append(*paths, subPath)
	}
	if root.Left != nil {
		binaryTreePaths(root.Left, subPath, paths)
	}
	if root.Right != nil {
		binaryTreePaths(root.Right, subPath, paths)
	}
}

//404: 左叶子之和
func SumOfLeftLeaves(root *BiTreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}
	sumlf := 0
	sumOfLeftLeaves(root, false, &sumlf)
	return sumlf
}
func sumOfLeftLeaves(root *BiTreeNode, fromLeft bool, sum *int) {
	if root.Left == nil && root.Right == nil {
		if fromLeft{
			*sum += root.Val.(int)
		}
	}
	if root.Left != nil {
		sumOfLeftLeaves(root.Left, true, sum)
	}
	if root.Right != nil {
		sumOfLeftLeaves(root.Right, false, sum)
	}
}

//563: 求坡度
func FindTilt(root *BiTreeNode) int {
	tilt, _ := findTilt(root)
	return tilt
}

func findTilt(root *BiTreeNode) (tiltSum, sum int){
	if root == nil {
		return 0,0
	}
	ltilt, lsum := findTilt(root.Left)
	rtilt, rsum := findTilt(root.Right)
	// self
	tiltSum += (ltilt + rtilt)
	if lsum > rsum{
		tiltSum += (lsum - rsum)
	}else {
		tiltSum += (rsum - lsum)
	}
	sum = lsum + rsum + root.Val.(int)
	return
}

//543: 求直径
func DiameterOfBinaryTree(root *BiTreeNode) int {
	diameter := 0
	diameterOfBinaryTree(root, &diameter)
	return diameter
}
func diameterOfBinaryTree(root *BiTreeNode, diameter *int) int {
	if root == nil {
		return 0
	}
	left := diameterOfBinaryTree(root.Left, diameter)
	right := diameterOfBinaryTree(root.Right, diameter)
	path := left + right // 是path ，而非高度
	if *diameter < path{
		*diameter = path
	}
	if left > right{
		return left + 1
	}else{
		return right + 1
	}
}

//235：LCA: 求最近公共祖先
func lowestCommonAncestor(root, p, q *BiTreeNode) *BiTreeNode {
	return nil
}
