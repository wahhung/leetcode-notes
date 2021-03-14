// 122. 买卖股票的最佳时机 II
package main

import "fmt"

// 贪心算法：每一步都做出当时看起来最佳的选择(局部最优)
// 解题思路：只要今天比昨天高，就可以交易
// Time:O(n)
// Space:O(1)
func maxProfit(prices []int) int {
	var profit int
	for i := 1; i < len(prices); i++ {
		profit += max(0, prices[i]-prices[i-1])
	}
	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main()  {
	nums := []int{7,1,5,3,6,4}
	fmt.Println(maxProfit(nums))
}