package BinarySearch

import (
	"fmt"
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

/* 1385. Find the Distance Value Between Two Arrays
** Given two integer arrays arr1 and arr2, and the integer d, return the distance value between the two arrays.
** The distance value is defined as the number of elements arr1[i] such that
** there is not any element arr2[j] where |arr1[i]-arr2[j]| <= d.
 */
//二分查找的本质是对可行区间的压缩。所以我们正好可以使用二分查找来确定arr2中是否有元素落在上述区间内
// 二分写法一：右边界起始为 数组最右侧下标
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	n := len(arr2)
	sort.Ints(arr2)
	fmt.Println(arr2)
	ans := 0
	for i := range arr1{
		// 对arr1中的每一个元素进行检索，判断arr2中是否有元素落在[arr1[i] - d, arr1[i] + d]区间内
		// 转化为求 arr1[i] - d <= arr2[j] <= arr1[i] + d
		j, k := 0, n-1
		for j <= k{
			mid := int(uint(j+k)>>1)
			if arr2[mid] >= arr1[i] - d && arr2[mid] <= arr1[i] + d{
				break
			}
			if arr2[mid] > arr1[i] + d {
				k = mid - 1
			}
			if arr2[mid] < arr1[i] - d{
				j = mid + 1
			}
		}
		if j > k{
			fmt.Println(arr1[i])
			ans++
		}
	}
	return ans
}
// 第二种写法： 右边界起始为 数组长度
func findTheDistanceValue2(arr1 []int, arr2 []int, d int) int {
	n := len(arr2)
	sort.Ints(arr2)
	//fmt.Println(arr2)
	ans := 0
	for i := range arr1{
		// 对arr1中的每一个元素进行检索，判断arr2中是否有元素落在[arr1[i] - d, arr1[i] + d]区间内
		// 转化为求 arr1[i] - d <= arr2[j] <= arr1[i] + d
		j, k := 0, n
		for j < k{
			mid := int(uint(j+k)>>1)
			if arr2[mid] >= arr1[i] - d && arr2[mid] <= arr1[i] + d{
				break
			}
			if arr2[mid] > arr1[i] + d {
				k = mid // 这里要能访问到
			}
			if arr2[mid] < arr1[i] - d{
				j = mid + 1
			}
		}
		if j >= k{ // 无法查找的界定
			ans++
		}
	}
	return ans
}

// 写法三
func findTheDistanceValue3(arr1 []int, arr2 []int, d int) int {
	n := len(arr2)
	sort.Ints(arr2)
	//fmt.Println(arr2)
	ans := 0
	for i := range arr1{
		// 对arr1中的每一个元素进行检索，判断arr2中是否有元素落在[arr1[i] - d, arr1[i] + d]区间内
		// 转化为求 arr1[i] - d <= arr2[j] <= arr1[i] + d
		j, k := 0, n
		for j <= k && j < n{ // 要加 = 包含 j == k 的边界情况，同时又不能允许越界
			mid := int(uint(j+k)>>1)
			if arr2[mid] >= arr1[i] - d && arr2[mid] <= arr1[i] + d{
				break
			}
			if arr2[mid] > arr1[i] + d {
				k = mid - 1
			}
			if arr2[mid] < arr1[i] - d{
				j = mid + 1
			}
		}
		if j > k || j == n{ // 越界情况也属于找不到情况
			ans++
		}
	}
	return ans
}

/* 34. Find First and Last Position of Element in Sorted Array
** Given an array of integers nums sorted in non-decreasing order,
** find the starting and ending position of a given target value.
** If target is not found in the array, return [-1, -1].
** You must write an algorithm with O(log n) runtime complexity.
 */
/* 此题可以借助 查询查找位置来定位左右边界：
** 1. search(target)   返回 左边界
** 2. search(target+1) 返回 右边界
*/
func searchRange(nums []int, target int) []int {
	n := len(nums)
	search := func(t int)int{ // 实际上是找左边界
		i, j := 0, n
		for i < j {
			mid := int(uint(i+j)>>1)
			if nums[mid] > target{
				j = mid - 1
			}else if nums[mid] < target {
				i = mid + 1
			}else{
				j = mid
			}
		}
		return i
	}
	return []int{search(target), search(target+1)-1}
}
/* 此题引出了 类似C++ lower  upper 2分查询上下边界的问题
** 1. 查找左边界的情况中，当中间元素命中要查找的元素时，将右边界固定为mid，这样循环便能在 [left, mid)范围内查找。
 		如果目标数在数组中只有一个，这样把这个元素跳过了吗？
 		答案是不会，在之后的迭代中，left 位置会逐步更新，最终和right位置一直保持在目标元素上。
** 2. 查找右边界的情况中，当中间元素命中要查找的数时，让左边界left等于中间元素的下标加1，事实上跳过了这个元素。
		因此，循环可以在 [mid + 1, right)的区间中迭代查找。最终，右指针会和和左指针汇合在最右一个目标元素的下标+1的位置。
		因而最终返回的是左指针（右指针）减一的的结果
 */
func searchRange_lower_upper(nums []int, target int) []int {
	n := len(nums)
	if n == 0{ return []int{-1, -1}}
	ans := []int{}
	i, j := 0, n
	// 查询左边界
	for i < j {
		mid := int(uint(i+j)>>1)
		if nums[mid] < target{
			i = mid + 1 // [left, mid)
		}else{
			j = mid
		}
	}
	if i >= n || nums[i] != target{
		ans = append(ans, -1)
	}else{
		ans = append(ans, i)
	}
	i, j = 0, n
	// 查询右边界
	for i < j {
		mid := int(uint(i+j)>>1)
		if nums[mid] <= target{
			i = mid + 1
		}else{
			j = mid
		}
	}
	if j > 0 && nums[j-1] == target{
		ans = append(ans, j)
	}else{
		ans = append(ans, -1)
	}
	return ans
}

/* 852. Peak Index in a Mountain Array
** Let's call an array arr a mountain if the following properties hold:
	arr.length >= 3
** There exists some i with 0 < i < arr.length - 1 such that:
	arr[0] < arr[1] < ... arr[i-1] < arr[i]
	arr[i] > arr[i+1] > ... > arr[arr.length - 1]
** Given an integer array arr that is guaranteed to be a mountain,
** return any i such that arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1].
 */
/* 由于 arr 数值各不相同, 因此峰顶元素左侧必然满足严格单调递增, 峰顶元素右侧必然不满足
** 因此 以峰顶元素为分割点的 arr 数组, 根据与 前一元素/后一元素 的大小关系，具有二段性
** 峰顶元素左侧满足 arr[i-1] < arr[i]
** 峰顶元素右侧满足 arr[i] < arr[i+1]
 */
func peakIndexInMountainArray(arr []int) int {
	i, j := 0, len(arr)-1 // len(arr)-1 无需访问， 因为肯定不在两端
	for i < j {
		mid := int(uint(i+j)>>1)
		if arr[mid] < arr[mid+1]{ // 左侧
			i = mid + 1
		}else if arr[mid] > arr[mid+1]{
			j = mid
		}
	}
	return i
}

/* 441. Arranging Coins
** You have n coins and you want to build a staircase with these coins.
** The staircase consists of k rows where the ith row has exactly i coins.
** The last row of the staircase may be incomplete.
** Given the integer n, return the number of complete rows of the staircase you will build.
 */
// 2022-04-22 调不出结果，mid 原因
func arrangeCoins(n int) int {
	sum := 0
	i, j := 0, n
	for i < j{
		//mid := int(uint(i+j)>>1)
		mid := int(uint(i+j)>>1) + 1 // 注意mid 情况
		//sum = (i+mid)*(mid-i+1)/2
		sum = mid*(mid+1)/2
		//fmt.Println(mid, sum)
		if sum > n{
			j = mid-1
		}else if sum <= n{
			i = mid
		}
	}
	return i
}
// 官方解答
func arrangeCoins2(n int) int {
	left, right := 1, n
	for left < right{
		mid := int(uint(left+right+1)>>1)
		if mid * (mid+1) <= 2*n{
			left = mid
		}else{
			right = mid-1
		}
	}
	return left
}

/* 33. Search in Rotated Sorted Array
** There is an integer array nums sorted in ascending order (with distinct values).
** Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length)
** such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
** For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].
** Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums,
** or -1 if it is not in nums.
** You must write an algorithm with O(log n) runtime complexity.
 */
// 2022-04-28 刷出此题
// 1. 分情况讨论
// 2. 由于 j 指向的是 数组长度， 因此缩 j 的时候  注意for 退出条件 i < j ， j = mid  否则会漏掉 i == j 情况下的值
func search(nums []int, target int) int {
	n := len(nums)
	i, j := 0, n
	if target >= nums[0]{
		for i < j {
			mid := int(uint(i+j)>>1)
			if nums[mid] > target{
				j = mid
			}else if nums[mid] < target {
				if nums[mid] < nums[0]{
					j = mid
				}else{
					i = mid+1
				}
			}else{
				return mid
			}
		}
	}else{
		for i < j{
			mid := int(uint(i+j)>>1)
			if nums[mid] > target{
				if nums[mid] >= nums[0]{
					i = mid+1
				}else{
					//j = mid-1
					j = mid
				}
			}else if nums[mid] < target {
				i = mid+1
			}else{
				return mid
			}
		}
	}
	return -1
}

/* 81. Search in Rotated Sorted Array II
** There is an integer array nums sorted in non-decreasing order (not necessarily with distinct values).
** Before being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length)
** such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
** For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,2,4,4].
** Given the array nums after the rotation and an integer target, return true if target is in nums, or false if it is not in nums.
** You must decrease the overall operation steps as much as possible.
 */
// 2022-04-28 未能刷出此题
// 参考case 输入
//  [1,1,1,2,1]
//  2
// 思维卡点在：
//	对于数组中有重复元素的情况，二分查找时可能会有 a[l]=a[mid]=a[r]，此时无法判断区间 [l,mid] 和区间 [mid+1,r] 哪个是有序的。
// [0,2] 和 区间 [3,4] 哪个是有序的
// 这种功情况的处理方式：
//		只能将当前二分区间的左边界加一, 右边界减一，然后在新区间上继续二分查找
func search_Error(nums []int, target int) bool {
	n := len(nums)
	i, j := 0, n
	if target >= nums[0]{
		for i < j {
			mid := int(uint(i+j)>>1)
			if nums[mid] > target{
				j = mid
			}else if nums[mid] < target {
				if nums[mid] < nums[0]{
					j = mid
				}else{
					i = mid+1
				}
			}else{
				return true
			}
		}
	}else{
		for i < j{
			mid := int(uint(i+j)>>1)
			if nums[mid] > target{
				//if nums[mid] >= nums[0]{
				if nums[mid] > nums[0]{
					i = mid+1
				}else{
					//j = mid-1
					j = mid
				}
			}else if nums[mid] < target {
				i = mid+1
			}else{
				return true
			}
		}
	}
	return false
}

func search2(nums []int, target int) bool {
	n := len(nums)
	if n == 0 { return false }
	if n == 1{ return nums[0] == target }
	l, r := 0, n-1
	for l <= r{
		mid := int(uint(l+r)>>1)
		if nums[mid] == target{
			return true
		}
		// 特殊的情况: 因为mid 不等于target 所以在 3 个值相等情况下，直接两边全部缩小
		if nums[l] == nums[mid] && nums[mid] == nums[r]{
			l++
			r--
		}else if nums[l] <= nums[mid]{
			if nums[l] <= target && target < nums[mid]{
				r = mid - 1
			}else{
				l = mid + 1
			}
		}else{
			if nums[mid] < target && target <= nums[n-1]{
				l = mid+1
			}else{
				r = mid -1
			}
		}
	}
	return false
}

/* 153. Find Minimum in Rotated Sorted Array[寻找旋转排序数组中的最小值]
** Suppose an array of length n sorted in ascending order is rotated between 1 and n times.
** For example, the array nums = [0,1,2,4,5,6,7] might become:
	[4,5,6,7,0,1,2] if it was rotated 4 times.
	[0,1,2,4,5,6,7] if it was rotated 7 times.
** Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].
** Given the sorted rotated array nums of unique elements, return the minimum element of this array.
** You must write an algorithm that runs in O(log n) time.
*/
// 2022-04-28 刷出此题
func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n
	for left < right{
		mid := int(uint(left+right)>>1)
		// 遗漏点-1 mid < n-1 设定 否则数组越界
		if mid < n-1 && nums[mid] > nums[mid+1]{
			return nums[mid+1]
		}
		if nums[mid] > nums[0]{
			left = mid + 1
		}else if nums[mid] < nums[n-1]{
			right = mid
		}else{ // 易漏点-2 搜不下去了 搜索区间长度为 1 即  nums[mid] == nums[0] 或者 nums[mid] == nums[n-1]
			return nums[mid]
		}
	}
	return nums[0]
}
// 官方解答 直接使用二分 比较 left right 两侧值，保证 断点在其中间
func FindMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high{
		mid := (low^high)>>1 + (low&high)
		if nums[mid] < nums[high]{
			high = mid
		}else{
			low = mid+1
		}
	}
	return nums[low]
}

/* 154. Find Minimum in Rotated Sorted Array II[寻找旋转排序数组中的最小值 II]
** Suppose an array of length n sorted in ascending order is rotated between 1 and n times.
** For example, the array nums = [0,1,4,4,5,6,7] might become:
	[4,5,6,7,0,1,4] if it was rotated 4 times.
	[0,1,4,4,5,6,7] if it was rotated 7 times.
Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].
Given the sorted rotated array nums that may contain duplicates, return the minimum element of this array.
You must decrease the overall operation steps as much as possible.
*/
/* 区别在于 有重复元素
** 考虑数组中的最后一个元素 x ：在最小值右侧的元素，它们的值一定都小于等于 x ， 而在最小值左侧的元素，他们的值一定都大于等于x。
** 依据上述性质，可以通过二分查找找出最小值。
** 一共分3种情况：
** 1. nums[mid] > nums[high]
** 2. nums[mid] < nums[high]
** 3. nums[mid] == nums[high]
** 由于重复元素的存在，我们并不能确定nums[mid]究竟在最小值的左侧还是右侧，因此我们不能莽撞地忽略某一部分的元素
** 唯一可以知道的是，由于它们的值相同，所以无论nums[high]是不是最小值，都有一个它的「替代品」nums[mid],因此依然可以继续缩
** 即忽略二分的右端点。
 */
func FindMinII(nums []int) int {
	n := len(nums)
	low, high := 0, n-1
	//是有序单调数组
	if nums[low] < nums[high]{
		return nums[low]
	}
	for low < high{
		mid := (low^high)>>1 + low&high
		if nums[mid] > nums[high]{
			low = mid+1
		}else if nums[mid] < nums[high]{
			high = mid
		}else{
			high--
		}
	}
	return nums[low]
}
func FindMinIILeft(nums []int) int {
	n := len(nums)
	low, high := 0, n-1
	//是有序单调数组
	if nums[low] < nums[high]{
		return nums[low]
	}
	for low < high{
		//如果二分后的数组是有序数组，则返回最左元素，即为最小
		if nums[low] < nums[high]{
			return nums[low]
		}
		mid := (low^high)>>1 + low&high
		//若最左小于mid元素，则最左到mid是严格递增的，那么最小元素必定在mid之后
		if nums[mid] > nums[low]{
			low = mid+1
		}else if nums[mid] < nums[high]{
			high = mid
		}else{
			high--
		}
	}
	return nums[low]
}




