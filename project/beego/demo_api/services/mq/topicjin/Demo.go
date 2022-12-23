package main

import (
	"demo_api/services/mq"
	"fmt"
)

func main() {
	mq.ConsumerEx("ycp_test.demo.topic", "topic", "*.test.*", callback)
}

func callback(msg string) {
	fmt.Printf("%s\n", msg)
}
