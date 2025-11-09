package main

import (
	"fmt"
	"sync"
)

func main() {

	//// 生产者消费者
	////ctx, cancel := context.WithCancel(context.Background())
	//
	//ch := make(chan int, 10)
	//wg := sync.WaitGroup{}
	//// 消费者
	//for i := 0; i < 3; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		defer wg.Done()
	//		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	//		defer cancel()
	//		for {
	//			select {
	//			case <-ctx.Done():
	//				fmt.Println("timeout", i)
	//				return
	//			case i, ok := <-ch:
	//				if !ok {
	//					return
	//				}
	//				fmt.Println("worker", i)
	//			}
	//		}
	//	}(i)
	//}
	//// 生产者
	//
	//go func() {
	//	for i := 0; i < 100; i++ {
	//		ch <- i
	//	}
	//	close(ch)
	//}()
	//wg.Wait()
	//fmt.Println("all done")

	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i += 2 {
			<-ch1
			fmt.Println("A:", i)
			ch2 <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= 100; i += 2 {
			<-ch2
			fmt.Println("B:", i)
			if i < 100 {
				ch1 <- struct{}{}
			}
		}
	}()

	// start
	ch1 <- struct{}{}

	wg.Wait()
}
