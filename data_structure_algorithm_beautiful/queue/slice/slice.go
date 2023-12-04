package main

import (
	"fmt"
)

func main() {
	// 创建一个简单队列
	// 如果head == tail 队列空
	// 如果tail == len(array) - 1
	// 整体做迁移 如果head == 0 队列满
	stack1 := createQueue[int](5)
	err := stack1.push(1)
	// 处理错误 后面的就不处理了
	if err != nil {
		return
	}
	stack1.push(2)
	fmt.Printf("队列容量%+v\n", cap(stack1.data))
	stack1.push(3)
	stack1.push(4)
	stack1.push(5)

	popData1 := stack1.pop()
	fmt.Printf("出队列元素%+v\n", *popData1)
	popData2 := stack1.pop()
	fmt.Printf("出队列元素%+v\n", *popData2)
	stack1.forEach()
	stack1.push(6)
	stack1.push(7)
	stack1.push(8)
	stack1.push(9)
	// 看是否自动扩容
	fmt.Printf("队列容量%+v\n", cap(stack1.data))
	popData3 := stack1.pop()
	fmt.Printf("出队列元素%+v\n", *popData3)
	err = stack1.push(10)
	// 处理错误 后面的就不处理了
	if err != nil {
		fmt.Printf(err.Error() + "\n")
	}
	stack1.forEach()
}

type queue[T int | string | map[string]string] struct {
	data []T
}

func createQueue[T int | string | map[string]string](len int) *queue[T] {
	return &queue[T]{
		data: make([]T, 0, len),
	}
}

func (s *queue[T]) push(item T) error {
	s.data = append(s.data, item)
	return nil
}

// 数组头部出队列
func (s *queue[T]) pop() *T {
	// 队列为空
	if len(s.data) == 0 {
		return nil
	}
	res := &s.data[0]
	s.data = s.data[1:]
	return res
}

func (s *queue[T]) forEach() {
	fmt.Printf("遍历队列：\n")
	for _, datum := range s.data {
		fmt.Printf("当前队列元素%+v\n", datum)
	}

}
