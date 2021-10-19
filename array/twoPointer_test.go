package array

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	for caseId, testCase := range []struct{
		nums []int
		want []int
	}{
		{[]int{-1,-1,0,1,3,6}, []int{0,1,1,1,9,36}},
		{[]int{0,1,3,6}, []int{0,1,9,36}},
		{[]int{1,3,6}, []int{1,9,36}},
		{[]int{-4, -2, -2, -1}, []int{1,4,4,16}},
	}{
		result := SortedSquares(testCase.nums)
		if !assert.Equal(t, testCase.want, result, fmt.Sprintf("case-%d: result=%v, want=%v", caseId, result, testCase.want)){
			break
		}
	}
}
