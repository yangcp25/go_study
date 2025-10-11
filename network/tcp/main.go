package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Listen("tcp", "127.0.0.1:8900")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for {
		acceptCome, err := conn.Accept()
		if err != nil {
			panic(err)
			return
		}
		go handleConnect(acceptCome)
	}
}

func handleConnect(come net.Conn) {
	addr := come.RemoteAddr().String()
	for {
		buf := make([]byte, 1024)
		n, err := come.Read(buf)
		if err != nil {
			break
		}
		fmt.Println("server review remote addr:", addr, string(buf[:n]))

		// 发送给客户端
		_, err = come.Write(buf[:n])
		if err != nil {
			break
		}
	}
}
