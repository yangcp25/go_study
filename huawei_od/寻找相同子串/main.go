package main

import (
	"fmt"
)

func main() {

	// ppptt   pt
	// AVERDXIVYERDIAN   RDXI
	// AVERDVYERDXIIAN   RDXI
	// RORROX   ROX
	// ROROROXOX   ROROX

	t, p := "", ""
	_, err := fmt.Scan(&t, &p)
	if err != nil {
		return
	}

	// i 字符的开始 j 指向 匹配的结尾
	// 如果 p[j-i] = p[j] i, j++ ; j-i+1<len(p) break
	// 如果 p[j-i] != p[j] ; i = j, p[j-i] != p[j] , i = j, j++, i = j
	// i

	i, j := 0, 0
	for j < len(t) && j-i+1 < len(p) {
	}

	if j-i+1 == len(p) {
		fmt.Println(i + 1)
	} else {
		fmt.Println("No")
	}

}
