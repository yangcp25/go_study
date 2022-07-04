package main

import "fmt"

type queueNode struct {
	Data any
	Next *queueNode
}

// Stack 使用链表实现队结构
type Stack struct {
	Length int
	Top    *queueNode
}

func buildStack() *Stack {
	top := &queueNode{
		Data: nil,
		Next: nil,
	}
	return &Stack{
		0,
		top,
	}
}

func newNode(v any, next *queueNode) *queueNode {
	return &queueNode{v, next}
}

// 实现队 ，先进先出
// 入队
func (stack *Stack) Push(v any) {
	top := stack.Top
	for top.Next != nil {
		top = top.Next
	}
	if top.Next == nil {
		newNodeData := newNode(v, nil)
		top.Next = newNodeData
	}
}

// 出队
func (stack *Stack) Pop() {
	if stack.Top.Next != nil {
		stack.Top = stack.Top.Next
		fmt.Printf("%v\n", stack.Top.Data)
	}
}

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
