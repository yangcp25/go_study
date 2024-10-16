package main

import (
	"errors"
	"fmt"
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

	find1 := tree.find(8)
	find2 := tree.find(7)
	fmt.Printf("找8：%v\n 找7：%v\n", find1, find2)

	// 测试删除
	tree.delete1(3)
	tree.delete1(7)
	tree.delete1(6)
	fmt.Println("1删除后遍历二叉排序树：")
	midOrderTraverse(tree.rootNode)

	tree2 := initNode(nil)
	tree2.Insert(3)
	tree2.Insert(2)
	tree2.Insert(5)
	tree2.Insert(1)
	tree2.Insert(4)
	tree2.Insert(7)
	tree2.Insert(6)
	tree2.delete2(3)
	tree2.delete1(7)
	tree2.delete1(6)
	//tree2.delete2(7)
	//tree2.delete2(6)
	fmt.Println("\n2删除后遍历二叉排序树：")
	midOrderTraverse(tree.rootNode)
	//fmt.Println("二叉排序树：")
	//fmt.Printf("%v\n", tree)
}

// Insert 二叉排序树插入
func (t *Tree) Insert(value int) any {
	// 如果是空树先初始化
	node := t.rootNode
	if node == nil {
		t.rootNode = newNode(value)
		return nil
	} else {
		for node != nil {
			if value < node.Data {
				// 当前左子树为空 可以插入
				if node.Left == nil {
					node.Left = newNode(value)
				}
				// 如果当前值比树节点的值小 则继续在左子树查找
				node = node.Left
			} else if value > node.Data {
				// 当前右子树为空 可以插入
				if node.Right == nil {
					node.Right = newNode(value)
				}
				// 如果当前值比树节点的值大 则继续在右子树查找
				node = node.Right
			} else { // 否则报差错
				return errors.New("已存在当前节点")
			}
		}

	}
	return nil
}

// 二叉排序树查找
func (t *Tree) find(value int) any {
	// 如果是空树先初始化
	node := t.rootNode
	for node != nil {
		if value < node.Data {
			// 如果当前值比树节点的值小 则继续在左子树查找
			node = node.Left
		} else if value > node.Data {
			// 如果当前值比树节点的值大 则继续在右子树查找
			node = node.Right
		} else { // 否则报差错
			return node
		}
	}
	return errors.New("不存在当前节点")
}

// 二叉排序树删除
func (t *Tree) delete1(deleteElm int) any {
	// 3种情况
	// (1)删除的节点没有子节点
	// (2)删除的节点只有一个节点
	// (3)删除的节点有2个节点
	node := t.rootNode
	var nodeParent *TreeNode = nil
	for node != nil && node.Data != deleteElm {
		nodeParent = node
		if deleteElm < node.Data {
			// 如果当前值比树节点的值小 则继续在左子树查找
			node = node.Left
		} else if deleteElm > node.Data {
			// 如果当前值比树节点的值大 则继续在右子树查找
			node = node.Right
		} else {
			break
		}
	}

	if node == nil {
		return errors.New("删除的节点不存在")
	}

	// 判断当前node 是父树的左子树还是右子树

	if node.Left != nil && node.Right != nil {
		rightNode := node.Right
		rightNodeP := node

		// 找到右子树最小的节点 或者找到左子树最大的节点替代删除的节点并删掉它
		for rightNode.Left != nil {
			rightNodeP = rightNode
			rightNode = rightNode.Left
		}
		rightNodeP.Left = nil
		minRightNode := rightNode

		//替换删除的node
		if node.Left != nil {
			t.rootNode.Data = minRightNode.Data
		} else {
			if nodeParent.Left == node {
				nodeParent.Left.Data = minRightNode.Data
			} else {
				nodeParent.Right.Data = minRightNode.Data
			}
		}

	} else if node.Left != nil || node.Right != nil {
		// 父树指向被删除的节点的 单个子树
		if node.Left != nil {
			if nodeParent.Left == node {
				nodeParent.Left = node.Left
			} else {
				nodeParent.Right = node.Left
			}
		} else {
			if nodeParent.Left == node {
				nodeParent.Left = node.Right
			} else {
				nodeParent.Right = node.Right
			}
		}
	} else {
		// 删除的是头结点
		if nodeParent == nil {
			t.rootNode = nil
		} else {
			if nodeParent.Left == node {
				nodeParent.Left = nil
			} else {
				nodeParent.Right = nil
			}
		}
	}
	return true
}
func (t *Tree) delete2(deleteElm int) any {
	// 3种情况
	// (1)删除的节点没有子节点
	// (2)删除的节点只有一个节点
	// (3)删除的节点有2个节点
	node := t.rootNode
	var nodeParent *TreeNode = nil
	for node != nil && node.Data != deleteElm {
		nodeParent = node
		if deleteElm < node.Data {
			// 如果当前值比树节点的值小 则继续在左子树查找
			node = node.Left
		} else if deleteElm > node.Data {
			// 如果当前值比树节点的值大 则继续在右子树查找
			node = node.Right
		}
	}

	if node == nil {
		return errors.New("删除的节点不存在")
	}

	// 判断当前node 是父树的左子树还是右子树

	if node.Left != nil && node.Right != nil {
		minRightNode := node.Right
		rightNodeP := node

		// 找到右子树最小的节点 或者找到左子树最大的节点替代删除的节点并删掉它
		for minRightNode.Left != nil {
			rightNodeP = minRightNode
			minRightNode = minRightNode.Left
		}
		rightNodeP.Left = nil
		node.Data = minRightNode.Data
	} else {
		// 找到删除节点 的子节点
		var child *TreeNode

		if node.Left != nil {
			child = node.Left
		} else if node.Right != nil {
			child = node.Right
		} else {
			child = nil
		}

		// 找到删除节点的子节点插入的位置
		if nodeParent == nil { // 头结点
			t.rootNode = nil
		} else if nodeParent.Left == node {
			nodeParent.Left = child
		} else if node.Right != node {
			nodeParent.Right = child
		}
	}
	return true
}
