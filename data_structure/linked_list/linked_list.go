package main

import "fmt"

// 定义node
type node struct {
	data interface{}
	next *node
}

// 定义链表
type linkedList struct {
	// 头结点
	nodeHead *node
}

// 定义接口功能
type linkedListFunc interface {
	// 创建链表，如果没有头结点 创建头结点
	createLinkedList()
	// 查找节点 <因为添加删除等都需要先查找节点>
	searchNode(node *node, searchVal int)
	// 添加节点
	createNode(node *node, v interface{})
	// 删除节点
	removeNode(node *node, v interface{})
	// 修改节点
	updateNode(node *node, oldVal interface{}, newVal interface{})
	// 遍历链表
	rangeNode(node *node)
}

func createLinkedList() *linkedList {
	// 头结点 data一般为空 ，或者拿来存链表长度
	headerPoint := &node{
		data: 0,
		next: nil,
	}
	linkedLists := &linkedList{headerPoint}
	return linkedLists
}

func createNode(v interface{}, nextNode *node) *node {
	return &node{
		v,
		nextNode,
	}
}

// 最简单的末尾加节点
func (linkedList *linkedList) insertNode(v interface{}) bool {
	// 指定头结点
	nodeData := linkedList.nodeHead
	// 拿到末尾节点的指针
	for nodeData.next != nil {
		nodeData = nodeData.next
	}
	if nodeData.next == nil {
		newNode := createNode(v, nil)
		nodeData.next = newNode
		return true
	} else {
		return false
	}
}

// 指定值后加节点
func (linkedList *linkedList) insertNodeByValue(v interface{}, searchVal interface{}) bool {
	// 指定头结点
	nodeData := linkedList.nodeHead
	// 拿到指针
	for nodeData.next != nil {
		if nodeData.next.data == searchVal {
			newNode := createNode(v, nodeData.next.next)
			nodeData.next.next = newNode
			return true
		}
		nodeData = nodeData.next
	}
	return false
}

// 删除
func (linkedList *linkedList) removeNode(v interface{}) bool {
	// 指定头结点
	nodeData := linkedList.nodeHead
	for nodeData.next != nil {
		fmt.Printf("%d,", nodeData.next.data)
		fmt.Printf("%d\n", nodeData.next.next)
		if nodeData.next.data == v {
			nodeData.next = nodeData.next.next
			break
		}
		nodeData = nodeData.next
	}
	return true
}

// 遍历
func (linkedList *linkedList) rangeNode() {
	// 指定头结点
	nodeData := linkedList.nodeHead
	for nodeData.next != nil {
		// 排除头结点 直接使用nodeData.next.data
		fmt.Printf("%v", nodeData.next.data)
		nodeData = nodeData.next
	}
}

// 修改
func (linkedList *linkedList) updateNode(oldValue interface{}, newValue interface{}) bool {
	// 指定头结点
	nodeData := linkedList.nodeHead
	for nodeData.next != nil {
		if nodeData.next.data == oldValue {
			nodeData.next.data = newValue
		}
		nodeData = nodeData.next
	}
	return true
}

func main() {
	// 测试
	linkedListOne := createLinkedList()
	linkedListOne.insertNode(1)
	linkedListOne.insertNode(2)
	linkedListOne.insertNode(3)

	linkedListOne.insertNodeByValue(9, 1)

	linkedListOne.rangeNode()
}
