package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//str1 := "a5"
	//str4 := "aBB9"
	//str3 := "abcdef"
	//str4 := "ab4a1c1111"
	str4 := "ab4a1117a777a77777"
	//str4 := "7777"
	//str4 := "7777a"
	//str4 := "aaaa"

	test1 := getMaxLenSub(str4)
	fmt.Println(test1)
}

func getMaxLenSub(str string) int {
	maxL := -1
	i, j := 0, 0
	str = strings.ToLower(str)
	findLetter := 0
	for j < len(str) {
		if isLetter(str[j]) {
			findLetter++
		}

		for findLetter > 1 {
			if isLetter(str[i]) {
				findLetter--
			}
			i++
		}

		if findLetter == 1 {
			l := j - i + 1
			if l >= 2 {
				maxL = max(maxL, l)
			}
		}
		j++
	}

	return maxL
}

func isLetter(str byte) bool {
	return unicode.IsLetter(rune(str))
}
