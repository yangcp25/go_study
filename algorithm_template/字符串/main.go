package main

func main() {

}

// 最长回文子串
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 1

	for i := 0; i < len(s); i++ {
		// 奇数长度回文
		len1 := expandAroundCenter(s, i, i)
		// 偶数长度回文
		len2 := expandAroundCenter(s, i, i+1)

		currMax := max(len1, len2)
		if currMax > maxLen {
			maxLen = currMax
			start = i - (currMax-1)/2
		}
	}

	return s[start : start+maxLen]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
