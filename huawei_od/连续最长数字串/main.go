package main

import (
	"fmt"
	"regexp"
)

func main() {

	//1234567890abcd9.+12345.678.9ed
	str := "1234567890abcd9.-1++2+12345.678.9ed"

	// ?0次或者1次 +1次或者多次 * 0次或者多次
	pattern := `[\+-]{0,1}\d+\.{0,1}\d+`

	reg := regexp.MustCompile(pattern)

	res := reg.FindAllString(str, -1)

	fmt.Println(res)
}
