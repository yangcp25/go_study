package main

import (
	"fmt"
	"math"
)

// 指针 和变量
func main() {
	//test1()
	//t := test2()
	//fmt.Printf("%v\n", &t.data)
	test4()
}

func test4() {
	fmt.Println(math.Sqrt(2))
}

// 如果返回的不是指针类型 那就是按值传递 ，否则是引用传递
func test2() *node {
	//
	t := test3()
	fmt.Printf("%v\n", &t.data)
	return t
}

type node struct {
	data int
}

func test3() *node {
	return &node{1}
}

func test1() {
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
	//
}
