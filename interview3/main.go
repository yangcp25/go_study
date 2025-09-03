package main

import "fmt"

// func main() {
//
//		//LRU
//		//实现一个 LRU 缓存，要求支持以下操作：
//		//- Get(key int)(int, bool)：如果缓存中存在该 key，返回其对应的值，否则返回 0,false。
//		//- Put(key, value int)：将一个 key-value 对插入缓存。如果缓存已经满了，淘汰最久未使用的元素。
//		//
//		//示例：
//		//func main() {
//		//    // 构造 一个 LRU cache
//		//    cache := Constructor(2)
//		//    cache.Put(1, 1)
//		//    cache.Put(2, 2)
//		//
//		//    // 测试缓存获取
//		//    fmt.Println(cache.Get(1)) // 输出: 1 true
//		//    fmt.Println(cache.Get(3)) // 输出: 0 false
//		//
//		//    cache.Put(3, 3)
//		//    fmt.Println(cache.Get(2)) // 输出: 0 false
//		//
//		//    cache.Put(4, 4)
//		//    fmt.Println(cache.Get(1)) // 输出: 0 false
//		//    fmt.Println(cache.Get(3)) // 输出: 3 true
//		//    fmt.Println(cache.Get(4)) // 输出: 4 true
//		//}
//	}
func main() {
	// 构造 一个 LRU cache
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)

	// 测试缓存获取
	fmt.Println(cache.Get(1)) // 输出: 1 true
	fmt.Println(cache.Get(3)) // 输出: 0 false

	cache.Put(3, 3)
	fmt.Println(cache.Get(2)) // 输出: 0 false

	cache.Put(4, 4)
	fmt.Println(cache.Get(1)) // 输出: 0 false
	fmt.Println(cache.Get(3)) // 输出: 3 true
	fmt.Println(cache.Get(4)) // 输出: 4 true
	// 测试缓存获取

	// 输出: 4 true
}

type Node struct {
	next *Node
	pre  *Node
	val  int
	key  int
}

type LruCache struct {
	capacity   int
	size       int
	cache      map[int]*Node
	head, tail *Node
}

func Constructor(capacity int) *LruCache {
	return &LruCache{
		capacity: capacity,
		cache:    make(map[int]*Node, capacity),
		head:     &Node{},
		tail:     &Node{},
	}
}

// 1 2 3 4
func (l *LruCache) Get(key int) (int, bool) {
	// 从map拿数据
	if v, ok := l.cache[key]; ok {
		return v.val, true
	} else {
		return -1, false
	}
}

func (l *LruCache) Put(key int, val int) {
	l.AddToHead(key, val)
	l.size++
	if l.size > l.capacity {
		l.RemoveTail()
	}
}

func InitNode(key, val int) *Node {
	return &Node{
		key: key,
		val: val,
	}
}

// 添加到头节点
func (l *LruCache) AddToHead(key, val int) {
	newNode := InitNode(key, val)
	newNode.next = l.head.next
	l.head.next = newNode
}

func (l *LruCache) RemoveNode(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (l *LruCache) RemoveTail(node *Node) {
	l.RemoveNode(l.tail)
}

func (l *LruCache) MoveToHead() {}

// 添加节点
// 删除节点
// 删除尾节点
