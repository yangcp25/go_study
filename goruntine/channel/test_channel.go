package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 避免关闭已经关闭的通道
	test8()
	// 超时机制实现
	//test7()
	//
	//test6()
	// 协程
	//test1()
	// 使用channel 传递消息
	//test2()

	// 非缓冲通道：如果读取操作时，通道没有数据，会堵塞；如果写入操作时，通道已有数据，会阻塞，只到通道的数据被读取，才会继续写入
	// 通道传输的值是值的副本；不是引用传递
	// 缓冲通道：
	//test3()
	// 通过通道传递结果 x 这是并行了
	//test4()
	// 单向通道：是指程序上规定通信的方向不是语言层面上的限制；语言上始终支持发送和接收
	//test5()
	//
}

func test8() {
	ch := make(chan int, 2)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		fmt.Println("发送方关了哟！")
		close(ch)
	}()

	for {
		num, ok := <-ch
		if !ok {
			fmt.Println("你关了我也关了嘛那")
			break
		}
		fmt.Println("接收到了", num)
	}
}

func test7() {
	ch := make(chan int)

	closeChan := make(chan bool)

	go func() {
		time.Sleep(1 * time.Second)
		closeChan <- true
	}()

	select {
	case data, ok := <-ch:
		if !ok {
			break
		}
		fmt.Println("数据来啦！、n\n", data)
	case <-closeChan:
		fmt.Println("超时了！")
	}
}

func test6() {
	// 使用select 判断通道的值
	chs := [3]chan int{
		make(chan int, 3),
		make(chan int, 3),
		make(chan int, 3),
	}
	index1 := rand.Intn(3)
	chs[index1] <- rand.Int()

	index2 := rand.Intn(3)
	chs[index2] <- rand.Int()

	index3 := rand.Intn(3)
	chs[index3] <- rand.Int()

	for i := 0; i < 3; i++ {
		select {
		case nums, ok := <-chs[0]:
			fmt.Println("选到第一个了before：", nums)
			if !ok {
				break
			}
			fmt.Println("选到第一个了：", nums)
		case nums, ok := <-chs[1]:
			if !ok {
				break
			}
			fmt.Println("选到第2个了：", nums)
		case nums, ok := <-chs[2]:
			if !ok {
				break
			}
			fmt.Println("选到第3个了：", nums)
		default:
			fmt.Println("个都没选到：")
		}
	}
}

func test5() {
	start := time.Now()

	ch := make(chan int, 20)
	go add5(ch)
	add5Foreach(ch)
	end := time.Now()
	timeLong := end.Sub(start).Seconds()
	fmt.Println("\n", timeLong)
}

func add5Foreach(ch <-chan int) {
	for i := range ch {
		fmt.Println("接收到的数据:", i)
	}
}
func add5(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func test4() {

}

func test3() {
	start := time.Now()

	ch := make(chan int, 20)
	go add3(ch)
	for i := range ch {
		fmt.Println("接收到的数据:", i)
	}
	end := time.Now()
	timeLong := end.Sub(start).Seconds()
	fmt.Println("\n", timeLong)
}

func add3(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
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
