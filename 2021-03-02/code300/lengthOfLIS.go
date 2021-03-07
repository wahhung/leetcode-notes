// 300. 最长递增子序列
package main

import "fmt"

// 严格递增 a < b

// 方法一
// 状态定义：dp[i]是前i个数字的最长子序列长度
// 转移方程：dp[i] = max(dp[i], dp[j]+1)
// 时间复杂度：O(N^2)
// 空间复杂度O：(N)
func lengthOfLIS1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	var dpMax int
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] { // 严格递增 <；非严格递增 <=
				// dp[j] 为 从0到j最长自增字串
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		dpMax = max(dpMax, dp[i])
	}

	// 加上在该位置的数字本身
	return dpMax + 1
}

// 方法二
// 状态定义：tails[count-1] 为长度为count的子序列的最后一个元素
// 转移方程：
//		二分法插入，查找到 nums[i] 元素在 tails[0]-tails[count] 中应该存放的位置
//		如果是插入的位置是count，即最后一个位置，则count++
// 取最长，从0开始
//
// 时间复杂度：O(nlogn) 动态规划+二分查找，二分查找时间复杂度O(logn)
// 空间复杂度： O(n)
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	tails := make([]int, n)
	count := 0
	for _, num := range nums {
		i, j := 0, count
		// 二分查找，找到num该放入到tails的位置
		for i < j {
			mid := (i + j) / 2
			if tails[mid] < num {
				i = mid + 1
			} else {
				j = mid
			}
		}
		tails[i] = num  // i == j
		// 如果 j 没有向前移动，证明num落在了count(最后一个)位置上，则最长字串长度+1
		if count == j {
			count++
		}
	}

	return count
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}
