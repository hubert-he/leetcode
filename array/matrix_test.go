package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	for caseId, testCase := range []struct{
		arr		[][]int
		want	[]int
	}{
		{[][]int{[]int{1,2,3}, []int{4,5,6}, []int{7,8,9}}, []int{1,2,3,6,9,8,7,4,5}},
		{[][]int{[]int{1,2,3,4}, []int{5,6,7,8}, []int{9,10,11,12}}, []int{1,2,3,4,8,12,11,10,9,5,6,7}},
	}{
		result := SpiralOrder(testCase.arr)
		assert.Equal(t, testCase.want, result, "case-%d result=%v want=%v", caseId, result, testCase.want)
	}
}

func TestGenerateMatrix(t *testing.T) {
	for caseId, testCase := range []struct{
		n		int
		want	[][]int
	}{
		{3, [][]int{[]int{1,2,3}, []int{8,9,4}, []int{7,6,5}}},
		{1, [][]int{[]int{1}}},
	}{
		result := GenerateMatrix(testCase.n)
		assert.Equal(t, testCase.want, result, "case-%d result=%v want=%v", caseId, result, testCase.want)
	}
}

func TestSpiralMatrixIII(t *testing.T) {
	for caseId, testCase := range []struct{
		rows,	cols 	int
		rStart,	cStart 	int
		want			[][]int
	}{
		{1,4,0,0, [][]int{[]int{0,0},[]int{0,1},[]int{0,2},[]int{0,3}}},
		{5,6,1,4,
			[][]int{
				[]int{1,4},[]int{1,5},[]int{2,5},[]int{2,4},[]int{2,3},[]int{1,3},
				[]int{0,3},[]int{0,4},[]int{0,5},[]int{3,5},[]int{3,4},[]int{3,3},[]int{3,2},[]int{2,2},
				[]int{1,2},[]int{0,2},[]int{4,5},[]int{4,4},[]int{4,3},[]int{4,2},[]int{4,1},[]int{3,1},
				[]int{2,1},[]int{1,1},[]int{0,1},[]int{4,0},[]int{3,0},[]int{2,0},[]int{1,0},[]int{0,0},
			},
		},
	}{
		result := SpiralMatrixIII(testCase.rows, testCase.cols, testCase.rStart, testCase.cStart)
		assert.Equal(t, testCase.want, result, "case-%d result=%v want=%v", caseId, result, testCase.want)
	}
}