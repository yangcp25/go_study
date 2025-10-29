package main

func main() {

}

// 53 最大子序和
func maxSubArray(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}

// 198. 打家劫舍
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return Max(nums[0], nums[1])
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = Max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = Max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 139 单词拆分
func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, w := range wordDict {
		wordSet[w] = true
	}

	dp := make(map[int]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// 70. 爬楼梯
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// 0-1 背包问题
func knapsack(weights, values []int, capacity int) int {
	n := len(weights)
	// dp[i][j] 表示前 i 个物品在容量 j 下的最大价值
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= capacity; j++ {
			if j >= weights[i-1] {
				// 可以选择放或不放
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weights[i-1]]+values[i-1])
			} else {
				// 放不下，只能不放
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][capacity]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func knapsackOptimized(weights, values []int, capacity int) int {
	n := len(weights)
	dp := make([]int, capacity+1)

	for i := 0; i < n; i++ {
		for j := capacity; j >= weights[i]; j-- {
			dp[j] = max(dp[j], dp[j-weights[i]]+values[i])
		}
	}
	return dp[capacity]
}

// 300. 最长递增子序列
func lengthOfLIS(nums []int) int {

}
