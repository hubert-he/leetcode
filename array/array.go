package array

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

/*
27. 移除元素  双指针
*/
func removeElement(nums []int, val int) int {
	i, j := 0, 0
	for j < len(nums) {
		if nums[j] == val {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
	return j - i
}

func removeElementII(nums []int, val int) int {
	i, j := 0, 0
	for j < len(nums) {
		if nums[j] != val {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
	return i
}
/* 189. Rotate Array
** Given an array, rotate the array to the right by k steps, where k is non-negative.
 */
/*10.14日可以想到的：
** 从位置 0 开始，最初令 tmp = nums[0],根据规则，位置 0 的元素会放至 (0+k)%n 的位置，互相交换位置，然后找下一个位置
** nums[0], nums[(i+k)%n] = nums[(i+k)%n],nums[0]
** 不断进行上述过程，直至回到初始位置 0
** 下面是没有想到的 《== 环状替换
** 但是有的情况，会发生当回到初始位置 0 时，有些数字可能还没有遍历到，此时我们应该从下一个数字开始重复的过程，可是这个时候怎么才算遍历结束呢？
** 我们不妨先考虑这样一个问题：从 0 开始不断遍历，最终回到起点 0 的过程中，我们遍历了多少个元素
** 关键点：
** 由于最终回到了起点，故该过程恰好走了整数数量的圈，不妨设为 a 圈；再设该过程总共遍历了 b 个元素。
** 因此，我们有 an=bk，即 an 一定为 n,k 的公倍数。又因为我们在第一次回到起点时就结束，因此 a 要尽可能小，故 an 就是 n,k 的最小公倍数 lcm(n,k)，
** 因此 b 就为 lcm(n,k)/k
** 这说明单次遍历会访问到 lcm(n,k)/k 个元素。为了访问到所有的元素，我们需要进行遍历的次数为
** n / (lcm(n,k)/k) => nk / lcm(n,k) = gcd(n,k) 即 n k 的最大公约数
** 另： 也可以使用另外一种方式完成代码：使用单独的变量 count 跟踪当前已经访问的元素数量，当 count=n 时，结束遍历过程
 */
func Rotate(nums []int, k int)  {
	n := len(nums)
	gcd := func(a, b int)int{
		for a != 0{
			a, b = b%a, a
		}
		return b
	}
	count := gcd(n, k)
	for i := 0; i < count; i++{
		j := (i+k)%n
		for j != i{
			nums[i], nums[j] = nums[j], nums[i]
			j = (j+k)%n
		}
	}
}
/*使用单独的变量 count 跟踪当前已经访问的元素数量，当 count=n 时，结束遍历过程*/
func RotateCount(nums []int, k int)  { // 这个应该是最优的（题目要求内存复杂度O(1)）
	n := len(nums)
	count := 1
	for i := 0; count < n; i++{
		j := (i+k)%n
		for j != i{
			nums[i], nums[j] = nums[j], nums[i]
			j = (j+k)%n
			count++
		}
		count++ // 务必记得加1
	}
}
/* 方法二： 数组翻转
** 原理： 当我们将数组的元素向右移动 k 次后，尾部 k%n 个元素会移动至数组头部，其余元素向后移动 k%n 个位置
** 该方法为数组的翻转：我们可以先将所有元素翻转，这样尾部的 k%n 个元素就被移至数组头部，
** 然后我们再翻转 [0, k%(n-1)] 区间的元素 以及 [k % n, n-1] 区间的元素即能得到最后的答案
 */
func RotateReverse(nums []int, k int)  {
	reverse := func(a []int){
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1{
			a[i], a[j] = a[j], a[i]
		}
	}
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

/*  组合
	1. 画出树型图
		1.1 分析搜到到重复列表的原因
		1.2 如何去重
			1.2.1 按顺序搜索，设置搜索起点 start
	2. 编码
		2.1 回溯、深度优先遍历、递归
	3. 剪枝
		3.1 剪枝非必须
		3.2 剪枝的前提：候选数组排序
	什么时候使用 used 数组，什么时候使用 start 变量
		排列问题，讲究顺序（即 [2, 2, 3] 与 [2, 3, 2] 视为不同列表时），需要记录哪些数字已经使用过，此时用 used 数组
		组合问题，不讲究顺序（即 [2, 2, 3] 与 [2, 3, 2] 视为相同列表时），需要按照某种顺序搜索，此时使用 start 变量。
*/
// 46. 全排列
func Permute(nums []int) [][]int {
	size := len(nums)
	if size == 0 {
		return nil
	}
	result := [][]int{}
	var dfs func(path, array []int)
	dfs = func(path, array []int) {
		length := len(array)
		if length <= 1 {
			path = append(path, array...)
			result = append(result, append([]int{}, path...))
			return
		}
		for i := 0; i < length; i++ {
			cur := array[0]
			array = array[1:]
			path = append(path, cur)
			dfs(path, array)
			array = array[0 : length-1]
			array = append(array, cur)
			path = path[:len(path)-1]
		}
	}
	dfs(nil, nums)
	return result
}

func PermuteII(nums []int) [][]int {
	result := [][]int{}
	length := len(nums)
	if length == 0 {
		return result
	}
	used := make([]bool, length)
	var dfs func([]int, int)
	dfs = func(path []int, depth int) {
		if depth == length {
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			result = append(result, pathCopy)
			return
		}
		for i := 0; i < length; i++ {
			if !used[i] {
				path = append(path, nums[i]) // 路径添加
				used[i] = true
				dfs(path, depth+1)
				used[i] = false
				path = path[:len(path)-1] // 路径恢复
			}
		}
	}
	dfs([]int{}, 0)
	return result
}

// 47. 全排列II
func PermuteUnique(nums []int) [][]int {
	result := [][]int{}
	length := len(nums)
	if length == 0 {
		return result
	}
	used := make([]bool, length)
	var dfs func([]int, int)
	dfs = func(path []int, depth int) {
		if depth == length {
			result = append(result, append([]int{}, path...))
			return
		}
		dup := map[int]interface{}{}
		for i := 0; i < length; i++ {
			if !used[i] {
				if _, ok := dup[nums[i]]; ok {
					continue
				}
				dup[nums[i]] = nil
				path = append(path, nums[i])
				used[i] = true
				dfs(path, depth+1)
				used[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	dfs([]int{}, 0)
	return result
}

func PermuteUniqueII(nums []int) (ans [][]int) {
	sort.Ints(nums) // 注意 nums变的有序
	length := len(nums)
	perm := []int{}
	vis := make([]bool, length)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == length {
			permCopy := make([]int, length)
			copy(permCopy, perm)
			ans = append(ans, permCopy)
		}
		for i, v := range nums {
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue // 有序后，查看是否与前一个重复
			}
			perm = append(perm, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return
}

// 784. Letter Case Permutation
func LetterCasePermutation(S string) []string {
	Sbyte := []byte(S)
	length := len(S)
	result := []string{}
	if length == 0 {
		return result
	}
	path := []byte{}
	var backtrack func(index int)
	backtrack = func(index int) {
		if index >= length {
			path_copy := make([]byte, len(path))
			copy(path_copy, path)
			result = append(result, string(path_copy))
			return
		}
		ch := Sbyte[index]
		switch {
		case ch >= 'A' && ch <= 'Z':
			path = append(path, ch)
			backtrack(index + 1)
			path[len(path)-1] = ch + 32
			backtrack(index + 1)
			path = path[:len(path)-1]
		case ch >= 'a' && ch <= 'z':
			path = append(path, ch)
			backtrack(index + 1)
			path[len(path)-1] = ch - 32
			backtrack(index + 1)
			path = path[:len(path)-1]
		default:
			path = append(path, ch)
			backtrack(index + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return result
}
func LetterCasePermutationII(S string) []string {
	result, Sbyte := []string{S}, []byte(S) // 1. 添加空集
	length := len(S)
	var backtrack func(int)
	backtrack = func(start int){
		for i := start; i < length; i++{//2. 遍历选择列表
			if Sbyte[i] > '9'{
				Sbyte[i] ^= (1<<5)  //大小写转换通过异或转换大小写
				result = append(result, string(Sbyte))// 记录子集
				// 剪枝条件：决策树子节点 < 父节点
				// 实现剪枝：传递 i+1作为start
				backtrack(i+1)
				Sbyte[i] ^= (1 << 5)// 回溯
			}
		}
	}
	backtrack(0)
	return result
}
/*思路：
如果下一个字符 ch 是字母，将当前已遍历过的字符串全排列复制两份，
在第一份的每个字符串末尾添加 lowercase(c)，在第二份的每个字符串末尾添加 uppercase(c)。
如果下一个字符 ch 是数字，将 ch 直接添加到每个字符串的末尾。
 */
func LetterCasePermutationIII(S string) []string {
	result, Sbyte := []string{}, []byte(S)
	path := make([][]byte, 1)
	//path := [][]byte{nil} 与上作用相同，如果期初未nil占位，之后append元素后nil占位被替换
	for _,ch := range Sbyte{
		n := len(path)
		for i := 0; i < n; i++{
			tmp := make([]byte, len(path[i]))
			// tmp := []byte{} 使用copy 务必保证source 长度
			//copy(tmp, path[i]) 不能在此处复制，会丢数据
			switch {
			case ch >= 'a' && ch <= 'z':
				copy(tmp, path[i])
				path[i] = append(path[i], ch)
				tmp = append(tmp, ch - 32)
				path = append(path, tmp)
			case ch >= 'A' && ch <= 'Z':
				copy(tmp, path[i])
				path[i] = append(path[i], ch)
				tmp = append(tmp, ch + 32)
				path = append(path, tmp)
			default:
				path[i] = append(path[i], ch)
			}
		}
	}
	for _,value := range path{
		result = append(result, string(value))
	}
	return result
}
func LetterCasePermutationDFS(S string) (result []string){
	length := len(S)
	var dfs func(path []byte, start int)
	dfs = func(path []byte, start int){
		if start == length{
			result = append(result, string(path))
			return
		}
		ch := S[start]
		switch {
		case ch >= 'a' && ch <= 'z':
			dfs(append(path, ch), start+1)
			dfs(append(path, ch - 'a' + 'A'), start + 1)
		case ch >= 'A' && ch <= 'Z':
			dfs(append(path, ch), start+1)
			dfs(append(path, ch - 'A' + 'a'), start + 1)
		default:
			dfs(append(path, ch), start+1)
		}
	}
	dfs([]byte{}, 0)
	return
}
/*2021-10-26 重新练习*/
func LetterCasePermutationDFS2(s string) []string {
	var dfs func([]byte)[][]byte
	dfs = func(str []byte)[][]byte{
		if len(str) <= 0{
			return nil
		}
		c := str[0]
		ret,head := [][]byte{}, [][]byte{[]byte{c}}
		if c >= 'a' && c <= 'z'{
			head = append(head, []byte{c-'a'+'A'})
		}
		if c >= 'A' && c <= 'Z'{
			head = append(head, []byte{c-'A'+'a'})
		}
		sub := dfs(str[1:])
		/* sub 有 nil 可能*/
		for i := range sub{
			for j := range head{
				ret = append(ret, append(head[j], sub[i]...))
			}
		}
		if len(sub) == 0{
			for j := range head{
				ret = append(ret, head[j])
			}
		}
		return ret
	}
	ans := []string{}
	t := dfs([]byte(s))
	for i := range t{
		ans = append(ans, string(t[i]))
	}
	return ans
}
// 39. Combination Sum
// 针对每个元素，都有选 或 不选择两个
func CombinationSum(candidates []int, target int) [][]int {
	length := len(candidates)
	result, path := [][]int{}, []int{} // 此时， result的值实际为nil
	var backtrack func(int,int)
	backtrack = func(sum, start int){
		if sum == 0{
			path_copy := make([]int, len(path))
			copy(path_copy, path)
			result = append(result, path_copy)
			return
		}
		if sum < 0 || start >= length{
			return
		}
		// 跳过，不选择candidates[start]
		backtrack(sum, start + 1)
		// 选择candidates[start]
		path = append(path, candidates[start])
		backtrack(sum - candidates[start], start)
		path = path[:len(path) - 1]
	}
	backtrack(target, 0)
	return result
}
func CombinationSumII(candidates []int, target int) (result [][]int){
	sort.Ints(candidates) // 剪枝
	length := len(candidates)
	path := []int{}
	//var backtrack func(int)
	var backtrack func(int, int) // 增加start变量，去重
	//backtrack = func(sum int){
	backtrack = func(sum, start int){
		if sum == target{
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			result = append(result, pathCopy)
			return
		}
		for i := start; i < length; i++{ // 控制决策树中每层的元素选取
			num := candidates[i]
			//fmt.Printf("befor start=> %v, sum => %d, add => %d\n", path, sum, num)
			// 剪枝，序列有序前提下
			if sum + num > target{
				//fmt.Printf("cut all start=> %v, sum => %d\n", path, sum)
				break
			}
			if i > 0 && num == candidates[i-1]{
				//fmt.Printf("cut start=> %v, sum => %d\n", path, sum)
				continue
			}
			path = append(path, num)
			//backtrack(sum + num, start+1)
			backtrack(sum+num, i) // 指定当层只能往后选择
			path = path[:len(path) - 1]
			//fmt.Printf("after start=> %v, sum => %d\n", path, sum)
		}
	}
	backtrack(0, 0)
	return
}

// 40. Combination Sum II
/*
由于我们需要求出所有和为target的组合，并且每个数只能使用一次，因此我们可以使用递归 + 回溯的方法来解决这个问题：
1. 我们用dfs(pos,rest)表示递归的函数，其中 pos 表示我们当前递归到了数组 candidates 中的第 pos 个数，
   而 rest 表示我们还需要选择和为 rest 的数放入列表作为一个组合；
2. 对于当前的第 pos 个数，我们有两种方法：选或者不选。如果我们选了这个数，那么我们调用 dfs(pos+1,rest−candidates[pos]) 进行递归，
   注意这里必须满足 rest ≥ candidates[pos]。如果我们不选这个数，那么我们调用 dfs(pos+1,rest) 进行递归；
3. 在某次递归开始前，如果 rest 的值为 0，说明我们找到了一个和为 target 的组合，将其放入答案中。
   每次调用递归函数前，如果我们选了那个数，就需要将其放入列表的末尾，该列表中存储了我们选的所有数。
   在回溯时，如果我们选了那个数，就要将其从列表的末尾删除。
标准的递归+回溯，但是存在重复组合
下面解法会出现重复解
 */
func CombinationSum2II(candidates []int, target int) [][]int {
	result,path := [][]int{}, []int{}
	length := len(candidates)
	used := make([]bool, length)
	var backtrack func(int, int)
	backtrack = func(sum, start int){
		if sum == 0{
			path_copy := make([]int, len(path))
			copy(path_copy, path)
			result = append(result, path_copy)
			return
		}
		if sum < 0 || start >= length{
			return
		}
		backtrack(sum, start + 1)
		if !used[start]{
			path = append(path, candidates[start])
			used[start] = true
			backtrack(sum - candidates[start], start)
			path = path[:len(path) - 1]
			used[start] = false
		}
	}
	backtrack(target, 0)
	return result
}
/*
这个方法最重要的作用是，可以让同一层级，不出现相同的元素。即
                  1
                 / \
                2   2  这种情况不会发生 但是却允许了不同层级之间的重复即：
               /     \
              5       5
                例2
                  1
                 /
                2      这种情况确是允许的
               /
              2
*/
func CombinationSum2(candidates []int, target int) [][]int {
	result,path := [][]int{}, []int{}
	length := len(candidates)
	if length == 0 || target <= 0{ //处理掉特殊情况
		return result
	}
	sort.Ints(candidates) // sort inplace，方便去重

	var backtrack func(int, int)
	backtrack = func(sum, start int){
		if sum == 0{
			path_copy := make([]int, len(path))
			copy(path_copy, path)
			result = append(result, path_copy)
			return
		}
		if sum < 0 || start >= length{
			return
		}
		for i := start; i < length; i++{
			if i != start && candidates[i] == candidates[i-1]{
				continue // 跳过和本轮循环中已访问过元素数值相同的元素们。
			}
			path = append(path, candidates[i])
			backtrack(sum - candidates[i], i+1)
			path = path[:len(path) - 1] // 回溯
		}

	}
	backtrack(target, 0)
	return result
}
// 77. Combinations
/*
剪枝：
说明： 如果 n = 7, k = 4，从 5 开始搜索就已经没有意义了，
	  这是因为：即使把 5 选上，后面的数只有 6 和 7，一共就 3 个候选数，凑不出 4 个数的组合。因此，搜索起点有上界
	  搜索起点的上界 = n - (k - path.size()) + 1
 */
func Combine(n int, k int) [][]int {
	result, path := [][]int{}, []int{}
	var backtrack func(target, start int)
	backtrack = func(target, start int){
		if target == 0{
			path_copy := make([]int, len(path))
			copy(path_copy, path)
			result = append(result, path_copy)
			return
		}
		for i := start; i <= n; i++{
			// 进行剪枝
			if n - i + 1 < target - len(path){
				fmt.Println(path)
				break
			}
			path = append(path, i)
			backtrack(target-1, i+1)
			path = path[:len(path) - 1]
		}

	}
	backtrack(k, 1)
	return result
}
/* 77. Combinations
考虑一个二进制数数字 x，它由 k 个 1 和 n−k 个 0 组成，如何找到它的字典序中的下一个数字 next(x)，这里分两种情况：
规则一：
	x 的最低位为 1，这种情况下，如果末尾由 t 个连续的 1，我们直接将倒数第 t 位的 1 和倒数第 t + 1 位的 0 替换，
	就可以得到 next(x)。如 0011=>0101， 0101=>0110, 1001=>1010, 1001111=>1100011。
规则二：
	x 的最低位为 0，这种情况下，末尾有 t 个连续的 0，而这 t 个连续的 0 之前有 m 个连续的 1，
	我们可以将倒数第 t + m 位置的 1 和倒数第 t + m + 1 位的 0 对换，然后把倒数第 t + 1位到倒数第 t + m - 1位的 1 移动到最低位。
	如 0110=>1001, 1010=>1100, 1011100=>1100011
链接：https://leetcode-cn.com/problems/combinations/solution/zu-he-by-leetcode-solution/
 */
func CombineII(n int, k int) (ans [][]int) {
	/* 初始化
	   将 temp 中 [0, k - 1] 每个位置 i 设置为 i + 1，即 [0, k - 1] 存 [1, k]
	 */
	temp := []int{}
	for i := 1; i <= k; i++{
		temp = append(temp, i)
	}
	temp = append(temp, n+1) //末尾加一位 n + 1 作为哨兵
	for i := 0; i < k; {
		comb := make([]int, k)
		copy(comb, temp[:k])
		ans = append(ans, comb)
		// 寻找第一个 temp[i] + 1 != temp[i + 1] 的位置 t
		// 我们需要把 [0, t - 1] 区间内的每个位置重置成 [1, t]
		for i = 0; i < k && temp[i]+1 == temp[i+1]; i++{
			temp[i] = i+1
		}
		// i 是第一个 temp[i] + 1 != temp[i + 1] 的位置
		temp[i]++
	}
	return
}
/* 2021-10-25 重新实现 采用递归 完成
** dfs(n, k) = dfs(n-1, k-1) + dfs(n-1, k)
** 分别对应选 与 不选 两种情况
** 注意 递归结束条件 和 剪枝条件
*/
func CombineRecursive(n int, k int) (ans [][]int) {
	// 根据 n 和 k 比较情况，进行剪枝
	if n < k{
		return
	}
	if n == k{// 递归出口-1
		t := []int{}
		for i := 1; i <= n; i++{
			t = append(t, i)
		}
		ans = append(ans, t)
	}
	if n > k{
		if k == 1{// 递归出口-2
			for i := 1; i <= n; i++{
				ans = append(ans, []int{i})
			}
		}else{// dfs(n, k) = dfs(n-1, k-1) + dfs(n-1, k)
			choose := CombineRecursive(n-1, k-1)
			ignore := CombineRecursive(n-1, k)
			for i := range choose{
				ans = append(ans, append(choose[i], n))
			}
			for i := range ignore{
				ans = append(ans, ignore[i])
			}
		}
	}
	return
}

/* Union Find 并查集
	0. 设定相距为1的值，为一个集合
    1. 用map记录 数字和对应的索引
    2. 用例有重复数字，map处理
    3. 连通数字和索引对应关系
    4. 遍历并查集数组，集合最大即为最长连续子序列
 */
// 128. Longest Consecutive Sequence
type UnionFindSet struct {
	id		[]int // 记录指向
	sz		[]int // 记录并查集某一子集大小
	count	int
}
func NewUnionFindSet(n int) *UnionFindSet {
	id, sz := make([]int, n), make([]int, n)
	for i := range id {
		id[i] = i
	}
	for i := range sz {
		sz[i] = 1
	}
	return &UnionFindSet{id: id, sz: sz, count: n}
}
func (u *UnionFindSet) Find(p int) int {
	for p != u.id[p] {
		p = u.id[p]
	}
	return p
}
func (u *UnionFindSet) Merge(p, q int){
	i := u.Find(p)
	j := u.Find(q)
	if i == j {
		return
	}
	if u.sz[i] < u.sz[j] {
		u.id[i] = j
		u.sz[j] += u.sz[i]
	}else{
		u.id[j] = i
		u.sz[i] += u.sz[j]
	}
	u.count--
}
func LongestConsecutiveUnionFind(nums []int) int {
	uf := NewUnionFindSet(len(nums))
	m := map[int]int{}
	for i := 0; i < len(nums); i++{
		cur := nums[i]
		if _,ok := m[cur]; ok {
			continue
		}
		m[cur] = i
		if idx, ok := m[cur - 1]; ok {
			uf.Merge(i, idx)
		}
		if idx, ok := m[cur + 1]; ok {
			uf.Merge(i, idx)
		}
	}
	res := 0
	for _,v := range uf.sz{
		if v > res{
			res = v
		}
	}
	return res
}

// 并查集写法2
func LongestConsecutiveUnionFind2(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	ans := 1
	// 初始化并查集
	rt,sz := map[int]int{}, map[int]int{}
	var find func(int)int
	var merge func(int, int)int
	find = func(v int)int{
		if rt[v] != v{
			rt[v] = find(rt[v])
		}
		return rt[v]
	}
	merge = func(u, v int) int{
		u = find(u)
		v = find(v)
		if u != v{ // u与v不在一个集合,将v合并到u的集合中去
			sz[u] += sz[v]
			rt[v] = rt[u]
		}
		return sz[u] // 返回当前集合的大小
	}
	for _,i := range nums {
		rt[i] = i;// rt[v] = v 代表v是自己所在的集合的根
		sz[i] = 1;
	}
	for _,i := range nums {
		// 连续数组，只需考虑v与v-1就能覆盖掉所有边的情况
		if _,ok := rt[i-1]; ok {
			length := merge(i, i-1)
			if ans < length{
				ans = length
			}
		}
	}
	return ans
}

// 并查集写法3
type unionFindSet struct{
	fa map[int]int
	sz map[int]int
}
func unionFindSetNew(nums []int) *unionFindSet{
	ufs := unionFindSet{fa: map[int]int{}, sz: map[int]int{}}
	for _,u := range nums{
		ufs.fa[u] = u
		ufs.sz[u] = 1
	}
	return &ufs
}
/* 优化1-路径压缩
  find执行的操作：从一个节点，不停的通过parent向上去寻找他的根节点，
  在这个过程中，我们相当于把从这个节点到根节点的这条路径上的所有的节点都给遍历了一遍，
  那么，让我们想一想，在find的同时，是否可以顺便加上一些其它的操作使得树的层数尽量变得更少呢？答案是可以的。
请看下面的例子：
  对于一个集合树来说，它的根节点下面可以依附着许多的节点，因此，我们可以尝试在find的过程中，
  从底向上，如果此时访问的节点不是根节点的话，那么我们可以把这个节点尽量的往上挪一挪，减少数的层数，
  这个过程就叫做路径压缩
 */
func (ufs *unionFindSet) Find(p int) int{
	for p != ufs.fa[p] {
		// 优化1：路径压缩
		// ufs.fa[p] = ufs.fa[ufs.fa[p]] // 方案一，只能压缩一部分，不能压缩至最低
		ufs.fa[p] = ufs.Find(ufs.fa[p]) // 方案二, 递归实现可转循环 将所有元素直接指向根节点，最低压缩
		p = ufs.fa[p]
	}
	return p
}
/* 优化2： 按秩合并减少树高度
   按秩合并是一种启发式合并，主要思想是合并的时候把小的树合并到大的树以减少工作量。
   秩的定义： 1. 树的高度  2. 树的节点数
   我们在路径压缩之后一般采用第二种，因为第一种在路径压缩之后就已经失去意义了，按照第二种合并可以减少一定的路径压缩的工作量
   单独采用按秩合并的话平摊查询时间复杂度同样为O(logN)
   如果我们把路径压缩和按秩合并合起来一起使用的话可以把查询复杂度下降到O(α(n))，其中α(n)为反阿克曼函数。
   阿克曼函数是一个增长极其迅速的函数，而相对的反阿克曼函数是一个增长极其缓慢的函数，所以我们在算时间复杂度的时候可以把他视作一个常数看待。
 */
func (ufs *unionFindSet) Merge(p, q int) int {
	// find p q 的根
	set1 := ufs.Find(p)
	set2 := ufs.Find(q)
	// 如果p q 有相同的根，即 set1 == set2，则直接返回
	// 表示当前的数与之前数已经构成连续序列
	if set1 == set2 {
		return ufs.sz[set1]
	}
	// 子树合并优化2： 压缩合并树的高度
	// 将较高的树作为根节点，将较矮的树连在其上
	// 两棵树高度相同时，不管哪个作为根节点，另一个脸上它之后，整棵树的高度就要加1
	// 后两种情况，合并后树的最大高度不变
	if ufs.sz[set1] < ufs.sz[set2]{
		ufs.fa[set1] = set2
		ufs.sz[set2] += ufs.sz[set1]
		return ufs.sz[set2]
	}else {
		ufs.fa[set2] = set1
		ufs.sz[set1] += ufs.sz[set2]
		return ufs.sz[set1]
	}
	// 子树合并
	ufs.fa[set2] = set1
	ufs.sz[set1] += ufs.sz[set2]
	return ufs.sz[set1]
}
func LongestConsecutiveUFSHash(nums []int) (count int) {
	ufs := unionFindSetNew(nums)
	for _,item := range nums{
		if _,ok := ufs.fa[item+1]; ok {
			length := ufs.Merge(item+1, item)
			if length > count {
				count = length
			}
		}
	}
	return
}




/*  leetcode官方解法
   考虑数组中每个数x,考虑以其为起点，不断尝试匹配 x+1 x+2 ... 是否存在，假设匹配到了x+y，那么以x起点的最长连续序列即为 x x+1 ... x+y 长度为
   y+1，然后不断的枚举并更新答案即可。
   对于匹配的过程，暴力的方法是 O(n) 遍历数组去看是否存在这个数，但其实更高效的方法是用一个哈希表存储数组中的数，
   这样查看一个数是否存在即能优化至 O(1)的时间复杂度

 */
func LongestConsecutiveHash(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	length := 1
	uniqueNums := map[int]bool{}
	for _,elem := range nums{
		uniqueNums[elem] = true
	}
	for elem,_ := range uniqueNums{
		start := elem
		cnt := 0
		_, ok := uniqueNums[start]
		for ok {
			start++
			_, ok = uniqueNums[start]
			cnt++
		}
		if cnt > length{
			length = cnt
		}
	}
	return length
}
// 优化法
/*
 我们会发现其中执行了很多不必要的枚举，如果已知有一个x,x+1,x+2,⋯,x+y 的连续序列，而我们却重新从x+1 x+2 或者是 x+y 处开始尝试匹配
 那么得到的结果肯定不会优于枚举 x 为起点的答案，因此我们在外层循环的时候碰到这种情况跳过即可。

那么怎么判断是否跳过呢？由于我们要枚举的数 x 一定是在数组中不存在前驱数x−1 的，
不然按照上面的分析我们会从 x−1 开始尝试匹配，因此我们每次在哈希表中检查是否存在 x−1 即能判断是否需要跳过了。

 */
func LongestConsecutiveHashImprove(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	length := 1
	uniqueNums := map[int]bool{}
	for _,elem := range nums{
		uniqueNums[elem] = true
	}
	for elem := range uniqueNums{
		//检查 elem - 1
		if _,ok := uniqueNums[elem - 1]; ok {
			continue
		}
		cnt := 1
		for uniqueNums[elem+1] {
			elem++
			cnt++
		}
		if cnt > length{
			length = cnt
		}
	}
	return length
}
func LongestConsecutive(nums []int) int {
	if len(nums) <= 0{
		return 0
	}
	mark := map[int]int{}
	ans := 1
	for _,elem := range nums{
		if _,ok := mark[elem]; !ok {
			mark[elem] = elem // 刚插入时，v的左右边界都是他本身
			for _,v := range []int{elem-1, elem+1} {// 遍历左右邻居
				if _,ok := mark[v]; ok {// 如果邻居已经在mark表中
					x,y := mark[elem], mark[v] // 用x代表v的另一端位置，用y代表v+i的另一端位置
					// 两个端点分别记录位置
					mark[x] = y
					mark[y] = x
					length := y - x
					if x > y{
						length = x - y
					}
					if ans < length+1{
						ans = length + 1
					}
				}
			}
		}
	}
	return ans
}
/* DP 未优化版本，提交会超时
    对于每个v我们都从v向v+1连一条线的话，输入数据就会成为一个有向无环图；
    也即求有向无环图求最长路的方法
    可以用一个基于hash的map记录答案。 mp[v]代表以v为起点的最长路的长度，同时有
	递推式：mp[v] = mp[v+1] + 1, if v+1 in mp
	基情况: mp[v] = 0, if v not in mp
 */
func LongestConsecutiveDP(nums []int) int {
	umap := map[int]int{}
	answer := 0
	var dfs func(int) int
	dfs = func(item int) (length int){
		if _,ok := umap[item];!ok{
			return 0
		}
		return dfs(item + 1) +1
	}
	for _,elem := range nums {
		umap[elem] = 0
	}
	for _,elem := range nums{
		length := dfs(elem)
		if length > answer{
			answer = length
		}
	}
	return answer
}
func LongestConsecutiveDPImprove(nums []int) int {
	umap := map[int]int{}
	answer := 0
	var dfs func(int) int
	dfs = func(item int) (length int){
		if _,ok := umap[item];!ok{
			return 0
		}
		if umap[item] != 0{
			return umap[item] //缓存剪枝
		}
		return dfs(item + 1) +1
	}
	for _,elem := range nums {
		umap[elem] = 0
	}
	for _,elem := range nums{
		length := dfs(elem)
		umap[elem] = length // 缓存剪枝
		if length > answer{
			answer = length
		}
	}
	return answer
}

//721: Account name
// 可以处理用户名不同，但有相同邮箱地址的情况
func AccountsMerge(accounts [][]string) [][]string {
	total := len(accounts)
	ufs := make([]int, total)
	for i := 0; i < total; i++{
		ufs[i] = i
	}
	var find func(int) int
	var union func(int,int)
	find = func(i int) int {
		if i != ufs[i]{// 路径压缩
			return find(ufs[i])
		}
		return i
	}
	union = func(i, j int) {
		pi := find(i)
		pj := find(j)
		if pi != pj{
			ufs[pi] = pj
		}
	}
	// 开始处理,找到要合并的项
	emailMap := map[string][]int{}
	for i := 0; i < total; i++{
		for _,mail := range accounts[i][1:]{
			if _,ok := emailMap[mail]; ok {
				emailMap[mail] = append(emailMap[mail], i)
			}else{
				emailMap[mail] = []int{i}
			}
		}
	}
	// union
	for _,values := range emailMap{
		length := len(values)
		if length != 0{
			start := values[0]
			for i := 1; i < length; i++{
				// union(start, i) 晕了晕了
				union(start, values[i])
			}
		}
	}
	// 核算结果
	result := map[int][]string{}
	for idx,account := range accounts{
		class := find(idx)
		if _,ok := result[class]; ok {
			result[class] = append(result[class], account[1:]...)
		}else{
			result[class] = append(result[class], account...)
		}
	}
	// answer
	answer := [][]string{}
	for _,item := range result {
		sort.Strings(item[1:])
		itemCopy := []string{item[0]}
		for i := 1; i < len(item); i++{
			if item[i] != item[i-1]{
				itemCopy = append(itemCopy, item[i])
			}
		}
		answer = append(answer, itemCopy)
	}
	return answer
}

// 官方解法
func accountsMerge(accounts [][]string) (ans [][]string) {
	email2index := map[string]int{}
	email2Name := map[string]string{}
	for _,account := range accounts {
		name := account[0]
		for _,mail := range account[1:]{
			if _, ok := email2index[mail]; !ok {
				email2index[mail] = len(email2index)
				email2Name[mail] = name //名称必定相同
			}
		}
	}
	// 并查集初始化
	parent := make([]int, len(email2index)) // 用每个email做集合
	for i := range parent{
		parent[i] = i
	}
	var find func(int) int
	find = func(i int) int {
		if parent[i] != i{ // 路径压缩
			parent[i] = find(parent[i])
		}
		return parent[i]
	}
	union := func(from, to int){
		pf, pt := find(from), find(to)
		parent[pf] = pt
	}
	// 分组合并，条件是 所属同一个用户
	for _, account := range accounts{
		firstIndex := email2index[account[1]]
		for _, email := range account[2:] {
			union(firstIndex, email2index[email])
		}
	}
	// 核算结果
	index2email := map[int][]string{}
	for email, idx := range email2index{
		idx = find(idx)
		index2email[idx] = append(index2email[idx], email)
	}
	for _,emails := range index2email{
		sort.Strings(emails)
		account := append([]string{email2Name[emails[0]]}, emails...)
		ans = append(ans, account)
	}
	return
}
// 抽象成邻接图， 使用dfs 计算图的 连通性
func AccountsMergeDfs(accounts [][]string) (ans [][]string) {
	// build graph 忽略姓名，直接把关联的账户连接起来即可
	graph := map[string][]string{}
	for _,account := range accounts{
		master := account[1]
		for _, email := range account[2:]{
			graph[master] = append(graph[master], email)
			graph[email] = append(graph[email], master)
		}
	}
	res := [][]string{}
	visited := map[string]bool{} // 集合实现
	var dfs func(string)[]string
	dfs = func(email string)(emails []string){
		if visited[email]{
			return
		}
		visited[email] = true
		emails = append(emails, email)
		for _, neighbor := range graph[email]{
			emails = append(emails, dfs(neighbor)...)
		}
		return
	}

	for _, account := range accounts{
		emails := dfs(account[1])
		if len(emails) > 0{
			data := []string{account[0]}
			sort.Strings(emails)
			data = append(data, emails...)
			res = append(res, data)
		}
	}
	return res
}

//17.07. Baby Names LCCI
func TrulyMostPopular(names []string, synonyms []string) []string {
	babyNames := map[string]int64{}
	map2index := map[string]int{}
	index2map := map[int]string{}
	for _, name := range names{
		r, err := regexp.Compile(`([a-zA-Z]+)\(([0-9]+)\)`)
		if err != nil {
			continue
		}
		tmp := r.FindStringSubmatch(name)
		if len(tmp) < 3{
			continue
		}
		babyNames[tmp[1]], err = strconv.ParseInt(tmp[2], 10, 64)
		if err != nil {
			continue
		}
		map2index[tmp[1]] = len(babyNames)-1
		index2map[len(babyNames)-1] = tmp[1]
	}
	ufs := make([]int, len(babyNames))
	for idx,_ := range ufs{
		ufs[idx] = idx
	}
	fmt.Println(index2map)
	var find func(int)int
	var union func(int, int)
	find = func(i int) int{
		if ufs[i] != i{
			ufs[i] = find(ufs[i])
		}
		return ufs[i]
	}
	union = func(i, j int){
		pi, pj := find(i), find(j)
		if pi != pj{
			tmp := []string{index2map[pi], index2map[pj]}
			sort.Strings(tmp)
			ufs[map2index[tmp[1]]] = ufs[map2index[tmp[0]]]
			babyNames[tmp[0]] += babyNames[tmp[1]]
		}
	}
	for _, item := range synonyms {
		r, err := regexp.Compile(`\(([a-zA-Z]+),([a-zA-Z]+)\)`)
		if err != nil {
			continue
		}
		tmp := r.FindStringSubmatch(item)
		if len(tmp) < 3{
			continue
		}
		union(map2index[tmp[1]], map2index[tmp[2]])
	}
	res := map[string]int64{}
	for key,_ := range babyNames{
		p := find(map2index[key])
		if _, ok := res[index2map[p]]; !ok{
			res[index2map[p]] = babyNames[index2map[p]]
		}
	}
	ans := []string{}
	for key, value := range res{
		ans = append(ans, fmt.Sprintf("%s(%d)", key, value))
	}
	return ans
}
// 结构体表示
func TrulyMostPopularII(names []string, synonyms []string) []string {
	type node struct {
		freq 		uint64
		setName 	string // 类别
	}
	babyNames := map[string]node{}
	for _, name := range names{
		 r, _ := regexp.Compile(`([A-Za-z]+)\(([0-9]+)\)`)
		 tmp := r.FindStringSubmatch(name)
		 freq,_ := strconv.ParseUint(tmp[2], 10, 64)
		 babyNames[tmp[1]]=node{freq: freq, setName: tmp[1]}
	}
	var find func(i string)string
	find = func(i string)string{
		// 需要注意 names不含有synonyms的情况
		if _,ok := babyNames[i]; !ok{
			return i
		}
		if i != babyNames[i].setName{
			cache := babyNames[i]
			cache.setName = find(cache.setName)// 路径压缩
			babyNames[i] = cache
			return babyNames[i].setName
		}
		return i
		//return babyNames[i].setName
	}
	var union func(i, j string)
	union = func(i, j string){
		pi := find(i)
		pj := find(j)
		if pi != pj{
			//babyNames[pi].freq += babyNames[pj].freq
			// 按照字典序合并, 字典序小的作为根
			picknames := []string{pi, pj}
			sort.Strings(picknames)
			namePicked := picknames[0]
			freq := babyNames[pi].freq + babyNames[pj].freq
			// TODO: 按秩合并
			cache := babyNames[pj]
			cache.setName = namePicked
			cache.freq = freq
			babyNames[pj] = cache
			babyNames[pi] = cache
		}
	}
	for _, name := range synonyms{
		r, _ := regexp.Compile(`\(([A-Za-z]+),([A-Za-z]+)\)`)
		tmp := r.FindStringSubmatch(name)
		union(tmp[1], tmp[2])
	}
	visited := map[string]bool{}
	ans := []string{}
	for _, item := range babyNames{
		setName := find(item.setName)
		freq := babyNames[setName].freq
		if !visited[setName]{
			visited[setName] = true
			//ans = append(ans, fmt.Sprintf("%s(%d)", item.setName, item.freq))  错误： 不能直接用item， 确定 item所在类的head的名子
			//ans = append(ans, fmt.Sprintf("%s(%d)", setName, item.freq)) 错误：freq不能直接去item的 需要取真正类head的freq，因为集合中有些元素没有更新freq
			ans = append(ans, fmt.Sprintf("%s(%d)", setName, freq))
		}
	}
	return ans
}
// 不使用正则表达式
func TrulyMostPopularIII(names []string, synonyms []string) []string {
	type node struct {
		freq 		uint64
		setName 	string // 类别
	}
	babyNames := map[string]node{}
	for i := 0 ; i < len(names);  i ++{
		a := strings.Index(names[i],"(")
		b := strings.Index(names[i],")")
		name := names[i][:a]
		temp := names[i][a+1:b]
		freq, _ := strconv.ParseUint(temp, 10, 64)
		babyNames[name]=node{freq: freq, setName: name}
	}
	var find func(i string)string
	find = func(i string)string{
		// 需要注意 names不含有synonyms的情况
		if _,ok := babyNames[i]; !ok{
			return i
		}
		if i != babyNames[i].setName{
			cache := babyNames[i]
			cache.setName = find(cache.setName)// 路径压缩
			babyNames[i] = cache
			return babyNames[i].setName
		}
		return i
		//return babyNames[i].setName
	}
	var union func(i, j string)
	union = func(i, j string){
		pi := find(i)
		pj := find(j)
		if pi != pj{
			//babyNames[pi].freq += babyNames[pj].freq
			// 按照字典序合并, 字典序小的作为根
			picknames := []string{pi, pj}
			sort.Strings(picknames)
			namePicked := picknames[0]
			freq := babyNames[pi].freq + babyNames[pj].freq
			// TODO: 按秩合并
			cache := babyNames[pj]
			cache.setName = namePicked
			cache.freq = freq
			babyNames[pj] = cache
			babyNames[pi] = cache
		}
	}
	for i := 0 ; i < len(synonyms) ; i ++{
		a := strings.Index(synonyms[i],"(")
		b := strings.Index(synonyms[i],",")
		c := strings.Index(synonyms[i],")")
		temp1 := synonyms[i][a+1:b]
		temp2 := synonyms[i][b+1:c]
		union(temp1,temp2)
	}
	visited := map[string]bool{}
	ans := []string{}
	for _, item := range babyNames{
		setName := find(item.setName)
		freq := babyNames[setName].freq
		if !visited[setName]{
			visited[setName] = true
			//ans = append(ans, fmt.Sprintf("%s(%d)", item.setName, item.freq))  错误： 不能直接用item， 确定 item所在类的head的名子
			//ans = append(ans, fmt.Sprintf("%s(%d)", setName, item.freq)) 错误：freq不能直接去item的 需要取真正类head的freq，因为集合中有些元素没有更新freq
			ans = append(ans, fmt.Sprintf("%s(%d)", setName, freq))
		}
	}
	return ans
}
// 200. Number of islands
/*
知识点：1. golang 中 数组是可以比较的 2. 数组可以作为map的key
 */
func NumIslands(grid [][]byte) int {
	index2map := map[[2]int][2]int{}
	for i, row := range grid{
		for j,elem := range row{
			if elem == '1'{
				loc := [2]int{i,j}
				index2map[loc] = loc
			}
		}
	}
	var find func(m [2]int) [2]int
	find = func(m [2]int) [2]int {
		if m != index2map[m]{
			return find(index2map[m])
		}
		return m
	}
	var union func(x, y [2]int)
	union = func(x, y [2]int){
		px, py := find(x), find(y)
		// if px != px 错误点-1： 低级错误
		if px != py {
			//index2map[x] = py  错误点-2： 是类 低级错误
			index2map[px] = py
		}
	}
	rowLen := len(grid)
	for i := 0; i < rowLen; i++ {
		colLen := len(grid[i])
		for j := 0; j < colLen; j++{
			if j+1 < colLen && grid[i][j] == '1' && grid[i][j+1] == '1'{
				union([2]int{i,j}, [2]int{i, j+1})
			}
			if i+1 < rowLen && grid[i][j] == '1' && grid[i+1][j] == '1'{
				union([2]int{i,j}, [2]int{i+1, j})
			}
		}
	}
	islandSet := map[[2]int]bool{}
	//for _,elem := range index2map{ 统计的是key， value可能重复的 低级错误
	for elem,_ := range index2map{
		island := find(elem)
		if !islandSet[island]{
			islandSet[island] = true
		}
	}
	return len(islandSet)
}

func NumIslandsDFS(grid [][]byte) int {
	total, rowLen, colLen := 0, len(grid), len(grid[0])
	var dfs func(int, int)
	dfs = func(r, c int){
		grid[r][c] = 0 // 置0
		if c+1 < colLen && grid[r][c+1] == '1'{
			dfs(r, c+1)
		}
		if r+1 < rowLen && grid[r+1][c] == '1'{
			dfs(r+1, c)
		}
		// dfs 与并差集不同，需要4个方向来判定
		if c - 1 >= 0 && grid[r][c-1] == '1'{
			dfs(r, c-1)
		}
		if r - 1 >= 0 && grid[r-1][c] == '1'{
			dfs(r-1, c)
		}
	}
	for i, row := range grid{
		for j, elem := range row{
			if elem == '1'{
				total++ // 岛屿的数量就是我们进行dfs搜索的次数
				dfs(i, j)
			}
		}
	}
	return total
}
// 利用按秩压缩
type UnionFind struct {
	count		int
	parent		[]int
	rank		[]int
}

func ConstructUnionFindByNumIslands(grid [][]byte)*UnionFind{
	// 初始化
	m, n, cnt := len(grid),len(grid[0]), 0
	rank, parent := make([]int, m*n), make([]int, m*n)
	for i := 0; i < m; i++{
		for j := 0; j < n; j++{
			if grid[i][j] == '1'{
				parent[i * n + j] = i * n + j // 2. 类
				cnt++ // 1. 计数
			}
			rank[i * n + j] = 0 // 3. 秩
		}
	}
	return &UnionFind{count: cnt, parent: parent, rank: rank}

}

func (ufs *UnionFind) find(i int)int{
	if ufs.parent[i] != i {
		ufs.parent[i] = ufs.find(ufs.parent[i]) // 1. 路径压缩
	}
	return ufs.parent[i]
}

func(ufs *UnionFind) union(x, y int){
	rootx, rooty := ufs.find(x), ufs.find(y)
	if rootx != rooty{ // 4. 按秩压缩： 查询高度小的 插入到 查询高度大的
		if ufs.rank[rootx] > ufs.rank[rooty]{
			ufs.parent[rooty] = rootx
		}else if ufs.rank[rootx] < ufs.rank[rooty]{
			ufs.parent[rootx] = rooty
		}else{
			ufs.parent[rooty] = rootx
			ufs.rank[rootx] += 1 // 2. 按秩压缩，查询高度记录
		}
		ufs.count-- // 3. union 后 计数递减
	}
}

func NumIslandsUFSImprove(grid [][]byte) int {
	if len(grid) <= 0{
		return 0
	}
	nr, nc := len(grid), len(grid[0])
	pUfs := ConstructUnionFindByNumIslands(grid)
	for r := 0; r < nr; r++{
		for c := 0; c < nc; c++{
			if grid[r][c] == '1' {
				grid[r][c] = '0'
				cur := r * nc + c
				if r - 1 >= 0 && grid[r-1][c] == '1'{
					pUfs.union(cur, (r-1) * nc + c)
				}
				if r + 1 < nr && grid[r+1][c] == '1'{
					pUfs.union(cur, (r+1) * nc + c)
				}
				if c - 1 >= 0 && grid[r][c-1] == '1'{
					pUfs.union(cur, r * nc + c - 1)
				}
				if c + 1 < nc && grid[r][c+1] == '1'{
					pUfs.union(cur, r * nc + c + 1)
				}
			}
		}
	}
	return pUfs.count
}

/* 327. Count of Range Sum（区间和的个数）
** Given an integer array nums and two integers lower and upper,
** return the number of range sums that lie in [lower, upper] inclusive.
** Range sum S(i, j) is defined as the sum of the elements in nums between indices i and j inclusive, where i <= j.
** Constraints：
** 1 <= nums.length <= 105
** -2的31次方 <= nums[i] <= 2的31次方 - 1   注意溢出的情况
** -10的5次方 <= lower <= upper <= 10的5次方
** The answer is guaranteed to fit in a 32-bit integer.
 */
/*
	前缀和：设前缀和数组为preSum 则问题转换为 求所有的下标对(i,j)，满足 preSum[j] - preSum[i] 属于 [lower, upper]
    O(n^2) 会超时
	另外注意到条件-2的31次方 <= nums[i] <= 2的31次方 - 1， 注意溢出的情况
*/
func CountRangeSum(nums []int, lower int, upper int) int {
	length := len(nums)
	if length <= 0{
		return 0
	}
	// 先计算前缀和
	// 易错点-2 需要long long 类型，因为可能会溢出，见case-2情况
	preSum := []int64{0} // 首元素置0，前缀和可以标识单元素
	for i := 0; i < length; i++{
		preSum = append(preSum, preSum[i] + int64(nums[i]))
	}
	ans := 0
	// 下面是遍历preSum  此处需要优化，平方级别复杂度，见解法II 利用归并排序降低时间复杂度
	for i := 1; i < len(preSum); i++{
		for j := 0; j < i; j++{
			//diff := preSum[j] - preSum[i] 错误点-1： 顺序颠倒
			diff := preSum[i] - preSum[j]
			if diff >= int64(lower) && diff <= int64(upper){
				ans++
			}
		}
	}
	return ans
}
type prefixSum struct {
	sum int64 // 当前前缀和
	idx int // 前缀和所在数组的位置下标
}
/* 二分查找/插入
** 前缀和条件转换为： prefixSum[j] - lower >= x >= prefixSum[j] - upper x由prefixSum[0...j-1]选出
** 通过维护一个有序数组sumArr，表示当前的已知区间和，然后每次向右推进右端点，得到一个新的前缀和，
** 在升序存放prefixSum[0...j-1]的数组中，二分查找：
** 1. 找第一个 l 使得 x 大于等于 prefixSum[j] - upper
** 2. 找第一个 r 使得 x 大于 prefixSum[j] - lower
** r - l 就是当前右端点prefixSum[j]可以组成的区间和个数
** 然后再把prefixSum[j] 加入到 有序数组sumArr（插入位置也是二分搜索得到），准备为下一个prefixSum[j]做准备
** const(
		lower_bound	= iota	// 返回第一个大于等于给定值所在的位置
		upper_bound			// 返回第一个大于给定元素值所在的位置
		insert_loction		// 返回给定元素值待插入的位置
	)
 */
/* 在得出 prefixSum[j] - lower >= x >= prefixSum[j] - upper x由prefixSum[0...j-1]选出
** 此时需要一个数据结构来支持下面的2个操作：
** 1. 查询： 给定一个范围[left, right]，查询数据结构中该范围内的元素个数，即范围[prefixSum[j] - lower, prefixSum[j] - upper]
** 2. 更新： 给定一个元素 x， 需要将它添加到数据结构中，即给定元素 prefixSum[j]
** 从而
** 首先将 0 放入数据结构中，随后我们从小到大枚举 ，查询 [prefixSum[j] - lower, prefixSum[j] - upper] 范围内的元素个数并计入答案。
** 在查询完成之后，我们将 P[j]P[j] 添加进数据结构，就可以进行下一次枚举
*/
func CountRangeSumBinarySearch(nums []int, lower int, upper int) int {
	// prefixSumArr := []prefixSum{}
	prefixSumArr := []prefixSum{{0,-1}}
	search := func(n int, f func(int)bool)int{
		i, j := 0, n
		for i < j{
			mid := int(uint(i+j)>>1) // 防止溢出的一种方法
			if !f(mid){
				i = mid + 1 // f(i-1)== false
			}else{
				j = mid
			}
		}
		return i
	}
	ans, sum := 0, int64(0)
	for i := range nums{
		sum += int64(nums[i])
		n := len(prefixSumArr)
		lower_bound := func(mid int)bool{// 返回第一个大于等于给定值所在的位置
			return prefixSumArr[mid].sum >= sum - int64(upper)
		}
		upper_bound := func(mid int)bool{// 返回第一个大于给定元素值所在的位置
			return prefixSumArr[mid].sum >     sum - int64(lower)
		}
		insert_loction := func(mid int)bool {// 返回给定元素值待插入的位置
			return prefixSumArr[mid].sum > sum
		}
		l := search(n, lower_bound)
		r := search(n, upper_bound)
		ans += r - l
		loc := search(n, insert_loction)
		prefixSumArr = append(prefixSumArr, prefixSum{})
		copy(prefixSumArr[loc+1:], prefixSumArr[loc:])
		prefixSumArr[loc] = prefixSum{sum, i}
	}
	if len(nums) < 6{
		fmt.Println(prefixSumArr)
	}
	return ans
}

/* 上面二分查找方法中：
** 在最坏情况维护的二分查找数组，查找时可能会退化到O(n)，而不是稳定O(log(n))，
** 为了维护稳定的O(log(n))，我们需要使用平衡树来代替之前的有序数组，进行二分查找
 */
func CountRangeSumByMap(nums []int, lower int, upper int) int {
	return 0
}
/* 归并排序: 求所有的下标对(i,j), 满足 preSum[j] - preSum[i] 属于 [lower, upper]
** 给定两个升序排列的数组 n1 和 n2，尝试找出所有的下标对 (i, j),满足 n2[j] - n1[i] 属于 [lower, upper]
** 在已知两个序列升序的前提下，在 n2 中维护2个指针 left, right 均指向 n2 的起始位置
** 开始遍历 n1，考察n1的第一个元素，不断的将指针 left 向右移动，直至 n2[left] >= n1[0]+lower。
** 此时 left 及其右边的元素均大于等于 n1[0]+lower,
** 随后再不断地将指针right向右移动，直至n2[right] > n1[0] + upper，则 r 左边的元素均小于或等于 n1[0] + upper。
** 此时 right - left 即满足的下标对个数
** 由于 n1 是递增的，不难发现 left 和 right 只可能向右移动。不断重复此过程，对应n1的每个下标，都记录响应的区间[left, right)的大小
** 我们采用归并排序的方式，能够得到左右两个数组排序后的形式，以及对应的下标对数量。
** 对于原数组而言，若要找出全部的下标对数量，只需要再额外找出左端点在左侧数组，同时右端点在右侧数组的下标对数量，而这正是我们此前讨论的问题
 */
func CountRangeSumMergeSort(nums []int, lower, upper int) int {
	// 前缀和
	prefixSum := make([]int, len(nums) + 1)
	for i, v := range nums{
		prefixSum[i+1] = prefixSum[i] + v
	}
	var mergeCount func([]int) int
	mergeCount = func(arr []int) int {
		n := len(arr)
		if n <= 1{
			return 0
		}
		n1 := append([]int{}, arr[:n/2]...)
		n2 := append([]int{}, arr[n/2:]...)
		cnt := mergeCount(n1) + mergeCount(n2)
		// 此时 n1 和 n2 均有序，升序
		// 算法开始：统计下标对的数量
		left, right := 0, 0
		for _, v := range n1{
			// 直至 n2[left] >= n1[0]+lower
			for left < len(n2) && n2[left] - v < lower{
				left++
			}
			//直至n2[right] > n1[0] + upper
			for right < len(n2) && n2[right] - v <= upper {
				right++
			}
			cnt += right - left
		}
		// n1 和 n2 归并填入 arr
		p1, p2 := 0, 0
		for i := range arr {
			if  p1 < len(n1) && (p2 == len(n2) || n1[p1] <= n2[p2]) {
				arr[i] = n1[p1]
				p1++
			}else{
				arr[i] = n2[p2]
				p2++
			}
		}
		return cnt
	}
	return mergeCount(prefixSum)
}
/* 树状数组-
wiki: A Fenwick tree or binary indexed tree is
      a data structure that can efficiently update elements and calculate prefix sums in a table of numbers.
** 树状数组单次更新或查询的复杂度为 O(logN)
** ==> 前因 <==
** 主要是用来解决前缀和的问题，前缀和有3个操作：
** prefixSum(idx)：直接返回前缀和数组prefixSumArr[idx + 1]即可。该操作为O(1)时间复杂度
** rangeSum(from_idx, to_idx)：直接返回prefixSumArr[to_idx + 1] - prefixSumArr[from_idx]即可。该操作为O(1)操作。
** update(idx, delta)：更新操作需要更新prefixSumArr数组中每一个受此更新影响的前缀和，即从idx其到最后一个位置的前缀和。该操作为O(n)时间复杂度
** 而树状数组就是来解决这个update问题的，即为了在保证求和操作依然高效的前提下优化update(idx, delta) 操作的时间复杂度
** ==> 基本思想 <==
** Binary Indexed Tree事实上是将根据数字的二进制表示来对数组中的元素进行逻辑上的分层存储
** Binary Indexed Tree求和的基本思想在于，给定需要求和的位置i，例如13，
** 我们可以利用其二进制表示法来进行分段（或者说分层）求和：
** 13 = 2^3 + 2^2 + 2^0，则prefixSum(13) = RANGE(1, 8) + RANGE(9, 12) + RANGE(13, 13)
** （注意此处的RANGE(x, y)表示数组中第x个位置到第y个位置的所有数字求和）
 */
type fenwick struct {
	tree []int
}
func (f fenwick) len()int{
	return len(f.tree)
}
func (f fenwick) inc(i int){
	for ; i < f.len(); i += i&-i{
		f.tree[i]++
	}
}
func (f fenwick) sum(i int)(res int){
	for ; i > 0; i &= i-1{
		res += f.tree[i]
	}
	return
}
func (f fenwick) query(l, r int)(res int){
	return f.sum(r) - f.sum(l-1)
}
func CountRangeSumFenwick(nums []int, lower, upper int) (cnt int) {
	n := len(nums)
	// 计算前缀和 preSum，以及后面统计时会用到的所有数字 allNums
	allNums := make([]int, 1, 3*n+1)
	preSum := make([]int, n+1)
	for i := range nums{
		preSum[i+1] = preSum[i] + nums[i]
		allNums = append(allNums, preSum[i+1], preSum[i+1]-lower, preSum[i+1]-upper)
	}
	// 将 allNums 离散化
	sort.Ints(allNums)
	k := 1
	kth := map[int]int{allNums[0]: k}
	for i := 1; i <= 3*n; i++{
		if allNums[i] != allNums[i-1]{
			k++
			kth[allNums[i]] = k
		}
	}
	// 遍历 preSum，利用树状数组计算每个前缀和对应的合法区间数
	t := fenwick{make([]int, k+1)}
	t.inc(kth[0])
	for _, sum := range preSum[1:]{
		left, right := kth[sum-upper], kth[sum-lower]
		cnt += t.query(left, right)
		t.inc(kth[sum])
	}
	return
}

// 17.10. Find Majority Element LCCI
/*
	摩尔投票法
 */
func majorityElement(nums []int) int {
	candidate := -1
	count := 0
	for _, num := range nums {
		if count == 0{
			candidate = num
		}
		if num == candidate{
			count++
		}else{
			count --
		}
	}
	count = 0
	for _, num := range nums{
		if num == candidate{
			count++
		}
	}
	if count * 2 > len(nums){
		return candidate
	}
	return -1
}
/* 并查集 以及 图系列
 */
/* 399. Evaluate Division
	抽象为并查集处理 带权有向图
	由于变量之间的倍数关系具有传递性，处理有传递性关系的问题，可以使用「并查集」，
	我们需要在并查集的「合并」与「查询」操作中 维护这些变量之间的倍数关系
	1. 同在一个集合中的两个变量就可以通过某种方式计算出它们的比值。具体来说，可以把 不同的变量的比值转换成为相同的变量的比值
	2. 如果两个变量不在同一个集合中， 返回 -1.0

	对于任意两点x y  假设他们在并查集中具有共同的parent, 并且
	v[x] / v[parent] = a
	v[y] / v[parent] = b ==> v[x] / v[y] = a / b
 */

func CalcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 给方程组中的每个变量编号
	id := map[string]int{}
	for _, op := range equations{
		if _, ok := id[op[0]]; !ok{
			id[op[0]] = len(id)
		}
		if _, ok := id[op[1]]; !ok {
			id[op[1]] = len(id)
		}
	}
	parent := make([]int, len(id))
	w := make([]float64, len(id))
	// 初始化
	for i := range parent{
		parent[i] = i
		w[i] = 1
	}

	var find func(int)int
	var union func(int, int, float64)
	find = func(x int) int {
		if parent[x] != x{
			f := find(parent[x])
			/* a/b=3; b/c=0.5; a/c=3*0.5 ==> a = (3*05)c*/
			w[x] *= w[parent[x]]
			parent[x] = f
		}
		return parent[x]
	}
	union = func(i, j int, val float64){
		pi, pj := find(i), find(j)
		//w[pi] = val * w[pj] / w[pi]
		w[pi] = val * w[j] / w[i]
		parent[pi] = pj
	}
	for i, eq := range equations{
		union(id[eq[0]], id[eq[1]], values[i])
	}
	fmt.Println(parent, w)
	precision := math.Pow10(5)
	ans := make([]float64, len(queries))
	for i, q := range queries{
		start, hasS := id[q[0]]
		end, hasE := id[q[1]]
		if hasS && hasE && find(start) == find(end) {
			ans[i] = math.Floor((w[start] / w[end])*precision + 0.5) / precision
		} else {
			ans[i] = -1
		}
	}
	return ans
}
/* 图BST解法：
   我们可以将整个问题建模成一张图：给定图中的一些点（变量），以及某些边的权值（两个变量的比值），
   试对任意两点（两个变量）求出其路径长（两个变量的比值）
*/
func CalcEquationBST(equations [][]string, values []float64, queries [][]string) []float64 {
	// 编号
	id := map[string]int{}
	for _, eq := range equations{
		for i := 0; i < 2; i++{
			if _, ok := id[eq[i]]; !ok {
				id[eq[i]] = len(id)
			}
		}
	}
	// 构键图 --- 邻接表实现
	type edge struct{
		to		int
		weight	float64
	}
	graph := make([][]edge, len(id))
	for i, eq := range equations{
		v, w := id[eq[0]], id[eq[1]]
		graph[v] = append(graph[v], edge{w, values[i]})
		graph[w] = append(graph[w], edge{v, 1/values[i]})
	}
	fmt.Println(graph)
	/* bfs 遍历
	 构建完图之后，对于任何一个查询，就可以从起点出发，通过广度优先搜索的方式，不断更新起点与当前点之间的路径长度，直到搜索到终点为止
	*/
	bfs := func(s, e int) float64{
		ratios := make([]float64, len(graph))
		ratios[s] = 1
		q := []int{s}
		for len(q) > 0{
			v := q[0]
			q = q[1:]
			if v == e{
				return ratios[v]
			}
			for _, item := range graph[v]{
				if w := item.to; ratios[w] == 0{
					ratios[w] = ratios[v] * item.weight
					q = append(q, w)
				}
			}
		}
		return -1
	}
	// 查询
	precision := math.Pow10(5)
	ans := make([]float64, len(queries))
	for i, q := range queries{
		start, hasS := id[q[0]]
		end, hasE := id[q[1]]
		if hasS && hasE {
			ans[i] = math.Floor(bfs(start, end)*precision + 0.5) / precision
		}else{
			ans[i] = -1
		}
	}
	return ans
}
/*
  对于查询数量很多的情形，如果为每次查询都独立搜索一次，则效率会变低。为此，我们不妨对图先做一定的预处理，随后就可以在较短的时间内回答每个查询。
  利用floyd算法预先算出任意两点间的距离
 */
func calcEquationFloyd(equations [][]string, values []float64, queries [][]string) []float64 {
	nodes := map[string]int{}
	for _, item := range equations{
		for _, e := range item{
			if _, ok := nodes[e]; !ok{
				nodes[e] = len(nodes)
			}
		}
	}
	size := len(nodes)
	// adj matrix
	graph := make([][]float64, size)
	for i := range graph{
		graph[i] = make([]float64, size)
	}
	for i, eq := range equations{
		n1, n2 := nodes[eq[0]], nodes[eq[1]]
		graph[n1][n2] = values[i]
		graph[n2][n1] = 1/values[i]
	}
	/* floyd alg: 复杂度为 N的立方, 空间负责度为 N 的平方
	DP: f[k][i][j] = min{ f[k-1][i][j], f[k-1][i][k]+f[k-1][k][j] }
	for k := 0; k < size; k++{
		for i := 0; i < size; i++{
			for j := 0; j < size; j++{
				graph[i][j] = min(graph[i][j], graph[i][k]+graph[k][j])
			}
		}
	}
	 */
	for k := range graph{
		for i := range graph{
			for j := range graph{
				if graph[i][k] > 0 && graph[k][j] > 0{
					graph[i][j] = graph[i][k]*graph[k][j]
				}
			}
		}
	}
	ans := make([]float64, len(queries))
	for i, q := range queries{
		from, fromOk := nodes[q[0]]
		to, toOk := nodes[q[1]]
		if toOk && fromOk && graph[from][to] != 0{
			ans[i] = graph[from][to]
		} else{
			ans[i] = -1
		}
	}
	return ans
}

/* 990. Satisfiability of Equality Equations
	可以将每个变量看作图中的一个节点，把相等关系 == 看作是连接两个节点的边， 由于表示相等关系的等式方程具有传递性
	即 若a==b 和 b==c 成立，则 a == c也成立。也即相等的变量属于同一个连通分量。
	看到连通分量，因此可以 union set来处理
	1. 遍历所有等式，构造并查集。同一个等式中的两个变量属于同一个连通分量
	2. 遍历所有的不等式。同一个不等式中的两个变量不能属于同一个连通分量，因此对两个变量分别查找其所在的连通分量，如果两个变量在同一个连通分量中，
       则产生矛盾，返回false
    3. 如果遍历完所有的不等式没有发现矛盾，则返回true
 */
func equationsPossible(equations []string) bool {
	parent := make([]int, 26)
	for i := 0; i < 26; i++{
		parent[i] = i
	}
	var union func(int, int)
	var find func(int) int
	for _, str := range equations{
		if str[1] == '='{
			index1 := int(str[0] - 'a')
			index2 := int(str[3] - 'a')
			union(index1, index2)
		}
	}

	for _, str := range equations {
		if str[1] == '!'{
			index1 := int(str[0] - 'a')
			index2 := int(str[3] - 'a')
			if find(index1) == find(index2) {
				return false
			}
		}
	}
	union = func(i, j int){
		parent[find(i)] = find(j)
	}
	find = func(i int) int{
		for parent[i] != i {
			parent[i] = parent[parent[i]] // 路径压缩
			i = parent[i]
		}
		return i
	}
	return true
}
/* 88. Merge Sorted Array

 */
func Merge(nums1 []int, m int, nums2 []int, n int)  {
	// two pointer
	i, j := 0, 0
	for i < m && j < n{
		if nums1[i] <= nums2[j]{
			i++
		}else{
			copy(nums1[i+1:], nums1[i:])
			nums1[i] = nums2[j]
			i++
			m++ // m 随之增加
			j++
		}
	}
	if i >= m && j < n{
		copy(nums1[i:], nums2[j:])
	}
}
/*
  从后往前 填充
 */
func Merge2(nums1 []int, m int, nums2 []int, n int)  {
	p1, p2, tail := m-1, n-1, m+n-1;
	for p1 >= 0 || p2 >= 0{ // 选择cur 填充到尾部
		cur := 0
		if p1 == -1{
			cur = nums2[p2]
			p2--
		}else if p2 == -1{
			cur = nums1[p1]
			p1--
		}else if nums1[p1] > nums2[p2]{
			cur = nums1[p1]
			p1--
		}else{
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
		tail--
	}
}

/* 453. Minimum Moves to Equal Array Elements
  Given an integer array nums of size n, return the minimum number of moves required to make all array elements equal.
  In one move, you can increment n - 1 elements of the array by 1.
Example 1:
	Input: nums = [1,2,3]
	Output: 3
	Explanation: Only three moves are needed (remember each move increments two elements):
	[1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
Example 2:
	Input: nums = [1,1,1]
	Output: 0
 */
/* 暴力： 关键点： 数组中的 最大值 与 最小值 相等
	O(k * n^2) k 为最大值与最小值的差
 */
func MinMoves(nums []int) int {
	n := len(nums)
	min, max, count := 0, n-1, 0
	for {
		for i := 0; i < n; i++ {
			if nums[max] < nums[i] {
				max = i
			}
			if nums[min] > nums[i] {
				min = i
			}
		}
		// 条件：最大值与最小值相等 即为条件
		if nums[max] == nums[min] {
			break
		}
		for i := 0; i < n; i++{
			if i != max{
				nums[i]++
			}
		}
		count++
	}
	return count
}
/* 改进：为了让最小元素等于最大元素，至少需要加 k 次，之后最大元素可能发生变化。因此可以一次性增加增量 k=max-min
	并将移动次数增加k，然后对整个数组遍历，找到最大值 最小值，重复这一过程 直至 最大值与最小值相等
	O(n^2)
 */
func MinMoves2(nums []int) int {
	n := len(nums)
	min, max, count := 0, n-1, 0
	for {
		for i := 0; i < n; i++{
			if nums[max] < nums[i]{
				max = i
			}
			if nums[min] > nums[i]{
				min = i
			}
		}
		diff := nums[max] - nums[min]
		if diff == 0{
			break
		}
		count += diff
		for i := 0; i < n; i++{
			if i != max{
				nums[i] += diff
			}
		}
	}
	return count
}
/*改进：排序 加速获得最大最小值
	用diff = max - min 更新数列
	1. 在每一步计算diff 之后正在更新有序数组的元素。如何在不遍历数组的情况下查询最大 最小值。在第一步中，最后的元素即为最大值。
	  因此 diff = a[n-1] - a[0] 我们对除最后一个元素以外增加diff
	2. 更新后的数组起始元素a'[0] 变成了 a[0]+diff = a[n-1] 因此 a'[0]变为上一步最大元素a[n-1]。由于数组有序，直到 i-2 的元素都满足
	  a[j] >= a[j-1]。故更新后 a'[n-2] 即为最大  a[0] 依然是最小元素
	3. 于是 在第二次更新时， diff = a[n-2] - a[0]  更新后， a''[0] 会成为 a'[n-2] 于是 最大元素为 a[n-3]
	4. 继续如此，每一步用最大最小值的差 更新数组
	5. 优化：不需要每次更新数组值，这是因为 即使是在更新元素之后，我们要登记的diff 差值也不变，因为max 和 min 增加的数字相同
	6. 于是，在排序数组后， moves = SUM(a[i] - a[0]) i从1到n-1
 */
func MinMoves3(nums []int) int {
	sort.Ints(nums)
	n, count := len(nums), 0
	for i := n-1; i > 0; i--{
		count += (nums[i] - nums[0])
	}
	return count
}
/* DP
考虑有序数组a 不考虑整个问题，而是分解问题
假设 直到 i-1 位置的元素已经相等， 我们只需考虑 i 位的元素，将差值diff=a[i]-a[i-1]加到总移动次数上，使得第 i 位也相等。
但当我们想要继续这一步时， a[i] 之后的元素也会被增加diff 亦即 a[j] += diff, 其中 j > i
但是在实现本方法时，不需要对这样的a[j]进行增加，相反 我们把moves 的数量增加到当前元素(a[i])中， a'[i] = a[i] + moves
对数组排序，一直更新moves以使得直到当前的元素相等，而不改变除了当前元素之外的元素。在整个数组扫描完毕后，moves即为答案。
 */
func MinMovesDP(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := 1; i < len(nums); i++{
		diff := (ans + nums[i]) - nums[i-1]
		nums[i] += ans
		ans += diff
	}
	return ans
}
/* 数学计算
将除了一个元素之外的全部元素+1，等价于将该元素-1，因为我们只对元素的相对大小感兴趣。因此，该问题简化为需要进行的减法次数
我们只需要将所有的数都减到最小的数即可
moves = SUM(a[i]) - min(a) * n ; n 为数组长度， i属于[0, n)
由于 SUM(a[i]) 可能非常大，造成整数越界，可以即时计算moves
moves = SUM(a[i] - min(a)) i属于[0, n)
 */
func MinMovesBest(nums []int) int {
	ans := 0
	min := math.MaxInt32
	for i := range nums{
		if min > nums[i]{
			min = nums[i]
		}
	}
	for i := range nums{
		ans += (nums[i] - min)
	}
	return ans
}

/* 414. Third Maximum Number
  Given integer array nums, return the third maximum number in this array.
  If the third maximum does not exist, return the maximum number.
 */
// nums中包含MinInt32的可能, 无法使用3指针处理，需要借助int64 避开MinInt32 判断
func ThirdMax(nums []int) int {
	f := math.MinInt32
	s,t := f,f
	ff,fs,ft := false, false, false // nums中包含MinInt32的可能
	for _,v := range nums{
		if f <= v{
			if f != v && ff{
				if fs{
					t = s
					ft = true
				}
				s = f
				fs = true
			}
			f = v
			ff = true
		}else if s <= v{
			if s != v && fs{
				t = s
				ft = true
			}
			s = v
			fs=true
		}else if t <= v{
			ft = true
			t = v
		}
	}
	if !ft {
		return f
	}
	return t
}

func ThirdMax2(nums []int) int {
	var f, s, t int64 = math.MinInt64, math.MinInt64, math.MinInt64
	for _, v := range nums{
		num := int64(v)
		if num > f{// 如果比第一个大
			num, f = f, num
		}
		if num < f && num > s{ // 比第一个小 但比第二个大
			num, s = s, num
		}
		if num < s && num > t{
			num, t = t, num
		}
	}
	if t != math.MinInt64{
		return int(t)
	}
	return int(f)
}

/* 645. Set Mismatch
You have a set of integers s, which originally contains all the numbers from 1 to n.
Unfortunately, due to some error, one of the numbers in s got duplicated to another number in the set, which results in repetition of one number and loss of another number.
You are given an integer array nums representing the data status of this set after the error.
Find the number that occurs twice and the number that is missing and return them in the form of an array.
 */
func FindErrorNums(nums []int) []int {
	n := len(nums)
	m := make([]int, n+1)
	ans := []int{}
	for i := range nums{
		if m[nums[i]] == 0{
			m[nums[i]] = 1
		}else{
			ans = append(ans, nums[i])
		}
	}
	for i := 1; i < n+1; i++{
		if m[i] == 0{
			ans = append(ans, i)
			break
		}
	}
	return ans
}
/* 寻找丢失的数字相对复杂，可能有以下两种情况：
  1. 如果丢失的数字大于 1 且小于 n，则一定存在相邻的两个元素的差等于 2，这两个元素之间的值即为丢失的数字
  2. 如果丢失的数字是 1 或者 n， 则需要额外判断
 */
func FindErrorNums2(nums []int) []int {
	ans := make([]int, 2)
	sort.Ints(nums)
	pre := 0
	for _, v := range nums{
		if v == pre{
			ans[0] = v
		}else if v-pre > 1{
			ans[1] = pre+1
		}
		pre = v
	}
	n := len(nums)
	if nums[n-1] != n{
		ans[1] = n
	}
	return ans
}
/* 位运算
知识点-1： 异或性质
知识点-2： 负数的异或运算
以 10 ^ -10 为例：
  0000 0000 0000 0000 0000 0000 0000 1010
^ 1111 1111 1111 1111 1111 1111 1111 0110
= 1111 1111 1111 1111 1111 1111 1111 1100

  0000 0000 0000 0000 0000 0000 0000 1010
& 1111 1111 1111 1111 1111 1111 1111 0110
= 0000 0000 0000 0000 0000 0000 0000 0010   lowbit 最低不同位
重复的数字在数组中出现 2 次，丢失的数字在数组中出现 0 次，其余的每个数字在数组中出现 1 次
重复的数字和丢失的数字的出现次数的奇偶性相同，且和其余的每个数字的出现次数的奇偶性不同。
如果在数组的 n 个数字后面再添加从 1 到 n 的每个数字，得到 2n 个数字，则在 2n 个数字中，重复的数字出现 3 次，丢失的数字出现 1 次，其余的每个数字出现 2 次。
根据出现次数的奇偶性，可以使用异或运算求解。
用 x 和 y 分别表示重复的数字和丢失的数字。异或运算满足交换律和结合律, a^a = 0   0 ^ a = a
xor = x^x^x^y = x^y
x与y 不同，故xor != 0
令lowbit=xor & (-xor)， 则lowbit为x和y的二进制表示中的最低不同位，可用lowbit区分x和y
得到lowbit 后， 可以将上述2n个数分成2组，
第一组的每个数字a 都满足 a % lowbit = 0
第二组的每个数组b 都满足 b & lowbit != 0
创建两个变量 num1 num2 初始值为0  再次遍历上述2n个数字，对于每个数字a 如果 a&lowbit==0,则另num1 = num1 ^ a， 否则 num2=num2^a
遍历结束后，num1为第一组的全部数字的异或结果， num2为第二组全部数字异或结果。
因为同一个数字只能出现在其中的一组，且除了x和y外，每个数字一定在其中的一组出现2次，因此 num1 和 num2 分别对应 x 和 y中的一个数字。
为了知道num1 和 num2 与 x和y 对应关系。需要再次遍历数组nums即可。
如果数组中存在元素等于num1，则x==num1 y == num2 反之 x = num2  y== num1
 */
func FindErrorNumsBit(nums []int) []int {
	xor := 0
	for _, v := range nums{
		xor ^= v
	}
	n := len(nums)
	for i := 1; i <= n; i++{
		xor ^= i
	}
	lowbit := xor & (-xor)
	num1, num2 := 0, 0
	for _, v := range nums{
		if v & lowbit == 0{
			num1 ^= v
		}else {
			num2 ^= v
		}
	}
	for i := 1; i <= n; i++{
		if i & lowbit == 0{
			num1 ^= i
		}else {
			num2 ^= i
		}
	}
	for _, v := range nums{
		if v == num1{
			return []int{num1, num2}
		}
	}
	return []int{num2, num1}
}
/* 因为值的范围在[1,n]，我们可以运用「桶排序」的思路，根据 nums[i] = i + 1的对应关系使用 O(n) 的复杂度将每个数放在其应该落在的位置里。
   然后线性扫描一遍排好序的数组，找到不符合 nums[i] = i + 1对应关系的位置，从而确定重复元素和缺失元素是哪个值。
 */
func FindErrorNumsSwap(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++{
		for nums[i] != i+1 && nums[nums[i] - 1] != nums[i]{ // 保证每次迭代，都有一个元素放入正确位置
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	a, b := -1, -1
	for i := 0; i < n; i++{
		if nums[i] != i+1{
			a = nums[i]
			if i == 0{ // 避免i-1情况
				b = 1
			}else{
				b = nums[i-1]+1
			}
		}
	}
	return []int{a,b}
}

/* 153. Find Minimum in Rotated Sorted Array[寻找旋转排序数组中的最小值]
** Suppose an array of length n sorted in ascending order is rotated between 1 and n times.
** For example, the array nums = [0,1,2,4,5,6,7] might become:
	[4,5,6,7,0,1,2] if it was rotated 4 times.
	[0,1,2,4,5,6,7] if it was rotated 7 times.
** Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].
** Given the sorted rotated array nums of unique elements, return the minimum element of this array.
** You must write an algorithm that runs in O(log n) time.
 */
func FindMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high{
		mid := (low^high)>>1 + (low&high)
		if nums[mid] < nums[high]{
			high = mid
		}else{
			low = mid+1
		}
	}
	return nums[low]
}
/* 154. Find Minimum in Rotated Sorted Array II[寻找旋转排序数组中的最小值 II]
** Suppose an array of length n sorted in ascending order is rotated between 1 and n times.
** For example, the array nums = [0,1,4,4,5,6,7] might become:
	[4,5,6,7,0,1,4] if it was rotated 4 times.
	[0,1,4,4,5,6,7] if it was rotated 7 times.
Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].
Given the sorted rotated array nums that may contain duplicates, return the minimum element of this array.
You must decrease the overall operation steps as much as possible.
 */
/* 区别在于 有重复元素
** 考虑数组中的最后一个元素 x ：在最小值右侧的元素，它们的值一定都小于等于 x ， 而在最小值左侧的元素，他们的值一定都大于等于x。
** 依据上述性质，可以通过二分查找找出最小值。
** 一共分3种情况：
** 1. nums[mid] > nums[high]
** 2. nums[mid] < nums[high]
** 3. nums[mid] == nums[high]
** 由于重复元素的存在，我们并不能确定nums[mid]究竟在最小值的左侧还是右侧，因此我们不能莽撞地忽略某一部分的元素
** 唯一可以知道的是，由于它们的值相同，所以无论nums[high]是不是最小值，都有一个它的「替代品」nums[mid],因此依然可以继续缩
** 即忽略二分的右端点。
 */
func FindMinII(nums []int) int {
	n := len(nums)
	low, high := 0, n-1
	//是有序单调数组
	if nums[low] < nums[high]{
		return nums[low]
	}
	for low < high{
		mid := (low^high)>>1 + low&high
		if nums[mid] > nums[high]{
			low = mid+1
		}else if nums[mid] < nums[high]{
			high = mid
		}else{
			high--
		}
	}
	return nums[low]
}
func FindMinIILeft(nums []int) int {
	n := len(nums)
	low, high := 0, n-1
	//是有序单调数组
	if nums[low] < nums[high]{
		return nums[low]
	}
	for low < high{
		//如果二分后的数组是有序数组，则返回最左元素，即为最小
		if nums[low] < nums[high]{
			return nums[low]
		}
		mid := (low^high)>>1 + low&high
		//若最左小于mid元素，则最左到mid是严格递增的，那么最小元素必定在mid之后
		if nums[mid] > nums[low]{
			low = mid+1
		}else if nums[mid] < nums[high]{
			high = mid
		}else{
			high--
		}
	}
	return nums[low]
}

/*Offer 51. 数组中的逆序对  LCOF
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
示例 1:
输入: [7,5,6,4]
输出: 5
 */
/* 归并排序 */
func ReversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}
func mergeSort(nums []int, start, end int)int{
	if start >= end{
		return 0
	}
	mid := (start ^ end)>> 1 + start & end
	cnt := mergeSort(nums, start, mid) + mergeSort(nums, mid+1, end)
	tmp := []int{}
	i, j := start, mid+1 // 避开了 mid
	for i <= mid && j <= end{
		if nums[i] <= nums[j]{
			tmp = append(tmp, nums[j])
			cnt += j - (mid + 1)
			i++
		}else{
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for ;i <= mid;i++{
		tmp = append(tmp, nums[i])
		cnt += end - (mid+1) + 1
	}
	for ; j <= end; j++{
		tmp = append(tmp, nums[j])
	}
	for i := start; i <= end; i++{
		nums[i] = tmp[i-start]
	}
	return cnt
}
/* 315. Count of Smaller Numbers After Self
** You are given an integer array nums and you have to return a new counts array.
** The counts array has the property where counts[i] is the number of smaller elements to the right of nums[i].
*/
/*方法一：二分查找 复杂度：O(n(n+logn))*/
func CountSmaller(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}
	ans := make([]int, n)
	sorted := []int{}
	search := func(target int)int{// 返回插入位置
		loc := len(sorted)
		i, j := 0, len(sorted) - 1
		for i <= j{
			mid := (i^j)>>1 + i&j
			if sorted[mid] < target{
				i = mid + 1
			}else{
				loc = mid
				j = mid - 1
			}
		}
		return loc
	}
	for i := n-1; i >= 0; i--{
		index := search(nums[i])
		// 切片中间插入元素，append方法效率低，提交超时
		//sorted = append(sorted[:index], append([]int{nums[i]}, sorted[index:]...)...)
		// 下面这个方式，效率比上面要高
		sorted = append(sorted, 0) // 扩充空间
		copy(sorted[index+1:], sorted[index:])
		sorted[index] = nums[i]
		ans[i] = index
	}
	return ans
}

/* 1539. Kth Missing Positive Number
** Given an array arr of positive integers sorted in a strictly increasing order, and an integer k.
** Find the kth positive integer that is missing from this array.
** 如果数组是无序的，此方法是已知最优的
 */
func FindKthPositive(arr []int, k int) int {
	m := map[int]bool{}
	for i := range arr{
		m[arr[i]] = true
	}
	ans := 0
	for j := 1; k > 0; j++{
		if !m[j]{
			k--
			ans = j
		}
	}
	return ans
}
/* 优化-1: 时间复杂度O(n+k),空间O(n)
** 如何把 map 给取消掉，发现有个特性没有用到，即数组是升序的
** 因此可借助此特性处理
 */
func FindKthPositiveIter(arr []int, k int) int {
	var ans, cnt, i int
	n := len(arr)
	for j := 1; cnt < k; j++{
		if i < n && arr[i] == j{
			i++
		}else{
			ans = j
			cnt++
		}
	}
	return ans
}

/*优化-2 Binary Search
** 利用arr[i]与其下标i关系
** 一个不缺失元素的序列，会有arr[i]=i+1这种关系，而在缺失元素之后，会有arr[i]>i+1 可转换为==> arr[i]-i-1 > 0
** 缺失一个的时候，相差 1， 两个则相差2，依次类推，缺失越多，两者差距越大，要找第 K 个缺失的，换言之，只要 arr[i]-i-1 == k 这便是要找的数子
** arr[i]-i-1 == k ==> arr[i] - i == k+1  由于数组arr中并不一定存在k+1,故 找第一个大于等于的
 */
func FindKthPositiveBinarySearch(arr []int, k int) int {
	left, right := 0, len(arr)
	for left < right{
		mid := (left ^ right)>>1 + left&right
		if arr[mid] - mid >= k+1{
			right = mid
		}else{
			left = mid + 1
		}
	}
	return k  + left
}

/* 273. Integer to English Words
** Convert a non-negative integer num to its English words representation.
*/
var(  // 注意对 0 的设置
	singles	= []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	teens	= []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens	= []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	thousands	= []string{"", "Thousand", "Million", "Billion"}
)
func NumberToWords(num int) string {
	if num == 0{
		return "Zero"
	}
	sb := strings.Builder{}
	toEng := func(num int){
		if num >= 100{
			sb.WriteString(singles[num/100])
			sb.WriteString(" Hundred ")
			num %= 100
		}
		if num >= 20{
			sb.WriteString(tens[num/10]+" ")
			num %= 10
		}
		if 0 < num && num < 10{
			sb.WriteString(singles[num]+" ")
		}else if num >= 10{
			sb.WriteString(teens[num-10]+" ")
		}
	}
	// 迭代
	for i, unit := 3, int(1e9); i >= 0; i--{
		if curNum := num / unit; curNum > 0{
			num -= curNum * unit
			toEng(curNum) // 计算3位一组
			sb.WriteString(thousands[i])
			sb.WriteByte(' ')
		}
		unit /= 1000
	}
	return strings.TrimSpace(sb.String())
}
/*自己实现*/
func numberToWords(num int) string {
	if num == 0{ // 易漏点-1 特殊情况
		return "Zero"
	}
	toEng := func(num int)string{
		t := []string{}
		if num >= 100{
			t = append(t, singles[num/100], "Hundred")
			num %= 100
		}
		if num >= 20{
			 t = append(t, tens[num/10])
			 num %= 10
		}
		// 排除 数字为 0 情况
		if num > 0 && num < 10{
			t = append(t, singles[num])
		}else if num >= 10{
			t = append(t, teens[num-10])
		}
		return strings.Join(t, " ")
	}
	unit := 0
	ans := []string{}
	for num > 0{
		n := num % 1000
		if n > 0{ // 易错点-2 针对 1000000 情况， 否则出现 one Million Thousand
			ans = append([]string{toEng(n), thousands[unit]}, ans...)
		}
		unit++
		num /= 1000
	}
	return strings.Trim(strings.Join(ans, " "), " ")
}

func NumberToWordsReCurive(num int) string {
	if num == 0{ // 易漏点-1 特殊情况
		return "Zero"
	}
	sb := strings.Builder{}
	var toEng func(int)
	toEng = func(num int){
		switch {
		case num == 0:
		case num < 10:
			sb.WriteString(singles[num]+" ")
		case num < 20:
			sb.WriteString(teens[num-10]+" ")
		case num < 100:
			sb.WriteString(tens[num/10] + " ")
			toEng(num%10)
		default:
			sb.WriteString(singles[num/100] + "Hundred ")
			toEng(num % 100)
		}
	}
	for i, unit := 3, int(1e9); i >= 0; i--{
		if curNum := num / unit; curNum > 0{
			num -= curNum * unit
			toEng(curNum)
			sb.WriteString(thousands[i])
			sb.WriteByte(' ')
		}
		unit /= 1000
	}
	return strings.TrimSpace(sb.String())
}
/* 1863. Sum of All Subset XOR Totals
The XOR total of an array is defined as the bitwise XOR of all its elements, or 0 if the array is empty.
For example, the XOR total of the array [2,5,6] is 2 XOR 5 XOR 6 = 1.
Given an array nums, return the sum of all XOR totals for every subset of nums. 
Note: Subsets with the same elements should be counted multiple times.
An array a is a subset of an array b if a can be obtained from b by deleting some (possibly zero) elements of b.
Example 1:
	Input: nums = [1,3]
	Output: 6
	Explanation: The 4 subsets of [1,3] are:
	- The empty subset has an XOR total of 0.
	- [1] has an XOR total of 1.
	- [3] has an XOR total of 3.
	- [1,3] has an XOR total of 1 XOR 3 = 2.
	0 + 1 + 3 + 2 = 6
 */
func SubsetXORSum(nums []int) int {
	ans := 0
	n := len(nums)
	var dfs func(val, idx int)
	dfs = func(val, idx int){
		if idx == n {
			ans += val
			return
		}
		dfs(val ^ nums[idx], idx+1) // 选idx
		dfs(val, idx+1)				// 不选idx
	}
	dfs(0,0)
	return ans
}


/* 1470. Shuffle the Array
** Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].
** Return the array in the form [x1,y1,x2,y2,...,xn,yn].
*/
/*空间复杂度O(1)解法
** 题目限制了每一个元素 nums[i] 最大只有可能是 1000，这就意味着每一个元素只占据了 10 个 bit。（2^10 - 1 = 1023 > 1000）
** 而一个int 有 32 bit，所以我们还可以使用剩下的 22 个 bit 做存储。实际上，每个 int，我们再借 10 个 bit 用就好了
** 每一个 nums[i] 的最低的十个 bit（0-9 位），我们用来存储原来 nums[i] 的数字；
** 再往前的十个 bit（10-19 位），我们用来存储重新排列后正确的数字是什么
*/
func Shuffle(nums []int, n int) []int {
	for i := 0; i < 2*n; i++{
		// 首先计算 nums[i] 对应的重新排列后的索引 j
		j := 2*i
		if i >= n{
			j = 2 * (i-n) + 1
		}
		//取 nums[i] 的低 10 位（nums[i] & 1023），即 nums[i] 的原始信息，把他放到 nums[j] 的高十位上
		nums[j] |= (nums[i] & 1023) << 10
	}
	//每个元素都取高 10 位的信息
	for i := range nums{
		nums[i] = nums[i]>>10
	}
	return nums
}
/* 空间复杂度O(1)解法二
** 题目中限制每一个元素 nums[i] 都大于 0。我们可以使用负数做标记, 标记当前 nums[i] 存储的数字，是不是重新排列后的正确数字
** 如果是，存负数；如果不是，存正数（即原本的数字，还需处理）
** 每次处理一个nums[i] 计算这个 nums[i] 应该放置的正确位置 j。
** 但是，nums[j] 还没有排列好，所以我们暂时把 nums[j] 放到 nums[i] 的位置上来，并且记录上，此时 nums[i] 的元素本来的索引是 j。
** 现在，我们就可以安心地把 nums[i] 放到 j 的位置了。同时，因为这已经是 nums[i] 正确的位置，取负数，即标记这个位置已经存放了正确的元素
** 之后，我们继续处理当前的 nums[i]，注意，此时这个新的 nums[i]，本来的索引是 j。
** 所以我们根据 j 算出它应该存放的位置，然后把这个位置的元素放到 nums[i] 中，取负做标记
** 这个过程以此类推。这就是代码中 while 循环做的事情。
** 直到 nums[i] 的值也是负数，说明 i 的位置也已经是重新排列后的正确元素了，我们就可以看下一个位置了。
 */
func Shuffle2(nums []int, n int) []int {
	for i := range nums{
		// 在 for 循环中，如果某一个元素已经是小于零了，说明这个位置已经是正确元素了，可以忽略
		if nums[i] > 0{
			// j 描述当前的 nums[i] 对应的索引，初始为 i
			j := i
			// 计算 j 索引的元素，也就是现在的 nums[i]，应该放置的索引
			for nums[i] > 0{
				if j < n{
					j = 2*j
				}else{
					j = 2*(j-n)+1
				}
				// 把 nums[i] 放置到 j 的位置，
				// 同时，把 nums[j] 放到 i 的位置，在下一轮循环继续处理
				nums[i], nums[j] = nums[j], nums[i]
				// 使用负号标记上，现在 j 位置存储的元素已经是正确的元素了
				nums[j] = -nums[j]
			}
		}
	}
	for i := range nums{
		nums[i] = -nums[i]
	}
	return nums
}
/*1528. Shuffle String
Given a string s and an integer array indices of the same length.
The string s will be shuffled such that the character at the ith position moves to indices[i] in the shuffled string.
Return the shuffled string.
原地修改
*/
func RestoreString(s string, indices []int) string {
	ans := []byte(s)
	for i, c := range ans{
		if indices[i] != i{ // 避免处理重复的封闭路径
			idx := indices[i] // 当前字符需要被移动的目标位置
			for idx != i{
				ans[idx], c = c, ans[idx] // 在覆写 s[idx] 之前，先将其原始值赋给变量 c
				indices[idx], idx = idx, indices[idx] // 将封闭路径中的 indices 数组的值设置成下标自身
			}
			ans[i] = c
			indices[i] = i // 每处理一个封闭路径，就将该路径上的indices 数组的值设置成下标自身
		}
	}
	return string(ans)
}