package main

import (
	"fmt"
	"sync"
)

func main() {
	//nums := []int{1, 2, 3, 4, 5, 6}

	ch := make(chan int)
	ch2 := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(19)

	go func() {
		for v := range ch {
			fmt.Println("go 1", v)
			v += 1
			ch2 <- v
			wg.Done()
		}
		//for i := 0; i < len(nums); i++ {
		//	if i%2 == 0 {
		//		fmt.Println("go runtine 1", nums[i])
		//	} else {
		//		ch <- nums[i]
		//	}
		//	wg.Done()
		//}
		//close(ch)
	}()

	go func() {
		for v := range ch2 {
			fmt.Println("go 2", v)
			v += 1
			ch <- v
			wg.Done()
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
		close(ch2)
	}()
	ch <- 1

	select {}
}
