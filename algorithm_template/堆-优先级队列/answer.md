
```go
//前 K 个高频元素（Top K Frequent Elements）
package main

import "container/heap"

type Pair struct {
	num  int
	freq int
}

type MinHeap []Pair

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].freq < h[j].freq }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Pair)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	freq := map[int]int{}
	for _, n := range nums {
		freq[n]++
	}

	h := &MinHeap{}
	heap.Init(h)

	for num, f := range freq {
		heap.Push(h, Pair{num, f})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]int, 0, k)
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(Pair).num)
	}
	return res
}

```