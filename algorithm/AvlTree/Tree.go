package main

import "fmt"

// TreeNode 节点
type TreeNode struct {
	Data        int
	Left, Right *TreeNode
	Height      int
}

// BalanceFactor 计算节点的平衡因子
func (node *TreeNode) BalanceFactor() int {
	leftNode, rightNode := 0, 0
	if node.Left != nil {
		leftNode = node.Left.Height
	}
	if node.Right != nil {
		rightNode = node.Right.Height
	}

	return leftNode - rightNode
}

// 增加一个节点
func newNode(data int) *TreeNode {
	return &TreeNode{
		Data:   data,
		Left:   nil,
		Right:  nil,
		Height: 0,
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
