package array

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

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
func AddBinary(a string, b string) string {
	ans := []byte{}
	left := 0
	for sa,sb := len(a) - 1, len(b) - 1; sa >= 0 || sb >= 0;{
		if sa < 0 {
			if b[sb] == '1'{
				if left > 0{
					ans = append([]byte{'0'}, ans...)
					left = 1
				}else{
					ans = append([]byte{'1'}, ans...)
					left = 0
				}
			}else{
				if left > 0{
					ans = append([]byte{'1'}, ans...)
				}else{
					ans = append([]byte{'0'}, ans...)
				}
				left = 0
			}
		}else if sb < 0{
			if a[sa] == '1'{
				if left > 0{
					ans = append([]byte{'0'}, ans...)
					left = 1
				}else{
					ans = append([]byte{'1'}, ans...)
					left = 0
				}
			}else{
				if left > 0{
					ans = append([]byte{'1'}, ans...)
				}else{
					ans = append([]byte{'0'}, ans...)
				}
				left = 0
			}
		}else {
			if a[sa] == '1' && b[sb] == '1'{
				if left > 0{
					ans = append([]byte{'1'}, ans...)
				}else{
					ans = append([]byte{'0'}, ans...)
				}
				left = 1
			}else if a[sa] == '1' || b[sb] == '1'{
				if left > 0{
					ans = append([]byte{'0'}, ans...)
					left = 1
				}else{
					ans = append([]byte{'1'}, ans...)
					left = 0
				}
			}else{
				if left > 0{
					ans = append([]byte{'1'}, ans...)
				}else{
					ans = append([]byte{'0'}, ans...)
				}
				left = 0
			}
		}
		sa--
		sb--
	}
	if left > 0{
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
