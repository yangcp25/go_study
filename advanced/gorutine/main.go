package main

import (
	"fmt"
	"time"
)

func main() {

	println("hello world")

	go func() {
		for {
			println("hello worldx")
			time.Sleep(time.Second)
			println("hello world222")
		}
	}()

	fmt.Println("hello world3")
}
