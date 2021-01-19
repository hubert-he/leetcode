package main

import "fmt"

// 常规解法：按列计算
func trap(height []int) int {
	sum := 0
	for j := 0; j < len(height); j++ {
		max_left, max_right := 0, len(height) -1
		for i := 0; i <= j; i++{
			if height[i] > height[max_left] {
				max_left = i
			}
		}
		for i := j; i < len(height); i++ {
			if height[i] > height[max_right] {
				max_right = i
			}
		}
		//fmt.Println(height[max_left], height[max_right])
		if  height[max_left] > height[max_right]{
			sum += height[max_right] - height[j]
		} else {
			sum += height[max_left] - height[j]
		}
	}
	return sum
}
// 动态规划解法
func trap2(height []int) int {
	if len(height) == 0 {
		return 0
	}  // 要考虑[] 空slice
	lmax := make([]int, len(height))
	rmax := make([]int, len(height))
	sum := 0
	lmax[0] = height[0]
	rmax[len(height) - 1] = height[len(height) - 1]
	for i := 1; i < len(height); i++ {
		if height[i] > lmax[i-1]{
			lmax[i] = height[i]
		}else {
			lmax[i] = lmax[i-1]
		}
	}
	for i := len(height) - 2; i >=0; i-- {
		if height[i] > rmax[i+1] {
			rmax[i] = height[i]
		} else {
			rmax[i] = rmax[i+1]
		}
	}
	for i := 0; i < len(height); i++ {
		if lmax[i] > rmax[i] {
			sum += rmax[i] - height[i]
		} else {
			sum += lmax[i] - height[i]
		}
	}
	return sum
}
// 双指针解法
func trap3(height []int) int {
	left := 0
	right := len(height) - 1
	var sum, left_max, right_max int
	for left < right {
		// 每次循环只有一个方向移动，为何只有一个方向，是为确定left_max right_max 谁最小的 情况
		// 如果发现左边比右边大，就会一直计算从右边计算，因为此时right_max 肯定小于左边
		if height[left] < height[right] {
			if height[left] >= left_max {
				left_max = height[left]
			} else {
				fmt.Printf("%d : %d \n", left, left_max - height[left])
				sum += left_max - height[left]
			}
			left++
		} else {
			if height[right] >= right_max {
				right_max = height[right]
			} else {
				sum += right_max - height[right]
			}
			right--
		}
	}
	return sum
}

// 单调栈解法
func trap4(height []int) int {
	 stack := []int{}
	 sum := 0
	 for i := 0; i < len(height); i++ {
	 	fmt.Println(stack)
		for len(stack) > 0 && height[i] > height[stack[len(stack) - 1]]{
			mid := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			if len(stack) > 0 {
				top := stack[len(stack) - 1]
				//stack = stack[:len(stack) - 1]
				if height[top] > height[i]{
					sum += (height[i] - height[mid]) * (i - top - 1)
				} else {
					sum += (height[top] - height[mid]) * (i - top - 1)
				}
			}
			fmt.Println(sum)
		 }
		 stack = append(stack, i)
	 }
	return sum
}

func main(){
	testcase := []int{0,4,0,2,1,0,1,5,2,1,2,1}
	fmt.Println(trap(testcase))
	t2 := []int{4,2,0,3,2,5}
	fmt.Println(trap2(testcase))
	fmt.Println(trap(t2))

	fmt.Println(trap3(testcase))
	fmt.Println(trap3(t2))

	fmt.Println(trap4(testcase))
	fmt.Println(trap4(t2))
}