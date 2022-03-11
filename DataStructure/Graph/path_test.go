package Graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortestAlternatingPaths(t *testing.T) {
	for caseId, testCase := range []struct{
		n			int
		redEdges	[][]int
		blueEdges	[][]int
		want		[]int
	}{
		{3, [][]int{[]int{0,1}, []int{0,2}}, [][]int{[]int{1,0}},
			[]int{0,1,1},
		}, // 此 case 测试 环的情况处理
		{5,
			[][]int{[]int{0,1}, []int{1,2}, []int{2,3}, []int{3,4}},
			[][]int{[]int{1,2},[]int{2,3},[]int{3,1}},
			[]int{0,1,2,3,7},
		}, // 此case 用来验证 4 节点循环导入路径
		{5,
			[][]int{[]int{3,2}, []int{4,1}, []int{1,4}, []int{2,4}},
			[][]int{[]int{2,3},[]int{0,4},[]int{4,3}, []int{4,4}, []int{4,0}, []int{1, 0}},
			[]int{0,2,-1,-1,1},
		}, // 不是以红色开始的 情况，也即需要你比较 红 蓝 2种情况的
		{3, [][]int{[]int{0,1}, []int{1,2}}, [][]int{}, []int{0,1,-1}},
		{3, [][]int{[]int{0,1}}, [][]int{[]int{2,1}}, []int{0,1,-1}},
	}{
		result := shortestAlternatingPaths_DFS(testCase.n, testCase.redEdges, testCase.blueEdges)
		if !assert.Equal(t, testCase.want, result,
			"case-%d: result=%v, but want=%v", caseId, result, testCase.want){
			break
		}
	}
}
