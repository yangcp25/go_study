package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MaxLevel = 16   // 最大层级
	P        = 0.25 // 每层概率
)

// 节点结构
type node struct {
	val     int
	forward []*node
}

// 跳表结构
type skipList struct {
	head  *node
	level int
}

// 创建新节点
func newNode(val, level int) *node {
	return &node{
		val:     val,
		forward: make([]*node, level),
	}
}

// 创建跳表
func newSkipList() *skipList {
	rand.Seed(time.Now().UnixNano())
	return &skipList{
		head:  newNode(-1, MaxLevel),
		level: 1,
	}
}

// 随机层数
func randomLevel() int {
	level := 1
	for rand.Float64() < P && level < MaxLevel {
		level++
	}
	return level
}

// 查找
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

// 插入
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

// 删除
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

// 测试用例
func main() {
	sl := newSkipList()
	sl.Insert(3)
	sl.Insert(6)
	sl.Insert(7)
	sl.Insert(9)
	sl.Insert(12)
	sl.Insert(19)
	sl.Insert(17)
	sl.Insert(26)
	sl.Insert(21)
	sl.Insert(25)

	fmt.Println(sl.Search(19)) // true
	fmt.Println(sl.Search(15)) // false

	sl.Delete(19)
	fmt.Println(sl.Search(19)) // false
}
