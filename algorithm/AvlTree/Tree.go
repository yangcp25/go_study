package main

import "fmt"

// TreeNode 节点
type TreeNode struct {
	Data        int
	Left, Right *TreeNode
	Height      int
}

// Traverse 遍历
func (tree *Tree) Traverse() {
	tree.rootNode.midOrderTraverse()
}

// 递归方式中序遍历树
func (node *TreeNode) midOrderTraverse() {
	if node == nil {
		return
	}
	node.Left.midOrderTraverse()
	node.printNode()
	node.Right.midOrderTraverse()
}

// BalanceFactor 计算节点的平衡因子
func (node *TreeNode) BalanceFactor() int {
	leftHeight, rightHeight := 0, 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}
	if node.Right != nil {
		rightHeight = node.Right.Height
	}

	return leftHeight - rightHeight
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
	fmt.Printf("节点%v;平衡因子:%d\n", node.Data, node.BalanceFactor())
}

func (node *TreeNode) UpdateHeight() {
	if node == nil {
		return
	}

	leftHeight, rightHeight := 0, 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}
	if node.Right != nil {
		rightHeight = node.Right.Height
	}

	maxHeight := leftHeight
	if leftHeight < rightHeight {
		maxHeight = rightHeight
	}

	node.Height = maxHeight + 1
}

// 定义树，保存根节点就可以

type Tree struct {
	rootNode *TreeNode
}

func initNode(val int) *Tree {
	return &Tree{
		&TreeNode{
			Data:   val,
			Height: 1,
		},
	}
}
