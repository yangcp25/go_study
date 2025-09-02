package main

import (
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	"server"
	"server/server/api"
	"time"
)

func main() {
	// 在后台启动 gRPC 服务
	go server.StartService()
	// 等待服务启动
	time.Sleep(1000 * time.Millisecond)

	// 执行 XA 事务
	xa()

	select {}
}

const (
	// DefaultHTTPServer default url for http server. used by test and examples
	DefaultHTTPServer = "http://localhost:36789/api/dtmsvr"
	// DefaultJrpcServer default url for http json-rpc server. used by test and examples
	DefaultJrpcServer = "http://localhost:36789/api/json-rpc"
	// DefaultGrpcServer default url for grpc server. used by test and examples
	DefaultGrpcServer = "127.0.0.1:36790"
	busiServerGrpc1   = "127.0.0.1:8083/server.GrpcServer/Test"
	busiServerGrpc2   = "127.0.0.1:8085/server.GrpcServer2/Test2"
)

func xa() {
	gid := dtmgrpc.MustGenGid(DefaultGrpcServer)
	fmt.Println("XA GID:", gid)

	// 创建 XA 全局事务
	err := dtmgrpc.XaGlobalTransaction(DefaultGrpcServer, gid, func(xa *dtmgrpc.XaGrpc) error {
		req1 := &api.ReqMsg{}
		req2 := &api.ReqMsg2{}
		reply1 := &api.RpMsg{}  // 假设你的响应消息是 RpMsg
		reply2 := &api.RpMsg2{} // 假设你的响应消息是 RpMsg2

		//method1 := "/server.GrpcServer/Test"
		//method2 := "/server.GrpcServer2/Test2"

		// 正确的调用顺序：请求消息，地址，方法，响应消息
		err := xa.CallBranch(req1, busiServerGrpc1, reply1)
		if err != nil {
			return err
		}

		err = xa.CallBranch(req2, busiServerGrpc2, reply2)
		return err
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("XA transaction finished successfully, GID:", gid)
}
