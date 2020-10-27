// 展示如何在程序里造成竞争状态
// go build -race 可以检查竞态
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64 // counter 是所有goroutine都要增加其值的变量
	wg sync.WaitGroup
	mutex sync.Mutex
)

func main() {
	wg.Add(2)
	// 竞争出现
	// go incCounter(1)
	// go incCounter(2)

	// 加锁避免竞争
	// go incCounterAtomic(1)
	// go incCounterAtomic(2)

	// mutex
	go incCounterMutex(1)
	go incCounterMutex(2)

	wg.Wait()
	fmt.Println("Final Counter: ", counter)
}

func incCounter(id int){
	defer wg.Done()
	for count := 0; count < 2; count++ {
		value := counter // 捕获counter的值

		runtime.Gosched() // 当前goroutine从线程退出，并放回到队列

		value++
		counter = value
	}
}
// 原子函数能够以很底层的加锁机制来同步访问整型变量和指针
func incCounterAtomic(id int){
	defer wg.Done()
	for count := 0; count < 2; count++ {
		// 安全地对counter加1
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}

func incCounterMutex(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		// 同一时刻只允许一个goroutine进入此临界区
		mutex.Lock()
		{ // 使用大括号 只是为了让临界区看起来更清晰
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock() // 释放锁，允许其他goroutine进入
	}
}
