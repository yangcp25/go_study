package main

import (
	"fmt"
	"sync"
)

func main() {
	//DoubleChanPrint()
	//DoubleChanPrint2()
	DoubleChanPrint3()
}

// 交替打印1-100

func DoubleChanPrint() {
	chan1 := make(chan struct{})
	chan2 := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i < 100; i += 2 {
			<-chan1
			fmt.Println("chan1 : ", i)
			chan2 <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= 100; i += 2 {
			<-chan2
			fmt.Println("chan2 : ", i)
			if i < 100 {
				chan1 <- struct{}{}
			}
		}
	}()

	chan1 <- struct{}{}
	wg.Wait()
	close(chan1)
	close(chan2)
}
func DoubleChanPrint2() {
	chan1 := make(chan int) // 用于发送奇数 (1, 3, 5...)
	chan2 := make(chan int) // 用于发送偶数 (2, 4, 6...)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for v := range chan1 { // 接收奇数，发送偶数
			fmt.Println("chan1 : ", v) // 打印奇数
			if v < 99 {
				chan2 <- v + 1
			} else {
				// v == 99，发送100后，就可以关闭chan2了
				chan2 <- v + 1 // 发送 100
				close(chan2)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for v := range chan2 { // 接收偶数，发送奇数
			fmt.Println("chan2 : ", v) // 打印偶数
			if v < 100 {               // v == 100时，是最后一个数字，不应该再发送了
				chan1 <- v + 1
			} else {
				// v == 100, 是最后一个数字，什么都不用做，
				// 只是等待 chan2 被关闭后，循环自然结束。
				close(chan1)
			}
		}
		// chan2 被关闭后，循环结束，执行 defer wg.Done()
	}()

	chan1 <- 1 // 启动
	wg.Wait()
	//close(chan1)
}
func DoubleChanPrint3() {
	chan1 := make(chan struct{}) // 用于发送奇数 (1, 3, 5...)
	chan2 := make(chan struct{}) // 用于发送偶数 (2, 4, 6...)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i < 100; i += 2 {
			<-chan1
			fmt.Println("chan1 : ", i)
			chan2 <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= 100; i += 2 {
			<-chan2
			fmt.Println("chan2 : ", i)
			if i < 100 {
				chan1 <- struct{}{}
			}
		}
	}()

	chan1 <- struct{}{} // 启动
	wg.Wait()
}
