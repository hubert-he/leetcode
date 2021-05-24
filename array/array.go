package array

import (
	"fmt"
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









