package String

import "fmt"
// 算法哲学： 降低复杂度一种思维是 预处理一些信息，来降低主算法运行时间

/* KMP算法的主要思想是提前判断如何重新开始查找，而这种判断只取决于模式字符串本身
** 为了优化时间复杂度，提出了 预处理patter机制，引入了next 数组
** KMP 的 next 数组相当于告诉我们：当模式串中的某个字符跟文本串中的某个字符匹配失配时，模式串下一步应该跳到哪个位置。
** 如果模式串中在j 处的字符跟文本串在 i 处的字符匹配失配时，
** 下一步用 next [j] 处的字符继续跟文本串 i 处的字符匹配，相当于模式串向右移动 j - next[j] 位
 */

// 查找首次的KMP 子串匹配
func kmpMatch(s, pattern string, getNext func(string)[]int) bool{
	n := len(s)
	next := getNext(pattern)
	i, j := 0, -1
	for i < n && j < len(pattern){
		if j == -1 || s[i] == pattern[j]{
			i, j = i+1, j+1
		}else{
			j = next[j]
		}
	}
	if j == len(pattern){
		return true
	}
	return false
}

func kmpSearch(s, pattern string, getNext func(string)[]int) bool{
	n := len(s)
	next := getNext(pattern)
	for i,j := 0, 0; i < n; i++{
		if j == len(pattern) - 1 && s[i] == pattern[j]{
			fmt.Println("Found pattern at %d", i - j)
			j = next[j] // 继续向后匹配
		}
		if s[i] == pattern[j]{
			i++
			j++
		}else{
			j = next[i]
			if j < 0{ // next[i] == -1情况
				i++
				j++
			}
		}
	}
}
/*
** 先 求出最长公共前后缀长度，然后向右平移1位，构成next数组，然后初值赋为 -1 , 下面 getNext_0  与 getNext_1 为此类方式获得next数组
** next[i]可以看作直接计算某个i字符对应的 next 值，就是看这个字符之前(i-1)的字符串中有多大长度的相同前缀后缀
** 0  1  2  3  4  5  6  7  8   <== 下标
** A  B  A  B  C  A  B  A  A
** 0  0  1  2  0  1  2  3  1   <== 最大公共前后缀元素的长度，注意最后一个A 的 长度计算
** -1 0  0  1  2  0  1  2  3   <== next 数组， 恰好是Pattern 待比较的回滚元素的下标
**
** 例如要找next[8] = ? 的值
** 首先知道的是 p[0...7] 即 A  B  A  B  C  A  B  A 【A】可存在的 公共前后缀有
 		1. A
		2. A  B
		3. A  B  A  B
** 可以知道的信息是离 8 最近的 前后缀是ABA，然后 确定 ABAB 是否 可以与 ABA+p[8] 构成功公共前后缀
** 但是p[8] = A  因此不能构成 ABAB 长度为 3+1 的公共前后缀，所以 A B A 这个公共前后缀是不能用的
** 然后继续 发现 p[0...7] 次短的公共前后缀 为 A , p[8] = A 因此结果为 next[8] = next[0]+1
 */
func getNext_0(p string) []int{
	n := len(p)
	next := make([]int, n)
	i := 1
	len := 0
	for i < n{
		if p[i] == p[len]{
			len++
			next[i] = len
			i++
		}else{
			if len > 0{
				len = next[len-1]
			}else{
				next[i] = len
				i++
			}

		}
	}
	for i := n-1; i > 0; i--{
		next[i] = next[i-1]
	}
	next[0] = -1
	return next
}

func getNext_1(p string) []int{
	n := len(p)
	next := make([]int, n) // next[0] = 0
	for i := 1; i < n; i++{
		k := next[i-1]
		for p[i] != p[k] && k > 0{ // 对应不相等的情况， 此处的选好可以看 上面的 表（模式子串，前缀，后缀）
			k = next[k-1]  // 务必注意 不是 k = next[k], 此处的循环在往 更小的 公共前后缀里找可以匹配的长度
		}
		if p[i] == p[k]{
			next[i] = k + 1
		}else{ // k == 0 了， 即没有找到 匹配的公共前后缀
			next[i] = 0
		}
	}
	for i := n-1; i > 0; i--{
		next[i] = next[i-1]
	}
	next[0] = -1
	return next
}

/* getNext_0 的另外一种 直接求next写法， 即 主循环 i 下 计算 next[i+1] 的值， 避开后面的拷贝动作
** 通过前面例子的推算（其实存在一个状态机），发现了新的结果可以有前面的子状态相关的
** next[j] = k 表示已经有 p[0],p]1],...,p[k-1]  == p[j-k], p[j-k+1], ..., p[j-1]
** next[j] = k 本质上为 p[j] 之前的模式串子串中， 有长度为 k 的相同前缀和后缀。
** 如何 已知 next[0...j] 推出 next[j+1] 的状态方程
** 这里因为相等不相等条件，分成了2个状态方程：
** 1. p[next[j]] == p[k] == p[j] 的情况下，next[j+1] = next[j] + 1 = k + 1
** 2. p[next[j]] == p[k] != p[j] 的情况下，
		如果此时 p[ next[k] ] == p[j]，则 next[j + 1] = next[k] + 1
		否则继续递归前缀索引 k = next[k]，而后重复此过程。
	这相当于在字符 p[j+1] 之前不存在长度为 k+1 的前缀"p0 p1, …, pk-1 pk"跟后缀“pj-k pj-k+1, …, pj-1 pj"相等，
	那么是否可能存在另一个值 t+1 < k+1，使得长度更小的前缀 “p0 p1, …, pt-1 pt” 等于长度更小的后缀 “pj-t pj-t+1, …, pj-1 pj” 呢？
	如果存在，那么这个 t+1 便是 next[j+1] 的值，
	此相当于利用已经求得的 next 数组（next [0, ..., k, ..., j]）进行 P 串前缀跟 P 串后缀的匹配。
 */
func getNext_2(p string) []int{
	n := len(p)
	next := make([]int, n)
	next[0] = -1
	k := -1 // p[k] 表示前缀
	i := 0 // 填充next, p[i] 表示后缀
	for i < n-1{ // 想想为何 是 n-1? next 算的是上一个的最长公共前后缀的大小
		if k == -1 || p[i] == p[k]{ // p[k] 表示前缀， p[i] 表示后缀
			i++
			k++
			next[i] = k
		}else{ // 注意此处其实是一个循环
			k = next[k]
		}
	}
	return next
}
// getNext_2 的容易理解的版本，明显看出不匹配时是一个循环
func getNext_3(p string) []int{
	n := len(p)
	next := make([]int, n)
	next[0] = -1 // next[1] = 0 肯定的
	for i := 1; i < n-1; i++{// 根据 i 求 next[i+1]
		k := next[i]
		for p[k] != p[i] && k > 0{
			k = next[k]
		}
		if p[k] == p[i] {
			next[i+1] = k + 1
		}else{
			next[i+1] = 0
		}
	}
	return next
}

/* 进一步优化版本
** https://wiki.jikexueyuan.com/project/kmp-algorithm/define.html
 */
func getNextBest_1(p string)[]int{ // 此代码 不正确 待调整！！！
	n := len(p)
	next := make([]int, n)
	next[0] = -1
	for i := 1; i < n-1; i++{
		//p[k]表示前缀，p[i]表示后缀
		k := next[i]
		for p[k] != p[i] && k > 0{
			k = next[k]
		}
		if p[k] == p[i]{
			if p[i+1] == p[k+1]{ // 如果出现 p[i] == p[ next[i] ], 不能直接赋值，需要返回下一个前后缀串的长度
				next[i+1] = next[k+1]
			}else{
				next[i+1] = k + 1
			}
		}else{
			next[i+1] = 0
		}
	}
	return next
}
/* 进一步优化版本
** 优化的很精妙
 */
func getNextBest_2(p string)[]int{
	n := len(p)
	next := make([]int, n)
	next[0] = -1
	i := 1
	k := -1
	for i < n-1{
		//p[k]表示前缀，p[i]表示后缀
		if k == -1 || p[i] == p[k]{
			i++
			k++
			if p[i] != p[k]{
				next[i] = k
			}else{
				next[i] = next[k]
			}
		}else{
			k = next[k]
		}
	}
	return next
}


/* 基于自动机-DFA 实现的 KMP
** 只需要知道位于文本字符串上该位置的字符与模式字符串上的哪个位置的字符进行匹配
** 把j停留在的位置理解成一种状态，模式字符串一共有0，1，2，3，…，M-1这M个状态，
** 那么用KMP算法成功查找到模式字符串第一次出现的位置的过程，其实也是一种j从0不断跳转，直到跳转到M-1状态并满足t[i] == p[j]的过程。
** 需要构造一个int类型的二维数组dfa，用来模拟一个确定有限状态自动机：
	第一维：对应需要比较的文本字符串的字符,Ascii 码需要 255， Unicode就会比较多
	第二维：表示要比较的模式字符串所处的状态；
** dfa元素的值：表示比较后应该跳转到的状态；
 */
func kmpMatch_DFA(s, pattern string) bool{
	dfa := getDFA(pattern)
	state := 0
	n := len(pattern) // DFA 状态总数
	for i := range s{
		if state == n{ return true }
		state = dfa[int(s[i])][state]
	}
	return false
}
/* Match    Transition:  If in state j  and next char c == p[j] go to j+1
** Mismatch Transiton:   If in state j and next char c != p[j] then
						「 the last j-1 characters of input are p[1...j-1], followed by c  」
** To compute dfa[c][j]: Simulate p[1...j-1] on DFA and take transition c , 其中 state x 表示  p[1...j-1]
** Running time : Takes only O(1) time if we maintain state x
** The key is keeping track of the state where the machine would be if we had backed up or if we had run it on the pattern shifted over one.
 */
func getDFA(p string)[][]int{
	m := len(p)
	R := 255 // 表示ASCII 字符
	dfa := make([][]int, R)
	for i := range dfa{
		dfa[i] = make([]int, m)
	}
	dfa[p[0]][0] = 1
	for x, j := 0, 1; j < m; j++{
		// 1. copy mismatch cases
		for c := 0; c < R; c++{
			dfa[c][j] = dfa[c][x]
		}
		// 2. set match case
		dfa[p[j]][j] = j + 1
		// 3. upate restart state
		x = dfa[p[j]][x]
	}
	return dfa
}