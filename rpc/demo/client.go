package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Goods struct {
	Id   int32
	Name string
}

// 编写一个rpc服务并注册
func main() {
	handler, _ := rpc.DialHTTP("tcp", ":8084")

	arg := Goods{
		Id:   2,
		Name: "wangi",
	}
	var res Goods
	err := handler.Call("Goods.AddGoods", arg, &res)
	if err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Println(res)

	arg = Goods{
		Id:   3,
		Name: "wangi3",
	}
	err = handler.Call("Goods.GetGoods", arg, &res)

	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println(res.Name)
}
