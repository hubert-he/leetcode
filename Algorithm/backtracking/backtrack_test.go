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

func TestAddOperators(t *testing.T){
	for caseId, testCase := range []struct{
		num		string
		target	int
		want	[]string
	}{
		{"10009", 9, []string{"1*0*0*0+9","1*0*0+0+9","1*0*0-0+9","1*0+0*0+9","1*0+0+0+9","1*0+0-0+9","1*0-0*0+9","1*0-0+0+9","1*0-0-0+9","10*0*0+9","10*0+0+9","10*0-0+9","100*0+9"}},
		{"105", 5, []string{"1*0+5","10-5"}},
		{"109", 9, []string{"1*0+9"}},
		{"1009", 9, []string{"1*0*0+9","1*0+0+9","1*0-0+9","10*0+9"}},
		{"1005", 5, []string{"1*0*0+5","1*0+0+5","1*0-0+5","10*0+5","10+0-5","10-0-5"}},
		{"00", 0, []string{"0*0", "0+0", "0-0"}},
		{"000", 0, []string{"0*0*0","0*0+0","0*0-0","0+0*0","0+0+0","0+0-0","0-0*0","0-0+0","0-0-0"}},
		{"232", 8, []string{"2*3+2","2+3*2"}},
		{"123", 6, []string{"1*2*3", "1+2+3"}},
		{"3456237490", 9191, []string{}},
	}{
		result := AddOperators(testCase.num, testCase.target)
		ok := assert.ElementsMatch(t, testCase.want, result, fmt.Sprintf("case-%d Failed: result=%v, want=%v", caseId, result, testCase.want))
		if !ok { break }
	}
}
