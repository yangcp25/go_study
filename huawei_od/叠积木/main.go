package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	text := input.Text()
	str := strings.Split(text, " ")
	nums := make([]int, 0)
	for _, v := range str {
		temp, _ := strconv.Atoi(v)
		nums = append(nums, temp)
	}

	// 每层的最大个数
	maxH := len(nums)

	for i := 1; i <= maxH; i++ {

	}

}
