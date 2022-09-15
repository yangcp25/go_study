package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Goods struct {
	Id   int32
	Name string
}

func (receiver Goods) AddGoods(req Goods, res *Goods) error {
	fmt.Println(req)
	*res = Goods{
		Id:   req.Id,
		Name: req.Name,
	}
	return nil
}

func (receiver Goods) GetGoods(req Goods, res *Goods) error {
	fmt.Println(req)
	*res = Goods{
		Id:   req.Id,
		Name: req.Name,
	}
	return nil
}

// 编写一个rpc服务并注册
func main() {
	goods := new(Goods)
	rpc.RegisterName("Goods", goods)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":8084")

	if err != nil {
		log.Fatal("启动服务失败！", err)
	}

	err = http.Serve(listener, nil)

	if err != nil {
		log.Fatal("启动http服务失败!", err)
	}

}
