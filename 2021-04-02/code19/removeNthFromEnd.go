package code19


type ListNode struct {
	Val int
	Next *ListNode
}

// 19. 删除链表的倒数第 N 个结点
// 解题方法：快慢指针(和剑指 Offer 22. 链表中倒数第k个节点类似)
// 解题思路：
//		1：快指针从head开始，慢指针从cur(root->head, cur=root)开始
//		2：快指针从head先走 n 步
//		3：快慢指针同时走，快指针走完，慢指针停留在 n - k + 1(倒数 k+1 位)
//		4：返回 root.Next
// Time:O(n) Space:O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	root := &ListNode{Next: head}
	cur := root
	for i := 0; i < n && head != nil; i++ {
		head = head.Next
	}
	for head != nil {
		head = head.Next
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return root.Next
}
