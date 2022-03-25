package greedy

import (
	"math"
	"sort"
	"../../utils"
)

/* 330. Patching Array
** Given a sorted integer array nums and an integer n,
** add/patch elements to the array such that any number in the range [1, n] inclusive can be formed by
** the sum of some elements in the array.
** Return the minimum number of patches required.
 */
/* 定理：对于正整数 x，如果区间 [1,x-1] 内的所有数字都已经被覆盖，且 x 在数组中，则区间 [1,2x−1] 内的所有数字也都被覆盖
** 假设数字 x  缺失，则至少需要在数组中补充一个小于或等于 x 的数，才能覆盖到 x， 否则无法覆盖到 x
** 如果区间[1, x-1] 内的所有数字都已经被覆盖，则从贪心的角度考虑，补充 x 之后即可覆盖到 x， 且满足补充的数字个数最少。
** 在补充[1,2x-1]内的所有数字都被覆盖，下一个缺失的数字一定不会小于 2x
** 贪心方案： 每次找到未被数组 nums 覆盖的最小的整数 x , 并在数组中补充 x， 然后寻找下一个未被覆盖的最小的整数，
** 重复上述步骤直到区间[1, n]中的所有数字都被覆盖。
 */
func minPatches(nums []int, n int) int {
	ans := 0
	for i, x := 0, 1; x <= n; {
		if i < len(nums) && nums[i] <= x{
			x += nums[i]
			i++
		}else{
			x *= 2
			ans++
		}
	}
	return ans
}
/* 以[1,5,10]的例子为例:
** 我们从1开始遍历,并且维护一个指向nums的下标.
** >> 一开始是1，而我们看到当前nums数组的第一个元素就是1,所以不需要其他操作.直接跳到2，并且让pos指向nums的第二个元素；
** 现在,我们的目标数是2,但是当前pos指向的数却是5,显然我们只能自己填充一个2,所以让res+1;
** 既然我们已经填过2了,而在2之前可以被覆盖的最长区间长度是1,所以当前可以遍历到的最大区间长度变成了3(即 2 + 1);
** >> 然后,我们可以忽略3,直接跳到4(因为上一步已经知道3在最大覆盖范围内了)。
** 我们发现4同样比当前pos所指向的nums元素小,所以我们得填入4，即让res+1;
** 既然已经填入4了,而我们知道在4之前可以覆盖的连续区间是(1-3),所以当前可以覆盖的最大区间被扩展到了7(即 4 + 3)。
** >> 接下来我们可以直接跳过5、6、7来到8,而当前pos所指向的元素是5,所以当前可覆盖的区间大小又可以加上5了(7+5 = 12),
** 并让pos指向下一个元素
** >> 最后我们跳过了7-12，从13开始遍历，这时候pos所指的元素是10,所以覆盖范围变成了12 + 10 = 22 >20，说明可以完全覆盖指定区间了！
** 到这里大概能够看出端倪 ：
** 我们不断维持一个从1开始的可以被完全覆盖的区间,
** 举个例子,当前可以完全覆盖区间是[1,k]，而当前pos所指向的nums中的元素为B,
** 说明在B之前(因为是升序，所以都比B小)的所有元素之和可以映射到1-----k，
** 而当我们把B也加入进去后，显然，可映射范围一定向右扩展了B个，也就是变成了1---k+B，这就是解题的思路
 */
func minPatches_2(nums []int, n int) int {
	curr_range := 0
	ans := 0
	for x, pos := 1, 0; x <= n;{
		if pos >= len(nums) || x < nums[pos]{
			ans++
			curr_range += x
		}else{
			curr_range += nums[pos]
			pos++
		}
		x = curr_range + 1
	}
	return ans
}

/* 991. Broken Calculator
** There is a broken calculator that has the integer startValue on its display initially.
** In one operation, you can:
	multiply the number on display by 2, or
	subtract 1 from the number on display.
** Given two integers startValue and target,
** return the minimum number of operations needed to display target on the calculator.
 */
/* 逆向
** 除了对 X 执行乘 2 或 减 1 操作之外，我们也可以对 Y 执行除 2（当 Y 是偶数时）或者加 1 操作
** 这样做的动机是我们可以总是贪心地执行除 2 操作：
** 1. 当 Y 是偶数，如果先执行 2 次加法操作，再执行 1 次除法操作，
	我们可以通过先执行 1 次除法操作，再执行 1 次加法操作以使用更少的操作次数得到相同的结果 [(Y+2) / 2 vs Y/2 + 1] 前者多一步
** 2. 当 Y 是奇数，如果先执行 3 次加法操作，再执行 1 次除法操作，
	我们可以将其替代为顺次执行加法、除法、加法操作以使用更少的操作次数得到相同的结果 [(Y+3) / 2 vs (Y+1) / 2 + 1] 前者多一步
 */
/* 正向思维：在X<Y时要实现操作数最小，要将X逼近Y的1/2值或1/4值或1/8值或...再进行*2操作，
** 			难点在于要判断要逼近的是1/2值还是1/4值还是其他值，逻辑复杂
** 逆向思维：在Y>X时Y只管/2，到了Y<X时在+1逼近 说白了就是，正向思维采用的是先小跨度的-1操作，再大跨度的*2操作；
** 			逆向思维采用的是先大跨度的/2操作，再小跨度的-1操作 然而事实上往往是先大后小的解决问题思维在实现起来会比较简单
 */
func brokenCalc(startValue int, target int) int {
	ans := 0
	for target > startValue{
		ans++
		if target & 0x1 == 1 {
			target += 1
		}else {
			target >>= 1
		}
	}
	return ans + startValue - target // startValue > target 只能不断减一
	//return ans + target - startValue
}

/* 极限思想
** startValue < target情况
** cnt1：统计多少乘法
** cnt2：统计多少减法
** 显然我们必须把 startValue 乘到恰好比 target 大的数，否则无法减， 因此先求 cnt1
** 如何求cnt2：
** 假设减法穿插在各个乘法之间，规律
** 1. 如果在第一次乘法前减，那么最终等价于减去2的cnt1次方
** 2. 如果在第二次乘法前减，最终等价于减去2的cnt1-1次方
** 以此类推，由于每次可以减多个1，因此最终要乘个系数，减了a * 2^cnt1 + b * 2^(cnt1-1) + ...
** 那么这个系数 a,b,c等等是多少呢，贪心即可，a越大越好，其次到b, c...
** newtarget - a * 2^cnt1 + b * 2^(cnt1-1) + ... = target
** newtarget 为 多次乘得到的而大于target的
 */
/* 举个例子 startValue = 1   target = 100
** cnt1=7 即128 是 距离100最近的, newtarget = 2^7
** 1*2^7 - 0*2^7 - 0*2^6 - 0*2^5 - 1*2^4 - 1*2^3 - 1*2^2 = target=100
** 转换为计算式为
** ( ( ( 1*2*2*2 - 1) * 2 - 1) * 2 - 1)*2*2 = 100
 */
func brokenCalc2(startValue int, target int) int {
	if startValue > target { return startValue - target }
	cnt1, cnt2 := 0, 0
	for startValue < target{
		startValue <<= 1
		cnt1++
	}
	if startValue == target { return cnt1  }
	// 求 cnt2
	diff := startValue - target
	for i := cnt1; i >= 0; i--{
		t := int(math.Pow(2, float64(i)))
		coeff := diff / t
		diff %= t
		cnt2 += coeff
		if diff == 0 {
			break
		}
	}
	return cnt1+cnt2
}

/* 976. Largest Perimeter Triangle
** Given an integer array nums, return the largest perimeter of a triangle with a non-zero area,
** formed from three of these lengths.
** If it is impossible to form any triangle of a non-zero area, return 0.
 */
// 此方法超时，由下面公式可以发现 k 继续往下走已经没有意义,故继续优化
func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for i := len(nums)-1; i >= 0; i--{
		for j := i-1; j >= 0; j--{
			for k := j-1; k >= 0; k--{
				if nums[k] + nums[j] > nums[i]{
					return nums[k] + nums[j] + nums[i]
				}
			}
		}
	}
	return 0
}

func largestPerimeter2(nums []int) int {
	sort.Ints(nums)
	for i := len(nums)-1; i >= 2; i--{
		if nums[i-1] + nums[i-2] > nums[i]{
			return nums[i-2] + nums[i-1] + nums[i]
		}
	}
	return 0
}

/* 908. Smallest Range I
** You are given an integer array nums and an integer k.
** In one operation, you can choose any index i where 0 <= i < nums.length and change nums[i] to nums[i] + x
** where x is an integer from the range [-k, k]. You can apply this operation at most once for each index i.
** The score of nums is the difference between the maximum and minimum elements in nums.
** Return the minimum score of nums after applying the mentioned operation at most once for each index in it.
 */
/*
** 假设 A 是原始数组，B 是修改后的数组，我们需要最小化 max(B) - min(B)，也就是分别最小化 max(B) 和最大化 min(B)。
** max(B) 最小可能为 max(A) - K， 因为 max(A) 不可能再变得更小。
** min(B) 最大可能为 min(A) + K，
** 所以结果 max(B) - min(B) 至少为 ans = (max(A) - K) - (min(A) + K)
** 我们可以用一下修改方式获得结果（如果 ans >= 0）：
	如果 A[i] <= min(A) + k, 那么 B[i] = min(A)+k
	如果 A[i] >= max(A) - k, 那么 B[i] = max(A)-k
	否则 A[i] = B[i]
** 如果 ans < 0，最终结果会有 ans = 0，同样利用上面的修改方式。
*/
func smallestRangeI(nums []int, k int) int {
	small, big := math.MaxInt32, math.MinInt32
	for _, c := range nums{
		if small > c {
			small = c
		}
		if big < c {
			big = c
		}
	}
	ans := (big - k) - (small + k)
	if ans < 0{
		ans = 0
	}
	return ans
}

/* 910. Smallest Range II
** You are given an integer array nums and an integer k.
** For each index i where 0 <= i < nums.length, change nums[i] to be either nums[i] + k or nums[i] - k.
** The score of nums is the difference between the maximum and minimum elements in nums.
** Return the minimum score of nums after changing the values at each index.
 */
/* 思路： 如最小差值-1 较小的 A[i] 将增加，较大的 A[i] 将变小。
** 题目要求每个元素要么向上移动 K 的距离，要么向下移动 K 的距离，然后要求这个新数组的“最大最小值的距离尽可能地小”。
** 此时最优的策略是把这个数组拆成左右两半，把左边那一半上移K，把右边那一半下移K。 题解中画出了一个坐标图
** 当我们选择在 i 这一点“切一刀”的时候，也就是 A[0] ~ A[i] 的元素都上移，A[i + 1] ~ A[A.length - 1] 的元素都下移。
** 此时 B 点的值是 A[i] + K，D 点的值是 A[A.length - 1] - K。
** 新数组的最大值要么是 B 点要么是 D 点，也就是说新数组的最大值是 Max(A[i] + K, A[A.length - 1] - K)。
** 同样道理，此时 A 点的值是 A[0] + K，C 点的值是 A[i + 1] - K。
** 新数组的最小值要么是 A 点要么是 C 点，也就是说新数组的最小值是 Min(A[0] + K, A[i + 1] - K)。
** 因此，题目需要的“新数组的最大值和最小值的差值”，就是 Max(A[i] + K, A[A.length - 1] - K) - Min(A[0] + K, A[i + 1] - K)。
** 挨个遍历一下所有可能的 i 的值，然后取上面算式的最小值即可
 */

// 忽视了 d1 d2 有可能负数的情况
func smallestRangeII_error(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	ans := nums[n-1] - nums[0]
	//在分割为2部分的 两个端点出计算
	// 左边是+k， 而 右边是 -k 以使得最大 最下 差值最小
	for i := 1; i < n-1; i++{
		// 情况1： k 很小的时候
		d1 := (nums[n-1] - k) - (nums[0] + k)
		//if d1 < 0 { d1 = -d1 }
		// 情况2： k 很大的时候
		d2 := (nums[i-1] + k) - (nums[i] - k)
		//if d2 < 0 { d2 = -d2 }
		ans = utils.Min(ans, d1, d2)
	}
	return ans
}

func smallestRangeII(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	ans := nums[n-1] - nums[0]
	// 不断的分割为 2 部分，然后分情况计算最大值，最小值
	// 整体排序后，前部分+k， 后部分-k，尽量缩小最大最小的差额
	for i := 1; i < n; i++{
		maxium := utils.Max(nums[i-1]+k, nums[n-1]-k)
		minium := utils.Min(nums[0]+k, nums[i]-k)
		ans = utils.Min(ans, maxium-minium)
	}
	return ans
}














