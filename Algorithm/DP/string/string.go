package string

import (
	"fmt"
	"strings"
)

/* 392. Is Subsequence
** Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
** A subsequence of a string is a new string
** that is formed from the original string by deleting some (can be none) of the characters
** without disturbing the relative positions of the remaining characters.
** (i.e., "ace" is a subsequence of "abcde" while "aec" is not).
** Follow up: Suppose there are lots of incoming s, say s1, s2, ..., sk where k >= 109,
** 		and you want to check one by one to see if t has its subsequence.
** 		In this scenario, how would you change your code?
 */
/* 此题可以用双指针，大量操作用于在t 中找下一个匹配的字符
** 后续进阶 DP
** 预处理出对于t 的每一个位置，从该位置开始往后每一个字符第一次出现的位置
** 动态规划实现预处理
** dp[i][j] 表示字符串 t 从 i 开始往后字符 j 第一次出现的位置
** dp[i][j] = i  			若 t[i] == j
**          = dp[i+1][j] 	若 t[i] != j
** 每次 O(1) 找出 t 中下一个位置
** 同类题目：
** 1055. Shortest Way to Form String
** 1182. 与目标颜色间的最短距离
 */
func isSubsequence(s string, t string) bool {
	ns, nt := len(s), len(t)
	if nt == 0{
		if ns == 0{
			return true
		}
		return false
	}
	//dp[i][j] 表示字符串 t 从 i 开始往后字符 j 第一次出现的位置
	dp := make([][26]int, nt)
	for i := 0; i < 26; i++{
		dp[nt-1][i] = nt
	}
	dp[nt-1][t[nt-1]-'a'] = nt-1
	for i := nt-2; i >= 0; i--{
		for j := 0; j < 26; j++{
			c := int(t[i] - 'a')
			if c == j{
				dp[i][j] = i
			}else{
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	j := 0
	for i := range s{
		loc := int(s[i] - 'a')
		if j >= nt || dp[j][loc] == nt{
			return false
		}
		j = dp[j][loc] + 1
	}
	return true
}
// 针对特殊情况的处理
func isSubsequence2(s string, t string) bool {
	nt := len(t)
	dp := make([][26]int, nt+1) // 额外申请一个，方便特视情况处理
	for i := range dp[nt]{ // 设置保护边界
		dp[nt][i] = nt // 表示不存在
	}
	for i := nt-1; i >= 0; i--{
		for j := range dp[i]{
			c := int(t[i] - 'a')
			if c == j{
				dp[i][j] = i
			}else{
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	j := 0
	for i := range s{
		c := int(s[i] - 'a')
		if dp[j][c] == nt{
			return false
		}
		j = dp[j][c] + 1
	}
	return true
}

/* 1055. Shortest Way to Form String
** A subsequence of a string is a new string
** that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters.
** (i.e., "ace" is a subsequence of "abcde" while "aec" is not).
Given two strings source and target,
** return the minimum number of subsequences of source such that their concatenation equals target. If the task is impossible, return -1.
 */
/* 2012-12-06 DP思路 卡在 当前匹配到的字符串 这一层
** 1. 当有字符不在源字符串中， 可返回-1
** 2. dp[i]: 表示以 i 结尾的目标字符串使用源串可构造出的最少数量
** 3. 转移方程： 设当前匹配到的字符串是tmp 这个时候分2种情况：
**    3.1 tmp 是 源串的子序列， 则 dp[i] = dp[i-1]
**	  3.2 tmp 不是 源串的子序列， 则 dp[i] = dp[i-1] + 1 且重新初始 tmp 作为当前匹配到的字符串
** 初始边界： dp[0] = 1
** 例子：source="abc",target="abcbc"
** 遍历目标字符串target，当前遍历的下标设为i。
	初始情况：tmp="a",dp[0]=1
	i=1,tmp="ab", tmp是source的子序列,所以dp[1]=dp[0]=1
	i=2,tmp="abc", tmp是source的子序列,所以dp[2]=dp[1]=1
	i=3,tmp="abcb", tmp不是source的子序列,所以dp[3]=dp[2]+1=2，且更新tmp为当前字符的字符串，即tmp="b"
	i=4,tmp="bc",tmp是source的子序列,所以dp[4]=dp[3]=2
	结束，返回dp[4]=2
 */
func ShortestWay(source string, target string) int {
	dp := 1
	n := len(target)
	tmp := []byte{}
	for i := 1; i <= n; i++{
		tmp = append(tmp, target[i])
		if !isSubsequence(source, string(tmp)){
			tmp = []byte{target[i]}
			dp = dp + 1
		}
	}
	return dp
}

/*712. Minimum ASCII Delete Sum for Two Strings
** Given two strings s1 and s2, return the lowest ASCII sum of deleted characters to make two strings equal.
*/
/* 状态定义： dp[i][j]: 表示字符串s1[i:] 和 s2[j:] 达到相等所需删除的字符的 ASCII 值的最小和，最终的答案为 dp[0][0]
** 1. 当 s1[i:] 和 s2[j:] 中的某一个字符串为空时， dp[i][j]的值即为另一个非空字符串的所有字符的ASCII值的最小和
**    例如：当s2[j:]为空时，此时有 j = len(s2) 状态转移方程： dp[i][j] = s1.asciiSumFromPos(i) 转换为递推=>
**    dp[i][j] = dp[i+1][j] + s1.asciiAtPos(i)
**2. 两个字符串都非空时，
** 2.1 如果 s1[i] = s2[j], 那么当前位置的两个字符相同，他们不需要被删除，状态转移方程：
		dp[i][j] = dp[i+1][j+1]
** 2.2 如果 s1[i] != s2[j],那么我们至少要删除s1[i] 和 s2[j]两个字符中的一个，方程：
** 		dp[i][j] = min(dp[i+1][j] + s1.asciiAtPos(i), dp[i][j+1] + s2.asciiAtPos(j))
*/
func minimumDeleteSum(s1 string, s2 string) int {
	n1, n2 := len(s1), len(s2)
	dp := make([][]int, n1+1)
	for i := range dp{
		dp[i] = make([]int, n2+1)
	}
	// 初始化, 对应第一种情况
	for i := n1-1; i >= 0; i--{
		dp[i][n2] = dp[i+1][n2] + int(s1[i])
	}
	for i := n2-1; i >= 0; i--{
		dp[n1][i] = dp[n1][i+1] + int(s2[i])
	}
	for i := n1-1; i >= 0; i--{
		for j := n2-1; j >= 0; j--{
			if s1[i] == s2[j]{
				dp[i][j] = dp[i+1][j+1]
			}else{
				dp[i][j] = min(dp[i+1][j]+int(s1[i]), dp[i][j+1]+int(s2[j]))
			}
		}
	}
	return dp[0][0]
}
/* 方法二： 转换为 LCS 去处理
** dp[i][j] 为S1, S2最大公共子串的所有字母的和，那么该问题的解为 Sum(S1,S2) - 2 * dp[s1.size()][s2.size()]
*/
func MinimumDeleteSum_LCS(s1 string, s2 string) int {
	n1, n2 := len(s1), len(s2)
	sum := 0
	for i := range s1{
		sum += int(s1[i])
	}
	for i := range s2{
		sum += int(s2[i])
	}
	// lcs 求最大公共子序列
	dp := make([][]int, n1+1)
	for i := range dp{
		dp[i] = make([]int, n2+1)
	}
	for i := 1; i <= n1; i++{
		for j := 1; j <= n2; j++{
			if s1[i-1] == s2[j-1]{
				dp[i][j] = dp[i-1][j-1] + int(s1[i-1]) // s1[i] == s2[j]
			}else{
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return sum - dp[n1][n2]*2
}

/* 647. Palindromic Substrings
** Given a string s, return the number of palindromic substrings in it.
** A string is a palindrome when it reads the same backward as forward.
** A substring is a contiguous sequence of characters within the string.
 */
/* 枚举思路-1：枚举出所有的子串，然后再判断这些子串是否是回文
	复杂度立方级， 枚举子串 平方级别， 判断回文 线性级别
** 枚举思路-1：枚举每一个可能的回文中心，然后用两个指针分别向左右两边拓展，当两个指针指向的元素相同的时候就拓展，否则停止拓展
	复杂度平方级， 枚举回文中心的是线性， 对于每个回文中心拓展的次数也是线性
** 	核心： 分奇偶确定 字符串中心位置。 即 如果回文长度是奇数，那么回文中心是一个字符；如果回文长度是偶数，那么中心是两个字符
**  处理方法一： 2个循环分别处理 偶数长度 和 奇数长度的 回文
**  处理方法二： 也可用一个循环处理，理论：
** 		长度为 n 的字符串会生成 2n−1 组回文中心[li, ri], 其中 li = i/2, ri = li + i%2
**		这样我们只要从 0 到 2n-2遍历i， 就可以得到所有可能的回文中心，把两种情况统一起来
**		为什么是2*n-1 表示 n + n-1 <== 单个字符为中心的 以及 双字符为中心的 相加
** 		情况类似于 aaa ==> a#a#a 这种情况
** 处理方法三： 更直观的方式奇偶两种情况统一起来，即在每个字符两边加间隔符后，由于 奇 + 偶 = 奇 以及 奇 + 奇 = 奇
** 		得到新字符串必定奇数个，然后利用中心扩展计算个数
*/
// 处理方法二
func countSubstrings(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < 2*n-1; i++{
		// 计算两个端点的初始状态，
		// 单个节点的时候， left 和 right 指向同一个中心, left = i/2  right = left + 0
		// 两个节点的时候， left 和 right 相差一个， left = i/2  right = left + 1
		left := i/2
		right := left + i % 2 //奇偶
		for left >= 0 && right < n && s[left] == s[right]{
			// 往两边扩散
			left--
			right++
			ans++
		}
	}
	return ans
}

// Manacher算法的第一阶段 来求解， 即处理方法三
func countSubstrings_1(s string) int {
	nstr := []byte{'#'}
	n := len(s)
	for i := 0; i < n; i++{
		nstr = append(nstr, s[i], '#')
	}
	ans := 0
	for i := 0; i < len(nstr); i++{
		j := 1
		for i-j >= 0 && i+j < len(nstr) && nstr[i-j] == nstr[i+j]{
			j++
		}
		// 理解 奇偶 直接 j/2
		// 统计答案, 当前贡献为 (f[i] - 1) / 2 向上取整
		ans += j/2
	}
	return ans
}
// 利用在源串中加 前后缀 来省略 下标越界的 条件判断
func countSubstrings_2(s string) int {
	nstr := []byte{'$', '#'}
	n := len(s)
	for i := 0; i < n; i++{
		nstr = append(nstr, s[i], '#')
	}
	nstr = append(nstr, '!')
	ans := 0
	//for i := 0; i < len(nstr); i++{
	for i := 1; i < len(nstr)-1; i++{
		j := 1
		//for i-j >= 0 && i+j < len(nstr) && nstr[i-j] == nstr[i+j]{
		for nstr[i-j] == nstr[i+j]{ // 利用在源串中加 前后缀 来省略 下标越界的 条件判断
			j++
		}
		// 理解 奇偶 直接 j/2
		// 统计答案, 当前贡献为 (f[i] - 1) / 2 向上取整
		ans += j/2
	}
	return ans
}
// 要想降低复杂度，类似于KMP， 充分利用已经得到的信息。
// 已知-1： 回文字符串都是对称的
// 已知-2： 如果一个长回文字符串的对称点左面包含一个小的回文字符串，那么对称过去到右面也必然会包含一个小的回文字符串
/* 已知-3: 回文字符串边界的情况讨论: 观察对称点左面出现的这个小回文字符串，这个字符串有三种情况
** 1. 如果左侧小回文字符串的左边界在大回文字符串的左边界之内，那么右面对称出的小回文字符串也不会触碰到大回文字符串的右边界
	比如“dacaxacad”这个字符串，左侧的“aca”没有超过这个大回文字符串的左边界，那么右面对称出的“aca”也不会超过右边界。
	也就是说，在这种情况下，右面这个小回文字符串的长度与对称的小回文字符串的长度相等，「 绝对不会超过 」这个大回文字符串。
** 2. 如果左侧小回文字符串的左边界超过了大回文字符串的左边界，那这个右面对称出的小回文字符串会正好触碰到大回文字符串的右边界，但是不会超出。
	比如观察这个字符串“dcbabcdxdcbabce”。左侧的小回文字符串「 dcbabcd 」的边界超出了大回文字符串「 cbabcdxdcbabc 」的左边界
	对称过来的右侧的小回文字符串的边界会刚好卡在大回文字符串的右边界。
	这是由于大回文字符串右边界之外的下一个字母（此处是“e”）绝对不是左边界的那个字母“d”的对称，所以右边的小回文字符串延申到边界之后也无法继续延申下去了
	在这种情况下，右面这个小回文字符串的右边界与大回文字符串的右边界相同，那么这个小回文字符串的长度也「绝对不会超过」这个大回文字符串
** 3. 如果左侧小回文字符串的左边界正好卡在大回文字符串的左边界上，那么右面对称出的小回文字符串有可能会继续延伸下去，超过大回文字符串的右边界。
	比如观察这个字符串“abcdcbabcdxdcbabcdxdcb"，左边的小回文字符串「dcbabcd」的左边界正好卡在大回文字符串「dcbabcdxdcbabcd」的左边界上，
	那么对称过来的大回文字符串是有可能继续延申下去的。比如在这个例子中，右面以“a”为对称点的小回文字符串「bcdxdcbabcdxdcb」一直能向右延申到整个字符串的结尾。
	在这种情况下，右面这个小回文字符串的右边界至少与大回文字符串的有边界相同，并且有可能会延申。也就是说这个小回文字符串的长度可能会超过这个大回文字符串
** 为何会这样，因为是从左到右扫描遍历的
** 综合上面三种情况，得知 情况1和2 是不需要再次中心扩展搜寻回文，故而跳过很多字母，情况3 需要继续中心搜索
** 用 f(i) 来表示以 s 的第 i 位为回文中心，可以拓展出的最大回文半径，那么f(i)−1 就是以 i 为中心的最大回文串长度（消除掉#间隔符后的）
** Manacher 算法依旧需要枚举 s 的每一个位置并先假设它是回文中心，但是它会利用已经计算出来的状态来更新 f(i)，而不是向「中心拓展」一样盲目地拓展
** 当我们知道一个 i 对应的 f(i) 的时候，我们就可以很容易得到它的右端点为 i + f(i) - 1
** Manacher 算法如何通过已经计算出的状态来更新 f(i)
** Manacher 算法要求我们维护「当前最大的回文的右端点 rMax」 以及 这个回文右端点对应的回文中心iMax。
** 顺序遍历s， 假设当前遍历的下标为 i。 我们知道在求解 f(i)之前我们应当已经得到了从[1,i-1]所有的 f ， 并且当前已经有了一个最大回文右端点rMax
** 以及它对应的回文中心iMax
**
*/
func countSubstrings_3(s string) int {
	nstr := []byte{'$','#'}
	n := len(s)
	for i := 0; i < n; i++{
		nstr = append(nstr, s[i], '#')
	}
	nstr = append(nstr, '!')
	n = len(nstr)
	f := make([]int, n) // 中间数组
	ans := 0
	iMax := 0  // 这个回文右端点对应的回文中心
	rMax := 0 //当前最大的回文的右端点
	for i := 1; i < n; i++{
		// 初始化f[i]
		if i <= rMax{// 说明 i 被包含在当前最大回文子串内
			f[i] = min(rMax - i + 1, f[2 * iMax - i]) // f[2 * iMax - i] 为 i 以 iMax 为镜像的位置
		}else{
			f[i] = 1 // 不在的话，走正常的中心扩展计算， 这里的 f[i] 类似于countSubstrings_2 中的 j
		}
		// 初始化之后，我们可以保证此时的s[i+f(i)−1]=s[i−f(i)+1],要继续拓展这个区间，
		// 我们就要继续判断s[i+f(i)] 和 s[i - f(i)]s[i−f(i)] 是否相等
		// 如果相等将 f(i) 自增, 否则退出循环
		// 另外需要注意的是不能让下标越界，有一个很简单的办法，就是在开头加一个 $，并在结尾加一个 !，
		// 这样开头和结尾的两个字符一定不相等，循环就可以在这里终止。
		for nstr[i+f[i]] == nstr[i-f[i]]{ // 利用 源字符串加 前后缀 来防止下标越界
			f[i]++ //这里的 f[i] 类似于countSubstrings_2 中的 j
		}
		// 动态维护 iMax 和 rMax 中心和最右边界 同步更新
		if i + f[i] - 1 > rMax{
			iMax = i
			rMax = i + f[i] - 1
		}
		// 统计答案, 当前贡献为 (f[i] - 1) / 2 上取整
		ans += f[i]/2
	}
	return ans
}

/* DP
** 状态：dp[i][j]: 表示字符串s在[i,j]区间的子串是否是一个回文串
** 状态转移： 当 s[i] = s[j] && (j-i < 2 || dp[i+1][j-1]) 时， dp[i][j] = true, 否则为false
** 1. 当只有一个字符时，比如 a 自然是一个回文串
** 2. 当有两个字符时，如果是相等的，比如 aa，也是一个回文串
** 3. 根据子状态推断
 */
func countSubstrings_DP(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp{
		dp[i] = make([]bool, n)
	}
	ans := 0
	for j := 0; j < n; j++{
		for i := 0; i <= j; i++{
			if s[i] == s[j] && ( j - i < 2 || dp[i+1][j-1]){
				dp[i][j] = true
				ans++
			}
		}
	}
	return ans
}

/* 5. Longest Palindromic Substring
** Given a string s, return the longest palindromic substring in s.
 */
/* 中心扩展 */
func longestPalindrome(s string) string {
	ns := []byte{'^', '#'}
	for i := range s{
		ns = append(ns, s[i], '#')
	}
	n := len(ns)
	//f := make([]int, n)
	ns = append(ns, '$')
	ans := []byte{}
	for i := 1; i < n; i++{
		j := 1
		for ns[i-j] == ns[i+j]{
			j++
		}
		if len(ans) < (j-1)*2+1{
			ans = ns[i-j+1:i+j-1]
		}
	}
	//return strings.Replace(string(ans), "#", "", -1)
	//双指针就地删除字符#
	i := 0
	for j := 0; j < len(ans); j++{
		if ans[j] != '#'{
			ans[i] = ans[j]
			i++
		}
	}
	return string(ans[:i])
}
/* 额外构造数组 f， 通过manacher 算法 降低复杂度
** 引入了回文半径或者臂长的概念, 表示 中心扩展算法向外扩展的长度。即如果一个中心位置的最大回文字符串长度为 2 * r + 1, 其半径为 r
** 现定义 i 为当前需要扩展的位置， j 为最近的最大回文串的中心位置， 其半径为 r
** 如果 j + r > i 表示 i 在其覆盖范围内，则可以跳过某些字符，来进行中心扩展判断
** 为了方便处理此情况， 再次引入 i 的 镜像位置，镜像位置计算公式： j - (i-j) => 2*j-i，
** 以 i 为中心的回文串是 大回文串j的 右侧   以 i的镜像为中心的回文串 是 大回文串j 的 左侧
** 设 镜像点(2*j-i)的臂长/半径为 n
** 另外可知 中心点 j 回文串的边界范围是 [j-r, j+r]
情况-1：如果左侧小回文字符串的左边界在大回文字符串的左边界之内，那么右面对称出的小回文字符串也不会触碰到大回文字符串的右边界
		右面这个小回文字符串的长度与对称的小回文字符串的长度相等，「 绝对不会超过 」这个大回文字符串
		判断条件即 j + r > i + n ==> j + r - i > n
情况-2：如果左侧小回文字符串的左边界超过了大回文字符串的左边界，那这个右面对称出的小回文字符串会正好触碰到大回文字符串的右边界，但是不会超出
		右面这个小回文字符串的右边界与大回文字符串的右边界相同，那么这个小回文字符串的长度也「绝对不会超过」这个大回文字符串
		判断条件即 j = r == i + n ==> j + r - i == n
情况-3：如果左侧小回文字符串的左边界正好卡在大回文字符串的左边界上，那么右面对称出的小回文字符串有可能会继续延伸下去，超过大回文字符串的右边界
		右面这个小回文字符串的右边界至少与大回文字符串的有边界相同，并且有可能会延申。也就是说这个小回文字符串的长度可能会超过这个大回文字符串
		判断条件即 j - r == 2*j-i - n ==> j - i + r == n
所有情况 合并：
** 如果 不在覆盖位置， 即 j + r < i 则 继续原来的中心扩展判断, 并记录右半径在最右边的回文字符串，将其中心作为 j ，避免重复计算
 */
func longestPalindrome_Manacher(s string) string {
	ns := []byte{'^', '#'}
	for i := range s{
		ns = append(ns, s[i], '#')
	}
	n := len(ns)
	f := make([]int, n)
	ns = append(ns, '$')
	iMax := 0  // 这个回文右端点对应的回文中心
	rMax := 0 //当前最大的回文的右端点
	ans := []byte{}
	for i := 1; i < n; i++{
		// 初始化f[i]
		if i <= rMax{// 说明 i 被包含在当前最大回文子串内
			f[i] = min(rMax - i + 1, f[2 * iMax - i]) // f[2 * iMax - i] 为 i 以 iMax 为镜像的位置
			// 情况-1 的时候  f[i] 为 i镜像位置的回文半径， 即 f[2 * iMax - i]
			// 情况-2 的时候， f[i] 为 rMax - i + 1
			// 情况-3 的时候， f[i] 为 rMax - i + 1
		}else{
			f[i] = 1 // 不在的话，走正常的中心扩展计算， 这里的 f[i] 类似于countSubstrings_2 中的 j
		}
		j := f[i] // 合并了情况1-3 为 初始化f[i] 从此处继续中心判断
		//j := 1
		for ns[i-j] == ns[i+j]{
			j++
		}

		// 更新 iMax 和 rMax 中心和最右边界 同步更新f[i]的值
		f[i] = j
		if i + f[i] - 1 > rMax{ // 判断是否超过当前的最右边界
			iMax = i
			rMax = i + f[i] - 1
		}
		if len(ans) < (j-1)*2+1{ // i 为中心的 臂长/半径 为 j-1
			ans = ns[i-j+1:i+j-1]
		}
	}
	//return strings.Replace(string(ans), "#", "", -1)
	//双指针就地删除字符#
	i := 0
	for j := 0; j < len(ans); j++{
		if ans[j] != '#'{
			ans[i] = ans[j]
			i++
		}
	}
	return string(ans[:i])
}

/* 471. Encode String with Shortest Length
** Given a string s, encode the string such that its encoded length is the shortest.
** The encoding rule is: k[encoded_string],
** where the encoded_string inside the square brackets is being repeated exactly k times.
** k should be a positive integer.
** If an encoding process does not make the string shorter, then do not encode it.
** If there are several solutions, return any of them.
** 此题目也属于区间DP 问题
** 459.重复的子字符串  --- 找到连续重复的子字符串，我们才能进行编码(压缩)。
 */
/* DP 部分：
** 设s(i,j)表示子串s[i,…,j]。串的长度为len=j-i+1
** 用d[i][j]表示s(i,j)的最短编码串。当s(i,j)有连续重复子串时，s(i,j)可编码为”k[重复子串]”的形式.d[i][j]= "k[重复子串]"
** 当len < 5时，s(i,j)不用编码。d[i][j]=s(i,j)
** 当len > 5时，s(i,j)的编码串有两种可能。
** 现将s(i,j)分成两段s(i,k)和s(k+1,j)，(i <= k <j) 推导出状态方程
** d[i][j] = d[i][k] + d[k+1][j]  当d[i][k].length + d[k+1][j].length < d[i][j].length时
**
** 题目难点部分：快速求出字符串中连续的重复子串
** 枚举子串逐个查找，用上kmp，lcp类的算法，进行加速
** 另有一个方法： 对字符串s，s与s拼接成t=s+s
** 在 t 中 从索引位置 1 开始查找 s 如果查找到，即 在位置 p 处开始， t 中出现了 s
** 注意： t 中肯定可以查找到 s （从索引位置1开始搜索的前提下）
** 当 p >= len(s) 时，说明s中没有连续的重复子串, 不能压缩
** 当 p < len(s) 时，说明s中有连续重复子串并且 连续重复子串是 s[0:p], 重复个数为 len(s) / p
**
 */
func encode(s string) (ans string) {
	n := len(s)
	dp := make([][]string, n)
	for i := range dp{
		dp[i] = make([]string, n)
	}
	var dfs func(start, end int)string
	dfs = func(start, end int) string{
		if start > end { return ""}
		if len(dp[start][end]) > 0{
			return dp[start][end]
		}
		length := end - start + 1
		ss := s[start:end+1]
		if length < 5 { return ss }
		ret := ss // 初始最大
		p := strings.Index((ss+ss)[1:], ss) + 1 // 从索引1开始查找是否有重复子串
		if p > 0 && p < length { // ss 存在重复子串
			ret = fmt.Sprintf("%d[%s]", length/p, dfs(start, start+p-1))
			dp[start][end] = ret
			return ret
		}
		// 动态规划部分
		for mid := start; mid < end; mid++{
			s1 := dfs(start, mid)
			s2 := dfs(mid+1, end)
			if len(s1) + len(s2) < len(ret){
				ret = s1+s2
			}
		}
		dp[start][end] = ret
		return ret
	}
	return dfs(0, n-1)
}
// 区间DP
// dp[i][j] 来自 1. 存在连续的重复子串 2. 分成2段dp[i][k] 和 dp[k+1][j]
func encodeDP(s string) (ans string) {
	n := len(s)
	dp := make([][]string, n)
	for i := range dp{
		dp[i] = make([]string, n)
	}
	for length := 1; length <= n; length++{ // 从长度开始枚举
		for i := 0; i + length <= n; i++{
			j := i + length - 1
			ss := s[i:j+1]
			dp[i][j] = ss
			//if length > 5{
			if length > 4{
				p := strings.Index((ss+ss)[1:], ss) + 1
				if p > 0 && p < len(ss){
					//dp[i][j] = fmt.Sprintf("%d[%s]", len(ss)/p, ss[:p]) 不能直接ss[:p] 可能源串有压缩情况
					dp[i][j] = fmt.Sprintf("%d[%s]", len(ss)/p, dp[i][i+p-1])
				}else{
					for k := i; k < j; k++{ // 注意不要与 切片操作搞混
						//if len(dp[i][k+1]) + len(dp[k+1][j+1]) < len(dp[i][j]){
						if len(dp[i][k]) + len(dp[k+1][j]) < len(dp[i][j]){
							dp[i][j] = dp[i][k] + dp[k+1][j]
						}
					}
				}
			}
		}
	}
	return dp[0][n-1]
}

/* 10. Regular Expression Matching
** Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:
	'.' Matches any single character.​​​​
	'*' Matches zero or more of the preceding element.
** The matching should cover the entire input string (not partial).
*/
/* dp[i][j] 表示 s 前 i 个字符 与 p中前 j个字符是否匹配。根据 p 的第 j 个字符的匹配情况讨论
** 分情况讨论：
** 1. p的j 个字符是一个小写字母，那么我们必须在 s 中匹配一个相同的小写字母
		dp[i][j] = dp[i-1][j-1]  s[i] == p[j]
				 = false		 s[i] != p[j]
** 2. p的第j个字符是 *，那么就表示可以对 p 的第 j-1 个字符匹配任意次数
	2.1 匹配0次情况， dp[i][j] = dp[i][j-2] 即 我们「浪费」了一个字符 + 星号的组合，没有匹配任何 s 中的字符
	2.2 匹配1次的情况， dp[i][j] = dp[i-1][j-2],  if s[i] = p[j-1]
	2.3 匹配2次的情况， dp[i][j] = dp[i-2][j-2],  if s[i-1] = s[i] = p[j-1]
	2.4 匹配x次的情况， dp[i][j] = dp[i-x][j-2],  if s[i-x-1] = s[i-1] = s[i] = p[j-1]
	如果我们通过这种方法进行转移，那么我们就需要枚举这个组合到底匹配了 s 中的几个字符，会增导致时间复杂度增加，并且代码编写起来十分麻烦
	换个角度考虑这个问题：字母 + 星号的组合在匹配的过程中，本质上只会有两种情况：
		1. 匹配 s 末尾的一个字符，将该字符扔掉，而该组合还可以继续进行匹配；
		2. 不匹配字符，将该组合扔掉，不再进行匹配。
	dp[i][j] = dp[i-1][j] || dp[i][j-2]   if s[i] = p[j-1]
			 = dp[i][j-2] 				  if s[i] != p[j-1]
** 3. 在任意情况下，只要 p[j] 是 . ，那么 p[j] 一定成功匹配 s 中的任意一个小写字母
	dp[i][j] = dp[i-1][j-1]
** 总和3类情况得出状态转移方程：
	dp[i][j] = dp[i-1][j] || dp[i][j-2]  matches(s[i], p[j-1]) p[j] = *
			 = dp[i][j-2]  				 p[j] = *
			 = dp[i-1][j-1] 			 matches(s[i], p[j]) p[j] != *
			 = false					 p[j] != *
** 动态规划的边界条件为 dp[0][0]=true，即两个空字符串是可以匹配的
 */
func IsMatch(s string, p string) bool {
	sn, pn := len(s)+1, len(p)+1
	dp := make([][]bool, sn)
	for i := range dp{
		dp[i] = make([]bool, pn)
	}
	match := func(i, j int)bool{
		// 增加溢出判断
		if i < 0 || j < 0 || i >= len(s) || j >= len(p){
			return false
		}
		if s[i] == p[j] || p[j] == '.'{
			return true
		}
		return false
	}
	// 初始化: 主循环是从1开始的，即有元素开始，所以要初始化 s 为空串时，状态计算
	dp[0][0] = true
	for i := 1; i < pn; i++{// 主要*号 会匹配0个
		if p[i-1] == '*'{
			if i < 2{// 首字符是 *
				dp[0][i] = true
			}else{
				dp[0][i] = dp[0][i-2]
			}
		}
	}
	for i := 1; i < sn; i++{
		for j := 1; j < pn; j++{
			if p[j-1] == '*'{
				//if match(i, j-2){
				if match(i-1, j-2){
					dp[i][j] = dp[i-1][j]
					if j >= 2{
						dp[i][j] = dp[i][j] || dp[i][j-2]
					}
				}else{
					if j >= 2{
						dp[i][j] = dp[i][j-2]
					}
				}
			}else{ // 非 *
				if match(i-1, j-1){
					dp[i][j] = dp[i-1][j-1]
				}else{
					dp[i][j] = false
				}
			}
		}
	}
	return dp[sn-1][pn-1]
}

/** 44. Wildcard Matching
** Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*' where:
	'?' Matches any single character.
	'*' Matches any sequence of characters (including the empty sequence).
	The matching should cover the entire input string (not partial).
Constraints:
	0 <= s.length, p.length <= 2000
	s contains only lowercase English letters.
	p contains only lowercase English letters, '?' or '*'.
 */
//2022-01-26 刷出此题
/* 模式 p 中的任意一个字符都是独立的，即不会和前后的字符互相关联，形成一个新的匹配模式。 情况会少
** dp[i][j] 表示字符串s的前 i 个字符和模式 p 的前 j 个字符是否能匹配。
** 考虑模式 p 的第 j 个字符pj, 与之对应的是字符串s 中的 第 i 个字符 si
** 情况1： pj 是字母/问号, si 是字母
		dp[i][j] = dp[i-1][j-1]   pj == si || pj == '?'
				 = false          其他情况
** 情况2：pj 是星号，星号可以匹配零或任意多个小写字母，分两种情况：
	2.1 应用这个星号
		dp[i][j] = dp[i-1][j]
	2.2 不应用这个星号
		dp[i][j] = dp[i][j-1]

边界情况：dp[0][0] = true
		dp[i][0] = false
		dp[0][j] 分情况讨论：星号才能匹配空字符串
		dp[0][j] = true p的前 j 个字符均为星号时
 */
func isWildMatch(s string, p string) bool {
	ns, np := len(s), len(p)
	dp := make([][]bool, ns+1)
	for i := range dp{
		dp[i] = make([]bool, np+1)
	}
	dp[0][0] = true
	for i := 1; i <= np; i++{
		if p[i-1] == '*'{
			if i < 2{
				dp[0][i] = true
				continue
			}
			dp[0][i] = dp[0][i-1]
		}
	}
	for i := 0; i <= ns; i++{
		for j := 1; j <= np; j++{
			if p[j-1] == '*'{
				dp[i][j] = dp[i][j-1]
				/*
				if i > 0{
					dp[i][j] = dp[i][j] || dp[i-1][j-1]
				}*/
				for k := 0; k < i; k++{ // 官方题解优化
					dp[i][j] = dp[i][j] || dp[k][j-1]
				}
			}else{
				if i == 0{
					dp[i][j] = false
					continue
				}
				if s[i-1] == p[j-1] || p[j-1] == '?'{
					dp[i][j] = dp[i-1][j-1]
				}else{
					dp[i][j] = false
				}
			}
		}
	}
	//fmt.Println(dp)
	return dp[ns][np]
}
// 官方题解
func isWildMatchDP(s string, p string) bool {
	sn, pn := len(s), len(p)
	dp := make([][]bool, sn+1)
	for i := range dp{
		dp[i] = make([]bool, pn+1)
	}
	dp[0][0] = true
	for i := 1; i <= pn; i++{
		if p[i-1] == '*'{
			dp[0][i] = true
			continue
		}
		break
	}
	for i := 1; i <= sn; i++{
		for j := 1; j <= pn; j++{
			if p[j-1] == '*'{
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			}else if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[sn][pn]
}
/* 针对 * 的处理 采用贪心枚举
** 模式 p 的种类：
** 1. 模式 p 是以 * 开头 和 以 * 结尾
** 2. 模式 p 不是 以 * 开头
** 3. 模式 p 不是 以 * 结尾
** 将情况 2 和 情况 3 转换为 情况 1
 */
func isWildMatchGreedy(s string, p string) bool {
	ns, np := len(s), len(p)
	// 情况2： 模式p 不是以 * 开头
	i, j := 0, 0
	for ; i < np && j < ns && p[i] != '*'; i, j = i+1, j+1{
		if p[i] == '?' || s[j] == p[i]{
			continue
		}else{
			return false
		}
	}
	s, p = s[j:], p[i:]
	ns, np = len(s), len(p)
	if np == 0{
		return ns == 0
	}
	// 情况3： 模式p 不是以 * 结尾
	i,j = np-1, ns-1
	for ; i >= 0 && j >= 0 && p[i] != '*'; i,j = i-1,j-1{
		if p[i] == '?' || s[j] == p[i]{
			continue
		}else{
			return false
		}
	}
	s, p = s[:j+1], p[:i+1]
	if len(p) == 0{
		return len(s) == 0
	}
	// 情况-1 的处理
	/* 注释的代码存在忽略问号的问题， 必须使用原始的方式
	subs := strings.FieldsFunc(p, func(c rune)bool{ return c == '*'})
	str := s
	for i := range subs{
		pos := strings.Index(str, subs[i]) 忽略 问号
		for pos < len(s){

		}
		if pos < 0{
			return false
		}
		str = str[pos+len(subs[i]):]
	}

	// return p[len(p)-1] == '*' 忽略问号
	 */
	allstart := func(idx int)bool{ // 余下的都是 星号
		for i := idx; i < len(p); i++{
			if p[i] != '*'{
				return false
			}
		}
		return true
	}
	i, j = 0, 0 // i for p  j for s
	for j < len(s) && i < len(p){
		if p[i] == '*'{
			i++
		}else{// s中必须存在p的子串
			k := j
			for ; k < len(s); k++{
				m := k
				n := i
				for ;m < len(s)&& n < len(p) &&p[n] != '*'; m, n = m+1, n+1{
					if s[m] == p[n] || p[n] == '?'{
						continue
					}
					break
				}
				if p[n] == '*'{
					i, j = n, m
					break
				}
			}
			if k >= len(s){
				return false
			}
		}
	}
	return allstart(i)
}