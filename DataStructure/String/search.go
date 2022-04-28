package String

/* 214. Shortest Palindrome
** You are given a string s. You can convert s to a palindrome by adding characters in front of it.
** Return the shortest palindrome you can find by performing this transformation.
 */
/* 简单思路：枚举 s 的前缀字符串 其中最长的 为目标
** 方法一： Rabin-Karp 字符串哈希算法来找出最长的s的前缀回文串
** 将字符串看成一个 base 进制的数，它对应的十进制值就是哈希值。
** 显然，两个字符串的哈希值相等，当且仅当这两个字符串本身相同。
** 然而如果字符串本身很长，其对应的十进制值在大多数语言中无法使用内置的整数类型进行存储
** 因此，我们会将十进制值对一个大质数 mod 进行取模。此时：
	1. 如果两个字符串的哈希值在取模后不相等，那么这两个字符串本身一定不相同
	2. 如果两个字符串的哈希值在取模后相等，并不能代表这两个字符串本身一定相同
** 然而，我们在编码中使用的 base 和 mod 对于测试数据本身是「黑盒」的，
** 也即 并不存在一组测试数据，使得对于任意的 base 和 mod，都会产生哈希碰撞，导致错误的判断结果
** 一般来说，我们选取一个大于字符集大小（即字符串中可能出现的字符种类的数目）的质数作为 base，
** 再选取一个在字符串长度平方级别左右的质数作为 mod，产生哈希碰撞的概率就会很低。
*/
func shortestPalindrome(s string) string {
	const base, mod = 131, 1e9 + 7
	n := len(s)
	left, right, mul := 0, 0, 1
	best := -1
	// 技巧：从前往后计算 与从后往前计算
	for i := range s{
		left = (left * base + int(s[i] - '0')) % mod
		right = (right + mul * int(s[i]-'0')) % mod
		if left == right { best = i }
		mul = mul * base % mod
	}
	add := ""
	if best != n-1{
		add = s[best+1:]
	}
	b := []byte(add)
	// 逆转b
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1{
		b[i], b[j] = b[j], b[i]
	}
	return string(b) + s
}

/* 方法二： KMP
**
 */
func shortestPalindrome_kmp(s string) string {

}















