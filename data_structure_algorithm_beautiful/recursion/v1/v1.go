package main

import "fmt"

func main() {
	sum := getSum(3)
	fmt.Printf("%d阶的总走法是%d\n", 3, sum)

	sum = getSum(5)
	fmt.Printf("%d阶的总走法是%d\n", 5, sum)
}

// 递归求解一次性走1阶或者2阶 所有的走法
func getSum(n int) (sum int) {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return getSum(n-1) + getSum(n-2)
}
