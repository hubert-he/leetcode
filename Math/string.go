package Math

/* 43. Multiply Strings
** Given two non-negative integers num1 and num2 represented as strings,
** return the product of num1 and num2, also represented as a string.
Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.
 */
// 2022-03-07 刷出此题，注意 3个易错点
// 复杂度在 mn*n^2。 mn次乘法，字符串相加操作n次，每次m+n长度字符串， n*(m+n)
func Multiply(num1 string, num2 string) string {
	if num2 == "0" || num1 == "0" { return "0" } // 易错点-3： 特殊情况
	ans := []byte{}
	tailZero, n := 0, len(num2)
	//for i := range num2{
	for i := n-1; i >= 0; i--{
		t := mul(num1, num2[i])
		//fmt.Println(num1, string(num2[i]), string(t))
		for k := 0; k < tailZero; k++{
			t = append(t, '0')
		}
		//fmt.Println(string(ans), string(t))
		ans = add(ans, t)
		//fmt.Println(string(ans), string(t))
		tailZero++
	}
	return string(ans)
}
func mul(a string, b byte)[]byte{
	ret := []byte{}
	x := b - '0'
	n := len(a)
	var carry byte
	//for i := range a{ // 易错点-1： 要逆序
	for i := n-1; i >= 0; i--{
		prod := (a[i] - '0')*x + carry
		carry = prod / 10
		ret = append([]byte{prod % 10 + '0'}, ret...)
	}
	// 易错点-2：遗漏最后的进位
	if carry != 0{
		ret = append([]byte{carry+'0'}, ret...)
	}
	return ret
}
func add(a, b []byte)[]byte{
	var carry byte
	na, nb := len(a), len(b)
	//i, j := 0, 0
	i, j := na-1, nb-1
	ret := []byte{}
	//for i < na && j < nb{
	for i >= 0 && j >= 0{
		sum := a[i]-'0' + b[j] - '0' + carry
		carry = sum / 10
		ret = append([]byte{sum%10 + '0'}, ret...)
		//i, j = i+1, j+1
		i, j = i-1, j-1
	}
	for i >= 0 {
		sum := a[i]-'0' + carry
		carry = sum / 10
		ret = append([]byte{sum%10 + '0'}, ret...)
		i--
	}
	for j >= 0 {
		sum := b[j]-'0' + carry
		carry = sum / 10
		ret = append([]byte{sum%10 + '0'}, ret...)
		j--
	}
	// 易错点-2：遗漏最后的进位
	if carry != 0{
		ret = append([]byte{carry+'0'}, ret...)
	}
	return ret
}

// 方法二：官方解答，复杂度降低为 mn
// 该算法是通过两数相乘时，乘数某位与被乘数某位相乘，与产生结果的位置的规律来完成
// 由于 num1 和 num2 的乘积的最大长度为 m + n，因此创建长度为 m + n 的数组 ansArr 用于存储乘积
// 对于任意 0 <= i < m 和 0 <= j < n， num1[i] * num2[j] 的结果位于 ansArr[i+j+1]
// 如果 ansArr[i+j+1] >= 10 则将进位部分 加到 ansArr[i+j]
// 最后 Trim 掉结果的前缀 0
func Multiply2(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	ansArr := make([]int, m+n)
	for i := m-1; i >= 0; i--{
		x := int(num1[i] - '0')
		for j := n-1; j >= 0; j--{
			y := int(num2[j] - '0')
			ansArr[i+j+1] += x*y
		}
	}
	for i := m+n-1; i > 0; i--{
		ansArr[i-1] += ansArr[i] / 10
		ansArr[i] %= 10
	}
	ans := []byte{}
	idx := 0
	if ansArr[0] == 0{ idx = 1 }
	for ; idx < m+n; idx++{
		ans = append(ans, byte(ansArr[idx])+'0')
	}
	return string(ans)
}
/* 方法三： 把两个数相乘看成是两个多项式相乘
** num1 和 num2 相乘的积的长度为 m+n-1 或 m+n
** 与二进制表示相同， 每个十进制数 都可以转换为 累加和 样式
** C(10) = SUM(c[i] * 10^i) i 从0到m+n-2
** c[i] = SUM(a[k]*b[i-k]) k 从 0 到 i
** 实际推理  10的最高幂次为 m + n - 2
** 123 * 456 ==>  ( 1*10^2 + 2*10^1 + 3*10^0 ) * ( 4*10^2 + 5*10^1 + 6*10^0 )
	==> (1*4)*10^4 + ( 1*5 + 2*4 )*10^3 + ( 1*6 + 2*5 + 3*4 )*10^2 + ( 2*6 + 3*5 )*10^1 + ( 3*6 )*10^0
	==> 4 * 10^4 + 13 * 10^3 + 28 * 10^2 + 27 * 10^1 + 18 * 10^0 	<== 18的十位进位
	==> 4 * 10^4 + 13 * 10^3 + 28 * 10^2 + 28 * 10^1 + 8 * 10^0		<== 28 进位
	==> 4 * 10^4 + 13 * 10^3 + 30 * 10^2 + 8 * 10^1 + 8 * 10^0		<== 30 进位
	==> 4 * 10^4 + 16 * 10^3 + 0 * 10^2 + 8 * 10^1 + 8 * 10^0		<== 16 进位
	==> 5 * 10^4 + 6 * 10^3 + 0 * 10^2 + 8 * 10^1 + 8 * 10^0		<== 无需 进位
	==> 最后结果 56088
** 从而得出算法：
	1. 顺序求解 c[i] 每次O(i) 地选取下标和为 i 的一组(a[k], b[i-k])
	2. 求到 c[i]序列后，再处理进位即可得到答案
** 进阶：c[i] = SUM(a[k]*b[i-k]) k 从 0 到 i
** c[i] 序列起始 是 a[i] 序列与 b[i] 序列的卷积，可以用一种叫做快速傅立叶变换（Fast Fourier Transform, FFT）的方法来加速卷积计算。
** 使得时间复杂度降低到O(clogc), c 是 不小于m+n的最小的 2 的整数幂。
 */
func Multiply3(num1 string, num2 string) string {
	return "快速傅里叶求卷积FTT"
}