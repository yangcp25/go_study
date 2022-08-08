package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	// 定义锁，保证并发安全
	Lock sync.Mutex
	// 管道 存放资源
	Sources chan io.Closer
	// 函数 用于创建资源
	Fn func() (io.Closer, error)
	// 标识符 表示是否关闭的状态
	IsClosed bool
}

var errorPoolClosed = errors.New("对象池已关闭！")

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size < 0 {
		return nil, errors.New("容量必须大于0")
	}
	return &Pool{
		Sources: make(chan io.Closer, size),
		Fn:      fn,
	}, nil
}

// Acquire 获取资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.Sources:
		if !ok {
			return nil, errorPoolClosed
		}
		log.Println("Acquire:", "---------通道资源")
		//os.Exit(2)
		return r, nil
	default:
		log.Println("Acquire:", "新增资源")
		return p.Fn()
	}
}

// Release 释放资源（回收资源）
func (p *Pool) Release(r io.Closer) {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	if p.IsClosed {
		err := r.Close()
		if err != nil {
			return
		}
	}

	// 回收资源
	select {
	case p.Sources <- r:
		log.Println("Release", "回收到队列里去")
	default:
		log.Println("Release", "资源被关闭")
	}
}

// 关闭资源池

func (p *Pool) Close() {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	// 已被关闭，直接返回
	if p.IsClosed {
		return
	}

	//
	p.IsClosed = true
	// 关闭通道
	close(p.Sources)
	for r := range p.Sources {
		err := r.Close()
		if err != nil {
			return
		}
	}
}
