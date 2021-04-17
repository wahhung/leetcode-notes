// 剑指 Offer 48. 最长不含重复字符的子字符串
// 双指针 + 动态规划
// 1.定义 
//      哈希表data,存储每个字符最后一次出现的位置，
//      i 指针data出现相同字符的最后一个位置 i = max(i, data[s[j]])
//      j当前位置
//      res 结果最大长度
// 2.循环，出现相同字符串，更新 i，每次循环，res = max(res, j - i)
// Time:O(n) Space:O(1) 字符的 ASCII 是 0～127 是有限的
func lengthOfLongestSubstring(s string) int {
    data := map[byte]int{}
    i := -1
    var res int
    for j := 0; j < len(s); j++ {
        if val, ok := data[s[j]]; ok {
            i = max(i, val)
        }
        data[s[j]] = j
        res = max(res, j - i)
    }
    return res
}

func max(a, b int) int {
    if a >= b {
        return a
    }
    return b
}
