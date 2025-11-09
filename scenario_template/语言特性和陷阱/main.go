package main

import (
	"bytes"
	"fmt"
	"sync"
	"unsafe"
)

func main() {
	nums := []int{1, 3, 5, 7, 9}
	modifySlice(nums, 0)
	fmt.Println("nums", nums)
	// 扩容 外部会看不见修改的数据
	b := appendSlice(nums)
	fmt.Println("nums", nums)
	fmt.Println("b", b)
	modifySlice(b, 2)

	fmt.Println("nums", nums)
	fmt.Println("b", b)

	modifySlice(nums, 3)
	fmt.Println("nums", nums)
	fmt.Println("b", b)

	n := deferDO()
	fmt.Println("n", n)

	s := make([]int, 0, 5)
	for i := 0; i < 6; i++ {
		s = append(s, i)
	}
	//appendS(s)
	fmt.Println(cap(s), s)

	// 字符串拼接
	JoinStr()

	// 字符串 byte 类型转换
	TransferByte()

	// sync.map的使用
	syncMap()
}

func syncMap() {
	map1 := sync.Map{}
	map1.Store("a", 1)
	if v, ok := map1.Load("a"); ok {
		fmt.Println(v)
	}
	map1.Delete("a")

	if v, ok := map1.Load("a"); ok {
		fmt.Println(v)
	} else {
		fmt.Println("key not found")
	}
}

func TransferByte() {
	s := []byte("<UNK>")
	s2 := *(*string)(unsafe.Pointer(&s))
	fmt.Println(s2)
}

func JoinStr() {
	s1 := "ss"
	s2 := "ss2"

	var buf bytes.Buffer
	buf.WriteString(s1)
	buf.WriteString(s2)

	fmt.Println(buf.String())

	// unsafe

}

func appendS(s []int) {
	for i := 0; i < 6; i++ {
		s = append(s, i)
	}
}

func deferDO() (n int) {
	defer fmt.Println("world1")
	defer fmt.Println("world2")

	n = 1
	defer func(n int) {
		n = n + 1
	}(n)

	return n
}

func modifySlice(nums []int, val int) {
	nums[0] = val
}

func appendSlice(nums []int) []int {
	nums = append(nums, 9)
	return nums
}
