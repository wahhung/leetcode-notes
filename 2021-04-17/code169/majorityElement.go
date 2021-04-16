// 169. 多数元素
// 投票法
// 1. 因为本题中众数个数 > n / 2，利用投票法
// 2. 当count==0时，记录元素 target = num
// 3. 下一个如果等于target，count++，否则 count--，
// 4. count == 0 重复2
// Time：O(n) Space:O(1)
func majorityElement(nums []int) int {
    var count int
    var target int
    for _, num := range nums {
        if count == 0 {
            target = num
        }
        if target == num {
            count++
        } else {
            count--
        }
    }
    return target
}
