// 剑指 Offer 24. 反转链表
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法1：遍历
// Time: O(n)
// Space: O(1)
func reverseList1(head *ListNode) *ListNode {
	var prev, next *ListNode
	for head != nil {
		// 将当前结点的下一个结点保存起来, 并将Next指针指向 prev
		next, head.Next = head.Next, prev
		// 指向下移
		prev, head = head, next
	}
	return prev
}

// 方法二：递归
// Time: O(n)
// Space: O(n)，递归栈空间
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 递归找到最后一个结点，后面会转变为头结点
	prev := reverseList(head.Next)
	// 将当前的下一个结点的Next指针指向当前结点，并清空当前结点的Next，即:1->1 转换成 2->1
	head.Next.Next,head.Next = head, nil
	return prev
}
