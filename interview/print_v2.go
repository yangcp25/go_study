package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
			if i%2 == 1 {
				fmt.Println(i)
			}
		}
	}()
	go func() {
		for i := 1; i <= 100; i++ {
			<-ch
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
	}()

	for {
	}
}
