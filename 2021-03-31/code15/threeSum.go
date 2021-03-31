// 15. 三数之和
package main

import (
	"fmt"
	"sort"
)

// 先排序，后双指针
func threeSum(nums []int) [][]int {
	target := make([][]int, 0)
	if len(nums) < 3 {
		return target
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			total := nums[i] + nums[left] + nums[right]
			if total == 0 {
				n := len(target) - 1
				if n < 0 || target[n][1] != nums[left] || target[n][2] != nums[right] {
					// 保证上一轮数字和本次不同
					target = append(target, []int{nums[i], nums[left], nums[right]})
				}
				left++
				right--
			} else if total < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return target
}

func main() {
	arr := []int{-2, 0, 0, 2, 2}
	fmt.Println(threeSum(arr))
}
