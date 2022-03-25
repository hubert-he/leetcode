package Math

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	for caseId, testCase := range []struct{
		nums []int
		want []int
	}{
		{[]int{9}, []int{1,0}},
		{[]int{1,2,3}, []int{1,2,4}},
		{[]int{4,3,2,1}, []int{4,3,2,2}},
		{[]int{0}, []int{1}},
	}{
		result := PlusOne(testCase.nums)
		ok := assert.Equal(t, testCase.want, result,
			"case-%d: result=%v but want=%v", caseId, result, testCase.want)
		if !ok{
			break
		}
	}
}

func TestAddBinary(t *testing.T) {
	for caseId, testCase := range []struct{
		binstring [2]string
		want string
	}{
		{[2]string{"1010", "1011"}, "10101"},
		{[2]string{"10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
			"110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"},
			"110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000",
		},
	}{
		result := AddBinary(testCase.binstring[0], testCase.binstring[1])
		ok := assert.Equal(t, testCase.want, result,
			"case-%d: result=%v but want=%v", caseId, result, testCase.want)
		if !ok{
			break
		}
		result = AddBinary2(testCase.binstring[0], testCase.binstring[1])
		ok = assert.Equal(t, testCase.want, result,
			"case-%d: result=%v but want=%v", caseId, result, testCase.want)
		if !ok{
			break
		}
		result = AddBinaryBitBig(testCase.binstring[0], testCase.binstring[1])
		ok = assert.Equal(t, testCase.want, result,
			"case-%d: result=%v but want=%v", caseId, result, testCase.want)
		if !ok{
			break
		}
	}
}

func TestCountPrimes(t *testing.T){
	for caseId, testCase := range []struct{
		n	int
		want int
	}{
		{10, 4},
		{1, 0},
		{2, 0},
		{5000000, 348513},
	}{
		result := CountPrimes(testCase.n)
		if result != testCase.want{
			t.Errorf("case-%d: result = %d want=%d", caseId, result, testCase.want)
		}
		result = CountPrimesEratosthenes(testCase.n)
		if result != testCase.want{
			t.Errorf("CountPrimesEratosthenes-%d: result = %d want=%d", caseId, result, testCase.want)
		}
	}
}

func TestPermutation(t *testing.T){
	for caseId, testCase := range []struct{
		nums	[]int
		want 	[][]int
	}{
		{[]int{1,2,3}, [][]int{
			[]int{1,2,3}, []int{1,3,2},
			[]int{2,1,3}, []int{2,3,1},
			[]int{3,1,2}, []int{3,2,1},
		}},
		{[]int{1,1,3}, [][]int{
			[]int{1,1,3}, []int{1,3,1},	[]int{3,1,1},
		}},
		{[]int{3,2,2,3}, [][]int{
			[]int{3,2,2,3}, []int{3,2,3,2},
			[]int{3,3,2,2}, []int{2,3,2,3},
			[]int{2,3,3,2}, []int{2,2,3,3},
		}},
	}{
		result := permutation(testCase.nums)
		if !assert.ElementsMatchf(t, testCase.want, result, "case-%d: faild, result=%#v, \n want=%#v",
			caseId, result, testCase.want){
			break
		}
	}
}
