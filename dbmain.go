// 如何使用pool 来共享一组模拟的数据库连接
package main

import (
	"./pool"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines = 25
	pooledResources = 2
)

type dbConnection struct {
	ID int32
}

// Close实现io.Closer接口，以便DBConnection可以被池管理
// close用来完成任意资源的释放管理
func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

// idCounter 用来给每个连接分配一个独一无二的id
var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: new connection", id)
	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxGoroutines; query++ {
		// 每个goroutine需要自己复制一份要查询值的副本，不然
		// 所有的查询会共享同一个查询变量
		// 因此不能直接闭包变量query
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("Shutdown Program")
	p.Close()
}

func performQueries(query int, p *pool.Pool){
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer p.Release(conn)
	time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
	log.Printf("QID [%d] CID [%d]\n", query, conn.(*dbConnection).ID) // 通过类型断言拿到元素值
}