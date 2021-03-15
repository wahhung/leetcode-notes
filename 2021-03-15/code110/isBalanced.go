// 110. 平衡二叉树
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 平衡二叉树：二叉树的每个结点的左右子树的高度差不超过1
// 解题思路：利用前序遍历，先判断迭代找到最左子叶结点，再判断其父结点的左右结点是否平衡，然后往上迭代
// Time: O(n^2)，高度为d的树height(p)调用d次，所以时间复杂度为O(nlogn),最坏情况下，二叉树形成链结构，为O(n^2)
// Space:O(n) 递归调用的栈
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if abs(height(root.Left) - height(root.Right)) > 1 {
		return false
	}
	if !isBalanced(root.Left) {
		return false
	}
	return isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}