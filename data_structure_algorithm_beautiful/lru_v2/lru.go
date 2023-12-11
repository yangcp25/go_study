package main

import "fmt"

func main() {
	// 实现一个双向循环链表
	linkedObj := getLinked[int](5)
	linkedObj.headInsert(6)
	linkedObj.headInsert(5)
	linkedObj.headInsert(4)
	linkedObj.headInsert(3)
	linkedObj.headInsert(2)
	linkedObj.headInsert(1)
	linkedObj.headInsert(0)
	//fmt.Printf("当前节点: %+v\n", linkedObj)
	//fmt.Printf("当前节点: %+v\n", linkedObj.head.next.data)
	linkedObj.headForeach()
	//linkedObj.tailForeach()
}

type linked[T int | string | map[string]string] struct {
	head   *node[T]
	length int
	limit  int
}

type node[T int | string | map[string]string] struct {
	data T
	next *node[T]
	prev *node[T]
}

func getLinked[T int | string | map[string]string](limit int) *linked[T] {
	return &linked[T]{
		head:   nil,
		length: 0,
		limit:  limit,
	}
}

func createNode[T int | string | map[string]string](data T) *node[T] {
	return &node[T]{
		data: data,
		next: nil,
		prev: nil,
	}
}

// 从头部插入
func (l *linked[T]) headInsert(data T) bool {
	newNode := createNode(data)

	if l.head == nil {
		l.head = newNode
		l.head.next = newNode
		l.head.prev = newNode
		l.length++
		return true
	}

	// 原头结点
	currentNode := l.head
	headNode := currentNode

	l.head = newNode
	newNode.next = currentNode
	currentNode.prev = newNode

	// 找到尾结点
	for {
		if currentNode.next == headNode {
			break
		}
		currentNode = currentNode.next
	}

	if l.length >= l.limit {
		currentNode.prev.next = l.head
		l.head.prev = currentNode.prev
	} else {
		l.head.prev = currentNode
		l.length++
	}
	return true
}

func (l *linked[T]) delete(node *node[T]) {

}

// 从头部遍历
func (l *linked[T]) headForeach() {
	headNode := l.head
	fmt.Printf("从头结点遍历：\n")
	for {
		fmt.Printf("当前节点: %+v\n", headNode.data)
		if headNode.next == l.head {
			break
		}
		headNode = headNode.next
	}
}

// 从尾部遍历
func (l *linked[T]) tailForeach() {
	endNode := l.head.prev
	fmt.Printf("从尾结点遍历：\n")
	for {
		fmt.Printf("当前节点: %+v\n", endNode.data)
		if endNode.prev == l.head.prev {
			break
		}
		endNode = endNode.prev
	}
}
