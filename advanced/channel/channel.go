package main

import (
	"fmt"
	"sort"
	"time"
)

//var c = make(chan int)
//var a string
//
//func main() {
//	//go f()
//	//c <- 0
//	//print(a)
//	test2()
//}
//
//func f() {
//	a = "hello, world"
//	<-c
//}
//
//func test2() {
//	ch := make(chan int, 5)
//	ch <- 18
//	ch <- 18
//	close(ch)
//	x, ok := <-ch
//	if ok {
//		fmt.Println("received: ", x)
//	}
//
//	x, ok = <-ch
//	if ok {
//		fmt.Println("received: ", x)
//	}
//
//	x, ok = <-ch
//	if !ok {
//		fmt.Println("received fail: ", x)
//	}
//}

//	func main() {
//		ch := make(chan int)
//		go readChan(ch) // 先启动接收者
//		ch <- 10        // 发送数据，此时接收者已经在等待接收
//		time.Sleep(time.Second * 2)
//	}
//
//	func readChan(ch chan int) {
//		for {
//			val, ok := <-ch
//			fmt.Println("read ch: ", val)
//			if !ok {
//				break
//			}
//		}
//	}
//
// 代码片段5
func main() {
	ch := make(chan int)

	go writeChan(ch)

	for {
		val, ok := <-ch
		fmt.Println("read ch: ", val)
		if !ok {
			break
		}
	}

	sort.Sort()
	slices.SortI

	time.Sleep(time.Second)
	fmt.Println("end")
}

func writeChan(ch chan int) {
	for i := 0; i < 4; i++ {
		ch <- i
	}
}
