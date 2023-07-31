package main

import (
	"fmt"
	"strings"
)

// Handler函数类型
type Handler func(string)

// Logger middleware
func Logger(next Handler) Handler {
	fmt.Println("进入Logger")
	return func(s string) {
		println("输入参数:", s)
		next(s)
	}
}

// Uppercase middleware
func Uppercase(next Handler) Handler {
	return func(s string) {
		s = strings.ToUpper(s)
		next(s)
	}
}

// Handler函数
func print(s string) {
	println("输出:", s)
}

func main() {
	a := Uppercase(print)
	Logger(a)
	//h("hello")
}
