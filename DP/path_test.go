package DP

import "testing"

func TestUniquePaths(t *testing.T){
	for caseID, testCase := range []struct{
		grid [2]int
		want int
	}{
		{[2]int{1, 1}, 1},
		{[2]int{3, 2}, 3},
		{[2]int{3, 7}, 28},
	}{
		result := UniquePaths(testCase.grid[0], testCase.grid[1])
		if result != testCase.want{
			t.Errorf("case-%d: result = %d want = %d", caseID, result, testCase.want)
			break
		}
	}
}

func TestUniquePathsWithObstacles(t *testing.T){
	for caseID, testCase := range []struct{
		grid [][]int
		want int
	}{
		{[][]int{[]int{0}}, 1},
		{[][]int{[]int{1}}, 0},
		{[][]int{[]int{1, 0}}, 0},
		{[][]int{[]int{0,0}, []int{1,1}, []int{0, 0}}, 0},
		{[][]int{[]int{0,1}, []int{0,0}}, 1},
		{[][]int{[]int{0,0,0}, []int{0,1,0}, []int{0,0,0}}, 2},
	}{
		result := UniquePathsWithObstaclesI(testCase.grid)
		if result != testCase.want{
			t.Errorf("case-%d: result = %d want = %d", caseID, result, testCase.want)
			break
		}
	}
}

func TestUniquePathsIII(t *testing.T){
	for caseId, testCase := range []struct{
		grid [][]int
		want int
	}{
		{[][]int{[]int{1,0}, []int{0,2}}, 0},
		{[][]int{[]int{1,-1}, []int{0,2}}, 1},
		{[][]int{[]int{1,-1}, []int{-1,2}}, 0},
		{[][]int{[]int{1,0,0,0}, []int{0,0,0,0}, []int{0,0,2,-1}}, 2},
		{[][]int{[]int{1,0,0,0}, []int{0,0,0,0}, []int{0,0,0,2}}, 4},
	}{
		result := UniquePathsIII(testCase.grid)
		if result != testCase.want{
			t.Errorf("case-%d: result=%d want=%d", caseId, result, testCase.want)
			break
		}
	}
}
