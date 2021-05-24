package array

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermute(t *testing.T) {
	for caseId, testCase := range []struct{
		nums []int
		want [][]int
	}{
		{[]int{1,2,3}, [][]int{{1,2,3},{1,3,2},{2,1,3},{2,3,1},{3,1,2},{3,2,1}}},
		{[]int{5,4,6,2}, [][]int{{5,4,6,2},{5,4,2,6},{5,6,4,2},{5,6,2,4},{5,2,4,6},{5,2,6,4},{4,5,6,2},{4,5,2,6},
											{4,6,5,2},{4,6,2,5},{4,2,5,6},{4,2,6,5},{6,5,4,2},{6,5,2,4},{6,4,5,2},{6,4,2,5},
											{6,2,5,4},{6,2,4,5},{2,5,4,6},{2,5,6,4},{2,4,5,6},{2,4,6,5},{2,6,5,4},{2,6,4,5}}},
	}{
		result := Permute(testCase.nums)
		assert.Equal(t, len(testCase.want), len(result), fmt.Sprintf("case-%d failed result=>%d, want=>%d", caseId, len(result), len(testCase.want)))
		assert.ElementsMatch(t, testCase.want, result, fmt.Sprintf("case-%d failed result=>%d, want=>%d", caseId, len(result), len(testCase.want)))
	}
}

func TestPermuteUniqueII(t *testing.T) {
	for caseId, testCase := range []struct{
		nums []int
		want [][]int
	}{
		{[]int{1,2,3}, [][]int{{1,2,3},{1,3,2},{2,1,3},{2,3,1},{3,1,2},{3,2,1}}},
		{[]int{1,1,2}, [][]int{{1,1,2},{1,2,1},{2,1,1}}},
	}{
		result := PermuteUniqueII(testCase.nums)
		assert.Equal(t, len(testCase.want), len(result), fmt.Sprintf("case-%d failed result=>%d, want=>%d", caseId, len(result), len(testCase.want)))
		assert.ElementsMatch(t, testCase.want, result, fmt.Sprintf("case-%d failed result=>%v, want=>%v", caseId, result, testCase.want))
	}
}

func TestLetterCasePermutation(t *testing.T){
	for caseId, testCase := range []struct{
		str string
		want []string
	}{
		{"0", []string{"0"}},
		{"a", []string{"a", "A"}},
		{"3z4", []string{"3z4","3Z4"}},
		{"a1b2", []string{"a1b2","a1B2","A1b2","A1B2"}},
		{"1234", []string{"1234"}},
	}{
		result := LetterCasePermutation(testCase.str)
		ok := assert.ElementsMatch(t, testCase.want, result, fmt.Sprintf("case-%d Failed: result = %s, but want = %s", caseId, result, testCase.want))
		if !ok {
			break
		}
	}
}
func TestLetterCasePermutationII(t *testing.T){
	for caseId, testCase := range []struct{
		str string
		want []string
	}{
		{"0", []string{"0"}},
		{"a", []string{"a", "A"}},
		{"3z4", []string{"3z4","3Z4"}},
		{"a1b2", []string{"a1b2","a1B2","A1b2","A1B2"}},
		{"1234", []string{"1234"}},
	}{
		result := LetterCasePermutationII(testCase.str)
		ok := assert.ElementsMatch(t, testCase.want, result, fmt.Sprintf("case-%d Failed: result = %s, but want = %s", caseId, result, testCase.want))
		if !ok {
			break
		}
	}
}

func TestLetterCasePermutationIII(t *testing.T){
	for caseId, testCase := range []struct{
		str string
		want []string
	}{
		{"0", []string{"0"}},
		{"a", []string{"a", "A"}},
		{"a1b2", []string{"a1b2","a1B2","A1b2","A1B2"}},
		{"3z4", []string{"3z4","3Z4"}},
		{"1234", []string{"1234"}},
	}{
		result := LetterCasePermutationIII(testCase.str)
		ok := assert.ElementsMatch(t, testCase.want, result, fmt.Sprintf("case-%d Failed: result = %s, but want = %s", caseId, result, testCase.want))
		if !ok {
			break
		}
	}
}
func TestLetterCasePermutationDFS(t *testing.T){
	for caseId, testCase := range []struct{
		str string
		want []string
	}{
		{"0", []string{"0"}},
		{"a", []string{"a", "A"}},
		{"a1b2", []string{"a1b2","a1B2","A1b2","A1B2"}},
		{"3z4", []string{"3z4","3Z4"}},
		{"1234", []string{"1234"}},
	}{
		result := LetterCasePermutationDFS(testCase.str)
		ok := assert.ElementsMatch(t, testCase.want, result, fmt.Sprintf("case-%d Failed: result = %s, but want = %s", caseId, result, testCase.want))
		if !ok {
			break
		}
	}
}

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

func TestCombine(t *testing.T){
	for caseId, testCase := range []struct{
		n		int
		k		int
		want 	[][]int
	}{
		{4, 2, [][]int{{2,4}, {3,4},{2,3},{1,2},{1,3},{1,4}}},
	}{
		result := CombineII(testCase.n, testCase.k)
		ok := assert.ElementsMatch(t, testCase.want, result, fmt.Sprintf("case-%d Failed: result = %v, but want = %v", caseId, result, testCase.want))
		if !ok{
			break
		}
	}
}

func TestTrulyMostPopular(t *testing.T) {
	for caseId, testCase := range []struct{
		names		[]string
		synonyms	[]string
		want		[]string
	}{
		{[]string{"John(15)","Jon(12)","Chris(13)","Kris(4)","Christopher(19)"},
		 []string{"(Jon,John)","(John,Johnny)","(Chris,Kris)","(Chris,Christopher)"},
		 []string{"John(27)","Chris(36)"},
		},
	}{
		result := TrulyMostPopular(testCase.names, testCase.synonyms)
		ok := assert.ElementsMatch(t, testCase.want, result, fmt.Sprintf("case-%d: result=%v, but want=%v", caseId, result, testCase.want))
		if !ok {
			break
		}
	}
}

func TestTrulyMostPopularII(t *testing.T) {
	for caseId, testCase := range []struct{
		names		[]string
		synonyms	[]string
		want		[]string
	}{
		{[]string{"John(15)","Jon(12)","Chris(13)","Kris(4)","Christopher(19)"},
			[]string{"(Jon,John)","(John,Johnny)","(Chris,Kris)","(Chris,Christopher)"},
			[]string{"John(27)","Chris(36)"},
		},
		{
			[]string{"John(15)","Jon(12)","Chris(13)","Kris(4)","Christopher(19)"},
			[]string{"(Jon,John)","(John,Johnny)","(Chris,Kris)","(Chris,Christopher)","(Jon,J)"},
			[]string{"J(27)","Chris(36)"},
		},
	}{
		result := TrulyMostPopularII(testCase.names, testCase.synonyms)
		ok := assert.ElementsMatch(t, testCase.want, result, fmt.Sprintf("case-%d: result=%v, but want=%v", caseId, result, testCase.want))
		if !ok {
			break
		}
	}
}

func TestNumIslandsDFS(t *testing.T) {
	for caseId, testCase := range []struct {
		grid	[][]byte
		want	int
	}{
		{[][]byte{[]byte{'1','1','1'},[]byte{'0','1','0'},[]byte{'1','1','1'}}, 1},
		{[][]byte{[]byte{'1','1','1','1','0'},[]byte{'1','1','0','1','0'},[]byte{'1','1','0','0','0'},[]byte{'0','0','0','0','0'}}, 1},
		{[][]byte{[]byte{'1','1','0','0','0'},[]byte{'1','1','0','0','0'},[]byte{'0','0','1','0','0'},[]byte{'0','0','0','1','1'}}, 3},
	}{
		result := NumIslandsDFS(testCase.grid)
		if result != testCase.want{
			t.Errorf("case-%d: result = %d but want = %d", caseId, result, testCase.want)
			return
		}
	}
}