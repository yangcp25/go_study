package main

import (
	"errors"
	"fmt"
	"sync"
)

// todo 待做cas版本
func main() {
	// 双向循环链表实现队列 加锁实现并发安全
	linkedObj := getLinked[int](5)

	syncGroup := sync.WaitGroup{}
	syncGroup.Add(1050)

	for i := 0; i < 1000; i++ {
		i := i
		go func() {
			err := linkedObj.headPush(i)
			if err != nil {
				fmt.Printf(err.Error())
			}
			syncGroup.Done()
		}()
	}
	for i := 0; i < 50; i++ {
		go func() {
			data := linkedObj.tailPop()
			if data != nil {
				fmt.Println(*data)
			}
			syncGroup.Done()
		}()
	}
	//fmt.Printf("当前节点: %+v\n", linkedObj)
	//fmt.Printf("当前节点: %+v\n", linkedObj.head.next.data)

	syncGroup.Wait()
	linkedObj.headForeach()
}

type linked[T int | string | map[string]string] struct {
	head     *node[T]
	length   int
	limit    int
	headLock sync.Mutex
	tailLock sync.Mutex
}

type node[T int | string | map[string]string] struct {
	data *T
	next *node[T]
	prev *node[T]
}

func getLinked[T int | string | map[string]string](limit int) *linked[T] {
	return &linked[T]{
		head:     nil,
		length:   0,
		limit:    limit,
		headLock: sync.Mutex{},
		tailLock: sync.Mutex{},
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
	l.headLock.Lock()
	defer l.headLock.Unlock()
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
	l.tailLock.Lock()
	defer l.tailLock.Unlock()
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

	//currentNode.prev.next = headNodePos
	//headNodePos.prev = currentNode.prev
	if currentNode == headNodePos {
		// The list has only one element
		l.head = nil
	} else {
		currentNode.prev.next = headNodePos
		headNodePos.prev = currentNode.prev
	}

	l.length--
	return currentNode.data
}

// 从头部遍历
func (l *linked[T]) headForeach() {
	if l.head == nil {
		fmt.Printf("队空：\n")
		return
	}
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
	if l.head == nil {
		fmt.Printf("队空：\n")
		return
	}
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
