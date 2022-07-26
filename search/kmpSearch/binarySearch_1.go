package main

import (
	"fmt"
)

// 查找算法
// 字符串查找《KMP匹配》（）
func main() {
	s := "Hello,22 杨春坪!"
	p := "杨春坪"
	pos := strStrV1(s, p)
	fmt.Printf("Find \"%s\" at %d in \"%s\"\n", p, pos, s)
}

func strStrV1(s string, p string) int {
	n, m := len(s), len(p)
	i, j := 0, 0
	next := getNext(p)
	for i < n && j < m {
		if j == -1 || s[i] == p[j] {
			j++
			i++
		} else {
			j = next[j]
		}
	}
	if j == m {
		return i - j
	}
	return -1
}

func getNext(p string) []int {
	m := len(p)
	next := make([]int, m, m)
	next[0] = -1
	next[1] = 0
	i, j := 0, 1
	for j < m-1 {
		if i == -1 || p[i] == p[j] {
			j++
			i++
			next[j] = i
		} else {
			i = next[j]
		}
	}
	return next
}
