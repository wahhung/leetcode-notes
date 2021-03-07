// 120. 三角形最小路径和

// 解题思路：从后往前推，b列中(a, b) 和 (a+1, b) 之间最小的和 (a,b-1) 之和为b-1列最小，直至到 (0,0) 位置为最小路径
package main

import "fmt"

// 状态定义：dp[i][j] 为 (n-i, m-j) 位置的最佳路线，dp[0][0]为最优
// 转移方程：dp[i][j] = triangle[i][j] + min(dp[i+1][j], dp[i+1][j+1])
// 取最优，从后往前推演
// 时间复杂度 O(n^2)
// 空间复杂度 O(n^2)，如果需要更改 triangle 的值则是 O(1)
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 || len(triangle[n-1]) == 0 {
		return 0
	}
	dp := make([][]int, n)
	dp[n-1] = append(dp[n-1], triangle[n-1]...)
	for i := n - 2; i >= 0; i-- {
		dp[i] = make([]int, n)
		for j := 0; j < i+1; j++ {
			dp[i][j] = triangle[i][j] + min(dp[i+1][j], dp[i+1][j+1])
		}
	}
	return dp[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(triangle))
}
