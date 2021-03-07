// 746. 使用最小花费爬楼梯 。
package main

import "fmt"

// 状态定义：dp[i] 为 走到 n-i 步最有消耗体力
// 转移方程：dp[i] = cost[i] + min(dp[i+1], dp[i+2])，最优选择，从后往前推演
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n == 0 {
		return n
	}
	dp := cost[:]
	for i := n - 3; i >= 0; i-- {
		dp[i] = cost[i] + min(dp[i+1], dp[i+2])
	}
	return min(dp[0], dp[1])
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(cost))
}
