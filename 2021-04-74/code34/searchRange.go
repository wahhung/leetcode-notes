// 34. 在排序数组中查找元素的第一个和最后一个位置
// 解题方法：二分查找
// 解题思路：
//      1: 利用二分查找找到第一个 target ，如果不存在则结束查找
//      2: 利用二分查找到第一个 target+1，返回的下标 idx
//        2-1: 如果 nums[idx] == target，则表示没有比target大的
//        2-2: 如果 nums[idx] > target，则最后一个target位置为 idx-1
// Time:O(logn) Space:O(1)
func searchRange(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    }
    left := search(nums, target, 0, len(nums)-1)
    if nums[left] != target {
        return []int{-1, -1}
    }
    if right := search(nums, target+1, left, len(nums)-1); nums[right] == target {
        return []int{left, right}
    } else {
        return []int{left, right-1}
    }    
}

func search(nums []int, target, left, right int) int {
    for left < right {
        mid := (left+right) / 2
        if target <= nums[mid] {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}
