package main

import "fmt"

type animal interface {
	move()
}
type dog struct{}

func (d dog) move() {
	fmt.Println("dog moving")
}

type cat struct{}

func (c cat) move() {
	fmt.Println("cat moving")
}

func main() {
	var a animal
	a = dog{}
	a.move()
	a = cat{}
	a.move()
}
