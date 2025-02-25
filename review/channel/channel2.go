package main

import (
	"fmt"
)

func main() {
	//nums := []int{0, 1, 2, 3, 4}
	ch := make(chan int, 1000)
	exit := make(chan struct{})
	//wg := sync.WaitGroup{}
	//wg.Add(2)
	go func() {
		for i := 0; i < 100000; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		//for v := range ch {
		//	fmt.Println(v)
		//	wg.Done()
		//}
		//defer wg.Done()
		for {
			if v, ok := <-ch; ok {
				fmt.Println(v)
			} else {
				break
			}
		}
		exit <- struct{}{}
	}()
	<-exit
}
