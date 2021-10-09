package bits

import (
	"math/bits"
)
func max(nums ...int)int{
	ans := nums[0]
	for i := 1; i < len(nums); i++{
		if nums[i] > ans {
			ans = nums[i]
		}
	}
	return ans
}
// 190 Reverse Bits
const (
	m1 = 0x55555555
	m2 = 0x33333333
	m4 = 0x0F0F0F0F
	m8 = 0x00FF00FF
	m16 = 0x0000FFFF
)
// 分治，自底向上
func reverseBits(num uint32) uint32 {
	n := num
	n = n >> 1 & m1 | n & m1 << 1
	n = n >> 2 & m2 | n & m2 << 2
	n = n >> 4 & m4 | n & m4 << 4
	n = n >> 8 & m8 | n & m8 << 8
	n = n >> 16 & m16 | n & m16 << 16
	return n
}
// 231 Power of Two
/* 	方法一： n & (n - 1) ： 如果 nn 是正整数并且 n & (n - 1) = 0，那么 n 就是 2 的幂
	方法二：n & (-n) ：如果 n 是正整数并且 n & (-n) = n, n就是2的幂
		其中-n是n的相反数，是一个负数。该位运算技巧可以直接获取n二进制表示的最低位的1.
    	由于负数是补码规则在计算机存储的，-n 的二进制表示为n的二进制表示的每一位取反再加上1，因此原理：
  		假设n的二进制表示为(a10...0)2 其中a表示若干个高位，1表示最低位的那个1, 0...0表示后面的若干0
  		进行按位与运算，高位全部变为 0，最低位的 1 以及之后的所有 0 不变，这样我们就获取了 n 二进制表示的最低位的 1
	方法三：判断是否为最大 2 的幂的约数
	在题目给定的32位有符号整数的范围内，最大的2的幂为2^30 = 1073741824。只需判断n是否是2^30的约数即可。
*/
func isPowerOfTwo(n int) bool{
	return n > 0 && (n & (n-1)) == 0
}
func isPowerOfTwo2(n int) bool {
	return n > 0 && n & -n == n
}
func isPowerOfTwo3(n int) bool {
	const big = 1 << 30
	return n > 0 && big%n == 0
}

// 342. Power of Four
/* 方法一： 二进制表示中1的位置
	如果 n 是 4 的幂，那么 n 的二进制表示中有且仅有一个 1，并且这个 1 出现在从低位开始的第偶数个二进制位上（这是因为这个 1 后面必须有偶数个 0）
    这里我们规定最低位为第 0 位，例如 n=16 时，n 的二进制表示为(10000)_2  唯一的 1 出现在第 4 个二进制位上，因此 n 是 4 的幂
​	由于题目保证了 n 是一个 32 位的有符号整数，因此我们可以构造一个整数 mask，它的所有偶数二进制位都是 0，所有奇数二进制位都是 1。
    这样一来，我们将 n 和  mask 进行按位与运算，如果结果为 0，说明 n 二进制表示中的 1 出现在偶数的位置，否则说明其出现在奇数的位置。
	根据上面的思路，mask 的二进制表示为：
		mask = (10101010101010101010101010101010)_2
		mask = (AAAAAAAA)_{16}  或者 （2AAAAAAA)_{16}
	方法二：取模性质
	如果 n 是 4 的幂，那么它可以表示成 4^x 的形式，我们可以发现它除以 3 的余数一定为 1，即：
		4^x == (3+1)^x == 1^x == 1  (mod 3)
	如果n是2的幂却不是4的幂，那么它可以表示成4^x*2的形式，此时除3的余数一定为2.因此我们可以判断通过n 除以 3的余数是否为1来判断n是否是4的幂
 */
func isPowerOfFour(n int) bool {
	return n > 0 && n&(n-1) == 0 && n&0xaaaaaaaa == 0
}
func isPowerOfFour2(n int) bool{
	return n > 0 && n & (n-1) == 0 && n % 3 == 1
}

/*	461. Hamming Distance
	方法1：使用内置函数
		大多数编程语言都内置了计算二进制表达中 1 的数量的函数:
        C/C++: __builtin_popcount(x^y);
		JAVA: Integer.bitCount(x^y);
		GO: bits.OnesCount(uint(x^y))
	方法2: >> 使用
	方法3：Brian Kernighan 算法
		核心概念：记 f(x) 表示 x 和 x-1进行 & 运算的结果, 即 f(x) = x & (x-1), 那么f(x)恰好为 x 删去其二进制表示中最右侧的 1 的结果。
      基于该概念，当计算出 s = x ^ y后，只需要不断让 s = f(s)， 直到 s = 0 即可。这样每循环一次， s 都会删去其二进制表示中 最右侧的 1
       最终循环的次数即为 s 的二进制表示中 1 的数量。
 */
func hammingDistance(x, y int) (ans int) {
	return bits.OnesCount(uint(x^y))
}
func hammingDistance2(x int, y int) int {
	ans := 0
	for z := x ^ y; z > 0; z = z >> 1 {
	/*	bad idea
		if z & 1 == 1{
			ans++
		}
	 */
		ans += z & 1
	}
	return ans
}

func hammingDistance3(x, y int) (ans int) {
	s := x ^ y
	for s != 0 {
		ans++
		s = s & (s - 1)
	}
	return ans
}

/* 191. Number of 1 Bits
Write a function that takes an unsigned integer and returns the number of '1' bits it has (also known as the Hamming weight).
Note:
Note that in some languages, such as Java, there is no unsigned integer type. In this case, the input will be given as a signed integer type. It should not affect your implementation, as the integer's internal binary representation is the same, whether it is signed or unsigned.
In Java, the compiler represents the signed integers using 2's complement notation. Therefore, in Example 3, the input represents the signed integer. -3.
Example 1:
	Input: n = 00000000000000000000000000001011
	Output: 3
	Explanation: The input binary string 00000000000000000000000000001011 has a total of three '1' bits.
Example 2:
	Input: n = 00000000000000000000000010000000
	Output: 1
	Explanation: The input binary string 00000000000000000000000010000000 has a total of one '1' bit.
Example 3:
	Input: n = 11111111111111111111111111111101
	Output: 31
	Explanation: The input binary string 11111111111111111111111111111101 has a total of thirty one '1' bits.
Constraints: The input must be a binary string of length 32.
Follow up: If this function is called many times, how would you optimize it?
*/
func hammingWeight(num uint32) int {
	ans := 0
	for num > 0{
		ans ++
		num = num & (num - 1)
	}
	return ans
}

/* 05.03. Reverse Bits LCCI
   You have an integer and you can flip exactly one bit from a 0 to a 1.
   Write code to find the length of the longest sequence of 1s you could create.
 */
// 05.03. Reverse Bits LCCI
/*
	自己实现的
*/
func reverseBits0503(num int) int {
	bits := [33]int{}
	for i := 32; i > 0; i--{
		if num & 0x1 == 0x1{
			bits[i] = 1
		}
		num >>= 1
	}
	first := true // 起始计算
	prev, curr := 0, 0
	ans := 0
	for i := 0; i <= 32; i++{
		if bits[i] == 1{
			curr++
		}else{
			if first {
				ans = max(ans, prev + curr)
			} else {
				ans = max(ans, prev + curr + 1)
			}
			prev = curr
			curr = 0
		}
	}
	return ans
}

func reverseBits05032(num int) int {
	cur := 0 // 当前位置为止连续1的个数，遇 0 归 0， 遇 1 加 1
	insert := 0 // 在当前位置变1， 往前数连续 1 的最大个数， 遇到 0 变为 cur + 1， 遇到 1 加 1
	ans := 1
	for i := 0; i < 32; i++{
		//if num & (1 << i) == 1{  低级错误
		if (num & (1 << i)) != 0 {
			cur++
			insert++
		}else{
			insert = cur + 1
			cur = 0
		}
		ans = max(ans, insert)
	}
	return ans
}
/* 476. Number Complement 求补码
The complement of an integer is the integer you get when you flip all the 0's to 1's and all the 1's to 0's in its binary representation.
For example, The integer 5 is "101" in binary and its complement is "010" which is the integer 2.
Given an integer num, return its complement.
注意是非负整数  和 正整数 两种情况
 */
/*
变量				二进制
num				00001101
mask			11110000
~mask ^ num		00000010
 */
func FindComplement(num int) int {
	if num == 0{
		return 1
	}
	// 取反: 利用异或 5的二进制是101 取反就是010 实际上就是101和111的异或运算
	// 111的求解过程就是获取5最高为1总位数3的全为1的的过程
	n := 1
	t := num
	for t > 0{
		t >>= 1
		n <<= 1
	}
	n = n - 1// 全反转为1
	return n ^ num
}

/* 405. Convert a Number to Hexadecimal
  Given an integer num, return a string representing its hexadecimal representation.
  For negative integers, two’s complement method is used.
  All the letters in the answer string should be lowercase characters,
  and there should not be any leading zeros in the answer except for the zero itself.
  Note: You are not allowed to use any built-in library method to directly solve this problem.
 */
/*
	知识点1： 计算机内部就是补码
	知识点2：goLang 中没有逻辑左/右移， 只有算术左/右移。
			即 sign = a >> 31 如果a是正数 sign结果为0， 如果a是负数 sign结果为-1
			用算术移位实现逻辑移位： 当a 为负数时， 把右移时高位补充的1变0就可以了：
				sign = a >> 31 & 1 这样a为负数时得到的sign结果为1， 正数时为 0
 */
func ToHex(num int) string {
	if num == 0{ // 首先要排除特殊情况
		return "0"
	}
	m := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	ans := []byte{}
	for num != 0{
		t := num & 0xF
		ans = append([]byte{m[t]}, ans...)
		num >>= 4 // golang 只有算术右移，故需要额外一步去掉符号位
		num &= 0xFFFFFFF
	}
	return string(ans)
}

/* 371. Sum of Two Integers
** Given two integers a and b, return the sum of the two integers without using the operators + and -.
 */
/* 对于整数 a 和 b：
在不考虑进位的情况下，其无进位加法结果为 a 异或 b。
而所有需要进位的位为 a & b，进位后的进位结果为 ( a&b ) << 1。
因此 可以将整数a和b的和，拆分为 a 和 b的无进位加法结果与进位结果的和
因为每一次拆分都可以让需要进位的最低位至少左移一位，又因为 a 和 b
可以取到负数
计算过程：
从b开始
1. a b 求无进位加法
2. 求 a b 的进位
3. 第一次 a b 无进位加法的值 与 第一次产生的进位值 进行 无进位加法
 */
func GetSum(a int, b int) int {
	for b != 0{
		carry := (a&b) << 1
		a ^= b
		b = carry
	}
	return a
}
// 更容易理解的版本
func GetSum2(a int, b int) int {
	carry := (a&b)<<1
	a ^= b
	for carry != 0{
		b = carry
		carry = (a&b)<<1
		a ^= b
	}
	return a
}
// 此题的引申 求中位值
// 使用的上题原理：a + b = c;  ==>  (a与b的无进位加法) + (a 与 b 的进位值) = c
// 二分查找为了避免溢出都会写left + (right - left)/2，这里增加另外一种编码方法
// 在golang sort 库函数中 求 mid 为了防止溢出：mid := int(uint(a+b) >> 1)
func GetMid(a, b int)int{
	// (进位 + 无进位和)/2 => ((a & b)<<1 + (a ^ b))>>1,化简可得
	return (a & b) + (a^b)>>1
}


