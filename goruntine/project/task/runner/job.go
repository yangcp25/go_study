package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// 需要一个结构体存放任务
// (1)接收是否中断的系统信号 (2)系统中断的错误类型 （3）超时的错误 (4) 任务列表 数组函数
var ErrorTimeOut = errors.New("超时错误")
var InterruptOut = errors.New("系统中断")

type task struct {
	interrupt chan os.Signal
	complete  chan error
	// 单向通道 右边 发送 左边接收
	timeOut <-chan time.Time
	list    []func(int)
}

func New(timeOut time.Duration) *task {
	return &task{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeOut:   time.After(timeOut),
	}
}

func (t *task) Add(funcItem ...func(int)) {
	t.list = append(t.list, funcItem...)
}

func (t *task) Start() error {
	// 系统中断 监听
	signal.Notify(t.interrupt, os.Interrupt)

	go func() {
		t.complete <- t.run()
	}()

	select {
	case err := <-t.complete:
		return err
	case <-t.timeOut:
		return ErrorTimeOut
	}
}

func (t *task) run() error {
	for i, itemFunc := range t.list {
		if t.getInterrupt() {
			return InterruptOut
		}
		itemFunc(i)
	}
	return nil
}

func (t *task) getInterrupt() bool {
	select {
	case <-t.interrupt:
		signal.Stop(t.interrupt)
		return true
	default:
		return false
	}
}
