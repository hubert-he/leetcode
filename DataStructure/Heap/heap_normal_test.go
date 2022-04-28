package Heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeapSort(t *testing.T){
	for caseId, testCase := range []struct{
		arr 	[]int
		want	[]int
	}{
		{[]int{10, 5, 8, 3, 4, 6, 7, 1, 2}, []int{1, 2, 3, 4, 5, 6, 7, 8, 10}},
	}{
		result := heap_sort(testCase.arr)
		if assert.Equal(t, testCase.want, result,
			"case-%d: result=%v, but want=%v", caseId, result, testCase.want){
			break
		}
	}
}
