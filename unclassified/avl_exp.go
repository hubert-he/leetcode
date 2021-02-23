package unclassified

import (
	"errors"
	"fmt"
	"math/rand"
)

type node struct {
	key    int
	value  int
	height int
	left   *node
	right  *node
}

type avlTree struct {
	size int
	Root *node
}

func NewAvlTree() *avlTree {
	return new(avlTree)
}

func (tree *avlTree) GetSize() int {
	return tree.size
}

// 获取高度
func (nd *node) getHeight() int {
	if nd == nil {
		return 0
	}
	return nd.height
}

// 获取的平衡因子
func (nd *node) getBalanceFactor() int {
	if nd == nil {
		return 0
	}
	return nd.left.getHeight() - nd.right.getHeight()
}

func max(int1, int2 int) int {
	if int1 >= int2 {
		return int1
	}
	return int2
}

// 添加/更新节点
func (tree *avlTree) Add(key, val int) {
	isAdd, nd := tree.Root.add(key, val)
	tree.size += isAdd
	tree.Root = nd
}

// 递归写法:向树的root节点中插入key,val
// 返回1,代表加了节点
// 返回0,代表没有添加新节点,只更新key对应的value值
func (nd *node) add(key, val int) (int, *node) {
	if nd == nil {
		return 1, &node{key, val, 1, nil, nil}
	}

	isAdd := 0
	if key < nd.key {
		isAdd, nd.left = nd.left.add(key, val)
	} else if key > nd.key {
		isAdd, nd.right = nd.right.add(key, val)
	} else { // nd.key == key
		// 对value值更新,节点数量不增加,isAdd = 0
		nd.value = val
	}

	// 更新节点高度和维护平衡
	nd = nd.updateHeightAndBalance(isAdd)
	return isAdd, nd
}

func (tree *avlTree) Remove(key int) error {
	if tree.Root == nil {
		return errors.New(
			"failed to remove,avlTree is empty.")
	}
	var isRemove int
	isRemove, tree.Root = tree.Root.remove(key)
	tree.size -= isRemove
	return nil
}

// 删除nd为根节点中key节点,返回更新了高度和维持平衡的新nd节点
// 返回值int == 1 表明有节点被删除,0 表明没有节点被删除
func (nd *node) remove(key int) (int, *node) {
	// 找不到key对应node,返回0,nil
	if nd == nil {
		return 0, nil
	}

	var retNode *node
	var isRemove int
	switch {
	case key < nd.key:
		isRemove, nd.left = nd.left.remove(key)
		retNode = nd
	case key > nd.key:
		isRemove, nd.right = nd.right.remove(key)
		retNode = nd
	default:
		switch {
		case nd.left == nil: // 待删除节点左子树为空的情况
			retNode = nd.right
		case nd.right == nil: // 待删除节点右子树为空的情况
			retNode = nd.left
		default:
			// 待删除节点左右子树均不为空的情况
			// 找到比待删除节点大的最小节点,即右子树的最小节点
			retNode = nd.right.getMinNode()
			nd.key = retNode.key
			nd.value = retNode.value
			nd.right.remove(retNode.key)
			retNode = nd
		}
		isRemove = 1
	}

	// 前面删除节点后,返回retNode有可能为空,这样在执行下面的语句会产生异常，加这步判定避免异常
	if retNode == nil {
		return isRemove, retNode
	}

	retNode = retNode.updateHeightAndBalance(isRemove)
	return isRemove, retNode
}

// 对节点y进行向右旋转操作，返回旋转后新的根节点x
//        y                              x
//       / \                           /   \
//      x   T4     向右旋转 (y)        z     y
//     / \       - - - - - - - ->    / \   / \
//    z   T3                       T1  T2 T3 T4
//   / \
// T1   T2
func (y *node) rightRotate() *node {
	x := y.left
	y.left = x.right
	x.right = y

	// 更新height值: 旋转后只有x,y的子树发生变化,所以只更新x,y的height值
	// x的height值和y的height值相关,先更新y,再更新x
	y.height = 1 + max(y.left.getHeight(), y.right.getHeight())
	x.height = 1 + max(x.left.getHeight(), x.right.getHeight())
	return x
}

// 对节点y进行向左旋转操作，返回旋转后新的根节点x
//    y                             x
//  /  \                          /   \
// T1   x      向左旋转 (y)       y     z
//     / \   - - - - - - - ->   / \   / \
//   T2  z                     T1 T2 T3 T4
//      / \
//     T3 T4
func (y *node) leftRotate() *node {
	x := y.right
	y.right = x.left
	x.left = y

	// 更新height值
	y.height = 1 + max(y.left.getHeight(), y.right.getHeight())
	x.height = 1 + max(x.left.getHeight(), x.right.getHeight())
	return x
}

func (tree *avlTree) Contains(key int) bool {
	return tree.Root.contains(key)
}

// 以root为根的树中是否包含key节点
func (nd *node) contains(key int) bool {
	if nd == nil {
		return false
	}

	if nd.key == key {
		return true
	}

	if key < nd.key {
		return nd.left.contains(key)
	}
	return nd.right.contains(key)
}

// 中序遍历打印出key,val,height
func (tree *avlTree) PrintInOrder() {
	resp := [][]int{}
	tree.Root.printInOrder(&resp)
	fmt.Println(resp)
}

func (nd *node) printInOrder(resp *[][]int) {
	if nd == nil {
		return
	}
	nd.left.printInOrder(resp)
	*resp = append(*resp, []int{nd.key, nd.value, nd.height})
	nd.right.printInOrder(resp)
}

// 中序遍历所有key,用来辅助判断是否是二分搜索树
func (nd *node) traverseInOrderKey(resp *[]int) {
	if nd == nil {
		return
	}

	nd.left.traverseInOrderKey(resp)
	*resp = append(*resp, nd.key)
	nd.right.traverseInOrderKey(resp)
}

// 判断avlTree是否仍然是一颗二分搜索树
// 思路: 二分搜索数如果用中序遍历时,所有元素都是从小到大排列
func (tree *avlTree) IsBST() bool {
	buf := []int{}
	tree.Root.traverseInOrderKey(&buf)
	length := len(buf)
	for i := 1; i < length; i++ {
		if buf[i-1] > buf[i] {
			return false
		}
	}
	return true
}

// 判断avlTree是否是一颗平衡二叉树(每个节点的左右子树高度差不能超过1)
func (tree *avlTree) IsBalanced() bool {
	return tree.Root.isBalanced()
}

func (nd *node) isBalanced() bool {
	if nd == nil {
		return true
	}

	if nd.getBalanceFactor() > 1 || nd.getBalanceFactor() < int(-1) {
		return false
	}
	// 能到这里说明:当前该节点满足平衡二叉树条件
	// 再去判断该节点的左右子树是否也是平衡二叉树
	return nd.left.isBalanced() && nd.right.isBalanced()
}

// 找出以nd为根节点中最小值的节点
func (nd *node) getMinNode() *node {
	if nd.left == nil {
		return nd
	}
	return nd.left.getMinNode()
}

// 更新节点高度和维护平衡
func (nd *node) updateHeightAndBalance(isChange int) *node {
	// 0说明无改变,不必更新
	if isChange == 0 {
		return nd
	}

	// 更新高度
	nd.height = 1 + max(nd.left.getHeight(),
		nd.right.getHeight())

	// 平衡维护
	if nd.getBalanceFactor() > 1 {
		// 左左LL
		if nd.left.getBalanceFactor() >= 0 {
			nd = nd.rightRotate()
		} else { // 左右LR
			nd.left = nd.left.leftRotate()
			nd = nd.rightRotate()
		}
		return nd
	}

	if nd.getBalanceFactor() < int(-1) {
		// 右右RR
		if nd.right.getBalanceFactor() <= 0 {
			nd = nd.leftRotate()
		} else { // 右左RL
			nd.right = nd.right.rightRotate()
			nd = nd.leftRotate()
		}
	}
	return nd
}

func main() {
	a := NewAvlTree()
	nums := []int{}
	a.PrintInOrder()
	fmt.Println("----")
	for i := 0; i < 10; i++ {
		key := rand.Intn(1000)
		a.Add(key, 0)
		nums = append(nums, key)
		if !a.IsBST() || !a.IsBalanced() {
			fmt.Println("a++>代码有错误", a.IsBST(), a.IsBalanced())
			break
		}
	}
	fmt.Println("----")
	a.PrintInOrder()
	fmt.Println(nums[5])
	a.Remove(nums[5])
	a.Remove(847)
	a.PrintInOrder()
	/*
	for i := 0; i < 1000; i++ {
		nums = append(nums, i)
	}
	for _, v := range nums {
		a.Remove(v)
		if !a.IsBST() || !a.IsBalanced() {
			fmt.Println("代码有错误", a.IsBST(), a.IsBalanced())
			break
		}
	}
*/
	fmt.Println("测试结束")
	fmt.Println("is BST:", a.IsBST())
	fmt.Println("is Balanced:", a.IsBalanced())
}