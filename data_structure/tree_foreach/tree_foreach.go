package main

import "fmt"

// 树的三种遍历
func main() {
	// 简化版树结构
	tree := newNode("A")
	tree1 := newNode("B")
	tree2 := newNode("C")

	tree.Left = tree1
	tree.Right = tree2

	tree.Left.Left = newNode("D")
	tree.Left.Right = newNode("E")

	tree.Right.Left = newNode("F")
	tree.Right.Right = newNode("G")

	// 前序遍历
	preOrderTraverse(tree)
	fmt.Println()
	midOrderTraverse(tree)
	fmt.Println()
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
	midOrderTraverse(t.Left)
	t.printNode()
	midOrderTraverse(t.Right)
}

func lastOrderTraverse(t *Tree) {
	if t == nil {
		return
	}
	lastOrderTraverse(t.Left)
	lastOrderTraverse(t.Right)
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
