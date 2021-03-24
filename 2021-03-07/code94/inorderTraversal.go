// 94. 二叉树的中序遍历
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历 左->根->又

// 递归版本
// 时间复杂度：O(n) 每个元素仅被访问一次
// 空间复杂度: O(n)，递归过程中栈的开销，平均情况为 O(logn), 最坏情况树呈连状，为O(n)
// nums 不算额外空间
func inorderTraversal(root *TreeNode) []int {
	nums := make([]int, 0)
	inorder(root,&nums)
	return nums
}

func inorder(root *TreeNode, nums *[]int)  {
	if root == nil {
		return
	}
	inorder(root.Left, nums)
	*nums = append(*nums, root.Val)
	inorder(root.Right, nums)
}

// 非递归版本
// 时间复杂度：O(n) 每个元素仅被访问一次
// 空间复杂度: O(n)，递归过程中栈的开销，平均情况为 O(logn), 最坏情况树呈连状，为O(n)
func inorderTraversal1(root *TreeNode) []int {
	nums := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nums = append(nums, root.Val)
		root = root.Right
	}
	return nums
}
