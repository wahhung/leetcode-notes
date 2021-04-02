package main

//剑指 Offer 22. 链表中倒数第k个节点
type ListNode struct {
	Val int
	Next *ListNode
}

// 解法：统计
// 解题思路：先统计，返回第N-k个结点
// Time:O(n) Space:O(1)
func getKthFromEnd(head *ListNode, k int) *ListNode {
	var n int
	cur := head
	for cur != nil {
		cur = cur.Next
		n++
	}
	n -= k
	for head != nil {
		if n == 0 {
			return head
		}
		head = head.Next
		n--
	}
	return head
}

// 解题方法：快慢指针
// 解题思路：
//		1：快指针从head开始，慢指针从root(root->head)开始
//		2：快指针从head先走 n 步
//		3：快慢指针同时走，快指针走完，慢指针停留在 n - k + 1(倒数 k+1 位)
//		4：返回 root.Next
// Time:O(n) Space:O(1)
func getKthFromEnd1(head *ListNode, k int) *ListNode {
	root := &ListNode{Next: head}
	for i := 0; i < k && head != nil; i++ {
		head = head.Next
	}
	for head != nil {
		head = head.Next
		root = root.Next
	}
	return root.Next
}

