package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println("hello")
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

	res := levelOrder(node)

	fmt.Println(res)

}

func levelOrder(root *TreeNode) (res [][]int) {
	res = make([][]int, 0)
	temp := make([]TreeNode, 0)
	temp = append(temp, *root)
	for len(temp) > 0 {
		//
		count := len(temp)
		for count > 0 {
			val := temp[0]
			temp = temp[1:]
			res = append(res, []int{val.Val})
			if val.Left != nil {
				temp = append(temp, *val.Left)
			}
			if val.Right != nil {
				temp = append(temp, *val.Right)
			}
			count--
		}
	}
	return res
}
