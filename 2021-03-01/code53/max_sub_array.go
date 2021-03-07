// 53. 最大子序和

// 动态规划
// 解题思路：max 表示到目前为止，最大的集合和，pre为上一个元素至之前最大和，for循环，判断当前值nums[i]，和 pre + nums[i] 取最大，即为本次最大，记为pre，比较pre和max的大小，记为max
package main

import "fmt"
// 状态定义： 到i 位置，maxSum 表示目前最大连续子数组和，pre表示已i为最后一个数往前相加最大和
// 转移方程：
// 		 pre = max(num[i], pre + nums[i])
// 		 maxSum = max(maxSum, pre)
// 找最长，从0开始推演
// 时间复杂度 O(n)
// 空间复杂度O(1)
func maxSubArray(nums []int) int {
	pre, maxSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		pre = max(nums[i], pre+nums[i])
		maxSum = max(maxSum, pre)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
