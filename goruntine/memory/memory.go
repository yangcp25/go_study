package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 协程
	// 使用共享内存传递消息
	test1()
	//未使用协程
	test2()

}

func test2() {
	start := time.Now()

	for i := 0; i < 1; i++ {
		add2(i)
	}
	end := time.Now()

	timeLong := end.Sub(start).Seconds()

	fmt.Println("未使用协程:\n", timeLong)
}
func add2(a int) {
	count++
	c := a + 1
	fmt.Printf("a+1 = %v \n", c)
}

var count int = 0

func test1() {
	start := time.Now()

	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go add(i, lock)
	}

	for {
		lock.Lock()
		c := count
		lock.Unlock()
		if c >= 10 {
			break
		}
	}
	end := time.Now()

	timeLong := end.Sub(start).Seconds()

	fmt.Println("\n", timeLong)
}
func add(a int, lock *sync.Mutex) {
	lock.Lock()
	count++
	c := a + 1
	fmt.Printf("a+1 = %v \n", c)
	lock.Unlock()
}
