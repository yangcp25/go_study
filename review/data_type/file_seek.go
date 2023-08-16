package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.text")
	if err != nil {
		fmt.Println(err)
	}

	n, err := f.Write([]byte("123"))
	if err != nil {
		fmt.Println(err)
	}
	f.Write([]byte("4"))
	f.Write([]byte("5"))
	f.Seek(0, 0)
	f.Write([]byte("0"))
	f.Seek(0, 2)
	f.Write([]byte("6"))
	fmt.Println(n)
}
