// 144. 二叉树的前序遍历
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历 根->左->右

// 递归版本
// 时间复杂度：O(n)
// 空间复杂度: O(n)，递归过程中栈的开销，平均情况为 O(logn), 最坏情况树呈连状，为O(n)
func preorderTraversal(root *TreeNode) []int {
	nums := make([]int, 0)
	preorder(root,&nums)
	return nums
}

func preorder(root *TreeNode, nums *[]int)  {
	if root == nil {
		return
	}
	*nums = append(*nums, root.Val)
	preorder(root.Left, nums)
	preorder(root.Right, nums)
}

// 非递归版本
// 时间复杂度：O(n)
// 空间复杂度: O(n)，递归过程中栈的开销，平均情况为 O(logn), 最坏情况树呈连状，为O(n)
func preorderTraversal1(root *TreeNode) []int {
	nums := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			nums = append(nums, root.Val)
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return nums
}