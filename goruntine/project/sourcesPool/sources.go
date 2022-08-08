package main

import (
	"io"
	"log"
	"math/rand"
	"sourcesPool/pool"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines = 10
	maxResources  = 2
)

var idCounter int32

type dbConnect struct {
	Id int32
}

func (d dbConnect) Close() error {
	log.Println("Close:", "关闭的Id", d.Id)
	return nil
}

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("createConnection:", "新增链接ID", id)

	return &dbConnect{
		id,
	}, nil
}

func main() {
	// 使用withGroup 协同 多个协程状态
	wg := &sync.WaitGroup{}
	wg.Add(maxGoroutines)
	poolObj, err := pool.New(createConnection, maxResources)

	if err != nil {
		log.Println(err)
	}

	for i := 0; i < maxGoroutines; i++ {
		go func(i int) {
			// 模拟sql查询
			performQuery(i, poolObj)
			wg.Done()
		}(i)
	}

	wg.Wait()

	log.Println("查询跑完了！")

	wg.Add(maxGoroutines)

	// 为了试一下 使用到通道资源
	log.Println("开跑第二次！")

	for i := 0; i < maxGoroutines; i++ {
		go func(i int) {
			// 模拟sql查询
			performQuery(i, poolObj)
			wg.Done()
		}(i)
	}
	wg.Wait()

	poolObj.Close()
}

//
func performQuery(i int, obj *pool.Pool) {
	conn, err := obj.Acquire()

	if err != nil {
		log.Println(err)
		return
	}

	defer obj.Release(conn)
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	log.Printf("查询ID[%d]:数据库连接ID[%d]\n", i, conn.(*dbConnect).Id)
}
