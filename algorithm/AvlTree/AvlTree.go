package main

import (
	"errors"
	"fmt"
)

// 树的三种遍历
func main() {
	// 简化版树结构
	avlTree := initNode(3)
	avlTree.Insert(2)
	avlTree.Insert(1)
	avlTree.Insert(4)
	avlTree.Insert(5)
	avlTree.Insert(6)
	avlTree.Insert(7)
	avlTree.Insert(10)
	avlTree.Insert(9)
	avlTree.Insert(8)
	fmt.Println("中序遍历二叉排序树：")
	avlTree.Traverse()
	fmt.Println()
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
		// 对于当前节点来说 走的左子树
		if value < node.Data {
			// 每次递归返回的都是根节点
			node.Left = node.Left.Insert(value)
			// 查看平衡因子，相当于检查了每个节点的平衡因子
			bf := node.BalanceFactor()
			// 取等于2的 就是取最小不平衡子树
			if bf == 2 {
				// 在最小不平衡子树判断是什么类型的
				// 因为value < node.Data 说明先走的左子树 然后这里判断走的左还是右
				if value < node.Left.Data { // LL型
					newTreeNode = RightRotate(node)
				} else { //LR型
					newTreeNode = LeftRightRotate(node)
				}
			}
		} else if value > node.Data {
			// 每次递归返回的都是根节点
			node.Right = node.Right.Insert(value)
			// 查看平衡因子，相当于检查了每个节点的平衡因子
			bf := node.BalanceFactor()
			// 取等于-2的 就是取最小不平衡子树
			if bf == -2 {
				// 在最小不平衡子树判断是什么类型的
				// 因为value > node.Data 说明先走的右子树 然后这里判断走的左还是右
				if value > node.Right.Data { // LL型
					newTreeNode = LeftRotate(node)
				} else { //LR型
					newTreeNode = RightLeftRotate(node)
				}
			}
		}
	}
	// 返回的是根节点
	// 没有发生平衡被破坏直接返回根节点就行
	if newTreeNode == nil {
		node.UpdateHeight()
		return node
	} else {
		newTreeNode.UpdateHeight()
		return newTreeNode
	}
}

func RightRotate(node *TreeNode) *TreeNode {
	// 左子树作为根节点
	rootNode := node.Left
	var nodeRight *TreeNode
	if rootNode != nil {
		nodeRight = rootNode.Right
	}
	// 根节点作为左子树的右子树
	rootNode.Right = node
	// 左子树的右子树 挂靠到根节点的左子树
	node.Left = nodeRight
	// 更新变化了的节点高度
	node.UpdateHeight()
	rootNode.UpdateHeight()

	return rootNode
}

func LeftRotate(node *TreeNode) *TreeNode {
	rootNode := node.Right
	nodeLeft := rootNode.Left
	rootNode.Left = node
	node.Right = nodeLeft

	node.UpdateHeight()
	rootNode.UpdateHeight()

	return rootNode
}

func RightLeftRotate(node *TreeNode) *TreeNode {
	node.Right = RightRotate(node.Right)
	return LeftRotate(node)
}
func LeftRightRotate(node *TreeNode) *TreeNode {
	node.Left = LeftRotate(node.Left)
	return RightRotate(node)
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
