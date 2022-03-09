package Test

import (
	"fmt"
	"testing"
)
// 算法复杂度考察，此题的时间复杂度为 O(根号N)
func  TestAlg(t *testing.T)  {
	const N = 100
	i, cnt := 0, 0
	for cnt < N{
		fmt.Println(cnt)
		i += 2
		cnt += i
	}
}
