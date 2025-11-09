package main

import (
	"errors"
	"sync"
	"sync/atomic"
)

func main() {
	//
	//f, err := os.Open("/Users/ycp/work/code/own/go_study/scenario_template/pratice/urls.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer f.Close()
	//var buf bytes.Buffer
	////buf.ReadFrom(f)
	////fmt.Println(buf.String())
	//
	//io.ReadAll(&buf)

	//f2, err := os.ReadFile("urls.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//fmt.Println(string(f2))

	// 按行读取

	//scanner := bufio.NewScanner(f)
	//
	//var urls []string
	//for scanner.Scan() {
	//	urls = append(urls, scanner.Text())
	//}
	//fmt.Println(urls)

	// 分块读取

	//chunk := make([]byte, 100)
	//
	//count := 0
	//for {
	//	n, err := f.Read(chunk)
	//	if err != nil {
	//		return
	//	}
	//	if n > 0 {
	//		fmt.Println("received data :", string(chunk[:n]))
	//		count += n
	//	}
	//
	//	if err == io.EOF {
	//		break
	//	}
	//
	//}
	//
	//fmt.Println(string(chunk))

	// 交替打印

	//chan1 := make(chan struct{})
	//chan2 := make(chan struct{})
	//
	//wg := sync.WaitGroup{}
	//
	//wg.Add(2)
	//
	//go func() {
	//	defer wg.Done()
	//	for i := 1; i < 100; i += 2 {
	//		<-chan1
	//		fmt.Printf("chan1: %d\n", i)
	//		chan2 <- struct{}{}
	//	}
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	for i := 2; i <= 100; i += 2 {
	//		<-chan2
	//		fmt.Printf("chan2: %d\n", i)
	//		if i < 100 {
	//			chan1 <- struct{}{}
	//		}
	//	}
	//}()
	//
	//chan1 <- struct{}{}
	//
	//wg.Wait()
	//
	//close(chan1)
	//close(chan2)

	//chan1 := make(chan int)
	//chan2 := make(chan int)
	//
	//wg := sync.WaitGroup{}
	//wg.Add(2)
	//go func() {
	//	defer wg.Done()
	//	for v := range chan1 {
	//		fmt.Println("chan1", v)
	//		chan2 <- v + 1
	//		if v == 99 {
	//			close(chan2)
	//		}
	//	}
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	for v := range chan2 {
	//		fmt.Println("chan2", v)
	//		if v < 99 {
	//			chan1 <- v + 1
	//		} else {
	//			close(chan1)
	//		}
	//	}
	//}()
	//chan1 <- 1
	//wg.Wait()

	//
	//f, err := os.Open("/Users/ycp/work/code/own/go_study/scenario_template/pratice/urls.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer f.Close()
	//scanner := bufio.NewScanner(f)
	//urls := make([]string, 0)
	//for scanner.Scan() {
	//	urls = append(urls, scanner.Text())
	//}
	//
	//ch := make(chan struct{}, 2)
	//
	//wg := sync.WaitGroup{}
	//
	//for _, url := range urls {
	//	wg.Add(1)
	//	ch <- struct{}{}
	//	go func(url string) {
	//		defer wg.Done()
	//		defer func() {
	//			<-ch
	//		}()
	//		content, err := http.Get(url)
	//		if err != nil {
	//			fmt.Println(err)
	//			return
	//		}
	//		defer content.Body.Close()
	//		all, err := io.ReadAll(content.Body)
	//		if err != nil {
	//			fmt.Println(err)
	//			return
	//		}
	//		fmt.Println(string(all))
	//
	//	}(url)
	//}
	//
	//wg.Wait()

	//// 控制并发数量的模板
	//ch := make(chan struct{}, 3)
	//wg := sync.WaitGroup{}
	//runtime.GOMAXPROCS(1)
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	ch <- struct{}{}
	//	go func(i int) {
	//		defer wg.Done()
	//		defer func() {
	//			<-ch
	//		}()
	//		fmt.Println(i)
	//		//time.Sleep(time.Second)
	//	}(i)
	//}
	//
	//wg.Wait()

	// 任务池

	// 生产者消费者
}

type workerPool struct {
	ch     chan worker
	wg     sync.WaitGroup
	closed atomic.Bool
}

type worker func()

func newWorkerPool(workerSize, pollSize int) *workerPool {
	poll := &workerPool{
		ch: make(chan worker, pollSize),
	}
	poll.wg.Add(workerSize)
	for i := 0; i < workerSize; i++ {
		go func() {
			defer poll.wg.Done()
			for w := range poll.ch {
				w()
			}
		}()
	}

	return poll
}

func (p *workerPool) Submit(f func()) (err error) {
	if p.closed.Load() {
		return errors.New("pool is closed")
	}
	p.ch <- f
	return
}

func (p *workerPool) Close() {
	if p.closed.Swap(true) {
		return
	}
	close(p.ch)
	p.wg.Wait()
}
