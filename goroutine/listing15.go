// 展示如何使用atomic包里的Store和Load类函数来提供对数值类型的安全访问
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	_ "sync/atomic"
	"time"
)

var (
	// shutdown 是通知正在执行的goroutine 停止工作 的标志
	shutdown  int64
	wg sync.WaitGroup
)

func main(){
	runtime.GOMAXPROCS(8)
	wg.Add(2)
	go doWork("A")
	go doWork("B")

	// 给定goroutine执行的时间
	time.Sleep(1 * time.Second)
	// 该停止工作，安全设置shutdown标志
	fmt.Println("shutdown now")
	// shutdown = 1
	atomic.StoreInt64(&shutdown, 1)
	wg.Wait()
}
// 用来模拟执行工作的goroutine
// 检测之前的shutdown标志来决定是否提前终止
func doWork(name string){
	defer wg.Done()
	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(500 * time.Millisecond)
		// if shutdown == 1{
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}
