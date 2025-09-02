package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	// 读取数据
	var str = []string{"hello", "world", "22"}
	check := make(map[string]int)
	for _, v := range str {
		if _, ok := check[v]; ok {
			check[v]++
		} else {
			check[v] = 1
		}
	}
	//sort.Slice(check, func(i, j int) bool {
	//return check[strconv.Itoa(i)] <
	//})
}
