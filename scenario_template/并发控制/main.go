package main

import (
	"fmt"
	"sync"
)

func main() {

}

// 实现生产者消费者

func ProductConsumer() {
	urls := []string{"a", "b", "c", "d", "e"}
	jobs := make(chan string, len(urls))
	results := make(chan string, len(urls))

	// 启动3个消费者
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for url := range jobs {
				results <- fmt.Sprintf("worker %d fetched %s", id, url)
			}
		}(i)
	}

	// 生产者
	go func() {
		for _, url := range urls {
			jobs <- url
		}
		close(jobs)
	}()

	// 等消费者处理完后关闭结果通道
	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println(r)
	}

}
