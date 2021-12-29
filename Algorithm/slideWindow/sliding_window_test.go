package slideWindow

import "testing"

func TestMinWindow(t *testing.T) {
	for caseId, testCase := range []struct{
		s, t	string
		want	string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "b", ""},
		{"a", "aa", ""},
		{"bba", "ab", "ba"},
	}{
		result := MinWindow(testCase.s, testCase.t)
		if result != testCase.want{
			t.Errorf("case-%d Failed: result=%s, but want=%s", caseId, result, testCase.want)
			break
		}
	}
}
