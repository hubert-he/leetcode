package Array

import (
	"math/bits"
	"sort"
)

/* 1356. Sort Integers by The Number of 1 Bits
** You are given an integer array arr.
** Sort the integers in the array in ascending order by the number of 1's in their binary representation and
** in case of two or more integers have the same number of 1's you have to sort them in ascending order.
** Return the array after sorting it.
** Constraints:
	1 <= arr.length <= 500
	0 <= arr[i] <= 10^4
 */
// 2022-02-21 刷出此题
func sortByBits(arr []int) []int {
	sort.Ints(arr)
	m := make([][]int, 32)
	bits := func(x int)(cnt int){
		for x != 0{
			cnt++
			x &= (x-1)
		}
		return
	}
	for i := range arr{
		t := bits(arr[i])
		m[t] = append(m[t], arr[i])
	}
	ans := []int{}
	for i := range m{
		ans = append(ans, m[i]...)
	}
	return ans
}

// 方法二： 借助题目的限制, 提升性能
/* 计算数字中1的个数，乘以10000000（题目中不会大于 10^4）然后加上原数字，放入数组中，
** 并对数组进行排序，最后 % 10000000 获取原来的数组，填充到原数组返回即可。
** 将 统计值放入高位，低位为原数 以此来排序
 */

func sortByBits2(arr []int) []int {
	const high int = 1e7
	m := make([]int, len(arr))
	for i := range arr{
		m[i] = bits.OnesCount32(uint32(arr[i])) * high + arr[i]
	}
	sort.Ints(m)
	for i := range m {
		m[i] %= high
	}
	return m
}

// 方法三 ：借助go 库，
// 以及采用递推预处理，
// 定义 bit[i] 为数字 i 二进制表示下数字 1 的个数，得到递推公式： bit[i] = bit[i>>1] + (i & 1)
func sortByBits3(arr []int) []int {
	const limit int = 1e4+1
	bit := make([]int, limit)
	for i := 1; i < limit; i++{
		bit[i] = bit[i>>1] + i & 1
	}
	less := func(i, j int) bool{
		x, y := arr[i], arr[j]
		cx, cy := bit[x], bit[y]
		return cx < cy || cx == cy && x < y
	}
	sort.Slice(arr, less)
	return arr
}

/* 249. Group Shifted Strings
** We can shift a string by shifting each of its letters to its successive letter.
** For example, "abc" can be shifted to be "bcd".
** We can keep shifting the string to form a sequence.
** For example, we can keep shifting "abc" to form the sequence: "abc" -> "bcd" -> ... -> "xyz".
** Given an array of strings strings, group all strings[i] that belong to the same shifting sequence.
** You may return the answer in any order.
 */
// 2022-02-21 刷出此题， 需要增加测试case 的多样性思考，比如存在重复的字符串情况
func groupStrings(strings []string) [][]string {
	//m := map[string]bool{} // 需要统计词频，考虑 ["a", "a"] 这种情况
	m := map[string]int{}
	for i := range strings{
		m[strings[i]]++
	}
	ans := [][]string{}
	for i := range strings{
		sub := []string{}
		for j := 0; j <= 26; j++{
			t := []byte(strings[i])
			for k := range t{
				//t[k] = (t[k] + byte(j))%('a'+26) ❌
				t[k] = (t[k]+byte(j))%26 + 'a'
			}
			//if m[string(t)] > 0{
			for m[string(t)] > 0{
				sub = append(sub, string(t))
				m[string(t)]--
			}
		}
		if len(sub) > 0{
			ans = append(ans, sub)
		}
	}
	return ans
}
