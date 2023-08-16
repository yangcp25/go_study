package main

import (
	"fmt"
	"os"
)

func main() {
	test := []byte("中国")
	test2 := []rune("中国")
	fmt.Println(test)
	fmt.Println(test2)
	test3 := []byte{1, 2, 3}
	fmt.Println(test3)

	f, err := os.Create("test.text")
	if err != nil {
		fmt.Println(err)
	}

	_, err = f.Write([]byte{65, 66, 67})
	if err != nil {
		fmt.Println(err)
	}
}
