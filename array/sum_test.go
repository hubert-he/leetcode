package array

import "testing"

func TestMaxmiumScore(t *testing.T) {
	for caseId, testCase := range []struct{
		cards		[]int
		cnt			int
		want		int
	}{
		{[]int{10, 9}, 1, 10},
		{[]int{1,2,8,9}, 3, 18},
		{[]int{3,3,1}, 1, 0},
		{[]int{9,5,9,1,6,10,3,4,5,1}, 2, 18},
		{[]int{7,1,5,8,3,3,1,2}, 7, 28},
	}{
		result := MaxmiumScore(testCase.cards, testCase.cnt)
		if result != testCase.want{
			t.Errorf("MaxmiumScore-case-%d: result=%d, but want=%d", caseId, result, testCase.want)
			break
		}
		result = MaxmiumScorePrefixSum(testCase.cards, testCase.cnt)
		if result != testCase.want{
			t.Errorf("MaxmiumScorePrefixSum-case-%d: result=%d, but want=%d", caseId, result, testCase.want)
			break
		}
	}
}

func TestCountKDifference(t *testing.T) {
	for caseId, testCase := range []struct{
		nums		[]int
		k			int
		want		int
	}{
		{[]int{1,2,2,1}, 1, 4},
		{[]int{1,3}, 3, 0},
		{[]int{3,2,1,5,4}, 2, 3},
		{[]int{7,7,8,3,1,2,7,2,9,5}, 6, 6},
	}{
		result := CountKDifference(testCase.nums, testCase.k)
		if result != testCase.want{
			t.Errorf("case-%d: result=%d, but want=%d", caseId, result, testCase.want)
			break
		}
		result = CountKDifferenceDetail(testCase.nums, testCase.k)
		if result != testCase.want{
			t.Errorf("CountKDifferenceDetail-case-%d: result=%d, but want=%d", caseId, result, testCase.want)
			break
		}
	}
}

func TestMinStartValue(t *testing.T){
	for caseId, testCase := range []struct{
		nums		[]int
		want		int
	}{
		{[]int{-3,2,-3,4,2}, 5},
		{[]int{1,3}, 1},
		{[]int{1,-2,-3}, 5},
	}{
		result := MinStartValue(testCase.nums)
		if result != testCase.want {
			t.Errorf("case-%d: result=%d but want=%d", caseId, result, testCase.want)
			break
		}
	}
}