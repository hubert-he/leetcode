package stack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseInPlace(t *testing.T){
	for caseId, testCase := range []struct {
		stack	[]int
		want 	[]int
	}{
		{[]int{},[]int{}},
		{[]int{1}, []int{1}},
		{[]int{1,2}, []int{2,1}},
		{[]int{1,2,3,4,5}, []int{5,4,3,2,1}},
	}{
		result := make([]int, len(testCase.want))
		copy(result, testCase.stack)
		ReverseInPlace(&result)
		ok := assert.Equal(t, testCase.want, result, fmt.Sprintf("case-%d Failed: result=>%v but want=>%v", caseId, result, testCase.want))
		if !ok {
			break
		}
	}
}

func TestStack_Reverse(t *testing.T) {
	for caseId, testCase := range []struct {
		array	[]interface{}
		want 	[]interface{}
	}{
		{[]interface{}{},[]interface{}{nil}},
		{[]interface{}{1}, []interface{}{1}},
		{[]interface{}{1,2}, []interface{}{2,1}},
		{[]interface{}{1,2,3,4,5}, []interface{}{5,4,3,2,1}},
	}{
		st := Construct(testCase.array)
		st.Reverse()
		ok := assert.Equal(t, testCase.want, st.array,
			fmt.Sprintf("case-%d Failed: result=>%v but want=>%v", caseId, st.array, testCase.want))
		if !ok {
			break
		}
	}
}

func TestStack_Push(t *testing.T) {
	for caseId, testCase := range []struct{
		arr		[]interface{}
		want 	[]interface{}
	}{
		{[]interface{}{1}, []interface{}{1}},
		{[]interface{}{1,2}, []interface{}{2,1}},
		{[]interface{}{1,2,3,4,5}, []interface{}{5,4,3,2,1}},
	}{
		st := New()
		for _, tc := range testCase.arr{
			st.Push(tc)
		}
		assert.Equal(t, testCase.want, st.array, fmt.Sprintf("case-%d Failed: array=%v shoule be %v", caseId, st.array, testCase.want))
	}
}

func TestStack_Pop(t *testing.T) {
	for caseId, testCase := range []struct{
		arr		[]interface{}
		times 	int
		want 	[]interface{}
	}{
		{[]interface{}{1}, 1, []interface{}{}},
		{[]interface{}{1,2}, 1, []interface{}{2}},
		{[]interface{}{1,2}, 3, []interface{}{}},
		{[]interface{}{1,2,3,4,5}, 3, []interface{}{4,5}},
	}{
		st := Construct(testCase.arr)
		for i := 0; i < testCase.times; i++{
			st.Pop()
		}
		assert.Equal(t, testCase.want, st.array, fmt.Sprintf("case-%d Failed: array=%v shoule be %v", caseId, st.array, testCase.want))
	}
}

func TestStack_Empty(t *testing.T) {
	for caseId, testCase := range []struct{
		arr		[]interface{}
		want 	bool
	}{
		{[]interface{}{1}, false},
		{[]interface{}{1,2}, false},
		{[]interface{}{}, true},
	} {
		st := Construct(testCase.arr)
		assert.Equal(t, testCase.want, st.Empty(), fmt.Sprintf("case-%d failed: shoule be %t", caseId, testCase.want))
	}
}

func TestStack_Top(t *testing.T) {
		for caseId, testCase := range []struct{
			arr		[]interface{}
			want 	interface{}
		}{
			{[]interface{}{1}, 1},
			{[]interface{}{2,1}, 2},
			{[]interface{}{}, nil},
		}{
			st := Construct(testCase.arr)
			status := assert.Equal(t, testCase.want, st.Top(), fmt.Sprintf("case-%d failed: top=%v shoule be %v", caseId, st.Top(), testCase.want))
			if !status{
				break
			}
		}
}

func TestEvalRPN(t *testing.T) {
	for caseId, testCase := range []struct{
		tokens		[]string
		want		int
	}{
		{[]string{"18"}, 18}, // 特殊的例子
		{[]string{"2","1","+","3","*"}, 9},
		{[]string{"4","13","5","/","+"}, 6},
		{[]string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}, 22},
	}{
		result := EvalRPN(testCase.tokens)
		status := assert.Equal(t, testCase.want, result, fmt.Sprintf("case-%d failed: result=%d want=%d", caseId, result, testCase.want))
		if !status{
			break
		}
	}
}
