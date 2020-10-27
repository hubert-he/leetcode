// 创建一个goroutine池并完成工作
package main

import (
	"./work"
	"log"
	"sync"
	"time"
)

const submitCnt = 10

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
}

// Task实现Worker接口
func (m *namePrinter) Task(){
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main() {
	p := work.New(2)
	var wg sync.WaitGroup
	wg.Add(submitCnt * len(names))

	for i := 0; i < submitCnt; i++ {
		// 迭代names切片
		for _, name := range names {
			// 创建一个namePrinter并提供指定的名字
			np := namePrinter{
				name: name,
			}
			go func() {
				// 将任务提交执行，当Run返回时 我们就知道任务已经处理完成了
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	p.Shutdown()
}


