package main

import (
	"fmt"
	"os"
	"syscall"
)

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "错误：", err)
		os.Exit(1)
	}
}

func main() {
	// 1. 创建 socket
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	check(err)
	defer syscall.Close(fd)

	// 2. 绑定到本地地址 0.0.0.0:8888
	sa := &syscall.SockaddrInet4{Port: 8888}
	// IP 全 0 表示 INADDR_ANY
	check(syscall.Bind(fd, sa))

	// 3. 开始监听（backlog 128）
	check(syscall.Listen(fd, 128))
	fmt.Println("raw socket 服务器已启动，端口 8888")

	// 4. 接受连接
	nfd, rsa, err := syscall.Accept(fd)
	check(err)
	fmt.Printf("接收到客户端：%v\n", rsa)
	defer syscall.Close(nfd)

	// 5. 全双工：一个 goroutine 从 stdin 发出去，一个 goroutine 从 socket 读进来
	go func() {
		buf := make([]byte, 1024)
		for {
			// 从标准输入读
			n, err := syscall.Read(syscall.Stdin, buf)
			if err != nil {
				fmt.Fprintln(os.Stderr, "stdin 读取错误：", err)
				return
			}
			if n > 0 {
				_, err = syscall.Write(nfd, buf[:n])
				if err != nil {
					fmt.Fprintln(os.Stderr, "向客户端写入错误：", err)
					return
				}
			}
		}
	}()

	// 主 goroutine 负责从 socket 读并写到 stdout
	buf := make([]byte, 1024)
	for {
		n, err := syscall.Read(nfd, buf)
		if err != nil {
			fmt.Fprintln(os.Stderr, "从客户端读取错误：", err)
			return
		}
		if n > 0 {
			// 打印到标准输出
			_, err = syscall.Write(syscall.Stdout, buf[:n])
			if err != nil {
				fmt.Fprintln(os.Stderr, "写到 stdout 错误：", err)
				return
			}
		}
	}
}
