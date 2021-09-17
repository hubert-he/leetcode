package main

import (
	x "./heap"
	"fmt"
)

func main(){
	fmt.Println(x.MaxSlidingWindowHeap([]int{9,10,9,-7,-4,-8,2,-6}, 5))
}