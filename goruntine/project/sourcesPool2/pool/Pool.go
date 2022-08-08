package pool

import (
	"sync"
)

type Worker interface {
	Task()
}
type Pool struct {
	// 任务管道
	Worker chan Worker
	// 协程控制
	wg sync.WaitGroup
}

func New(maxGoroutines int) *Pool {
	p := Pool{
		Worker: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)
	// 放入协程循环里面

	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.Worker {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

func (p *Pool) Run(w Worker) {
	p.Worker <- w
}

func (p *Pool) Close() {
	close(p.Worker)
	p.wg.Wait()
}
