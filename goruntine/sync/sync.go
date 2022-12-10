package main

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"
)

func main() {
	// 使用cond 设置发送通知
	//test()
	test2()
}

func test2() {
	buffer := newDataBuff()
	for i := 1; i < 3; i++ {
		go buffer.read(i)
	}

	for i := 0; i < 10; i++ {
		go func(val int) {
			msg := fmt.Sprintf("%d", val)
			buffer.put([]byte(msg))
		}(i)
		time.Sleep(10 * time.Millisecond)
	}
}

func test() {
	buffer := newDataBuff()
	go buffer.read(1)

	go func(val int) {
		msg := fmt.Sprintf("我来了哦!!%d", val)
		buffer.put([]byte(msg))
	}(2)
	time.Sleep(100 * time.Millisecond)
}

func (receiver dataBuff) put(val []byte) (int, error) {
	// 打开写锁
	receiver.Lock.Lock()
	defer receiver.Lock.Unlock()
	n, err := receiver.Buff.Write(val)
	receiver.Cond.Broadcast()
	return n, err
}
func (receiver dataBuff) read(i int) {
	// 打开读锁
	receiver.Lock.RLock()
	defer receiver.Lock.RUnlock()
	var data []byte
	var d byte
	var err error
	for {
		d, err = receiver.Buff.ReadByte()
		if err != nil {
			// 达到了末尾
			if err == io.EOF {
				// 打印
				if string(data) != "" {
					fmt.Printf("%d:%s\n", i, data)
				}
				// 阻塞协程等待下次通知
				receiver.Cond.Wait() // wait 操作会打开读锁，在初始化的时候会传入读锁，所以写入器可以获取写锁
				data = data[:0]
				continue
			}
		}
		data = append(data, d)
	}
}

type dataBuff struct {
	Buff *bytes.Buffer
	Lock *sync.RWMutex
	Cond *sync.Cond
}

func newDataBuff() *dataBuff {
	buffer := make([]byte, 0)
	dataBuffObj := &dataBuff{
		bytes.NewBuffer(buffer),
		new(sync.RWMutex),
		nil,
	}
	dataBuffObj.Cond = sync.NewCond(dataBuffObj.Lock.RLocker())

	return dataBuffObj
}
