package main

import "fmt"

func main() {
	ans := 0
	n, num := 10, ""
	fmt.Scanln(&n, &num)
	//fmt.Scanln(os.Stdin, "%s", &num)
	for i := range num {
		x := num[i]
		ans *= n
		if x >= '0' && x <= '9' {
			ans += int(x - '0')
		} else {
			ans += int(10 + x - 'A')
		}
	}
	fmt.Println(ans)
}
