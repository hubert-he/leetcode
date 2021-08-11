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