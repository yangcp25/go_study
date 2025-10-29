package main

import (
	"fmt"
	"sync"
)

func main() {
	// 文件url 并发抓取
	urls := []string{"www.baidu.com", "www.baidu.com/2", "www.baidu.com/3", "www.baidu.com", "www.baidu.com/2", "www.baidu.com/3"}

	wg := sync.WaitGroup{}
	wg.Add(len(urls))
	res := make([]string, len(urls))
	for _, url := range urls {
		go func() {
			defer wg.Done()
			res = append(res, url)
		}()
	}
	wg.Wait()
	fmt.Println(res)
	//sync.cond

	ch := make(chan string, 3)

	notify := make(chan struct{})
	go func() {
		defer close(ch)
		for i := 0; i < len(res); i++ {
			go func() {
				ch <- res[i]
			}()
		}
	}()
	go func() {
		defer func() {
			notify <- struct{}{}
		}()
		for c := range ch {
			fmt.Println(c)
		}
	}()

	<-notify
	//select {}
}

// 生产者读取文件
// 消费订阅url 进行打印

type Producer struct {
	cond sync.Cond
}
