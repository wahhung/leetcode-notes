// 23. 合并K个升序链表
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 顺序合并
// 解题思路：将N个链表转换为多次合并两个链表
// Time:O(k^2 * n), Space:O(1)
func mergeKLists1(lists []*ListNode) *ListNode {
	var head *ListNode
	for i := 0; i < len(lists); i++ {
		head = mergeTwoList(head, lists[i])
	}
	return head
}

func mergeTwoList(a, b *ListNode) *ListNode {
	root := &ListNode{}
	cur := root
	for a != nil && b != nil {
		if a.Val <= b.Val {
			cur.Next = a
			a = a.Next
		} else {
			cur.Next = b
			b = b.Next
		}
		cur = cur.Next
	}
	if a != nil {
		cur.Next = a
	}
	if b != nil {
		cur.Next = b
	}
	return root.Next
}

// 分治合并
// 将lists分开[0,1][2,3]...[n-2,n-1]的两个链表，递归合并
// Time:O(kn * logk) Space:O(logk)
func mergeKLists2(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, left, right int) *ListNode {
	if left > right {
		return nil
	}
	if left == right {
		return lists[left]
	}
	mid := (left + right) / 2
	l1 := merge(lists, left, mid)
	l2 := merge(lists, mid+1, right)
	return mergeTwoList(l1, l2)
}

// 小顶堆
// 维护数组的链表的首个结点成小顶堆
// 每次从堆顶取出元素，堆顶的链表换成第二个元素，然后再调整小顶堆
// 直到堆为空
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 顺序合并
// 解题思路：将N个链表转换为多次合并两个链表
// Time:O(k^2 * n), Space:O(1)
func mergeKLists1(lists []*ListNode) *ListNode {
	var head *ListNode
	for i := 0; i < len(lists); i++ {
		head = mergeTwoList(head, lists[i])
	}
	return head
}

func mergeTwoList(a, b *ListNode) *ListNode {
	root := &ListNode{}
	cur := root
	for a != nil && b != nil {
		if a.Val <= b.Val {
			cur.Next = a
			a = a.Next
		} else {
			cur.Next = b
			b = b.Next
		}
		cur = cur.Next
	}
	if a != nil {
		cur.Next = a
	}
	if b != nil {
		cur.Next = b
	}
	return root.Next
}

// 分治合并
// 将lists分开[0,1][2,3]...[n-2,n-1]的两个链表，递归合并
// Time:O(kn * logk) Space:O(logk)
func mergeKLists2(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, left, right int) *ListNode {
	if left > right {
		return nil
	}
	if left == right {
		return lists[left]
	}
	mid := (left + right) / 2
	l1 := merge(lists, left, mid)
	l2 := merge(lists, mid+1, right)
	return mergeTwoList(l1, l2)
}

// 小顶堆
// 维护数组的链表的首个结点成小顶堆
// 每次从堆顶取出元素，堆顶的链表换成第二个元素，然后再调整小顶堆
// 直到堆为空
