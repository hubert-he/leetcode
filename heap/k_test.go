package heap

import (
	"fmt"
	"runtime"
	"testing"
)

func init(){
	_, codePath, _, _ := runtime.Caller(0)
	fmt.Println(codePath[:len(codePath)-3])
}

func TestSmallestK(t *testing.T) {
}
