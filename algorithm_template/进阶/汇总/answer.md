当然可以 💪
下面是一份为 Go 后端工程师准备的 **《高阶数据结构算法模板.md》**，精选了在百度、字节、美团、阿里等社招常考的高阶数据结构实现模板。
内容覆盖：

* LRU 缓存
* 跳表 SkipList
* Trie 前缀树
* 并查集 Union-Find
* 堆 Heap（优先队列）
* 线段树 Segment Tree
* TopK（分块 + 小根堆）

---

# 高阶数据结构算法模板（Go语言）

## 🧠 1. LRU 缓存

```go
package main

import "container/list"

type LRUCache struct {
    capacity int
    cache    map[int]*list.Element
    list     *list.List
}

type entry struct {
    key, value int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
        capacity: capacity,
        cache:    make(map[int]*list.Element),
        list:     list.New(),
    }
}

func (l *LRUCache) Get(key int) int {
    if e, ok := l.cache[key]; ok {
        l.list.MoveToFront(e)
        return e.Value.(entry).value
    }
    return -1
}

func (l *LRUCache) Put(key, value int) {
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
```

---

## 🪜 2. 跳表（SkipList）

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

const (
    MaxLevel = 16
    P        = 0.25
)

type node struct {
    val     int
    forward []*node
}

type skipList struct {
    head  *node
    level int
}

func newNode(val, level int) *node {
    return &node{val: val, forward: make([]*node, level)}
}

func newSkipList() *skipList {
    rand.Seed(time.Now().UnixNano())
    return &skipList{head: newNode(-1, MaxLevel), level: 1}
}

func randomLevel() int {
    level := 1
    for rand.Float64() < P && level < MaxLevel {
        level++
    }
    return level
}

func (sl *skipList) Search(target int) bool {
    cur := sl.head
    for i := sl.level - 1; i >= 0; i-- {
        for cur.forward[i] != nil && cur.forward[i].val < target {
            cur = cur.forward[i]
        }
    }
    cur = cur.forward[0]
    return cur != nil && cur.val == target
}

func (sl *skipList) Insert(val int) {
    update := make([]*node, MaxLevel)
    cur := sl.head
    for i := sl.level - 1; i >= 0; i-- {
        for cur.forward[i] != nil && cur.forward[i].val < val {
            cur = cur.forward[i]
        }
        update[i] = cur
    }
    level := randomLevel()
    if level > sl.level {
        for i := sl.level; i < level; i++ {
            update[i] = sl.head
        }
        sl.level = level
    }
    newNode := newNode(val, level)
    for i := 0; i < level; i++ {
        newNode.forward[i] = update[i].forward[i]
        update[i].forward[i] = newNode
    }
}

func (sl *skipList) Delete(val int) bool {
    update := make([]*node, MaxLevel)
    cur := sl.head
    for i := sl.level - 1; i >= 0; i-- {
        for cur.forward[i] != nil && cur.forward[i].val < val {
            cur = cur.forward[i]
        }
        update[i] = cur
    }
    cur = cur.forward[0]
    if cur == nil || cur.val != val {
        return false
    }
    for i := 0; i < sl.level; i++ {
        if update[i].forward[i] != cur {
            break
        }
        update[i].forward[i] = cur.forward[i]
    }
    for sl.level > 1 && sl.head.forward[sl.level-1] == nil {
        sl.level--
    }
    return true
}
```

---

## 🌲 3. Trie 前缀树

```go
type TrieNode struct {
    children map[rune]*TrieNode
    isEnd    bool
}

type Trie struct {
    root *TrieNode
}

func Constructor() Trie {
    return Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
    node := t.root
    for _, ch := range word {
        if node.children[ch] == nil {
            node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
        }
        node = node.children[ch]
    }
    node.isEnd = true
}

func (t *Trie) Search(word string) bool {
    node := t.root
    for _, ch := range word {
        if node.children[ch] == nil {
            return false
        }
        node = node.children[ch]
    }
    return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
    node := t.root
    for _, ch := range prefix {
        if node.children[ch] == nil {
            return false
        }
        node = node.children[ch]
    }
    return true
}
```

---

## ⚙️ 4. 并查集（Union-Find）

```go
type UnionFind struct {
    parent, rank []int
}

func NewUnionFind(n int) *UnionFind {
    p := make([]int, n)
    r := make([]int, n)
    for i := range p {
        p[i] = i
    }
    return &UnionFind{p, r}
}

func (uf *UnionFind) Find(x int) int {
    if uf.parent[x] != x {
        uf.parent[x] = uf.Find(uf.parent[x])
    }
    return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
    rootX, rootY := uf.Find(x), uf.Find(y)
    if rootX == rootY {
        return
    }
    if uf.rank[rootX] < uf.rank[rootY] {
        uf.parent[rootX] = rootY
    } else if uf.rank[rootX] > uf.rank[rootY] {
        uf.parent[rootY] = rootX
    } else {
        uf.parent[rootY] = rootX
        uf.rank[rootX]++
    }
}
```

---

## 🧩 5. 堆（最小堆 / 最大堆）

```go
import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // 小根堆
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}
```

---

## 🪜 6. 线段树（Segment Tree）

```go
type SegmentTree struct {
    tree, nums []int
}

func NewSegmentTree(nums []int) *SegmentTree {
    st := &SegmentTree{nums: nums, tree: make([]int, 4*len(nums))}
    st.build(0, 0, len(nums)-1)
    return st
}

func (st *SegmentTree) build(node, start, end int) {
    if start == end {
        st.tree[node] = st.nums[start]
        return
    }
    mid := (start + end) / 2
    st.build(2*node+1, start, mid)
    st.build(2*node+2, mid+1, end)
    st.tree[node] = st.tree[2*node+1] + st.tree[2*node+2]
}

func (st *SegmentTree) Query(node, start, end, L, R int) int {
    if R < start || L > end {
        return 0
    }
    if L <= start && end <= R {
        return st.tree[node]
    }
    mid := (start + end) / 2
    return st.Query(2*node+1, start, mid, L, R) + st.Query(2*node+2, mid+1, end, L, R)
}
```

---

## 💎 7. 分块 + 小根堆求前 K 大数

```go
import (
    "container/heap"
)

func TopK(nums []int, k int) []int {
    h := &IntHeap{}
    heap.Init(h)
    for _, num := range nums {
        if h.Len() < k {
            heap.Push(h, num)
        } else if num > (*h)[0] {
            heap.Pop(h)
            heap.Push(h, num)
        }
    }
    return *h
}
```
