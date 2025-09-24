package main

import "fmt"

func main() {
	//

	head := &ListNode{
		Val:  1,
		Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}},
	}

	head2 := &ListNode{
		Val:  3,
		Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}},
	}

	newLink := sortLinkV2(head, head2)

	for newLink != nil {
		fmt.Println(newLink.Val)
		newLink = newLink.Next
	}
}

func sortLink(list1 *ListNode, list2 *ListNode) *ListNode {
	newLink := &ListNode{}
	head := newLink
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			newLink.Next = list1
			list1 = list1.Next
		} else {
			newLink.Next = list2
			list2 = list2.Next
		}
		newLink = newLink.Next
	}
	if list1 != nil {
		newLink.Next = list1
	}

	if list2 != nil {
		newLink.Next = list2
	}
	return head
}

func sortLinkV2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val <= list2.Val {
		list1.Next = sortLinkV2(list1.Next, list2)
		return list1
	} else {
		list2.Next = sortLinkV2(list1, list2.Next)
		return list2
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
