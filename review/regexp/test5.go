package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 编译正则
	re := regexp.MustCompile(`(^2[0-4][0-9]|^25[0-5]|^1[0-9]?[0-9]?|^[1-9][0-9]|^[0-9])\.(2[0-4][0-9]|25[0-5]|1[0-9]?[0-9]?|[1-9][0-9]|[0-9])\.(2[0-4][0-9]|25[0-5]|1[0-9]?[0-9]?|[1-9][0-9]|[0-9])\.(2[0-4][0-9]$|25[0-5]$|1[0-9]?[0-9]?$|[1-9][0-9]|[0-9]$)`)

	tests := []string{
		"192.168.0.1",
		"255.255.255.255",
		"256.100.0.1", // 无效
		"1.2.3.4",
	}

	for _, ip := range tests {
		if matches := re.FindStringSubmatch(ip); matches != nil {
			// matches[0] 是整个匹配，matches[1]~matches[4] 分别是四个捕获组
			fmt.Printf("有效 IP: %s → 段1=%s, 段2=%s, 段3=%s, 段4=%s\n",
				ip, matches[1], matches[2], matches[3], matches[4])
		} else {
			fmt.Printf("无效 IP: %s\n", ip)
		}
	}
}
