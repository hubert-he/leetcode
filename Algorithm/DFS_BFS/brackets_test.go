package DFS_BFS

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBraceExpansionII(t *testing.T){
	for caseId, testCase := range []struct{
		s		string
		want	[]string
	}{
		{"{a,b}{c,{d,e}}", []string{"ac","ad","ae","bc","bd","be"}},
		{"{a,b}c{d,e}f", []string{"acdf","acef","bcdf","bcef"}},
		{"{{a,z},a{b,c},{ab,z}}", []string{"a","ab","ac","z"}},
	}{
		result := BraceExpansionII_2(testCase.s)
		if !assert.Equal(t, testCase.want, result, "case-%d Failed: result=%v but want=%v", caseId, result, testCase.want){
			break
		}
	}
}

func TestRemoveInvalidParentheses(t *testing.T){
	for caseId, testCase := range []struct{
		s		string
		want	[]string
	}{
		{"()", []string{"()"}},
		{")(", []string{""}},
		{")(f", []string{"f"}}, // 检验去重
		{"(a)())()", []string{"(a())()","(a)()()"}},
		{"()())()", []string{"(())()","()()()"}},
		{"((()((s((((()", []string{"()s()","()(s)","(()s)"}}, // 超时case
	}{
		result := removeInvalidParentheses_BFS(testCase.s)
		if !assert.ElementsMatch(t, testCase.want, result, "case-%d Failed: result=%v but want=%v", caseId, result, testCase.want){
			break
		}
	}
}