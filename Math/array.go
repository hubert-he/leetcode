package Math

import "sort"

/* 2195. Append K Integers With Minimal Sum
** You are given an integer array nums and an integer k.
** Append k unique positive integers that do not appear in nums to nums such that the resulting total sum is minimum.
** Return the sum of the k integers appended to nums.
 */
// 采用字典会超时
func minimalKSum(nums []int, k int) int64 {
	m := map[int]bool{}
	for _, c := range nums{
		m[c] = true
	}
	ans, i := 0, 1
	for k > 0{
		if !m[i]{
			ans += i
			k--
		}
		i++
	}
	return int64(ans)
}
//arithmetic sequence(等差数列) 或arithmetic progression（算术数列）
// 2022-03-08 查看提示 等差数列 刷出此题
func minimalKSum_arithmetic_seq(nums []int, k int) int64 {
	n := len(nums)
	sort.Ints(nums)
	sum, ans := 0, 0
	start := 0
	for i := 1; i <= n; i++{
		if nums[i-1] == start{
			continue
		}
		dist := nums[i-1] - start - 1
		if dist == 0{// 优化-1 元素相邻 直接回去，避免计算
			start = nums[i-1]
			sum = start
			continue
		}
		sum += nums[i-1]
		sn := (nums[i-1]-start+1)*(nums[i-1]+start)/2
		if k >= dist{
			k -= dist
			start = nums[i-1]
			ans += sn - sum
			sum = start
		}
		if k <= 0{
			break
		}
	}
	//fmt.Println(k, start, ans)
	/* 可用等差数列求和优化掉
	   for k > 0{
	       ans += start+k
	       k--
	   }*/
	ans += k*(start+1+start+k)/2 // 优化-2：尾巴等差求和直接计算
	return int64(ans)
}
/* 最大公约数
** 方法一： 穷举 即 从 1 开始枚举直至达到两数较小的，保存更新同时被两数整除的整数，最后的哪个数即为最大公约数
** 方法二： 更相减损术
	任意给定两个正整数，若是偶数，则用2约简
	以较大的数减较小的数，接着把所得的差与较小的数比较，并以大数减小数。
	继续这个操作，直到所得的减数和差相等为止
	a=1
	for num1 % 2 == 0 && num2 % 2 == 0{
		num1, num2 = num1 / 2, num2 / 2
		a *= 2
	}
	if num1 < num2 { num1, num2 = num2, num1 }
	for num2 != 0{ // 第二步
		num1, num2 = num2, num1-num2
	}
** 方法三： 欧几里德算法: 辗转相除法  最优
** 依托的原理：gcd(a, b) = gcd(b, a mod b)， a,b的最大公约数就是b,a mod b的最大公约数
** 算法步骤：
	1. a % b 得余数 r
	2. 若 r == 0， 则 b 即为 GCD
	3. 若 r != 0,  则 a = b, b = r(a%b), 返回步骤1
**
 */
/* 最小公倍数LCM算法
** 方法一：最大公约数方法
	最小公倍数=两整数的乘积 ÷ 最大公约数
** 方法二：乘穷举法
 */
/* 2197. Replace Non-Coprime(非互质的) Numbers in Array
** You are given an array of integers nums. Perform the following steps:
	1. Find any two adjacent numbers in nums that are non-coprime.
	2. If no such numbers are found, stop the process.
	3. Otherwise, delete the two numbers and replace them with their LCM(Least Common Multiple——最小公倍数)
	4. Repeat this process as long as you keep finding two adjacent non-coprime numbers.
** Return the final modified array.
** It can be shown that replacing adjacent non-coprime numbers in any arbitrary order(任意顺序) will lead to the same result.
** The test cases are generated such that the values in the final array are less than or equal to 10^8.
** Two values x and y are non-coprime if GCD(x, y) > 1 where
** GCD(x, y) is the Greatest Common Divisor(最大公约数) of x and y.
 */
// 2022-03-08 错误刷出，不过的case： [287,41,49,287,899,23,23,20677,5,825]
// 	错误答案：	[2009,899,20677,825]   错误原因： gcd(899,20677) = 899
//  正确答案：	[2009,20677,825]
//  错误原因：思维误区： 为考虑到向前的方向, 因为后项们计算lcm后可能会与之前记录的数存在公约数， 导致结果中有些项未能合并
func replaceNonCoprimes_error(nums []int) []int {
	n := len(nums)
	gcd := func(a, b int)int{
		if a < b { a, b = b, a}
		for b != 0{
			a, b = b, a%b
		}
		return a
	}
	coprime := func(a, b int)bool{
		return gcd(a, b) <= 1
	}
	lcm := func(a, b int) int{
		return a*b / gcd(a,b)
	}
	ans := []int{}
	for i := 1; i < n;{
		if i >= n { // 易错点-2：[1] 以及 [517,11,121,517,3,51,3,1887,5] 末尾处理
			ans = append(ans, nums[i-1])
			i++
			continue
		}
		// for !coprime(nums[i-1], nums[i]){ // 易错点-1： 子循环时候，务必确定大条件
		for i < n && !coprime(nums[i-1], nums[i]){
			nums[i] = lcm(nums[i-1], nums[i])
			i++
		}
		//ans = append(ans, nums[i])
		ans = append(ans, nums[i-1])
		i++
	}
	return ans
}
// 基于上面的错误可知， 栈适合此合并运算，运算符为 lcm
func replaceNonCoprimes(nums []int) []int {
	n := len(nums)
	gcd := func(a, b int)int{
		if a < b { a, b = b, a}
		for b != 0{
			a, b = b, a%b
		}
		return a
	}
	coprime := func(a, b int)bool{
		return gcd(a, b) <= 1
	}
	lcm := func(a, b int) int{
		return a*b / gcd(a,b)
	}
	st := []int{}
	for i := 0; i < n; i++{
		if len(st) == 0{
			st = append(st, nums[i])
			continue
		}
		t := nums[i]
		for len(st) > 0{
			top := st[len(st)-1]
			if !coprime(t, top){
				st = st[:len(st)-1]
				t = lcm(top, t)
			}else{
				break
			}
		}
		st = append(st, t)
	}
	return st
}