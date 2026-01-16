package main

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 206. 反转链表（Reverse Linked List）
// 1 2 3 nil
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func reverseListRec(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListRec(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// 141. 环形链表
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head.Next, head
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}

// 21. 合并两个有序链表
//func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//
//}

// 19. 删除链表倒数第 N 个节点
/**
 * Definition for a singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	slow, fast := dummy, dummy
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

// 25. K 个一组反转链表
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val  int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 检查链表长度是否 >= k
	cur := head
	count := 0
	for cur != nil && count < k {
		cur = cur.Next
		count++
	}
	if count < k {
		// 剩余节点不足 k 个，不反转
		return head
	}

	// 反转前 k 个节点
	var prev *ListNode
	cur = head
	for i := 0; i < k; i++ {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	// 递归反转后续节点
	head.Next = reverseKGroup(cur, k)

	// prev 是反转后的新头节点
	return prev
}

// reverseKGroup 迭代法（循环版）
func reverseKGroup2(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	end := dummy

	for {
		// 找到每一组的末尾节点
		count := 0
		for count < k && end != nil {
			end = end.Next
			count++
		}
		// 如果不足 k 个，则不反转
		if end == nil {
			break
		}

		// 记录下一组的起始节点
		start := pre.Next
		nextGroup := end.Next
		end.Next = nil // 断开本组链表

		// 反转当前这组
		pre.Next = reverse(start)

		// 连接反转后的尾部到下一组
		start.Next = nextGroup

		// 移动指针，准备下一轮
		pre = start
		end = pre
	}

	return dummy.Next
}

// 反转链表（标准迭代写法）
func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
