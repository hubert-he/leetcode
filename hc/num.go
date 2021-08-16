package hc

import (
	"fmt"
	"sync"
	"sync/atomic"
)
var a int64
func main() {
	ch := make(chan int64,100)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		wg1 := sync.WaitGroup{}
		defer wg.Done()
		defer close(ch)
		for i := 0; i <10 ; i++ {
			wg1.Add(1)
			go func() {
				defer wg1.Done()
				for  {
					if atomic.CompareAndSwapInt64(&a,a,a+1) {
						b:= a
						ch <-b
						return
					}
				}
			}()
		}
		wg1.Wait()
	}()
	wg.Wait()
	for aa := range ch {
		fmt.Println(aa)
	}
}