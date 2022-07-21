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
	for i := 0; i < (m - n + 1); i++ {
		j := 0
		for ; j < n; j++ {
			if s[i+j] != p[j] {
				break
			} else {
				// 最后一个
				if j == n-1 {
					return i
				}
			}
		}
	}
	return -1
}
