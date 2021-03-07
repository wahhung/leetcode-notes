//64. 最小路径和
package main

import "fmt"

// 状态定义：grid[i][j] 为走到格子 (n-i,m-j) 的最优路径
// 转移方程：grid[i][j] = grid[i][j] + min(grid[i+1][j], grid[i][j+1])，找最优从后往前推演
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if i+1 == len(grid) && j+1 == len(grid[i]) {
				continue
			}
			if i+1 >= len(grid) {
				grid[i][j] += grid[i][j+1]
			} else if j+1 >= len(grid[i]) {
				grid[i][j] += grid[i+1][j]
			} else {
				grid[i][j] += min(grid[i][j+1], grid[i+1][j])
			}
		}
	}
	return grid[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(grid))
}
