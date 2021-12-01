package linear

import "testing"

func TestStoneGame(t *testing.T){
	for caseId, testCase := range []struct{
		piles		[]int
		want		bool
	}{
		{[]int{5,4,3,5}, true},
		{[]int{3,7,2,3}, true},
		{[]int{7,7,12,16,41,48,41,48,11,9,34,2,44,30,27,12,11,39,31,8,23,11,47,25,15,23,4,17,11,50,16,50,38,34,48,27,16,24,22,48,50,10,26,27,9,43,13,42,46,24}, true},
	}{
		result := StoneGame(testCase.piles)
		if result != testCase.want{
			t.Errorf("case-%d result=%t want=%t", caseId, result, testCase.want)
			break
		}
		result = StoneGameDP(testCase.piles)
		if result != testCase.want{
			t.Errorf("StoneGameDP case-%d result=%t want=%t", caseId, result, testCase.want)
			break
		}
	}
}

