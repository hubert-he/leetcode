package main

import (
	"./unclassified"
	"fmt"
)

func main() {
	var testdata = []int{8,1,2,2,3,3}
	ret := unclassified.SmallerNumbersThanCurrentSolution(testdata)
	fmt.Println(ret)
}
