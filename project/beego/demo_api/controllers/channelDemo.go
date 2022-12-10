package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"math/rand"
	"time"
)

type ChannelDemoController struct {
	beego.Controller
}

func (c *ChannelDemoController) Test() {
	/*chanTest := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go add(i, chanTest)
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", <-chanTest)
	}

	var ch [10]chan int

	for i := 0; i < 10; i++ {
		go add2(i, ch[i])
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", ch[i])
	}*/

	//
	chanTest := make(chan int, 10)
	go add3(chanTest)

	for i := range chanTest {
		fmt.Printf("%d\n", i)
	}

}
func (c *ChannelDemoController) TestSelect() {

	ch := make(chan int)

	ch <- 2

	select {
	case num := <-ch:
		fmt.Printf("随机数：%d\n", num)
	}

	fmt.Println("xx")
}

func test2(ch chan int) {

}

func test1(ch chan int) {

}

func add3(test chan int) {
	for i := 0; i < 100; i++ {
		test <- i
	}
	close(test)
}

func add(i int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
	c <- i
}
func add2(i int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
	c <- i
}
