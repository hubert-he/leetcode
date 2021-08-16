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
		{[][]int{[]int{1,0,0,0}, []int{0,0,0,0}, []int{0,0,2,-1}}, 2},
		{[][]int{[]int{1,0}, []int{0,2}}, 0},
		{[][]int{[]int{1,-1}, []int{0,2}}, 1},
		{[][]int{[]int{1,-1}, []int{-1,2}}, 0},
		{[][]int{[]int{1,0,0,0}, []int{0,0,0,0}, []int{0,0,0,2}}, 4},
	}{
		result := UniquePathsIII1(testCase.grid)
		if result != testCase.want{
			t.Errorf("case-%d: result=%d want=%d", caseId, result, testCase.want)
			break
		}
	}
}

func TestCountRoutesDFS(t *testing.T){
	for caseId, testCase := range []struct{
		location	[]int
		path		[3]int
		want		int
	}{
		{[]int{22,74,92,86,12,68,64,19,79,10,69,13,62,18,87,88,33,96,78,73,57,42,91,17,55,26,27,67,60,46,72,41}, [3]int{30,29,47}, 535415296},
		{[]int{2,1,5}, [3]int{0,0,3}, 2},
		{[]int{2,3,6,8,4}, [3]int{1,3,5}, 4},
		{[]int{4,3,1}, [3]int{1,0,6}, 5},
		{[]int{5,2,1}, [3]int{0,2,3}, 0},
		{[]int{1,2,3}, [3]int{0,2,40}, 615088286},
	}{
		result := CountRoutesDP(testCase.location, testCase.path[0], testCase.path[1], testCase.path[2])
		if result != testCase.want{
			t.Errorf("case-%d: result=%d want=%d", caseId, result, testCase.want)
			break
		}
	}
}
