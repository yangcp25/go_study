package main

import (
	"fmt"
	"strings"
)

func main() {
	//目标版本：1.0.1.59
	//
	//当前版本：1.0.0.90
	curVersion := "1.0.0.90"
	version := "1.0.1.59"

	res := isUpdate(curVersion, version)

	fmt.Println(res)
}

func isUpdate(curVersion, version string) bool {
	curVersionSlice := strings.Split(curVersion, ".")
	versionSlice := strings.Split(version, ".")

	for i := 0; i < len(curVersion); i++ {
		if versionSlice[i] > curVersionSlice[i] {
			return true
		}
	}

	return false
}
