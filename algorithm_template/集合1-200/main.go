package main

import (
	"container/heap"
	"container/list"
	"fmt"
	"math"
	"math/rand"
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

// 54. 螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	//
	m, n := len(matrix), len(matrix[0])
	left, right, top, bottom := 0, n-1, 0, m-1
	res := make([]int, 0)
	for top <= bottom && left <= right {
		// 上
		for col := left; col <= right; col++ {
			res = append(res, matrix[top][col])
		}
		top++
		// 右
		for row := top; row <= bottom; row++ {
			res = append(res, matrix[row][right])
		}
		right--

		if top > bottom || left > right {
			break
		}

		// 下
		for col := right; col >= left; col-- {
			res = append(res, matrix[bottom][col])
		}
		bottom--
		// 左边

		for row := bottom; row >= top; row-- {
			res = append(res, matrix[row][left])
		}
		left++
	}

	return res
}

// 21 合并2个有序链表
/**
 * Definition for a singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}

	if list1 != nil {
		head.Next = list1
	} else {
		head.Next = list2
	}

	return dummy.Next
}

// 704. 二分查找

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// 9 回文数

func isPalindrome(x int) bool {
	// 边界
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reverted := 0
	for x > reverted {
		reverted = reverted*10 + x%10
		x /= 10
	}

	// 偶数位 or 奇数位
	return x == reverted || x == reverted/10
}

// 46 全排列

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make(map[int]bool)
	var backtrack func(nums []int, path []int, used map[int]bool)

	backtrack = func(nums []int, path []int, used map[int]bool) {
		if len(path) >= len(nums) {
			temp := make([]int, len(nums))
			copy(temp, path)
			res = append(res, temp)
		}
		for i := 0; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			path = append(path, nums[i])
			used[nums[i]] = true
			backtrack(nums, path, used)
			path = path[:len(path)-1]
			used[nums[i]] = false
		}
	}

	backtrack(nums, path, used)

	return res
}

// 27 移除元素
func removeElement(nums []int, val int) int {
	i, j := 0, 0

	for j < len(nums) {
		if nums[j] != val {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
	return i
}

// 41 缺失的第一个正数

func firstMissingPositive(nums []int) int {
	n := len(nums)
	// 找到每个数的位置
	for i := 0; i < n; i++ {
		for nums[i] >= 1 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	//
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

// 55 跳跃游戏
func canJump(nums []int) bool {
	maxReach := 0

	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
		if maxReach < i+nums[i] {
			maxReach = i + nums[i]
		}
	}
	return true
}

// 23 合并k个升序链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	h := MinHeap{}
	heap.Init(&h)
	n := len(lists)

	for i := 0; i < n; i++ {
		if lists[i] != nil {
			heap.Push(&h, lists[i])
		}
	}
	dummy := &ListNode{}
	head := dummy
	for h.Len() > 0 {
		node := heap.Pop(&h).(*ListNode)
		head.Next = node
		head = head.Next

		// 把下一个放进堆
		if node.Next != nil {
			heap.Push(&h, node.Next)
		}
	}

	return dummy.Next
}

type MinHeap []*ListNode

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h *MinHeap) Push(v any) {
	*h = append(*h, v.(*ListNode))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 32 最长的有效括号

func longestValidParentheses(s string) int {
	n := len(s)
	stack := make([]int, 0)
	stack = append(stack, -1)
	maxLen := 0
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			// 尝试出站
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			}
			currLen := i - stack[len(stack)-1]
			if currLen > maxLen {
				maxLen = currLen
			}
		}
	}
	return maxLen
}

// 45 跳跃游戏2
func jump(nums []int) int {
	end := 0      // 表示当前层能到的最远距离
	farthest := 0 // 扫描当前层能达到的最远距离
	step := 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}

		if i == end {
			step++
			end = farthest
		}
	}

	return step
}

// 160 相交链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]bool)

	for headA != nil {
		m[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := m[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

// 143 重排链表
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// 1. 找中点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 2. 反转后半段
	second := reverse(slow.Next)
	slow.Next = nil // 断开前后

	// 3. 合并
	first := head
	for second != nil {
		tmp1 := first.Next
		tmp2 := second.Next

		first.Next = second
		second.Next = tmp1

		first = tmp1
		second = tmp2
	}
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func reorderList2(head *ListNode) {
	if head == nil {
		return
	}
	nodes := []*ListNode{}
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}

// 912 排序数组
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}
func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	p := rand.Intn(end-start+1) + start
	nums[end], nums[p] = nums[p], nums[end]

	mid := partition(nums, start, end)
	quickSort(nums, start, mid-1)
	quickSort(nums, mid+1, end)
}
func partition(nums []int, start, end int) int {
	pivot := nums[end]
	left, right := start, start
	for right < end {
		if nums[right] < pivot {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
	nums[left], nums[end] = nums[end], nums[left]
	return left
}

// 归并排序
func sortArray3(nums []int) []int {
	n := len(nums)

	if n <= 1 {
		return nums
	}

	mid := n / 2
	left := sortArray(nums[:mid])
	right := sortArray(nums[mid:])

	return merge3(left, right)
}

func merge3(left, right []int) []int {
	m, n := len(left), len(right)
	res := make([]int, 0, m+n)
	i, j := 0, 0
	for i < m && j < n {
		if left[i] <= right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	res = append(res, left[i:]...)
	res = append(res, right[j:]...)

	return res
}

// 139 单词拆分
func wordBreak(s string, wordDict []string) bool {
	// 前i个字符能否被前j个字典拼成
	m := len(s)
	set := make(map[string]bool)
	for _, v := range wordDict {
		set[v] = true
	}
	dp := make([]bool, m+1)
	dp[0] = true

	for i := 1; i < m; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && set[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[m]

}
