// 使用互斥锁来定义一段需要同步访问的代码临界区资源的同步访问
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	cwg sync.WaitGroup
)

func init(){
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// 创建一个无缓冲的通道
	/*
	court := make(chan int)
	cwg.Add(2)

	go player("Nadal", court)
	go player("Djokovic", court)

	court <- 1
	*/
	baton := make(chan int)
	cwg.Add(1)
	go Runner(baton)
	baton <- 1

	cwg.Wait()
}

func player(name string, court chan int){
	// 在函数退出时调用Done来通知main函数工作已经完成
	defer cwg.Done()
	for {
		ball, ok := <-court
		if !ok { // 检查通道是否被关闭，游戏结束
			fmt.Printf("Player %s won\n", name)
			return
		}
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			close(court)
			return
		}
		fmt.Printf("Player %s Hit %d \n", name, ball)
		ball++
		court <- ball
	}
}

func Runner(baton chan int){
	var newRunner int
	runner := <-baton // 等待接力棒

	fmt.Printf("Runner %d Running with Baton\n", runner)

	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", runner)
		go Runner(baton)
	}
	time.Sleep(100 * time.Millisecond)
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		cwg.Done()
		return
	}
	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)
	baton <- newRunner
}

