package unionSet

import "testing"

func TestNumIslandsDFS(t *testing.T) {
	for caseId, testCase := range []struct {
		grid	[][]byte
		want	int
	}{
		{[][]byte{[]byte{'1','1','1'},[]byte{'0','1','0'},[]byte{'1','1','1'}}, 1},
		{[][]byte{[]byte{'1','1','1','1','0'},[]byte{'1','1','0','1','0'},[]byte{'1','1','0','0','0'},[]byte{'0','0','0','0','0'}}, 1},
		{[][]byte{[]byte{'1','1','0','0','0'},[]byte{'1','1','0','0','0'},[]byte{'0','0','1','0','0'},[]byte{'0','0','0','1','1'}}, 3},
	}{
		result := NumIslands(testCase.grid)
		if result != testCase.want{
			t.Errorf("case-%d: result = %d but want = %d", caseId, result, testCase.want)
			return
		}
	}
}
