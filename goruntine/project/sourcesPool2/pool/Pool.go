package pool

import (
	"sync"
)

type Worker interface {
	Task()
}
type Pool struct {
	// 任务管道
	worker chan Worker
	// 协程控制
	wg sync.WaitGroup
}

func New(maxGoroutines int) *Pool {
	p := Pool{
		worker: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)
	// 放入协程循环里面

	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.worker {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}
