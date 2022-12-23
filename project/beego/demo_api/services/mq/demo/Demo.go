package main

import (
	"demo_api/services/mq"
	"fmt"
)

func main() {
	mq.Consumer("", "ycp_mq_test", callback)
}

func callback(msg string) {
	fmt.Printf("%s\n", msg)
}
