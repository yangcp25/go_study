package main

import "fmt"

func main() {
	// 实现一个lru 淘汰算法
	// linked 结构体
	// node 节点 ： data prev next
	// 更新lru
	// 如果没有满
	// 将新的数据加入到头结点
	// 队满 ： 删除尾结点
	// 将新数据加入头结点
	linkedObj := getLinked[int](5)
	linkedObj.insert(6)
	linkedObj.insert(5)
	linkedObj.insert(4)
	linkedObj.insert(3)
	linkedObj.insert(2)
	linkedObj.insert(1)
	linkedObj.insert(0)
	//fmt.Printf("当前节点: %+v\n", linkedObj)
	fmt.Printf("当前节点: %+v\n", linkedObj.head.next.data)
	linkedObj.foreach()
}

type linked[T int | string | map[string]string] struct {
	head   *node[T]
	length int
	limit  int
}

type node[T int | string | map[string]string] struct {
	data T
	next *node[T]
}

func getLinked[T int | string | map[string]string](limit int) *linked[T] {
	headNode := &node[T]{}
	return &linked[T]{
		head:   headNode,
		length: 0,
		limit:  limit,
	}
}

func createNode[T int | string | map[string]string](data T) *node[T] {
	return &node[T]{
		data: data,
		next: nil,
	}
}

func (l *linked[T]) insert(data T) bool {
	newNode := createNode(data)

	headNode := l.head.next
	newNode.next = l.head.next
	l.head.next = newNode

	if l.length == l.limit {
		prevNode := headNode
		for headNode.next != nil {
			prevNode = headNode
			headNode = headNode.next
		}
		prevNode.next = nil
	} else {
		l.length++
	}
	return true
}

func (l *linked[T]) foreach() {
	headNode := l.head.next
	for headNode.next != nil {
		headNode = headNode.next
		fmt.Printf("当前节点: %+v\n", headNode.data)
	}
}
