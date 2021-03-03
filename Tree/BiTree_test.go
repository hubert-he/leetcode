package Tree

import "testing"

func TestHasPathSum(t *testing.T) {
	for _, test := range []struct{
		nums 		[]interface{}
		targetSum	int
		want 		bool
	}{
		{[]interface{}{}, 0, false},
		{[]interface{}{1,2}, 1, false},
		{[]interface{}{5,4,8,11,nil,13,4,7,2,nil,nil,nil,1, 22}, 22, true},
		{[]interface{}{2, 1, 3, nil, 4, nil, 7, nil, nil, 5, 6}, 234, false},
	}{
		tree := GenerateBiTree(test.nums)
		if HasPathSum(tree, test.targetSum) != test.want{
			t.Errorf("tree: %v should has path sum value %d", test.nums, test.targetSum )
		}
	}
}

func TestInvertBiTree(t *testing.T){
	for ii, test := range []struct{
		nums []interface{}
		want []interface{}
	}{
		{[]interface{}{}, []interface{}{}},
		{[]interface{}{1}, []interface{}{1}},
		{[]interface{}{1,2}, []interface{}{1, nil, 2}},
		{[]interface{}{1,2,3,4,5,6,7}, []interface{}{1,3,2,7,6,5,4}},
		{[]interface{}{5,4,8,11,nil,13,4,7,2,nil,nil,nil,1, 22}, []interface{}{5, 8, 4, 4, 13, nil, 11, 1, nil, nil, nil, 2, 7, nil, nil, nil, nil, nil, 22}},
		{[]interface{}{2, 1, 3, nil, 4, nil, 7, nil, nil, 5, 6}, []interface{}{2, 3, 1, 7, nil, 4, nil, 6, 5, nil, nil}},
	}{
		tree := GenerateBiTree(test.nums)
		reversedTree := InvertBiTree(tree)
		result := Serialization(reversedTree)
		for index, value := range test.want{
			if value != result[index]{
				t.Errorf("Invert Error ==> case-%d : %v", ii, result)
				break
			}
		}
	}
}

func TestBinaryTreePaths(t *testing.T) {
	for caseId, testCase := range []struct{
		nums []interface{}
		want []string
	}{
		{[]interface{}{}, nil},
		{[]interface{}{1},[]string{"1"}},
		{[]interface{}{1,2}, []string{"1->2"}},
		{[]interface{}{1,2,3,4,5,6,7}, []string{"1->2->4", "1->2->5", "1->3->6", "1->3->7"}},
		{[]interface{}{1,2,3,nil,5}, []string{"1->2->5", "1->3"}},
	}{
		tree := GenerateBiTree(testCase.nums)
		result := BinaryTreePaths(tree)
		for index, value := range testCase.want {
			if value != result[index]{
				t.Errorf("case-%d: %s | %s ", caseId, result[index], value)
			}
		}
	}
}

func TestSumOfLeftLeaves(t *testing.T) {
	for caseId, testCase := range []struct{
		nums []interface{}
		want int
	}{
		{[]interface{}{}, 0},
		{[]interface{}{1}, 0},
		{[]interface{}{3,9,20,nil,nil,15,7}, 24},
	}{
		tree := GenerateBiTree(testCase.nums)
		result := SumOfLeftLeaves(tree)
		if result != testCase.want{
			t.Errorf("case-%d: result = %d, but want %d", caseId, result, testCase.want)
		}
	}
}

func TestFindTilt(t *testing.T) {
	for caseId, testCase := range []struct{
		nums []interface{}
		want int
	}{
		{[]interface{}{}, 0},
		{[]interface{}{1}, 0},
		{[]interface{}{1,2,3}, 1},
		{[]interface{}{4,2,9,3,5,nil,7}, 15},
		{[]interface{}{21,7,14,1,1,2,2,3,3}, 9},
	}{
		tree := GenerateBiTree(testCase.nums)
		result := FindTilt(tree)
		if result != testCase.want{
			t.Errorf("case-%d: result = %d, but want %d", caseId, result, testCase.want)
		}
	}
}

func TestDiameterOfBinaryTree(t *testing.T) {
	for caseId, testCase := range []struct{
		nums []interface{}
		want int
	}{
		{[]interface{}{}, 0},
		{[]interface{}{1,2,3,4,5}, 3},
	}{
		tree := GenerateBiTree(testCase.nums)
		result := DiameterOfBinaryTree(tree)
		if result != testCase.want{
			t.Errorf("case-%d: result = %d, but want %d", caseId, result, testCase.want)
		}
	}
}
