package Tree

import "testing"

func TestPathSumIII(t *testing.T) {
	for caseId, testCase := range []struct{
		nums []interface{}
		sum  int
		want int
	}{
		{[]interface{}{1,-2,-3,1,3,-2,nil,-1}, 3, 1},
		{[]interface{}{10,5,-3,3,2,nil,11,3,-2,nil,1}, 8, 3},
		{[]interface{}{5,4,8,11,nil,13,4,7,2,nil,nil,5,1}, 22, 3},
		{[]interface{}{0,1,1}, 1, 4},
		{[]interface{}{1}, 0, 0},
		{[]interface{}{1}, 1, 1},
	}{
		tree := GenerateBiTree(testCase.nums)
		result := PathSumIII(tree, testCase.sum)
		if result != testCase.want{
			t.Errorf("case-%d: result = %d, but want %d", caseId, result, testCase.want)
		}
	}
}

func TestPathSumIV(t *testing.T) {
	for caseId, testCase := range []struct{
		nums []int
		want int
	}{
		{[]int{113,221}, 4},
		{[]int{113, 215, 221}, 12},
	}{
		result := PathSumIV(testCase.nums)
		if result != testCase.want {
			t.Errorf("case-%d Failed: result=%d, but want %d", caseId, result, testCase.want)
			break
		}
	}
}

func TestPathSumIVBFS(t *testing.T) {
	for caseId, testCase := range []struct{
		nums []int
		want int
	}{
		{[]int{113,221}, 4},
		{[]int{113, 215, 221}, 12},
	}{
		result := PathSumIVBFS(testCase.nums)
		if result != testCase.want {
			t.Errorf("case-%d Failed: result=%d, but want %d", caseId, result, testCase.want)
			break
		}
	}
}
