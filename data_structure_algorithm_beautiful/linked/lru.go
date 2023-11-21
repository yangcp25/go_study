package main

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
}

type linked[T int | string | map[string]string] struct {
	head   *node[T]
	length int
	limit  int
}

type node[T int | string | map[string]string] struct {
	data T
	prev *node[T]
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
		prev: nil,
		next: nil,
	}
}

func (l *linked[T]) insert(data T) bool {
	if l.length == 0 {
		l.head.data = data
		return true
	}
	if l.limit == l.length {
		l.head.prev.data = data
		l.head = l.head.prev
		return true
	}
	newNode := createNode(data)
	newNode.next = l.head
	l.head.prev = newNode
	l.length++
	return true
}

func (l *linked[T]) delete() {
	if l.length == 0 {
		return
	}
}
