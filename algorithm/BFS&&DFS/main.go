package fs

import "fmt"

func main() {
	//     1
	//   2   3
	// 5    4
	root := &TreeNode{
		Val: 1,
	}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}

	root.Right.Left = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 5}
	list := make([]int, 0)
	TreeDFS(root, &list)
	fmt.Println(list)
	listMid := make([]int, 0)
	TreeDFSMid(root, &listMid)
	fmt.Println(listMid)
	list2 := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		count := len(queue)
		for i := 0; i < count; i++ {
			node := queue[0]
			queue = queue[1:]
			list2 = append(list2, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	fmt.Println(list2)

	// 通过 list listMid 构造二叉树
	// [1 2 5 3 4] 先序
	// [2 5 1 3 4] 中序
	// 1 [2 5] [2 5] i = 2 [3 4]
	// [2 5 3 4] [2 5]
	// [5 3 4] [5]

	copyTreeRoot := BuildTree(list, listMid)
	fmt.Println(copyTreeRoot)
	list3 := make([]int, 0)
	TreeDFS(copyTreeRoot, &list3)
	fmt.Println("list3:", list3)

	//SliceTest()

	//
}

func BuildTree(first []int, mid []int) *TreeNode {
	if len(first) == 0 {
		return nil
	}
	root := &TreeNode{Val: first[0]}
	node := first[0]
	i := 0
	for ; i < len(mid); i++ {
		if mid[i] == node {
			break
		}
	}
	root.Left = BuildTree(first[1:i+1], mid[:i])
	root.Right = BuildTree(first[i+1:], mid[i+1:])
	return root
}

func TreeDFS(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	*list = append(*list, root.Val)
	TreeDFS(root.Left, list)
	TreeDFS(root.Right, list)
}
func TreeDFSMid(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	TreeDFS(root.Left, list)
	*list = append(*list, root.Val)
	TreeDFS(root.Right, list)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SliceTest() {
	slice1 := []int{1, 2, 3, 4, 5}

	fmt.Println(slice1[0:1])
}
