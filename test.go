package main

import "fmt"

func main() {
	link := NewLinkNode()
	link.head.next = &Node{
		data: 1,
		next: nil,
	}

	link.AddNode(&Node{
		data: 2,
		next: nil,
	})

	link.AddNode(&Node{
		data: 3,
		next: nil,
	})

	head := link.head
	for head.next != nil {
		head = head.next
		fmt.Println(head.data)
	}

	link.ReverseV2()

	head = link.head
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}

}

// 1 - 2 - 3

type Node struct {
	data any
	next *Node
}

type LinkNode struct {
	head *Node
}

func NewLinkNode() *LinkNode {
	return &LinkNode{head: &Node{}}
}

func (receiver LinkNode) AddNode(node *Node) error {
	head := receiver.head
	for head.next != nil {
		head = head.next
	}
	head.next = node
	return nil
}

func (receiver *LinkNode) Reverse() error {
	head := receiver.head
	data := make([]any, 0)
	for head.next != nil {
		head = head.next
		data = append(data, head.data)
	}
	node := &Node{
		data: data[len(data)-1],
	}
	receiver.head.next = node
	for j := len(data) - 2; j >= 0; j-- {
		newNode := &Node{
			data: data[j],
			next: nil,
		}
		node.next = newNode
		node = newNode
	}
	return nil
}

func (receiver *LinkNode) ReverseV2() error {
	current := receiver.head
	var pre *Node
	for current != nil {
		next := current.next
		current.next = pre
		pre = current
		current = next
	}
	receiver.head = pre
	return nil
}

// nil -> 1 -> 2 -> 3
// node.next.next = nil
// node.next.next = nil
