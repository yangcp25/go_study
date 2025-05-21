package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {

	content, err := os.ReadFile("/Users/ycp/work/code/own/go_study/review/regexp/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	list := strings.Split(string(content), "\n")
	rege2 := regexp.MustCompile(`(^2[0-4][0-9]|^25[0-5]|^1[0-9]?[0-9]?|^[1-9][0-9]|^[0-9])\.(2[0-4][0-9]|25[0-5]|1[0-9]?[0-9]?|[1-9][0-9]|[0-9])\.(2[0-4][0-9]|25[0-5]|1[0-9]?[0-9]?|[1-9][0-9]|[0-9])\.(2[0-4][0-9]$|25[0-5]$|1[0-9]?[0-9]?$|[1-9][0-9]|[0-9]$)`)

	rege3 := regexp.MustCompile(`^1+0+$`)

	for _, line := range list {
		ips := strings.Split(line, "~")
		res2 := rege2.FindAllStringSubmatch(ips[0], -1)
		res3 := rege3.FindAllStringSubmatch(ips[0], -1)
		fmt.Println(ips[0], len(res2) != 0, ips[1], len(res3) == 0)

	}

	return
	rege := regexp.MustCompile(`^(2[0-4][0-9])$|^(25[0-5])$`)
	res := rege.FindAllStringSubmatch("256", -1)

	fmt.Println(res)

	res2 := rege2.FindAllStringSubmatch("196.121.115.160", -1)

	fmt.Println(res2)

	//fmt.Println(res3)
}
