package main

import (
	"fmt"
)

func main() {
	//parent := context.Background()
	//ctx, _ := context.WithCancel(parent)
	//ctx, _ := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	//ctx, _ := context.WithTimeout(parent, time.Now().Add(5*time.Second))
	//defer func() {
	//	print(111)
	//}()
	//defer func() {
	//	print(2222)
	//}()
	//
	//ch := make(chan int)
	////close(ch)
	//fmt.Println(ch)
	//fmt.Println(ch)
	//fmt.Println(ch)
	//fmt.Println(ch)
	//go func() {
	//	test := <-ch
	//
	//	fmt.Println(test)
	//}()
	//ch <- 3

	//var test1 []int
	//test1 = make([]int, 1)
	//test1 = append(test1, 1)
	//test1 = append(test1, 2)
	//test1 = append(test1, 3)
	//test1 = append(test1, 5)
	//test1 = append(test1, 6)

	//fmt.Println(test1[0])

	//var test2 map[int]int
	//test2 = make(map[int]int)
	//test2[1] = 2
	//test2[2] = 3
	//fmt.Println(test2)

	//sync.Once{}
	//sync.Cond{}
	var test3 chan int

	test3 = make(chan int, 3)
	go func() {
		fmt.Println(<-test3)
		fmt.Println(<-test3)
	}()
	test3 <- 3
	close(test3)
	if v, ok := <-test3; ok {
		fmt.Println(v)
	}

	for {

	}

}
