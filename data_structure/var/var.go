package main

import "fmt"

// 指针 和变量
func main() {
	/*a := 1
	b := &a

	fmt.Printf("%v \n", b)
	fmt.Printf("%v", *b)*/
	type test struct {
		data int
	}

	a := test{
		1,
	}

	b := &a

	c := b
	fmt.Printf("%v \n", a)
	c.data = 3

	d := test{
		2,
	}
	c = &d
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", *b)
	fmt.Printf("%v \n", b.data)
	fmt.Printf("%v \n", c)
}
