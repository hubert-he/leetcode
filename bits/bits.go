package bits

import "math/bits"
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