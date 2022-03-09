package Stack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextGreaterElements(t *testing.T){
	for caseId, testCase := range []struct{
		nums	[]int
		want 	[]int
	}{
		{[]int{100,1,11,1,120,111,123,1,-1,-100}, []int{120,11,120,120,123,123,-1,100,100,100}},
		{[]int{1,8,-1,-100,-1,222,1111111,-111111}, []int{8,222,222,-1,222,1111111,-1,1}},
		{[]int{1,1,1,1,1}, []int{-1,-1,-1,-1,-1}},
		{[]int{1,5,3,6,8}, []int{5,6,6,8,-1}},
		{[]int{5,4,3,2,1}, []int{-1,5,5,5,5}},
		{[]int{1,2,3,4,3}, []int{2,3,4,-1,4}},
		{[]int{1,2,1}, []int{2,-1,2}},
	}{
		result := NextGreaterElements(testCase.nums)
		ok := assert.ElementsMatch(t, testCase.want, result,
			fmt.Sprintf("case-%d Failed: result = %v, but want = %v", caseId, result, testCase.want))
		if !ok {
			break
		}
	}
}
