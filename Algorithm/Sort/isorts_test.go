package Sort

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {

}
func BenchmarkQuickSort(t *testing.B) {
	testArray := getTestArr()
	for i := 0; i < t.N; i++ {
		QuickSort(testArray)
	}
}

func TestConcurrentQuickSort(t *testing.T) {
	for caseId, testCase := range []struct{
		nums	[]int
		want	[]int
	}{
		{[]int{5,2,3,1}, []int{1,2,3,5}},
		{[]int{5,1,1,2,0,0}, []int{0,0,1,1,2,5}},
	}{
		chanWait := make(chan struct{})
		go ConcurrentQuickSort(testCase.nums, chanWait)
		<- chanWait
		for i := range testCase.nums{
			if testCase.nums[i] != testCase.want[i]{
				t.Errorf("case-%d failed result=%v want=%v", caseId, testCase.nums, testCase.want)
				return
			}
		}
	}
}

func BenchmarkConcurrentQuickSort(b *testing.B) {
	testArray := getTestArr()
	for i := 0; i < b.N; i++ {
		chanWait := make(chan struct{})
		go ConcurrentQuickSort(testArray, chanWait)
		<- chanWait
	}
}

func getTestArr() []int {
	rand.Seed(time.Now().UnixNano())
	numRand := 10000000
	testArr := make([]int, numRand)
	for i := 0; i < numRand; i++ {
		testArr[i] = rand.Intn(numRand * 5)
	}
	return testArr
}

func TestRelativeSortArray(t *testing.T) {
	for caseId, testCase := range []struct{
		arr1	[]int
		arr2	[]int
		want	[]int
	}{
		{[]int{2,3,1,3,2,4,6,7,9,2,19}, []int{2,1,4,3,9,6}, []int{2,2,2,1,4,3,3,9,6,7,19}},
	}{
		result := RelativeSortArray(testCase.arr1, testCase.arr2)
		assert.Equal(t, testCase.want, result, "case-%d result=%v want=%v", caseId, result, testCase.want)
	}
}

func TestReorderLogFiles(t *testing.T) {
	for caseId, testCase := range []struct{
		logs	[]string
		want	[]string
	}{
		{[]string{"let1 art can","let2 own kit dig","let3 art zero"},
			[]string{"let1 art can","let3 art zero","let2 own kit dig"}},
		{[]string{"dig1 8 1 5 1","dig2 3 6"}, []string{"dig1 8 1 5 1","dig2 3 6"}},
		{[]string{"dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"},
			[]string{"let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5 1","dig2 3 6"}},
		{[]string{"a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo"},
			[]string{"g1 act car","a8 act zoo","ab1 off key dog","a1 9 2 3 1","zo4 4 7"}},
	}{
		result := ReorderLogFiles(testCase.logs)
		assert.Equal(t, testCase.want, result, "case-%d result=%v want=%v", caseId, result, testCase.want)
		result = ReorderLogFilesLib(testCase.logs)
		assert.Equal(t, testCase.want, result, "case-%d result=%v want=%v", caseId, result, testCase.want)
	}
}

func TestMergeInPlace(t *testing.T) {
	for caseId, testCase := range []struct{
		A		[]int
		m		int
		B		[]int
		n		int
		want	[]int
	}{
		{[]int{0}, 0, []int{1}, 1, []int{1}},
		{[]int{2,0}, 1, []int{1}, 1, []int{1,2}},
		{[]int{-1,0,0,0,3,0,0,0,0,0,0}, 5, []int{-1,-1,0,0,1,2}, 6, []int{-1,-1,-1,0,0,0,0,0,1,2,3}},
		{[]int{1,2,3,0,0,0}, 3, []int{2,5,6}, 3, []int{1,2,2,3,5,6}},
	}{
		result := make([]int, len(testCase.A))
		copy(result, testCase.A)
		MergeInPlace(result, testCase.m, testCase.B, testCase.n)
		assert.Equal(t, testCase.want, result, "case-%d: result=%v want=%v", caseId, result, testCase.want)
		result = make([]int, len(testCase.A))
		copy(result, testCase.A)
		Merge2Pointer(result, testCase.m, testCase.B, testCase.n)
		assert.Equal(t, testCase.want, result, "case-%d: result=%v want=%v", caseId, result, testCase.want)
	}
}

