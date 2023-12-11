package main

import (
	"errors"
	"fmt"
)

func main() {
	stack1 := createStack[int](3)
	err := stack1.push(1)
	// 处理错误 后面的就不处理了
	if err != nil {
		return
	}
	stack1.push(2)
	stack1.push(3)
	stack1.push(4)
	stack1.push(5)

	popData1 := stack1.pop()
	fmt.Printf("弹出栈元素%+v\n", *popData1)
	popData2 := stack1.pop()
	fmt.Printf("弹出栈元素%+v\n", *popData2)

	stack1.push(6)
	stack1.push(7)

	stack1.forEach()
}

type stack[T int | string | map[string]string] struct {
	data  []T
	limit int
}

func createStack[T int | string | map[string]string](len int) *stack[T] {
	return &stack[T]{
		data:  make([]T, 0, len),
		limit: len,
	}
}

func (s *stack[T]) push(item T) error {
	if len(s.data) >= s.limit {
		return errors.New("超过栈长度！")
	}
	s.data = append(s.data, item)
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
