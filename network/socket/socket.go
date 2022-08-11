package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	socket()
}
func socket() {
	//dialFunc()
	TcpFunc()
}

func dialFunc() {
	net.Dial("tcp", "127.0.0.1:8082")
	net.Dial("udp", "127.0.0.1:8082")
	// icmp
	net.Dial("ip4:icmp", "www.baidu.com")
	net.Dial("ip4:icmp", "127.0.0.1")
}

func TcpFunc() {
	if len(os.Args) != 2 {
		fmt.Print("请输入正确的参数")
	}
	fmt.Printf("%v\n", os.Args)
	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	checkErr(err)
	// 发送数据
	_, err = conn.Write([]byte("你好的很"))
	checkErr(err)

	result, err := readConn(conn)
	checkErr(err)
	fmt.Printf("%v", string(result))
}

func readConn(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var getBytesBuff [512]byte
	for {
		n, err := conn.Read(getBytesBuff[0:])
		result.Write(getBytesBuff[0:n])

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

	}

	return result.Bytes(), nil
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}
