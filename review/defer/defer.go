package main

import "fmt"

func main() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

}
func test(a int) int {
	fmt.Println(a)
	return a
}
