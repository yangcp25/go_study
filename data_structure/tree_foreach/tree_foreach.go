package main

import "fmt"

// 树的三种遍历
func main() {
	// 简化版树结构
	tree := newNode(0)
	tree1 := newNode(1)
	tree2 := newNode(2)

	tree.Left = tree1
	tree.Right = tree2

	// 前序遍历
	preOrderTraverse(tree)
	midOrderTraverse(tree)
	lastOrderTraverse(tree)
	// 中序遍历
	// 后序遍历
}

func preOrderTraverse(t *Tree) {
	if t == nil {
		return
	}
	t.printNode()
	preOrderTraverse(t.Left)
	preOrderTraverse(t.Right)
}

func midOrderTraverse(t *Tree) {
	if t == nil {
		return
	}
	preOrderTraverse(t.Left)
	t.printNode()
	preOrderTraverse(t.Right)
}

func lastOrderTraverse(t *Tree) {
	if t == nil {
		return
	}
	preOrderTraverse(t.Left)
	preOrderTraverse(t.Right)
	t.printNode()
}

type Tree struct {
	Data  interface{}
	Left  *Tree
	Right *Tree
}

func newNode(data interface{}) *Tree {
	return &Tree{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

func (t *Tree) printNode() {
	fmt.Printf("%v", t.Data)
}
