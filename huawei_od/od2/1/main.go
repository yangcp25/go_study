package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var m int
	fmt.Scan(&m)
	nums := make([]int, m)
	for i := 0; i < m; i++ {
		var t int
		fmt.Scan(&t)
		nums[i] = t
	}
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 0
	// 前m个程序员n个需求需要的最小时间
	// dp[i][j] = Max(dp[i-1][j], dp[])
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			day := j / nums[i-1]
			if j%nums[i-1] != 0 || j < nums[i-1] {
				day += 1
			}
			dp[i][j] = Max(dp[i-1][j], dp[i-1][j]+day)
		}
	}

	fmt.Println(dp[m][n])
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
