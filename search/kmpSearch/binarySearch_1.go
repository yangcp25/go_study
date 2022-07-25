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
	m, n := len(s), len(p)
	i, j := 0, 0
	next := make([]int, n)
	next = getNext(p)
	for i < m {
		if j == -1 || s[i] == p[j] {
			j++
			i++
		} else {
			j = next[j]
		}
	}
	return -1
}

func getNext(p string) []int {

}
