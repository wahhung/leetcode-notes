// 16. 最接近的三数之和
package main

import (
	"fmt"
	"math"
)

// 先排序，后双指针
func threeSumClosest(nums []int, target int) int {
	sort(nums, 0, len(nums)-1)
	best := math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			total := nums[i] + nums[left] + nums[right]
			fmt.Println(left, right, total)
			if total == target {
				return total
			}
			best = chooseBest(best, total, target)
			if total < target {
				left++
			} else {
				right--
			}
		}
	}
	return best
}

func chooseBest(a, b, target int) int {
	if abs(a-target) < abs(b-target) {
		return a
	}
	return b
}

func sort(nums []int, start, end int) {
	if start >= end {
		return
	}
	left, right := start, end
	target := nums[end]
	for left < right {
		for nums[left] <= target && left < right {
			left++
		}
		for nums[right] >= target && left < right {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	nums[left], nums[end] = target, nums[left]
	sort(nums, start, left-1)
	sort(nums, left+1, end)
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func main() {
	arr := []int{-1, 2, 1, -4}
	fmt.Println(threeSumClosest(arr, 1))
}
