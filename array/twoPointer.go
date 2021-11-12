package array

/* 977. Squares of a Sorted Array
** Given an integer array nums sorted in non-decreasing order,
** return an array of the squares of each number sorted in non-decreasing order.
*/
func SortedSquares(nums []int) []int {
	n := len(nums)
	ans := []int{}
	i := 0
	for i < n{
		if nums[i] < 0{
			i++
		}else{
			break
		}
	}
	a, b := nums[:i], nums[i:]
	slow, fast := len(a)-1, 0
	for slow >= 0 || fast < len(b){
		if  slow >= 0 && fast < len(b){
			av, bv := a[slow]*a[slow], b[fast]*b[fast]
			if av < bv{
				ans = append(ans, av)
				slow--
			}else{
				ans = append(ans, bv)
				fast++
			}
		}else if slow >= 0{
			ans = append(ans, a[slow]*a[slow])
			slow--
		}else{
			ans = append(ans, b[fast]*b[fast])
			fast++
		}
	}
	return ans
}

/* 844. Backspace String Compare
** Given two strings s and t, return true if they are equal when both are typed into empty text editors.
** '#' means a backspace character.
** Note that after backspacing an empty text, the text will continue empty.
** 双指针 值得注意的 一个case
"ab##"
"c#d#"
 */
/* 定义 skip 表示当前待删除的字符的数量。每次我们遍历到一个字符：
**
 */
func BackspaceCompare(s string, t string) bool {
	i, j := len(s)-1, len(t)-1
	skipS, skipT := 0, 0
	for i >= 0 || j >= 0{
		for i >= 0{ // 迭代删除："c#d#" 为 ""  就是这个迭代 思路 卡住了
			if s[i] == '#'{
				skipS++
				i--
			}else if skipS > 0{
				skipS--
				i--
			}else{
				break
			}
		}
		for j >= 0{
			if t[j] =='#'{
				skipT++
				j--
			}else if skipT > 0{
				skipT--
				j--
			}else{
				break
			}
		}
		// 此时 进行到真正的字符
		if i >= 0 && j >= 0{
			if s[i] != t[j]{
				return false
			}
		}else if i >= 0 || j >= 0{
			return false
		}
		i--
		j--
	}
	return true
}
/* 438. Find All Anagrams in a String
** Given two strings s and p, return an array of all the start indices of p's anagrams in s.
** You may return the answer in any order.
** An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
** typically using all the original letters exactly once.
 */
/*滑动窗口*/
func FindAnagrams(s string, p string) []int {
	ns, np := len(s), len(p)
	ans := []int{}
	if ns < np{
		return ans
	}
	sCnt, pCnt := [26]int{}, [26]int{}
	// 前np个字符进行统计
	for i := 0; i < np; i++{
		sCnt[s[i]-'a']++
		pCnt[p[i]-'a']++
	}
	// 如果数组相等，则找到第一个异位词
	if sCnt == pCnt{
		ans = append(ans, 0)
	}
	// 继续遍历剩余字符串，在sCnt中每次增加一个新字母，去除一个旧字母
	for i := np; i < ns; i++{
		// 滑动窗口的基本使用
		toDel := s[i-np] - 'a'
		toAdd := s[i] - 'a'
		sCnt[toDel]--
		sCnt[toAdd]++
		if sCnt == pCnt{
			ans = append(ans, i-np+1)
		}
	}
	return ans
}
/* 除直接比较统计数组是否相等外， 还可以用双指针来表示滑动窗口的两侧边界
** 当滑动窗口的长度等于p的长度时，表示找到一个异位词，两种方式的时间复杂度都是O(n)级别的
** 1. 定义滑动窗口的左右两个指针left，right
** 2. right一步一步向右走遍历s字符串
** 3. right当前遍历到的字符加入s_cnt后不满足p_cnt的字符数量要求，将滑动窗口左侧字符不断弹出，也就是left不断右移，直到符合要求为止
** 4. 当滑动窗口的长度等于p的长度时，这时的s子字符串就是p的异位词
*/
func FindAnagrams2Pointer(s string, p string) []int {
	ns, np := len(s), len(p)
	ans := []int{}
	if ns < np{
		return ans
	}
	sCnt, pCnt := [26]int{}, [26]int{}
	// 先统计标志字符串
	for i := 0; i < np; i++{
		pCnt[p[i]-'a']++
	}
	left := 0
	for right := 0; right < ns; right++{
		curR := s[right] - 'a'
		sCnt[curR]++
		for sCnt[curR] > pCnt[curR] {
			curL := s[left]-'a'
			sCnt[curL]--
			left++
		}
		if right - left + 1 == np{
			ans = append(ans, left)
		}
	}
	return ans
}








