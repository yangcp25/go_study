package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// auuuuecfd
	// xauuuuecfd
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	k, _ := strconv.Atoi(input.Text())

	input.Scan()
	str := input.Text()

	str = strings.ToLower(str)

	i, j := 0, 0
	findK := 0
	maxL := 0
	for j < len(str) {
		if i != j {
			if !IsYuan(string(str[j])) && findK == k {
				findK = 0
				j++
				i = j
			}
			if IsYuan(string(str[j])) && findK == k {
				maxL = max(maxL, i-j+1)
			}

			if !IsYuan(string(str[j])) {
				findK++
			}
		} else {
			if !IsYuan(string(str[j])) {
				j++
				i = j
			} else {
				maxL = max(maxL, i-j+1)
				j++
			}
		}
	}

	fmt.Println(maxL)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func IsYuan(s string) bool {
	if s == "a" || s == "e" || s == "i" || s == "o" || s == "u" {
		return true
	}
	return false
}
