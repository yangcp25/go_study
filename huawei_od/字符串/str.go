package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "password__a12345678_timeout_100"
	str2 := "aaa__password_\"a12_45678\"timeout__100_\"\"_"

	test1 := removePassword(str1, 1)
	test2 := removePassword(str2, 2)

	fmt.Println(test1, test2)
}

func removePassword(str string, k int) string {
	n := len(str)

	stack := make([]string, 0)
	flag := false
	ans := make([]string, 0)
	for i := 0; i < n; i++ {
		tempStr := string(str[i])
		if tempStr != "_" && tempStr != "\"" {
			stack = append(stack, tempStr)
		} else {
			if tempStr == "_" && !flag {
				if len(stack) > 0 {
					if len(ans) == k {
						ans = append(ans, "******")
					} else {

						ans = append(ans, strings.Trim(strings.Join(stack, ""), "_"))
					}
				}
				stack = stack[:0]
			} else {
				if !flag && tempStr == "\"" {
					stack = append(stack, tempStr)
					flag = true
				} else {
					if tempStr == "\"" {
						if len(ans) == k {
							ans = append(ans, "******")
						} else {
							stack = append(stack, tempStr)
							if len(stack) == 2 {
								ans = append(ans, "\"\"")
							} else {
								if len(stack) > 0 {
									ans = append(ans, strings.Trim(strings.Join(stack, ""), "_"))
								}
							}
						}
						stack = stack[:0]
						flag = false
					} else {
						stack = append(stack, tempStr)
					}
				}
			}
		}
	}

	if len(stack) > 0 {
		if len(ans) == k {
			ans = append(ans, "******")
		} else {
			ans = append(ans, strings.Trim(strings.Join(stack, ""), "_"))
		}
	}
	if k > len(ans) {
		return "ERROR"
	}

	return strings.Trim(strings.Join(ans, "_"), "_")
}
