//143. 重排链表
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 解题思路：
//		1：先找到中间结点，然后分为前后两段
//		2：后半段反转
//      3：再将前半和后半依次交错合并
// Time: O(n)
// Space:O(1)
func reorderList(head *ListNode) {
	// 寻找中间结点，并分段
	mid := findMid(head)
	next := mid.Next
	mid.Next = nil
	// 反转后半段
	next = reverse(next)
	// 合并
	merge(head, next)
}

func findMid(head *ListNode) *ListNode {
	fast := head
	for fast != nil && fast.Next != nil {
		head = head.Next
		fast = fast.Next.Next
	}
	return head
}

func reverse(head *ListNode) *ListNode {
	var prev, next *ListNode
	for head != nil {
		next, head.Next = head.Next, prev
		prev, head = head, next
	}
	return prev
}

func merge(first, second *ListNode) {
	for first != nil && second != nil {
		next1, next2 := first.Next, second.Next
		first.Next, second.Next = second, next1
		first, second = next1, next2
	}
}
