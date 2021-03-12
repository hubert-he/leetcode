package Tree

import (
	"fmt"
	"math"
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

func postOrderMorris(root *BiTreeNode) (list []interface{}) {
	dump := &BiTreeNode{nil, root, nil}
	cur := dump
	for cur != nil {
		if cur.Left == nil {
			cur = cur.Right
			continue
		}
		prev := cur.Left
		for prev.Right != nil && prev.Right != cur {
			prev = prev.Right
		}
		if prev.Right == nil {
			prev.Right = cur
			cur = cur.Left
		}else{
			prev.Right = nil
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
func IsSubtree_2dfs(s *BiTreeNode, t *BiTreeNode)( bool) {
	if s == nil{
		return false
	}
	is_subtree := isEqual(s, t)
	if is_subtree{
		return true
	}
	is_subtree = IsSubtree(s.Left,  t)
	if is_subtree{
		return true
	}
	is_subtree = IsSubtree(s.Right, t)
	if is_subtree{
		return true
	}
	return false
}

func isEqual(s *BiTreeNode, t *BiTreeNode) (equal bool){
	if s == nil && t == nil {
		return true
	}
	if (s == nil && t != nil) || (s != nil && t == nil){
		return false
	}
	if s.Val != t.Val{
		return false
	}
	left := isEqual(s.Left, t.Left)
	if left == false{
		return false
	}
	right := isEqual(s.Right, t.Right)
	if right == false{
		return false
	}
	return true
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