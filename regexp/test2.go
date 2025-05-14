package main

import (
	"fmt"
	"regexp"
)

func main() {
	rege := regexp.MustCompile(`^(2[0-4][0-9])$|^(25[0-5])$`)
	res := rege.FindAllStringSubmatch("256", -1)

	fmt.Println(res)

	rege2 := regexp.MustCompile(`(^[0-9]|^1[0-9][0-9]?|^2[0-4][0-9]|^25[0-5])\.([0-9]|1[0-9][0-9]?|2[0-4][0-9]|25[0-5])`)
	res2 := rege2.FindAllStringSubmatch("1.", -1)

	fmt.Println(res2)
}
