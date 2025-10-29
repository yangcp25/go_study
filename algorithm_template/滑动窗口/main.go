package main

import "fmt"

func main() {
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}

// 最长无重复子串
func lengthOfLongestSubstring(s string) int {
	left, right, res := 0, 0, 0
	m := [128]int{}
	for right < len(s) {
		c := s[right]
		right++
		m[c]++
		for m[c] > 1 {
			d := s[left]
			left++
			m[d]--
		}

		if res < right-left {
			res = right - left
		}
	}
	return res
}
