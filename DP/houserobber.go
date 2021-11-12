package DP

import (
	"../Tree"
	"fmt"
	"math"
)
/* 198. House Robber
  You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed,
  the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected
  and it will automatically contact the police if two adjacent houses were broken into on the same night.
  Given an integer array nums representing the amount of money of each house,
  return the maximum amount of money you can rob tonight without alerting the police.
Example 1:
	Input: nums = [1,2,3,1]
	Output: 4
	Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
	Total amount you can rob = 1 + 3 = 4.
Example 2:
	Input: nums = [2,7,9,3,1]
	Output: 12
	Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
	Total amount you can rob = 2 + 9 + 1 = 12.
 */
/*
	二维状态定义：
	res[i][0]: 表示第i间房子 不偷
    res[i][1]: 表示第间房子 偷
    状态方程：
	res[i][0] = max(res[i-1][0], res[i-1][1])
    res[i][1] = res[i-1][0] + nums[i]
    初始条件：res[0][0] = 0  res[0][1] = nums[0]
 */
func RobI(nums []int)int{
	s := len(nums)
	if s <= 0{
		return 0
	}
	// res := make([][2]int, s)
	res := make([][2]int, 2) // 滚动数组
	res[0][0], res[0][1] = 0, nums[0]
	for i := 1; i < s; i++{
		res[i%2][0] = max(res[(i-1)%2][0], res[(i-1)%2][1])
		res[i%2][1] = res[(i-1)%2][0] + nums[i]
	}
	return max(res[(s-1)%2][0], res[(s-1)%2][1])
}
// 2021-10-15 优化一下
func RobI1015(nums []int) int {
	// 二维思路：
	// dp[n][0] = max(dp[n-1][0], dp[n-1][1])
	// dp[n][1] = dp[n-1][0] + nums[n]
	// dp[n] = max(dp[n][0], dp[n][1])
	// 一维思路参考官方题解
	// dp[x] = max(dp[x-2]+nums[x], dp[x-1])
	dp := make([]int, 2)
	n := len(nums)
	for i := 1; i <= n; i++{
		t := dp[0]
		dp[0] = max(dp[0], dp[1])
		dp[1] = t+nums[i-1]
	}
	return max(dp...)
}

/*
  转换为一维DP
  1. 考虑最简单的情况：
	如果只有一间房屋，则偷窃该房屋，可以偷窃到最高总金额。如果只有两间房屋，则由于两间房屋相邻，不能同时偷窃，只能偷窃其中的一间房屋，
	因此选择其中金额较高的房屋进行偷窃，可以偷窃到最高总金额
  2. 如果房屋数量大于两间，设第k间房间，有两个情况：
     a. 偷第k间房，那一定不能偷k-1,偷窃金额最大为：前k-2间房的最高总金额 + 第k间房的金额
     b. 不偷第k间房，偷窃最大金额为 偷k-1间房所得的最大金额
  设定一维状态：dp[i]表示前i间房能偷到的最大金额
  状态转移方程：dp[i] = max(dp[i-2] + nums[i], dp[i-1])
  初始条件：dp[0] = nums[0]  dp[1] = max(nums[0], nums[1])
  最终答案：dp[n-1], n为数组长度
 */
func RobIDP(nums []int) int{
	s := len(nums)
	if s <= 0{
		return 0
	}
	if s == 1{
		return nums[0]
	}
	dp := [2]int{nums[0], max(nums[0], nums[1])}
	for i := 2; i < s; i++{
		dp[i%2] = max(dp[(i-2)%2]+nums[i], dp[(i-1)%2])
	}
	return dp[(s-1)%2]
}

/* 213. House Robber II
  You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed.
  All houses at this place are arranged in a circle. That means the first house is the neighbor of the last one.
  Meanwhile, adjacent houses have a security system connected,
  and it will automatically contact the police if two adjacent houses were broken into on the same night.
  Given an integer array nums representing the amount of money of each house,
  return the maximum amount of money you can rob tonight without alerting the police.
Example 1:
	Input: nums = [2,3,2]
	Output: 3
	Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2), because they are adjacent houses.
Example 2:
	Input: nums = [1,2,3,1]
	Output: 4
	Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
	Total amount you can rob = 1 + 3 = 4.
Example 3:
	Input: nums = [0]
	Output: 0
 */

func RobII(nums []int) int {
	s := len(nums)
	if s <= 0{
		return 0
	}
	if s == 1{
		return nums[0]
	}
	dp0 := [2]int{0, nums[1]}
	dp1 := [2]int{nums[0], nums[0]}
	for i := 2; i < s; i++{
		if i == s-1{
			dp1[i%2] = dp1[(i-1)%2]
		}else{
			dp1[i%2] = max(nums[i]+dp1[(i-2)%2], dp1[(i-1)%2])
		}
		dp0[i%2] = max(nums[i]+dp0[(i-2)%2], dp0[(i-1)%2])
	}
	return max(dp0[(s-1)%2], dp1[(s-1)%2])
}

/*2021-10-15更新： 可以预先计算出 选第一 和 不选第一个导致的序列情况， 最后比较可能情况*/
func RobII1015(nums []int)int{
	n := len(nums)
	if n < 2{
		return nums[0]
	}
	dp0 := []int{0, 0} // 不选第一个导致的序列
	dp1 := []int{0, nums[0]} // 选第一个导致的序列
	for i := 2; i <= n; i++{
		t := dp0[0]
		dp0[0] = max(dp0[0], dp0[1])
		dp0[1] = t + nums[i-1]
		t = dp1[0]
		dp1[0] = max(dp1[0], dp1[1])
		dp1[1] = t + nums[i-1]
	}
	return max(dp0[0], dp0[1], dp1[0])
}

// 2021-11-12 重刷此题
// 存在环，尝试 升维度 和 细分情况 然后汇总结果
func RobII012(nums []int) int {
	n := len(nums)
	//细分情况
	dp0, dp1 := [2]int{0, 0}, [2]int{0, nums[0]}
	if n < 2{ // 仅有一个元素的特殊情况，容易遗漏
		return dp1[1]
	}
	// 照搬dp 方程
	for i := 1; i < n; i++{
		dp0[0], dp0[1] = max(dp0[0], dp0[1]), dp0[0] + nums[i]
		dp1[0], dp1[1] = max(dp1[0], dp1[1]), dp1[0] + nums[i]
	}
	return max(dp0[0], dp0[1], dp1[0])
}

/* 官方题解
  注意到与上一题不同的地方，即当房屋数量不超过两间时，最多只能偷窃一间房屋，因此不需要考虑首尾相连的问题。
  如果房屋数量大于两间，就必须考虑首尾相连的问题，第一间房屋和最后一间房屋不能同时偷窃。
  如何才能保证第一间房屋和最后一间房屋不同时偷窃呢？
  如果偷窃了第一间房屋，则不能偷窃最后一间房屋，因此偷窃房屋的范围是第一间房屋到最后第二间房屋；
  如果偷窃了最后一间房屋，则不能偷窃第一间房屋，因此偷窃房屋的范围是第二间房屋到最后一间房屋。
  状态定义，以及状态方程和初始条件同上，区别在与结果计算
  分别取（start,end)=(0,n-2) 和 （start，end) = (1, n-1)进行计算，取两个dp[end]中的最大值
 */
func _rob(nums []int) int {
	first, second := nums[0], max(nums[0], nums[1])
	for _, v := range nums[2:] {
		first, second = second, max(first+v, second)
	}
	return second
}
func RobII2(nums []int) int {
	s := len(nums)
	if s <= 0{
		return 0
	}
	if s == 1{
		return nums[0]
	}
	if s == 2{
		return max(nums[0], nums[1])
	}
	return max(_rob(nums[:s-1]), _rob(nums[1:]))
}


/* 337. House Robber III
  The thief has found himself a new place for his thievery again. There is only one entrance to this area, called root.
  Besides the root, each house has one and only one parent house.
  After a tour, the smart thief realized that all houses in this place form a binary tree.
  It will automatically contact the police if two directly-linked houses were broken into on the same night.
  Given the root of the binary tree, return the maximum amount of money the thief can rob without alerting the police.
 */
/*
   方法一 穷举
 */
func Rob(root *Tree.BiTreeNode) (ans int) {
	if root == nil {
		return 0
	}
	// 选择 root
	sum := root.Val.(int)
	if root.Left != nil {
		sum += Rob(root.Left.Left) + Rob(root.Left.Right)
	}
	if root.Right != nil {
		sum += Rob(root.Right.Left) + Rob(root.Right.Right)
	}
	return max(sum, Rob(root.Left) + Rob(root.Right))
}
/*
  memo优化：上一把 递归运算存在很多重复计算
 */
func RobMemo(root *Tree.BiTreeNode)(ans int){
	memo := map[*Tree.BiTreeNode]int{}
	var dfs func(node *Tree.BiTreeNode)(res int)
	dfs = func(node *Tree.BiTreeNode)(res int){
		if node == nil {
			return
		}
		if value, ok := memo[root]; ok {
			return value
		}
		sum := node.Val.(int)
		if root.Left != nil {
			sum += dfs(node.Left.Left) + dfs(node.Left.Right)
		}
		if root.Right != nil {
			sum += dfs(node.Right.Left) + dfs(node.Right.Right)
		}
		value := max(sum, dfs(node.Left) + dfs(node.Right))
		memo[node] = value
		return value
	}
	return dfs(root)
}
/*  经典题目： 树型DP  递归+DP
 DP解法：每个节点可以选择偷或不偷两种状态，并且相连节点不能一起偷
  1. 当前节点选择偷时，那么两个孩子节点不能偷
  2. 当前节点选择不偷时，两个孩子节点只需要拿最多的钱出来就可以（两个孩子节点偷不偷没关系)
  使用一个大小为2的数组来表示 res := [2]int{}, 0 代表不偷， 1代表偷
  任何一个节点能偷到的最大钱的状态可定义为：
   1. 当前节点选择不偷：当前节点能偷到的最多 = 左孩子能偷到的最多钱 + 右孩子能偷到的最多钱
		root[0] = max(rob(root.Left)[0], rob(root.Left)[1]) + max(rob(root.Right)[0], rob(root.Right)[1])
   2. 当前节点选择偷：当前节点能偷到的最多 = 左孩子选择自己不偷时能得到的最多钱 + 右孩子选择不偷时能得到的最多钱 + 当前节点的钱
		root[1] = rob(root.Left)[0] + rob(root.Right)[0] + root.Val
*/
func RobDP(root *Tree.BiTreeNode)(ans int){
	var dfs func(node *Tree.BiTreeNode)[2]int
	dfs = func(node *Tree.BiTreeNode)(res [2]int){
		if node == nil {
			return
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		res[0] = max(left[0], left[1]) + max(right[0], right[1])
		res[1] = left[0] + right[0] + node.Val.(int)
		return
	}
	res := dfs(root)
	return max(res[0], res[1])
}

/* 256. Paint House
  There are a row of n houses, each house can be painted with one of the three colors: red, blue or green.
  The cost of painting each house with a certain color is different.
  You have to paint all the houses such that no two adjacent houses have the same color.
  The cost of painting each house with a certain color is represented by a n x 3 cost matrix.
  For example,
  costs[0][0] is the cost of painting house 0 with color red;
  costs[1][2] is the cost of painting house 1 with color green, and so on...
  Find the minimum cost to paint all houses.
Note:
All costs are positive integers.
Example:
	Input: [[17,2,17],[16,16,5],[14,3,19]]
	Output: 10
	Explanation: Paint house 0 into blue, paint house 1 into green, paint house 2 into blue.
				 Minimum cost: 2 + 5 + 3 = 10.
 */
/*
  dp[i][0] = nums[i][0] + min(dp[i-1][1], dp[i-1][2])
  dp[i][1] = nums[i][1] + min(dp[i-1][0], dp[i-1][2])
  dp[i][2] = nums[i][2] + min(dp[i-1][0], dp[i-1][1])
 */
func MinCostI(cost [][3]int) int {
	s := len(cost)
	if s <= 0{
		return 0
	}
	dp := make([][3]int, 2)
	dp[0][0], dp[0][1], dp[0][2] = cost[0][0], cost[0][1], cost[0][2]
	for i := 1; i < s; i++{
		pre := (i-1)%2
		dp[i%2][0] = cost[i][0] + min(dp[pre][1], dp[pre][2])
		dp[i%2][1] = cost[i][1] + min(dp[pre][0], dp[pre][2])
		dp[i%2][2] = cost[i][2] + min(dp[pre][0], dp[pre][1])
	}
	last := (s-1)%2
	return min(dp[last][0], dp[last][1], dp[last][2])
}
/* 解法2：
   状态定义：dp[i][j]表示刷到第i+1房子时用颜色j的最小花费
   转移方程：dp[i][j] = dp[i][j] + min( dp[i-1][(j+1)%3], dp[i-1][(j+2)%3] )
   %3 的理解： 如果当前的房子要用红色刷，则上一个房子只能用绿色或蓝色来刷，那么要求刷到当前房子，
             且当前房子用红色刷的最小花费就等于当前房子用红色刷的钱加上刷到上一个房子用绿色和刷到上一个房子用蓝色中的较小值
   初始条件：
   结果：
 */
func MinCostI2(cost [][3]int) int {
	s := len(cost)
	if s <= 0{
		return 0
	}
	dp := make([][3]int, 2)
	dp[0][0], dp[0][1], dp[0][2] = cost[0][0], cost[0][1], cost[0][2]
	for i := 1; i < s; i++{
		pre := (i-1)%2
		for j := 0; j < 3; j++{
			dp[i%2][j] = min(dp[pre][(j+1)%3], dp[pre][(j+2)%3]) + cost[i][j]
		}
	}
	last := (s-1)%2
	return min(dp[last][0], dp[last][1], dp[last][2])
}
/* 265. Paint House II
  There are a row of n houses, each house can be painted with one of the k colors.
  The cost of painting each house with a certain color is different.
  You have to paint all the houses such that no two adjacent houses have the same color.
  The cost of painting each house with a certain color is represented by a n x k cost matrix.
  For example,
  costs[0][0] is the cost of painting house 0 with color 0;
  costs[1][2] is the cost of painting house 1 with color 2, and so on...
  Find the minimum cost to paint all houses.
Note:
All costs are positive integers.
Example:
	Input: [[1,5,3],[2,9,4]]
	Output: 5
	Explanation: Paint house 0 into color 0, paint house 1 into color 2. Minimum cost: 1 + 4 = 5;
				 Or paint house 0 into color 2, paint house 1 into color 0. Minimum cost: 3 + 2 = 5.
Follow up:
Could you solve it in O(nk) runtime?
 */
/*  避免循环k次查找最小值
这题的解法的思路还是用DP，
但是在找不同颜色的最小值不是遍历所有不同颜色，而是用 min1 和 min2 来记录之前房子的最小和第二小的花费的颜色，
如果当前房子颜色和 min1 相同，那么用 min2 对应的值计算，反之用 min1 对应的值，这种解法实际上也包含了求次小值的方法，
感觉也是一种很棒的解题思路
 */
func MinCostII(cost [][]int) int {
	s := len(cost)
	if s <= 0{
		return 0
	}
	k := len(cost[0])
	min1, min2 := -1, -1
	dp := make([][]int, 2)
	for i := range dp{
		dp[i] = make([]int, k)
	}
	for i := 0; i < k; i++{
		dp[0][i] = cost[0][i]
	}
	for i := 1; i < s; i++{
		last1, last2 := min1, min2
		min1, min2 = -1, -1
		for j := range cost[i]{
			if j != last1{
				if last1 < 0{
					dp[i%2][j] = cost[i][j]
				}else{
					dp[i%2][j] = cost[i][j] + dp[(i-1)%2][last1]
				}
			}else{
				if last2 < 0{
					dp[i%2][j] = cost[i][j]
				}else{
					dp[i%2][j] = cost[i][j] + dp[(i-1)%2][last2]
				}
			}
			if min1 < 0 || dp[i%2][j] < dp[i%2][min1] {
				min2 = min1
				min1 = j
			}else if min2 < 0 || dp[i%2][j] < dp[i%2][min2]{
				min2 = j
			}
		}
	}
	return dp[(s-1)%2][min1]
}

/*
不需要建立二维 dp 数组，直接用三个变量就可以保存需要的信息即可
 */
func MinCostII2(cost [][]int) int {
	s := len(cost)
	if s <= 0{
		return 0
	}
	min1, min2, idx := 0,0,-1
	for i := 0; i < s; i++{
		m1, m2, idl := math.MaxInt32, math.MaxInt32, -1
		for j := 0; j < len(cost[i]); j++{
			c := 0
			if j == idx{
				c = cost[i][j] + min2
			}else{
				c = cost[i][j] + min1
			}
			// update
			if c < m1{
				m2 = m1
				m1 = c
				idl = j
			}else if c < m2{
				m2 = c
			}
		}
		min1 = m1
		min2 = m2
		idx = idl
	}
	return min1
}
/*  1473. Paint House III
  There is a row of m houses in a small city, each house must be painted with one of the n colors (labeled from 1 to n),
  some houses that have been painted last summer should not be painted again.
  A neighborhood is a maximal group of continuous houses that are painted with the same color.
  For example: houses = [1,2,2,3,3,2,1,1] contains 5 neighborhoods [{1}, {2,2}, {3,3}, {2}, {1,1}].
  Given an array houses, an m x n matrix cost and an integer target where:
  houses[i]: is the color of the house i, and 0 if the house is not painted yet.
  cost[i][j]: is the cost of paint the house i with the color j + 1.
  Return the minimum cost of painting all the remaining houses in such a way that there are exactly target neighborhoods.
  If it is not possible, return -1.
Example 1:
	Input: houses = [0,0,0,0,0], cost = [[1,10],[10,1],[10,1],[1,10],[5,1]], m = 5, n = 2, target = 3
	Output: 9
	Explanation: Paint houses of this way [1,2,2,1,1]
	This array contains target = 3 neighborhoods, [{1}, {2,2}, {1,1}].
	Cost of paint all houses (1 + 1 + 1 + 1 + 5) = 9.
Example 2:
	Input: houses = [0,2,1,2,0], cost = [[1,10],[10,1],[10,1],[1,10],[5,1]], m = 5, n = 2, target = 3
	Output: 11
	Explanation: Some houses are already painted, Paint the houses of this way [2,2,1,2,2]
	This array contains target = 3 neighborhoods, [{2,2}, {1}, {2,2}].
	Cost of paint the first and last house (10 + 1) = 11.
Example 3:
	Input: houses = [0,0,0,0,0], cost = [[1,10],[10,1],[1,10],[10,1],[1,10]], m = 5, n = 2, target = 5
	Output: 5
Example 4:
	Input: houses = [3,1,2,3], cost = [[1,1,1],[1,1,1],[1,1,1],[1,1,1]], m = 4, n = 3, target = 3
	Output: -1
	Explanation: Houses are already painted with a total of 4 neighborhoods [{3},{1},{2},{3}] different of target = 3.
 */
/*  3维DP
	定义状态：
		dp[i][j][k]表示0-i的房子都涂上颜色，末尾的第i个房子的颜色是j，并且它属于第k个街区时的最少花费
    状态转移方程：
    	因为是DP，因此必须考虑i-1的情况，这关系到cost 和 街区数量 的计算，因此需要对其进行 枚举 所有情况
	  设第i-1房子的颜色为p，
	  情况-1：如果第i间房子已经涂色，
		只有满足 j == houses[i] 的状态才是有意义的，此时又分2种情况
			<a> j == p => 那么第 i−1 个房子和第 i 个房子属于同一个街区，状态转移方程
				dp[i][j][k] = dp[i-1][j][k]   j == houses[i] && j== p
			<b> j != p => 属于不同街区
				dp[i][j][k] = dp[i-1][p][k-1]  j == houses[i] && j != p
		综合<a> <b>两种情况，状态转移方程：
				dp[i][j][k] = min(dp[i-1][j][k], dp[i-1][p][k-1])   j == houses[i]
							= math.MaxInt32							j != houses[i]
		其余状态均为math.MaxInt32 即无效状态,方程如下：
			dp[i][j][k] = math.MaxInt32  house[i] != -1 && house[i] != j
	  情况-2：第i间房子未涂色，即 houses[i] = 0， 此时可以刷任意颜色，不存在无效状态, 花费为cost[i][j]
			同情况-1， 根据i 和 i-1 颜色是否相同来决定第i间房子是否形成一个新的街区
			转移方程为：
				dp[i][j][k] = cost[i][j] + min(dp[i-1][j][k], dp[i-1][p][k-1])  p != j
	初始状态：
		dp[0][j][0] = math.MaxInt32  若 houses[0] != -1 && j != houses[0] 无效状态
					= 0				 若 house[0] != -1 && house[0] == j   不用涂色
					= cost[0][j]	 若 house[i] == -1
		当 i == 0 且 k != 0时， dp[0][j][k] = math.MaxInt32
 */
func minCostIII(houses []int, cost [][]int, m int, n int, target int) int {
	// 将颜色调整为从0开始编号，0标记为-1
	for i := range houses{
		houses[i]--
	}
	// dp初始化
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, target)
			for k := range dp[i][j]{
				dp[i][j][k] = math.MaxInt32
			}
		}
	}
	for i := 0; i < m; i++{
		for j := 0; j < n; j++{
			if houses[i] != -1 && houses[i] != j{
				continue
			}
			for k := 0; k < target; k++{
				for p := 0; p < n; p++{
					if j == p{ // 与i-1同色情况
						if i == 0{
							if k == 0{
								dp[i][j][k] = 0
							}
						}else{
							dp[i][j][k] = min(dp[i][j][k], dp[i-1][j][k])
						}
					}else{// 与i-1不同色情况
						if i > 0 && k > 0{
							dp[i][j][k] = min(dp[i][j][k], dp[i-1][p][k-1])
						}
					}
				}
				if dp[i][j][k] != math.MaxInt32 && houses[i] == -1{// 有效状态且未涂色
					dp[i][j][k] += cost[i][j]
				}
			}
		}
	}
	ans := math.MaxInt32
	for _, res := range dp[m-1]{
		ans = min(ans, res[target - 1])
	}
	if ans == math.MaxInt32{
		return -1
	}
	return ans
}
/* 暴力DFS， 画出递归树 能很明显发现 可以循环迭代：递归树节点是房间号，边是颜色，如果没有neighbor限制，可以直接2层for迭代完成
   通过代码，亦可看出这是尾递归，很容易转换为迭代
 */
func MinCostIIIBFS(houses []int, cost [][]int, m int, n int, target int) int {
	ans := math.MaxInt32
	cache := make([][][]int, m+1)
	for i := range cache{
		cache[i] = make([][]int, n+1)
		for j := range cache[i]{
			cache[i][j] = make([]int, target + 1)
		}
	}
	var dfs func(curRoom, lastColor, neigh, sum int)
	dfs = func(curRoom, lastColor, neigh, sum int){
		if sum >= ans || neigh > target{ // 终止1-超过之前结算，或 neighbor 过大
			return
		}
		if curRoom == m{
			if neigh == target{
				ans = min(ans, sum)
			}
			return // 终止2-正常终止
		}
		if curRoom - neigh > int(math.Abs(float64(target - m))) {
			fmt.Println(curRoom, neigh, target, m)
			// 剪枝：差额超过，提前结束
			return
		}
		//neighCnt := neigh
		//fmt.Println("->", curRoom, houses[curRoom], neigh, lastColor)
		if houses[curRoom] == 0{ //未涂色
			for i := 0; i < n; i++{
				color := i+1 // houses里的颜色值从 1 开始计
				neighCnt := neigh
				if lastColor != color{ // lastColor初始为-1
					neighCnt++
				}
				//fmt.Println(curRoom, neighCnt)
				dfs(curRoom+1, color, neighCnt, sum + cost[curRoom][i])
			}
		}else{ // 已涂色
			neighCnt := neigh
			//fmt.Println(houses[curRoom])
			if lastColor != houses[curRoom]{
				neighCnt++
			}
			dfs(curRoom + 1, houses[curRoom], neighCnt, sum)
			//dfs(curRoom + 1, lastColor, neighCnt, sum)
		}
	}
	dfs(0, -1, 0, 0)
	if ans == math.MaxInt32{
		return -1
	}else {
		return ans
	}
}

func MinCostIIIDFSDP(houses []int, cost [][]int, m int, n int, target int) int {
	INF := math.MaxInt32
	vis := make([][][]bool, m+1)
	cache := make([][][]int, m+1)
	for i := range vis{
		vis[i] = make([][]bool, n+1)
		cache[i] = make([][]int, n+1)
		for j := range vis[i]{
			vis[i][j] = make([]bool, target+1)
			cache[i][j] = make([]int, target+1)
		}
	}
	/*	u : 当前处理到的房间编号
		last : 上一次处理的房间颜色
		cnt : 当前形成的分区数量
		sum : 当前的涂色成本
	 */
	var dfs func(int, int, int, int)int
	dfs = func(u, last, cnt, sum int) int{
		if cnt > target{
			return INF
		}
		if vis[u][last][cnt] {
			return cache[u][last][cnt]
		}
		if u == m {
			if cnt == target{
				return 0
			}else{
				return INF
			}
		}
		result := INF
		color := houses[u]
		//fmt.Println(color)
		neighCnt := 1 // u == 0情况 默认
		if color == 0{// 未涂色
			for i := 1; i <= n; i++{
				if u != 0{
					if last == i{
						neighCnt = cnt
					}else{
						neighCnt = cnt+1
					}
				}
				cur := dfs(u+1, i, neighCnt, sum + cost[u][i-1])
				result = min(result, cur + cost[u][i-1])
			}
		}else {
			if u != 0{
				if last == color {
					neighCnt = cnt
				}else{
					neighCnt = cnt + 1
				}
			}
			cur := dfs(u+1, color, neighCnt, sum)
			result = min(result, cur)
		}
		vis[u][last][cnt] = true
		cache[u][last][cnt] = result
		return result
	}
	ans := dfs(0,0,0,0)
	if ans == INF{
		return -1
	}else{
		return ans
	}
}
