package main

import "sync"

func main() {

	//实现一个单例模式
	//1) 协程安全
	//2) 你认为性能最好的实现方式

}

type Client struct {
	conf string
	once sync.Once
}

func NewClient(conf string) *Client {
	snyc := sync.Once{}
	var f *Client
	snyc.Do(func() {
		f = &Client{
			conf: "",
		}
	})
	return f
}
