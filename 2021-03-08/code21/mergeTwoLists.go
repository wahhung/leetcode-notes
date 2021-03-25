//21. 合并两个有序链表
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法1：迭代
// 思路分析：
//		新建一个 root 结点，prev 结点记录当前指针位置
//		通过迭代方式，不断取 min(l1.Val, l2.Val) 的结点连接到 prev 后面，
//		直到某一个链表全部取完，另一个链表则将余下部分连接到 prev 后面
// Time: O(m+n) m、n分别为l1、l2的长度
// Space: O(1) 
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{Val: -1}
	prev := root
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next, l1 = l1, l1.Next
		} else {
			prev.Next, l2 = l2, l2.Next
		}
		prev = prev.Next
	}
	if l1 == nil {
		prev.Next = l2
	} else {
		prev.Next = l1
	}
	return root.Next
}

// 方法2：递归
// 思路分析：
//		如果 l1 或 l2 为空，证明递归结束
//		比较 l1 和 l2 谁的 Val更小，然后将递归添加到该结点的 Next
// Time: O(m+n) m、n分别为l1、l2的长度
// Space: O(m+n) 递归的栈空间
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}
