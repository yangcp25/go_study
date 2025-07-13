package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// socket bind listen
	con, err := net.Listen("tcp", ":8099")
	if err != nil {
		panic(err)
	}
	conn, err := con.Accept()
	go func() {
		if _, err := io.Copy(conn, os.Stdin); err != nil {
			fmt.Errorf("server to client :%w", err)
		}
	}()
	if _, err := io.Copy(os.Stdout, conn); err != nil {
		fmt.Errorf("from client :%w", err)
	}
}
