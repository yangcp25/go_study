```go
package main

// 53 最大子序和
func maxSubArray(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if  res < nums[i]{
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
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = Max(nums[0], nums[1])
	// f(x)= max(f(i-1), f(i-2)+f(i))
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

	dp := make([]bool, len(s)+1)
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
	if n < 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func climbStairs2(n int) int {
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

```