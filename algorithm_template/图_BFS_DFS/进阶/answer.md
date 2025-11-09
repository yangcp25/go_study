**“图论”不是只考岛屿数量（DFS/BFS），而是整个算法体系中的重头戏之一。**
在面试中，图题一般占 10%~15%，但出现时往往是「高质量面试题」，考察逻辑 + DFS/BFS 模型。

---

## 🧭 一、面试中最常考的图论题型（按难度分类）

| 分类                    | 常见题目                              | 技巧/模板           |
| --------------------- | --------------------------------- | --------------- |
| ✅ **连通性问题**           | 200. 岛屿数量、695. 岛屿的最大面积、547. 省份数量  | DFS/BFS 遍历模板    |
| ✅ **拓扑排序**            | 207. 课程表、210. 课程表 II              | BFS + 入度表模板     |
| ✅ **环检测（有向图/无向图）**    | 802. 找到最终安全状态、261. 无向图是否是树        | DFS 递归状态（灰白标记法） |
| ✅ **最短路径**            | 743. 网络延迟时间、1631. 最小体力消耗路径        | Dijkstra / BFS  |
| ✅ **并查集（Union-Find）** | 200. 岛屿数量（进阶）、684. 冗余连接、547. 省份数量 | Union-Find 模板   |
| ✅ **图的克隆/遍历**         | 133. 克隆图                          | DFS / BFS 构建新节点 |
| ✅ **图的拓展搜索**          | 127. 单词接龙、815. 公交路线               | BFS 多层次建图       |

---

## ⚙️ 二、核心模板代码（Go）

### 🧩 1. DFS 遍历（岛屿数量类）

```go
func numIslands(grid [][]byte) int {
    m, n := len(grid), len(grid[0])
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    var dfs func(int, int)
    dfs = func(i, j int) {
        if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
            return
        }
        grid[i][j] = '0'
        for _, d := range dirs {
            dfs(i+d[0], j+d[1])
        }
    }

    count := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                count++
                dfs(i, j)
            }
        }
    }
    return count
}
```

---

### 🚦 2. BFS 遍历（课程表 / 层级搜索）

```go
func canFinish(numCourses int, prerequisites [][]int) bool {
    indeg := make([]int, numCourses)
    graph := make([][]int, numCourses)
    for _, p := range prerequisites {
        a, b := p[0], p[1]
        graph[b] = append(graph[b], a)
        indeg[a]++
    }

    q := []int{}
    for i := 0; i < numCourses; i++ {
        if indeg[i] == 0 {
            q = append(q, i)
        }
    }

    count := 0
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        count++
        for _, nxt := range graph[cur] {
            indeg[nxt]--
            if indeg[nxt] == 0 {
                q = append(q, nxt)
            }
        }
    }

    return count == numCourses
}
```

---

### 🔗 3. 并查集（Union-Find 模板）

```go
type UnionFind struct {
    parent []int
    rank   []int
}

func NewUF(n int) *UnionFind {
    uf := &UnionFind{
        parent: make([]int, n),
        rank:   make([]int, n),
    }
    for i := 0; i < n; i++ {
        uf.parent[i] = i
    }
    return uf
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

### 🧭 4. 拓扑排序（课程表 II）

```go
func findOrder(numCourses int, prerequisites [][]int) []int {
    indeg := make([]int, numCourses)
    graph := make([][]int, numCourses)
    for _, p := range prerequisites {
        a, b := p[0], p[1]
        graph[b] = append(graph[b], a)
        indeg[a]++
    }

    q := []int{}
    for i := 0; i < numCourses; i++ {
        if indeg[i] == 0 {
            q = append(q, i)
        }
    }

    res := []int{}
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        res = append(res, cur)
        for _, nxt := range graph[cur] {
            indeg[nxt]--
            if indeg[nxt] == 0 {
                q = append(q, nxt)
            }
        }
    }

    if len(res) != numCourses {
        return []int{} // 有环
    }
    return res
}
```

---

### 🧬 5. 图的克隆（133. 克隆图）

```go
type Node struct {
    Val       int
    Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    visited := map[*Node]*Node{}

    var dfs func(*Node) *Node
    dfs = func(n *Node) *Node {
        if clone, ok := visited[n]; ok {
            return clone
        }
        clone := &Node{Val: n.Val}
        visited[n] = clone
        for _, nei := range n.Neighbors {
            clone.Neighbors = append(clone.Neighbors, dfs(nei))
        }
        return clone
    }

    return dfs(node)
}
```
最短路径（Dijkstra
```go
package main

import (
	"container/heap"
	"fmt"
)

type Edge struct {
	to, weight int
}
type Node struct {
	id, dist int
}
type MinPQ []Node

func (pq MinPQ) Len() int            { return len(pq) }
func (pq MinPQ) Less(i, j int) bool  { return pq[i].dist < pq[j].dist }
func (pq MinPQ) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *MinPQ) Push(x interface{}) { *pq = append(*pq, x.(Node)) }
func (pq *MinPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func dijkstraWithPath(n int, graph [][]Edge, start int) ([]int, []int) {
	dist := make([]int, n)
	prev := make([]int, n)
	for i := range dist {
		dist[i] = 1e9
		prev[i] = -1
	}
	dist[start] = 0

	pq := &MinPQ{}
	heap.Init(pq)
	heap.Push(pq, Node{start, 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Node)
		if cur.dist > dist[cur.id] {
			continue
		}
		for _, e := range graph[cur.id] {
			if dist[cur.id]+e.weight < dist[e.to] {
				dist[e.to] = dist[cur.id] + e.weight
				prev[e.to] = cur.id
				heap.Push(pq, Node{e.to, dist[e.to]})
			}
		}
	}
	return dist, prev
}

func getPath(prev []int, start, end int) []int {
	path := []int{}
	for end != -1 {
		path = append([]int{end}, path...)
		end = prev[end]
	}
	if path[0] != start {
		return []int{} // 无路径
	}
	return path
}

func main() {
	// 构建图（邻接表）
	graph := [][]Edge{
		{{1, 4}, {2, 1}}, // 0→1(4), 0→2(1)
		{{3, 1}},         // 1→3(1)
		{{1, 2}, {3, 5}}, // 2→1(2), 2→3(5)
		{},               // 3
	}

	dist, prev := dijkstraWithPath(4, graph, 0)

	end := 3
	fmt.Println("0 到 3 的最短距离:", dist[end])
	fmt.Println("0 到 3 的路径:", getPath(prev, 0, end))
}


```
---

## 🧱 三、建议优先练习的图题列表（Top 面试常考）

| 类型   | 题号  | 名称      | 模板 |
| ---- | --- | ------- | -- |
| DFS  | 200 | 岛屿数量    | ✅  |
| DFS  | 695 | 岛屿的最大面积 | ✅  |
| BFS  | 127 | 单词接龙    | ✅  |
| 拓扑排序 | 207 | 课程表     | ✅  |
| 拓扑排序 | 210 | 课程表 II  | ✅  |
| 并查集  | 547 | 省份数量    | ✅  |
| 并查集  | 684 | 冗余连接    | ✅  |
| 图克隆  | 133 | 克隆图     | ✅  |
| 最短路径 | 743 | 网络延迟时间  | ✅  |

---
