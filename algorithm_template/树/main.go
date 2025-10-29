package main

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDeep := maxDepth(root.Left)
	rightDeep := maxDepth(root.Right)

	return max(leftDeep, rightDeep) + 1
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	depth := 0
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}
	return depth
}

// 236. 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	return right
}

// 124. 二叉树中的最大路径和
// maxPathSum 返回整棵树的最大路径和
func maxPathSum(root *TreeNode) int {
	// 记录全局最大路径和（初始化为最小值）
	maxSum := math.MinInt32

	// 定义一个内部函数，用于计算“以当前节点为起点”的最大贡献值
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0 // 空节点对路径贡献为 0
		}

		// 递归计算左右子树的最大贡献值
		// 只有当贡献值大于 0 时才选取（小于 0 说明会拖后腿）
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		// 当前节点作为“拐点”时的路径和：
		// 左最大 + 右最大 + 当前节点值
		priceNewPath := node.Val + leftGain + rightGain

		// 更新全局最大路径和（比较是否更优）
		maxSum = max(maxSum, priceNewPath)

		// 返回“单边最大贡献值”给上层父节点
		// 父节点只能选择左或右的一条路径来延伸
		return node.Val + max(leftGain, rightGain)
	}

	// 从根节点开始递归
	maxGain(root)

	// 返回结果
	return maxSum
}

// 工具函数：取最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 测试用例
func main() {
	/*
	       构建如下二叉树：
	           -10
	           /  \
	          9   20
	             /  \
	            15   7
	   预期最大路径和 = 15 + 20 + 7 = 42
	*/
	root := &TreeNode{Val: -10}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	result := maxPathSum(root)
	fmt.Println("最大路径和 =", result) // 输出 42
}

// 98 验证二叉搜索树
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt, math.MaxInt)
}

func helper(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}

	return helper(root.Left, min, root.Val) && helper(root.Right, root.Val, max)
}

func isValidBST2(root *TreeNode) bool {
	var stack []*TreeNode
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}

// 二叉树的前序遍历

// 二叉树的中序遍历
// 二叉树的后序遍历
