package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//
	//waitGroupDo()
	// 固定大小的worker pool
	//workerPool(3)

	// 爬虫限制并发数限制请求频率
	//crawlerLimit()

	// 实现生产者消费者
	//ProductConsumer()
	// 消息订阅
	//MQSubscribePublish()
	// 限流
	tokenBucketDo()
}

func tokenBucketDo() {
	tb := NewTokenBucket(5, 20)
	for i := 0; i < 20; i++ {
		if tb.Allow() {
			fmt.Println("allow", i, time.Now())
		} else {
			fmt.Println("reject", i, time.Now())
		}
	}
}

type TokenBucket struct {
	tokens chan struct{}
}

func NewTokenBucket(rant, capacity int) *TokenBucket {
	tb := TokenBucket{tokens: make(chan struct{}, capacity)}
	// 初始时先装满
	for i := 0; i < capacity; i++ {
		tb.tokens <- struct{}{}
	}
	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(rant))
		for {
			select {
			case <-ticker.C:
				tb.tokens <- struct{}{}
			default:
			}
		}
	}()
	return &tb
}

func (tb *TokenBucket) Allow() bool {
	select {
	case <-tb.tokens:
		return true
	default:
		return false
	}
}

func MQSubscribePublish() {
	broker := NewBroker()
	coumser := broker.Subscribe("t1")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range coumser {
			fmt.Println("coumser:", v)
		}
	}()

	broker.Publish("t1", "msg1")
	broker.Publish("t1", "msg2")
	broker.UnSubscribe("t1", coumser)
	wg.Wait()
}

type Broker struct {
	rw   sync.RWMutex
	subs map[string][]chan string
}

func NewBroker() *Broker {
	return &Broker{
		subs: make(map[string][]chan string),
	}
}
func (this *Broker) Subscribe(topic string) (ch chan string) {
	ch = make(chan string, 10)
	this.rw.Lock()
	this.subs[topic] = append(this.subs[topic], ch)
	this.rw.Unlock()
	return
}
func (this Broker) Publish(topic string, data string) {
	this.rw.RLock()
	t := this.subs[topic]
	this.rw.RUnlock()
	for _, ch := range t {
		select {
		case ch <- data:
		default:
		}
	}
}

func (this Broker) UnSubscribe(topic string, ch chan string) {
	this.rw.RLock()
	subs := this.subs[topic]
	this.rw.RUnlock()

	for i, v := range subs {
		if v == ch {
			subs = append(subs[:i], subs[i+1:]...)
			close(ch)
			break
		}
	}

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

// 使用场景
/**
1 网络IO（api调用） 需要控制并发数量
2 数据库需要控制连接数
3 cpu密集 设置并发数量为cpu核心数
4 内存分配频繁
*/
func workerPool(num int) {
	ch := make(chan struct{}, num)
	for i := 0; i < 10; i++ {
		ch <- struct{}{}
		fmt.Println("workerPool i:", i)
		<-ch
	}
}

func crawlerLimit() {

}
func waitGroupDo() {
	wg := &sync.WaitGroup{}

	fmt.Println("waitGroupDo start ")
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}

	wg.Wait()
}
