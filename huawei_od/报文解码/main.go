package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	// 3[k]2[mn]  3[m12[c]]
	str := ""
	fmt.Scan(&str)
	// 数字 [ ]字符

	num := ""
	queue := make([]rune, 0)
	for _, s := range str {
		if string(s) != "]" {
			queue = append(queue, s)
		} else {
			temp := make([]string, 0)
			for i := len(queue) - 1; i >= 0; i-- {
				if string(queue[i]) == "[" {
					queue = queue[:i]
					break
				} else {
					temp = append(temp, string(queue[i]))
				}
			}
			d := make([]int, 0)
			for i := len(queue) - 1; i >= 0; i-- {
				if IsLetter(rune(byte(queue[i]))) {
					queue = queue[:i+1]
					break
				} else {
					dI, _ := strconv.Atoi(string(queue[i]))
					d = append(d, dI)
				}
			}
			sum := 0
			for k, v := range d {
				sum += int(float64(v) * math.Pow(float64(10), float64(k-1)))
			}
			strs := strings.Repeat(strings.Join(temp, ""), sum)

		}
	}
}

func IsNum(s rune) bool {
	return unicode.IsDigit(s)
}

func IsLetter(s rune) bool {
	return unicode.IsLetter(s)
}
