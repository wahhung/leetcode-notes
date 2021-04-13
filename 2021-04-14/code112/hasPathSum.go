package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//112. 路径总和
// 解题思路：广度优先
// 1：nodeList 保存结点，使用时删除， 利用层遍历将node append进去
// 2: valList 保存结点对于的值
// 3：循环终止 len(nodeList) == 0 || (node.Left == nil && node.Right == nil && valList[0] == targetSum
// Time:O(n), Space:O(n)
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	nodeList := []*TreeNode{root}
	valList := []int{root.Val}
	for len(nodeList) > 0 {
		node := nodeList[0]
		nodeList = nodeList[1:]
		val := valList[0]
		valList = valList[1:]
		if node.Left == nil && node.Right == nil {
			if val == targetSum {
				return true
			}
			continue
		}
		if node.Left != nil {
			nodeList = append(nodeList, node.Left)
			valList = append(valList, node.Left.Val + val)
		}
		if node.Right != nil {
			nodeList = append(nodeList, node.Right)
			valList = append(valList, node.Right.Val + val)
		}
	}
	return false
}
