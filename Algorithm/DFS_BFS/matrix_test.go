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
