package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	for {
		n, _ := fmt.Scan(&str)
		if n == 0 {
			break
		} else {
			fmt.Println(str)
			strSlice := strings.Split(str, " ")
			fmt.Println(len(strSlice[len(strSlice)-1]))
			//return
		}
	}
}
