package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println(runtime.NumCPU())
	// 如果配置大于1的逻辑处理器给goroutine的话，goroutine会并行运行(前提是多个物理处理器)
	runtime.GOMAXPROCS(1) // 分配一个逻辑处理器给调度器使用
	var wg sync.WaitGroup // wg用来等待程序完成，WaitGroup是一个计数信号量，记录并维护运行的goroutine，值大于0则wait方法就会阻塞
	wg.Add(2) // 计数加2， 表示要等待两个goroutine

	fmt.Println("Start Goroutines")
	go func() {
		// 在gorouting函数退出的时候调用Done来通知main函数工作已经完成
		defer wg.Done() // 递减WaitGroup信号量值
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}

	}()

	go func(){
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A' + 26; char++{
				fmt.Printf("%c ", char)
			}
		}

	}()

	// 等待goroutine结束
	fmt.Println("Waiting To finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
