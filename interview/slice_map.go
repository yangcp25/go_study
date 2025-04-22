//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	//str1 := []string{"a", "b", "c"}
//	str1 := make([]string, 0)
//	str1 = append(str1, "a")
//	str1 = append(str1, "b")
//	str1 = append(str1, "c")
//	str2 := str1[1:]
//	//str2 = "new"
//	fmt.Println(str1)
//	str2 = append(str2, "z", "x", "y")
//	fmt.Println(str1)
//}

package main

import (
	"fmt"
)

type Student struct {
	Age int
}

// string 转出切片数组是否会发生内存拷贝，如果会，如何优化性能
func main() {
	//kv := map[string]Student{"menglu": {Age: 21}}
	//fmt.Println(kv)
	//kv["menglu"].Age = 22
	s := []Student{{Age: 21}}
	s[0].Age = 22
	fmt.Println(s)
}
