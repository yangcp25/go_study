package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// socket connect
	conn, err := net.Dial("tcp", ":8099")
	if err != nil {
		panic(err)
	}
	go func() {
		if _, err := io.Copy(conn, os.Stdin); err != nil {
			fmt.Errorf(" to server :%w", err)
		}
	}()
	if _, err := io.Copy(os.Stdout, conn); err != nil {
		fmt.Errorf("from server :%w", err)
	}
}
