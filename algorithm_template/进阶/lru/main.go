package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 构造 一个 LRU cache
	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(1, 1)
	cache.Put(2, 3)
	cache.Put(4, 1)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))

}

type Node struct {
	next *Node
	pre  *Node
	val  int
	key  int
}

type LRUCache struct {
	capacity   int
	size       int
	cache      map[int]*Node
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node, capacity),
		head:     &Node{},
		tail:     &Node{},
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

// 1 2 3 4
func (l *LRUCache) Get(key int) int {
	// 从map拿数据
	if _, ok := l.cache[key]; !ok {
		return -1
	}
	node := l.cache[key]
	l.MoveToHead(node)
	return node.val
}

func (l *LRUCache) Put(key int, val int) {
	if _, ok := l.cache[key]; ok {
		l.cache[key].val = val
		l.MoveToHead(l.cache[key])
	} else {
		l.size++
		node := InitNode(key, val)
		l.cache[key] = node
		l.AddToHead(node)
		if l.size > l.capacity {
			node := l.RemoveTail()
			l.size--
			delete(l.cache, node.key)
		}
	}

}

func InitNode(key, val int) *Node {
	return &Node{
		key: key,
		val: val,
	}
}

// 添加到头节点
func (l *LRUCache) AddToHead(newNode *Node) *Node {
	newNode.next = l.head.next
	newNode.pre = l.head
	l.head.next.pre = newNode
	l.head.next = newNode
	return newNode
}

func (l *LRUCache) RemoveNode(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (l *LRUCache) RemoveTail() *Node {
	node := l.tail.pre
	//l.tail.pre = node.pre
	l.RemoveNode(node)
	return node
}

func (l *LRUCache) MoveToHead(node *Node) {
	l.RemoveNode(node)
	l.AddToHead(node)
}

type LRUCacheV2 struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type entry struct {
	key, value int
}

func ConstructorV2(capacity int) LRUCacheV2 {
	return LRUCacheV2{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (l *LRUCacheV2) Get(key int) int {
	if e, ok := l.cache[key]; ok {
		l.list.MoveToFront(e)
		return e.Value.(entry).value
	}
	return -1
}

func (l *LRUCacheV2) Put(key, value int) {
	if e, ok := l.cache[key]; ok {
		l.list.MoveToFront(e)
		e.Value = entry{key, value}
		return
	}
	if l.list.Len() == l.capacity {
		back := l.list.Back()
		delete(l.cache, back.Value.(entry).key)
		l.list.Remove(back)
	}
	e := l.list.PushFront(entry{key, value})
	l.cache[key] = e
}
