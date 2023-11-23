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
	linkedObj.insert(6)
	linkedObj.insert(5)
	linkedObj.insert(4)
	linkedObj.insert(3)
	linkedObj.insert(2)
	linkedObj.insert(1)
	linkedObj.insert(0)
	//fmt.Printf("当前节点: %+v\n", linkedObj)
	//fmt.Printf("当前节点: %+v\n", linkedObj.head.next.data)
	//linkedObj.foreach()
	linkedObj.foreachRev()
}

type linked[T int | string | map[string]string] struct {
	head   *node[T]
	end    *node[T]
	length int
	limit  int
}

type node[T int | string | map[string]string] struct {
	data T
	next *node[T]
	prev *node[T]
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
		prev: nil,
	}
}

func (l *linked[T]) insert(data T) bool {
	newNode := createNode(data)

	newNode.prev = l.head
	newNode.next = l.head.next
	l.head.next = newNode

	headNode := l.head
	for headNode.next != nil {
		headNode = headNode.next
	}

	l.end = headNode

	if l.length >= l.limit {
		l.delete(l.end)
	} else {
		l.length++
	}

	return true
}

func (l *linked[T]) delete(node *node[T]) {

}

func (l *linked[T]) foreach() {
	headNode := l.head
	for headNode.next != nil {
		headNode = headNode.next
		fmt.Printf("当前节点: %+v\n", headNode.data)
	}
}

func (l *linked[T]) foreachRev() {
	endNode := l.end
	for endNode.prev != nil {
		fmt.Printf("当前节点: %+v\n", endNode.data)
		endNode = endNode.prev
	}
}
