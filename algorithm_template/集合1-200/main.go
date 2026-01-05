package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
}

// 1.两数之和
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i <= len(nums); i++ {
		if j, ok := m[target-nums[i]]; ok {
			return []int{j, i}
		}
		m[nums[i]] = i
	}
	return []int{-1, -1}
}

// 3. 最长无重复子串
func lengthOfLongestSubstring(s string) int {
	count, left, right := 0, 0, 0
	win := make(map[byte]int)
	for right < len(s) {
		c := s[right]
		win[c]++
		right++
		for win[c] > 1 {
			d := s[left]
			left++
			win[d]--
		}
		if count < right-left {
			count = right - left
		}
	}
	return count
}

// 146. lru缓存

type Node struct {
	key, val   int
	next, prev *Node
}

type LRUCache struct {
	capacity   int
	data       map[int]*Node
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		data:     make(map[int]*Node),
		head:     &Node{},
		tail:     &Node{},
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.data[key]; ok {
		l.MoveToHead(node)
		return node.val
	}
	return -1
}

func (l *LRUCache) Put(key, val int) {
	if node, ok := l.data[key]; ok {
		l.MoveToHead(node)
		node.val = val
		return
	}

	newNode := NewNode(key, val)

	if l.capacity > len(l.data) {
		//
		tail := l.RemoveTail()
		delete(l.data, tail.key)
	}

	l.AddToHead(newNode)
}

func NewNode(key, val int) *Node {
	return &Node{
		key: key,
		val: val,
	}
}

func (l *LRUCache) RemoveNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache) AddToHead(node *Node) {
	node.next = l.head.next
	node.prev = l.head
	l.head.next.prev = node
	l.head.next = node

}

func (l *LRUCache) MoveToHead(node *Node) {
	l.RemoveNode(node)
	l.AddToHead(node)
}

func (l *LRUCache) RemoveTail() *Node {
	tail := l.tail.prev
	l.RemoveNode(tail)
	return tail
}

//
// 146. lru缓存v2 直接使用container
