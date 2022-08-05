package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// withGroup
	//test()

	// Once
	// test2
	//test2()

	// content包
	//testContent()
	testContent2()

	// 原子操作
	//test3()
}

func testContent2() {
	total := 10

	var num int32
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	valueCtx := context.WithValue(ctx, "key", "val")
	defer cancelFunc()
	for i := 0; i < total; i++ {
		go addContent(&num, i, func() {
			if atomic.LoadInt32(&num) == int32(total) {
				fmt.Printf("key:%s\n", valueCtx.Value("key"))
				cancelFunc()
			}
		})
	}
	select {
	case <-ctx.Done():
		fmt.Print("结束了！")
	}
}
func test3() {
	var x int32 = 32
	y := atomic.LoadInt32(&x)
	fmt.Printf("%d\n", y)

	var a int32 = 1
	var b int32 = 2
	var c int32 = 1

	atomic.CompareAndSwapInt32(&a, c, b)
	atomic.CompareAndSwapInt32(&a, b, c)
	fmt.Printf("a:%d,b:%d,c:%d,\n", a, b, c)

	var d int32 = 1
	var e int32
	atomic.StoreInt32(&e, atomic.LoadInt32(&d))
	fmt.Printf("d:%d,e:%d\n", d, e)

	var f int32 = 1
	var j int32 = 3
	old := atomic.SwapInt32(&f, j)
	fmt.Printf("new:%d,old:%d\n", f, old)

	var v atomic.Value

	v.Store(100)

	fmt.Printf("v:%d\n", v.Load())
}

// content 包
func testContent() {
	total := 10

	var num int32
	ctx, cancelFunc := context.WithCancel(context.Background())
	for i := 0; i < total; i++ {
		go addContent(&num, i, func() {
			if atomic.LoadInt32(&num) == int32(total) {
				cancelFunc()
			}
		})
	}
	select {
	case <-ctx.Done():
		fmt.Print("结束了！")
	}
}

func addContent(a *int32, b int, cancelFunc func()) {
	defer func() {
		cancelFunc()
	}()
	for i := 0; ; i++ {
		curNum := atomic.LoadInt32(a)
		newNum := curNum + 1
		time.Sleep(200 * time.Millisecond)
		if atomic.CompareAndSwapInt32(a, curNum, newNum) {
			fmt.Printf("number当前值:%d [%d -%d]\n", *a, b, i)
			break
		}

	}
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
