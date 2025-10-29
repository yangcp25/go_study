package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 7, 7, 7, 9, 10, 11, 12, 13, 14, 15}
	target := 7
	// 5 7

	start, end := -1, -1
	for k, num := range nums {
		if target == num && start == -1 {
			start = k
		}
		if target == num {
			end = k
		}
	}

	fmt.Println(start, end)

}
