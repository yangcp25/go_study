package main

import (
	"fmt"
	"regexp"
)

func main() {
	ip := "192.168.0.1"
	pattern := `^((?:25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9][0-9]|0))\.((?:25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9][0-9]|0))\.((?:25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9][0-9]|0))\.((?:25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9][0-9]|0))$`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(ip)
	fmt.Println(matches)
	if matches != nil {
		fmt.Printf("IP地址: %s\n", matches[0])
		fmt.Printf("段1: %s\n", matches[1])
		fmt.Printf("段2: %s\n", matches[2])
		fmt.Printf("段3: %s\n", matches[3])
		fmt.Printf("段4: %s\n", matches[4])
	} else {
		fmt.Println("无效的IP地址")
	}
}
