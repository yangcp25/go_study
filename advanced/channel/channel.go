package main

var c = make(chan int)
var a string

func main() {
	go f()
	c <- 0
	print(a)
}

func f() {
	a = "hello, world"
	<-c
}
