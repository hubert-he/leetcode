package BinarySearch

import (
	"math"
	"sort"
	"strings"
)

/* 1060. Missing Element in Sorted Array
** Given an integer array nums which is sorted in ascending order and all of its elements are unique and given also an integer k,
** return the kth missing number starting from the leftmost number of the array.
 */
// 2021-11-24 刷出此题
// 找到一个最小的区间 [l, r] 包含答案
func MissingElement(nums []int, k int) int {
	n := len(nums)
	// 处理边界
	dis := nums[n-1] - nums[0] - n + 1
	if dis < k{
		return nums[n-1] + k - dis
	}
	i, j := 0, n-1
	for j - i > 1{
		mid := int(uint(i+j)>>1)
		d := (nums[mid] - nums[i]) - (mid - i)
		if k > d{
			k -= d
			i = mid
		}else{
			j = mid
		}
	}
	return nums[i]+k
}
// 2021-12-20 刷出此题
func missingElement(nums []int, k int) int {
	i, j := 0, len(nums)
	diff := (nums[j-1] - nums[i]) - (j - i - 1)
	if k > diff{
		return nums[j-1] + k - diff
	}
	for i <= j {
		mid := int(uint(i+j)>>1)
		diff := (nums[mid] - nums[0]) - (mid - 0)
		//fmt.Println(mid, diff)
		if diff < k{
			i = mid + 1
		}else if diff > k{
			j = mid - 1
		}else{
			j = mid - 1
		}
	}
	//fmt.Println("=>", i, j, nums[i])
	return nums[j] + k - (nums[j] - nums[0] - j)
}

// 官方题解
func MissingElement2(nums []int, k int) int {
	n := len(nums)
	missing := func(idx int)int{
		return nums[idx] - nums[0] - idx
	}
	if k > missing(n-1){ // 边界，k超出nums最大可能值
		return nums[n-1] + k - missing(n-1)
	}
	left, right := 0, n-1
	// 找到left等于right，即 missing(left - 1) < k <= missing(left)
	for left < right{
		mid := int(uint(left+right)>>1)
		if missing(mid) < k{
			left = mid+1
		}else{
			right = mid
		}
	}
	// kth missing number is larger than nums[idx - 1] and smaller than nums[idx]
	return nums[left-1] + k - missing((left-1))
}

/*1901. Find a Peak Element II
** A peak element in a 2D grid is an element that is strictly greater than all of its adjacent neighbors to the left, right, top, and bottom.
** Given a 0-indexed m x n matrix mat where no two adjacent cells are equal, find any peak element mat[i][j] and return the length 2 array [i,j].
** You may assume that the entire matrix is surrounded by an outer perimeter with the value -1 in each cell.
** You must write an algorithm that runs in O(m log(n)) or O(n log(m)) time.
** Note: No two adjacent cells are equal.
 */
// 二分思路： 相邻元素各不相同
func findPeakGrid(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	maxRow := func(r int)(c int){
		if r < 0 || r >= m{
			return -1
		}
		for i := 1; i < n; i++{
			if mat[r][i] > mat[r][c]{
				c = i
			}
		}
		return
	}
	i, j := 0, m-1
	ans := []int{}
	for i <= j{
		mid := int(uint(i+j) >> 1)
		up, cur, down := maxRow(mid-1), maxRow(mid), maxRow(mid+1)
		upVal, curVal, downVal := -1, -1, -1
		if up != -1{
			upVal = mat[mid-1][up]
		}
		if cur != -1{
			curVal = mat[mid][cur]
		}
		if down != -1{
			downVal = mat[mid+1][down]
		}
		// 中间行最大，并且又是行内最大值，所以找到 直接 return
		if curVal >= upVal && curVal >= downVal{ // 注意 等号
			ans = append(ans, mid, cur)
			break
		}
		if upVal > curVal && upVal > downVal{
			j = mid - 1
		}else{
			i = mid + 1
		}
	}
	return ans
}

/* 1231. Divide Chocolate
** You have one chocolate bar that consists of some chunks. Each chunk has its own sweetness given by the array sweetness.
You want to share the chocolate with your k friends so you start cutting the chocolate bar into k + 1 pieces using k cuts, each piece consists of some consecutive chunks.
Being generous, you will eat the piece with the minimum total sweetness and give the other pieces to your friends.
Find the maximum total sweetness of the piece you can get by cutting the chocolate bar optimally.
*/
func MaximizeSweetness(sweetness []int, k int) int {
	var dfs func(arr []int, part int)[][][]int
	dfs = func(arr []int, part int)(ret [][][]int){
		n := len(arr)
		if n == 0 || part == 0{
			return
		}
		if n < part{
			return
		}
		if part == 1 {
			ret = append(ret, [][]int{arr})
			return
		}
		if n == part{
			t := [][]int{}
			for i := range arr{
				t = append(t, []int{arr[i]})
			}
			ret = append(ret, t)
			return
		}
		for i := 0; n - i >= part ; i++{
			sub := dfs(arr[i+1:], part-1)
			//fmt.Println(len(sub))
			for j := range sub{
				sub[j] = append(sub[j], arr[:i+1])
				ret = append(ret, sub[j])
			}
		}
		return
	}
	all := dfs(sweetness, k+1)
	ans := 0
	/*
	   for i := range all{
	       fmt.Println(all[i])
	   }
	*/
	for i := range all{
		t := math.MaxInt32
		for j := range all[i]{
			s := sum(all[i][j]...)
			if t > s{
				t = s
			}
		}
		ans = max(ans, t)
	}
	return ans
}
func sum(nums ...int)int{
	m := 0
	for _, c := range nums{
		m += c
	}
	return m
}

func MaximizeSweetnessBS(sweetness []int, k int) int {
	sum := 0
	minVal := math.MaxInt32
	for i := range sweetness{
		sum += sweetness[i]
		if minVal > sweetness[i]{
			minVal = sweetness[i]
		}
	}
	if k == 0{
		return sum
	}
	// 计算sweet最小甜度块时， 能够切出块的最大数量
	calccount := func(sweet int)int{
		total := 0
		cnt := 0
		for i := range sweetness{
			total += sweetness[i]
			if total >= sweet{
				cnt++
				total = 0
			}
		}
		return cnt
	}
	// 求 right_bound
	left, right := minVal, sum
	for left <= right{ // 搜索终止： [right+1, right] [left,left-1]
		mid := int(uint(left+right)>>1)
		count := calccount(mid)
		if k+1 < count{ // 证明还有变大的可能
			left = mid + 1
		}else if k+1 > count{
			right = mid - 1
		}else{
			left = mid + 1
		}
	}
	return right
}

/* 287. Find the Duplicate Number
** Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.
** There is only one repeated number in nums, return this repeated number.
** You must solve the problem without modifying the array nums and uses only constant extra space.
 */
/* 此题可以使用 floyd 快慢指针 同时也可bit 现在用二分
** 定义 cnt[i] 表示 nums 数组中小于等于 i 的数有多少个
** 假设我们重复的数是 target，那么[1,target−1]里的所有数满足[target,n] 里的所有数满足 cnt[i]>i，具有单调性
 */
func findDuplicate(nums []int) int {
	n := len(nums)
	l, r := 1, n-1
	ans := -1
	for l <= r{
		mid := int(uint(l+r)>>1)
		cnt := 0
		for i := range nums{
			if nums[i] <= mid{
				cnt++
			}
		}
		if cnt <= mid{
			l = mid + 1
		}else{// cnt > mid 结果可能在 [l, r]中
			r = mid - 1
			ans = mid
		}
	}
	return ans
}
// 另外一种写法
func findDuplicate2(nums []int) int {
	n := len(nums)
	l, r := 1, n-1
	for l < r{ // 搜索空间 [l, r) == > 终止条件 [r, r)
		mid := int(uint(l+r)>>1)
		cnt := 0
		for i := range nums{
			if nums[i] <= mid{
				cnt++
			}
		}
		if cnt <= mid{
			l = mid + 1
		}else{// cnt > mid 结果可能在 [l, r]中
			r = mid
		}
	}
	return r
}

/* 4. Median(中位数) of Two Sorted Arrays
** Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
** The overall run time complexity should be O(log (m+n)).
 */
/* 不需要合并两个有序数组，只要找到中位数的位置即可。
** 不变量：由于两个数组的长度已知，因此中位数对应的两个数组的下标之和也是已知的
** 维护两个指针，初始时分别指向两个数组的下标 0 的位置，每次将指向较小值的指针后移一位，直到到达中位数的位置
** 如果一个指针已经到达数组末尾，则只需要移动另一个数组的指针
** 当m + n 为奇数，median = 两个有序数组中的第(m+n)/2个元素，为偶数， median = (m+n)/2 与 (m+n)/2 + 1 的平均值
** 这道题可以转化成寻找两个有序数组中的第 k 小的数， 其中 k 为 (m+n)/2 or (m+n)/2 + 1
** 1. 如果A[k/2-1] < B[k/2-1] 则 比 A[k/2-1]小的数最多只有A的前 k/2-1个数和B的前k/2-1个数。
	即 比 A[k/2-1] 小的数 最多只有 k-2个，因此 A[k/2-1] 不可能是第 k 个数，因此 A[0] 至 A[k/2-1] 都不可能是 第 k 个数，全部排除
** 2. 如果 A[k/2-1] > B[k/2-1], 同 1 可排除 B[0] 至 B[k/2-1]
** 3. 如果A[k/2-1] == B[k/2-1], 则可以归为第一种情况处理
** 比较完 A[k/2-1] 和 B[k/2-1] 后，可以排除 k/2个不可能是第k小的数，同时继续对剩余的新数组进行二分，并根据排除的数的个数，减少k的值
** 后续缩小范围，有3个特殊情况：
** 1. 如果A[k/2-1] 或 B[k/2-1] 越界，那么我们可以选取对应数组中的最后一个元素。
	在这种情况下，我们必须根据排除数的个数减少k的值，而非直接减去k/2
** 2. 如果一个数组为空，说明该数组中的所有元素都被排除，我们可以直接返回另一个数组中第k小的数
** 3. 如果 k=1，我们只要返回两个数组首元素的最小值即可
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	median := (m+n)/2
	p1, p2 := 0, 0
	for p1 < m && p2 < n{
		k := median / 2
		if k < m && k < n{
			if nums1[k] < nums2[k]{
				p1 += k
			}else{
				p2 += k
			}
		}else if k < m{
			if nums1[k] < nums2[n-1]{
				p1 += k
			}else{

			}
		}else if k < n{

		}
	}
	return 0
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	getKthNum := func(k int)int{// k 为从 1 开始的序号 表示 个数
		idx1, idx2 := 0, 0
		for {
			if idx1 == len(nums1){
				return nums2[idx2 + k - 1]
			}
			if idx2 == len(nums2){
				return nums1[idx1 + k -1]
			}
			if k == 1{
				return min(nums1[idx1], nums2[idx2])
			}
			half := k/2
			newIdx1 := min(idx1+half, len(nums1)) - 1
			newIdx2 := min(idx2+half, len(nums2)) - 1
			if nums1[newIdx1] < nums2[newIdx2]{
				k -= newIdx1 - idx1 + 1
				idx1 = newIdx1 + 1
			}else  if nums1[newIdx1] > nums2[newIdx2]{
				k -= newIdx2 - idx2 + 1 // 扣掉 k/2 个
				idx2 = newIdx2 + 1
			}else{ // equal
				k -= newIdx1 - idx1 + 1
				idx1 = newIdx1 + 1
			}
		}
		return 0
	}
	median := total/2
	if total & 0x1 == 0{
		median0 := total/2 - 1
		return float64(getKthNum(median0+1) + getKthNum(median+1)) / 2.0
	}else{
		return float64(getKthNum(median+1))
	}
}
/* 方法2： 划分数组
** 在统计中，中位数被用来： 将一个集合划分为两个长度相等的子集，其中一个子集中的元素总是大于另一个子集中的元素
** 首先，在任意位置 i 将 A 划分成两个部分：left_A 和 right_A
** 由于A中有 m 个元素， 所以有 m+1种 划分方法
** 对于B， 同理
** 将 left_A 和 left_B 放入一个集合，并将 right_A 和 right_B 放入另一个集合，构成新集合 left_part 和 right_part
** 当 A 和 B 的总长度是偶数的时候，如果可以确认：
	1. len(left_part) == len(right_part)
	2. max(left_part) <= min(right_part)
** 中位数就是前一部分的最大值和后一部分的最小值的平均值： median = ( max(left_part) + min(right_part) ) / 2
** 当 A 和 B 的总长度是奇数的时候，如果可以确认：
	1. len(left_part) == len(right_part) + 1
	2. max(left_part) <= min(right_part)
** 中位数就是前一部分的最大值： median = max(left_part)
** 要确保满足2个条件，只需要保证：
	1. i + j = m-i + n-j (m+n 为 偶数)  或 i + j = m-i + n-j + 1 （当 m+n 为奇数）
	   将 i， j 全部移动到左侧，推导出  i + j = (m+n+1) / 2
	2. 0 <= i <= m, 0 <= j <= n 若规定 len(A) <= len(B), 即 m <= n
	  这样对于任意 i 属于 [0, m] 都有 j = (m+n+1)/2 - i 属于 [0, n]，
	  那么在 [0, m] 的范围内枚举 i 并 得到 j
	  特例： A 比较 大， 交换 A 和 B
			m > n j 可能会出现负数
	3. B[j-1] <= A[i] 以及 A[i-1] <= B[j] 即前一部分的最大值小于等于后一部分的最小值
** 假设 B[j-1] A[i]  A[i-1]  B[j] 总存在， 对于 i = 0 = m  j = 0 = n 临界条件
** 规定 A[-1] = B[-1] = 负无穷  A[m] = B[n] = 正无穷
** 当一个数组不出现在前一部分时，对应的值为负无穷， 就不会对前一部分的最大值产生影响
** 当一个数组不出现在后一部分时，对应的值为正无穷，就不会对后一部分的最小值产生影响
** 在 [0, m] 中找到 i，使得
	B[j-1] <= A[i] 且 A[i-1] <= B[j]， j = (m+n+1)/2 - i
	等价于 A[i-1] <= B[j], j = (m+n+1)/2 - i
	证明：
	a. 当 i 从 0 - m 递增时， A[i-1] 递增， B[j] 递减， 所以一定存在一个最大的 i 满足 A[i-1] <= B[j]
	b. 如果 i 是最大的，那么说明 i+1 不满足。 将 i+1 带入可以得到 A[i] > B[j-1] 也即 B[j-1] < A[i]
** 因此， 在 i [0,m] 区间 二分查找，找到 最大满足 A[i-1] <= B[j] 的 i 值，就得到了划分方法
** 此时，划分前一部分元素中的最大值，以及划分后一部分元素中的最小值，才可能作为就是这两个数组的中位数
 */
func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n { // 这个交换方法值得学习
		return findMedianSortedArrays(nums2, nums1)
	}
	// nums1 数量少
	left, right := 0, m
	median1, median2 := 0, 0
	for left <= right{
		i := (left+right) / 2
		j := (m+n+1) / 2 - i
		nums_im1 := math.MinInt32
		if i != 0{
			nums_im1 = nums1[i-1]
		}
		nums_i := math.MaxInt32
		if i != m{
			nums_i = nums1[i]
		}

		nums_jm1 := math.MinInt32
		if j != 0 {
			nums_jm1 = nums2[j-1]
		}
		nums_j := math.MaxInt32
		if j != n {
			nums_j = nums2[j]
		}
		if nums_im1 <= nums_j { // A[i-1] <= B[j]
			median1 = max(nums_im1, nums_jm1)
			median2 = min(nums_i, nums_j)
			left = i + 1
		}else{ // A[i-1] > B[j] 减少 i 增大 j
			right = i - 1 // 由于 i j 彼此关联， 所以减少 i 必然会导致 j 增加
		}
	}
	if (m+n) % n == 0{
		return float64(median1 + median2) / 2.0
	}
	return float64(median1)
}
// 好理解的代码
func findMedianSortedArrays4(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	A, B := nums1, nums2
	if m > n {
		A, B = nums2, nums1
		m, n = n, m
	}
	iMin, iMax := 0, m
	for iMin <= iMax{
		i := (iMax + iMin) / 2
		j := (m+n+1)/2 - i
		if j != 0 && i != m && B[j-1] > A[i]{ // i 需要增大
			iMin = i + 1
		}else if i != 0 && j != n && A[i-1] > B[j]{ // i 需要减少
			iMax = i - 1
		}else{// 达到要求，并且将边界条件列出来单独考虑 即 B[j-1] <= A[i] && A[i-1] <= B[j]
			maxLeft := 0
			if i == 0 {
				maxLeft = B[j-1]
			}else if j == 0{
				maxLeft = A[i-1]
			}else{
				maxLeft = max(A[i-1], B[j-1])
			}
			if (m+n) & 0x1 == 1{ // 奇数
				return float64(maxLeft)
			}

			minRight := 0
			if i == m {
				minRight = B[j]
			}else if j == n {
				minRight = A[i]
			}else{
				minRight = min(B[j], A[i])
			}
			return float64(maxLeft + minRight) / 2.0
		}
	}
	strings.LastIndex()
	return 0.0
}

/* 2080. Range Frequency Queries
** Design a data structure to find the frequency of a given value in a given subarray.
** The frequency of a value in a subarray is the number of occurrences of that value in the subarray.
** Implement the RangeFreqQuery class:
	RangeFreqQuery(int[] arr) Constructs an instance of the class with the given 0-indexed integer array arr.
	int query(int left, int right, int value) Returns the frequency of value in the subarray arr[left...right].
** A subarray is a contiguous sequence of elements within an array.
** arr[left...right] denotes the subarray that contains the elements of nums between indices left and right (inclusive).
 */
type RangeFreqQuery struct {
	qf map[int][]int
	//pos [1e4 + 1]sort.IntSlice 可使用hash  直接调用 sort.IntSlice 方法
}
/*
func (q *RangeFreqQuery) Query(left, right, value int) int {
	p := q.pos[value] // value 在 arr 中的所有下标位置
	return p[p.Search(left):].Search(right + 1) // 在下标位置上二分，求 [left,right] 之间的下标个数，即为 value 的频率
}
*/

func Constructor(arr []int) RangeFreqQuery {
	ret := RangeFreqQuery{}
	ret.qf = map[int][]int{}
	for i, c := range arr{
		ret.qf[c] = append(ret.qf[c], i)
	}
	return ret
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	t := this.qf[value]
	/* 线性查找 超时
	   for _, c := range this.qf[value]{
	       if c >= left && c <= right{
	           ans++
	       }
	   }    */
	/* 二分查找 left right 的下标位置 */
	search := func(target int)int{
		i, j, mid := 0, len(t)-1, 0
		for i <= j{
			mid = int(uint(i+j)>>1)
			if t[mid] < target{
				i = mid + 1
			}else if t[mid] > target{
				j = mid - 1
			}else{
				return mid
			}
		}
		return i
	}
	l := search(left)
	// r := search(right)
	r := search(right+1) //关键的一步，在golang中  upper_bound 的实现
	//fmt.Println(t)
	//fmt.Println(l, r)
	return r - l
	/* 使用库函数
	   l := sort.SearchInts(t, left)
	   r := sort.SearchInts(t, right+1)//upper_bound  关键的一步，在golang中  upper_bound 的实现
	   return r - l
	*/
}






















