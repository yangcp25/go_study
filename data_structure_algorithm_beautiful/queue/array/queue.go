package main

import (
	"errors"
	"fmt"
)

func main() {
	// 创建一个简单队列
	// 如果head == tail 队列空
	// 如果tail == len(array) - 1
	// 整体做迁移 如果head == 0 队列满
	stack1 := createQueue[int]()
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
	fmt.Printf("出队列元素%+v\n", *popData1)
	popData2 := stack1.pop()
	fmt.Printf("出队列元素%+v\n", *popData2)

	stack1.push(6)
	stack1.push(7)
	stack1.push(8)
	stack1.push(9)
	stack1.pop()
	err = stack1.push(10)
	// 处理错误 后面的就不处理了
	if err != nil {
		fmt.Printf(err.Error() + "\n")
	}
	stack1.forEach()
}

// 默认小一点的空间 size 为5 数组空间=size+1
type queue[T int | string | map[string]string] struct {
	data [6]T
	head int
	tail int
}

func createQueue[T int | string | map[string]string]() *queue[T] {
	return &queue[T]{
		data: [6]T{},
	}
}

func (s *queue[T]) push(item T) error {
	// len(s.data) - 1 这里固定为5
	if len(s.data)-1 == s.tail {
		if s.head == 0 {
			return errors.New("队列满！")
		}
		// 做迁移
		currentTail := s.tail - s.head
		// 队列整体往前移动
		for i := 0; i < currentTail; i++ {
			s.data[i] = s.data[i+s.head]
		}
		s.head = 0
		s.tail = currentTail
	}
	s.data[s.tail] = item
	s.tail++
	return nil
}

// 数组头部出队列
func (s *queue[T]) pop() *T {
	// 队列为空
	if s.head == s.tail {
		return nil
	}
	res := &s.data[s.head]
	s.head++
	return res
}

func (s *queue[T]) forEach() {
	fmt.Printf("遍历队列：\n")
	for i := s.head; i < s.tail; i++ {
		fmt.Printf("当前队列元素%+v\n", s.data[i])
	}
}
