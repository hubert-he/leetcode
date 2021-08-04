package DP

import "testing"
import "../Tree"

func TestRob(t *testing.T){
	for caseId, testCase := range []struct{
		nums []interface{}
		want int
	}{
		{[]interface{}{3,2,3,nil,3,nil,1}, 7},
		{[]interface{}{3,4,5,1,3,nil,1}, 9},
		{[]interface{}{3,10,1,1,3,nil,5}, 15},
	}{
		tree := Tree.GenerateBiTree(testCase.nums)
		result := Rob(tree)
		if result != testCase.want{
			t.Errorf("case-%d Failed: result=%d, but want=%d", caseId, result, testCase.want)
		}
	}
}

func TestMinCostI(t *testing.T){
	for caseId, testCase := range []struct{
		nums [][3]int
		want int
	}{
		{[][3]int{[3]int{17,2,17},[3]int{16,16,5}, [3]int{14,3,19}}, 10},
	}{
		result := MinCostI(testCase.nums)
		if result != testCase.want{
			t.Errorf("case-%d Failed: result=%d, but want=%d", caseId, result, testCase.want)
		}
		result = MinCostI2(testCase.nums)
		if result != testCase.want{
			t.Errorf("case-%d Failed: result=%d, but want=%d", caseId, result, testCase.want)
		}
	}
}