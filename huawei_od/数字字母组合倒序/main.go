package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// I am an 20-years out--standing @ * -stu- dent
func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	s := input.Text()

	pattern := "[^0-9a-zA-z\\-]"
	reg := regexp.MustCompile(pattern)

	sSlice := reg.Split(s, -1)

	fmt.Println(sSlice)
}
