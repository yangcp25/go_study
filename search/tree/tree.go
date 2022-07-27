package main

import "fmt"

// 查找算法
// 字符串查找tree树
func main() {
	s := []string{"杨春坪", "test", "go", "laravel"}
	p := "杨春坪"

	tree := initTree()

	for _, word := range s {
		tree.insert(word)
	}

	test1 := tree.find(p)

	fmt.Printf("%v", test1)

	test2 := tree.find("sssss")
	fmt.Printf("%v", test2)
}

// Node 定义节点
type Node struct {
	data     string
	isEnd    bool
	children map[rune]*Node
}

// Tree 树定义 保存根节点
type Tree struct {
	// 当前节点
	root *Node
}

//创建一个节点
func createNode(data string) *Node {
	return &Node{
		data:     data,
		isEnd:    false,
		children: make(map[rune]*Node),
	}
}

// 初始化tree树
func initTree() *Tree {
	node := createNode("/")
	return &Tree{node}
}

// 插入单词
func (t *Tree) insert(str string) {
	root := t.root
	for _, code := range str {
		val, ok := root.children[code]
		if !ok {
			val = createNode(string(code))
			root.children[code] = val
		}
		root = val
	}
	root.isEnd = true
}

// 匹配单词
func (t *Tree) find(str string) bool {
	root := t.root
	for _, code := range str {
		val, ok := root.children[code]
		if !ok {
			return false
		}
		root = val
	}

	// 不是完全匹配只是子集
	if root.isEnd != true {
		return false
	}
	return true
}
