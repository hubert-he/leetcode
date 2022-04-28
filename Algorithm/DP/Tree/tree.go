package Tree

import "math"

func min(nums ...int)int{
	m := nums[0]
	for _, c := range nums{
		if m > c { m = c }
	}
	return m
}

/* 968. Binary Tree Cameras
** You are given the root of a binary tree.
** We install cameras on the tree nodes where each camera at a node can monitor its parent, itself, and its immediate children.
** Return the minimum number of cameras needed to monitor all nodes of the tree.
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
// 2022-04-08 第一次见面毫无思路，思维需要训练
/* 本题的难度在于如何从左、右子树的状态，推导出父节点的状态.
** 如果某棵树的所有节点都被监控，则称该树被「覆盖」
** 假设当前节点为 root，其左右孩子为 left, right。如果要覆盖以 root 为根的树，有两种情况：
	1. 在root处安放camera, 则left， right 一定会被覆盖。此时只需要保证left 的两颗子树被覆盖，以及right的两颗子树也被覆盖即可
	2. 不在root处安放camera，则除了覆盖root的两颗子树外，孩子left，right 之一必须要安装camera，从而保证root被覆盖
** 根据上面的讨论，能够分析出，对于每个节点 root, 需要维护三种类型的状态：
	状态 a：root 必须放置摄像头的情况下，覆盖整棵树需要的摄像头数目。
	状态 b：覆盖整棵树需要的摄像头数目，无论 root 是否放置摄像头
	状态 c：覆盖两棵子树需要的摄像头数目，无论节点 root 本身是否被监控到。
** 根据上面的定义，一定有 a >= b >= c
** 对于节点root而言，设其左右孩子left,right 对应的状态变量分别为(la, lb, lc)以及(ra,rb,rc)。根据上面的讨论，可以推出求解a和b 的过程：
	a = lc + rc + 1
	b = min(a, min(la+rb, ra+lb))
** 对于 c 而言，要保证两颗子树被完全覆盖，
	要么root放一个camera，需要的camera数目为 a
	要么root处不放camera， 此时两颗子树分别保证自己被覆盖，需要的camera 为 lb + rb
	c = min(a, lb + rb)
** 特殊情况处理：left right 孩子为nil的情况
	1. 对于root，如果其某个孩子为空，则不能通过在该孩子处放置摄像头的方式，监控到当前节点。
		因此，该孩子对应的变量 a 应当返回一个大整数，用于标识不可能的情形
 */
func minCameraCover(root *TreeNode) int {
	var dfs func(node *TreeNode)(a, b, c int)
	dfs = func(node *TreeNode)(a, b, c int){
		if node == nil { return math.MaxInt32, 0, 0 }
		la, lb, lc := dfs(node.Left)
		ra, rb, rc := dfs(node.Right)
		a = lc + rc + 1
		b = min(a, min(la+rb, ra+lb))
		c = min(a, lb+rb)
		return
	}
	_, b, _ := dfs(root)
	return b
}

/* 另外的树形DP 题解
** 与打家劫舍III 类似，对于每个节点，要么打劫，要么不打劫，描述一个子树的最大收益需要两个变量：它的根节点，以及是否打劫根节点
** 但是此题多了一种情况，即 除了是否安放相机，还有一个是 是否被监控到，并且被监控的节点也可以安放相机
** 因此，需要3个变量去描述节点的状态：节点本身（代表不同子树）、是否放相机、是否被监控
** 令 minCam 为 以当前节点为根节点的子树，需要放置的最少相机个数
** minCam[root] <= minCam[left] + minCam[right] + 1
** minCam(当前节点，是否安放相机， 是否被父节点和自己监控)
	// 因为递归是父亲调儿子,对于当前节点，它只知道父亲和自己是否监控自己，不知道儿子是否监控自己
** 对于一个子树的根节点，它的状态无非下面3种：
	1. 安放有相机
		1.1 左右子节点没有相机，被父节点监控
			<== 1 + minCam(node.Left, false, true) + minCam(node.Right, false, true)
		1.2 左孩子节点有相机，右孩子无相机，被父节点监控
			<== 1 + minCam(node.Left, true, true) + minCam(node.Right, false, true)
		1.3 右孩子节点有相机，左孩子无相机，被父节点监控
			<== 1 + minCam(node.Left, false, true) + minCam(node.Right, true, true)
		1.4 左右节点都有相机 （可减掉）
	2. 不需要安放相机，被 其父节点监控到
		2.1 左右子节点均有相机
			<= minCam(node.Left, true, true) + minCam(node.Right, true, true)
		2.2 左孩子有相机，右孩子没有相机，但是被它的子节点监控到
			<= minCam(node.Left, true, true) + minCam(node.Right, false, false)
			<= minCam(node.Left, true, true) +
			min{ minCam(node.Right.Left, true, false) + minCam(node.Right.Right, false, false),
				 minCam(node.Right.Left, false, false) + minCam(node.Right.Right, true, false) }
		2.3 右孩子有相机，左孩子没有相机，但是被左孩子的子节点监控到
			<= minCam(node.Left, false, false) + minCam(node.Right, true, true)
		2.4 左孩子与右孩子都没有相机，但都被他们的子节点监控到
			<= minCam(node.Left, false, false) + minCam(node.Right, false, false)
			<= 即求第 3 种情况
	3. 不需要安放相机，被 其子节点监控到
		3.1 左子节点有相机，右子节点没有相机，但被它的孩子监控到
			<= minCam(node.Left, true, true) + minCam(node.Right, false, false)
		3.2 右子节点有相机，左子节点没有相机，但被它的孩子监控到
			<= minCam(node.Left, false, false) + minCam(node.Right, true, true)
		3.3 左右子节点都有相机
			<= minCam(node.Left, true, true) + minCam(node.Right, true, true)
** 以上3种情况分别求min 然后汇总求min
 */
func minCameraCover_DFS(root *TreeNode) int {
	var dfs func(node *TreeNode, camera, watched bool) int
	dfs = func(node *TreeNode, camera, watched bool) int{
		//if node == nil { return 0 }
		if node == nil {
			if camera{ return math.MaxInt32 } // 不会发送节点nil 但是有camera
			return 0
		}
		if camera{ // 安放有相机
			a := dfs(node.Left, false, true) + dfs(node.Right, false, true)
			b := dfs(node.Left, true, true) + dfs(node.Right, false, true)
			c := dfs(node.Left, false, true) + dfs(node.Right, true, true)
			return 1 + min(a, b, c)
		}else{ // node 没有安放相机
			if watched{ // 父节点监控到了
				a := dfs(node.Left, true, true) + dfs(node.Right, true, true)
				b := dfs(node.Left, true, true) + dfs(node.Right, false, false)
				c := dfs(node.Left, false, false) + dfs(node.Right, true, true)
				d := dfs(node.Left, false, false) + dfs(node.Right, false, false)
				return min(a, b, c, d)
			}else{
				a := dfs(node.Left, true, true) + dfs(node.Right, false, false)
				b := dfs(node.Left, false, false) + dfs(node.Right, true, true)
				c := dfs(node.Left, true, true) + dfs(node.Right, true, true)
				return min(a, b, c)
			}
		}
	}
	a := dfs(root, true, true)
	// b := dfs(root, false, true) root 无父节点，此情况不存在
	c := dfs(root, false, false)
	return min(a, c)
}
// 上面的方法 TLE 原因是重复计算太多，借助dp cache 缓存结果
// 这里针对每个节点，计算出三种情况的可能值，返回给上一层，省了 dp 缓存的开销
/* 	树形 DP 不像常规 DP 那样在迭代中 “填表”，而是在递归遍历中 “填表”。
	没有开辟 DP 数组去存中间状态，而是通过子调用将状态返回出去，提供给父调用。
	动态规划是根据过去的状态求出当前状态，按顺序一个个求。这里是沿着一棵树去求解子问题，可以理解为在一棵树上填表。
	可以理解为，递归调用栈把中间计算结果暂存了，子调用的结果交给父调用，它就销毁了，并没有记忆化。
	当然，你也可以开辟容器，key 是节点（子树），value 是子树的计算结果。
	但没有必要，因为中间计算结果不需要存储，所以这算是降维优化了。
	随着递归的出栈，子调用不断向上返回，子问题（子树）被一个个解决。
**  最后求出大问题：整个树的最小相机数。
 */
func minCameraCover_DP(root *TreeNode) int {
	var dfs func(node *TreeNode) (withCam, noCamWatchByDad, noCamWatchBySon int)
	dfs = func(node *TreeNode) (withCam, noCamWatchByDad, noCamWatchBySon int){
		if node == nil {
			return math.MaxInt32, 0, 0
		}
		left_withCam, left_noCamWatchByDad, left_noCamWatchBySon := dfs(node.Left)
		right_withCam, right_noCamWatchByDad, right_noCamWatchBySon := dfs(node.Right)
		// 下面相当于状态转移方程
		withCam = 1 + min(
			left_noCamWatchByDad + right_noCamWatchByDad,
			left_withCam + right_noCamWatchByDad,
			right_withCam + left_noCamWatchByDad,
			)
		noCamWatchByDad = min(
			left_noCamWatchBySon + right_noCamWatchBySon,
			left_noCamWatchBySon + right_withCam,
			left_withCam + right_withCam,
			left_withCam + right_noCamWatchBySon,
			)
		noCamWatchBySon = min(
			left_withCam + right_withCam,
			left_withCam + right_noCamWatchBySon,
			left_noCamWatchBySon + right_withCam,
			)
		return
	}
	// root 节点没有父节点，因此noCamWatchByDad 不可以参与比较，否则计算结果偏小
	a, b, _ := dfs(root)
	return min(a, b)
	//a, b, c := dfs(root)
	//return min(a, b, c)
}