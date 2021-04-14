 // 19. 删除链表的倒数第 N 个结点
 // 解题思路：快慢指针
 // 1: 快指针从第0位开始，走n步(root.Next = head)
 // 2: 慢指针从第0位开始，快慢指针同时走一步，同时定义prev指针
 // 3: 循环结束条件，快指针 == nil || 快指针.Next == nil, 这时慢指针刚好停在 倒数n+1个点上
 // 4: 返回 root.Next
 // Time:O(n) Space:O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    root := &ListNode{Next:head}
    fast, slow := root, root
    for i := 0; i < n; i++ {
        fast = fast.Next
    }
    for fast != nil && fast.Next != nil {
        fast = fast.Next
        slow = slow.Next
    }
    slow.Next = slow.Next.Next
    return root.Next
}
