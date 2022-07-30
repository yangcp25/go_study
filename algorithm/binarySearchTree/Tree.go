package main

import "fmt"

// 节点
type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

// 增加一个节点
func newNode(data int) *TreeNode {
	return &TreeNode{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

func (node *TreeNode) printNode() {
	fmt.Printf("%v", node.Data)
}

// 定义树，保存根节点就可以

type Tree struct {
	rootNode *TreeNode
}

func initNode(node *TreeNode) *Tree {
	return &Tree{
		node,
	}
}
