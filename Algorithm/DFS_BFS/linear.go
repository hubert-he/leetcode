package DFS_BFS

import (
	"fmt"
	"math"
)

/* 127. Word Ladder
** A transformation sequence from word beginWord to word endWord using a dictionary wordList is a sequence of words beginWord -> s1 -> s2 -> ... -> sk
** such that:
 	Every adjacent pair of words differs by a single letter.
	Every si for 1 <= i <= k is in wordList. Note that beginWord does not need to be in wordList.
	sk == endWord
** Given two words, beginWord and endWord, and a dictionary wordList,
** return the number of words in the shortest transformation sequence from beginWord to endWord, or 0 if no such sequence exists.

本题 需要学习的地方有 2 个：
1.  构建图的方式
2.  双向BFS的思路

同类题目：
	126. Word Ladder II
	433. Minimum Genetic Mutation
	752. Open the Lock
 */
/* 思路一： 广度优先搜索 + 优化建图 */
/* 本题要求的是最短转换序列的长度，看到最短首先想到的就是广度优先搜索，
** 提到广度优先搜索自然而然的就能想到图，但是本题并没有直截了当的给出图的模型，因此我们需要把它抽象成图的模型
** 可以把每个单词都抽象为一个点，如果两个单词可以只改变一个字母进行转换，那么说明他们之间有一条双向边。
** 因此我们只需要把满足转换条件的点相连，就形成了一张图
** 基于该图，我们以 beginWord 为图的起点，以 endWord 为终点进行广度优先搜索，寻找 beginWord 到 endWord 的最短路径。
*/
/* 构建图
** 1. 为了方便表示，我们先给每一个单词标号，即给每个单词分配一个 id。
**    创建一个由单词 word 到 id 对应的映射 wordId，并将 beginWord 与 wordList 中所有的单词都加入这个映射中。
**    之后我们检查 endWord 是否在该映射内，若不存在，则输入无解。我们可以使用哈希表实现上面的映射关系
** 2. 然后我们需要建图，依据朴素的思路，我们可以枚举每一对单词的组合，判断它们是否恰好相差一个字符，以判断这两个单词对应的节点是否能够相连。
**    但是这样效率太低，我们可以优化建图
** 3. 具体地，我们可以创建虚拟节点。对于单词 hit，我们创建三个虚拟节点 *it、h*t、hi*，并让 hit 向这三个虚拟节点分别连一条边即可。
**	  如果一个单词能够转化为 hit，那么该单词必然会连接到这三个虚拟节点之一。
**    对于每一个单词，我们枚举它连接到的虚拟节点，把该单词对应的 id 与这些虚拟节点对应的 id 相连即可
** 4. 最后我们将起点加入队列开始广度优先搜索，当搜索到终点时，我们就找到了最短路径的长度。
**    注意因为添加了虚拟节点，所以我们得到的距离为实际最短路径长度的两倍。
** 	  同时我们并未计算起点对答案的贡献，所以我们应当返回距离的一半再加一的结果
*/
func LadderLength_BFS(beginWord string, endWord string, wordList []string) int {
	wordId := map[string]int{}
	graph := [][]int{}
	addWord := func(word string)int{
		id, has := wordId[word]
		if !has {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	// 技巧： 虚拟中间节点的添加
	addEdge := func(word string)int{
		id1 := addWord(word)
		s := []byte(word)
		for i, b := range s {
			s[i] = '*'
			id2 := addWord(string(s))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			s[i] = b
		}
		return id1
	}
	for _, word := range wordList{
		addEdge(word)
	}
	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0
	}
	// 开始DFS 求解最短路径
	dist := make([]int, len(wordId))
	for i := range dist{
		dist[i] = math.MaxInt32
	}
	dist[beginId] = 0
	q := []int{beginId}
	for len(q) > 0{
		v := q[0]
		q = q[1:]
		if v == endId{
			return dist[endId]/2 + 1
		}
		for _, w := range graph[v] {
			if dist[w] == math.MaxInt32{
				dist[w] = dist[v] + 1
				q = append(q, w)
			}
		}
	}
	return 0
}

/** 思路2： 双向广度优先搜索来降低搜索空间
** 根据给定字典构造的图可能会很大，而广度优先搜索的搜索空间大小依赖于每层节点的分支数量
** 假如每个节点的分支数量相同，搜索空间会随着层数的增长指数级的增加。
** 考虑一个简单的二叉树，每一层都是满二叉树的扩展，节点的数量会以 2 为底数呈指数增长
** 所以，如果数据范围太大beginWord.length >= 10 ,朴素BFS 「搜索空间爆炸」
** 如果我们的 wordList 足够丰富（包含了所有单词），
** 对于一个长度为 10 的 beginWord​ 替换一次字符可以产生 10 * 25 个新单词（每个替换点可以替换另外 25 个小写字母），
** 第一层就会产生 250 个单词；第二层会产生超过 6 * 10^4 个新单词 ...
** 随着层数的加深，这个数字的增速越快，这就是「搜索空间爆炸」问题。
** 在朴素的 BFS 实现中，空间的瓶颈主要取决于搜索空间中的最大宽度
** 如果使用两个同时进行的广搜可以有效地减少搜索空间。
** 一边从 beginWord 开始，另一边从 endWord 开始。我们每次从两边各扩展一层节点，当发现某一时刻两边都访问过同一顶点时就停止搜索。
** 这就是双向广度优先搜索，它可以可观地减少搜索空间大小，从而提高代码运行效率
** 「双向 BFS」的基本实现思路如下：
	1. 创建「两个队列」分别用于两个方向的搜索；
	2. 创建「两个哈希表」用于「解决相同节点重复搜索」和「记录转换次数」；
	3. 为了尽可能让两个搜索方向“平均”，每次从队列中取值进行扩展时，先判断哪个队列容量较少；
	4. 如果在搜索过程中「搜索到对方搜索过的节点」，说明找到了最短路径。
伪代码：
	d1、d2 为两个方向的队列
	m1、m2 为两个方向的哈希表，记录每个节点距离起点的
	// 只有两个队列都不空，才有必要继续往下搜索
	// 如果其中一个队列空了，说明从某个方向搜到底都搜不到该方向的目标节点
	while(!d1.isEmpty() && !d2.isEmpty()) {
		if (d1.size() < d2.size()) {
			update(d1, m1, m2);
		} else {
			update(d2, m2, m1);
		}
	}
	// update 为从队列 d 中取出一个元素进行「一次完整扩展」的逻辑
	void update(Deque d, Map cur, Map other) {}
*/
func LadderLength_BiBFS(beginWord string, endWord string, wordList []string) int {
	m := map[string]bool{}
	for i := range wordList{
		m[wordList[i]] = true
	}
	if _, ok := m[endWord]; !ok{
		return 0
	}
	// update 代表从 queue 中取出一个单词进行扩展，
	// cur 为当前方向的距离字典；other 为另外一个方向的距离字典
	// 注意slice 作为参数，需要指针类型， 否则不更新
	update := func(queue *[]string, cur map[string]int, other map[string]int)int{
		q := *queue
		word := q[0]
		q = q[1:]
		sb := []byte(word)
		for i := range sb{
			origin := sb[i]
			var c byte
			for c = 'a'; c <= 'z'; c++{
				if c == origin{
					continue
				}
				sb[i] = c
				newString := string(sb)
				if m[newString]{
					// 如果该字符串在「当前方向」被记录过（拓展过），跳过即可
					if _, ok := cur[newString]; ok {
						continue
					}
					// 如果该字符串在「另一方向」出现过，说明找到了联通两个方向的最短路
					if _, ok := other[newString]; ok {
						//fmt.Println(newString, cur[word], cur[newString], other[newString])
						//return cur[newString] + other[newString] + 1  注意，只有一侧看到了, 即cur 第一次看到，other 早已看到过
						return cur[word] + other[newString] + 1
					}else{
						// 加入q 继续寻找
						q = append(q, newString)
						cur[newString] = cur[word] + 1
					}
				}
			}
			sb[i] = origin
		}
		return -1
	}
	// 定义2个队列
	begQueue, endQueue := []string{beginWord}, []string{endWord}
	// mBeg 和 mEnd 分别记录两个方向出现的单词是经过多少次转换而来
	mBeg, mEnd := map[string]int{beginWord: 0}, map[string]int{endWord: 0}
	/* 只有两个队列都不空，才有必要继续往下搜索
	 * 如果其中一个队列空了，说明从某个方向搜到底都搜不到该方向的目标节点
	 * 例如，如果 d1 为空了，说明从 beginWord 搜索到底都搜索不到 endWord，反向搜索也没必要进行了
	 */
	for len(begQueue) > 0 && len(endQueue) > 0{
		ret := -1
		// 为了让两个方向的搜索尽可能平均，优先拓展队列内元素少的方向
		if len(begQueue) <= len(endQueue){
			ret = update(&begQueue, mBeg, mEnd)
		}else{
			ret = update(&endQueue, mEnd, mBeg)
		}
		if ret != -1{
			return ret + 1
		}
	}
	return 0
}

/* 126. Word Ladder II
** A transformation sequence from word beginWord to word endWord using a dictionary wordList is a sequence of words beginWord -> s1 -> s2 -> ... -> sk such that:
	Every adjacent pair of words differs by a single letter.
	Every si for 1 <= i <= k is in wordList. Note that beginWord does not need to be in wordList.
	sk == endWord
** Given two words, beginWord and endWord, and a dictionary wordList,
** return all the shortest transformation sequences from beginWord to endWord, or an empty list if no such sequence exists.
** Each sequence should be returned as a list of the words [beginWord, s1, s2, ..., sk].
*/
/* 用DFS的方法来处理： 分析出递归树
** 递归树：即从起始单词开始，与之相差一个字母的单词作为它的子节点进行遍历，直至找到结束单词。到这里输出所有可能单词串，用DFS 是比较方便的
** 题目额外的一个条件是输出最短的，也即需要找到最短的所有可能，求最短的即 BFS 最为方便。如何结合两者是一个问题点
** DFS BFS 都有一个共同的问题需要解决， 即 如何找到一个节点的所有孩子节点
** 方法一：遍历wordList来判断每个单词和当前单词是否只有一个字母不同，复杂度O(mn) 平方级
** 方法二：将要找的节点单词的每个位置换一个字符，然后看更改后的单词在不在 wordList 中
** 最后一个问题，即查找最短路径，如果找到的新的路径的长度比之前的路径短，就把之前的结果清空，重新找，如果是最小的长度，就加入到结果中
** 这个保存结果问题，在dfs 里 很好处理，但是BFS 处理需要些方法
*/
/*下面这个函数是DFS 的标准实现
** 1. 务必持续训练此DFS方法
** 2. 务必注意Golang 传参切片，发生的拷贝是浅拷贝
*/
func FindLadders_DFS(beginWord string, endWord string, wordList []string) [][]string {
	m := map[string]bool{}
	for i := range wordList{
		m[wordList[i]] = true
	}
	ans := [][]string{}
	if _, ok := m[endWord]; !ok{
		return ans
	}
	wordLen := len(beginWord)
	//minPath := math.MaxInt32 // 记录当前最短路径
	minPath := len(wordList) + 1 // 路径的最大可能长度
	getNext := func(node string)(res []string){
		if len(node) != wordLen{
			return
		}
		cur := []byte(node)
		for i := range cur{
			origin := cur[i]
			var c byte
			for c = 'a'; c <= 'z'; c++{
				if c == origin{
					continue
				}
				cur[i] = c
				if _, ok := m[string(cur)]; ok {
					res = append(res, string(cur))
				}
			}
			cur[i] = origin
		}
		return
	}
	var dfs func(node string, path []string)
	dfs = func(node string, path []string){
		path = append(path, node)
		dist := len(path)
		if node == endWord{
			// deep copy： 务必注意Golang中Slice作为参数的浅拷贝情况
			t := make([]string, dist)
			copy(t, path)
			if dist == minPath{
				ans = append(ans, t)
			}else if dist < minPath{
				ans = append([][]string{}, t)
				minPath = dist
			}
			return
		}
		if dist >= minPath{
			return
		}
		for _, word := range getNext(node){
			dfs(word, path)
		}
	}
	dfs(beginWord, []string{})
	return ans
}

/*DFS 进行优化
** 优化点-1：路径中已经含有当前单词，如果再把当前单词加到路径，那肯定会使得路径更长，所以跳过
** 优化点-2：利用BFS提前获得 最短长度
 */
func FindLadders_DFS_Fine(beginWord string, endWord string, wordList []string) [][]string {
	m, visited := map[string]bool{}, map[string]bool{}
	for i := range wordList{
		m[wordList[i]] = true
		visited[wordList[i]] = false
	}
	ans := [][]string{}
	if _, ok := m[endWord]; !ok{
		return ans
	}
	wordLen := len(beginWord)
	minPath := 1 // 路径的最大可能长度
	getNext := func(node string)(res []string){
		if len(node) != wordLen{
			return
		}
		cur := []byte(node)
		for i := range cur{
			origin := cur[i]
			var c byte
			for c = 'a'; c <= 'z'; c++{
				if c == origin{
					continue
				}
				cur[i] = c
				s := string(cur)
				if _, ok := m[s]; ok {
					res = append(res, string(cur))
					//优化点-1： 路径中已经含有当前单词，如果再把当前单词加到路径，那肯定会使得路径更长，所以跳过
					/*
					if !visited[s]{
						res = append(res, string(cur))
					} */
				}
			}
			cur[i] = origin
		}
		return
	}
	graph := map[string][]string{} // 构建邻接表
	bfs := func(){
		queue := []string{beginWord}
		found := false
		for len(queue) > 0{
			minPath++
			q := queue
			queue = nil
			subVisited := map[string]bool{}
			for i := range q{
				children := getNext(q[i])
				//graph[q[i]] = children
				t := []string{} // 易错点-2
				for _, word := range children{
					if !visited[word] {
						// visited[word] = true  易错点-1： 不能直接置为true，导致处在同一层的部分节点丢掉
						subVisited[word] = true
						if word == endWord{
							found = true
						}
						queue = append(queue, word)
						// graph[q[i]] = append(graph[q[i]], word) 易错点-2 可能会出现重复的项
						if _, ok := graph[q[i]]; !ok {
							t = append(t, word)
						}
					}
				}
				graph[q[i]] = append(graph[q[i]], t...) // 易错点-2
			}
			for word := range subVisited{
				visited[word] = true
			}
			if found{
				break
			}
		}
	}
	var dfs func(node string, path []string)
	dfs = func(node string, path []string){
		path = append(path, node)
		dist := len(path)
		if node == endWord{
			// deep copy： 务必注意Golang中Slice作为参数的浅拷贝情况
			t := make([]string, dist)
			copy(t, path)
			if dist == minPath{
				ans = append(ans, t)
			}else if dist < minPath{
				ans = append([][]string{}, t)
				minPath = dist
			}
			return
		}
		if dist >= minPath{
			return
		}
		for _, word := range graph[node] {
			dfs(word, path)
		}
		return
	}
	if _, ok := visited[beginWord]; ok {
		visited[beginWord] = true
	}
	bfs()
	fmt.Println(minPath, graph)
	dfs(beginWord, []string{})
	return ans
}

/* BFS 方式求解
** 注意两点：
** 1. slice 浅拷贝的问题
** 2. 学习 BFS 方式 如何保存路径值：即放在Queue中
 */
func findLadders_BFS(beginWord string, endWord string, wordList []string) [][]string {
	ans := [][]string{}
	m := map[string]bool{}
	for i := range wordList{
		m[wordList[i]] = true
	}
	if _, ok := m[endWord]; !ok {
		return ans
	}
	getNext := func(word string)[]string{
		sb := []byte(word)
		res := []string{}
		for i := range sb{
			origin := sb[i]
			var c byte = 'a'
			for c <= 'z'{
				if c == origin{
					c++
					continue
				}
				sb[i] = c
				if m[string(sb)]{
					res = append(res, string(sb))
				}
				c++
			}
			sb[i] = origin
		}
		return res
	}
	//BFS 的队列不用存储 String，直接去存到目前为止的路径
	//queue := []string{beginWord}
	queue := [][]string{[]string{beginWord}}
	isFound := false
	for len(queue) > 0{
		q := queue
		queue = nil
		for i := range q{
			path := q[i]
			nx := len(path)
			cur := path[nx-1]
			if cur == endWord{
				isFound = true
				ans = append(ans, path)
				continue
			}
			for _, nxtWord := range getNext(cur){
				// 此处 path 是slice 浅拷贝，循环其他次数时，path 值会被修改
				//queue = append(queue, append(path, nxt[i]))
				t := make([]string, len(path))
				copy(t, path)
				t = append(t, nxtWord)
				queue = append(queue, t)
			}
		}
		if isFound {
			break
		}
	}
	return ans
}
/* BFS 优化版本，上一个版本的bfs 递归树处理 PATH中没有考虑重复已经出现的， 导致占用内存超过限制
**
*/
func findLadders_BFS_Fine(beginWord string, endWord string, wordList []string) [][]string {
	ans := [][]string{}
	m, visited := map[string]bool{}, map[string]bool{}
	for i := range wordList{
		m[wordList[i]] = true
		visited[wordList[i]] = false
	}
	if _, ok := m[endWord]; !ok {
		return ans
	}
	getNext := func(word string)[]string{
		sb := []byte(word)
		res := []string{}
		for i := range sb{
			origin := sb[i]
			var c byte = 'a'
			for c <= 'z'{
				if c == origin{
					c++
					continue
				}
				sb[i] = c
				if m[string(sb)]{
					res = append(res, string(sb))
				}
				c++
			}
			sb[i] = origin
		}
		return res
	}
	//BFS 的队列不用存储 String，直接去存到目前为止的路径
	//queue := []string{beginWord}
	queue := [][]string{[]string{beginWord}}
	isFound := false
	for len(queue) > 0{
		q := queue
		queue = nil
		subVisited := []string{}// 记录bfs 当前层已经访问过的节点，不可直接使用visited 置true
		for i := range q{
			path := q[i]
			nx := len(path)
			cur := path[nx-1]
			subVisited = append(subVisited, cur)
			if cur == endWord{
				isFound = true
				ans = append(ans, path)
				continue
			}else{
				for _, nxtWord := range getNext(cur){
					if visited[nxtWord]{
						continue
					}
					// 此处 path 是slice 浅拷贝，循环其他次数时，path 值会被修改
					//queue = append(queue, append(path, nxt[i]))
					t := make([]string, len(path))
					copy(t, path)
					t = append(t, nxtWord)
					queue = append(queue, t)
				}
			}
		}
		for i := range subVisited{
			visited[subVisited[i]] = true
		}
		if isFound {
			break
		}
	}
	return ans
}

/* 433. Minimum Genetic Mutation
** A gene string can be represented by an 8-character long string, with choices from 'A', 'C', 'G', and 'T'.
** Suppose we need to investigate a mutation from a gene string start to a gene string end where one mutation is defined as one single character changed in the gene string.
	For example, "AACCGGTT" --> "AACCGGTA" is one mutation.
** There is also a gene bank bank that records all the valid gene mutations.
** A gene must be in bank to make it a valid gene string.
** Given the two gene strings start and end and the gene bank bank, return the minimum number of mutations needed to mutate from start to end.
** If there is no such a mutation, return -1.
** Note that the starting point is assumed to be valid, so it might not be included in the bank.
Constraints:
	start.length == 8
	end.length == 8
	0 <= bank.length <= 10
	bank[i].length == 8
	start, end, and bank[i] consist of only the characters ['A', 'C', 'G', 'T'].
*/
// 2021-12-23 刷出此题 bfs 最短路径 通用题型
func minMutation(start string, end string, bank []string) int {
	depth := -1
	validMap, visited := map[string]bool{}, map[string]bool{}
	for i := range bank{
		validMap[bank[i]] = true
		visited[bank[i]] = false
	}
	if _, ok := validMap[end]; !ok{
		return depth
	}
	typo := [4]byte{'A', 'C', 'G', 'T'}
	getNext := func(gene string)(genes []string){
		g := []byte(gene)
		for i := range g{
			origin := g[i]
			for j := range typo {
				if typo[j] == origin{
					continue
				}
				g[i] = typo[j]
				if validMap[string(g)]{
					genes = append(genes, string(g))
				}
			}
			g[i] = origin
		}
		return
	}
	queue := []string{start}
	for len(queue) > 0{
		depth++
		q := queue
		queue = nil
		subVisited := []string{}
		for i := range q{
			if q[i] == end{
				return depth
			}else {
				for _, nxt := range getNext(q[i]){
					if visited[q[i]]{
						continue
					}
					queue = append(queue, nxt)
				}
			}
		}
		for i := range subVisited{
			visited[subVisited[i]] = true
		}
	}
	return -1
}

/* 752. Open the Lock
** You have a lock in front of you with 4 circular wheels.
** Each wheel has 10 slots: '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'.
** The wheels can rotate freely and wrap around: for example we can turn '9' to be '0', or '0' to be '9'.
** Each move consists of turning one wheel one slot.
** The lock initially starts at '0000', a string representing the state of the 4 wheels.
** You are given a list of deadends dead ends, meaning if the lock displays any of these codes,
** the wheels of the lock will stop turning and you will be unable to open it.
** Given a target representing the value of the wheels that will unlock the lock,
** return the minimum total number of turns required to open the lock, or -1 if it is impossible.
 */
// 2021-12-24 刷出此题， 多次出错后，修正
func openLock(deadends []string, target string) int {
	const start = "0000"
	// 修正-2 ： visited 数组设置
	m, visited := map[string]bool{}, map[string]bool{start: true}
	for i := range deadends{
		m[deadends[i]] = true
	}
	/* 特殊情况，仅覆盖一例
	if _, ok := m[target]; ok{
		return depth
	}*/
	/* 下面修正特殊情况：
		1. start 和 target 在deadend列表中
		2. start == target 直接返回 0
	 */
	if m[start] || m[target]{
		return -1
	}
	if start == target {
		return 0
	}
	getNext := func(word string)(res []string){
		w := []byte(word)
		for i := range word{
			origin := w[i]
			w[i] += 1
			if origin == '9'{
				//w[i] = 0 修正-3 输入错误
				w[i] = '0'
			}
			if !m[string(w)] && !visited[string(w)]{
				res = append(res, string(w))
			}
			w[i] = origin - 1
			if origin == '0'{
				w[i] = '9'
			}
			if !m[string(w)] && !visited[string(w)]{
				res = append(res, string(w))
			}
			w[i] = origin
		}
		return
	}
	queue := []string{start}
	depth := 0
	for len(queue) > 0{
		depth++
		q := queue
		queue = nil
		/* 修正-2 visited 设置， 以及 判断target 地方的逻辑错误
		for i := range q{
			if q[i] == target{
				return depth
			}
			ret := getNext(q[i])
			for j := range ret{
				if !visited[ret[j]]{
					queue = append(queue, ret[j])
				}
			}
		}
		*/
		for i := range q{
			for _, nxt := range getNext(q[i]){
				if nxt == target{
					return depth
				}
				visited[nxt] = true
				queue = append(queue, nxt)
			}
		}
	}
	return -1
}

// 官方题解BFS: 注意 它记录深度的方法
func openLock_BFS(deadends []string, target string) int {
	const start = "0000"
	if target == start {
		return 0
	}
	dead := map[string]bool{}
	for _, s := range deadends {
		dead[s] = true
	}
	if dead[start] {
		return -1
	}
	getNext := func(word string)(res []string){
		w := []byte(word)
		for i := range word{
			origin := w[i]
			w[i] += 1
			if origin == '9'{
				w[i] = '0'
			}
			if !dead[string(w)]{
				res = append(res, string(w))
			}
			w[i] = origin - 1
			if origin == '0'{
				w[i] = '9'
			}
			if !dead[string(w)]{
				res = append(res, string(w))
			}
			w[i] = origin
		}
		return
	}
	type pair struct {
		status	string
		step	int
	}
	q := []pair{{start, 0}}
	visited := map[string]bool{start: true}
	for len(q) > 0{
		p := q[0]
		q = q[1:]
		for _, nxt := range getNext(p.status){
			if !visited[nxt]{
				if nxt == target{
					return p.step + 1
				}
				visited[nxt] = true
				q = append(q, pair{nxt, p.step + 1})
			}
		}
	}
	return -1
}
// 双向BFS 题解
func openLock_BiBFS(deadends []string, target string) int {
	const start = "0000"
	//m, visited := map[string]bool{}, map[string]bool{start: true} 不需要使用全局visited
	m := map[string]bool{}
	for i := range deadends {
		m[deadends[i]] = true
	}
	if m[start] || m[target] {
		return -1
	}
	if start == target {
		return 0
	}
	getNext := func(word string) (res []string) {
		w := []byte(word)
		for i := range word {
			origin := w[i]
			w[i] += 1
			if origin == '9' {
				w[i] = '0'
			}
			//if !m[string(w)] && !visited[string(w)] {
			if !m[string(w)] {
				res = append(res, string(w))
			}
			w[i] = origin - 1
			if origin == '0' {
				w[i] = '9'
			}
			//if !m[string(w)] && !visited[string(w)] {
			if !m[string(w)] {
				res = append(res, string(w))
			}
			w[i] = origin
		}
		return
	}
	startQ, endQ := []string{start}, []string{target}
	// depthStart, depthEnd := 0, 0 双向BFS 不能使用全局深度
	// startSeen, endSeen 借助map 初始值为 0 表示未访问的情况, 大于0 表示当前元素所处的层数深度
	//startSeen, endSeen := map[string]bool{start: true}, map[string]bool{target: true}
	startSeen, endSeen := map[string]int{start: 1}, map[string]int{target: 1}
	for len(startQ) > 0 && len(endQ) > 0 {
		if len(startQ) <= len(endQ){
			//depthStart += 1
			q := startQ
			startQ = nil
			for i := range q{
				for _, nxt := range getNext(q[i]){
					if startSeen[nxt] <= 0{ // 未被start 开始的bfs 访问过
						if endSeen[nxt] > 0 {
							//return depthStart + depthEnd
							return startSeen[q[i]] + endSeen[nxt] - 1
						}
						startSeen[nxt] = startSeen[q[i]] + 1
						startQ = append(startQ, nxt)
					}
				}
			}
		}else{
			//depthEnd += 1
			q := endQ
			endQ = nil
			for i := range q{
				for _, nxt := range getNext(q[i]){
					if endSeen[nxt] <= 0 {
						if startSeen[nxt] > 0{
							//return depthStart + depthEnd
							return endSeen[q[i]] + startSeen[nxt] -1
						}
						endSeen[nxt] = endSeen[q[i]] + 1
						endQ = append(endQ, nxt)
					}
				}
			}
		}
	}
	return -1
}
// 缩短代码长度
func openLock_BiBFS_fine(deadends []string, target string) int {
	const start = "0000"
	//m, visited := map[string]bool{}, map[string]bool{start: true} 不需要使用全局visited
	m := map[string]bool{}
	for i := range deadends {
		m[deadends[i]] = true
	}
	if m[start] || m[target] {
		return -1
	}
	if start == target {
		return 0
	}
	getNext := func(word string) (res []string) {
		w := []byte(word)
		for i := range word {
			origin := w[i]
			w[i] += 1
			if origin == '9' {
				w[i] = '0'
			}
			//if !m[string(w)] && !visited[string(w)] {
			if !m[string(w)] {
				res = append(res, string(w))
			}
			w[i] = origin - 1
			if origin == '0' {
				w[i] = '9'
			}
			//if !m[string(w)] && !visited[string(w)] {
			if !m[string(w)] {
				res = append(res, string(w))
			}
			w[i] = origin
		}
		return
	}
	getResult := func(queue *[]string, curSeen, otherSeen map[string]int)int{
		q := *queue
		*queue = nil
		for i := range q{
			for _, nxt := range getNext(q[i]){
				if curSeen[nxt] <= 0{
					if otherSeen[nxt] > 0{
						return curSeen[q[i]] + otherSeen[nxt] - 1
					}
					curSeen[nxt] = curSeen[q[i]] + 1
					*queue = append(*queue, nxt)
				}
			}
		}
		return -1
	}
	startQ, endQ := []string{start}, []string{target}
	// depthStart, depthEnd := 0, 0 双向BFS 不能使用全局深度
	// startSeen, endSeen 借助map 初始值为 0 表示未访问的情况, 大于0 表示当前元素所处的层数深度
	//startSeen, endSeen := map[string]bool{start: true}, map[string]bool{target: true}
	startSeen, endSeen := map[string]int{start: 1}, map[string]int{target: 1}
	for len(startQ) > 0 && len(endQ) > 0 {
		ret := -1
		if len(startQ) <= len(endQ){
			ret = getResult(&startQ, startSeen, endSeen)
		}else{
			ret = getResult(&endQ, endSeen, startSeen)
		}
		if ret != -1{
			return ret
		}
	}
	return -1
}










