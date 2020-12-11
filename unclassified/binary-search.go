package unclassified

import "fmt"

/* Binary Search 两大基本原则
	1. 每次迭代都要缩减搜索区域  	Shrink the search scope in every iteration/recursion
	2. 每次缩减都不能排除掉潜在答案 	Can NOT exclude potential answers during each shrinking
*/
func SearchArray(a []int, key int) int {
	return binarySearchII(a, key)
}
// 模板一：查找准确值
// 循环条件： 	l <= r
// 缩减搜索空间：	l = mid + 1  &&  r = mid - 1
func binarySearch(a []int, key int) int {
	mid := -1
	i,j := 0, len(a)-1
	for i <= j { // 等号容易忘记，匹配i=j a[i] == key这种情况
		//mid = (j - i)/2 +i   // +i 容易忘记，注意除2只是偏移量
		mid = (i + j) >> 1  // 优化
		if key < a[mid] {
			j = mid - 1
		} else if key == a[mid]{
			break;
		} else {
			i = mid + 1
		}
	}
	if (i <= j) { // 等号容易忘记
		return mid
	} else {
		return -1
	}
}

// 模板二： 查找模糊值
// 循环条件：		l < r
// 缩减搜索空间：	l = mid, r = mid - 1(最后出现的) 或者 l = mid + 1, r = mid（最先出现的）


// 万用型模板：
// 循环条件：		l < r - 1
// 缩减搜索空间： 	l = mid, r = mid
func binarySearchII(a []int, key int) int {
	mid := -1
	l,r := 0, len(a)-1
	for l < r - 1 { // 剩余2个 l r 指向
		//mid = (r - l)/2 + l   // +i 容易忘记，注意除2只是偏移量
		mid = (r + l) >> 1  // 优化
		fmt.Printf("1=>%d 2=>%d\n", mid, (r+l)>>1)
		if key == a[mid]{
			return mid
		} else if key < a[mid] {
			r = mid
		} else {
			l = mid
		}
	}
	if a[l] == key {
		return l
	}
	if a[r] == key {
		return r
	}
	return -1
}