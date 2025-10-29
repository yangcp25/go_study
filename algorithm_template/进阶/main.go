package main

func main() {

}

// 实现lru
type Node struct {
	key, val   int
	prev, next *Node
}

type LRUCache struct {
	capacity   int
	head, tail *Node
	cache      map[int]*Node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		head:     &Node{},
		tail:     &Node{},
		cache:    make(map[int]*Node, capacity+1),
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}
func (l LRUCache) Get(key int) int {
	if node, ok := l.cache[key]; ok {
		l.MoveToHead(node)
		return node.val
	}
	return -1
}
func (l LRUCache) Put(key int, value int) {
	if node, ok := l.cache[key]; ok {
		node.val = value
		l.MoveToHead(node)
	} else {
		node := &Node{key, value, nil, nil}
		l.cache[key] = node
		l.AddNodeToHead(node)
		if len(l.cache) > l.capacity {
			node := l.RemoveTail()
			delete(l.cache, node.key)
		}
	}
}

func (l LRUCache) MoveToHead(node *Node) {
	l.RemoveNode(node)
	l.AddNodeToHead(node)
}
func (l LRUCache) RemoveNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
func (l LRUCache) AddNodeToHead(node *Node) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}
func (l LRUCache) RemoveTail() *Node {
	node := l.tail.prev
	l.RemoveNode(node)
	return node
}

// 优先级队列

// 图
