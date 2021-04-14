package main

import (
	"math/rand"
	"time"
)

func main() {

}

// 数组中的第 K 个最大元素

// 基于堆排序的选择方法
// 堆是一棵完全二叉树，
// 最大堆：根节点的键值是所有堆结点键值中最大值，且每个结点的值都比其孩子的值大
// 最小堆：根结点的键值是所有堆结点键值中最小值，且每个结点的值都比其孩子的值小
// 解题思路：利用堆排序，进排序后面的 n-k个，即剩下的为k最大堆
// 	1:调整为最大堆，nums[i] > max(nums[2*i+1], nums[2*i+2])
// 	2:对(k, n) 的数组进行堆排序
// 	3:剩余部分再调整为最大堆，nums[0]即为第k个最大数
// Time:O(nlogn)
// Space:O(logn) 栈
func findKthLargest(nums []int, k int) int {
    // 父子结点的关系，只需要遍历一半即可,i~len(nums)
	for i := len(nums)/2 - 1; i >= 0; i-- {
		adjust(nums, i, len(nums))
	}
	for i := len(nums) - 1; i > len(nums) - k; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		adjust(nums, 0, i)
	}
    return nums[0]
}

// 最大堆
func adjust(nums []int, start, end int) {
	target := nums[start]
	for i := start * 2 + 1; i < end; i = i * 2 + 1 {
		if i + 1 < end && nums[i] < nums[i+1] {
			i++
		}
		if target < nums[i] {
			nums[start] = nums[i]
			start = i
		} else {
			break
		}
	}
	nums[start] = target
}

// 基于快速排序的选择方法
// Time:O(n)
// Space:O(logn)
func findKthLargest1(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(a []int, l, r, index int) int {
	q := randomPartition(a, l, r)
	if q == index {
		return a[q]
	} else if q < index {
		return quickSelect(a, q+1, r, index)
	}
	return quickSelect(a, l, q-1, index)
}

func randomPartition(a []int, l, r int) int {
	i := rand.Int()%(r-l+1) + 1
	a[i], a[r] = a[r], a[i]
	return partition(a, l, r)
}

func partition(a []int, l, r int) int {
	x := a[r]
	i := l - 1
	for j := l; j < r; j++ {
		if a[j] <= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}
