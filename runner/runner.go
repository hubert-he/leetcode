package runner

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"
)
/*
  单向通道：
  1. 类型 chan<- int: 只能发送的通道		只写
  2. 类型 <-chan int: 只能接收的int类型通道 只读
 */
type Runner struct{
	interrupt	chan os.Signal
	complete	chan error
	timeout	<-chan time.Time // <-chan time.Time  表示timeout是一个chan 类型，且chan 是只可读的，单向接收channel 中的类型是Time类型，允许接收但是不能发送
	tasks	[]func(int)
}

var ErrTimeout = errors.New("recv timeout")

var ErrInterrupt = errors.New("recv interrupt")

// Runner的工厂函数
func New(d time.Duration) *Runner {
	return &Runner{
		// 通道interrupt被初始化为缓冲区容量为1的通道，可以保证通道至少能接收一个来自go-runtime的os.Signal，确保runtime不会被阻塞。
		// 如果goroutine没有准备好接收这个值，这个值就会被丢弃
		interrupt: make(chan os.Signal, 1),
		// 无缓冲通道，当执行任务的goroutine完成时，会向这个通道发送一个error类型的值或nil值。之后会等待main函数接收这个值。
		// 一旦main接收了这个error值,goroutine就可以安全地终止
		complete: make(chan error),
		timeout: time.After(d), // time.After()表示time.Duration后返回一条time.Time类型的channel 消息,可以实现定时器
	}
}
// 可变参数
func (r *Runner) Add(tasks ...func(int)){
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)
	go func() {
		r.complete <- r.run()
	}()

	select { // select 多路复用，可以指定接收和发送情况，一般是指定接收情况，判断有通道数据流入 通道全空的情况，select执行等待
	case err:= <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return  nil
}

func (r *Runner) gotInterrupt() bool {
	// 一般来说，select语句在没有任何要接收的数据时会阻塞，但是
	// 有了default分支后就不会阻塞了。
	// default分支会将interrupt通道有os.Signal需要接收，就会接收并处理
	// 如果没有需要接收的os.Signal，就会执行default分支
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

func (r *Runner) TestSelect() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
			fmt.Printf("send-write to chan %d\n", i)
		}
	}
}

func (r *Runner) TestPolling(abort chan int) {
	select {
	case <-abort:
		fmt.Printf("about")
		return
	default: // select 尝试从abort通道中接收一个值，如果无值什么也不做。这是一个非阻塞的接收操作，重复这个动作称为对通道的轮询
	}
}
