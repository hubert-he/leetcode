package design

import "testing"
/*
["Skiplist","add","add","add","search","add","search","erase","erase","search"]
[[],[1],[2],[3],[0],[4],[1],[0],[1],[1]]
 */
func TestSkipList(t *testing.T) {
	for caseID, testCase := range []struct {
		operatation []string
		data        [][]int
	}{
		{[]string{"Skiplist", "add", "add", "add", "search", "add", "search", "erase", "erase", "search"},
			[][]int{[]int{}, []int{1}, []int{2}, []int{3}, []int{0}, []int{4}, []int{1}, []int{0}, []int{1}, []int{1}}},
	} {
		var skl Skiplist
		for i, u := range testCase.operatation{
			v := testCase.data[i][0]
			switch u {
			case "Skiplist":
				skl = ConstructorSkipList()
			case "add":
				skl.Add(v)
			case "search":
				skl.Search(v)
			case "erase":
				skl.Erase(v)
			}
		}
	}
}
