package array

import (
	"fmt"
	"sort"
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