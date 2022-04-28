package Trie

import "sort"

/* 642. Design Search Autocomplete System
** Design a search autocomplete system for a search engine.
** Users may input a sentence (at least one word and end with a special character '#').
** You are given a string array sentences and an integer array times both of length n where
** sentences[i] is a previously typed sentence and
** times[i] is the corresponding number of times the sentence was typed.
** For each input character except '#',
** return the top 3 historical hot sentences that have the same prefix as the part of the sentence already typed.
** Here are the specific rules:
	1. The hot degree for a sentence is defined as the number of times a user typed the exactly same sentence before.
	2. The returned top 3 hot sentences should be sorted by hot degree (The first is the hottest one).
		If several sentences have the same hot degree, use ASCII-code order (smaller one appears first).
	3. If less than 3 hot sentences exist, return as many as you can.
	4. When the input is a special character,
		it means the sentence ends, and in this case, you need to return an empty list.
** Implement the AutocompleteSystem class:
	1. AutocompleteSystem(String[] sentences, int[] times) Initializes the object with the sentences and times arrays.
	2. List<String> input(char c) This indicates that the user typed the character c.
		Returns an empty array [] if c == '#' and stores the inputted sentence in the system.
		Returns the top 3 historical hot sentences that have the same prefix as the part of the sentence already typed.
		If there are fewer than 3 matches, return them all.
Constraints:
	n == sentences.length
	n == times.length
	1 <= n <= 100
	1 <= sentences[i].length <= 100
	1 <= times[i] <= 50
	c is a lowercase English letter, a hash '#', or space ' '.
	Each tested sentence will be a sequence of characters c that end with the character '#'.
	Each tested sentence will have a length in the range [1, 200].
	The words in each input sentence are separated by single spaces.
	At most 5000 calls will be made to input.
*/
type AutocompleteSystem struct {
	root    Trie
	input   *Trie
	cache   []rune
	weight  map[string]int
}

func ConstructorAutocompleteSystem(sentences []string, times []int) AutocompleteSystem {
	this := AutocompleteSystem{root: Trie{next: map[rune]*Trie{}}, weight: map[string]int{} }
	this.input = &this.root
	for i := range sentences{
		this.root.Insert(sentences[i])
		this.weight[sentences[i]] = times[i]
	}
	return this
}

func (this *AutocompleteSystem) Input(c byte) []string {
	ans := []string{}
	if c == '#'{ //遗漏-1：题目要求 # 结尾后输出空，并且把此句子纳入前缀树种
		sentence := string(this.cache)
		this.root.Insert(sentence)
		this.weight[sentence]++
		this.cache = []rune{}
		this.input = &this.root
		return ans
	}
	this.cache = append(this.cache, rune(c))
	this.input = this.input.StartWith(string(c))
	if this.input == nil {
		//this.input = &this.root //遗漏-1： 此时不能直接reset状态
		//this.cache = []rune{}
		return ans
	}
	//fmt.Println(this.cache, this.input)
	s := FindAll(this.input)
	//fmt.Println(s, string(this.cache))

	/* s 有可能为 空，正好搜到一个词的时候, 思考下面例子
	    ** ["AutocompleteSystem","input","input","input","input","input","input","input","input","input","input","input","input"]
	[[["island"],[5]],["i"],[" "],["a"],["#"],["i"],[" "],["a"],["#"],["i"],[" "],["a"],["#"]]
	    for i := range s{
	        ans = append(ans, string(this.cache)+s[i])
	    }*/
	//fmt.Println(c, s)
	for i := range s{
		ans = append(ans, string(this.cache)+s[i])
	}
	//fmt.Println(ans)
	//fmt.Println(this.weight)
	// 遗漏点-4： 排序的是ans， 不是 s
	sort.Slice(ans, func(i, j int)bool{ // 遗漏点-3：不能仅靠weight 还要考虑weight相同的情况
		if this.weight[ans[i]] == this.weight[ans[j]]{
			return ans[i] < ans[j]
		}
		return this.weight[ans[i]] > this.weight[ans[j]]
	})
	//fmt.Println(ans)
	if len(ans) <= 3{ return ans}
	return ans[:3]
}

type TrieMap struct {
	next      map[rune]*TrieMap
	isEnd     bool
}
func(t *TrieMap)Insert(word string){
	root := t
	for _, ch := range word{
		if root.next[ch] == nil{
			root.next[ch] = &TrieMap{next: map[rune]*TrieMap{}}
		}
		root = root.next[ch]
	}
	root.isEnd = true
}

func(t *TrieMap)StartWith(prefix string)*TrieMap{
	if t == nil { return nil }
	root := t
	for _, ch := range prefix{
		if root.next[ch] == nil { return nil }
		root = root.next[ch]
	}
	return root
}
func FindAll(root *TrieMap)(ret []string){
	cur := []rune{}
	var dfs func(node *TrieMap)
	dfs = func(node *TrieMap){
		for k := range node.next{
			cur = append(cur, k)
			if node.next[k].isEnd{
				ret = append(ret, string(cur))
			}
			dfs(node.next[k])
			cur = cur[:len(cur)-1]
		}
	}
	dfs(root)
	if root.isEnd{// 遗漏-2： 可以直接返回，即前缀可构成一个关键词
		ret = append(ret, "")
	}
	return ret
}

/**
 * Your AutocompleteSystem object will be instantiated and called as such:
 * obj := Constructor(sentences, times);
 * param_1 := obj.Input(c);
 */