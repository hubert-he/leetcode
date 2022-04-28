package TwoPointer

import "math"

/* 633. Sum of Square Numbers
** Given a non-negative integer c, decide whether there're two integers a and b such that a^2 + b^2 = c.
Constraints:
	0 <= c <= 2^31 - 1
 */
// 2022-04-28 TLE
func judgeSquareSum_tle(c int) bool {
	for i := 0; i <= c; i++{
		target := c - i*i
		l, r := 0, c+1
		for l < r{
			mid := int(uint(l+r)>>1)
			prod := mid*mid
			if prod > target{
				r = mid
			}else if prod < target{
				l = mid+1
			}else{
				return true
			}
		}
	}
	return false
}
// 第一个循环 没必要全部 遍历
func judgeSquareSum_bs(c int) bool {
	for i := 0; i*i <= c; i++{ // TLE 原因
		target := c - i*i
		l, r := 0, c+1
		for l < r{
			mid := int(uint(l+r)>>1)
			prod := mid*mid
			if prod > target{
				r = mid
			}else if prod < target{
				l = mid+1
			}else{
				return true
			}
		}
	}
	return false
}

/* 双指针思路
** 不失一般性，可以假设 a <= b, a=0  b=根号c
** 1. 若 a*a + b*b == c 返回True
** 2. 若 a*a + b*b > c  b = b-1 继续查找
** 3. 若 a*a + b*b < c  a = a+1 继续查找
** 结束条件 a == b
** 为什么会这样，画出矩阵图后，就类似于240. 搜索二维矩阵 II 的那样
** 每一列从上到下升序，每一行从左到右升序。
 */
func judgeSquareSum(c int) bool {
	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right{
		prod := left * left + right * right
		if prod == c { return true }
		if prod > c {
			right -= 1
		}else{
			left += 1
		}
	}
	return false
}

/* 此题还有 费马平方和定理
** 一个非负整数 c 如果能够表示为两个整数的平方和，当且仅当 c 的所有形如 4k+3 的质因子的幂均为偶数。
** 对 c 进行质因数分解，再判断所有形如4k+3 的质因子的幂是否均为偶数即可
 */
func judgeSquareSum_math(c int) bool {
	for base := 2; base*base <= c; base++{
		// 如果不是因子，枚举下一个
		if c % base > 0{ continue }
		// 计算 base 的幂
		exp := 0
		for ; c%base == 0; c /= base{
			exp++
		}
		// 根据 Sum of two squares theorem 验证
		if base%4 == 3 && exp % 2 != 0{
			return false
		}
	}
	// 例如 11 这样的用例，由于上面的 for 循环里 base * base <= c ，base == 11 的时候不会进入循环体
	// 因此在退出循环以后需要再做一次判断
	return c%4 != 3
}
