package main

import (
	"fmt"
	"sync"
)

func main() {
	// 临时对象池
	test2()
}

func test2() {
	wg := sync.WaitGroup{}

	// 计数
	wg.Add(1)
	pool := &sync.Pool{
		New: func() interface{} {
			return "oooo!"
		},
	}

	go add2(pool, wg.Done)

	wg.Wait()

	fmt.Printf("%s", pool.Get())
	fmt.Printf("%s", pool.Get())

	/*pool.Put("test1")
	pool.Put("test2")
	fmt.Printf("%s\n", pool.Get())
	fmt.Printf("%s\n", pool.Get())
	fmt.Printf("%s\n", pool.Get())
	fmt.Printf("%s\n", pool.Get())
	fmt.Printf("%s\n", pool.Get())*/
}

func add2(pool *sync.Pool, done func()) {
	defer func() {
		done()
	}()

	pool.Put("嘻嘻放进去了！")
	pool.Put("嘻嘻放进去了2！")
}
