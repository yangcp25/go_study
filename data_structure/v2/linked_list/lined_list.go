package linkedlist

import (
	"github.com/google/uuid"
)

type Linked interface {
	Add(data any) error
	Edit(key string, newData any) error
	Delete(key string) error
	Get(data any) (key string, err error)
	Find() (nodeList []*Node[any], err error)
}

type LinkedList struct {
	node *Node[any]
}
type Node[T any] struct {
	key  string
	data T
	next *Node[any]
}

func CreateLinkerList() *LinkedList {
	uuId := uuid.New()
	key := uuId.String()

	tailNode := &Node[any]{
		key:  key,
		data: nil,
		next: nil,
	}
	return &LinkedList{node: tailNode}
}

func (l *LinkedList) Add(data any) error {
	newNode := &Node[any]{
		data: data,
		next: nil,
	}
	tailNode := l.node
	for l.node.next != nil {
		l.node = l.node.next
	}
	l.node.next = newNode
	l.node = tailNode
	return nil
}

func (l *LinkedList) Edit(key string, newData any) error {
	tailNode := l.node
	for l.node.next != nil {
		if l.node.key == key {
			break
		}
	}
	l.node.data = newData
	l.node = tailNode
	return nil
}

func (l *LinkedList) Delete(key string) error {
	tailNode := l.node
	for l.node.next != nil {
		if l.node.key == key {
			if l.node.next.next != nil {
				l.node.next = l.node.next.next
			} else {
				l.node.next = nil
			}
			break
		}
	}
	l.node = tailNode
	return nil
}

func (l *LinkedList) Get(data any) (key string, err error) {
	tailNode := l.node
	for l.node.next != nil {
		if l.node.data == data {
			break
		}
	}
	l.node = tailNode
	return
}

func (l *LinkedList) Find() (nodeList []*Node[any], err error) {
	tailNode := l.node
	for l.node.next != nil {
		nodeList = append(nodeList, l.node)
	}
	l.node = tailNode
	return
}
