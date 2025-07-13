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

	// 2. 连接到服务器 127.0.0.1:8888
	sa := &syscall.SockaddrInet4{Port: 8888}
	copy(sa.Addr[:], []byte{127, 0, 0, 1})
	check(syscall.Connect(fd, sa))
	fmt.Println("raw socket 已连接到 127.0.0.1:8888")

	// 3. 全双工：一个 goroutine 从 stdin 发出去，一个从 socket 读进来
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := syscall.Read(syscall.Stdin, buf)
			if err != nil {
				fmt.Fprintln(os.Stderr, "stdin 读取错误：", err)
				return
			}
			if n > 0 {
				_, err = syscall.Write(fd, buf[:n])
				if err != nil {
					fmt.Fprintln(os.Stderr, "写到服务器错误：", err)
					return
				}
			}
		}
	}()

	// 主 goroutine 负责从 socket 读并写到 stdout
	buf := make([]byte, 1024)
	for {
		n, err := syscall.Read(fd, buf)
		if err != nil {
			fmt.Fprintln(os.Stderr, "从服务器读取错误：", err)
			return
		}
		if n > 0 {
			_, err = syscall.Write(syscall.Stdout, buf[:n])
			if err != nil {
				fmt.Fprintln(os.Stderr, "写到 stdout 错误：", err)
				return
			}
		}
	}
}
