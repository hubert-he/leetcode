package array

/* 977. Squares of a Sorted Array
** Given an integer array nums sorted in non-decreasing order,
** return an array of the squares of each number sorted in non-decreasing order.
*/
func SortedSquares(nums []int) []int {
	n := len(nums)
	ans := []int{}
	i := 0
	for i < n{
		if nums[i] < 0{
			i++
		}else{
			break
		}
	}
	a, b := nums[:i], nums[i:]
	slow, fast := len(a)-1, 0
	for slow >= 0 || fast < len(b){
		if  slow >= 0 && fast < len(b){
			av, bv := a[slow]*a[slow], b[fast]*b[fast]
			if av < bv{
				ans = append(ans, av)
				slow--
			}else{
				ans = append(ans, bv)
				fast++
			}
		}else if slow >= 0{
			ans = append(ans, a[slow]*a[slow])
			slow--
		}else{
			ans = append(ans, b[fast]*b[fast])
			fast++
		}
	}
	return ans
}
