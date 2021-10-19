package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main(){
	fmt.Println(ReorderLogFilesLib([]string{"dig1 8 1 5 1","dig2 3 6"}))
}
func ReorderLogFilesLib(logs []string) []string {
	less := func(i, j int)bool{ // 传递索引，比直接传string要优秀
		part_i := strings.SplitN(logs[i], " ", 2)
		part_j := strings.SplitN(logs[j], " ", 2)
		isDigit_i := unicode.IsDigit(rune(part_i[1][0]))
		isDigit_j := unicode.IsDigit(rune(part_j[1][0]))
		if !isDigit_i && !isDigit_j{
			return part_i[1] < part_j[1] || (part_i[1] == part_j[1] && part_i[0] < part_j[0])
		}
		return !isDigit_i && isDigit_j
	}
	sort.SliceStable(logs, less)
	return logs
}