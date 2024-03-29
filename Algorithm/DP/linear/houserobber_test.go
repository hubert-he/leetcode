package linear

import "testing"
import "../../../DataStructure/Tree"

func TestRobDP(t *testing.T){
	for caseId, testCase := range []struct{
		nums []int
		want int
	}{
		{[]int{0}, 0},
		{[]int{1}, 1},
		{[]int{1,2,3,1}, 4},
		{[]int{2,7,9,3,1}, 12},
	}{
		result := RobI(testCase.nums)
		if result != testCase.want{
			t.Errorf("RobI-case-%d Failed: result=%d, but want=%d", caseId, result, testCase.want)
		}
		result = RobI1015(testCase.nums)
		if result != testCase.want{
			t.Errorf("RobI1015-case-%d Failed: result=%d, but want=%d", caseId, result, testCase.want)
		}
	}
}

func TestRobIIDP(t *testing.T){
	for caseId, testCase := range []struct{
		nums []int
		want int
	}{
		{[]int{0}, 0},
		{[]int{1}, 1},
		{[]int{2,3,2}, 3},
		{[]int{1,2,3,1}, 4},
		{[]int{1,2,3}, 3},
	}{
		result := RobII(testCase.nums)
		if result != testCase.want{
			t.Errorf("RobI-case-%d Failed: result=%d, but want=%d", caseId, result, testCase.want)
		}
		result = RobII1015(testCase.nums)
		if result != testCase.want{
			t.Errorf("RobI1015-case-%d Failed: result=%d, but want=%d", caseId, result, testCase.want)
		}
	}
}

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

func TestMinCostIIIBFS(t *testing.T){
	for caseId, testCase := range []struct{
		houses []int
		cost [][]int
		size [2]int
		neigh int
		want int
	}{
		{[]int{2,2,1}, [][]int{[]int{1,1}, []int{3, 4}, []int{4, 2}}, [2]int{3, 2}, 2, 0},
		{[]int{0,2,1,2,0}, [][]int{[]int{1,10}, []int{10, 1}, []int{10, 1}, []int{1,10}, []int{5, 1}}, [2]int{5, 2}, 3, 11},
		{[]int{0,0,0,0,0}, [][]int{[]int{1,10}, []int{10, 1}, []int{1, 10}, []int{10, 1}, []int{1,10}}, [2]int{5, 2}, 5, 5},
		{[]int{0,0,0,0,0}, [][]int{[]int{1,10}, []int{10, 1}, []int{10, 1}, []int{1,10}, []int{5, 1}}, [2]int{5, 2}, 3, 9},
		{[]int{3,1,2,3}, [][]int{[]int{1,1,1}, []int{1,1,1}, []int{1,1,1}, []int{1,1,1}}, [2]int{4, 3}, 3, -1},
	}{
		/*
		result := minCostIII(testCase.houses, testCase.cost, testCase.size[0], testCase.size[1], testCase.neigh)
		if result != testCase.want{
			t.Errorf("case-%d: result2=%d want=%d", caseId, result, testCase.want)
			break
		}
		 */
		result := MinCostIIIDFSDP(testCase.houses, testCase.cost, testCase.size[0], testCase.size[1], testCase.neigh)
		if result != testCase.want{
			t.Errorf("case-%d: result=%d want=%d", caseId, result, testCase.want)
			break
		}
	}
}