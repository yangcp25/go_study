package main

import (
	"fmt"
	"time"
)

func main() {
	// 协程
	//test1()
	// 使用channel 传递消息
	//test2()

	// 通过通道传递结果 x
	test3()
}

func test3() {
	start := time.Now()

	var ch chan int
	for i := 0; i < 1; i++ {
		go add3(i, ch)
	}

	end := time.Now()

	timeLong := end.Sub(start).Seconds()

	fmt.Println("\n", timeLong)
	res := <-ch
	fmt.Println("res:\n", res)
}

func add3(i int, ch chan int) {
	var count int
	temp := <-ch
	count += temp
	ch <- count
}

func test2() {
	start := time.Now()

	var ch [10]chan int
	for i := 0; i < 10; i++ {
		ch[i] = make(chan int)
		go add2(i, ch[i])
	}

	for i := 0; i < 10; i++ {
		<-ch[i]
	}

	end := time.Now()

	timeLong := end.Sub(start).Seconds()

	fmt.Println("\n", timeLong)
}

func add2(i int, ch chan int) {
	c := i + 1
	fmt.Printf("a+1 = %v \n", c)
	ch <- c
}

func test1() {
	start := time.Now()

	for i := 0; i < 10; i++ {
		go add(i)
	}

	end := time.Now()

	timeLong := end.Sub(start).Seconds()

	fmt.Println("\n", timeLong)
}
func add(a int) {
	//a + 1
	//c := a + 1
	//fmt.Printf("a+1 = %v \n", c)
}
