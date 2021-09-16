package DP

import (
	"testing"
)

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

func TestMinimumTotal(t *testing.T){
	for caseId, testCase := range []struct{
		nums		[][]int
		want		int
	}{
		{[][]int{[]int{2}, []int{3,4}, []int{6,5,7}, []int{4,1,8,3}}, 11},
		{[][]int{[]int{-10}}, -10},
	}{
		result := MinimumTotalDFS(testCase.nums)
		if result != testCase.want{
			t.Errorf("case-%d: result=%d want=%d", caseId, result, testCase.want)
			break
		}
	}
}

func TestMinFallingPathSum(t *testing.T){
	for caseId, testCase := range []struct{
		nums		[][]int
		want		int
	}{
		{[][]int{[]int{-48}}, -48},
		{[][]int{[]int{-48, 2}}, -48},
		{[][]int{[]int{1}, []int{2}}, 3},
		{[][]int{[]int{2,1,3}, []int{6,5,4}, []int{7,8,9}}, 13},
		{[][]int{[]int{-19, 57}, []int{-40, -5}}, -59},
	}{
		result := MinFallingPathSumDFS(testCase.nums)
		if result != testCase.want{
			t.Errorf("case-MinFallingPathSumDFS-%d: result=%d want=%d", caseId, result, testCase.want)
			break
		}
		result = MinFallingPathSum(testCase.nums)
		if result != testCase.want{
			t.Errorf("case-MinFallingPathSum-%d: result=%d want=%d", caseId, result, testCase.want)
			break
		}
	}
}

func TestMinFallingPathSumII(t *testing.T) {
	for caseId, testCase := range []struct{
		nums		[][]int
		want		int
	}{
		{[][]int{[]int{7}}, 7},
		{[][]int{[]int{1,2,3}, []int{4,5,6}, []int{7,8,9}}, 13},
	} {
		result := MinFallingPathSumII(testCase.nums)
		if result != testCase.want{
			t.Errorf("case-MinFallingPathSumII-%d: result=%d want=%d", caseId, result, testCase.want)
			break
		}
	}
}

func TestFindPaths(t *testing.T) {
	for caseId, testCase := range []struct{
		tc			[]int
		want		int
	}{
		{[]int{36,5,50,15,3}, 390153306},
		{[]int{8,7,16,1,5}, 102984580},
		{[]int{2,2,2,0,0}, 6},
		{[]int{1,3,3,0,1}, 12},
	} {
		if caseId != 0{
			result := FindPaths(testCase.tc[0], testCase.tc[1], testCase.tc[2], testCase.tc[3], testCase.tc[4])
			if result != testCase.want{
				t.Errorf("case-FindPaths-%d: result=%d want=%d", caseId, result, testCase.want)
				break
			}
		}
		result := FindPathsDFSDP(testCase.tc[0], testCase.tc[1], testCase.tc[2], testCase.tc[3], testCase.tc[4])
		if result != testCase.want{
			t.Errorf("case-FindPathsDFSDP-%d: result=%d want=%d", caseId, result, testCase.want)
			break
		}
		result = FindPathsDPBest(testCase.tc[0], testCase.tc[1], testCase.tc[2], testCase.tc[3], testCase.tc[4])
		if result != testCase.want{
			t.Errorf("case-FindPathsDFSDP-%d: result=%d want=%d", caseId, result, testCase.want)
			break
		}
	}
}

func TestPathsWithMaxScore(t *testing.T) {
	for caseId, testCase := range []struct{
		board		[]string
		want		[2]int
	}{
		{[]string{"E23","2X2","12S"}, [2]int{7,1}},
		{[]string{"E12","1X1","21S"}, [2]int{4,2}},
		{[]string{"E11","XXX","11S"}, [2]int{0,0}},
	}{
		result := PathsWithMaxScore(testCase.board)
		if result[0] != testCase.want[0] || result[1] != testCase.want[1]{
			t.Errorf("case-PathsWithMaxScore-%d: result=%d want=%d", caseId, result, testCase.want)
			break
		}
	}
}
