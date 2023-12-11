package main

import (
	"errors"
	"fmt"
)

func main() {
	stack1 := createStack[string](1024)
	err := stack1.push("a")
	// 处理错误 后面的就不处理了
	if err != nil {
		return
	}
	stack1.push("a")
	fmt.Printf("当前栈容量%+v\n", cap(stack1.data))
	stack1.push("b")
	stack1.push("c")
	stack1.push("d")
	fmt.Printf("当前栈容量%+v\n", cap(stack1.data))

	popData1 := stack1.pop()
	fmt.Printf("弹出栈元素%+v\n", *popData1)
	popData2 := stack1.pop()
	fmt.Printf("弹出栈元素%+v\n", *popData2)

	stack1.push("e")
	stack1.push("f")

	fmt.Printf("当前栈容量%+v\n", cap(stack1.data))

	stack1.forEach()
}

type stack[T int | string | map[string]string] struct {
	data  []T
	limit int
}

func createStack[T int | string | map[string]string](len int) *stack[T] {
	return &stack[T]{
		data:  make([]T, 0, len/1024), // 简化一下 初始时 只申请很小的长度
		limit: len,
	}
}

// 实现自动扩容 简单版扩容策略 增加为原容量的2倍
func (s *stack[T]) push(item T) error {
	length := len(s.data)
	if length >= s.limit {
		return errors.New("超过栈长度！")
	}

	if length+1 > cap(s.data) {
		newArray := make([]T, 0, length*2)
		newArray = append(newArray, s.data...)
		s.data = append(newArray, item)
	} else {
		s.data = append(s.data, item)
	}

	return nil
}

func (s *stack[T]) pop() *T {
	length := len(s.data)
	if length == 0 {
		return nil
	}
	res := &s.data[length-1]
	s.data = s.data[0 : length-1]
	return res
}

func (s *stack[T]) forEach() {
	fmt.Printf("遍历栈：\n")
	for _, item := range s.data {
		fmt.Printf("当前栈元素%+v\n", item)
	}
}
