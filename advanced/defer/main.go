package main

import "fmt"

func main() {
	defer func() {
		fmt.Println(222)
	}()

	defer func() {
		fmt.Println(3333)
	}()
}
