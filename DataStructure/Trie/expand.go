package Trie

/* 1858. Longest Word With All Prefixes
** Given an array of strings words, find the longest string in words such that every prefix of it is also in words.
** For example, let words = ["a", "app", "ap"]. The string "app" has prefixes "ap" and "a", all of which are in words.
** Return the string described above. If there is more than one string with the same length,
** return the lexicographically smallest one, and if no string exists, return "".
 */

type TrieTree struct{
	next      map[rune]*TrieTree
	isEnd     bool
}
func(t *TrieTree)Insert(word string){
	root := t
	for _, ch := range word{
		if root.next[ch] == nil{
			root.next[ch] = &TrieTree{next: map[rune]*TrieTree{}}
		}
		root = root.next[ch]
	}
	root.isEnd = true
}

func longestWord(words []string) string {
	//dict := TrieTree{}  // 主要map应用 nil 与 空 是不同的
	dict := TrieTree{next: map[rune]*TrieTree{}}
	height := 0
	for i := range words{
		dict.Insert(words[i])
		if height < len(words[i]){
			height = len(words[i])
		}
	}
	ans := ""
	cur := []rune{}
	// 使用回溯
	var dfs func(node *TrieTree, size int)
	dfs = func(node *TrieTree, size int){
		if size > len(ans){
			ans = string(cur)
		}else if size == len(ans){
			if ans > string(cur){
				ans = string(cur)
			}
		}
		for k := range node.next{
			if node.next[k].isEnd{
				cur = append(cur, k)
				dfs(node.next[k], size+1)
				cur = cur[:len(cur)-1]
			}
		}
	}
	dfs(&dict, 0)
	return ans
}

/* 212. Word Search II
** Given an m x n board of characters and a list of strings words, return all words on the board.
** Each word must be constructed from letters of sequentially adjacent cells,
** where adjacent cells are horizontally or vertically neighboring.
** The same letter cell may not be used more than once in a word.
 */
/* 刚开始刷此题，首先想到的竟然是，先 对 board 中每个字符开始 建立 Trie 树
** 而不是反向 从words 建立Trie树，这思维实属奇葩
 */
type TrieAlpha struct {
	next	[26]*TrieAlpha
	word 	string // 这个位置值的设定，可快速找到单词值
}
func (this *TrieAlpha) Insert(word string){
	root := this
	for i := range word{
		ch := word[i] - 'a'
		if root.next[ch] == nil{
			root.next[ch] = &TrieAlpha{}
		}
		root = root.next[ch]
	}
	root.word = word
}
/* 方法一：回溯+Trie
** 这里防止重复访问情况，使用了 visited矩阵，
** 更好的方法是，就地更改board的值
	深度优先搜索所有从当前正在遍历的单元格出发的、由相邻且不重复的单元格组成的路径。
	因为题目要求同一个单元格内的字母在一个单词中不能被重复使用；所以我们在深度优先搜索的过程中，每经过一个单元格，
	都将该单元格的字母临时修改为特殊字符（例如 #），以避免再次经过该单元格。
** 易漏点：
	1. 因为同一个单词可能在多个不同的路径中出现，所以我们需要使用哈希集合对结果集去重
	2. 在回溯的过程中，我们不需要每一步都判断完整的当前路径是否是 wordswords 中任意一个单词的前缀；
		而是可以记录下路径中每个单元格所对应的前缀树结点，每次只需要判断新增单元格的字母是否是上一个单元格对应前缀树结点的子结点即可
 */
func findWords(board [][]byte, words []string) []string {
	dirs := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(board), len(board[0])
	valid := func(r, c int)bool{
		if r >= m || r < 0 || c >= n || c < 0{
			return false
		}
		return true
	}
	visted := make([][]bool, m)
	for i := range visted{
		visted[i] = make([]bool, n)
	}
	t := &TrieAlpha{}
	for _, word := range words{
		t.Insert(word)
	}
	ans := []string{}
	seen := map[string]struct{}{} //  错误-2: 重复情况
	var dfs func(r, c int, root *TrieAlpha)
	dfs = func(r, c int, root *TrieAlpha){
		/* 错误-1
		        这块逻辑有问题，处理  [["a"]]  ["a"] case 会返回 [] 结果
		        if root == nil { return }
				ch := board[r][c] - 'a'
				if root.word == "" && root.next[ch] == nil {
					return
				}
		*/
		ch := board[r][c] - 'a'
		root = root.next[ch]
		if root == nil { return }
		if root.word != ""{
			seen[root.word] = struct{}{}
			//ans = append(ans, root.word) //  错误-2: 重复情况
		}
		visted[r][c] = true
		for _, d := range dirs{
			x, y := r + d.x, c + d.y
			if valid(x, y) && !visted[x][y]{
				// dfs(x, y, root.next[ch]) //  错误-1
				dfs(x, y, root)
			}
		}
		visted[r][c] = false
	}
	for i := range board{
		for j := range board[i]{
			dfs(i, j, t)
		}
	}
	for k := range seen{ //  错误-2: 重复情况
		ans = append(ans, k)
	}
	return ans
}

/* 方法二：删除前缀树情况
** 考虑以下情况。假设给定一个所有单元格都是 a 的二维字符网格和单词列表 ["a", "aa", "aaa", "aaaa"] 。
** 当我们使用方法一来找出所有同时在二维网格和单词列表中出现的单词时，我们需要遍历每一个单元格的所有路径，会找到大量重复的单词。
** 为了缓解这种情况，我们可以将匹配到的单词从前缀树中移除，来避免重复寻找相同的单词。因为这种方法可以保证每个单词只能被匹配一次；
** 所以我们也不需要再对结果集去重了。
 */
func findWords_improve(board [][]byte, words []string) (ans []string) {
	dirs := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(board), len(board[0])
	valid := func(r, c int)bool{
		if r >= m || r < 0 || c >= n || c < 0{
			return false
		}
		return true
	}
	t := &TrieAlpha{}
	for _, word := range words{
		t.Insert(word)
	}
	var dfs func(r, c int, root *TrieAlpha)
	dfs = func(r, c int, root *TrieAlpha){
		ch := board[r][c] - 'a'
		root = root.next[ch]
		if root == nil { return }
		if root.word != ""{
			ans = append(ans, root.word)
			root.word = "" // 置空 标记已经记录
		}
		// 执行删除任务
		if root.next != [26]*TrieAlpha{}{
			board[r][c] = '#'
			for _, d := range dirs{
				x, y := d.x + r, d.y + c
				if valid(x,y) && board[x][y] != '#'{
					dfs(x, y, root)
				}
			}
			//board[r][c] = ch
			board[r][c] = ch+'a'
		}
		if root.next == [26]*TrieAlpha{}{
			root.next[ch] = nil
		}
	}
	for i := range board{
		for j := range board[i]{
			dfs(i, j, t)
		}
	}
	return ans
}

/* 336. Palindrome Pairs
** Given a list of unique words, return all the pairs of the distinct indices (i, j) in the given list,
** so that the concatenation of the two words words[i] + words[j] is a palindrome.
 */
/* 枚举会 TlE
** 学习科学思考
** 对于一个字符串对(s1, s2), 若想要字符串 s1+s2为回文串，则必须满足以下条件之一：
** 1. len(s1) == len(s2), s1 是 s2 的翻转
** 2. len(s1) > len(s2) ，字符串 s1 拆分为t1+t2，其中 t1 是 s2 的翻转 并且 t2 必须是一个回文串
** 3. len(s1) < len(s2) ，字符串 s2 拆分为t1+t2，其中 t2 是 s1 的翻转 并且 t1 必须是一个回文串
** 这样，对于每一个字符串，我们令其为 s1 和 s2 中较长的那个，然后找到可能和它构成回文串的字符串即可
** 具体地说，我们枚举每一个字符串 k, 令其为 s1 和 s2 中较长的那一个，那么 k 可以被分成两个部分，t1 和 t2
** 1. 当 t1 为 回文时候， 符合情况3 我们只需要查询给定的字符串序列中是否包含 t2 的翻转
** 2. 当 t2 为 回文时候， 符合情况2 我们只需要查询给定的字符串序列中是否包含 t1 的翻转
** 也就是说，我们要枚举字符串 k 的每一个前缀和后缀，判断其是否为回文串。
** 如果是回文串，我们就查询其剩余部分的翻转是否在给定的字符串序列中出现即可
** 注意到空串也是回文串，所以我们可以将 k 拆解为 k+∅ 或 ∅+k，这样我们就能将情况 1 也解释为特殊的情况 2 或情况 3。
** 而要实现这些操作，我们只需要设计一个能够在一系列字符串中查询「某个字符串的子串的翻转」是否存在的数据结构，有两种实现方法：
	1. 我们可以使用字典树存储所有的字符串。在进行查询时，我们将待查询串的子串逆序地在字典树上进行遍历，即可判断其是否存在
	2. 我们可以使用哈希表存储所有字符串的翻转串。在进行查询时，我们判断带查询串的子串是否在哈希表中出现，就等价于判断了其翻转是否存在
** 例如 [cat, ac]， 以cat为主串，去查看是否有 ac
   cat 拆分可能：
	["", "cat"]   ""是回文，但cat 与 ac 不互为回文
	["c", "at"]   "c" 是回文，但 at 与 ac 不互为回文
	["ca", "t"]   "t" 是回文 并且 ca ac 互为回文
	["cat", ""]    "" 是回文， 但 cat 与 ac 不互为回文
 */
func palindromePairs_hashmap(words []string) [][]int {
	m := map[string][]int{}
	for i := range words {
		m[words[i]] = append(m[words[i]], i)
	}
	ans := [][]int{}
	for i, word := range words{
		/*特例：如果word 本身为回文，匹配空串 */
		if isPalindrome(word) {// 补充 // 此情况匹配不了["a", ""] 情况
			for _, j := range m[""]{
				if i != j{
					ans = append(ans, []int{j, i})
				}
			}
		}
		for k := range word{
			t1, t2 := word[:k], word[k:]
			if isPalindrome(t1){
				// 此情况匹配不了["a", ""] 情况，因为此情况下，t 实际上是 word 自身，会导致i == j 上面特例处理规避这种情况
				t := reverse(t2) // 要匹配t2， 因此 m[t] 放在左边
				for _, j := range m[t]{
					if i != j{
						ans = append(ans, []int{j, i})
					}
				}
			}
			if isPalindrome(t2){
				t := reverse(t1) // 要匹配t1， 因此 m[t] 放在右边
				for _, j := range m[t]{
					if i != j{
						ans = append(ans, []int{i, j})
					}
				}
			}
		}
	}
	return ans
}
// 下面的方法 提前将 word 反转， 预处理优化
func palindromePairs(words []string) [][]int {
	m := map[string][]int{}
	for i := range words{
		t := reverse(words[i])
		m[t] = append(m[t], i)
	}
	ans := [][]int{}
	for i, word := range words{
		if word != "" && isPalindrome(word) && m[""] != nil{
			for _, j := range m[""]{
				ans = append(ans, []int{j, i})
			}
		}
		for k := range word{
			left, right := word[:k], word[k:]
			if isPalindrome(left) && m[right] != nil{
				for _, j := range m[right]{
					if i != j{
						ans = append(ans, []int{j, i})
					}
				}
			}
			if isPalindrome(right) && m[left] != nil {
				for _, j := range m[left]{
					if i != j{
						ans = append(ans, []int{i, j})
					}
				}
			}
		}
	}
	return ans
}
func isPalindrome(s string)bool{
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j]{
			return false
		}
	}
	return true
}
func reverse(s string)string{
	t := []byte(s)
	for i, j := 0, len(t)-1; i < j; i, j = i+1, j-1{
		t[i], t[j] = t[j], t[i]
	}
	return string(t)
}
// 方式2： Trie 树存
type TrieIndex struct{
	next 	[26]*TrieIndex
	index	int
}

func(this *TrieIndex)Insert(word string, idx int) {
	root := this
	for i := range word{
		ch := word[i] - 'a'
		if root.next[ch] == nil {
			root.next[ch] = &TrieIndex{}
		}
		root = root.next[ch]
	}
	root.index = idx
}
func(this *TrieIndex)findWord(word string)int{
	root := this
	//for i := range word{  逆序查询
	for i := len(word)-1; i >= 0; i--{
		ch := word[i] - 'a'
		if root.next[ch] == nil {
			return -1
		}
		root = root.next[ch]
	}
	return root.index
}
// 特别注意迷糊点
func palindromePairs_Trie(words []string) [][]int {
	root := &TrieIndex{}
	for i := range words{
		root.Insert(words[i], i)
	}
	ans := [][]int{}
	for i, word := range words{
		n := len(word)
		for k := 0; k <= n; k++{  // 迷糊点-1： k 到 n 位置 用来处理 边界情况
			if isPalindrome(word[k:n-1]) {
				left := root.findWord(word[:k-1])
				if left != -1 && left != i{
					ans = append( ans, []int{i, left} )
				}
			}
			if k != 0 && isPalindrome(word[:k-1]){ // 迷糊点-1： k != 0 用来处理 边界情况  两者联合， 错开 避免重复
				right := root.findWord(word[k:n-1])
				if right != -1 && right != i{
					ans = append(ans, []int{right, i})
				}
			}
		}
	}
	return ans
}

/* Manacher 算法 和 双Trie树
** 上述方法中， 对于每一个字符串 k，我们需要 子串长度平方的复杂度判断 k 的所有前缀与后缀是否是回文串
** 还需要地 子串长度平方的复杂度 判断所有前缀与后缀是否在给定字符串序列中出现
** 对于判断其所有前缀与后缀是否是回文串：
	利用 Manacher 线性处理是否为回文串
** 对于判断其所有前缀与后缀是否在给定的字符串序列中出现
	对于给定的字符串序列，分别正向与反向建立字典树，利用正向建立的字典树验证 k 的后缀的反转
	利用反向建立的字典树验证 k 的前缀反转
*/









