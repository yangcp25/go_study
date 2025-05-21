package main

import (
	"fmt"
	"regexp"
)

func main() {
	rege2 := regexp.MustCompile(`((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})(\.((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})){3}`)
	res := rege2.FindAllStringSubmatch("10.255.11.11", -1)

	fmt.Println(res)
}
