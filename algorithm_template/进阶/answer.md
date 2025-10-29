
## 1️⃣ LRU 缓存（经典面试题：146. LRU缓存机制）

```go
// LRUCache 双向链表 + 哈希表
type Node struct {
    key, val int
    prev, next *Node
}

type LRUCache struct {
    capacity int
    cache map[int]*Node
    head, tail *Node
}

func Constructor(capacity int) LRUCache {
    l := LRUCache{
        capacity: capacity,
        cache: make(map[int]*Node),
        head: &Node{},
        tail: &Node{},
    }
    l.head.next = l.tail
    l.tail.prev = l.head
    return l
}

func (l *LRUCache) Get(key int) int {
    if node, ok := l.cache[key]; ok {
        l.moveToHead(node)
        return node.val
    }
    return -1
}

func (l *LRUCache) Put(key int, value int) {
    if node, ok := l.cache[key]; ok {
        node.val = value
        l.moveToHead(node)
    } else {
        node := &Node{key: key, val: value}
        l.cache[key] = node
        l.addToHead(node)
        if len(l.cache) > l.capacity {
            removed := l.removeTail()
            delete(l.cache, removed.key)
        }
    }
}

// 将节点移动到头部（表示最近使用）
func (l *LRUCache) moveToHead(node *Node) {
    l.removeNode(node)
    l.addToHead(node)
}

func (l *LRUCache) removeNode(node *Node) {
    node.prev.next = node.next
    node.next.prev = node.prev
}

func (l *LRUCache) addToHead(node *Node) {
    node.prev = l.head
    node.next = l.head.next
    l.head.next.prev = node
    l.head.next = node
}

func (l *LRUCache) removeTail() *Node {
    node := l.tail.prev
    l.removeNode(node)
    return node
}
```

> ✅ 面试要点：
>
> * 哈希表 + 双向链表保证 `O(1)` 时间复杂度；
> * 链表顺序表示最近使用情况。

---

## 2️⃣ 分块查询 / 前 K 个最大最小（经典技巧：大数据、流数据 Top K）

**场景：** 给你一个大数组，需要频繁查询某个区间的最大值 / 最小值或 Top-K 元素，可以用**分块 / 堆**技巧。

### 2.1 分块 + 预处理

```go
package main
import "math"

type BlockQuery struct {
    data []int
    blocks []int
    blockSize int
}

func NewBlockQuery(arr []int) *BlockQuery {
    n := len(arr)
    blockSize := int(math.Sqrt(float64(n))) + 1
    blocks := make([]int, (n+blockSize-1)/blockSize)
    for i := range blocks {
        blocks[i] = math.MinInt32
    }
    for i, v := range arr {
        blk := i / blockSize
        if v > blocks[blk] {
            blocks[blk] = v
        }
    }
    return &BlockQuery{data: arr, blocks: blocks, blockSize: blockSize}
}

// 查询区间 [l, r] 的最大值
func (b *BlockQuery) QueryMax(l, r int) int {
    res := math.MinInt32
    for i := l; i <= r; {
        if i%b.blockSize == 0 && i+b.blockSize-1 <= r {
            blk := i / b.blockSize
            if b.blocks[blk] > res {
                res = b.blocks[blk]
            }
            i += b.blockSize
        } else {
            if b.data[i] > res {
                res = b.data[i]
            }
            i++
        }
    }
    return res
}
```

> ✅ 面试要点：
>
> * 分块大小一般取 √n；
> * 块内线性扫描，块间直接取预处理结果；
> * 可以延伸到**区间最小值 / Top K**。

---

### 2.2 Top K（堆实现）

```go
package main
import "container/heap"

// 最小堆获取前 K 大元素
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // 小根堆
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func TopK(nums []int, k int) []int {
    h := &IntHeap{}
    heap.Init(h)
    for _, v := range nums {
        heap.Push(h, v)
        if h.Len() > k {
            heap.Pop(h)
        }
    }
    res := make([]int, h.Len())
    for i := len(res)-1; i >=0 ; i-- {
        res[i] = heap.Pop(h).(int)
    }
    return res
}
```

> ✅ 面试要点：
>
> * 小根堆维护前 K 大，堆顶是最小值；
> * 大根堆同理可以维护 Top K 小；
> * 时间复杂度 `O(n log k)`，空间 `O(k)`。



---


```go
package main

// 547. 省份数量

func findCircleNumDFS(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	var dfs func(int)
	dfs = func(u int) {
		visited[u] = true
		for v := 0; v < n; v++ {
			if isConnected[u][v] == 1 && !visited[v] {
				dfs(v)
			}
		}
	}

	provinces := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			provinces++
			dfs(i)
		}
	}
	return provinces
}

//func main() {
//	mat := [][]int{
//		{1,1,0},
//		{1,1,0},
//		{0,0,1},
//	}
//	fmt.Println(findCircleNumDFS(mat)) // 输出 2
//}


type UnionFind struct {
	parent []int
	rank   []int
	count  int
}

func NewUnionFind(n int) *UnionFind {
	p := make([]int, n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
		r[i] = 1
	}
	return &UnionFind{parent: p, rank: r, count: n}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rx, ry := uf.Find(x), uf.Find(y)
	if rx == ry {
		return
	}
	// 按秩合并
	if uf.rank[rx] < uf.rank[ry] {
		uf.parent[rx] = ry
	} else if uf.rank[rx] > uf.rank[ry] {
		uf.parent[ry] = rx
	} else {
		uf.parent[ry] = rx
		uf.rank[rx]++
	}
	uf.count--
}

func findCircleNumUF(isConnected [][]int) int {
	n := len(isConnected)
	uf := NewUnionFind(n)
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ { // 只遍历上三角，避免重复
			if isConnected[i][j] == 1 {
				uf.Union(i, j)
			}
		}
	}
	// 统计根节点数量（也可以直接用 uf.count）
	return uf.count
}

//func main() {
//	mat := [][]int{
//		{1,1,0},
//		{1,1,0},
//		{0,0,1},
//	}
//	fmt.Println(findCircleNumUF(mat)) // 输出 2
//}

// 133. 克隆图

type Node struct {
	Val int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	// visited 映射原节点 → 新节点（防止重复克隆）
	visited := make(map[*Node]*Node)

	var dfs func(*Node) *Node
	dfs = func(n *Node) *Node {
		if n == nil {
			return nil
		}

		// 如果已克隆过，直接返回副本
		if cloned, ok := visited[n]; ok {
			return cloned
		}

		// 克隆当前节点（但先不克隆邻居）
		clone := &Node{Val: n.Val}
		visited[n] = clone

		// 克隆所有邻居
		for _, nei := range n.Neighbors {
			clone.Neighbors = append(clone.Neighbors, dfs(nei))
		}

		return clone
	}

	return dfs(node)
}


func cloneGraphBFS(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[*Node]*Node)
	queue := []*Node{node}

	// 克隆第一个节点
	visited[node] = &Node{Val: node.Val}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, nei := range curr.Neighbors {
			if _, ok := visited[nei]; !ok {
				// 克隆邻居节点
				visited[nei] = &Node{Val: nei.Val}
				queue = append(queue, nei)
			}
			// 加入邻居关系
			visited[curr].Neighbors = append(visited[curr].Neighbors, visited[nei])
		}
	}

	return visited[node]
}


```
