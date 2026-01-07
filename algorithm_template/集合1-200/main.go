package main

import (
	"container/list"
	"fmt"
	"sort"
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

// 146. lru缓存v2 直接使用container
type LRUCacheV2 struct {
	capacity int
	data     map[int]*list.Element
	list     *list.List
}

func Constructor2(capacity int) LRUCacheV2 {
	element := list.New()
	lru := LRUCacheV2{
		capacity: capacity,
		data:     make(map[int]*list.Element),
		list:     element,
	}
	return lru
}

type entry struct {
	key, val int
}

func (l *LRUCacheV2) Get(key int) int {
	if node, ok := l.data[key]; ok {
		l.list.MoveToFront(node)
		return node.Value.(entry).val
	}
	return -1
}

func (l *LRUCacheV2) Put(key, val int) {
	if node, ok := l.data[key]; ok {
		l.list.MoveToFront(node)
		node.Value = entry{key, val}
		return
	}

	e := l.list.PushFront(entry{key, val})
	l.data[key] = e
	if l.capacity > len(l.data) {
		//
		tail := l.list.Back()
		delete(l.data, tail.Value.(entry).key)
	}
	sort.Slice(l.data, func(i, j int) bool {})

}

// 42 接雨水

func trap(height []int) int {
	count, leftMax, rightMax := 0, 0, 0
	left, right := 0, len(height)-1
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				count += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				count += rightMax - height[right]
			}
			right--
		}
	}
	return count
}

// 15 三数之和

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	res := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
