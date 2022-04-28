package Tree

import (
	"fmt"
	"math"
	"time"
)

type BiTreeNode struct {
	Val		interface{}
	Left	*BiTreeNode
	Right	*BiTreeNode
	Next	*BiTreeNode // for next problem
	Parent	*BiTreeNode
}

type ListNode struct {
    Val int
    Next *ListNode
}

const (
	PreOrder = iota
	PreOrderIter
	PreOrderMorris
	MidOrder
	MidOrderIter
	MidOrderMorris
	PostOrder
	PostOrderIter
	PostOrderIterII
	PostOrderIterIII
	PostOrderMorris
	LayerOrder
)

func GenerateBiTree(values []interface{}) *BiTreeNode {
	var cursor, root *BiTreeNode
	if len(values) == 0 {
		return nil
	}
	var Queue = []*BiTreeNode{}
	root = &BiTreeNode{Val: values[0], Left: nil, Right: nil}
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
			cursor.Left = &BiTreeNode{Val: left, Left: nil, Right: nil}
			Queue = append(Queue, cursor.Left)
		}
		if right != nil {
			cursor.Right = &BiTreeNode{Val: right, Left: nil, Right: nil}
			Queue = append(Queue, cursor.Right)
		}

	}
	return root
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
	// TODO: 如何清理掉最后多余的nil
	last := len(result)
	for last > 0{
		if result[last-1] != nil {
			break;
		}
		last--
	}
	return result[:last]
}
/* 思维方式不对
func Tree2str(root *BiTreeNode) string{
	serial := []byte{}
	var dfs func(node *BiTreeNode)
	dfs = func(node *BiTreeNode){
		if node == nil{
			serial = append(serial, '(', ')')
			return
		}
		v := byte(node.Val.(int))
		serial  = append(serial, '(', '0'+v)
		if node.Left != nil || node.Right != nil{
			dfs(node.Left)
			dfs(node.Right)
		}else if node.Left != nil{
			dfs(node.Left)
		}else if node.Right != nil{
			dfs(node.Right)
		}
		serial = append(serial, ')' )
	}
	dfs(root)
	return string(serial[1:len(serial)-1])
}
 */
/*
分4种情况：
	1. 如果当前节点有两个孩子，那我们在递归时，需要在两个孩子的结果外都加上一层括号；
	2. 如果当前节点没有孩子，那我们不需要在节点后面加上任何括号；
	3. 如果当前节点只有左孩子，那我们在递归时，只需要在左孩子的结果外加上一层括号，而不需要给右孩子加上任何括号；
	4. 如果当前节点只有右孩子，那我们在递归时，需要先加上一层空的括号()表示左孩子为空，再对右孩子进行递归，并在结果外加上一层括号。
 */
func Tree2str(root *BiTreeNode) string {
	if root == nil{
		return ""
	}
	v := fmt.Sprintf("%d", root.Val.(int))
	if root.Left == nil && root.Right == nil{
		return v
	}
	if root.Right == nil{
		return fmt.Sprintf("%s(%s)", v, Tree2str(root.Left))
	}
	return fmt.Sprintf("%s(%s)(%s)", v, Tree2str(root.Left), Tree2str(root.Right))
}

func Tree2strIter(root *BiTreeNode) (serial string){
	if root == nil{
		return ""
	}
	stack := []*BiTreeNode{root}
	// 迭代得到前序遍历的方法略有不同，由于这里需要输出额外的括号，因此我们还需要用一个集合存储所有遍历过的节点
	// 当再次访问到节点时，出栈 以及 加闭括号 )
	visited := map[*BiTreeNode]interface{}{}  // 节省资源的 map 只关心key
	for len(stack) > 0{
		top := stack[0] // GoLang 中栈的便宜操作
		if _, ok := visited[top]; ok {
			stack = stack[1:]
			serial += ")"
		}else{
			visited[top]=nil
			// 先在答案末尾添加一个 (，表示一个节点的开始，然后判断该节点的子节点个数
			// 根据4种情况分别设置
			serial = fmt.Sprintf("%s(%d", serial, top.Val.(int))
			if top.Left == nil && top.Right != nil{
				// 如果它只有右孩子，那么我们在答案末尾添加一个 () 表示空的左孩子，再将右孩子入栈
				serial += "()"
			}
			if top.Right != nil{ // 前序遍历，右子树先入栈
				stack = append([]*BiTreeNode{top.Right}, stack...)
			}
			if top.Left != nil{
				stack = append([]*BiTreeNode{top.Left}, stack...)
			}
		}
	}
	return serial[1:len(serial)-1]
}
func PrintBiTree(root *BiTreeNode, t int) []interface{} {
	switch t {
	case PreOrder:
		return preOrder(root)
	case PreOrderIter:
		return preOrderIter(root)
	case PreOrderMorris:
		return preOrderMorris(root)
	case PostOrder:
		return postOrder(root)
	case PostOrderIter:
		return postOrderIter(root)
	case PostOrderIterII:
		return postOrderIterII(root)
	case PostOrderIterIII:
		return postOrderIterIII(root)
	case PostOrderMorris:
		return postOrderMorris(root)
	case MidOrder:
		return midOrder(root)
	case MidOrderIter:
		return midOrderIter(root)
	case MidOrderMorris:
		return midOrderMorris(root)
	case LayerOrder:
		return layerOrderDFS(root)
	default:
		return nil
	}
}

func preOrderMorris(root *BiTreeNode) (list []interface{}) {
	cur := root
	for cur != nil {
		if cur.Left == nil{
			list = append(list, cur.Val)
			cur = cur.Right
			continue
		}
		prev := cur.Left
		for prev.Right != nil && prev.Right != cur{
			prev = prev.Right
		}
		if prev.Right == nil{
			list = append(list, cur.Val)
			prev.Right = cur
			cur = cur.Left
		}else{ // prev.Right == prev的序列后继cur
			prev.Right = nil // 恢复树结构
			cur = cur.Right
		}
	}
	return
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

func midOrderMorris(root *BiTreeNode) (list []interface{}){
	cur := root
	for cur != nil{
		if cur.Left == nil{
			list = append(list, cur.Val)
			cur = cur.Right
			continue
		}
		prev := cur.Left
		for prev.Right != nil && prev.Right != cur{
			prev = prev.Right
		}
		if prev.Right == nil{
			prev.Right = cur
			cur = cur.Left
		}else{
			list = append(list, cur.Val)
			prev.Right = nil
			cur = cur.Right
		}
	}
	return
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

func postOrderMorris(root *BiTreeNode) (serial []interface{}) {
	dumbRoot := &BiTreeNode{Val: nil, Left: root, Right: nil}
	cur := dumbRoot
	for cur != nil {
		if cur.Left == nil {
			cur = cur.Right
			continue
		}
		predecessor := cur.Left
		for predecessor.Right != nil && predecessor.Right != cur {
			predecessor = predecessor.Right
		}
		if predecessor.Right == nil {
			predecessor.Right = cur
			cur = cur.Left
		}else{
			predecessor.Right = nil
			tmp := []interface{}{}
			for i := cur.Left; i != nil; i = i.Right{
				tmp = append([]interface{}{i.Val}, tmp...)
			}
			serial = append(serial, tmp...)
			cur = cur.Right
		}
	}
	return
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
	if root == nil {
		return nil
	}
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
func postOrderIterII(root *BiTreeNode) (serial []interface{}) {
	stack := []*BiTreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			// 逆向插入，不需要后面整体reverse: 后序遍历： 左 右 根
			serial = append([]interface{}{cur.Val}, serial...)
			stack = append([]*BiTreeNode{cur}, stack...)
			cur = cur.Right // 因为是逆向插入，先插入右边
		}else{
			cur = stack[0].Left
			stack = stack[1:]
		}
	}
	return
}
/* 不够简洁
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
 */

/*
   对于任一结点P，将其入栈，然后沿其左子树一直往下搜索，直到搜索到没有左孩子的结点，此时该结点出现在栈顶，但是此时不能将其出栈并访问，
   因为其右孩子还为被访问。
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

func Find(root *BiTreeNode, key interface{}) (node *BiTreeNode){
	if root == nil {
		return
	}
	st := []*BiTreeNode{root}
	for len(st) > 0 {
		if st[0].Val == key{
			return st[0]
		}
		tmp := []*BiTreeNode{}
		if st[0].Left != nil {
			tmp = append(tmp, st[0].Left)
		}
		if st[0].Right != nil {
			tmp = append(tmp, st[0].Right)
		}
		st = append(tmp, st[1:]...)
	}
	return
}

func addParent(root *BiTreeNode) ( map[*BiTreeNode]*BiTreeNode) {
	parents := make(map[*BiTreeNode]*BiTreeNode)
	var dfs func(*BiTreeNode, *BiTreeNode)
	dfs = func(node, parent *BiTreeNode){
		if node != nil{
			parents[node] = parent
			dfs(node.Left, node)
			dfs(node.Right, node)
		}
	}
	dfs(root, nil)
	return parents
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

//235：LCA: 求最近公共祖先-Lowest Common Ancestor
func LowestCommonAncestor(root, p, q *BiTreeNode) *BiTreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val{
		// 找到目标节点后，直接返回，不需要继续向下遍历
		return root
	}
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		// 其左右子树存在p和q，即找到LCA
		return root
	}
	if left == nil{
		return right
	}
	return left
}

func LowestCommonAncestorII(root, p, q *BiTreeNode) *BiTreeNode {
	var lca *BiTreeNode
	var dfs func(node *BiTreeNode) *BiTreeNode
	dfs = func(node *BiTreeNode) *BiTreeNode{
		if node == nil {
			return  nil
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		// node是lca 区分2种情况，node 是 p q 其中一个，或者 left right 都不为空
		if node == p || node == q{
			if left != nil || right != nil {
				lca = node
			}
			return node
		}else{
			if left != nil && right != nil {
				lca = node
				return node
			}
			if left != nil || right != nil {
				return node
			}
		}
		return nil
	}
	dfs(root)
	return lca
}

/* golang map用法 - 1
func LowestCommonAncestorHashMap(root, p, q *BiTreeNode) (lca *BiTreeNode) {
	// 匿名结构体充当map值，后面赋值不方便
	type Parent map[*BiTreeNode]struct{
		parent	*BiTreeNode
		visited	bool
	}
	parMap := Parent{}
	var dfs func(node *BiTreeNode)*BiTreeNode
	dfs = func(node *BiTreeNode)*BiTreeNode{
		if node == nil{
			return nil
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		if left != nil{
			parMap[left] = struct{
				parent	*BiTreeNode
				visited	bool
			}{node, false}
		}
		if right != nil{
			parMap[right] = struct{
				parent	*BiTreeNode
				visited	bool
			}{node, false}
		}
		return node
	}
	dfs(root)
	for i := p; i != nil;{
		if _, ok := parMap[i]; !ok{
			break
		}
		parMap[i] = struct{
			parent	*BiTreeNode
			visited	bool
		}{parMap[i].parent, true}
		i = parMap[i].parent
	}
	for i := q; i != nil;{
		if _, ok := parMap[i]; !ok || parMap[i].visited{
			lca = i
			break
		}
		i = parMap[i].parent
	}
	return lca
}
*/
func LowestCommonAncestorHashMap(root, p, q *BiTreeNode) (lca *BiTreeNode) {
	type ParNodeInfo struct {
		parent	*BiTreeNode
		visited	bool
	}
	type Parent map[*BiTreeNode]ParNodeInfo
	parMap := Parent{}  // parMap 值为 Tree.Parent{}
	var parMap2 Parent  // parMap2 值为 nil
	fmt.Printf("%#v  %#v %#v\n", parMap, parMap2, parMap[nil])
	// 输出 Tree.Parent{}  Tree.Parent(nil)  Tree.ParNodeInfo{parent:(*Tree.BiTreeNode)(nil), visited:false}
	// 如果map中对应的key不存在，就返回： value类型的零值
	var dfs func(node *BiTreeNode)*BiTreeNode
	dfs = func(node *BiTreeNode)*BiTreeNode{
		if node == nil{
			return nil
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		if left != nil{
			parMap[left] = ParNodeInfo{node, false}
		}
		if right != nil{
			parMap[right] = ParNodeInfo{node, false}
		}
		return node
	}
	dfs(root)
	for i := p; i != nil;{
		if _, ok := parMap[i]; !ok{
			break
		}
		// 非指针，嵌入结构体必须整体赋值
		parMap[i] = ParNodeInfo{parMap[i].parent, true}
		i = parMap[i].parent
	}
	for i := q; i != nil;{
		if _, ok := parMap[i]; !ok || parMap[i].visited{
			lca = i
			break
		}
		i = parMap[i].parent
	}
	return lca
}
// LCA 的离线解决： Tarjan算法
/*
   Tarjan算法是离线算法，即它不会在查询两个节点 LCA 这个事件发生时就及时的执行并给出答案，而是等到所有查询都给出后统一进行处理和求解
   算法从根节点root开始搜索，每次递归搜索所有的子树，然后处理跟当前根节点相关的所有lca查询，也就是说把树的每个子树当做一个并查集来查询
   伪代码：
   Tarjan(u)//marge和find为并查集合并函数和查找函数
   {
    	for each(u,v)    //访问所有u子节点v
    	{
        	Tarjan(v);        //继续往下遍历
        	marge(u,v);    //合并v到u上
        	标记v被访问过;
    	}
    	for each(u,e)    //访问所有和u有询问关系的e
    	{
        	如果e被访问过;
        	u,e的最近公共祖先为find(e);
    	}
	}
 */

//572: subtree of another tree
// 在判断「ss 的深度优先搜索序列包含 tt 的深度优先搜索序列」的时候，可以暴力匹配，也可以使用KMP 或者Rabin-Karp 算法，
//在使用Rabin-Karp 算法的时候，要注意串中可能有负值。

func IsSubtree_kmp(s, t *BiTreeNode) bool{
	maxEle := math.MinInt32
	getMaxElement(s, &maxEle)
	getMaxElement(t, &maxEle)
	lNull := maxEle + 1
	rNull := maxEle + 2
	s1, t1 := getDfsOrder(s, []int{}, lNull, rNull), getDfsOrder(t, []int{}, lNull, rNull)
	return kmp(s1, t1)
}
func kmp(s, t []int) bool{
	sLen, tLen := len(s), len(t)
	fail := make([]int, tLen)
	for i := 0; i < tLen; i++ {
		fail[i] = -1
	}
	for i, j := 1, -1; i < tLen; i++{
		for j != -1 && t[i] != t[j+1]{
			j = fail[j]
		}
		if t[i] == t[j+1]{
			j++
		}
		fail[i] = j
	}
	// 主搜索部分
	for i, j := 0, -1; i < sLen; i++ {
		// 内涵一层：失配，并且为开始状态，直接i++ 主串进一步
		for j != -1 && s[i] != t[j+1]{ // 失配，若非起始状态，则依据next数组移动位置
			j = fail[j]
		}
		if s[i] == t[j+1]{ //若两个字符相等，则主串和模式串各进一步
			j++
		}
		if j == tLen - 1{ // 匹配成功，返回。注：标准KMP算法中，是重置状态，继续匹配下一个位置
			return true
		}
	}
	return false
}
func getDfsOrder(t *BiTreeNode, list []int, lNull, rNull int) []int{
	if t == nil{
		return list
	}
	list = append(list, t.Val.(int))
	if t.Left != nil {
		list = getDfsOrder(t.Left, list, lNull, rNull)
	}else{
		list = append(list, lNull)
	}
	if t.Right != nil {
		list = getDfsOrder(t.Right, list, lNull, rNull)
	}else{
		list = append(list, rNull)
	}
	return list
}
func getMaxElement(t *BiTreeNode, maxEle *int){
	if t == nil {
		return
	}
	if t.Val.(int) > *maxEle{
		*maxEle = t.Val.(int)
	}
	getMaxElement(t.Left, maxEle)
	getMaxElement(t.Right, maxEle)
}
func IsSubtree_2dfs(s *BiTreeNode, t*BiTreeNode) bool{
	if s == nil {
		return false
	}
	return isEqual(s,t) || IsSubtree_2dfs(s.Left, t) || IsSubtree_2dfs(s.Right, t)
}
func isEqual(s *BiTreeNode, t *BiTreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s== nil || t == nil { // 上次条件判断，排除了全nil情况
		return false
	}
	if s.Val == t.Val{
		return isEqual(s.Left, t.Left) && isEqual(s.Right, t.Right)
	}
	return false
}
/* 实现比较渣
func checkSubTree(t1 *BiTreeNode, t2 *BiTreeNode) bool {
	var dfs func(root *BiTreeNode)bool
	dfs = func(root *BiTreeNode)bool{
		if root == nil && t2 != nil{ // 易漏1
			return false
		}
		if isEqual(root, t2){
			return true
		}
		return dfs(root.Left) || dfs(root.Right)
		// 错误-2
		//return isEqual(root.Left, t2) || isEqual(root.Right, t2)
	}
	return dfs(t1)
}
func isEqual(t1, t2 *BiTreeNode)bool{
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil{
		return false
	}
	if t1.Val != t2.Val{
		return false
	}
	return isEqual(t1.Left, t2.Left) && isEqual(t1.Right, t2.Right)
}
 */
// 剑指 Offer 27. 二叉树的镜像
func MirrorTree(root *BiTreeNode) *BiTreeNode {
	var dfs func(*BiTreeNode)
	dfs = func(node *BiTreeNode){
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		tmp := node.Left
		node.Left = node.Right
		node.Right = tmp
	}
	dfs(root)
	return root
}
// 剑指 Offer 28. Symmetric Binary tree
func IsSymmetric(root *BiTreeNode) bool {
	var dfs func(*BiTreeNode, *BiTreeNode)bool
	dfs = func(node1 *BiTreeNode, node2 *BiTreeNode) bool{
		if node1 == nil && node2 == nil {
			return true
		}
		if node1 == nil || node2 == nil {
			return false
		} // 提前结束
		if dfs(node1.Left, node2.Right){
			if node1.Val != node2.Val{
				return false
			}else{
				return dfs(node1.Right, node2.Left)
			}
		}else{
			return false
		}
	}
	return  dfs(root, root)
}
// 更简短，先序遍历
func IsSymmetric2(root *BiTreeNode) bool {
	if root == nil {
		return true
	}
	var dfs func(*BiTreeNode, *BiTreeNode)bool
	dfs = func(node1 *BiTreeNode, node2 *BiTreeNode)bool{
		if node1 == nil && node2 == nil {
			return true
		}
		if node1 == nil || node2 == nil || node1.Val != node2.Val{
			return false
		}
		return dfs(node1.Left, node2.Right) && dfs(node1.Right, node2.Left)
	}
	return dfs(root.Left, root.Right)
}

//993. 二叉树的堂兄弟节点
func isCousins(root *BiTreeNode, x int, y int) bool {
	xh, yh := -1, -1
	var xpar, ypar *BiTreeNode
	var dfs func(*BiTreeNode, *BiTreeNode, int)
	dfs = func(node, par *BiTreeNode, depth int){
		if node == nil {
			return
		}
		v := node.Val
		if v == x{
			xh = depth
			xpar = par
		}
		if v == y{
			yh = depth
			ypar = par
		}
		dfs(node.Left, node, depth+1)
		dfs(node.Right,node, depth+1)
	}
	dfs(root, nil, 0)
	if xh != -1 && xh == yh && xpar != nil && xpar != ypar{
		return true
	}else{
		return false
	}
}

//872. 叶子相似的树
func LeafSimilar(root1 *BiTreeNode, root2 *BiTreeNode) bool {
	s1 := LeafSequence(root1)
	s2 := LeafSequence(root2)
	if len(s1) != len(s2){
		return false
	}
	for idx, value := range s1{
		if value != s2[idx]{
			return false
		}
	}
	return true
}

func LeafSequence(root *BiTreeNode)(leafs []interface{}){
	var dfs func(node *BiTreeNode)
	if root == nil{
		return
	}
	dfs = func(node *BiTreeNode){
		if node.Left == nil && node.Right == nil{
			leafs = append(leafs, node.Val)
		}
		if node.Left != nil{
			dfs(node.Left)
		}
		if node.Right != nil{
			dfs(node.Right)
		}
	}
	dfs(root)
	return
}
//653: two sum问题，可以双指针或者hashMap
func FindTwoSum(root *BiTreeNode, sum int)(target []*BiTreeNode) {
	hashMap := map[int]*BiTreeNode{}
	var dfs func(node *BiTreeNode)
	dfs = func(node *BiTreeNode){
		if node == nil{
			return
		}
		u := node.Val.(int)
		if v, ok := hashMap[sum - u]; ok{
			target = append(target, node, v)
		}else{
			hashMap[u] = node
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return
}

// 1469 寻找所有的独生节点
/*
二叉树中，如果一个节点是其父节点的唯一子节点，则称这样的节点为 “独生节点” 。
 二叉树的根节点不会是独生节点，因为它没有父节点。
给定一棵二叉树的根节点 root ，返回树中 所有的独生节点的值所构成的数组 。
 数组的顺序 不限 。
示例 1：
输入：root = [1,2,3,null,4]
输出：[4]
解释：浅蓝色的节点是唯一的独生节点。
节点 1 是根节点，不是独生的。
节点 2 和 3 有共同的父节点，所以它们都不是独生的。
提示：
tree 中节点个数的取值范围是 [1, 1000]。
每个节点的值的取值范围是 [1, 10^6]。
*/
func GetLonelyNodes(root *BiTreeNode) (nodes []interface{}) {
	var dfs func(node, parent *BiTreeNode)
	dfs = func(node, parent *BiTreeNode){
		if node == nil{
			return
		}
		if parent != nil{
			if !(parent.Left != nil && parent.Right != nil){
				nodes = append(nodes, node.Val)
			}
		}
		dfs(node.Left, node)
		dfs(node.Right, node)
	}
	dfs(root, nil)
	return
}
// 623. Add One Row to Tree
func AddOneRow(root *BiTreeNode, val int, depth int) *BiTreeNode {
	if depth == 1{
		return &BiTreeNode{Val: val, Left: root}
	}
	var dfs func(*BiTreeNode, int)
	dfs = func(node *BiTreeNode, height int){
		if node == nil{
			return
		}
		if height == depth - 1{
			node.Left = &BiTreeNode{Val: val, Left: node.Left}
			node.Right = &BiTreeNode{Val: val, Right: node.Right}
		}
		dfs(node.Left, height + 1)
		dfs(node.Right, height + 1)
	}
	dfs(root, 1)
	return root
}
// depth直接与高度一致来做判断，条件判断多
func AddOneRow1(root *BiTreeNode, val int, depth int) *BiTreeNode {
	if depth <= 1{
		if root != nil{
			return &BiTreeNode{Val: val, Left: root}
		}else{
			return &BiTreeNode{Val: val}
		}
	}
	var dfs func(*BiTreeNode, *BiTreeNode, int)
	dfs = func(node, parent *BiTreeNode, height int){
		if node == nil{
			return
		}
		if height == depth-1{
			node.Left = &BiTreeNode{Val:val, Left: node.Left}
			node.Right = &BiTreeNode{Val:val, Right: node.Right}
			return
		}
		dfs(node.Left, node, height+1)
		dfs(node.Right, node, height+1)
	}
	dfs(root, nil, 1)
	return root
}
func AddOneRow2(root *BiTreeNode, val int, depth int) *BiTreeNode {
	if depth <= 1{
		if root != nil{
			return &BiTreeNode{Val: val, Left: root}
		}else{
			return &BiTreeNode{Val: val}
		}
	}
	var dfs func(*BiTreeNode, *BiTreeNode, int, bool)
	dfs = func(node, parent *BiTreeNode, height int, left bool){
		if height == depth{
			if left {
				parent.Left = &BiTreeNode{Val: val, Left: node}
			}else{
				parent.Right = &BiTreeNode{Val: val, Right: node}
			}
			return
		}
		if node == nil{
			return
		}
		dfs(node.Left, node, height+1, true)
		dfs(node.Right, node, height+1, false)
	}
	dfs(root, nil, 1, true)
	return root
}

// 652. Find Duplicate Subtrees
func FindDuplicateSubtrees(root *BiTreeNode) (subTreeRoot []*BiTreeNode) {
	serials := map[string]int{}
	count := map[string]bool{}
	var dfs func(node *BiTreeNode)int
	factor := 1
	dfs = func(node *BiTreeNode) int{
		if node == nil{
			return 0
		}
		tupleString := fmt.Sprintf("%d#%d#%d", node.Val, dfs(node.Left), dfs(node.Right))
		if _,ok := serials[tupleString]; !ok{
			factor++
			serials[tupleString] = factor
			count[tupleString] = true
		}else{
			if count[tupleString]{
				subTreeRoot = append([]*BiTreeNode{node}, subTreeRoot...)
				count[tupleString] = false
			}
		}
		return serials[tupleString]
	}
	dfs(root)
	return
}

func FindDuplicateSubtrees_TupleThing(root *BiTreeNode) (subTreeRoot []*BiTreeNode) {
	type tuple struct{
		root,left,right int
	}
	serials := map[tuple]int{}
	count := map[tuple]bool{}
	var dfs func(node *BiTreeNode)int
	factor := 1
	dfs = func(node *BiTreeNode) int{
		if node == nil{
			return 0
		}
		tupleThing := tuple{node.Val.(int), dfs(node.Left), dfs(node.Right)}
		if _,ok := serials[tupleThing]; !ok{
			factor++
			serials[tupleThing] = factor
			count[tupleThing] = true
		}else{
			if count[tupleThing]{
				subTreeRoot = append([]*BiTreeNode{node}, subTreeRoot...)
				count[tupleThing] = false
			}
		}
		return serials[tupleThing]
	}
	dfs(root)
	return
}

// 863. All Nodes Distance K in Binary Tree
func DistanceK(root *BiTreeNode, target *BiTreeNode, K int) []int {
	//adj := make([][]bool, 8)
	adj := [][]bool{}
	nodes := map[*BiTreeNode]int{}
	index := 0
	var dfs func(node, parent *BiTreeNode)
	dfs = func(node, parent *BiTreeNode){
		if node == nil {
			return
		}
		nodes[node] = index
		index++
		adj = append(adj, make([]bool, index))
		if parent != nil{
			adj[nodes[node]][nodes[parent]] = true
			//adj[nodes[parent]][nodes[node]] = true 二叉树的邻接矩阵是个对称矩阵
		}
		dfs(node.Left, node)
		dfs(node.Right, node)
	}
	dfs(root, nil)
	// 输出邻接矩阵
	//fmt.Println(adj)
	t := []int{nodes[target]}
	pre := []int{}
	for i := 0; i < K; i++{
		tmp := []int{}
		for _, value := range t{
			length := len(adj[value])
			for i := 0; i < len(adj); i++{
				v := false
				if length <= i{
					v = adj[i][value]
				}else{
					v = adj[value][i]
				}
				if v {
					m := 0
					for m < len(pre){
						if pre[m] == i{
							break
						}
						m++
					}
					if m >= len(pre){
						tmp = append(tmp, i)
					}
				}
			}
		}
		pre = t
		t = tmp
	}
	answer := []int{}
	for _, v := range t{
		for key,value := range nodes{
			if v == value{
				answer = append(answer, key.Val.(int))
				break
			}
		}
	}
	fmt.Println(answer)
	return nil
}

func DistanceK2(root *BiTreeNode, target *BiTreeNode, K int) []interface{} {
	answer := []interface{}{}
	var dfs func(node *BiTreeNode) int
	var findInSubTree func(node *BiTreeNode, disc int)
	dfs = func(node *BiTreeNode) int{
		if node == nil {
			return -1
		}
		if node == target{
			findInSubTree(target, K)
			return K
		}
		lk := dfs(node.Left)
		rk := dfs(node.Right)
		switch {
		case lk < 0 && rk < 0:
			return -1
		case lk < 0:
			if rk == 1{
				answer = append(answer, node.Val)
			}else{
				findInSubTree(node.Left, rk - 2)
			}
			return rk - 1
		case rk < 0:
			if lk == 1{
				answer = append(answer, node.Val)
			}else{
				findInSubTree(node.Right, lk - 2)
			}
		}
		return lk - 1
	}
	findInSubTree = func(node *BiTreeNode, disc int){
		if node != nil {
			if disc == 0{
				answer = append(answer, node.Val)
			}else{
				findInSubTree(node.Left, disc - 1)
				findInSubTree(node.Right, disc - 1)
			}
		}
	}
	dfs(root)
	return answer
}

func DistanceK3(root, target *BiTreeNode, K int) (answer []interface{}) {
	var dfs func(node *BiTreeNode) (disc int)
	var findKinSubtree func(*BiTreeNode, int)
	dfs = func(node *BiTreeNode)(disc int){
		if node == nil{
			return -1
		}
		// 情况1：子树中距离target节点的距离为K的所有节点加入answer
		if node == target{
			findKinSubtree(node, 0)
			return 0 // 定义target到target节点的距离位0
			//return 1 // 定义target到target节点的距离位1
		}
		// target 节点必定在node的左子树或者右子树上，不可能同时在左和右子树
		lk := dfs(node.Left)
		rk := dfs(node.Right)
		// 情况2： target在node左子树中：假设target距离node的距离为lk+1，找出右子树中距离target节点K-lk-1距离的所有节点加入结果集
		if lk != -1 {
		//  if lk == K { // 定义target到target节点的距离为1
			if lk + 1 == K{// 定义target到target节点的距离为0
				answer = append(answer, node.Val)
			}
			//转右子树找K-lk-1的节点
			//findKinSubtree(node.Right, lk + 1) // 定义target到target节点的距离位1
			findKinSubtree(node.Right, lk + 2) // 定义target到target节点的距离为0，+2原因是左右两支
			return lk + 1
		}
		// 情况3：target在node的右子树中
		if rk != -1 {
		//	if rk == K{ // 定义target到target节点的距离为1
			if rk + 1 == K{ // 定义target到target节点的距离为0
				answer = append(answer, node.Val)
			}
			//findKinSubtree(node.Left, rk + 1)
			findKinSubtree(node.Left, rk + 2)
			return rk + 1
		}
		//情况4：target 不在节点的子树中，无需处理
		return -1
	}
	findKinSubtree = func(node *BiTreeNode, disc int){
		if node == nil{
			return
		}
		if disc == K{
			answer = append(answer, node.Val)
			return
		}
		findKinSubtree(node.Left, disc + 1)
		findKinSubtree(node.Right, disc + 1)
	}
	dfs(root)
	return
}
// 对所有节点添加一个指向父节点的引用
func DistanceK4(root, target *BiTreeNode, K int) (answer []interface{}) {
	parents := addParent(root)
	queue := []*BiTreeNode{target}
	cnt := 0
	visited := make(map[*BiTreeNode]interface{})
	for len(queue) != 0 && cnt < K{
		tmp := []*BiTreeNode{}
		seen := make(map[*BiTreeNode]interface{})
		for _, value := range queue{
			seen[value] = nil
			if _, ok := visited[value.Left]; !ok && value.Left != nil {
				tmp = append(tmp, value.Left)
			}
			if _, ok := visited[value.Right]; !ok && value.Right != nil {
				tmp = append(tmp, value.Right)
			}
			if _, ok := visited[parents[value]]; !ok && parents[value] != nil {
				tmp = append(tmp, parents[value])
			}
		}
		queue = tmp
		visited = seen
		cnt++
	}
	for _, value := range queue{
		answer = append(answer, value.Val)
	}
	return
}

/*
  类似于斐波那契数列的求解：
   令 FBT(N) 作为所有含 N 个结点的可能的满二叉树的列表。
   每个满二叉树 T 含有 3 个或更多结点，在其根结点处有 2 个子结点。这些子结点 left 和 right 本身就是满二叉树。
   因此，对于N ≥ 3，我们可以设定如下的递归策略：FBT(N)=[对于所有的 xx，所有的树的左子结点来自FBT(x) 而右子结点来自FBT(N−1−x)]。
优化：
  	1. 通过简单的计数参数，没有满二叉树具有正偶数个结点。
	2. 我们应该缓存函数FBT之前的结果，这样我们就不必在递归中重新计算它们
 */
func AllPossibleFBT(n int) []*BiTreeNode {
	cache := make(map[int][]*BiTreeNode) // 优化：缓存中间结果
	var dfs func(N int)[]*BiTreeNode
	dfs = func(N int)[]*BiTreeNode{
		if N == 1{
			return []*BiTreeNode{&BiTreeNode{}}
		}
		// 优化：N 必须是奇数 否则不能构成FST
		if N % 2 == 0{
			return nil
		}
		// 优化：缓存中间结果
		if _, ok := cache[N]; ok {
			return cache[N]
		}
		var rootSet []*BiTreeNode
		for i := 0; i < N; i++{
			j := N - i - 1
			for _, left := range dfs(i){
				for _, right := range dfs(j) {
					root := &BiTreeNode{}
					root.Left = left
					root.Right = right
					rootSet = append(rootSet, root)
				}
			}
		}
		// 优化：缓存中间结果
		cache[N] = rootSet
		return rootSet
	}
	return dfs(n)
}
// 1325. Delete Leaves With a Given Value
func RemoveLeafNodes(root *BiTreeNode, target int) *BiTreeNode {
	if root == nil{
		return nil
	}
	if root.Left == nil && root.Right == nil && root.Val == target{
		return nil
	}
	root.Left = RemoveLeafNodes(root.Left, target)
	root.Right = RemoveLeafNodes(root.Right, target)
	if root.Left == nil && root.Right == nil && root.Val == target{
		return nil
	}
	return root
}
// 889. Construct Binary Tree from Preorder and Postorder Traversal
func ConstructFromPrePostorder(pre []int, post []int) *BiTreeNode {
	if len(pre) == 0{
		return  nil
	}
	//fmt.Println(pre, post)
	root := &BiTreeNode{Val: pre[0]}
	if len(pre) == 1{
		return root
	}
	/* 逻辑混乱
	locPost := FindElemFromSlice(post, pre[1])
	locPre := 0
	if locPost != 0{
		locPre = FindElemFromSlice(pre, post[locPost - 1])
	}
	fmt.Println(locPost, locPre, post, pre)
	root.Left = ConstructFromPrePost(pre[1:locPre+1], post[:locPost+1])
	root.Right = ConstructFromPrePost(pre[locPre+1:], post[locPost+1:len(post)-1])
	 */
	/*
	   我们令左分支有 L 个节点。我们知道左分支的头节点为 pre[1]，但它也出现在左分支的后序表示的最后。
	   所以 pre[1] = post[L-1]（因为结点的值具有唯一性），因此 L = post.indexOf(pre[1]) + 1。
	   现在在我们的递归步骤中，左分支由 pre[1 : L+1] 和 post[0 : L] 重新分支，而右分支将由 pre[L+1 : N] 和 post[L : N-1] 重新分支。
	 */
	L := FindElemFromSlice(post, pre[1]) + 1 // 左分支节点个数
	fmt.Println(pre, post, L)
	root.Left = ConstructFromPrePostorder(pre[1:L+1], post[:L])
	root.Right = ConstructFromPrePostorder(pre[L+1:], post[L:len(post) - 1])
	return root
}
func FindElemFromSlice(a []int, target int)int{
	for idx, value := range a{
		if target == value{
			return idx
		}
	}
	return -1
}

// 105. Construct Binary Tree from Preorder and Inorder Traversal
func ContructTreeFromPreInorder(preorder []int, inorder []int) *BiTreeNode {
	if len(preorder) <= 0{
		return nil
	}
	root := BiTreeNode{Val: preorder[0]}
	pos := FindElemFromSlice(inorder, root.Val.(int))
	root.Left = ContructTreeFromPreInorder(preorder[1:pos+1], inorder[:pos])
	root.Right = ContructTreeFromPreInorder(preorder[pos+1:], inorder[pos+1:])
	return &root
}

// 105. 迭代版
func ContructTreeFromPreInorder_iter(preorder []int, inorder []int) *BiTreeNode {
	if len(preorder) <= 0{
		return nil
	}
	root := &BiTreeNode{Val: preorder[0]}
	st := []*BiTreeNode{root}
	index := 0
	for _, value := range preorder[1:]{
		if st[0].Val == inorder[index]{ // 栈顶节点没有左孩子，value必须为栈中某个节点的右孩子,如何找到此节点
			// 栈中的节点的顺序和它们在前序遍历中出现的顺序是一致的，而且每一个节点的右儿子都还没有被遍历过，
			//那么这些节点的顺序和它们在中序遍历中出现的顺序一定是相反的
			node := st[0]
			for len(st) > 0 && st[0].Val == inorder[index]{
				node = st[0]
				st = st[1:]
				index++
			}
			node.Right = &BiTreeNode{Val: value}
			st = append([]*BiTreeNode{node.Right}, st...)
		}else{
			st[0].Left = &BiTreeNode{Val: value}
			st = append([]*BiTreeNode{st[0].Left}, st...)
		}
	}
	return root
}

// 106. Construct Binary Tree from Inorder and Postorder Traversal
func ContructTreeFromInPostorder(inorder []int, postorder []int) *BiTreeNode {
	if len(postorder) <= 0 {
		return nil
	}
	value := postorder[len(postorder) - 1]
	root := BiTreeNode{Val: value}
	pos := FindElemFromSlice(inorder, value)
	root.Left = ContructTreeFromInPostorder(inorder[:pos], postorder[:pos])
	root.Right = ContructTreeFromInPostorder(inorder[pos+1:], postorder[pos:len(postorder)-1])
	return &root
}
/*
对于后序遍历中的任意两个连续节点 u 和 v（在后序遍历中，u 在 v 的前面），根据后序遍历的流程，我们可以知道 u 和 v 只有两种可能的关系：
1. u 是 v 的右儿子。这是因为在遍历到 u 之后，下一个遍历的节点就是 v 的双亲节点，即 v；
2. v 没有右儿子，并且 u 是 v 的某个祖先节点（或者 v 本身）的左儿子。
   如果 v 没有右儿子，那么上一个遍历的节点就是 v 的左儿子。
   如果 v 没有左儿子，则从 v 开始向上遍历 v 的祖先节点，直到遇到一个有左儿子（且 v 不在它的左儿子的子树中）的节点va, 那么u 就是 va 的左儿子。
 */
func ContructTreeFromInPostorder_iter(inorder []int, postorder []int) *BiTreeNode {
	return nil
}
// 116. Populating Next Right Pointers in Each Node
func Connect(root *BiTreeNode) *BiTreeNode {
	if root == nil {
		return nil
	}
	var dfs func(node1, node2 *BiTreeNode)
	dfs = func(node1, node2 *BiTreeNode){
		if node1 == nil {
			return
		}
		node1.Next = node2
		dfs(node1.Left, node1.Right)
		dfs(node2.Left, node2.Right)
		dfs(node1.Right, node2.Left)
	}
	dfs(root.Left, root.Right)
	return root
}
/* 使用已建立的next指针
	1. 从根节点开始，第0层只有一个节点，故不需要连接，直接为第1层节点建立next指针即可。
    该算法中需要注意的一点是，当为第N层节点建立next指针时，处于第N-1层。当第 N 层节点的next指针全部建立完成后，转移至第N层，建立N+1层节点的next指针
    2. 遍历某一层的节点时，这层节点的next指针已经建立。因此只需知道这一层的最左节点 就可以按照链表方式遍历，不需要队列
    伪代码：
    leftmost = root
    while (leftmost.left != null){
        head = leftmost
        while head != null {
           1> left right child connection
		   2> using next pointer
           head = head.next
        }
        leftmost = leftmost.left
   }
 */
func Connect2(root *BiTreeNode) *BiTreeNode {
	if root == nil {
		return nil
	}
	for leftmost := root; leftmost.Left != nil; leftmost = leftmost.Left {
		for head := leftmost; head != nil;head = head.Next {
			head.Left.Next = head.Right
			if head.Next != nil {
				head.Right.Next = head.Next.Left
			}
		}
	}
	return root
}

/* 1448. Count Good Nodes in Binary Tree
	Given a binary tree root,
	a node X in the tree is named good if in the path from root to X there are no nodes with a value greater than X.
	Return the number of good nodes in the binary tree.
Example 1:
	Input: root = [3,1,4,3,null,1,5]
	Output: 4
	Explanation: Nodes in blue are good.
	Root Node (3) is always a good node.
	Node 4 -> (3,4) is the maximum value in the path starting from the root.
	Node 5 -> (3,4,5) is the maximum value in the path
	Node 3 -> (3,1,3) is the maximum value in the path.
Example 2:
	Input: root = [3,3,null,4,2]
	Output: 3
	Explanation: Node 2 -> (3, 3, 2) is not good, because "3" is higher than it.
Example 3:
	Input: root = [1]
	Output: 1
	Explanation: Root is considered as good.
*/
func GoodNodes(root *BiTreeNode) int {
	ans := 0
	fmt.Println("---")
	var dfs func(*BiTreeNode, int)
	dfs = func(node *BiTreeNode, curMax int){
		if node == nil{
			return
		}
		if node.Val.(int) >= curMax{
			ans++
			fmt.Println(node.Val)
			curMax = node.Val.(int)
		}
		dfs(node.Left, curMax)
		dfs(node.Right, curMax)
	}
	dfs(root, math.MinInt32)
	return ans
}
/* 1372. Longest ZigZag Path(交错路径) in a Binary Tree
	You are given the root of a binary tree.
	A ZigZag path for a binary tree is defined as follow:
	Choose any node in the binary tree and a direction (right or left).
	If the current direction is right, move to the right child of the current node; otherwise, move to the left child.
	Change the direction from right to left or from left to right.
	Repeat the second and third steps until you can't move in the tree.
	Zigzag length is defined as the number of nodes visited - 1. (A single node has a length of 0).
	Return the longest ZigZag path contained in that tree.
Example 1:
	Input: root = [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1,null,1]
	Output: 3
	Explanation: Longest ZigZag path in blue nodes (right -> left -> right).
Example 2:
	Input: root = [1,1,1,null,1,null,null,1,1,null,1]
	Output: 4
	Explanation: Longest ZigZag path in blue nodes (left -> right -> left -> right).
Example 3:
	Input: root = [1]
	Output: 0
 */
func LongestZigZag(root *BiTreeNode) int {
	maxZ := 0
	var dfs func(node *BiTreeNode)[2]int
	dfs = func(node *BiTreeNode)(ret [2]int){
		 if node.Left == nil{
		 	ret[0] = 0
		 }else{
		 	left := dfs(node.Left)
		 	ret[0] = left[1] + 1
		 }
		 if node.Right == nil {
		 	ret[1] = 0
		 }else{
		 	right := dfs(node.Right)
		 	ret[1] = right[0] + 1
		 }
		 maxZ = max(maxZ, ret[0], ret[1])
		 return ret
	}
	r := dfs(root)
	maxZ = max(maxZ, r[0], r[1])
	return maxZ
}
// 代码优化
func LongestZigZag2(root *BiTreeNode) int {
	maxZ := 0
	var dfs func(node *BiTreeNode)[2]int
	dfs = func(node *BiTreeNode)(ret [2]int){
		ret[0], ret[1] = 0, 0
		if node.Left != nil{
			left := dfs(node.Left)
			ret[0] = left[1] + 1
		}
		if node.Right != nil {
			right := dfs(node.Right)
			ret[1] = right[0] + 1
		}
		maxZ = max(maxZ, ret[0], ret[1])
		return ret
	}
	r := dfs(root)
	maxZ = max(maxZ, r[0], r[1])
	return maxZ
}
/* DP
  f(u): 从根节点u的路径上以u结尾并且u是它父亲的左儿子的最长交错路径
  g(u): 从根到节点u的路径上以u结尾且u是它父亲的右儿子的最长交错路径
  father(u): u的父节点
  方程：
	f(u) = g[fater[u]) + 1  u是左儿子
	g(u) = f[fater[u]) + 1  u是右儿子
  实现：
	维护两个数组： f 和 g, 利用从树的节点到整数的映射容器来实现，方便从节点的指针映射到f和g的函数值。
    可以通过DFS 和 BFS 来遍历树。
    BFS版本：(node,parent) 作为状态，其中node表示当前待计算f和g的节点，parent 表示其父节点。
    初始化：每个点的f和g 都为0
 */
func LongestZigZagDP(root *BiTreeNode) int {
	f, g := map[*BiTreeNode]int{root: 0}, map[*BiTreeNode]int{root: 0}
	q := [][2]*BiTreeNode{[2]*BiTreeNode{root, nil}}
	for len(q) > 0{
		top := q[0]
		q = q[1:]
		u, x := top[0], top[1]
		f[u], g[u] = 0, 0
		if x != nil {
			if x.Left == u  { f[u] = g[x] + 1 }
			if x.Right == u { g[u] = f[x] + 1 }
		}
		if u.Left != nil { q = append(q, [2]*BiTreeNode{u.Left, u}) }
		if u.Right != nil { q = append(q, [2]*BiTreeNode{u.Right, u})}
	}
	maxAns := 0
	for k := range f{
		maxAns = max(maxAns, f[k])
	}
	for k := range g{
		maxAns = max(maxAns, g[k])
	}
	return maxAns
}

/* 662. Maximum Width of Binary Tree
Given the root of a binary tree, return the maximum width of the given tree.
The maximum width of a tree is the maximum width among all levels.
The width of one level is defined as the length between the end-nodes (the leftmost and rightmost non-null nodes),
where the null nodes between the end-nodes are also counted into the length calculation.
It is guaranteed that the answer will in the range of 32-bit signed integer.
 */
/*
这个问题中的主要想法是给每个节点一个 position 值，
如果我们走向左子树，那么 position -> position * 2，
如果我们走向右子树，那么 position -> positon * 2 + 1。
当我们在看同一层深度的位置值 L 和 R 的时候，宽度就是 R - L + 1
 */
func WidthOfBinaryTree(root *BiTreeNode) int {
	type tuple struct {
		node	*BiTreeNode
		depth	int
		pos 	int
	}
	curDepth, left, ans := 0, 0, 0
	q := []tuple{{root, 0, 0}}
	for _, v := range q{
		if v.node != nil{
			q = append(q, tuple{v.node.Left, v.depth+1, v.pos*2})
			q = append(q, tuple{v.node.Right, v.depth+1, v.pos*2+1})
			if curDepth != v.depth{
				curDepth = v.depth
				left = v.pos
			}
			if ans < (v.pos - left + 1){
				ans = v.pos - left + 1
			}
		}
	}
	return ans
}

/* 366. Find Leaves of Binary Tree
** Given the root of a binary tree, collect a tree's nodes as if you were doing this:
	Collect all the leaf nodes.
	Remove all the leaf nodes.
	Repeat until the tree is empty.
 */
// 2022-04-07 刷出此题 转化为图，借助拓扑排序解决
type degree struct{
	d      int
	parent *BiTreeNode
	value  int
}
func findLeaves_graph(root *BiTreeNode) [][]int {
	const Node, Parent = 0, 1
	ans := [][]int{}
	if root == nil { return  ans }
	out := map[*BiTreeNode]degree{}
	st := [][2]*BiTreeNode{ [2]*BiTreeNode{root, nil} }
	queue := []degree{}
	for len(st) > 0{
		top := st[len(st)-1]
		cur := top[Node]
		par := top[Parent]
		st = st[:len(st)-1]
		d := 0
		if cur.Right != nil {
			st = append(st, [2]*BiTreeNode{cur.Right, cur})
			d++
		}
		if cur.Left != nil{
			st = append(st, [2]*BiTreeNode{cur.Left, cur})
			d++
		}
		out[cur] = degree{d, par, cur.Val.(int)}
		if d == 0{  queue = append(queue, out[cur]) }
	}
	for len(queue) > 0{
		q := queue
		queue = nil
		t := []int{}
		for _, deg := range q{
			t = append(t, deg.value)
			par := out[deg.parent]
			if deg.parent != nil{
				par.d--
				out[deg.parent] = par
				if out[deg.parent].d == 0{
					queue = append(queue, out[deg.parent])
				}
			}
		}
		ans = append(ans, t)
	}
	return ans
}

// 借助后序遍历，深度不再从根结点开始算，而是从叶子结点开始算
// 即，叶子结点的深度为0，然后往上累加，叶子的父结点深度为1
// 当某个结点的左右子树深度不相等，取更大值作为其深度
func findLeaves(root *BiTreeNode) [][]int {
	ans := [][]int{}
	var dfs func(node *BiTreeNode)int
	dfs = func(node *BiTreeNode)int{
		if node == nil { return -1 }
		left := dfs(node.Left)
		right := dfs(node.Right)
		height := max(left, right) + 1
		if height >= len(ans){
			ans = append(ans, []int{node.Val.(int)})
		}else{
			ans[height] = append(ans[height], node.Val.(int))
		}
		return height
	}
	dfs(root)
	return ans
}

/* 958. Check Completeness of a Binary Tree
** 判断是否为完全二叉树
 */
// 2022-04-24 刷出此题，
// 在这里 收录是 注意  利用统计 节点数量 是有局限性的
func isCompleteTree(root *BiTreeNode) bool {
	if root == nil { return false }
	queue := []*BiTreeNode{root}
	end := false
	//h := 0
	for len(queue) > 0{
		// n := len(queue)
		q := queue
		queue = nil
		for i := range q{
			if q[i].Left == nil{
				end = true
			}else{
				queue = append(queue, q[i].Left)
			}
			if end && (q[i].Left != nil || q[i].Right != nil ){
				return false
			}
			if q[i].Right == nil {
				end = true
			}else{
				queue = append(queue, q[i].Right)
			}
		}
		/*
		if n != (1<<h) && len(queue) != 0{
			return false
		}
		h++
		*/
	}
	return true
}
/* 方法二： 思路
** 这个问题可以简化成两个小问题：
	用 （depth, position) 表示每个节点的位置；
	确定如何定义所有节点都在最左边
** 一个更简单的表示深度和位置的方法是：
	用 1 表示根节点，对于任意一个节点 v，	它的左孩子为 2*v 右孩子为 2*v + 1。
	这就是我们用的规则，在这个规则下，一颗二叉树是完全二叉树当且仅当节点编号依次为 1, 2, 3, ... 且没有间隙。
** 算法：
	1. 对于根节点，我们定义其编号为 1 对于每个节点，左节点编号为 2*v， 右节点编号： 2*v+1
	2. 树中所有节点的编号按照广度优先搜索顺序正好是升序,也可以使用深度优先搜索，之后对序列排序
	3. 然后，我们检测编号序列是否为无间隔的1 2 3 4，。。。 事实上，我们只需要检查最后一个编号是否正确，因为最后一个编号的值最大
*/
// 主体思路就是 完全依据完全二叉树 与 数组的关系 如果一一对应，则就是完全二叉树，否则不是
func isCompleteTree2(root *BiTreeNode) bool {
	var dfs func(node, parent *BiTreeNode)
	dfs = func(node, parent *BiTreeNode){
		if node == nil {return }
		if parent.Left == node{
			node.Val = parent.Val.(int) * 2
		}
		if parent.Right == node{
			node.Val = parent.Val.(int) * 2 + 1
		}
		dfs(node.Left, node)
		dfs(node.Right, node)
	}
	root.Val = 1
	dfs(root.Left, root)
	dfs(root.Right, root)

	total := 1
	q := []*BiTreeNode{root}
	for len(q) > 0{
		queue := q
		q = nil
		for i := range queue{
			if queue[i].Val != total{
				return false
			}
			total++
			if queue[i].Left != nil{
				q = append(q, queue[i].Left)
			}
			if queue[i].Right != nil{
				q = append(q, queue[i].Right)
			}
		}
	}
	return true
}





