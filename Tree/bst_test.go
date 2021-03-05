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
