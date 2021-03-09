//876. 链表的中间结点
package main

type ListNode struct {
	Val  int
	Next *ListNode
}


// 解题思路：
//		定位中间结点，利用龟兔赛跑赛跑原理(快慢指针)
//		龟走一步兔子走两步，当兔子走完时，龟正在中间
// Time: O(n)
// Space: O(1)
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
