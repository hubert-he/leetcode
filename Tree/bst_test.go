package Tree

import "testing"

func TestBinarySearchTree_GetMinimumDifference(t *testing.T) {
	for caseId, testCase := range []struct{
		num []int
		want int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1,3}, 2},
		{[]int{1,nil,3,2}, 1},
	}{

	}
}
