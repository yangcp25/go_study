package main

import (
	"fmt"
	"regexp"
)

func main() {
	match := regexp.MustCompile(`^([ADWS])([1-9][0-9]?)$`)
	res := match.FindAllStringSubmatch("A10", -1)
	fmt.Println(res)
	res = match.FindAllStringSubmatch("A9", -1)
	fmt.Println(res)
}

func MatchIP() {
	rege := regexp.MustCompile(`^0\.[0-]`)
}
