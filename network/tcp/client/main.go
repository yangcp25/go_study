package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	clientConn, err := net.Dial("tcp", "127.0.0.1:8900")
	if err != nil {
		panic(err)
	}
	defer clientConn.Close()

	clientConn.Write([]byte("hello world"))
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := clientConn.Read(buf)
			if err != nil {
				fmt.Errorf("%+v", err)
				break
			}
			fmt.Println("client receive server reply ", string(buf[:n]))
		}
	}()
	for {
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		clientConn.Write([]byte(input.Text()))
	}
}
