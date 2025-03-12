package main

import (
	"fmt"
	"regexp"
)

// I am an 20-years out--standing @ * -stu- dent
func main() {
	s := ""

	fmt.Scan(&s)

	pattern := "[^0-9a-zA-z\\-]"
	reg := regexp.MustCompile(pattern)

	sSlice := reg.Split(s, -1)

	fmt.Println(sSlice)
}
