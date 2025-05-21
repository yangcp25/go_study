package main

import (
	"fmt"
	"regexp"
)

func main() {
	rege := regexp.MustCompile(`((\w){2,})\1`)

	test1 := rege.FindAllStringSubmatch("xxabxxab", -1)
	fmt.Println(test1)

	rege = regexp.MustCompile(``)

	test1 = rege.FindAllStringSubmatch("xxabxxab", -1)
	fmt.Println(test1)
}
