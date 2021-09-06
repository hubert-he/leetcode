package design

import "testing"
/*
["Skiplist","add","add","add","search","add","search","erase","erase","search"]
[[],[1],[2],[3],[0],[4],[1],[0],[1],[1]]
["Skiplist","add","add","add","add","add","add","add","add","add","erase","search","add","erase","erase","erase","add","search","search","search","erase","search","add","add","add","erase","search","add","search","erase","search","search","erase","erase","add","erase","search","erase","erase","search","add","add","erase","erase","erase","add","erase","add","erase","erase","add","add","add","search","search","add","erase","search","add","add","search","add","search","erase","erase","search","search","erase","search","add","erase","search","erase","search","erase","erase","search","search","add","add","add","add","search","search","search","search","search","search","search","search","search"]
[[],[16],[5],[14],[13],[0],[3],[12],[9],[12],
[3],[6],[7],[0],[1],[10],[5],[12],[7],[16],
[7],[0],[9],[16],[3],[2],[17],[2],[17],[0],
[9],[14],[1],[6],[1],[16],[9],[10],[9],[2],
[3],[16],[15],[12],[7],[4],[3],[2],[1],[14],
[13],[12],[3],[6],[17],[2],[3],[14],[11],[0],
[13],[2],[1],[10],[17],[0],[5],[8],[9],[8],
[11],[10],[11],[10],[9],[8],[15],[14],[1],[6],
[17],[16],[13],[4],[5],[4],[17],[16],[7],[14],
[1]]
 */
func TestSkipList(t *testing.T) {
	for caseID, testCase := range []struct {
		operatation []string
		data        [][]int
		want		[]interface{}
	}{
		{[]string{"Skiplist", "add", "add", "add", "search", "add", "search", "erase", "erase", "search"},
			[][]int{[]int{}, []int{1}, []int{2}, []int{3}, []int{0}, []int{4}, []int{1}, []int{0}, []int{1}, []int{1}},
		[]interface{}{true, true, true, true, false, true, true, false, true, false}},
		{[]string{"Skiplist","add","add","add","add","add","add","add","add","add",
			"erase","search","add","erase","erase","erase","add","search","search","search",
			"erase","search","add","add","add","erase","search","add","search","erase",
			"search","search","erase","erase","add","erase","search","erase","erase","search",
			"add","add","erase","erase","erase","add","erase","add","erase","erase",
			"add","add","add","search","search","add","erase","search","add","add",
			"search","add","search","erase","erase","search","search","erase","search","add",
			"erase","search","erase","search","erase","erase","search","search","add","add",
			"add","add","search","search","search","search","search","search","search","search",
			"search"},
			[][]int{[]int{}, []int{16}, []int{5}, []int{14}, []int{13}, []int{0}, []int{3}, []int{12}, []int{9}, []int{12},
				[]int{3}, []int{6}, []int{7}, []int{0}, []int{1}, []int{10}, []int{5}, []int{12}, []int{7}, []int{16},
				[]int{7}, []int{0}, []int{9}, []int{16}, []int{3}, []int{2}, []int{17}, []int{2}, []int{17}, []int{0},
				[]int{9}, []int{14}, []int{1}, []int{6}, []int{1}, []int{16}, []int{9}, []int{10}, []int{9}, []int{2},
				[]int{3}, []int{16}, []int{15}, []int{12}, []int{7}, []int{4}, []int{3}, []int{2}, []int{1}, []int{14},
				[]int{13},[]int{12},[]int{3},[]int{6},[]int{17},[]int{2},[]int{3},[]int{14},[]int{11},[]int{0},
				[]int{13},[]int{2},[]int{1},[]int{10},[]int{17},[]int{0},[]int{5},[]int{8},[]int{9},[]int{8},
				[]int{11},[]int{10},[]int{11},[]int{10},[]int{9},[]int{8},[]int{15},[]int{14},[]int{1},[]int{6},
				[]int{17},[]int{16},[]int{13},[]int{4},[]int{5},[]int{4},[]int{17},[]int{16},[]int{7},[]int{14},[]int{1},
			},
			[]interface{}{nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,
				true,false,nil,true,false,false,nil,true,true,true,
				true,false,nil,nil,nil,false,false,nil,false,false,
				true,true,false,false,nil,true,true,false,true,true,
				nil,nil,false,true,false,nil,true,nil,true,true,
				nil,nil,nil,false,false,nil,true,false,nil,nil,
				true,nil,false,false,false,true,true,false,true,nil,
				true,false,false,false,true,true,false,false,nil,nil,
				nil,nil,true,true,true,true,true,true,false,false,true}},
	} {
		var skl Skiplist
		for i, u := range testCase.operatation{
			switch u {
			case "Skiplist":
				skl = ConstructorSkipList()
			case "add":
				v := testCase.data[i][0]
				skl.Add(v)
			case "search":
				v := testCase.data[i][0]
				if skl.Search(v) != testCase.want[i]{
					t.Errorf("case-%d-%s-%d: result=%t, want=%t", caseID, u, v, skl.Search(v), testCase.want[i])
					skl.Print()
					return
				}
			case "erase":
				v := testCase.data[i][0]
				result := skl.Erase(v)
				if result != testCase.want[i]{
					t.Errorf("case-%d-%s-%d: result=%t, want=%t", caseID, u, v, skl.Search(v), testCase.want[i])
					return
				}
			}
		}
	}
}
