// 70. 爬楼梯
package main

// 方法一
// 状态定义：dp[i] 为 i-1 级楼梯有多少种方法
// 转移方程： dp[i] = dp[i-1] + dp[i-2]
// 动态规划：时间复杂度 O(n) 空间复杂度 O(n)
func climbStairs1(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 方法二
// 状态定义：cur i级楼梯当前有多少中方法，pre1 i-1 级楼梯有多少，pre2 i-2 有多少
// 转移方程：cur = pre1 + pre2
func climbStairs(n int) int {
	pre1, pre2, cur := 1, 1, 1
	for i := 2; i <= n; i++ {
		cur = pre1 + pre2
		pre1, pre2 = cur, pre1
	}
	return cur
}
