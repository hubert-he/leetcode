package Graph

import "testing"

func TestClosedIsland_DFS(t *testing.T) {
	for caseId, testCase := range []struct{
		grid		[][]int
		want		int
	}{
		{
			[][]int{
				[]int{1,1,1,1,1,1,1,0},[]int{1,0,0,0,0,1,1,0},
				[]int{1,0,1,0,1,1,1,0},[]int{1,0,0,0,0,1,0,1},
				[]int{1,1,1,1,1,1,1,0},
			}, 2,
		},
		{
			[][]int{
				[]int{0,0,1,0,0}, []int{0,1,0,1,0},
				[]int{0,1,1,1,0},
			}, 1,
		},
		{
			[][]int{
				[]int{1,1,1,1,1,1,1},[]int{1,0,0,0,0,0,1},
				[]int{1,0,1,1,1,0,1},[]int{1,0,1,0,1,0,1},
				[]int{1,0,1,1,1,0,1},[]int{1,0,0,0,0,0,1},
				[]int{1,1,1,1,1,1,1},
			}, 2,
		},
		{[][]int{
			[]int{0,0,1,1,0,1,0,0,1,0}, []int{1,1,0,1,1,0,1,1,1,0},
			[]int{1,0,1,1,1,0,0,1,1,0}, []int{0,1,1,0,0,0,0,1,0,1},
			[]int{0,0,0,0,0,0,1,1,1,0}, []int{0,1,0,1,0,1,0,1,1,1},
			[]int{1,0,1,0,1,1,0,0,0,1}, []int{1,1,1,1,1,1,0,0,0,0},
			[]int{1,1,1,0,0,1,0,1,0,1}, []int{1,1,1,0,1,1,0,1,1,0},
		}, 5},
	}{
		result := closedIsland_DFS(testCase.grid)
		if result != testCase.want{
			t.Errorf("case-%d: result=%d want=%d", caseId, result, testCase.want)
			break
		}
	}
}
