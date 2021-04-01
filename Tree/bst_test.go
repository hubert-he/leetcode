package Tree

import "testing"

func TestBinarySearchTree_GetMinimumDifference(t *testing.T) {
	for caseId, testCase := range []struct{
		nums []interface{}
		want int
	}{
		{[]interface{}{}, 0},
		{[]interface{}{1}, 0},
		{[]interface{}{1,nil, 3}, 2},
		{[]interface{}{1,nil,3,2}, 1},
	}{
		tree := NewBSTFromPlainList(testCase.nums)
		result := tree.GetMinimumDifference()
		if result != testCase.want{
			t.Errorf("case-%d: result is %d should be %d", caseId, result, testCase.want)
		}
	}
}

func TestBinarySearchTree_FindMode(t *testing.T) {
	for caseId, testCase := range []struct{
		nums []interface{}
		want []int
	}{
		{[]interface{}{}, nil},
		{[]interface{}{1}, []int{1}},
		{[]interface{}{1,3}, []int{1,3}},
		{[]interface{}{1,nil,2,2}, []int{2}},
		{[]interface{}{0,nil,2,2}, []int{2}},
	}{
		tree := NewBSTFromPlainList(testCase.nums)
		result := tree.FindMode()
		if len(result) != len(testCase.want) {
			t.Errorf("case-%d: result is %d should be %d", caseId, result, testCase.want)
			continue
		}
		for idx, value := range result{
			if value != testCase.want[idx]{
				t.Errorf("case-%d: result is %d should be %d", caseId, result, testCase.want)
				break
			}
		}

		result = tree.FindModeMorris()
		if len(result) != len(testCase.want) {
			t.Errorf("case-%d: result is %d should be %d", caseId, result, testCase.want)
			continue
		}
		for idx, value := range result{
			if value != testCase.want[idx]{
				t.Errorf("case-%d: result is %d should be %d", caseId, result, testCase.want)
				break
			}
		}
	}
}

func TestDeleteNode(t *testing.T){
	for caseId, testCase := range []struct{
		nums 	[]interface{}
		key 	int
		want 	[]interface{}
	}{
		{[]interface{}{}, 0, []interface{}{}},
		{[]interface{}{1}, 0, []interface{}{1}},
		{[]interface{}{1}, 1, []interface{}{}},
		{[]interface{}{2, 1}, 2, []interface{}{1}},
		{[]interface{}{2, 1}, 1, []interface{}{2}},
		{[]interface{}{50,30,70,nil,40,60,80}, 50,[]interface{}{60,30,70,nil,40,nil,80}}, // 删root
		{[]interface{}{5,3,6,2,4,nil,7}, 7,[]interface{}{5,3,6,2,4}}, // 删叶子
		{[]interface{}{5,3,6,2,4,nil,7}, 3,[]interface{}{5,4,6,2,nil,nil,7}}, // 删中间节点
		{[]interface{}{5,3,6,2,nil,nil,7}, 3,[]interface{}{5,2,6,nil,nil,nil,7}}, // 删中间节点,没有右子树
		{[]interface{}{10,5,11,2,9,nil,nil,nil,nil,7,nil,nil,8}, 5, []interface{}{10,7,11,2,9,nil,nil,nil,nil, 8}},
		{[]interface{}{5,3,6,2,4,nil,7}, 0,[]interface{}{5,3,6,2,4,nil,7}}, // 删不存在可以
	}{
		tree := NewBSTFromPlainList(testCase.nums)
		tree.DeleteNode2(testCase.key)
		result := Serialization(tree.root)
		if len(result) != len(testCase.want){
			t.Errorf("case-%d: result is %d should be %d", caseId, result, testCase.want)
			break
		}
		for idx, value := range result{
			if value != testCase.want[idx]{
				t.Errorf("case-%d: result is %d should be %d", caseId, result, testCase.want)
				break
			}
		}
	}
}

func TestConvertBiNode(t *testing.T) {
	for caseId, testCase := range []struct{
		nums []interface{}
		want []interface{}
	}{
		{[]interface{}{}, []interface{}{}},
		{[]interface{}{1}, []interface{}{1}},
		{[]interface{}{1,nil, 3}, []interface{}{1, nil, 3}},
		{[]interface{}{4,2,5,1,3,nil,6,0}, []interface{}{0, nil, 1, nil, 2, nil, 3, nil, 4, nil, 5, nil,6}},
	}{
		tree := NewBSTFromPlainList(testCase.nums)
		tree.ConvertBiNode()
		result := Serialization(tree.root)
		for idx, value := range result{
			if value != testCase.want[idx]{
				t.Errorf("case-%d: result is %d should be %d", caseId, result, testCase.want)
			}
		}
	}
}

func TestClosestValue(t *testing.T)  {
	for caseId, testCase := range []struct{
		nums []interface{}
		target float64
		want int
	}{
	//	{[]interface{}{}, },
	//	{[]interface{}{1}, },
		{[]interface{}{4,2,5,1,3}, 3.714286, 4},
		{[]interface{}{2,0,33,nil,1,25,40,nil,nil,11,31,34,45,10,18,29,32,nil,36,43,46,4,nil,12,24,26,30,nil,nil,35,39,42,44,nil,48,3,9,nil,14,22,nil,nil,27,nil,nil,nil,nil,38,nil,41,nil,nil,nil,47,49,nil,nil,5,nil,13,15,21,23,nil,28,37,nil,nil,nil,nil,nil,nil,nil,nil,8,nil,nil,nil,17,19,nil,nil,nil,nil,nil,nil,nil,7,nil,16,nil,nil,20,6}, 0.428571, 0},
		{[]interface{}{28,12,45,4,24,35,47,2,9,14,25,31,42,46,48,0,3,8,11,13,20,nil,26,30,33,41,43,nil,nil,nil,49,nil,1,nil,nil,7,nil,10,nil,nil,nil,17,22,nil,27,29,nil,32,34,36,nil,nil,44,nil,nil,nil,nil,6,nil,nil,nil,16,18,21,23,nil,nil,nil,nil,nil,nil,nil,nil,nil,37,nil,nil,5,nil,15,nil,nil,19,nil,nil,nil,nil,nil,40,nil,nil,nil,nil,nil,nil,39,nil,38}, 2.00000, 2},
	} {
		tree := NewBSTFromPlainList(testCase.nums)
		result := tree.ClosestValue(testCase.target)

		if result == nil || result.Val.(int) != testCase.want{
			t.Errorf("case-%d: result is %v should be %d", caseId, result, testCase.want)
		}
	}
}
