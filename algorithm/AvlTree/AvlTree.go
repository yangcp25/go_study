package main

import (
	"errors"
	"fmt"
	"math"
)

// 树的三种遍历
func main() {
	// 简化版树结构
	tree := initNode(nil)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(7)
	tree.Insert(6)
	fmt.Println("中序遍历二叉排序树：")
	midOrderTraverse(tree.rootNode)
	fmt.Println()
}

func midOrderTraverse(t *TreeNode) {
	if t == nil {
		return
	}
	midOrderTraverse(t.Left)
	t.printNode()
	midOrderTraverse(t.Right)
}

// Insert 二叉排序树插入
func (t *Tree) Insert(value int) {
	// 如果是空树先初始化
	t.rootNode = t.rootNode.Insert(value)
}

// Insert 二叉排序树插入
func (node *TreeNode) Insert(value int) *TreeNode {
	// 如果是空树先初始化
	var newTreeNode *TreeNode
	// 当前子树为空 可以插入
	if node == nil {
		return &TreeNode{
			Data:   value,
			Height: 1,
		}
	} else if node.Data == value {
		return node
	} else {
		if value < node.Data {
			node.Left = node.Left.Insert(value)
			// 查看平衡因子
			bf := node.BalanceFactor()
			bf = int(math.Abs(float64(bf)))
			if bf > 2 {
			}
		} else if value > node.Data {
			// 当前右子树为空 可以插入
			if node.Right == nil {
				node.Right = newNode(value)
			}
			// 如果当前值比树节点的值大 则继续在右子树查找
			return node.Right.Insert(value)
		}
	}
	// 返回的是根节点
	// 没有发生平衡被破坏直接返回根节点就行
	if newTreeNode == nil {
		return node
	} else {
		return newTreeNode
	}
}

// 二叉排序树查找
func (t *Tree) find(value int) any {
	if t.rootNode == nil {
		return errors.New("空树！")
	}
	// 如果是空树先初始化
	return t.rootNode.find(value)
}

// 查找
func (node *TreeNode) find(value int) any {
	if value == node.Data {
		// 如果当前值比树节点的值大 则继续在右子树查找
		return node
	} else if value < node.Data {
		// 如果当前值比树节点的值小 则继续在左子树查找
		if node.Left == nil {
			return nil
		}
		return node.Left.find(value)
	} else if value > node.Data {
		// 如果当前值比树节点的值大 则继续在右子树查找
		if node.Right == nil {
			return nil
		}
		return node.Right.find(value)
	} else { // 否则报差错
		return nil
	}
}
