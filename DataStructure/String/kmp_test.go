package String

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNext(t *testing.T) {
	for caseId, testCase := range []struct{
		p		string
		want	[]int
	}{
		{"abab", []int{-1,0,-1,0}},
		//{"abab", []int{-1,0,0,1}},
		{"ABABCABAA",[]int{-1, 0, 0, 1, 2, 0, 1, 2, 3}},
		{"ABCDABD", []int{-1, 0, 0, 0, 0, 1, 2}},
	}{
		result := getNextBest_1(testCase.p)
		if !assert.Equal(t, testCase.want, result,
			fmt.Sprintf("case-%d result: %v want=%v", caseId, result, testCase.want)){
			break
		}
	}
}

func TestKmpMatch(t *testing.T){
	for caseId, testCase := range []struct{
		s		string
		p		string
		want	bool
	}{
		{"BBC ABCDAB ABCDABCDABDE", "ABCDABD", true},
	}{
		result := kmpMatch(testCase.s, testCase.p, getNext_0)
		if result != testCase.want{
			t.Fatalf("case-%d result: %v want=%t", caseId, result, testCase.want)
		}
	}
}
