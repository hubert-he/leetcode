package heap

import (
	"fmt"
	"runtime"
	"testing"
)

func init(){
	_, codePath, _, _ := runtime.Caller(0)
	fmt.Println(codePath[:len(codePath)-3])
}

func TestSmallestK(t *testing.T) {
}

func TestMaxSlidingWindowHeap(t *testing.T){
	for caseId, testCase := range []struct{
		nums		[]int
		k 			int
		want		[]int
	}{
		{[]int{9,10,9,-7,-4,-8,2,-6}, 5, []int{10,10,9,2}},
		{[]int{1,3,-1,-3,5,3,6,7}, 3, []int{3,3,5,5,6,7}},
	}{
		result := MaxSlidingWindowHeap(testCase.nums, testCase.k)
		for i := range result{
			if result[i] != testCase.want[i]{
				t.Errorf("case-%d failed: result=%v want=%v", caseId, result, testCase.want)
				return
			}
		}
	}
}
