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

type node struct {
	next *node
	pre  *node
	val  int
}
type MapValue struct {
	node *node
}
type LruCache struct {
	capacity int
	data     map[int]MapValue
	head     *node
	tail     *node
}

func Constructor(capacity int) *LruCache {
	return &LruCache{
		capacity: capacity,
		//head:     &node{},
		//tail:     &node{},
		data: make(map[int]MapValue, capacity),
	}
}

// 1 2 3 4
func (l *LruCache) Get(key int) (int, bool) {
	// 从map拿数据
	if v, ok := l.data[key]; ok {
		if v.node.pre != nil {
			// 更新链表，将数据放到头部
			newNode := &node{
				val: v.node.val,
			}

			newNode.pre = nil
			newNode.next = l.head
			l.head = newNode
			// 删除该节点
			v.node.pre.next = v.node.next
		}
		return v.node.val, true
	} else {
		return 0, false
	}
}

func (l *LruCache) Put(key int, value int) {
	//head :=
	if v, ok := l.data[key]; ok {
		v.node.val = value
		// 更新到头节点
		//head.next = v.node
		//v.node.pre = head
		//v.node.next = v.node.pre.next
		return
	} else {
		newNode := MapValue{
			node: &node{
				val: value,
			},
		}
		l.data[key] = newNode
		if l.head == nil {
			l.head = newNode.node
		} else {
			newNode.node.pre = nil
			temp := l.head
			l.head = newNode.node
			newNode.node.next = temp
			temp.pre = newNode.node
		}
		length := len(l.data)
		// 更新尾节点
		if l.tail == nil {
			l.tail = newNode.node
		}
		if length > l.capacity {
			// 删除尾部节点 更新尾节点
			l.tail = l.tail.pre
			l.tail.next = nil
		}
	}
}
