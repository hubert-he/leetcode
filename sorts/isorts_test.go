package sorts

import (
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

