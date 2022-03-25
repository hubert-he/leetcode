package Math

import (
	"fmt"
	"math"
	"math/big"
	"sort"
	"strconv"
)

/* 数学相关的题目：
** 1.
*/

/* 66. Plus One
Given a non-empty array of decimal digits representing a non-negative integer, increment one to the integer.
The digits are stored such that the most significant digit is at the head of the list,
and each element in the array contains a single digit.
You may assume the integer does not contain any leading zero, except the number 0 itself.
Example 1:
	Input: digits = [1,2,3]
	Output: [1,2,4]
	Explanation: The array represents the integer 123.
Example 2:
	Input: digits = [4,3,2,1]
	Output: [4,3,2,2]
	Explanation: The array represents the integer 4321.
Example 3:
	Input: digits = [0]
	Output: [1]
 */
func PlusOne(digits []int) []int {
	s := len(digits)
	if s <= 0{
		return nil
	}
	carry := 1
	for i := s - 1; i >= 0; i--{
		digits[i] += carry
		if digits[i] > 9{
			carry = digits[i] / 10
			digits[i] %= 10
		}else{
			carry = 0
			break
		}
	}
	if carry > 0{ // 溢出
		digits = append([]int{carry}, digits...)
	}
	return digits
}
/* 67. Add Binary
Given two binary strings a and b, return their sum as a binary string.
 */
// 2022-03-08 刷出此题方法，
// 易错点：
//	1. 逆序逆序
//	2. carry 算好了
//	3. 计算到最后一定要注意是否有进位
func AddBinary(a string, b string) string {
	na, nb := len(a), len(b)
	ans := []byte{}
	i, j := na-1, nb-1
	carry := '0'
	for i >= 0 && j >= 0{
		if carry == '0'{
			if a[i] == '1' && b[j] == '1'{
				carry = '1'
				ans = append([]byte{'0'}, ans...)
			}else if a[i] == '1' || b[j] == '1'{
				ans = append([]byte{'1'}, ans...)
			}else{
				ans = append([]byte{'0'}, ans...)
			}
		}else{
			if a[i] == '1' && b[j] == '1'{
				ans = append([]byte{'1'}, ans...)
			}else if a[i] == '1' || b[j] == '1'{
				ans = append([]byte{'0'}, ans...)
			}else{
				ans = append([]byte{'1'}, ans...)
				carry = '0'
			}
		}
		i, j = i-1, j-1
	}
	for i >= 0{
		if carry == '0'{
			ans = append([]byte{a[i]}, ans...)
		}else{
			if a[i] == '0'{
				carry = '0'
				ans = append([]byte{'1'}, ans...)
			}else{
				ans = append([]byte{'0'}, ans...)
			}
		}
		i--
	}
	for j >= 0{
		if carry == '0'{
			ans = append([]byte{b[j]}, ans...)
		}else{
			if b[j] == '0'{
				ans = append([]byte{'1'}, ans...)
				carry = '0'
			}else{
				ans = append([]byte{'0'}, ans...)
			}
		}
		j--
	}
	if carry == '1'{
		ans = append([]byte{'1'}, ans...)
	}
	return string(ans)
}

func AddBinary2(a, b string) (ans string){
	carry := 0
	lenA, lenB := len(a)-1, len(b)-1
	n := lenA
	if n < lenB{
		n = lenB
	}
	for i := 0; i <= n; i++{
		if i <= lenA {
			carry += int(a[lenA - i] - '0')
		}
		if i <= lenB {
			carry += int(b[lenB - i] - '0')
		}
		ans = strconv.Itoa(carry % 2) + ans
		carry /= 2
	}
	if carry > 0{
		ans = "1" + ans
	}
	return
}
/* 好难理解！！！！！！！！
  使用bit运算
  把a 和 b 转换成整型数字x 和 y，接下来 x保存结果  y保存进位
  当进位不为0时：
	计算当前 x 和 y 的无进位相加结果： answer = x ^ y
    计算 x 和 y 的进位： carry = (x & y) << 1
    完成本次循环，更新 x = answer   y = carry
   最后返回 x 的 二进制形式
   原理：
     在第一轮计算中，
		answer的最后一位是 x 和 y 相加后的结果
		carry的倒数第二位是 x 和 y 最后一位相加的进位。
	接着每一轮中，由于carry是由 x 和 y 按位 与 并且左移得到的，空位会有 0 来填补，所以下面的计算过程中后面的数位不受影响
    而每一轮都可以得到一个低 i 位的答案 和 它向低 i+1位的进位，也就模拟了加法过程
 */
// 此方法仅供学习加法与位运算原理
func AddBinaryBit(a, b string) string {
	x, _ := strconv.ParseInt(a, 2, 64)
	y, _ := strconv.ParseInt(b, 2, 64)
	fmt.Println(x, y)
	for y != 0 {
		answer := x ^ y
		carry := (x & y) << 1
		x, y = answer, carry
	}
	return strconv.FormatInt(x, 2)
}
func AddBinaryBitBig(a, b string) string {
	x, y := new(big.Int), new(big.Int)
	x.SetString(a, 2)
	y.SetString(b, 2)
	big0 := new(big.Int)
	for y.Cmp(big0) > 0{
		answer,carry := new(big.Int), new(big.Int)
		answer = answer.Xor(x, y)
		carry = carry.Lsh(carry.And(x, y), 1)
		x, y = answer, carry
	}
	//fmt.Println(fmt.Sprintf("%b", x))
	return fmt.Sprintf("%b", x)
}

/* 69. Sqrt(x)
Given a non-negative integer x, compute and return the square root of x.
Since the return type is an integer, the decimal digits are truncated, and only the integer part of the result is returned.
Note: You are not allowed to use any built-in exponent function or operator, such as pow(x, 0.5) or x ** 0.5.
*/
// Binary Search
func MySqrt(x int) int {
	l, r := 0, x // 设定上下界
	ans := -1
	for l <= r{
		mid := l + (r - l) / 2
		if mid * mid <= x{
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
// 牛顿迭代法是一种可以用来快速求解函数零点的方法
func MySqrt2(x int) int {
	if x == 0{
		return 0
	}
	C, x0 := float64(x), float64(x)
	for {
		xi := 0.5*(x0 + C / x0)
		if math.Abs(x0 - xi) < 1e-7{
			break
		}
		x0 = xi
	}
	return int(x0)
}

/* 168. Excel Sheet Column Title
	进制转换 https://leetcode-cn.com/problems/excel-sheet-column-title/solution/excelbiao-lie-ming-cheng-by-leetcode-sol-hgj4/
	number' = (number - a0) / 26 = ((number - a0) + (a0 - 1)) / 26 = (number - 1) / 26
 */
func convertToTitle(columnNumber int) string {
	ans := []byte{}
	for columnNumber > 0{
		columnNumber--
		ans = append([]byte{'A'+byte(columnNumber%26)}, ans...)
		columnNumber /= 26
	}
	return string(ans)
}
/* 204. Count Primes
 */
func CountPrimes(n int) int {
	ans := 0
	for i := 2; i < n; i++{
		if isPrime2(i){
			ans++
		}
	}
	return ans
}
func isPrime(num int)bool{
	if num == 1{
		return false
	}
	t := int(math.Sqrt(float64(num)))
	for k := 2; k <= t; k++{
		if num % k == 0{
			return false
		}
	}
	return true
}
func isPrime2(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
/* 埃氏筛: 希腊数学家厄拉多塞  复杂度 O(nlogn)
   如果 x 是质数，那么大于 x 的 x 的倍数 2x,3x,… 一定不是质数.
   设isPrime[i]表示 i 是不是素数， 如果是则为true 否则false. 从小到大遍历每个数，如果这个数为质数，则将其所有的倍数都标记为合数，即false。
	这样运行结束后，就可以知道素数个数。
 */
func CountPrimesEratosthenes(n int) (cnt int){
	isPrime := make([]bool, n)
	for i := range isPrime{
		isPrime[i] = true
	}
	for i := 2; i < n; i++{
		if isPrime[i]{
			cnt++
			for j := 2 * i; j < n; j += i{ // 2i 3i 4i ....
				isPrime[j] = false
			}
		}
	}
	return
}

/* 507. Perfect Number
A perfect number is a positive integer that is equal to the sum of its positive divisors, excluding the number itself.
A divisor of an integer x is an integer that can divide x evenly.
Given an integer n, return true if n is a perfect number, otherwise return false.
*/
// TimeOut
func checkPerfectNumber(num int) bool {
	t := num/2
	sum := 0
	for i := 1; i <= t; i++{
		if num % i == 0{
			sum += i
		}
	}
	return sum == num
	/* 辣鸡
	if sum == num{
		return true
	}
	return false
	 */
}
/* 优化-1：
   从1至根号n 进行枚举。如果n有一个大于根号n的因数x， 那么它一定有一个小于根号n的因数 n/x。
   因此从1-根号n枚举n因子，当出现一个n的因数时，还需要加上 n/x.
   另外特殊情况，即 x == n/x情况
 */
func CheckPerfectNumber(num int) bool {
	if num == 1{
		return false
	}
	sum := 1
	for i := 2; i*i <= num; i++{ // 学习根号n的策略
		if num % i == 0{
			sum += i
			if i*i != num{
				sum += num / i
			}
		}
	}
	return sum == num
}
/* 欧几里得-欧拉定理
   每个偶数是完全数都可以写成 2^(p-1) * (2^p - 1)的形式，其中p为素数， 2^p -1 也是素数，称为梅森素数
	例如 6 = 2^1 * (2^2 - 1)
		28 = 2^2 * (2^3 - 1)
	由于目前奇完全数还未被发现，因此所有的完全数都可以写成上述形式。当n不超过 10的8次方时， p 也不会很大
    因此我们只要带入最小的若干个梅森素数 2, 3, 5, 7, 13, 17, 19, 312,3,5,7,13,17,19,31（形如 2^p - 1的素数，p是素数），
	将不超过10的8次方 的所有完全数计算出来即可
 */
func CheckPerfectNumberBest(num int) bool {
	primes := []int{2, 3, 5, 7, 13, 17, 19, 31} // 梅森素数
	pn := func(p int)int{
		return (1 << (p-1)) * ((1 << p) - 1)
	}
	for _, p := range primes{
		if pn(p) == num{
			return true
		}
	}
	return false
}

/* 172. Factorial Trailing Zeroes
	Given an integer n, return the number of trailing zeroes in n!.
	Follow up: Could you write a solution that works in logarithmic time complexity?
Example 2:
	Input: n = 5
	Output: 1
	Explanation: 5! = 120, one trailing zero.
Example 3:
	Input: n = 0
	Output: 0
	计算因子中5的个数
 */
func TrailingZeroes(n int) int {
	if n >= 5{
		return n/5 + TrailingZeroes(n/5)
	}else{
		return 0
	}
}

/* 1952. Three Divisors
Given an integer n, return true if n has exactly three positive divisors. Otherwise, return false.
An integer m is a divisor of n if there exists an integer k such that n = k * m.
*/
/* 方法一：遍历[1,n] 闭区间内的所有正整数
** 向上取整：比自己大的最小整数
** 向下取整：比自己小的最大整数
** 四舍五入：更接近自己的整数
** 方法二：内含数学： 对于任意一个大于等于根号n(向下取整)的正除数 x, n/x也是n的正除数，且一定小于等于根号n的向下取整
** 方法是只需遍历[1, 根号n向下取整]区间内的所有正整数，便可以统计出整数n的正除数数目。如果 x 被 n 整除，
** 那么 x 与 n/x 都是 n 的正除数。此时需要根据 x 与 n/x是否相等来确定新增的正除数数目。即
** 1. x == n/x  新增的数目为 1
** 2. x != n/x  新增的数目为 2
此题需要注意的两点：
1. 通过 i * i <= n 来求根号n
2. x 与 n/x是否相等来确定新增的正除数数目
*/
func IsThree(n int) bool {
	cnt := 0
	for i := 1; i*i <= n; i++{
		if n % i == 0{
			if i != n/i{
				cnt += 2
			}else{
				cnt++
			}
		}
		if cnt > 3{
			return false
		}
	}
	return cnt == 3
}
/* 1945. Sum of Digits of String After Convert
** You are given a string s consisting of lowercase English letters, and an integer k.
First, convert s into an integer by replacing each letter with its position in the alphabet (i.e., replace 'a' with 1, 'b' with 2, ..., 'z' with 26).
Then, transform the integer by replacing it with the sum of its digits. Repeat the transform operation k times in total.
For example, if s = "zbax" and k = 2, then the resulting integer would be 8 by the following operations:
Convert: "zbax" ➝ "(26)(2)(1)(24)" ➝ "262124" ➝ 262124
Transform #1: 262124 ➝ 2 + 6 + 2 + 1 + 2 + 4 ➝ 17
Transform #2: 17 ➝ 1 + 7 ➝ 8
Return the resulting integer after performing the operations described above.
 */
/*注意此方法未考虑到 数字溢出 */
func getLucky(s string, k int) int {
	num := 0
	for i := 0; i < len(s); i = i+1{
		n := int(s[i] - 'a' + 1)
		if n > 9{
			num *= int(math.Pow10(2))
		}else{
			num *= int(math.Pow10(1))
		}
		num += n
	}
	for i := 0; i < k; i++{
		t := 0
		for num > 0{
			t += num % 10
			num /= 10
		}
		num = t
	}
	return num
}
/*Trick: 注意到k最小从1开始计算，故可在溢出地方直接计算出一次transform 避免溢出*/
func GetLucky(s string, k int) int {
	var num rune // int 与 int32
	for _, c := range s{
		n := c - 'a' + 1
		num += n/10 + n % 10 // 提前算出第一次transform
	}
	for ;k > 1 && num > 9; k--{// 扣掉第一次，num如果小于9 可直接退出
		var sum rune
		for { // do ...  while 循环
			sum += num % 10
			if num /= 10; num == 0{
				break
			}
		}
		num = sum
	}
	return int(num)
}

/* 279. Perfect Squares
Given an integer n, return the least number of perfect square numbers that sum to n.
A perfect square is an integer that is the square of an integer; in other words, it is the product of some integer with itself.
For example, 1, 4, 9, and 16 are perfect squares while 3 and 11 are not.
限制条件： 1<= n <= 10000
贪心算法不适合
*/
/* 数学 -- 四平方和定理
	1. -- 任意一个正整数都可以被表示为至多四个正整数的平方和
	2. -- 当且仅当 n != 4^k * (8m + 7) 时， n可以被表示为至多三个正整数的平方和。
	      因此当 n == 4^k * (8m + 7)时， n只能被表示成 4 个正整数的平方和，可以直接返回4
当  n != 4^k * (8m + 7)， 情况可能为 1， 2， 3
answer=1： 则必有n 为完全平方数
answer=2： 则有 n = a^2 + b^2， 只需要枚举所有的 a ([1, 根号n]), 判断 n - a^2 是否为完全平方数即可
answer=3： 排除法。
*/

func NumSquaresMath(n int) int {
	checkAnswer4 := func(x int)bool{
		for x % 4 == 0{
			x /= 4
		}
		return x % 8 == 7
	}
	isPerfectSquare := func(x int)bool {
		y := int(math.Sqrt(float64(x)))
		return y*y == x
	}
	if isPerfectSquare(n) {
		return 1
	}
	if checkAnswer4(n) {
		return 4
	}
	for i := 1; i * i <= n; i++{
		j := n - i*i
		if isPerfectSquare(j) {
			return 2
		}
	}
	return 3
}
/* 650. 2 Keys Keyboard
** There is only one character 'A' on the screen of a notepad. You can perform two operations on this notepad for each step:
	Copy All: You can copy all the characters present on the screen (a partial copy is not allowed).
	Paste: You can paste the characters which are copied last time.
** Given an integer n, return the minimum number of operations to get the character 'A' exactly n times on the screen.
*/
/* 对一个数 分解质因子， 这里学习 分解一个数 为素数的乘积
** 另外此题也可以用DP 解决
*/
func minSteps(n int) int {
	ans := 0
	// 分解质因子: 对 n 进行质因数分解，统计所有质因数的和，即为最终的答案
	for i := 2; i * i <= n; i++{
		for n%i == 0 {
			fmt.Println(i)
			n /= i
			ans += i
		}
	}
	// 8 的话 8 = 2*2*2*1
	// 例如：986 分解为  2 * 17 * 29
	fmt.Println(n) // 985 分解为 197 * 5
	if n > 1 {
		ans += n
	}
	return ans
}

/* 372. Super Pow
** Your task is to calculate ab mod 1337
** where a is a positive integer and b is an extremely large positive integer given in the form of an array.
Constraints:
	1 <= a <= 2^31 - 1
	1 <= b.length <= 2000
	0 <= b[i] <= 9
	b does not contain leading zeros.
 */
/** 前提数学知识
** 1. 快速幂计算
** 2. 模与乘法的关系：乘法在取模的意义下满足分配律 (a*b)mod m = [ ( a mod m) * (b mod m)] mod m
 */
/** 方法一： 倒序遍历
** 设 a 的幂次为 n。根据题意，n 从最高位到最低位的所有数位构成了数组 b。记数组 b 的长度为 m
** 公式-1： n = SUM( 10^m-1-i ) * bi  // i 从0 到 m-1  例如： n = 123 ==> n = 1 * 10^2 + 2 * 10^1 + 3 * 10^0
** 公式-2： a^(x+y) = a^x * a^y  以及  a^(x*y) = (a^x)^y
** 联合公式 1和2，得出 a^n = MUL( (a^(10^m-1-i))^bi  ) i 从 0 到 m-1
** 例如 a^123 = a^(10^2)*a^1  * a^(10^1)*a^2 * a^(10^0)*a^3
** a的10的k次方 = a 的 10的k-1次方 的 10次方
 */
func superPow(a int, b []int) int {
	const mod = 1337
	pow := func(x, n int) int{
		res := 1
		for ;n > 0; n >>= 1{
			if n&1 == 1{
				res = (res*x)%mod
			}
			x = x * x % mod
		}
		return res
	}
	ans := 1
	for i := len(b)-1; i >= 0; i--{
		ans = ans * pow(a, b[i]) % mod
		a = pow(a, 10)
	}
	return ans
}
/* 方法二：秦九韶算法 正序遍历
** 秦九韶算法是一种将一元n次多项式的求值问题转化为n个一次式的算法
** 一般地，一元n次多项式的求值需要经过(n+1)*n/2次乘法和n次加法，而秦九韶算法只需要n次乘法和n次加法
** 例如 a的123次方
** 1234 = 123 * 10 + 4 = (12*10+3) * 10 + 4
	    = ( (1*10+2) * 10 + 3 ) * 10 + 4
		= ( ( (0*10 + 1) * 10 + 2) * 10 + 3 ) * 10 + 4
*/
func superPow_qinjiuzhao(a int, b []int) int {
	const mod = 1337
	pow := func(x, n int) int{
		res := 1
		for ;n > 0; n >>= 1{
			if n&1 == 1{
				res = (res*x)%mod
			}
			x = x * x % mod
		}
		return res
	}
	ans := 1
	for _, e := range b{
		ans = pow(ans, 10) * pow(a, e) % mod
	}
	return ans
}
/* 96. Unique Binary Search Trees
** Given an integer n, return the number of structurally unique BST's (binary search trees) which has exactly n nodes of unique values from 1 to n.
 */
/* 此题可以DP 解决
** 这里来学习 Catalan number 即卡特兰数 或者 明安图数（蒙古族数学家明安图）是组合数学中一种常出现于各种计数问题中的数列
*/
func numTrees(n int) int {
	return 0
}

/* 剑指 Offer 62. 圆圈中最后剩下的数字
** 约瑟夫环问题
** 0,1,···,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字（删除后从下一个数字开始计数）。求出这个圆圈里剩下的最后一个数字。
** 例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。
** 限制：
	1 <= n <= 10^5
	1 <= m <= 10^6
 */
// 可以转换为链表处理，但是数据量大
/** 拆分为较小子问题
** 如果我们知道对于一个长度 n-1的序列，留下的是第几个元素，那么我们就可以由此计算出长度为 n 的序列的答案 <== 如何 推出？？
** ** 我们将上述问题建模为函数 f(n, m)，该函数的返回值为最终留下的元素的序号。
** 1. 长度为 n 的序列会先删除第 m % n 个元素，然后剩下一个长度为 n-1的序列。那么可以求解 f(n-1, m)，就可以知道对于剩下的n-1个元素，
** 	  最终会留下第几个元素，设答案为 x = f(n-1, m)
** 2. 关键的一步，当我们知道了 f(n-1, m)对应的答案 x 之后，
**    也就可以知道，长度为 n 的序列最后一个删除的元素，应当是从 m%n 开始数的第 x 个元素。
** 3. 当序列长度为 1 时， 返回的编号肯定为 0
**    即 f(n, m) = （m % n + x) % n = (m+x) % n
 */
func lastRemaining(n int, m int) int {
	if n == 1{ return 0 }
	x := lastRemaining(n-1, m)
	return (m+x) % n
}
// 尾递归转迭代 DP
/* DP
  「1, 3」问题:	f(1) = 0					3
  「2, 3」问题:	f(2) = (f(1) + m) % 2		1	3
  「3, 3」问题:	f(3) = (f(2) + m) % 3		1	3	4
  「4, 3」问题:	f(4) = (f(3) + m) % 4		3	4	0	1
  「5, 3」问题:	f(5) = (f(4) + m) % 5		0	1	2	3	4
 */
func lastRemainingDP(n int, m int) int {
	ans := 0
	for i := 2; i <= n; i++{
		ans = (ans + m) % i
	}
	return  ans
}


/* 812. Largest Triangle Area
** Given an array of points on the X-Y plane points where points[i] = [xi, yi],
** return the area of the largest triangle that can be formed by any three different points.
** Answers within 10-5 of the actual answer will be accepted.
 */
// 不好处理的一个因素 是 计算三角形面积
/* 根据三角形的三个顶点计算出面积的方法有很种
** 1. 鞋带公式：
	Shoelace公式，也叫高斯面积公式，是一种数学算法，可求确定区域的一个简单多边形的面积
	该多边形是由它们顶点描述笛卡尔坐标中的平面.
	用户交叉相乘相应的坐标以找到包围该多边形的区域,并从周围的多边形中减去该区域以找到其中的多边形的区域
	之所以称为鞋带公式，是因为对构成多边形的坐标进行恒定的交叉乘积，就像系鞋带一样
** 2. 海伦公式:
	海伦公式又译作希伦公式、海龙公式、希罗公式、海伦－秦九韶公式。利用三角形的三条边的边长直接求三角形面积的公式
	S = 对 p(p-a)(p-b)(p-c) 开根号  p = (a+b+c)/2
** 3. 普通计算公式： S = 1/2 * a * b * sin(C)
 */
func largestTriangleArea(points [][]int) float64 {
	n := len(points)
	var ans float64
	area := func(P, Q, R []int)float64{
		return 0.5 * math.Abs(float64(P[0]*Q[1] + Q[0]*R[1] + R[0]*P[1] - P[1]*Q[0] - Q[1]*R[0] - R[1]*P[0]))
	}
	for i := 0; i < n; i++{
		for j := i+1; j < n; j++{
			for k := j+1; k < n; k++{
				a := area(points[i], points[j], points[k])
				if ans < a{
					ans = a
				}
			}
		}
	}
	return ans
}

/* 1502. Can Make Arithmetic Progression(等差数列) From Sequence
** A sequence of numbers is called an arithmetic progression if the difference between any two consecutive elements is the same.
** Given an array of numbers arr,
** return true if the array can be rearranged to form an arithmetic progression. Otherwise, return false.
 */
// 2022--2-14 刷出此题， 学习等差数列高效的判别方法： 2*ai = (ai-1 + ai+1)
func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	for i := 1; i < len(arr)-1; i++{
		if 2*arr[i] != (arr[i-1]+arr[i+1]){
			return false
		}
	}
	return true
}

/* 1232. Check If It Is a Straight Line
** You are given an array coordinates, coordinates[i] = [x, y], where [x, y] represents the coordinate of a point.
** Check if these points make a straight line in the XY plane.
 */
// 共线向量线性相关
/* 在给定的点集中，以任意一点 P0 为基准，如果其他所有点的, deltaY / deltaX 是不变的，那么点集内所有的点在同一条直线上。
** 但是这种做法会涉及到除数为 0的问题，即垂直于 x 轴的直线需要单独判断， 另外还会涉及浮点数运算。这是斜率方法
** 我们可以把点集中除了 P0 之外的点 Pi 都看成以 P0 为起点，以 Pi 为终点的向量，记为 vi， 并选择 v1 作为基准。
** 如果其他向量都与 v1 共线，那么点集合所有的点都共线
** 线性代数：如果二维向量 A 与 B 共线，那么他们线性相关，且有   |A, B| = 0  即它们拼成的二阶矩阵的行列式为 0。
 */
// 官方题解：将第一个节点 平移到坐标原点，经过两点的直线有且仅有一条，可以 通过前2个节点来确定这条直线
// 没有想到  平移策略
func checkStraightLine(coordinates [][]int) bool {
	deltaX, deltaY := coordinates[0][0], coordinates[0][1]
	for _, p := range coordinates{// 将第一个节点 平移到坐标原点， 直线过原点，故方程 A * x + B * y = 0
		p[0] -= deltaX
		p[1] -= deltaY
	}
	// 本质上是斜率： 除法转乘法
	A, B := coordinates[1][1], -coordinates[1][0]
	for _, p := range coordinates[2:]{
		x, y := p[0], p[1]
		if A*x + B*y != 0{
			return false
		}
	}
	return true
}

func checkStraightLine2(coordinates [][]int) bool {
	deltaX, deltaY := coordinates[0][0], coordinates[0][1]
	for _, p := range coordinates{// 将第一个节点 平移到坐标原点， 直线过原点，故方程 A * x + B * y = 0
		p[0] -= deltaX
		p[1] -= deltaY
	}
	// 本质上是斜率： 除法转乘法
	A, B := coordinates[1][1], coordinates[1][0]
	for _, p := range coordinates[2:]{
		x, y := p[0], p[1]
		if A*x != B*y{
			return false
		}
	}
	return true
}

/** 增补知识 全排列
** 剔除重复元素
 */
// DFS回溯
func permutation_DFS(nums []int) [][]int{
	n := len(nums)
	ans := [][]int{}
	// 易错点：repeat 判断不能从 0 开始查找，因为递归的是子数组
	// 查找范围 [start, end)
	isRepeat := func(start, end, value int)bool{
		//for i := 0; i < end; i++{
		for i := start; i < end; i++{
			if nums[i] == value{
				return true
			}
		}
		return false
	}
	var dfs func(start int)
	dfs = func(start int){
		if start == n{
			t := make([]int, n)
			copy(t, nums)
			ans = append(ans, t)
		}
		// 方法一：使用map来去重
		repeated := map[int]bool{}
		for i := start; i < n; i++{
			// 方法二：遍历前面元素来去重
			if isRepeat(start, i, nums[i]){
				continue
			}
			if repeated[nums[i]]{
			//	continue  // 方法一：使用map来去重
			}
			repeated[nums[i]] = true
			nums[start], nums[i] = nums[i], nums[start]
			dfs(start+1)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
	dfs(0)
	return ans
}
/* 全排列的非递归实现
** 非递归实现，也就是通常所说的字典序全排列算法
** 如何求某一个排列紧邻着的后一个字典序
** 设P是1～n的一个全排列: P=p1p2......pn = p1p2......pj−1 pj pj+1......pk−1 pk pk+1......pn，
** 求该排列的紧邻着的后一个字典序的步骤如下：
	1. 从排列的右端开始，找出第一个比右边数字小的数字的序号j，即 j=max(i|pi<pi+1)
	2. 在pj的右边的数字中，找出所有比pj大的数中最小的数字pk
	3. 交换 pj 与 pk
	4. 再将pj+1......pk−1pkpk+1...pnpj+1......pk−1pkpk+1...pn倒转得到排列
		p’=p1.....pjpn.....pk+1pkpk−1.....pj+1p1.....pjpn.....pk+1pkpk−1.....pj+1，这就是排列p的下一个排列
** 例如：p=839647521是数字1～9的一个排列。下面生成下一个排列的步骤如下：
	1. 自右至左找出排列中第一个比右边数字小的数字4
	2. 在该数字后的数字中找出比4大的数中最小的一个5
	3. 将5与4交换，得到839657421
	4. 将7421反转，得到839651247。这就是排列p的下一个排列。
** 以上这4步，可以归纳为：一找、二找、三交换、四翻转。
** 和C++的 STL库中的next_permutation()函数（#include<algorithm>）原理相同。
*/
func permutation(nums []int) [][]int{
	// 首先全序列排序，找到最小的字典序，然后依次遍历查找
	n := len(nums)
	sort.Ints(nums)
	next_permutation := func()bool{
		//1. 从排列的右端开始，找出第一个比右边数字小的数字的序号j
		j := -1
		for i := n-2; i >= 0; i--{
			if nums[i] < nums[i+1]{
				j = i
				break
			}
		}
		if j == -1{ return false }  // 未找到，不存在下一个
		//2. 在pj的右边的数字中，找出所有比pj大的数中最小的数字pk 注意：pk > pj
		k := -1
		for i := j+1; i < n; i++{
			// if nums[i] < nums[j]{ break } 易错点：必须<=, 考虑重复元素
			if nums[i] <= nums[j]{ break }
			k = i
		}
		if k < 0{ return false }
		//3. 交换 pj 与 pk
		nums[j], nums[k] = nums[k], nums[j]
		//4. 将[j+1,n]反转
		for o, p := j+1, n-1; o < p; o,p = o+1, p-1{
			nums[o], nums[p] = nums[p], nums[o]
		}
		return true
	}
	ans := [][]int{}
	for {
		t := make([]int, n)
		copy(t, nums)
		ans = append(ans, t)
		fmt.Println(t)
		if !next_permutation(){
			break
		}
	}
	return ans
}


