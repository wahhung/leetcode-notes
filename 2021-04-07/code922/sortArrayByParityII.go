// 922. 按奇偶排序数组 II
// 解题方法：双指针
// 解题思路：
//      1: 定于下标为i的保存偶数，j的保存基数
//      2: 如果 nums[i]%2 != 0,即需要第一个nums[j]%2==0与之替换
// Time:O(n), Space:O(1)
func sortArrayByParityII(nums []int) []int {
    for i, j := 0, 1; i < len(nums); i += 2 {
        if nums[i] % 2 == 0 {
            continue
        }
        for nums[j] % 2 == 1 {
            j += 2
        }
        nums[i], nums[j] = nums[j], nums[i]
    }
    return nums
}
