package main

import (
	"fmt"
	"unicode"
)

func main() {

	// abc1234019A334bc
	// 123s1234
	// aaaaa
	str := ""
	fmt.Scan(&str)

	i, j := 0, 0
	maxL := -1
	for j < len(str) {
		if isLetter(rune(str[j])) {
			j++
			i = j
		} else {
			if i == j {
				maxL = max(maxL, j-i+1)
			} else {
				if j-1 > 0 && str[j] >= str[j-1] {
					maxL = max(maxL, j-i+1)
				}
				if str[j] < str[j-1] {
					maxL = max(maxL, j-i)
					i = j + 1
				}
			}
			j++
		}
	}

	fmt.Println(maxL)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isLetter(s rune) bool {
	return unicode.IsLetter(s)
}
