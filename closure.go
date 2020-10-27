package main

import (
	"fmt"
	"sync"
)

func main() {
	var j int = 5
	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()
	a()
	j *= 2
	a()
	var wg sync.WaitGroup // wg用来等待程序完成
	wg.Add(6)
	// false, print 3 3 3
	values := []int{1,2,3}

	for _, val := range values {
		go func() {
			fmt.Println(val)
			defer wg.Done()
		}()
	}
	// true, print 1 2 3
	values = []int{1,2,3}
	for _, val := range values {
		go func(val interface{}) {
			fmt.Println(val)
			defer wg.Done()
		}(val)
	}
	fmt.Println("Waiting print")
	wg.Wait()
}