package main

import (
	"./array"
	"fmt"
)

func main(){
	for _, testCase := range []struct{
		str string
		want []string
	}{
		{"a1b2", []string{"a1b2","a1B2","A1b2","A1B2"}},
		{"3z4", []string{"3z4","3Z4"}},
		{"1234", []string{"1234"}},
	}{
		result := array.LetterCasePermutationIII(testCase.str)
		fmt.Println(result)
	}
}