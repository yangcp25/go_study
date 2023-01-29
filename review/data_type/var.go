package main

import "fmt"

func main() {
	test := uint(255)

	test2 := int8(test)

	fmt.Println(test2)

	testStr := []byte{'a', 'c', 't'}

	fmt.Println(testStr)

	testStr2 := []rune{'a', '这'}
	fmt.Println(string(testStr2))
}
