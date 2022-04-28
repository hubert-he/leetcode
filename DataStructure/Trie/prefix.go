package Trie
/* Trie tree
** In computer science, a trie, also called digital tree, radix tree or prefix tree,
** is a kind of search tree—an ordered tree data structure used to store a dynamic set or associative array
** where the keys are usually strings.
** Trie tree，又被称为字典树或前缀树。从名字我们可以推断出其可以被用来查找字符串。
** 字典树的特点：
	1. Trie tree用边来表示字母
	2. 有相同前缀的单词公用前缀节点。那么我们可以知道，在单词只包含小写字母的情况下，我们可以知道每个节点最多有26个子节点
	3. 整棵树的根节点是空的
	4. 每个单词结束的时候用特殊字符表示(比如上图的$)，在代码中可以单独建立一个bool字段来表示是否是单词结束处
** 基本操作：
	1. insert
	  从左到右扫描新单词，如果字母在相应根节点下没有出现过，就插入这个字母；否则沿着字典树往下走，看单词的下一个字母。
		根据编码方式：
		编码方式1： 按照输入顺序对其进行编码，会出现相同字母的编码可能不同
		编码方式2： 因为每个节点最多26个子节点，我可以可以按他们的字典序0-25编号，这里相同字母的编码相同
	2. search
** 存储方式：
	1. 数组模拟
	2. 类的形式

** 前缀树的应用：
1. autocomplete
2. spell checker
3. IP routing(longest prefix matching)
4. T9 predictive text
 */

/* 208. Implement Trie (Prefix Tree)
** A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys
** in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.
** Implement the Trie class:
	Trie() Initializes the trie object.
	void insert(String word) Inserts the string word into the trie.
	boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before),
		and false otherwise.
	boolean startsWith(String prefix) Returns true
		if there is a previously inserted string word that has the prefix prefix, and false otherwise.
 */
/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
/* 方法一：二维数组
** 使用二维数组 trie[] 来存储我们所有的单词字符。
** 使用 index 来自增记录我们到底用了多少个格子（相当于给被用到格子进行编号）。
** 使用 count[] 数组记录某个格子被「被标记为结尾的次数」（当 idx 编号的格子被标记了 n 次，则有 cnt[idx] = n  ）。
 */

type Trie_Array struct {
	// root [26]*[26]byte  // 本来想用此二维数组实现，无法处理单词终结条件
}


func ConstructorTrie_Array() Trie_Array {
	this := TrieArray{}
	return this
}


func (this *Trie_Array) Insert(word string)  {
	n := len(word)
	if n == 0 { return }
	root := this.root[word[0]]
	for i := 1; i < n; i++{
		ch := word[i]
		if root[ch] != nil {
			root = root[ch]
		}
	}
}


func (this *Trie_Array) Search(word string) bool {

}


func (this *Trie_Array) StartsWith(prefix string) bool {

}

/* 方法二：设置对象
** 使用固定数组实现，特别注意 要 指针使用
** 对照方法二 与 方法三：可以看到 slice 与 固定数组的 使用区别，特别注意赋值传递值的时候
 */
type TrieArray struct {
	next 	[26]*TrieArray
	end		bool // 标记到此可以构成一个单词
}

func ConstructorTrieArray() TrieArray {
	return TrieArray{}
}

// Insert_Error 实现中有个很大bug，即 root 是值拷贝，没用应用到真正的this中
func (this *TrieArray) Insert_Error(word string)  {
	root := this.next
	var end *TrieArray
	for i := range word{
		ch := int(word[i] - 'a')
		if root[ch] == nil{
			root[ch] = &TrieArray{}
		}
		end = root[ch]
		root = root[ch].next
	}
	end.end = true
}

func (this *TrieArray) Insert_Error_Fix(word string)  {
	root := &this.next // 采用指针
	var end *TrieArray
	for i := range word{
		ch := int(word[i] - 'a')
		if (*root)[ch] == nil{
			(*root)[ch] = &TrieArray{}
		}
		end = (*root)[ch]
		//(*root) = (*root)[ch].next
		root = &(*root)[ch].next  // 指针，真正修改地址
	}
	end.end = true

}

func (this *TrieArray) Search(word string) bool {
	root := this.next
	var end *TrieArray
	for i := range word{
		ch := word[i] - 'a'
		if root[ch] == nil { return false }
		end = root[ch]
		root = root[ch].next
	}
	return end.end
}


func (this *TrieArray) StartsWith(prefix string) bool {
	root := this.next
	for i := range prefix{
		ch := prefix[i] - 'a'
		if root[ch] == nil { return false }
		root = root[ch].next
	}
	return true
}

/* 方法三：不使用 [26] 数组方式， 使用slice
**  此时关注 Insert 函数，发现 Slice 直接传递的就是指针，而应用数组，传递过来的就是拷贝值
** 这是因为golang 全部为值传递方式，只是因为slice 本质上是一个结构体，内含指针
 */

type TrieSlice struct {
	next 	[]*TrieSlice
	end		bool // 标记到此可以构成一个单词
}

func ConstructorTrieSlice() TrieSlice {
	this := TrieSlice{next: make([]*TrieSlice, 26)}
	return this
}


func (this *TrieSlice) Insert(word string)  {
	root := this.next
	var end *TrieSlice
	for i := range word{
		ch := int(word[i] - 'a')
		if root[ch] == nil{
			root[ch] = &TrieSlice{next: make([]*TrieSlice, 26)}
		}
		end = root[ch]
		root = root[ch].next
	}
	end.end = true
}


func (this *TrieSlice) Search(word string) bool {
	root := this.next
	var end *TrieSlice
	for i := range word{
		ch := word[i] - 'a'
		//fmt.Println(root[ch], word[i])
		if root[ch] == nil { return false }
		end = root[ch]
		root = root[ch].next
	}
	return end.end
}


func (this *TrieSlice) StartsWith(prefix string) bool {
	root := this.next
	for i := range prefix{
		ch := prefix[i] - 'a'
		if root[ch] == nil { return false }
		root = root[ch].next
	}
	return true
}

/* 方法四：观察Insert 方法，完全没必要从next 开始，无需end 变量来跟踪
** 直接从 this 出发
 */
type Trie struct {
	next 	[26]*Trie
	end		bool // 标记到此可以构成一个单词
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string)  {
	root := this
	for i := range word{
		ch := word[i] - 'a'
		if root.next[ch] == nil {
			root.next[ch] = &Trie{}
		}
		root = root.next[ch]
	}
	root.end = true
}

func (this *Trie) Search(word string) bool {
	root := this
	for i := range word{
		ch := word[i] - 'a'
		if root.next[ch] == nil { return false }
		root = root.next[ch]
	}
	return root.end
}

func (this *Trie) StartsWith(prefix string) bool {
	root := this
	for i := range prefix{
		ch := prefix[i] - 'a'
		if root.next[ch] == nil { return false }
		root = root.next[ch]
	}
	return true
}







