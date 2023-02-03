package observer

import (
	"fmt"
	"reflect"
	"sync"
)

type IBus interface {
	Subscribe(topic string, handler interface{}) error
	Publish(topic string, args ...interface{})
}

type AsyncEventBus struct {
	handler map[string][]reflect.Value
	lock    sync.Mutex
}

func NewSyncEventBus() *AsyncEventBus {
	return &AsyncEventBus{
		handler: map[string][]reflect.Value{},
		lock:    sync.Mutex{},
	}
}

// 订阅

func (a *AsyncEventBus) Subscribe(topic string, handler interface{}) error {
	a.lock.Lock()
	defer a.lock.Unlock()

	v := reflect.ValueOf(handler)

	if v.Type().Kind() != reflect.Func {
		return fmt.Errorf("handler need func")
	}

	handlerObj, ok := a.handler[topic]

	if !ok {
		handlerObj = []reflect.Value{}
	}

	handlerObj = append(handlerObj, v)

	a.handler[topic] = handlerObj

	return nil
}

func (a *AsyncEventBus) Publish(topic string, arg ...interface{}) {
	handler, ok := a.handler[topic]

	if !ok {
		fmt.Println("未找到对应topic")
	}

	params := make([]reflect.Value, len(arg))

	for i, v := range arg {
		params[i] = reflect.ValueOf(v)
	}

	for _, value := range handler {
		go value.Call(params)
	}
}
