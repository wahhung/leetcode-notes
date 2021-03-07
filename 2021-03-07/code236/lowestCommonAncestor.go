// 236. 二叉树的最近公共祖先
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 循环终值条件：
//		当前节点为 nil (root==nil)，返回 null
//      当前节点等于p或q，返回 当前节点(root==p || root==q)
// 时间复杂度： O(n)
// 空间复杂度： O(n) 递归调用的栈深度取决二叉树的高度，二叉树最坏的情况为一条链，高度为n
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil &&  right == nil {
		return nil
	}
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
