package main

import "fmt"

func main() {
	res := make([]int, 0)
	for i := 0; i < 8; i++ {
		res = append(res, GenFibonacciSequence(i))
	}

	fmt.Println(res)
}

// 斐波那契
func GenFibonacciSequence(n int) int {
	if n <= 1 {
		return n
	}
	return GenFibonacciSequence(n-1) + GenFibonacciSequence(n-2)
}
