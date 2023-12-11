package main

import "fmt"

func main() {
	// 实现一个lru 淘汰算法
	// 双向循环链表
	// linked 结构体
	// node 节点 ： data prev next
	// 更新lru
	// 如果没有满
	// 将新的数据加入到头结点
	// 队满 ： 删除尾结点
	// 将新数据加入头结点
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
	linkedObj.tailForeach()
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
		l.length++
		newNode.next = newNode
		newNode.prev = newNode
		return true
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
	return true
}

func (l *linked[T]) delete(node *node[T]) {

}

// 从头部遍历
func (l *linked[T]) headForeach() {
	headNode := l.head
	headNodPos := headNode
	fmt.Printf("从头结点遍历：\n")
	for {
		fmt.Printf("当前节点: %+v\n", headNode.data)
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
		fmt.Printf("当前节点: %+v\n", endNode.prev.data)
		if endNode.prev == endNodePos {
			break
		}
		endNode = endNode.prev
	}
}
