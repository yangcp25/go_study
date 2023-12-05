package main

import (
	"errors"
	"fmt"
)

func main() {
	// 双向循环链表实现队列
	linkedObj := getLinked[int](5)
	err := linkedObj.headPush(6)
	if err != nil {
		fmt.Printf(err.Error())
	}
	err = linkedObj.headPush(5)
	if err != nil {
		fmt.Printf(err.Error())
	}
	err = linkedObj.headPush(4)
	if err != nil {
		fmt.Printf(err.Error())
	}
	err = linkedObj.headPush(3)
	if err != nil {
		fmt.Printf(err.Error())
	}
	err = linkedObj.headPush(2)
	if err != nil {
		fmt.Printf(err.Error())
	}
	err = linkedObj.headPush(1)
	if err != nil {
		fmt.Printf(err.Error())
	}
	err = linkedObj.headPush(0)
	if err != nil {
		fmt.Printf(err.Error())
	}
	//fmt.Printf("当前节点: %+v\n", linkedObj)
	//fmt.Printf("当前节点: %+v\n", linkedObj.head.next.data)
	item := linkedObj.tailPop()
	fmt.Printf("弹出节点: %+v\n", *item)
	item = linkedObj.tailPop()
	fmt.Printf("弹出节点: %+v\n", *item)
	linkedObj.headForeach()
	err = linkedObj.headPush(-1)
	if err != nil {
		fmt.Printf(err.Error())
	}
	linkedObj.tailForeach()
}

type linked[T int | string | map[string]string] struct {
	head   *node[T]
	length int
	limit  int
}

type node[T int | string | map[string]string] struct {
	data *T
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
		data: &data,
		next: nil,
		prev: nil,
	}
}

// 从头部插入
func (l *linked[T]) headPush(data T) error {

	if l.length >= l.limit {
		return errors.New("当前队满\n")
	}
	newNode := createNode(data)

	if l.head == nil {
		l.head = newNode
		l.length++
		newNode.next = newNode
		newNode.prev = newNode
		return nil
	}

	currentNode := l.head
	// 头结点位置
	headNodePos := l.head

	l.head = newNode
	newNode.next = currentNode
	currentNode.prev = newNode

	// 找尾结点
	for {
		if currentNode.next == headNodePos {
			break
		}
		currentNode = currentNode.next
	}

	if l.length >= l.limit {
		currentNode.prev.next = newNode
		newNode.prev = currentNode.prev
	} else {
		currentNode.next = newNode
		newNode.prev = currentNode
		l.length++
	}
	return nil
}

// 尾部弹出
func (l *linked[T]) tailPop() *T {
	if l.head == nil {
		return nil
	}

	currentNode := l.head
	// 头结点位置
	headNodePos := l.head
	// 找尾结点
	for {
		if currentNode.next == headNodePos {
			break
		}
		currentNode = currentNode.next
	}
	// 将尾结点的前一个节点作为尾节点
	currentNode.prev.next = headNodePos
	headNodePos.prev = currentNode.prev
	l.length--
	return currentNode.data
}

// 从头部遍历
func (l *linked[T]) headForeach() {
	headNode := l.head
	headNodPos := headNode
	fmt.Printf("从头结点遍历：\n")
	for {
		fmt.Printf("当前节点: %+v\n", *headNode.data)
		if headNode.next == headNodPos {
			break
		}
		headNode = headNode.next
	}
}

// 从尾部遍历
func (l *linked[T]) tailForeach() {
	endNode := l.head
	endNodePos := endNode
	fmt.Printf("从尾结点遍历：\n")
	for {
		fmt.Printf("当前节点: %+v\n", *endNode.prev.data)
		if endNode.prev == endNodePos {
			break
		}
		endNode = endNode.prev
	}
}
