package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// asdbuiodevauufgh
	// auauo
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
	// i j 都指向一个元音 + 瑕疵度 == k
	// 如果 j 不是元音 j++ 判断下 如何 i 不是元音 i = j；如果是 i不动 并且 瑕疵+1, 如果findK > k 需要移动i 直到 findk <= k
	// i
	// j 是 判断下 瑕疵度 , =  取max ; j++
	//
	for j < len(str) {
		if IsYuan(string(str[j])) {
			if findK == k {
				maxL = max(maxL, j-i+1)
			}
			j++
		} else {
			if !IsYuan(string(str[i])) {
				j++
				i = j
			} else {
				findK++
				for findK > k {
					if !IsYuan(string(str[i])) {
						findK--
					}
					i++
				}
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
