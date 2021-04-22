package stack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseInPlace(t *testing.T){
	for caseId, testCase := range []struct {
		stack	[]int
		want 	[]int
	}{
		{[]int{},[]int{}},
		{[]int{1}, []int{1}},
		{[]int{1,2}, []int{2,1}},
		{[]int{1,2,3,4,5}, []int{5,4,3,2,1}},
	}{
		result := make([]int, len(testCase.want))
		copy(result, testCase.stack)
		ReverseInPlace(&result)
		ok := assert.Equal(t, testCase.want, result, fmt.Sprintf("case-%d Failed: result=>%v but want=>%v", caseId, result, testCase.want))
		if !ok {
			break
		}
	}
}
