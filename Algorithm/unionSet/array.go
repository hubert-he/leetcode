package unionSet

import "fmt"

/* Union Find 并查集
	0. 设定相距为1的值，为一个集合
    1. 用map记录 数字和对应的索引
    2. 用例有重复数字，map处理
    3. 连通数字和索引对应关系
    4. 遍历并查集数组，集合最大即为最长连续子序列
*/
/* 128. Longest Consecutive Sequence
** Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
** You must write an algorithm that runs in O(n) time.
 */
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

// 2022-03-25 重刷此题，并查集 错误,
// 将下标作为集合元素，无法处理重复元素， 应该选用 nums 元素的值
func longestConsecutive_error(nums []int) int {
	n := len(nums)
	ufs,cnt := make([]int, n), make([]int, n)
	ans := 0
	m, vis := map[int]int{}, map[int]bool{}
	for i := range ufs{
		ufs[i] = i
		m[nums[i]] = i
		cnt[i] = 1
	}
	/*
	   find := func(x int)int{
	       px := ufs[x]
	       for px != x{
	           ufs[x] = ufs[px]
	           x = px
	           px = ufs[x]
	       }
	       return px
	   }*/
	find := func(x int)int{
		for x != ufs[x]{
			ufs[x] = ufs[ufs[x]]
			x = ufs[x]
		}
		return x
	}
	union := func(x, y int){
		px, py := find(x), find(y)
		ufs[px] = py
		if px != py {
			cnt[py] += cnt[px]
			if ans < cnt[py]{
				ans = cnt[py]
			}
		}
	}
	for i := range nums{
		if vis[i] { continue }
		vis[i] = true
		for _, r := range []int{1,-1}{
			if idx, ok := m[nums[i]+r]; ok {
				union(i, idx)
			}
		}
	}
	fmt.Println(ufs)
	return ans
}

/*  leetcode官方解法
** 考虑数组中每个数x,考虑以其为起点，不断尝试匹配 x+1 x+2 ... 是否存在，假设匹配到了x+y，那么以x起点的最长连续序列即为 x x+1 ... x+y 长度为
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
