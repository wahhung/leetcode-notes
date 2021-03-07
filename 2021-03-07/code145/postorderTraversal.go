// 145. 二叉树的后序遍历
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 后序遍历 左->右->根

// 递归版本
// 时间复杂度：O(n)
// 空间复杂度: O(n)，递归过程中栈的开销，平均情况为 O(logn), 最坏情况树呈连状，为O(n)
func postorderTraversal(root *TreeNode) []int {
	nums := make([]int, 0)
	postorder(root,&nums)
	return nums
}

func postorder(root *TreeNode, nums *[]int)  {
	if root == nil {
		return
	}
	postorder(root.Left, nums)
	postorder(root.Right, nums)
	*nums = append(*nums, root.Val)
}

// 非递归版本
// 时间复杂度：O(n)
// 空间复杂度: O(n)，递归过程中栈的开销，平均情况为 O(logn), 最坏情况树呈连状，为O(n)
func postorderTraversal1(root *TreeNode) []int {
	nums := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var prev *TreeNode // 记录right 是否遍历过了
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		if root.Right == nil || root.Right == prev {
			nums = append(nums, root.Val)
			stack = stack[:len(stack)-1]
			prev = root
			root = nil
		} else {
			root = root.Right
		}
	}
	return nums
}

