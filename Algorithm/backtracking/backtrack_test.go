package backtracking

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombinationSum(t *testing.T){
	for caseId, testCase := range []struct{
		nums []int
		target int
		want [][]int
	}{
		{[]int{2,3,6,7}, 7,[][]int{{2,2,3}, {7}}},
		{[]int{2,3,5}, 8, [][]int{{2,2,2,2},{2,3,3},{3,5}}},
		{[]int{2}, 1, [][]int{}},
		{[]int{1}, 2, [][]int{{1,1}}},
	}{
		result := CombinationSumII(testCase.nums, testCase.target)
		ok := assert.ElementsMatch(t, testCase.want, result, fmt.Sprintf("case-%d Failed: result = %v, but want = %v", caseId, result, testCase.want))
		if !ok{
			break
		}
	}
}

func TestCombinationSum2(t *testing.T){
	for caseId, testCase := range []struct{
		nums []int
		target int
		want [][]int
	}{
		{[]int{10,1,2,7,6,5}, 8, [][]int{{1,2,5}, {1,7},{2,6}}},
		{[]int{10,1,2,7,6,1,5}, 8, [][]int{{1,1,6}, {1,2,5}, {1,7}, {2,6}}},
		{[]int{2,5,2,1,2}, 5, [][]int{{1,2,2}, {5}}},
	}{
		result := CombinationSum2(testCase.nums, testCase.target)
		ok := assert.ElementsMatch(t, testCase.want, result, fmt.Sprintf("case-%d Failed: result = %v, but want = %v", caseId, result, testCase.want))
		if !ok{
			break
		}
	}
}
