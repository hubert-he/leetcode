package main

import (
	"./dp"
	xh "./heap"
	"fmt"
)

func main(){
	fmt.Println(DP.BoundKnapsackQueue(2, 5, []int{1,2}, []int{1,2}, []int{2,1}))
	fmt.Println("-----")
	fmt.Println(DP.BoundKnapsackBinary(2, 5, []int{1,2}, []int{1,2}, []int{2,1}))
	fmt.Println("-----")
	fmt.Println(DP.MultiPackage1( 2, 5, []int{1,2}, []int{1,2}, []int{2,1}))
	fmt.Println("-----")
	fmt.Println(DP.MultiPackage( 2, 5, []int{1,2}, []int{1,2}, []int{2,1}))
	fmt.Println("-----")
	fmt.Println(xh.MaxSlidingWindowHeap([]int{9,10,9,-7,-4,-8,2,-6}, 5))
}