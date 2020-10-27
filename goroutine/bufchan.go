// 有缓冲的通道和固定数目的goroutine来处理job
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4 //要使用的goroutine数量
	taskLoad = 10 // 要处理的job数量
)

var wg sync.WaitGroup
// init初始化包，go运行时在其他的代码执行前优先执行这个函数
func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	tasks := make(chan string, taskLoad) // 创建一个有缓冲的通道来管理工作
	wg.Add(numberGoroutines) // 启动goroutine来处理工作
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}
	// 增加一组工作
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task: %d", post)
	}
	// 当所有工作完成 关闭通道，以便所有的goroutine退出
	close(tasks)
	// 等待所有工作完成
	wg.Wait()
}
// worker作为goroutine启动来处理
// 从有缓冲的通道传入的工作
func worker(tasks chan string, worker int) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			// 意味着通道已经清空，并已被关闭
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}
		fmt.Printf("Worker: %d started %s\n", worker, task)
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep)*time.Millisecond)
		fmt.Printf("Worker: %d : Completed %s \n", worker, task)
	}
}
