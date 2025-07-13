package main

func main() {
	// 树的深度优先遍历

	node := &TreeNode{
		Val:   0,
		Right: nil,
		Left:  nil,
	}
	node.Left = &TreeNode{
		Val:   1,
		Right: nil,
		Left:  nil,
	}
	node.Right = &TreeNode{
		Val:   2,
		Right: nil,
		Left:  nil,
	}
	node.Left.Right = &TreeNode{
		Val:   3,
		Right: nil,
		Left:  nil,
	}
	TreeDFS(node)
	//TreeBFS(node)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TreeDFS(node *TreeNode) {

}
