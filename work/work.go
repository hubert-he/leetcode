//
package work

import "sync"

type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg sync.WaitGroup
}

func New(maxGoroutines int) *Pool{
	p := Pool{
		work: make(chan Worker),
	}
	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			// for-range循环一直阻塞，知道从work通道收到一个Worker接口值
			// 若收到一个值，就会执行这个值的Task方法。
			// 一旦woker通道关闭，for-range循环就会结束，并调用Done方法，goroutine终止
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

func (p *Pool) Run(w Worker) {
	p.work <- w
}

func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}