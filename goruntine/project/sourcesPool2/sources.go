package main

import (
	"log"
	"sourcesPool2/pool"
	"sync"
	"time"
)

var stringStr = []string{
	"python",
	"java",
	"php",
	"golang",
	"javascript",
}

const (
	maxGoroutines = 5
	maxResources  = 2
)

type langPrinter struct {
	Str string
}

func (l langPrinter) Task() {
	log.Println(l.Str)
	time.Sleep(time.Duration(1) * time.Second)
}

// 协程池 无缓冲版
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(maxGoroutines)

	poolObj := pool.New(maxResources)

	for i := 0; i < maxGoroutines; i++ {
		for _, str := range stringStr {
			printer := &langPrinter{Str: str}
			// 放入work
			go func() {
				poolObj.Run(printer)
				wg.Done()
			}()
		}
	}

	wg.Wait()
	poolObj.Close()
}
