package server

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"server/server/api"
	"sync"
)

type gServer2 struct {
	api.UnimplementedGrpcServer2Server
}

func (receiver *gServer2) Test2(ctx context.Context, rq *api.ReqMsg2) (rp *api.RpMsg2, err error) {
	fmt.Println("Test2:", rq)
	return &api.RpMsg2{}, errors.New("error")
}

type gServer struct {
	api.UnimplementedGrpcServerServer
}

func (receiver *gServer) Test(ctx context.Context, rq *api.ReqMsg) (rp *api.RpMsg, err error) {
	fmt.Println("Test:", rq)
	return &api.RpMsg{}, nil
}

// 模拟数据库
var accountA = struct {
	balance float64
	frozen  float64
}{
	balance: 100.0,
	frozen:  0.0,
}

func StartService() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		netS, err := net.Listen("tcp", ":8083")
		if err != nil {
			panic(err)
		}
		defer netS.Close()
		g := grpc.NewServer()
		api.RegisterGrpcServerServer(g, &gServer{})
		fmt.Println("success grpc1 8083")
		err = g.Serve(netS)
		if err != nil {
			return
		}
		fmt.Println("success grpc1 finished")
	}()

	go func() {
		defer wg.Done()
		netS, err := net.Listen("tcp", ":8085")
		if err != nil {
			panic(err)
		}
		defer netS.Close()
		g := grpc.NewServer()
		api.RegisterGrpcServer2Server(g, &gServer2{})
		fmt.Println("success grpc2 8085")
		err = g.Serve(netS)
		if err != nil {
			return
		}
		fmt.Println("success grpc2 finished")
	}()

	select {}

}
