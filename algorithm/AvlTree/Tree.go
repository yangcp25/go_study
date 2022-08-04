package main

import (
	"fmt"
	"math"
)

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

func (node *TreeNode) isBalanced() bool {
	if node == nil {
		return true
	}
	if node.Left == nil && node.Right == nil {
		return true
	} else if node.Left != nil && node.Right != nil {
		// 左子树比节点小 右子树比节点大
		if node.Left.Data > node.Data || node.Right.Data < node.Data {
			fmt.Printf("不是排序树！左子树是：%d;节点是：节点%d；右子树是：%d\n", node.Left.Data, node.Data, node.Right.Data)
		}
		// 平衡因子 < 2
		bf := node.BalanceFactor()
		if math.Abs(float64(bf)) > 1 {
			fmt.Printf("节点%d:平衡因子>1！左子树高度是：%d;节点高度是：%d；右子树高度是：%d\n", node.Data, node.Left.Height, node.Height, node.Right.Height)
		}
		// 节点的高度是左右子树最高的高度+1
		if node.Left.Height > node.Right.Height {
			if node.Height != node.Left.Height+1 {
				fmt.Printf("节点%d:节点高度不等于左子树+1,左子树高度是：%d;节点高度是：%d", node.Data, node.Left.Height, node.Height)
			}
		} else {
			if node.Height != node.Right.Height+1 {
				fmt.Printf("节点%d:节点高度不等于右子树+1,右子树高度是：%d;节点高度是：%d", node.Data, node.Right.Height, node.Height)
			}
		}
		// 递归判断左右子树
		if !node.Left.isBalanced() {
			return false
		}
		if !node.Right.isBalanced() {
			return false
		}
	} else {
		// 只有一个节点
		// 只有左子树
		// （1）左子树要比节点小 (2)左子树不能有子树
		if node.Left != nil {
			if node.Left.Data > node.Data || node.Left.Left != nil || node.Left.Right != nil {
				fmt.Printf("节点%d:（1）左子树要比节点小 (2)左子树不能有子树。左子树高度是：%d;节点高度是：%d；右子树高度是：%d\n", node.Data, node.Left.Height, node.Height, node.Right.Height)
			}
		} else { // 只有右子树
			// （1）右子树要比节点Da (2)左子树不能有子树
			if node.Right.Data < node.Data || node.Right.Left != nil || node.Right.Right != nil {
				fmt.Printf("节点%d:（（1）右子树要比节点Da (2)左子树不能有子树。左子树高度是：%d;节点高度是：%d；右子树高度是：%d\n", node.Data, node.Left.Height, node.Height, node.Right.Height)
			}
		}
	}
	return true
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
