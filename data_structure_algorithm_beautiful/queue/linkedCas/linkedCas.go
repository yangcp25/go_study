package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

// cas版本
func main() {
	casQueue1 := createCasQueue[int]()

	casQueue1.tailPush(1)
	casQueue1.tailPush(2)
	casQueue1.tailPush(3)
	casQueue1.tailPush(4)
	casQueue1.tailPush(5)

	data1 := casQueue1.headPop()
	fmt.Printf("弹出节点：%+v\n", *data1)

	casQueue1.tailPush(6)

	data2 := casQueue1.headPop()
	fmt.Printf("弹出节点：%+v\n", *data2)

	casQueue1.headForeach()
}

type casQueue[T int | string] struct {
	head *node[T] // 将head 申明成unsafe.Pointer 是因为后面的atomic.CompareAndSwapPointer 需要这种类型的指针
	tail *node[T]
}

type node[T int | string] struct {
	next *node[T]
	data T
}

// 初始化队列
func createCasQueue[T int | string]() *casQueue[T] {
	return &casQueue[T]{
		head: nil,
		tail: nil,
	}
}

// 创建节点
func createNode[T int | string](data T) *node[T] {
	return &node[T]{
		data: data,
	}
}

// 尾部插入
func (queue *casQueue[T]) tailPush(data T) bool {
	newNode := createNode(data)
	if queue.tail == nil {
		atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(queue.head)), unsafe.Pointer(queue.head), unsafe.Pointer(newNode))
		atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(queue.tail)), unsafe.Pointer(queue.tail), unsafe.Pointer(newNode))
		return true
	}
	queue.tail.next = newNode
	queue.tail = newNode
	return true
}

// 头部弹出
func (queue *casQueue[T]) headPop() (node *T) {
	if queue.head == nil {
		return nil
	}
	node = &queue.head.data
	queue.head = queue.head.next
	if queue.head == nil {
		queue.tail = nil
	}
	return
}

func (queue *casQueue[T]) headForeach() {
	cur := queue.head
	fmt.Print("从头部遍历当前队列：\n")
	for {
		fmt.Printf("当前节点：%+v\n", cur.data)
		if cur == queue.tail {
			break
		}
		cur = cur.next
	}
}
