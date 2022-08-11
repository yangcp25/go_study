package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"rpc/utils"
)

type MathService struct {
}

func (receiver MathService) Multiply(args *utils.Args, reply *int) error {
	*reply = args.A * args.B

	return nil
}

func (receiver MathService) Add(args *utils.Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}
func main() {
	math := new(MathService)
	rpc.Register(math)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("启动服务失败！", err)
	}

	err = http.Serve(listener, nil)

	if err != nil {
		log.Fatal("启动http服务失败!", err)
	}
}
