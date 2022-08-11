package main

import (
	"fmt"
	"net/rpc"
	"rpc/utils"
)

func main() {
	syncCall()
	asyncCall()
}

// 异步调用
func asyncCall() {
	var serverAddr = "localhost"
	client, err := rpc.DialHTTP("tcp", serverAddr+":8080")
	if err != nil {
		fmt.Printf("建立远程http失败:%v", err)
	}
	arg := &utils.Args{A: 10, B: 10}

	var reply int
	asyncCallRes := client.Go("MathService.Multiply", arg, &reply, nil)

	select {
	case <-asyncCallRes.Done:
		fmt.Print("调用完成!!!!")
		return
	}
}

// 同步调用
func syncCall() {
	var serverAddr = "localhost"
	client, err := rpc.DialHTTP("tcp", serverAddr+":8080")
	if err != nil {
		fmt.Printf("建立远程http失败:%v", err)
	}
	arg := &utils.Args{A: 10, B: 10}

	var reply int
	err = client.Call("MathService.Multiply", arg, &reply)
	if err != nil {
		fmt.Printf("远程调用失败:%v", err)
	}

	fmt.Printf("%d*%d=%d", arg.A, arg.B, reply)
}
