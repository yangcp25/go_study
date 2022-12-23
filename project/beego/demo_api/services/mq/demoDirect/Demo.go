package main

import (
	"demo_api/services/mq"
	"fmt"
)

func main() {
	mq.ConsumerEx("ycp_test.demo.direct", "direct", "one", callback)
}

func callback(msg string) {
	fmt.Printf("%s\n", msg)
}
