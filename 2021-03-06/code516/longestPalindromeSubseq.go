// 516. 最长回文子序列
package main

import "fmt"

// 状态定义：dp[n-i][j] 是子串s[j:i]的最长回文子序列
// 转移方程：
// 		事实上，相当于正方形取最优路径，所以j(移动)从后往前推演
// 		s[i]==s[j] dp[i][j] = dp[i-1][j+1] + 2 相当于在 s[j+1,...,i-1] -> s[j,j+1,...,i-1,i]
//      s[i]!=s[j] dp[i][j] = max(dp[i-1][j], dp[i][j+1]) s[j,...,i-1] 或 s[j+1,..,i]的最大值
// 时间复杂度 O(n^2) 空间复杂度 O(n^2)
func longestPalindromeSubseq1(s string) int {
	n := len(s)
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1 // 开始的地方，i==j,记1
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] {
				dp[i][j] = dp[i-1][j+1] + 2
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j+1])
			}
		}
		fmt.Println(dp[i])
	}
	// 结果放到最后一维的数组，所以n-1
	return dp[n-1][0]
}

// 该题可以
// 状态定义：
//		lastMax 计算s[j,i]时， s[j+1,...,i-1]的最大回文子序列
//      rightMax 计算s[j,i]时， s[j+1,...,i] 的最大回文子序列
// 转移方程：
//      事实上，相当于正方形取最优路径，所以j(移动)从后往前推演，
// 		s[i]==s[j] dp[i][j] = dp[i-1][j+1] + 2 相当于在 s[j+1,...,i-1] -> s[j,j+1,...,i-1,i]
//      s[i]!=s[j] dp[i][j] = max(dp[i-1][j], dp[i][j+1]) s[j,...,i-1] 或 s[j+1,..,i]的最大值
//      但是和正方形的区别在于只会计算一半的数据，为了降低空间复杂度，可以将dp压缩至一维数组
//      s[i]=s[j]  dp[j] = lastMax + 2
//      s[i]!=s[j] dp[j] = max(d[j], rightMax)
// 时间复杂度 O(n^2) 空间复杂度 O(n)
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		lastMax := 0
		curLastMax := 0
		for j := i - 1; j >= 0; j-- {
			curLastMax = dp[j] // 记录当前位置的值，相当于下一次的 d[i-1][j+1]
			if s[i] == s[j] {
				dp[j] = lastMax + 2 // 相当于 dp[i-1][j+1] + 2
			} else {
				dp[j] = max(curLastMax, dp[j+1]) // 相当于 max(dp[i-1][j], dp[i][j+1])
			}
			lastMax = curLastMax
		}
	}
	maxSum := 0
	for i := 0; i < n; i++ {
		maxSum = max(maxSum, dp[i])
	}
	return maxSum
}

// TODO 中心拓展法
// TODO 马拉车法

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	//fmt.Println(longestPalindromeSubseq("cbaabc"))
	fmt.Println(longestPalindromeSubseq("bbbab"))
}
