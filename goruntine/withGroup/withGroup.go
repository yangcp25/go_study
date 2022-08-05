package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// withGroup
	//test()

	// Once
	// test2
	//test2()

	// content包

}

func test2() {
	once := &sync.Once{}

	go doThins(once)
	go doThins(once)
	time.Sleep(1 * time.Second)
}

func doThins(once *sync.Once) {
	fmt.Println("开始!")
	once.Do(func() {
		fmt.Println("来了中间一次!")
	})
	fmt.Println("结束!")
}

func test() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		add(i, wg.Done)
	}
	wg.Wait()
}

func add(i int, done func()) {
	defer func() {
		done()
	}()
	c := i + 1
	fmt.Printf("c:%d\n", c)
}
