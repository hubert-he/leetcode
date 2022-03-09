package DFS_BFS

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWallsAndGates_DFS(t *testing.T) {
	for caseId, testCase := range []struct{
		rooms	[][]int
		want	[][]int
	}{
		{[][]int{[]int{2147483647,0,2147483647,2147483647,0,2147483647,-1,2147483647}},
			[][]int{[]int{1,0,1,1,0,1,-1,2147483647}}},
		{  },
	}{
		WallsAndGates_DFS(testCase.rooms)
		if assert.Equal(t, testCase.want, testCase.rooms,
			fmt.Sprintf("case-%d: result=%v, want = %v", caseId, testCase.rooms, testCase.want)){
			break
		}
	}
}
//cmd: go test -v matrix.go matrix_test.go helper.go -test.run TestShortestPathBinaryMatrix
func TestShortestPathBinaryMatrix(t *testing.T){
	for caseId, testCase := range []struct{
		grid	[][]int
		want	int
	}{
		{[][]int{[]int{0,1}, []int{1, 0}}, 2}, // 常规测试case
		{[][]int{[]int{0,0,0},[]int{1,1,0},[]int{1,1,0}}, 4}, // 常规测试case
		{[][]int{[]int{1,0,0}, []int{1,1,0}, []int{1,1,0}}, -1}, // 边界条件测试
		{[][]int{[]int{0,0,0,0,1}, []int{1,0,0,0,0}, []int{0,1,0,1,0}, []int{0,0,0,1,1}, []int{0,0,0,1,0}},
			-1}, // 常规测试case： 首尾截断不通的情况
		{[][]int{[]int{0}}, 1}, // 边界条件测试
		{[][]int{[]int{0,1,1,0,0,0}, []int{0,1,0,1,1,0},	[]int{0,1,1,0,1,0},
			[]int{0,0,0,1,1,0}, []int{1,1,1,1,1,0}, []int{1,1,1,1,1,0}}, 14}, // 注意此测试例子，特殊方向例子
	}{
		result := shortestPathBinaryMatrix_BFS(testCase.grid)
		//result := shortestPathBinaryMatrix_DP(testCase.grid)
		if assert.Equal(t, testCase.want, result,
			fmt.Sprintf("case-%d: result=%v, want = %v", caseId, result, testCase.want)){
			break
		}
	}
}
