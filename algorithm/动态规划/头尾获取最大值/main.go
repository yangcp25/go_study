package main

import "fmt"

func main() {

	// 只能从数组的头部或者尾部获取数据，求在n次获取之后的最大值
	nums := []int{1, 2, 3, 4, 5, 6, 7, 6, 3, 1}

	res := getMax(nums, 3)

	fmt.Println(res)
}

func getMax(nums []int, k int) int {
	dp := make([]int, k+1)
	dp[0] = 0
	for i := 1; i <= k; i++ {
		for j := 0; j < i; j++ {
			dp[i] = max(dp[i-1]+nums[j], dp[i-1]+nums[len(nums)-j-1])
		}
	}
	return dp[k]
}
