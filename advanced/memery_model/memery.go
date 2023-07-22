package main

import (
	"time"
)

func main() {
	hello()
	time.Sleep(3)
}

var a string

func f() {
	print(a)
}
func hello() {
	go f()
	a = "hello, world"
}

func f2() {
	a = "hello, world"
}
func hello2() {
	go f()
	go f2()
}
