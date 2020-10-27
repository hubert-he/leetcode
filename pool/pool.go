// 管理用户定义的一组资源
package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool管理一组可以安全地在多个goroutine间共享的资源。被管理的资源必须实现io.Closer接口
type Pool struct {
	m sync.Mutex
	resources chan io.Closer // 用来保存共享的资源
	closed	bool
	factory func()(io.Closer, error)
}
var ErrPoolClosed = errors.New("Pool has been closed.")

// 创建一个用来管理资源的池，这个池需要一个可以分配新资源的函数，并规定池大小
func New(fn func()(io.Closer, error), size uint) (*Pool, error){
	if size <= 0 {
		return nil, errors.New("Size value too small.")
	}
	return &Pool{
		factory: fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// Acquire从池中获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("Acquire: ", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default: // 无空闲资源，提供一个新资源
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}
}
// 将使用后的资源放回池中
func (p *Pool) Release(r io.Closer) {
	// 保证本操作和Close操作的安全
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}
	select {
	case p.resources<- r: // 试图将资源放入队列
		log.Println("Release", "In Queue")
	default: // pool 满 则关闭
		log.Println("Release:", "Closing")
		r.Close()
	}
}

func (p *Pool) Close(){
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed{
		return
	}
	p.closed = true
	// 在清空通道资源之前，将通道关闭
	// 如果不这样做，会发生deadlock
	close(p.resources)
	// for-range循环会一直阻塞，知道channel关闭，因此要先关闭通道
	for r:= range p.resources {
		r.Close()
	}

}



















