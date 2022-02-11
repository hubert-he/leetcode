package Graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlienOrder(t *testing.T){
	for caseId, testCase := range []struct{
		s		[]string
		want	string
	}{
		{[]string{"wrt","wrf","er","ett","rftt"}, "wertf"},
		{[]string{"z","x"}, "zx"},
		{[]string{"z","x","z"}, ""},
		// 特别case
		{[]string{ "ab","adc" }, "abcd"},
		{[]string{"z","z"}, "z"},
		{[]string{"abc","ab"}, ""},
		{[]string{"aba"}, "ab"},
		{[]string{"wnlb"}, "blnw"},
		{[]string{"vlxpwiqbsg","cpwqwqcd"}, "bdgilpqsvcwx"},
		{[]string{"z","x","a","zb","zx"}, ""},
	}{
		result := AlienOrder(testCase.s)
		if !assert.ElementsMatch(t, testCase.want, result, "case-%d Failed: result=%v but want=%v", caseId, result, testCase.want){
			break
		}
	}
}
