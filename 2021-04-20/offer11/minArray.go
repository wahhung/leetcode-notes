// 剑指 Offer 11. 旋转数组的最小数字
// 二分查找
// 1.定义left->0, right->len(numbers)-1，循环继续条件 left < right
// 2.取mid
//  2-1.如果nums[mid] < nums[right]，则说明后半段有序，target在前半段，right=mid
//  2-2.如果nums[mid] > nums[right]，说明目标在后半段，且不是nums[mid], left=mid-1
//  2-3.如果nums[mid] == nums[right]，不足以判断target在哪一段，right--
// Time:O(logn) Space:O(1)
func minArray(numbers []int) int {
    if len(numbers) == 0 {
        return -1
    }
    left, right := 0, len(numbers) - 1
    for left < right {
        mid := (left + right) / 2
        if numbers[mid] < numbers[right] {
            right = mid
        } else if numbers[mid] > numbers[right] {
            left = mid + 1
        } else {
            right--
        }
    }
    return numbers[left]
}

// 二分查找的 left= mid or mid+1，需要判断mid的结构是不是被剔除
