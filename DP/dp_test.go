package DP

import "testing"

func Test_MaxProfitBrust(t *testing.T){
	for id, testCase := range []struct{
		prices []int
		want int
	}{
		{[]int{7,1,5,3,6,4}, 7},
		{[]int{1,2,3,4,5}, 4},
		{[]int{7,6,4,3,1}, 0},
	}{
		result := MaxProfitBrust(testCase.prices)
		if result != testCase.want{
			t.Errorf("case-%d failed result = %d want = %d", id, result, testCase.want)
		}
	}
}
