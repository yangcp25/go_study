package main

import (
	"container/heap"
	"container/list"
	"fmt"
	"math"
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

func findKthLargest(nums []int, k int) int {
	h := &minHeadp{}
	heap.Init(h)

	for _, v := range nums {
		heap.Push(h, v)

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return heap.Pop(h).(int)
}

type minHeadp []int

func (h *minHeadp) Len() int {
	return len(*h)
}

func (h minHeadp) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeadp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeadp) Push(n interface{}) {
	*h = append(*h, n.(int))
}

func (h *minHeadp) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 560 和为k的子数组

func subarraySum(nums []int, k int) int {
	set := map[int]int{0: 1}
	sum, count := 0, 0
	for _, v := range nums {
		sum += v
		if num, ok := set[sum-k]; ok {
			count += num
		}
		set[sum]++
	}
	return count
}

// 56 合并区间

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	start, end := intervals[0][0], intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if end >= intervals[i][0] {
			end = Max(intervals[i][1], end)
		} else {
			res = append(res, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		}
	}
	res = append(res, []int{start, end})
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 88 合并2个有序数组

func merge2(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1

	for j >= 0 {
		if i >= 0 && nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

// 4. 寻找2个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	m, n := len(nums1), len(nums2)

	left, right := 0, m

	for left <= right {
		i := (left + right) / 2
		j := (m+n+1)/2 - i

		ALeft := math.MinInt
		if i > 0 {
			ALeft = nums1[i-1]
		}

		ARight := math.MaxInt
		if i < m {
			ARight = nums1[i]
		}

		BLeft := math.MinInt
		if j > 0 {
			BLeft = nums2[j-1]
		}

		BRight := math.MaxInt
		if j < n {
			BRight = nums2[j]
		}

		if ALeft <= BRight && BLeft <= ARight {
			if (m+n)%2 == 1 {
				return float64(Max(ALeft, BLeft))
			}

			return float64(Max(ALeft, BLeft)+Min(ARight, BRight)) / 2
		}

		if ALeft > BRight {
			right = i - 1
		} else {
			left = i + 1
		}
	}

	return 0
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// 14. 最长公共前缀

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	i := 0
Outlook:
	for ; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i < len(strs[j]) && strs[0][i] == strs[j][i] {

			} else {
				break Outlook
			}
		}
	}

	return strs[0][0:i]
}
