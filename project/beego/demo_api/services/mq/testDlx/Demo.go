package main

import (
	"demo_api/services/mq"
	"fmt"
)

func main() {
	mq.ConsumerDlx("ycp_dlx.test.a", "ycp_dlx_a", "ycp_dlx.test.b", "ycp_dlx_b", 10000, callback)
}

func callback(msg string) {
	fmt.Printf("%s\n", msg)
}
