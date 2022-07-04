package main

import "fmt"

type StackNode struct {
	Data any
	Next *StackNode
}

// Stack 使用链表实现栈结构
type Stack struct {
	Length int
	Top    *StackNode
}

func buildStack() *Stack {
	top := &StackNode{
		Data: nil,
		Next: nil,
	}
	return &Stack{
		0,
		top,
	}
}

func newNode(v any, next *StackNode) *StackNode {
	return &StackNode{v, next}
}

// 实现栈 ，先进后出
// 入栈
func (stack *Stack) Push(v any) {
	top := stack.Top
	newNodeData := newNode(v, top.Next)
	top.Next = newNodeData
}

// 出栈
func (stack *Stack) Pop() {
	if stack.Top.Next != nil {
		stack.Top = stack.Top.Next
		//fmt.Printf("%v\n", stack.Top.Data)
	}
}

// 遍历
func (stack *Stack) Range() {
	top := stack.Top
	for top.Next != nil {
		fmt.Printf("%v\n", top.Next.Data)
		top = top.Next
	}
}
func main() {
	stack := buildStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	stack.Range()

	stack.Pop()
	stack.Pop()
	stack.Range()
}
