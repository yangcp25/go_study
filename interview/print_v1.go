package main

import (
	"fmt"
	"sync"
)

func main() {

	//g := sync.WaitGroup{}
	//g.Add(100)
	////
	//ch := make(chan int)
	//ch2 := make(chan int)
	//
	//go func() {
	//	for i := range ch {
	//		//fmt.Println("chan 1", i)
	//		i++
	//		ch2 <- i
	//		if i > 100 {
	//			close(ch2)
	//		}
	//		g.Done()
	//	}
	//}()
	//
	//go func() {
	//	for i := range ch2 {
	//		//fmt.Println("chan 2", i)
	//		i++
	//		ch <- i
	//		if i > 100 {
	//			close(ch)
	//		}
	//		g.Done()
	//	}
	//}()
	//
	//ch <- 1
	//
	//g.Wait()
	//close(ch)
	//close(ch2)

	//printV2()
	printV2()
}

func printV2() {
	g := sync.WaitGroup{}
	g.Add(2)

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer g.Done()
		for {
			i, ok := <-ch1
			if !ok {
				close(ch2)
				return
			}
			if i > 100 {
				close(ch2)
				return
			}
			fmt.Println("ch1", i)
			ch2 <- i + 1
		}
	}()

	go func() {
		defer g.Done()
		for {
			i, ok := <-ch2
			if !ok {
				close(ch1)
				return
			}
			if i > 100 {
				close(ch1)
				return
			}
			fmt.Println("ch2", i)
			ch1 <- i + 1
		}
	}()

	ch1 <- 1
	g.Wait()
}

func testV3() {

	wg := sync.WaitGroup{}
	wg.Add(2)
	type msg struct{ num, turn int }

	ch := make(chan msg)

	go func() { // G1
		defer wg.Done()
		for m := range ch {
			if m.num > 100 {
				close(ch) // 统一关闭
				return
			}
			if m.turn == 1 {
				fmt.Println("G1:", m.num)
				ch <- msg{m.num + 1, 2}
			} else {
				ch <- m
			}
		}
	}()

	go func() { // G2
		defer wg.Done()
		for m := range ch {
			if m.num > 100 {
				// chan 已由 G1 关闭，这里不再 close
				return
			}
			if m.turn == 2 {
				fmt.Println("G2:", m.num)
				ch <- msg{m.num + 1, 1}
			} else {
				ch <- m
			}
		}
	}()

	ch <- msg{1, 1}
	wg.Wait()
}
