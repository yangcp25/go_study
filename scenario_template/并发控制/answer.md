```go
// 获取url爬虫 并发控制

f, err := os.Open("/Users/ycp/work/code/own/go_study/scenario_template/pratice/urls.txt")
if err != nil {
fmt.Println(err)
return
}
defer f.Close()
scanner := bufio.NewScanner(f)
urls := make([]string, 0)
for scanner.Scan() {
urls = append(urls, scanner.Text())
}

ch := make(chan struct{}, 2)

wg := sync.WaitGroup{}

for _, url := range urls {
wg.Add(1)
ch <- struct{}{}
go func(url string) {
defer wg.Done()
defer func() {
<-ch
}()
content, err := http.Get(url)
if err != nil {
fmt.Println(err)
return
}
defer content.Body.Close()
all, err := io.ReadAll(content.Body)
if err != nil {
fmt.Println(err)
return
}
fmt.Println(string(all))

}(url)
}

wg.Wait()

```

```go
// 控制并发数量的模板
	ch := make(chan struct{}, 3)
	wg := sync.WaitGroup{}
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		ch <- struct{}{}
		go func(i int) {
			defer wg.Done()
			defer func() {
				<-ch
			}()
			fmt.Println(i)
			//time.Sleep(time.Second)
		}(i)
	}

	wg.Wait()

	// 超时控制


```

```go
// 超时控制
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // 一定记得调用，释放资源

	result := make(chan string, 1)

	go func() {
		// 模拟一个慢任务（比如数据库查询）
		time.Sleep(3 * time.Second)
		result <- "task finished"
	}()

	select {
	case res := <-result:
		fmt.Println("✅ got result:", res)
	case <-ctx.Done():
		fmt.Println("⏰ timeout:", ctx.Err())
	}
}

```