package string

import "testing"

func TestIsMatch(t *testing.T) {
	for caseId, testCase := range []struct{
		s		string
		p		string
		want	bool
	}{
		{"sa", "*", false},
		{"sa", "*sa", true},
		{"sa", "s*a", true},
		{"sa", "s*sa", true},
		{"", "", true},
		{"", "a", false},
		{"", "a*", true},
		{"", "*", true},
		{"", ".*", true},
		{"", "**", true},
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
	}{
		result := IsMatch(testCase.s, testCase.p)
		if result != testCase.want{
			t.Errorf("case-%d: result=%v but want=%v", caseId, result, testCase.want)
			break
		}
	}
}

func TestIsWildMatch(t *testing.T) {
	for caseId, testCase := range []struct{
		s		string
		p		string
		want	bool
	}{
		{"sa", "*aaa*bbb*ccc", false}, // 0
		{"sa", "*", true},
		{"sa", "*sa", true},
		{"sa", "s*a", true},
		{"sa", "s*sa", false},
		{"", "", true},
		{"", "a", false}, // 6
		{"", "a*", false},
		{"", "*", true},
		{"", "**", true},
		{"aa", "a", false},
		{"aa", "a*", true}, // 11
		{"ab", "?*", true},
		{"abcabczzzde", "*abc???de*", true},
		{"mississippi", "m??*ss*?i*pi", false},
	}{
		result := isWildMatchGreedy(testCase.s, testCase.p)
		if result != testCase.want{
			t.Errorf("case-%d: result=%v but want=%v", caseId, result, testCase.want)
			break
		}
	}
}
